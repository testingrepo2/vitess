/*
Copyright 2019 The Vitess Authors.

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

package tabletmanager

import (
	"fmt"
	"regexp"
	"time"

	"vitess.io/vitess/go/vt/logutil"
	"vitess.io/vitess/go/vt/vterrors"

	"golang.org/x/net/context"

	"vitess.io/vitess/go/vt/hook"
	"vitess.io/vitess/go/vt/mysqlctl"
	"vitess.io/vitess/go/vt/topotools"

	tabletmanagerdatapb "vitess.io/vitess/go/vt/proto/tabletmanagerdata"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
)

// This file contains the implementations of RPCAgent methods.
// Major groups of methods are broken out into files named "rpc_*.go".

// Ping makes sure RPCs work, and refreshes the tablet record.
func (agent *TabletManager) Ping(ctx context.Context, args string) string {
	return args
}

// GetPermissions returns the db permissions.
func (agent *TabletManager) GetPermissions(ctx context.Context) (*tabletmanagerdatapb.Permissions, error) {
	return mysqlctl.GetPermissions(agent.MysqlDaemon)
}

// SetReadOnly makes the mysql instance read-only or read-write.
func (agent *TabletManager) SetReadOnly(ctx context.Context, rdonly bool) error {
	if err := agent.lock(ctx); err != nil {
		return err
	}
	defer agent.unlock()

	return agent.MysqlDaemon.SetReadOnly(rdonly)
}

// ChangeType changes the tablet type
func (agent *TabletManager) ChangeType(ctx context.Context, tabletType topodatapb.TabletType) error {
	if err := agent.lock(ctx); err != nil {
		return err
	}
	defer agent.unlock()
	return agent.changeTypeLocked(ctx, tabletType)
}

// ChangeType changes the tablet type
func (agent *TabletManager) changeTypeLocked(ctx context.Context, tabletType topodatapb.TabletType) error {
	// We don't want to allow multiple callers to claim a tablet as drained. There is a race that could happen during
	// horizontal resharding where two vtworkers will try to DRAIN the same tablet. This check prevents that race from
	// causing errors.
	if tabletType == topodatapb.TabletType_DRAINED && agent.Tablet().Type == topodatapb.TabletType_DRAINED {
		return fmt.Errorf("Tablet: %v, is already drained", agent.TabletAlias)
	}

	tablet := agent.Tablet()
	tablet.Type = tabletType
	// If we have been told we're master, set master term start time to Now
	// and save it topo immediately.
	if tabletType == topodatapb.TabletType_MASTER {
		tablet.MasterTermStartTime = logutil.TimeToProto(time.Now())

		// change our type in the topology, and set masterTermStartTime on tablet record if applicable
		_, err := topotools.ChangeType(ctx, agent.TopoServer, agent.TabletAlias, tabletType, tablet.MasterTermStartTime)
		if err != nil {
			return err
		}
	} else {
		tablet.MasterTermStartTime = nil
	}

	// updateState will invoke broadcastHealth if needed.
	agent.updateState(ctx, tablet, "ChangeType")

	// Let's see if we need to fix semi-sync acking.
	if err := agent.fixSemiSyncAndReplication(agent.Tablet().Type); err != nil {
		return vterrors.Wrap(err, "fixSemiSyncAndReplication failed, may not ack correctly")
	}
	return nil
}

// Sleep sleeps for the duration
func (agent *TabletManager) Sleep(ctx context.Context, duration time.Duration) {
	if err := agent.lock(ctx); err != nil {
		// client gave up
		return
	}
	defer agent.unlock()

	time.Sleep(duration)
}

// ExecuteHook executes the provided hook locally, and returns the result.
func (agent *TabletManager) ExecuteHook(ctx context.Context, hk *hook.Hook) *hook.HookResult {
	if err := agent.lock(ctx); err != nil {
		// client gave up
		return &hook.HookResult{}
	}
	defer agent.unlock()

	// Execute the hooks
	topotools.ConfigureTabletHook(hk, agent.TabletAlias)
	return hk.Execute()
}

// RefreshState reload the tablet record from the topo server.
func (agent *TabletManager) RefreshState(ctx context.Context) error {
	if err := agent.lock(ctx); err != nil {
		return err
	}
	defer agent.unlock()

	return agent.refreshTablet(ctx, "RefreshState")
}

// RunHealthCheck will manually run the health check on the tablet.
func (agent *TabletManager) RunHealthCheck(ctx context.Context) {
	agent.runHealthCheck()
}

// IgnoreHealthError sets the regexp for health check errors to ignore.
func (agent *TabletManager) IgnoreHealthError(ctx context.Context, pattern string) error {
	var expr *regexp.Regexp
	if pattern != "" {
		var err error
		if expr, err = regexp.Compile(fmt.Sprintf("^%s$", pattern)); err != nil {
			return err
		}
	}
	agent.mutex.Lock()
	agent._ignoreHealthErrorExpr = expr
	agent.mutex.Unlock()
	return nil
}
