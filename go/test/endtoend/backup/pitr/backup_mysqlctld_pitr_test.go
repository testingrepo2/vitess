/*
Copyright 2022 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mysqlctld

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"vitess.io/vitess/go/mysql"
	backup "vitess.io/vitess/go/test/endtoend/backup/vtctlbackup"
	"vitess.io/vitess/go/test/endtoend/cluster"
)

func waitForReplica(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pMsgs := backup.ReadRowsFromPrimary(t)
	for {
		rMsgs := backup.ReadRowsFromReplica(t)
		if len(pMsgs) == len(rMsgs) {
			// success
			return
		}
		select {
		case <-ctx.Done():
			assert.FailNow(t, "timeout waiting for replica to catch up")
			return
		case <-time.After(time.Second):
			//
		}
	}
}

// TestIncrementalBackupMysqlctld - tests incremental backups using myslctld
func TestIncrementalBackupMysqlctld(t *testing.T) {
	defer cluster.PanicHandler(t)
	// setup cluster for the testing
	code, err := backup.LaunchCluster(backup.Mysqlctld, "xbstream", 0, nil)
	require.NoError(t, err, "setup failed with status code %d", code)
	defer backup.TearDownCluster()

	backup.InitTestTable(t)

	rowsPerPosition := map[string]int{}
	backupPositions := []string{}

	recordRowsPerPosition := func(t *testing.T) {
		pos := backup.GetReplicaPosition(t)
		msgs := backup.ReadRowsFromReplica(t)
		if _, ok := rowsPerPosition[pos]; !ok {
			backupPositions = append(backupPositions, pos)
			rowsPerPosition[pos] = len(msgs)
		}
	}

	var fullBackupPos mysql.Position
	t.Run("full backup", func(t *testing.T) {
		backup.InsertRowOnPrimary(t, "before-full-backup")
		waitForReplica(t)
		manifest, _ := backup.TestReplicaFullBackup(t)
		fullBackupPos = manifest.Position
		require.False(t, fullBackupPos.IsZero())
		//
		msgs := backup.ReadRowsFromReplica(t)
		pos := mysql.EncodePosition(fullBackupPos)
		backupPositions = append(backupPositions, pos)
		rowsPerPosition[pos] = len(msgs)
	})

	lastBackupPos := fullBackupPos
	backup.InsertRowOnPrimary(t, "before-incremental-backups")

	tt := []struct {
		name              string
		writeBeforeBackup bool
		fromFullPosition  bool
		autoPosition      bool
		expectError       string
	}{
		{
			name: "first incremental backup",
		},
		{
			name:              "make writes, succeed",
			writeBeforeBackup: true,
		},
		{
			name:        "fail, no binary logs to backup",
			expectError: "no binary logs to backup",
		},
		{
			name:              "make writes again, succeed",
			writeBeforeBackup: true,
		},
		{
			name:              "auto position, succeed",
			writeBeforeBackup: true,
			autoPosition:      true,
		},
		{
			name:         "fail auto position, no binary logs to backup",
			autoPosition: true,
			expectError:  "no binary logs to backup",
		},
		{
			name:              "auto position, make writes again, succeed",
			writeBeforeBackup: true,
			autoPosition:      true,
		},
		{
			name:             "from full backup position",
			fromFullPosition: true,
		},
	}
	var fromFullPositionBackups []string
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.writeBeforeBackup {
				backup.InsertRowOnPrimary(t, "")
			}
			// we wait for 1 second because backups are written to a directory named after the current timestamp,
			// in 1 second resolution. We want to avoid two backups that have the same pathname. Realistically this
			// is only ever a problem in this end-to-end test, not in production.
			// Also, we gie the replica a chance to catch up.
			time.Sleep(1100 * time.Millisecond)
			waitForReplica(t)
			recordRowsPerPosition(t)
			// configure --incremental-from-pos to either:
			// - auto
			// - explicit last backup pos
			// - back in history to the original full backup
			var incrementalFromPos mysql.Position
			if !tc.autoPosition {
				incrementalFromPos = lastBackupPos
				if tc.fromFullPosition {
					incrementalFromPos = fullBackupPos
				}
			}
			manifest, backupName := backup.TestReplicaIncrementalBackup(t, incrementalFromPos, tc.expectError)
			if tc.expectError != "" {
				return
			}
			defer func() {
				lastBackupPos = manifest.Position
			}()
			if tc.fromFullPosition {
				fromFullPositionBackups = append(fromFullPositionBackups, backupName)
			}
			require.False(t, manifest.FromPosition.IsZero())
			require.NotEqual(t, manifest.Position, manifest.FromPosition)
			require.True(t, manifest.Position.GTIDSet.Contains(manifest.FromPosition.GTIDSet))

			gtidPurgedPos, err := mysql.ParsePosition(mysql.Mysql56FlavorID, backup.GetReplicaGtidPurged(t))
			require.NoError(t, err)
			fromPositionIncludingPurged := manifest.FromPosition.GTIDSet.Union(gtidPurgedPos.GTIDSet)

			expectFromPosition := lastBackupPos.GTIDSet.Union(gtidPurgedPos.GTIDSet)
			if !incrementalFromPos.IsZero() {
				expectFromPosition = incrementalFromPos.GTIDSet.Union(gtidPurgedPos.GTIDSet)
			}
			require.Equalf(t, expectFromPosition, fromPositionIncludingPurged, "expected: %v, found: %v", expectFromPosition, fromPositionIncludingPurged)
		})
	}

	testRestores := func(t *testing.T) {
		for _, r := range rand.Perm(len(backupPositions)) {
			pos := backupPositions[r]
			testName := fmt.Sprintf("%s, %d records", pos, rowsPerPosition[pos])
			t.Run(testName, func(t *testing.T) {
				restoreToPos, err := mysql.DecodePosition(pos)
				require.NoError(t, err)
				backup.TestReplicaRestoreToPos(t, restoreToPos, "")
				msgs := backup.ReadRowsFromReplica(t)
				count, ok := rowsPerPosition[pos]
				require.True(t, ok)
				assert.Equalf(t, count, len(msgs), "messages: %v", msgs)
			})
		}
	}
	t.Run("PITR", func(t *testing.T) {
		testRestores(t)
	})
	t.Run("remove full position backups", func(t *testing.T) {
		// Delete the fromFullPosition backup(s), which leaves us with less restore options. Try again.
		for _, backupName := range fromFullPositionBackups {
			backup.RemoveBackup(t, backupName)
		}
	})
	t.Run("PITR-2", func(t *testing.T) {
		testRestores(t)
	})
}
