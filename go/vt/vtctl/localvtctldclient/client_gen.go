// Code generated by localvtctldclient-generator. DO NOT EDIT.

/*
Copyright 2021 The Vitess Authors.

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

package localvtctldclient

import (
	"context"

	"google.golang.org/grpc"

	"vitess.io/vitess/go/vt/vtctl/internal/grpcshim"

	vtctldatapb "vitess.io/vitess/go/vt/proto/vtctldata"
	vtctlservicepb "vitess.io/vitess/go/vt/proto/vtctlservice"
)

// AddCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) AddCellInfo(ctx context.Context, in *vtctldatapb.AddCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.AddCellInfoResponse, error) {
	return client.s.AddCellInfo(ctx, in)
}

// AddCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) AddCellsAlias(ctx context.Context, in *vtctldatapb.AddCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.AddCellsAliasResponse, error) {
	return client.s.AddCellsAlias(ctx, in)
}

// ApplyRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplyRoutingRules(ctx context.Context, in *vtctldatapb.ApplyRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyRoutingRulesResponse, error) {
	return client.s.ApplyRoutingRules(ctx, in)
}

// ApplySchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplySchema(ctx context.Context, in *vtctldatapb.ApplySchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplySchemaResponse, error) {
	return client.s.ApplySchema(ctx, in)
}

// ApplyShardRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplyShardRoutingRules(ctx context.Context, in *vtctldatapb.ApplyShardRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyShardRoutingRulesResponse, error) {
	return client.s.ApplyShardRoutingRules(ctx, in)
}

// ApplyVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplyVSchema(ctx context.Context, in *vtctldatapb.ApplyVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyVSchemaResponse, error) {
	return client.s.ApplyVSchema(ctx, in)
}

type backupStreamAdapter struct {
	*grpcshim.BidiStream
	ch chan *vtctldatapb.BackupResponse
}

func (stream *backupStreamAdapter) Recv() (*vtctldatapb.BackupResponse, error) {
	select {
	case <-stream.Context().Done():
		return nil, stream.Context().Err()
	case <-stream.Closed():
		// Stream has been closed for future sends. If there are messages that
		// have already been sent, receive them until there are no more. After
		// all sent messages have been received, Recv will return the CloseErr.
		select {
		case msg := <-stream.ch:
			return msg, nil
		default:
			return nil, stream.CloseErr()
		}
	case err := <-stream.ErrCh:
		return nil, err
	case msg := <-stream.ch:
		return msg, nil
	}
}

func (stream *backupStreamAdapter) Send(msg *vtctldatapb.BackupResponse) error {
	select {
	case <-stream.Context().Done():
		return stream.Context().Err()
	case <-stream.Closed():
		return grpcshim.ErrStreamClosed
	case stream.ch <- msg:
		return nil
	}
}

// Backup is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) Backup(ctx context.Context, in *vtctldatapb.BackupRequest, opts ...grpc.CallOption) (vtctlservicepb.Vtctld_BackupClient, error) {
	stream := &backupStreamAdapter{
		BidiStream: grpcshim.NewBidiStream(ctx),
		ch:         make(chan *vtctldatapb.BackupResponse, 1),
	}
	go func() {
		err := client.s.Backup(in, stream)
		stream.CloseWithError(err)
	}()

	return stream, nil
}

type backupShardStreamAdapter struct {
	*grpcshim.BidiStream
	ch chan *vtctldatapb.BackupResponse
}

func (stream *backupShardStreamAdapter) Recv() (*vtctldatapb.BackupResponse, error) {
	select {
	case <-stream.Context().Done():
		return nil, stream.Context().Err()
	case <-stream.Closed():
		// Stream has been closed for future sends. If there are messages that
		// have already been sent, receive them until there are no more. After
		// all sent messages have been received, Recv will return the CloseErr.
		select {
		case msg := <-stream.ch:
			return msg, nil
		default:
			return nil, stream.CloseErr()
		}
	case err := <-stream.ErrCh:
		return nil, err
	case msg := <-stream.ch:
		return msg, nil
	}
}

func (stream *backupShardStreamAdapter) Send(msg *vtctldatapb.BackupResponse) error {
	select {
	case <-stream.Context().Done():
		return stream.Context().Err()
	case <-stream.Closed():
		return grpcshim.ErrStreamClosed
	case stream.ch <- msg:
		return nil
	}
}

// BackupShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) BackupShard(ctx context.Context, in *vtctldatapb.BackupShardRequest, opts ...grpc.CallOption) (vtctlservicepb.Vtctld_BackupShardClient, error) {
	stream := &backupShardStreamAdapter{
		BidiStream: grpcshim.NewBidiStream(ctx),
		ch:         make(chan *vtctldatapb.BackupResponse, 1),
	}
	go func() {
		err := client.s.BackupShard(in, stream)
		stream.CloseWithError(err)
	}()

	return stream, nil
}

// ChangeTabletType is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ChangeTabletType(ctx context.Context, in *vtctldatapb.ChangeTabletTypeRequest, opts ...grpc.CallOption) (*vtctldatapb.ChangeTabletTypeResponse, error) {
	return client.s.ChangeTabletType(ctx, in)
}

// CleanupSchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) CleanupSchemaMigration(ctx context.Context, in *vtctldatapb.CleanupSchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.CleanupSchemaMigrationResponse, error) {
	return client.s.CleanupSchemaMigration(ctx, in)
}

// CreateKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) CreateKeyspace(ctx context.Context, in *vtctldatapb.CreateKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.CreateKeyspaceResponse, error) {
	return client.s.CreateKeyspace(ctx, in)
}

// CreateShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) CreateShard(ctx context.Context, in *vtctldatapb.CreateShardRequest, opts ...grpc.CallOption) (*vtctldatapb.CreateShardResponse, error) {
	return client.s.CreateShard(ctx, in)
}

// DeleteCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteCellInfo(ctx context.Context, in *vtctldatapb.DeleteCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteCellInfoResponse, error) {
	return client.s.DeleteCellInfo(ctx, in)
}

// DeleteCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteCellsAlias(ctx context.Context, in *vtctldatapb.DeleteCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteCellsAliasResponse, error) {
	return client.s.DeleteCellsAlias(ctx, in)
}

// DeleteKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteKeyspace(ctx context.Context, in *vtctldatapb.DeleteKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteKeyspaceResponse, error) {
	return client.s.DeleteKeyspace(ctx, in)
}

// DeleteShards is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteShards(ctx context.Context, in *vtctldatapb.DeleteShardsRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteShardsResponse, error) {
	return client.s.DeleteShards(ctx, in)
}

// DeleteSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteSrvVSchema(ctx context.Context, in *vtctldatapb.DeleteSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteSrvVSchemaResponse, error) {
	return client.s.DeleteSrvVSchema(ctx, in)
}

// DeleteTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteTablets(ctx context.Context, in *vtctldatapb.DeleteTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteTabletsResponse, error) {
	return client.s.DeleteTablets(ctx, in)
}

// EmergencyReparentShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) EmergencyReparentShard(ctx context.Context, in *vtctldatapb.EmergencyReparentShardRequest, opts ...grpc.CallOption) (*vtctldatapb.EmergencyReparentShardResponse, error) {
	return client.s.EmergencyReparentShard(ctx, in)
}

// ExecuteFetchAsApp is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ExecuteFetchAsApp(ctx context.Context, in *vtctldatapb.ExecuteFetchAsAppRequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteFetchAsAppResponse, error) {
	return client.s.ExecuteFetchAsApp(ctx, in)
}

// ExecuteFetchAsDBA is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ExecuteFetchAsDBA(ctx context.Context, in *vtctldatapb.ExecuteFetchAsDBARequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteFetchAsDBAResponse, error) {
	return client.s.ExecuteFetchAsDBA(ctx, in)
}

// ExecuteHook is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ExecuteHook(ctx context.Context, in *vtctldatapb.ExecuteHookRequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteHookResponse, error) {
	return client.s.ExecuteHook(ctx, in)
}

// FindAllShardsInKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) FindAllShardsInKeyspace(ctx context.Context, in *vtctldatapb.FindAllShardsInKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.FindAllShardsInKeyspaceResponse, error) {
	return client.s.FindAllShardsInKeyspace(ctx, in)
}

// GetBackups is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetBackups(ctx context.Context, in *vtctldatapb.GetBackupsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetBackupsResponse, error) {
	return client.s.GetBackups(ctx, in)
}

// GetCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetCellInfo(ctx context.Context, in *vtctldatapb.GetCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoResponse, error) {
	return client.s.GetCellInfo(ctx, in)
}

// GetCellInfoNames is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetCellInfoNames(ctx context.Context, in *vtctldatapb.GetCellInfoNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoNamesResponse, error) {
	return client.s.GetCellInfoNames(ctx, in)
}

// GetCellsAliases is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetCellsAliases(ctx context.Context, in *vtctldatapb.GetCellsAliasesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellsAliasesResponse, error) {
	return client.s.GetCellsAliases(ctx, in)
}

// GetFullStatus is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetFullStatus(ctx context.Context, in *vtctldatapb.GetFullStatusRequest, opts ...grpc.CallOption) (*vtctldatapb.GetFullStatusResponse, error) {
	return client.s.GetFullStatus(ctx, in)
}

// GetKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetKeyspace(ctx context.Context, in *vtctldatapb.GetKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspaceResponse, error) {
	return client.s.GetKeyspace(ctx, in)
}

// GetKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetKeyspaces(ctx context.Context, in *vtctldatapb.GetKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspacesResponse, error) {
	return client.s.GetKeyspaces(ctx, in)
}

// GetPermissions is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetPermissions(ctx context.Context, in *vtctldatapb.GetPermissionsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetPermissionsResponse, error) {
	return client.s.GetPermissions(ctx, in)
}

// GetRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetRoutingRules(ctx context.Context, in *vtctldatapb.GetRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetRoutingRulesResponse, error) {
	return client.s.GetRoutingRules(ctx, in)
}

// GetSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSchema(ctx context.Context, in *vtctldatapb.GetSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSchemaResponse, error) {
	return client.s.GetSchema(ctx, in)
}

// GetSchemaMigrations is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSchemaMigrations(ctx context.Context, in *vtctldatapb.GetSchemaMigrationsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSchemaMigrationsResponse, error) {
	return client.s.GetSchemaMigrations(ctx, in)
}

// GetShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetShard(ctx context.Context, in *vtctldatapb.GetShardRequest, opts ...grpc.CallOption) (*vtctldatapb.GetShardResponse, error) {
	return client.s.GetShard(ctx, in)
}

// GetShardRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetShardRoutingRules(ctx context.Context, in *vtctldatapb.GetShardRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetShardRoutingRulesResponse, error) {
	return client.s.GetShardRoutingRules(ctx, in)
}

// GetSrvKeyspaceNames is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvKeyspaceNames(ctx context.Context, in *vtctldatapb.GetSrvKeyspaceNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvKeyspaceNamesResponse, error) {
	return client.s.GetSrvKeyspaceNames(ctx, in)
}

// GetSrvKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvKeyspaces(ctx context.Context, in *vtctldatapb.GetSrvKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvKeyspacesResponse, error) {
	return client.s.GetSrvKeyspaces(ctx, in)
}

// GetSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvVSchema(ctx context.Context, in *vtctldatapb.GetSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemaResponse, error) {
	return client.s.GetSrvVSchema(ctx, in)
}

// GetSrvVSchemas is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvVSchemas(ctx context.Context, in *vtctldatapb.GetSrvVSchemasRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemasResponse, error) {
	return client.s.GetSrvVSchemas(ctx, in)
}

// GetTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetTablet(ctx context.Context, in *vtctldatapb.GetTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletResponse, error) {
	return client.s.GetTablet(ctx, in)
}

// GetTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetTablets(ctx context.Context, in *vtctldatapb.GetTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletsResponse, error) {
	return client.s.GetTablets(ctx, in)
}

// GetTopologyPath is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetTopologyPath(ctx context.Context, in *vtctldatapb.GetTopologyPathRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTopologyPathResponse, error) {
	return client.s.GetTopologyPath(ctx, in)
}

// GetVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetVSchema(ctx context.Context, in *vtctldatapb.GetVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetVSchemaResponse, error) {
	return client.s.GetVSchema(ctx, in)
}

// GetVersion is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetVersion(ctx context.Context, in *vtctldatapb.GetVersionRequest, opts ...grpc.CallOption) (*vtctldatapb.GetVersionResponse, error) {
	return client.s.GetVersion(ctx, in)
}

// GetWorkflows is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetWorkflows(ctx context.Context, in *vtctldatapb.GetWorkflowsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetWorkflowsResponse, error) {
	return client.s.GetWorkflows(ctx, in)
}

// InitShardPrimary is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) InitShardPrimary(ctx context.Context, in *vtctldatapb.InitShardPrimaryRequest, opts ...grpc.CallOption) (*vtctldatapb.InitShardPrimaryResponse, error) {
	return client.s.InitShardPrimary(ctx, in)
}

// MoveTablesComplete is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) MoveTablesComplete(ctx context.Context, in *vtctldatapb.MoveTablesCompleteRequest, opts ...grpc.CallOption) (*vtctldatapb.MoveTablesCompleteResponse, error) {
	return client.s.MoveTablesComplete(ctx, in)
}

// MoveTablesCreate is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) MoveTablesCreate(ctx context.Context, in *vtctldatapb.MoveTablesCreateRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowStatusResponse, error) {
	return client.s.MoveTablesCreate(ctx, in)
}

// PingTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) PingTablet(ctx context.Context, in *vtctldatapb.PingTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.PingTabletResponse, error) {
	return client.s.PingTablet(ctx, in)
}

// PlannedReparentShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) PlannedReparentShard(ctx context.Context, in *vtctldatapb.PlannedReparentShardRequest, opts ...grpc.CallOption) (*vtctldatapb.PlannedReparentShardResponse, error) {
	return client.s.PlannedReparentShard(ctx, in)
}

// RebuildKeyspaceGraph is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RebuildKeyspaceGraph(ctx context.Context, in *vtctldatapb.RebuildKeyspaceGraphRequest, opts ...grpc.CallOption) (*vtctldatapb.RebuildKeyspaceGraphResponse, error) {
	return client.s.RebuildKeyspaceGraph(ctx, in)
}

// RebuildVSchemaGraph is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RebuildVSchemaGraph(ctx context.Context, in *vtctldatapb.RebuildVSchemaGraphRequest, opts ...grpc.CallOption) (*vtctldatapb.RebuildVSchemaGraphResponse, error) {
	return client.s.RebuildVSchemaGraph(ctx, in)
}

// RefreshState is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RefreshState(ctx context.Context, in *vtctldatapb.RefreshStateRequest, opts ...grpc.CallOption) (*vtctldatapb.RefreshStateResponse, error) {
	return client.s.RefreshState(ctx, in)
}

// RefreshStateByShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RefreshStateByShard(ctx context.Context, in *vtctldatapb.RefreshStateByShardRequest, opts ...grpc.CallOption) (*vtctldatapb.RefreshStateByShardResponse, error) {
	return client.s.RefreshStateByShard(ctx, in)
}

// ReloadSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReloadSchema(ctx context.Context, in *vtctldatapb.ReloadSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaResponse, error) {
	return client.s.ReloadSchema(ctx, in)
}

// ReloadSchemaKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReloadSchemaKeyspace(ctx context.Context, in *vtctldatapb.ReloadSchemaKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaKeyspaceResponse, error) {
	return client.s.ReloadSchemaKeyspace(ctx, in)
}

// ReloadSchemaShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReloadSchemaShard(ctx context.Context, in *vtctldatapb.ReloadSchemaShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaShardResponse, error) {
	return client.s.ReloadSchemaShard(ctx, in)
}

// RemoveBackup is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RemoveBackup(ctx context.Context, in *vtctldatapb.RemoveBackupRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveBackupResponse, error) {
	return client.s.RemoveBackup(ctx, in)
}

// RemoveKeyspaceCell is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RemoveKeyspaceCell(ctx context.Context, in *vtctldatapb.RemoveKeyspaceCellRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveKeyspaceCellResponse, error) {
	return client.s.RemoveKeyspaceCell(ctx, in)
}

// RemoveShardCell is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RemoveShardCell(ctx context.Context, in *vtctldatapb.RemoveShardCellRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveShardCellResponse, error) {
	return client.s.RemoveShardCell(ctx, in)
}

// ReparentTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReparentTablet(ctx context.Context, in *vtctldatapb.ReparentTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.ReparentTabletResponse, error) {
	return client.s.ReparentTablet(ctx, in)
}

type restoreFromBackupStreamAdapter struct {
	*grpcshim.BidiStream
	ch chan *vtctldatapb.RestoreFromBackupResponse
}

func (stream *restoreFromBackupStreamAdapter) Recv() (*vtctldatapb.RestoreFromBackupResponse, error) {
	select {
	case <-stream.Context().Done():
		return nil, stream.Context().Err()
	case <-stream.Closed():
		// Stream has been closed for future sends. If there are messages that
		// have already been sent, receive them until there are no more. After
		// all sent messages have been received, Recv will return the CloseErr.
		select {
		case msg := <-stream.ch:
			return msg, nil
		default:
			return nil, stream.CloseErr()
		}
	case err := <-stream.ErrCh:
		return nil, err
	case msg := <-stream.ch:
		return msg, nil
	}
}

func (stream *restoreFromBackupStreamAdapter) Send(msg *vtctldatapb.RestoreFromBackupResponse) error {
	select {
	case <-stream.Context().Done():
		return stream.Context().Err()
	case <-stream.Closed():
		return grpcshim.ErrStreamClosed
	case stream.ch <- msg:
		return nil
	}
}

// RestoreFromBackup is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RestoreFromBackup(ctx context.Context, in *vtctldatapb.RestoreFromBackupRequest, opts ...grpc.CallOption) (vtctlservicepb.Vtctld_RestoreFromBackupClient, error) {
	stream := &restoreFromBackupStreamAdapter{
		BidiStream: grpcshim.NewBidiStream(ctx),
		ch:         make(chan *vtctldatapb.RestoreFromBackupResponse, 1),
	}
	go func() {
		err := client.s.RestoreFromBackup(in, stream)
		stream.CloseWithError(err)
	}()

	return stream, nil
}

// RetrySchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RetrySchemaMigration(ctx context.Context, in *vtctldatapb.RetrySchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.RetrySchemaMigrationResponse, error) {
	return client.s.RetrySchemaMigration(ctx, in)
}

// RunHealthCheck is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RunHealthCheck(ctx context.Context, in *vtctldatapb.RunHealthCheckRequest, opts ...grpc.CallOption) (*vtctldatapb.RunHealthCheckResponse, error) {
	return client.s.RunHealthCheck(ctx, in)
}

// SetKeyspaceDurabilityPolicy is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetKeyspaceDurabilityPolicy(ctx context.Context, in *vtctldatapb.SetKeyspaceDurabilityPolicyRequest, opts ...grpc.CallOption) (*vtctldatapb.SetKeyspaceDurabilityPolicyResponse, error) {
	return client.s.SetKeyspaceDurabilityPolicy(ctx, in)
}

// SetShardIsPrimaryServing is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetShardIsPrimaryServing(ctx context.Context, in *vtctldatapb.SetShardIsPrimaryServingRequest, opts ...grpc.CallOption) (*vtctldatapb.SetShardIsPrimaryServingResponse, error) {
	return client.s.SetShardIsPrimaryServing(ctx, in)
}

// SetShardTabletControl is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetShardTabletControl(ctx context.Context, in *vtctldatapb.SetShardTabletControlRequest, opts ...grpc.CallOption) (*vtctldatapb.SetShardTabletControlResponse, error) {
	return client.s.SetShardTabletControl(ctx, in)
}

// SetWritable is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetWritable(ctx context.Context, in *vtctldatapb.SetWritableRequest, opts ...grpc.CallOption) (*vtctldatapb.SetWritableResponse, error) {
	return client.s.SetWritable(ctx, in)
}

// ShardReplicationAdd is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ShardReplicationAdd(ctx context.Context, in *vtctldatapb.ShardReplicationAddRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationAddResponse, error) {
	return client.s.ShardReplicationAdd(ctx, in)
}

// ShardReplicationFix is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ShardReplicationFix(ctx context.Context, in *vtctldatapb.ShardReplicationFixRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationFixResponse, error) {
	return client.s.ShardReplicationFix(ctx, in)
}

// ShardReplicationPositions is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ShardReplicationPositions(ctx context.Context, in *vtctldatapb.ShardReplicationPositionsRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationPositionsResponse, error) {
	return client.s.ShardReplicationPositions(ctx, in)
}

// ShardReplicationRemove is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ShardReplicationRemove(ctx context.Context, in *vtctldatapb.ShardReplicationRemoveRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationRemoveResponse, error) {
	return client.s.ShardReplicationRemove(ctx, in)
}

// SleepTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SleepTablet(ctx context.Context, in *vtctldatapb.SleepTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.SleepTabletResponse, error) {
	return client.s.SleepTablet(ctx, in)
}

// SourceShardAdd is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SourceShardAdd(ctx context.Context, in *vtctldatapb.SourceShardAddRequest, opts ...grpc.CallOption) (*vtctldatapb.SourceShardAddResponse, error) {
	return client.s.SourceShardAdd(ctx, in)
}

// SourceShardDelete is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SourceShardDelete(ctx context.Context, in *vtctldatapb.SourceShardDeleteRequest, opts ...grpc.CallOption) (*vtctldatapb.SourceShardDeleteResponse, error) {
	return client.s.SourceShardDelete(ctx, in)
}

// StartReplication is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) StartReplication(ctx context.Context, in *vtctldatapb.StartReplicationRequest, opts ...grpc.CallOption) (*vtctldatapb.StartReplicationResponse, error) {
	return client.s.StartReplication(ctx, in)
}

// StopReplication is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) StopReplication(ctx context.Context, in *vtctldatapb.StopReplicationRequest, opts ...grpc.CallOption) (*vtctldatapb.StopReplicationResponse, error) {
	return client.s.StopReplication(ctx, in)
}

// TabletExternallyReparented is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) TabletExternallyReparented(ctx context.Context, in *vtctldatapb.TabletExternallyReparentedRequest, opts ...grpc.CallOption) (*vtctldatapb.TabletExternallyReparentedResponse, error) {
	return client.s.TabletExternallyReparented(ctx, in)
}

// UpdateCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) UpdateCellInfo(ctx context.Context, in *vtctldatapb.UpdateCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateCellInfoResponse, error) {
	return client.s.UpdateCellInfo(ctx, in)
}

// UpdateCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) UpdateCellsAlias(ctx context.Context, in *vtctldatapb.UpdateCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateCellsAliasResponse, error) {
	return client.s.UpdateCellsAlias(ctx, in)
}

// UpdateThrottlerConfig is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) UpdateThrottlerConfig(ctx context.Context, in *vtctldatapb.UpdateThrottlerConfigRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateThrottlerConfigResponse, error) {
	return client.s.UpdateThrottlerConfig(ctx, in)
}

// Validate is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) Validate(ctx context.Context, in *vtctldatapb.ValidateRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateResponse, error) {
	return client.s.Validate(ctx, in)
}

// ValidateKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateKeyspace(ctx context.Context, in *vtctldatapb.ValidateKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateKeyspaceResponse, error) {
	return client.s.ValidateKeyspace(ctx, in)
}

// ValidateSchemaKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateSchemaKeyspace(ctx context.Context, in *vtctldatapb.ValidateSchemaKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateSchemaKeyspaceResponse, error) {
	return client.s.ValidateSchemaKeyspace(ctx, in)
}

// ValidateShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateShard(ctx context.Context, in *vtctldatapb.ValidateShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateShardResponse, error) {
	return client.s.ValidateShard(ctx, in)
}

// ValidateVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateVSchema(ctx context.Context, in *vtctldatapb.ValidateVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateVSchemaResponse, error) {
	return client.s.ValidateVSchema(ctx, in)
}

// ValidateVersionKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateVersionKeyspace(ctx context.Context, in *vtctldatapb.ValidateVersionKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateVersionKeyspaceResponse, error) {
	return client.s.ValidateVersionKeyspace(ctx, in)
}

// ValidateVersionShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateVersionShard(ctx context.Context, in *vtctldatapb.ValidateVersionShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateVersionShardResponse, error) {
	return client.s.ValidateVersionShard(ctx, in)
}

// WorkflowDelete is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) WorkflowDelete(ctx context.Context, in *vtctldatapb.WorkflowDeleteRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowDeleteResponse, error) {
	return client.s.WorkflowDelete(ctx, in)
}

// WorkflowStatus is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) WorkflowStatus(ctx context.Context, in *vtctldatapb.WorkflowStatusRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowStatusResponse, error) {
	return client.s.WorkflowStatus(ctx, in)
}

// WorkflowSwitchTraffic is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) WorkflowSwitchTraffic(ctx context.Context, in *vtctldatapb.WorkflowSwitchTrafficRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowSwitchTrafficResponse, error) {
	return client.s.WorkflowSwitchTraffic(ctx, in)
}

// WorkflowUpdate is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) WorkflowUpdate(ctx context.Context, in *vtctldatapb.WorkflowUpdateRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowUpdateResponse, error) {
	return client.s.WorkflowUpdate(ctx, in)
}
