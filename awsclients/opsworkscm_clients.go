package awsclients

import (
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"go.temporal.io/sdk/workflow"
)

type OpsWorksCMClient interface {
       AssociateNode(ctx workflow.Context, input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error)
       AssociateNodeAsync(ctx workflow.Context, input *opsworkscm.AssociateNodeInput) *OpsworkscmAssociateNodeResult

       CreateBackup(ctx workflow.Context, input *opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error)
       CreateBackupAsync(ctx workflow.Context, input *opsworkscm.CreateBackupInput) *OpsworkscmCreateBackupResult

       CreateServer(ctx workflow.Context, input *opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error)
       CreateServerAsync(ctx workflow.Context, input *opsworkscm.CreateServerInput) *OpsworkscmCreateServerResult

       DeleteBackup(ctx workflow.Context, input *opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error)
       DeleteBackupAsync(ctx workflow.Context, input *opsworkscm.DeleteBackupInput) *OpsworkscmDeleteBackupResult

       DeleteServer(ctx workflow.Context, input *opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error)
       DeleteServerAsync(ctx workflow.Context, input *opsworkscm.DeleteServerInput) *OpsworkscmDeleteServerResult

       DescribeAccountAttributes(ctx workflow.Context, input *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error)
       DescribeAccountAttributesAsync(ctx workflow.Context, input *opsworkscm.DescribeAccountAttributesInput) *OpsworkscmDescribeAccountAttributesResult

       DescribeBackups(ctx workflow.Context, input *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error)
       DescribeBackupsAsync(ctx workflow.Context, input *opsworkscm.DescribeBackupsInput) *OpsworkscmDescribeBackupsResult

       DescribeEvents(ctx workflow.Context, input *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error)
       DescribeEventsAsync(ctx workflow.Context, input *opsworkscm.DescribeEventsInput) *OpsworkscmDescribeEventsResult

       DescribeNodeAssociationStatus(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error)
       DescribeNodeAssociationStatusAsync(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) *OpsworkscmDescribeNodeAssociationStatusResult

       DescribeServers(ctx workflow.Context, input *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error)
       DescribeServersAsync(ctx workflow.Context, input *opsworkscm.DescribeServersInput) *OpsworkscmDescribeServersResult

       DisassociateNode(ctx workflow.Context, input *opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error)
       DisassociateNodeAsync(ctx workflow.Context, input *opsworkscm.DisassociateNodeInput) *OpsworkscmDisassociateNodeResult

       ExportServerEngineAttribute(ctx workflow.Context, input *opsworkscm.ExportServerEngineAttributeInput) (*opsworkscm.ExportServerEngineAttributeOutput, error)
       ExportServerEngineAttributeAsync(ctx workflow.Context, input *opsworkscm.ExportServerEngineAttributeInput) *OpsworkscmExportServerEngineAttributeResult

       ListTagsForResource(ctx workflow.Context, input *opsworkscm.ListTagsForResourceInput) (*opsworkscm.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *opsworkscm.ListTagsForResourceInput) *OpsworkscmListTagsForResourceResult

       RestoreServer(ctx workflow.Context, input *opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error)
       RestoreServerAsync(ctx workflow.Context, input *opsworkscm.RestoreServerInput) *OpsworkscmRestoreServerResult

       StartMaintenance(ctx workflow.Context, input *opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error)
       StartMaintenanceAsync(ctx workflow.Context, input *opsworkscm.StartMaintenanceInput) *OpsworkscmStartMaintenanceResult

       TagResource(ctx workflow.Context, input *opsworkscm.TagResourceInput) (*opsworkscm.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *opsworkscm.TagResourceInput) *OpsworkscmTagResourceResult

       UntagResource(ctx workflow.Context, input *opsworkscm.UntagResourceInput) (*opsworkscm.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *opsworkscm.UntagResourceInput) *OpsworkscmUntagResourceResult

       UpdateServer(ctx workflow.Context, input *opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error)
       UpdateServerAsync(ctx workflow.Context, input *opsworkscm.UpdateServerInput) *OpsworkscmUpdateServerResult

       UpdateServerEngineAttributes(ctx workflow.Context, input *opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error)
       UpdateServerEngineAttributesAsync(ctx workflow.Context, input *opsworkscm.UpdateServerEngineAttributesInput) *OpsworkscmUpdateServerEngineAttributesResult

       WaitUntilNodeAssociated(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) error}

type OpsWorksCMStub struct {
}

func NewOpsWorksCMStub() OpsWorksCMClient {
    return &OpsWorksCMStub{}
}

type OpsworkscmAssociateNodeResult struct {
	Result workflow.Future
}

func (r *OpsworkscmAssociateNodeResult) Get(ctx workflow.Context) (*opsworkscm.AssociateNodeOutput, error) {
    var output opsworkscm.AssociateNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmCreateBackupResult struct {
	Result workflow.Future
}

func (r *OpsworkscmCreateBackupResult) Get(ctx workflow.Context) (*opsworkscm.CreateBackupOutput, error) {
    var output opsworkscm.CreateBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmCreateServerResult struct {
	Result workflow.Future
}

func (r *OpsworkscmCreateServerResult) Get(ctx workflow.Context) (*opsworkscm.CreateServerOutput, error) {
    var output opsworkscm.CreateServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDeleteBackupResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDeleteBackupResult) Get(ctx workflow.Context) (*opsworkscm.DeleteBackupOutput, error) {
    var output opsworkscm.DeleteBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDeleteServerResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDeleteServerResult) Get(ctx workflow.Context) (*opsworkscm.DeleteServerOutput, error) {
    var output opsworkscm.DeleteServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDescribeAccountAttributesResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDescribeAccountAttributesResult) Get(ctx workflow.Context) (*opsworkscm.DescribeAccountAttributesOutput, error) {
    var output opsworkscm.DescribeAccountAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDescribeBackupsResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDescribeBackupsResult) Get(ctx workflow.Context) (*opsworkscm.DescribeBackupsOutput, error) {
    var output opsworkscm.DescribeBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDescribeEventsResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDescribeEventsResult) Get(ctx workflow.Context) (*opsworkscm.DescribeEventsOutput, error) {
    var output opsworkscm.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDescribeNodeAssociationStatusResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDescribeNodeAssociationStatusResult) Get(ctx workflow.Context) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
    var output opsworkscm.DescribeNodeAssociationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDescribeServersResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDescribeServersResult) Get(ctx workflow.Context) (*opsworkscm.DescribeServersOutput, error) {
    var output opsworkscm.DescribeServersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmDisassociateNodeResult struct {
	Result workflow.Future
}

func (r *OpsworkscmDisassociateNodeResult) Get(ctx workflow.Context) (*opsworkscm.DisassociateNodeOutput, error) {
    var output opsworkscm.DisassociateNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmExportServerEngineAttributeResult struct {
	Result workflow.Future
}

func (r *OpsworkscmExportServerEngineAttributeResult) Get(ctx workflow.Context) (*opsworkscm.ExportServerEngineAttributeOutput, error) {
    var output opsworkscm.ExportServerEngineAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *OpsworkscmListTagsForResourceResult) Get(ctx workflow.Context) (*opsworkscm.ListTagsForResourceOutput, error) {
    var output opsworkscm.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmRestoreServerResult struct {
	Result workflow.Future
}

func (r *OpsworkscmRestoreServerResult) Get(ctx workflow.Context) (*opsworkscm.RestoreServerOutput, error) {
    var output opsworkscm.RestoreServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmStartMaintenanceResult struct {
	Result workflow.Future
}

func (r *OpsworkscmStartMaintenanceResult) Get(ctx workflow.Context) (*opsworkscm.StartMaintenanceOutput, error) {
    var output opsworkscm.StartMaintenanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmTagResourceResult struct {
	Result workflow.Future
}

func (r *OpsworkscmTagResourceResult) Get(ctx workflow.Context) (*opsworkscm.TagResourceOutput, error) {
    var output opsworkscm.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmUntagResourceResult struct {
	Result workflow.Future
}

func (r *OpsworkscmUntagResourceResult) Get(ctx workflow.Context) (*opsworkscm.UntagResourceOutput, error) {
    var output opsworkscm.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmUpdateServerResult struct {
	Result workflow.Future
}

func (r *OpsworkscmUpdateServerResult) Get(ctx workflow.Context) (*opsworkscm.UpdateServerOutput, error) {
    var output opsworkscm.UpdateServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type OpsworkscmUpdateServerEngineAttributesResult struct {
	Result workflow.Future
}

func (r *OpsworkscmUpdateServerEngineAttributesResult) Get(ctx workflow.Context) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
    var output opsworkscm.UpdateServerEngineAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) AssociateNode(ctx workflow.Context, input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error) {
    var output opsworkscm.AssociateNodeOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.AssociateNode", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) AssociateNodeAsync(ctx workflow.Context, input *opsworkscm.AssociateNodeInput) *OpsworkscmAssociateNodeResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.AssociateNode", input)
    return &OpsworkscmAssociateNodeResult{Result: future}
}

func (a *OpsWorksCMStub) CreateBackup(ctx workflow.Context, input *opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error) {
    var output opsworkscm.CreateBackupOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.CreateBackup", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) CreateBackupAsync(ctx workflow.Context, input *opsworkscm.CreateBackupInput) *OpsworkscmCreateBackupResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.CreateBackup", input)
    return &OpsworkscmCreateBackupResult{Result: future}
}

func (a *OpsWorksCMStub) CreateServer(ctx workflow.Context, input *opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error) {
    var output opsworkscm.CreateServerOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.CreateServer", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) CreateServerAsync(ctx workflow.Context, input *opsworkscm.CreateServerInput) *OpsworkscmCreateServerResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.CreateServer", input)
    return &OpsworkscmCreateServerResult{Result: future}
}

func (a *OpsWorksCMStub) DeleteBackup(ctx workflow.Context, input *opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error) {
    var output opsworkscm.DeleteBackupOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DeleteBackup", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DeleteBackupAsync(ctx workflow.Context, input *opsworkscm.DeleteBackupInput) *OpsworkscmDeleteBackupResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DeleteBackup", input)
    return &OpsworkscmDeleteBackupResult{Result: future}
}

func (a *OpsWorksCMStub) DeleteServer(ctx workflow.Context, input *opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error) {
    var output opsworkscm.DeleteServerOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DeleteServer", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DeleteServerAsync(ctx workflow.Context, input *opsworkscm.DeleteServerInput) *OpsworkscmDeleteServerResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DeleteServer", input)
    return &OpsworkscmDeleteServerResult{Result: future}
}

func (a *OpsWorksCMStub) DescribeAccountAttributes(ctx workflow.Context, input *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error) {
    var output opsworkscm.DescribeAccountAttributesOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeAccountAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DescribeAccountAttributesAsync(ctx workflow.Context, input *opsworkscm.DescribeAccountAttributesInput) *OpsworkscmDescribeAccountAttributesResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeAccountAttributes", input)
    return &OpsworkscmDescribeAccountAttributesResult{Result: future}
}

func (a *OpsWorksCMStub) DescribeBackups(ctx workflow.Context, input *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error) {
    var output opsworkscm.DescribeBackupsOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeBackups", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DescribeBackupsAsync(ctx workflow.Context, input *opsworkscm.DescribeBackupsInput) *OpsworkscmDescribeBackupsResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeBackups", input)
    return &OpsworkscmDescribeBackupsResult{Result: future}
}

func (a *OpsWorksCMStub) DescribeEvents(ctx workflow.Context, input *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error) {
    var output opsworkscm.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeEvents", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DescribeEventsAsync(ctx workflow.Context, input *opsworkscm.DescribeEventsInput) *OpsworkscmDescribeEventsResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeEvents", input)
    return &OpsworkscmDescribeEventsResult{Result: future}
}

func (a *OpsWorksCMStub) DescribeNodeAssociationStatus(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
    var output opsworkscm.DescribeNodeAssociationStatusOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeNodeAssociationStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DescribeNodeAssociationStatusAsync(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) *OpsworkscmDescribeNodeAssociationStatusResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeNodeAssociationStatus", input)
    return &OpsworkscmDescribeNodeAssociationStatusResult{Result: future}
}

func (a *OpsWorksCMStub) DescribeServers(ctx workflow.Context, input *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error) {
    var output opsworkscm.DescribeServersOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeServers", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DescribeServersAsync(ctx workflow.Context, input *opsworkscm.DescribeServersInput) *OpsworkscmDescribeServersResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DescribeServers", input)
    return &OpsworkscmDescribeServersResult{Result: future}
}

func (a *OpsWorksCMStub) DisassociateNode(ctx workflow.Context, input *opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error) {
    var output opsworkscm.DisassociateNodeOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.DisassociateNode", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) DisassociateNodeAsync(ctx workflow.Context, input *opsworkscm.DisassociateNodeInput) *OpsworkscmDisassociateNodeResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.DisassociateNode", input)
    return &OpsworkscmDisassociateNodeResult{Result: future}
}

func (a *OpsWorksCMStub) ExportServerEngineAttribute(ctx workflow.Context, input *opsworkscm.ExportServerEngineAttributeInput) (*opsworkscm.ExportServerEngineAttributeOutput, error) {
    var output opsworkscm.ExportServerEngineAttributeOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.ExportServerEngineAttribute", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) ExportServerEngineAttributeAsync(ctx workflow.Context, input *opsworkscm.ExportServerEngineAttributeInput) *OpsworkscmExportServerEngineAttributeResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.ExportServerEngineAttribute", input)
    return &OpsworkscmExportServerEngineAttributeResult{Result: future}
}

func (a *OpsWorksCMStub) ListTagsForResource(ctx workflow.Context, input *opsworkscm.ListTagsForResourceInput) (*opsworkscm.ListTagsForResourceOutput, error) {
    var output opsworkscm.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) ListTagsForResourceAsync(ctx workflow.Context, input *opsworkscm.ListTagsForResourceInput) *OpsworkscmListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.ListTagsForResource", input)
    return &OpsworkscmListTagsForResourceResult{Result: future}
}

func (a *OpsWorksCMStub) RestoreServer(ctx workflow.Context, input *opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error) {
    var output opsworkscm.RestoreServerOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.RestoreServer", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) RestoreServerAsync(ctx workflow.Context, input *opsworkscm.RestoreServerInput) *OpsworkscmRestoreServerResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.RestoreServer", input)
    return &OpsworkscmRestoreServerResult{Result: future}
}

func (a *OpsWorksCMStub) StartMaintenance(ctx workflow.Context, input *opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error) {
    var output opsworkscm.StartMaintenanceOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.StartMaintenance", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) StartMaintenanceAsync(ctx workflow.Context, input *opsworkscm.StartMaintenanceInput) *OpsworkscmStartMaintenanceResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.StartMaintenance", input)
    return &OpsworkscmStartMaintenanceResult{Result: future}
}

func (a *OpsWorksCMStub) TagResource(ctx workflow.Context, input *opsworkscm.TagResourceInput) (*opsworkscm.TagResourceOutput, error) {
    var output opsworkscm.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) TagResourceAsync(ctx workflow.Context, input *opsworkscm.TagResourceInput) *OpsworkscmTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.TagResource", input)
    return &OpsworkscmTagResourceResult{Result: future}
}

func (a *OpsWorksCMStub) UntagResource(ctx workflow.Context, input *opsworkscm.UntagResourceInput) (*opsworkscm.UntagResourceOutput, error) {
    var output opsworkscm.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) UntagResourceAsync(ctx workflow.Context, input *opsworkscm.UntagResourceInput) *OpsworkscmUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.UntagResource", input)
    return &OpsworkscmUntagResourceResult{Result: future}
}

func (a *OpsWorksCMStub) UpdateServer(ctx workflow.Context, input *opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error) {
    var output opsworkscm.UpdateServerOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.UpdateServer", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) UpdateServerAsync(ctx workflow.Context, input *opsworkscm.UpdateServerInput) *OpsworkscmUpdateServerResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.UpdateServer", input)
    return &OpsworkscmUpdateServerResult{Result: future}
}

func (a *OpsWorksCMStub) UpdateServerEngineAttributes(ctx workflow.Context, input *opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
    var output opsworkscm.UpdateServerEngineAttributesOutput
    err := workflow.ExecuteActivity(ctx, "OpsWorksCM.UpdateServerEngineAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksCMStub) UpdateServerEngineAttributesAsync(ctx workflow.Context, input *opsworkscm.UpdateServerEngineAttributesInput) *OpsworkscmUpdateServerEngineAttributesResult {
    future := workflow.ExecuteActivity(ctx, "OpsWorksCM.UpdateServerEngineAttributes", input)
    return &OpsworkscmUpdateServerEngineAttributesResult{Result: future}
}

func (a *OpsWorksCMStub) WaitUntilNodeAssociated(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) error {
    return workflow.ExecuteActivity(ctx, "OpsWorksCM.WaitUntilNodeAssociated", input).Get(ctx, nil)
}

func (a *OpsWorksCMStub) WaitUntilNodeAssociatedAsync(ctx workflow.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "OpsWorksCM.WaitUntilNodeAssociated", input)
}

