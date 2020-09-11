package awsclients

import (
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"go.temporal.io/sdk/workflow"
)

type MigrationHubClient interface {
       AssociateCreatedArtifact(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error)
       AssociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) *MigrationhubAssociateCreatedArtifactResult

       AssociateDiscoveredResource(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error)
       AssociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) *MigrationhubAssociateDiscoveredResourceResult

       CreateProgressUpdateStream(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error)
       CreateProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) *MigrationhubCreateProgressUpdateStreamResult

       DeleteProgressUpdateStream(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error)
       DeleteProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) *MigrationhubDeleteProgressUpdateStreamResult

       DescribeApplicationState(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error)
       DescribeApplicationStateAsync(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) *MigrationhubDescribeApplicationStateResult

       DescribeMigrationTask(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error)
       DescribeMigrationTaskAsync(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) *MigrationhubDescribeMigrationTaskResult

       DisassociateCreatedArtifact(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error)
       DisassociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) *MigrationhubDisassociateCreatedArtifactResult

       DisassociateDiscoveredResource(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error)
       DisassociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) *MigrationhubDisassociateDiscoveredResourceResult

       ImportMigrationTask(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error)
       ImportMigrationTaskAsync(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) *MigrationhubImportMigrationTaskResult

       ListApplicationStates(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error)
       ListApplicationStatesAsync(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) *MigrationhubListApplicationStatesResult

       ListCreatedArtifacts(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error)
       ListCreatedArtifactsAsync(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) *MigrationhubListCreatedArtifactsResult

       ListDiscoveredResources(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error)
       ListDiscoveredResourcesAsync(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) *MigrationhubListDiscoveredResourcesResult

       ListMigrationTasks(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error)
       ListMigrationTasksAsync(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) *MigrationhubListMigrationTasksResult

       ListProgressUpdateStreams(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error)
       ListProgressUpdateStreamsAsync(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) *MigrationhubListProgressUpdateStreamsResult

       NotifyApplicationState(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error)
       NotifyApplicationStateAsync(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) *MigrationhubNotifyApplicationStateResult

       NotifyMigrationTaskState(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error)
       NotifyMigrationTaskStateAsync(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) *MigrationhubNotifyMigrationTaskStateResult

       PutResourceAttributes(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error)
       PutResourceAttributesAsync(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) *MigrationhubPutResourceAttributesResult
}

type MigrationHubStub struct {
}

func NewMigrationHubStub() MigrationHubClient {
    return &MigrationHubStub{}
}

type MigrationhubAssociateCreatedArtifactResult struct {
	Result workflow.Future
}

func (r *MigrationhubAssociateCreatedArtifactResult) Get(ctx workflow.Context) (*migrationhub.AssociateCreatedArtifactOutput, error) {
    var output migrationhub.AssociateCreatedArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubAssociateDiscoveredResourceResult struct {
	Result workflow.Future
}

func (r *MigrationhubAssociateDiscoveredResourceResult) Get(ctx workflow.Context) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
    var output migrationhub.AssociateDiscoveredResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubCreateProgressUpdateStreamResult struct {
	Result workflow.Future
}

func (r *MigrationhubCreateProgressUpdateStreamResult) Get(ctx workflow.Context) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
    var output migrationhub.CreateProgressUpdateStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubDeleteProgressUpdateStreamResult struct {
	Result workflow.Future
}

func (r *MigrationhubDeleteProgressUpdateStreamResult) Get(ctx workflow.Context) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
    var output migrationhub.DeleteProgressUpdateStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubDescribeApplicationStateResult struct {
	Result workflow.Future
}

func (r *MigrationhubDescribeApplicationStateResult) Get(ctx workflow.Context) (*migrationhub.DescribeApplicationStateOutput, error) {
    var output migrationhub.DescribeApplicationStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubDescribeMigrationTaskResult struct {
	Result workflow.Future
}

func (r *MigrationhubDescribeMigrationTaskResult) Get(ctx workflow.Context) (*migrationhub.DescribeMigrationTaskOutput, error) {
    var output migrationhub.DescribeMigrationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubDisassociateCreatedArtifactResult struct {
	Result workflow.Future
}

func (r *MigrationhubDisassociateCreatedArtifactResult) Get(ctx workflow.Context) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
    var output migrationhub.DisassociateCreatedArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubDisassociateDiscoveredResourceResult struct {
	Result workflow.Future
}

func (r *MigrationhubDisassociateDiscoveredResourceResult) Get(ctx workflow.Context) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
    var output migrationhub.DisassociateDiscoveredResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubImportMigrationTaskResult struct {
	Result workflow.Future
}

func (r *MigrationhubImportMigrationTaskResult) Get(ctx workflow.Context) (*migrationhub.ImportMigrationTaskOutput, error) {
    var output migrationhub.ImportMigrationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubListApplicationStatesResult struct {
	Result workflow.Future
}

func (r *MigrationhubListApplicationStatesResult) Get(ctx workflow.Context) (*migrationhub.ListApplicationStatesOutput, error) {
    var output migrationhub.ListApplicationStatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubListCreatedArtifactsResult struct {
	Result workflow.Future
}

func (r *MigrationhubListCreatedArtifactsResult) Get(ctx workflow.Context) (*migrationhub.ListCreatedArtifactsOutput, error) {
    var output migrationhub.ListCreatedArtifactsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubListDiscoveredResourcesResult struct {
	Result workflow.Future
}

func (r *MigrationhubListDiscoveredResourcesResult) Get(ctx workflow.Context) (*migrationhub.ListDiscoveredResourcesOutput, error) {
    var output migrationhub.ListDiscoveredResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubListMigrationTasksResult struct {
	Result workflow.Future
}

func (r *MigrationhubListMigrationTasksResult) Get(ctx workflow.Context) (*migrationhub.ListMigrationTasksOutput, error) {
    var output migrationhub.ListMigrationTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubListProgressUpdateStreamsResult struct {
	Result workflow.Future
}

func (r *MigrationhubListProgressUpdateStreamsResult) Get(ctx workflow.Context) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
    var output migrationhub.ListProgressUpdateStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubNotifyApplicationStateResult struct {
	Result workflow.Future
}

func (r *MigrationhubNotifyApplicationStateResult) Get(ctx workflow.Context) (*migrationhub.NotifyApplicationStateOutput, error) {
    var output migrationhub.NotifyApplicationStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubNotifyMigrationTaskStateResult struct {
	Result workflow.Future
}

func (r *MigrationhubNotifyMigrationTaskStateResult) Get(ctx workflow.Context) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
    var output migrationhub.NotifyMigrationTaskStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MigrationhubPutResourceAttributesResult struct {
	Result workflow.Future
}

func (r *MigrationhubPutResourceAttributesResult) Get(ctx workflow.Context) (*migrationhub.PutResourceAttributesOutput, error) {
    var output migrationhub.PutResourceAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) AssociateCreatedArtifact(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error) {
    var output migrationhub.AssociateCreatedArtifactOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.AssociateCreatedArtifact", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) AssociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) *MigrationhubAssociateCreatedArtifactResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.AssociateCreatedArtifact", input)
    return &MigrationhubAssociateCreatedArtifactResult{Result: future}
}

func (a *MigrationHubStub) AssociateDiscoveredResource(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
    var output migrationhub.AssociateDiscoveredResourceOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.AssociateDiscoveredResource", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) AssociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) *MigrationhubAssociateDiscoveredResourceResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.AssociateDiscoveredResource", input)
    return &MigrationhubAssociateDiscoveredResourceResult{Result: future}
}

func (a *MigrationHubStub) CreateProgressUpdateStream(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
    var output migrationhub.CreateProgressUpdateStreamOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.CreateProgressUpdateStream", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) CreateProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) *MigrationhubCreateProgressUpdateStreamResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.CreateProgressUpdateStream", input)
    return &MigrationhubCreateProgressUpdateStreamResult{Result: future}
}

func (a *MigrationHubStub) DeleteProgressUpdateStream(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
    var output migrationhub.DeleteProgressUpdateStreamOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.DeleteProgressUpdateStream", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) DeleteProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) *MigrationhubDeleteProgressUpdateStreamResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.DeleteProgressUpdateStream", input)
    return &MigrationhubDeleteProgressUpdateStreamResult{Result: future}
}

func (a *MigrationHubStub) DescribeApplicationState(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error) {
    var output migrationhub.DescribeApplicationStateOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.DescribeApplicationState", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) DescribeApplicationStateAsync(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) *MigrationhubDescribeApplicationStateResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.DescribeApplicationState", input)
    return &MigrationhubDescribeApplicationStateResult{Result: future}
}

func (a *MigrationHubStub) DescribeMigrationTask(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error) {
    var output migrationhub.DescribeMigrationTaskOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.DescribeMigrationTask", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) DescribeMigrationTaskAsync(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) *MigrationhubDescribeMigrationTaskResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.DescribeMigrationTask", input)
    return &MigrationhubDescribeMigrationTaskResult{Result: future}
}

func (a *MigrationHubStub) DisassociateCreatedArtifact(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
    var output migrationhub.DisassociateCreatedArtifactOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.DisassociateCreatedArtifact", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) DisassociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) *MigrationhubDisassociateCreatedArtifactResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.DisassociateCreatedArtifact", input)
    return &MigrationhubDisassociateCreatedArtifactResult{Result: future}
}

func (a *MigrationHubStub) DisassociateDiscoveredResource(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
    var output migrationhub.DisassociateDiscoveredResourceOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.DisassociateDiscoveredResource", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) DisassociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) *MigrationhubDisassociateDiscoveredResourceResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.DisassociateDiscoveredResource", input)
    return &MigrationhubDisassociateDiscoveredResourceResult{Result: future}
}

func (a *MigrationHubStub) ImportMigrationTask(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error) {
    var output migrationhub.ImportMigrationTaskOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.ImportMigrationTask", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) ImportMigrationTaskAsync(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) *MigrationhubImportMigrationTaskResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.ImportMigrationTask", input)
    return &MigrationhubImportMigrationTaskResult{Result: future}
}

func (a *MigrationHubStub) ListApplicationStates(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error) {
    var output migrationhub.ListApplicationStatesOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.ListApplicationStates", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) ListApplicationStatesAsync(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) *MigrationhubListApplicationStatesResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.ListApplicationStates", input)
    return &MigrationhubListApplicationStatesResult{Result: future}
}

func (a *MigrationHubStub) ListCreatedArtifacts(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error) {
    var output migrationhub.ListCreatedArtifactsOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.ListCreatedArtifacts", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) ListCreatedArtifactsAsync(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) *MigrationhubListCreatedArtifactsResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.ListCreatedArtifacts", input)
    return &MigrationhubListCreatedArtifactsResult{Result: future}
}

func (a *MigrationHubStub) ListDiscoveredResources(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error) {
    var output migrationhub.ListDiscoveredResourcesOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.ListDiscoveredResources", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) ListDiscoveredResourcesAsync(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) *MigrationhubListDiscoveredResourcesResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.ListDiscoveredResources", input)
    return &MigrationhubListDiscoveredResourcesResult{Result: future}
}

func (a *MigrationHubStub) ListMigrationTasks(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error) {
    var output migrationhub.ListMigrationTasksOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.ListMigrationTasks", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) ListMigrationTasksAsync(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) *MigrationhubListMigrationTasksResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.ListMigrationTasks", input)
    return &MigrationhubListMigrationTasksResult{Result: future}
}

func (a *MigrationHubStub) ListProgressUpdateStreams(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
    var output migrationhub.ListProgressUpdateStreamsOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.ListProgressUpdateStreams", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) ListProgressUpdateStreamsAsync(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) *MigrationhubListProgressUpdateStreamsResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.ListProgressUpdateStreams", input)
    return &MigrationhubListProgressUpdateStreamsResult{Result: future}
}

func (a *MigrationHubStub) NotifyApplicationState(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error) {
    var output migrationhub.NotifyApplicationStateOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.NotifyApplicationState", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) NotifyApplicationStateAsync(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) *MigrationhubNotifyApplicationStateResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.NotifyApplicationState", input)
    return &MigrationhubNotifyApplicationStateResult{Result: future}
}

func (a *MigrationHubStub) NotifyMigrationTaskState(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
    var output migrationhub.NotifyMigrationTaskStateOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.NotifyMigrationTaskState", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) NotifyMigrationTaskStateAsync(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) *MigrationhubNotifyMigrationTaskStateResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.NotifyMigrationTaskState", input)
    return &MigrationhubNotifyMigrationTaskStateResult{Result: future}
}

func (a *MigrationHubStub) PutResourceAttributes(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error) {
    var output migrationhub.PutResourceAttributesOutput
    err := workflow.ExecuteActivity(ctx, "MigrationHub.PutResourceAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubStub) PutResourceAttributesAsync(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) *MigrationhubPutResourceAttributesResult {
    future := workflow.ExecuteActivity(ctx, "MigrationHub.PutResourceAttributes", input)
    return &MigrationhubPutResourceAttributesResult{Result: future}
}
