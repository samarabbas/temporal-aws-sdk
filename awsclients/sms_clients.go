package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sms"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SMSClient interface {
    CreateApp(ctx workflow.Context, input *sms.CreateAppInput) (*sms.CreateAppOutput, error)
    CreateAppAsync(ctx workflow.Context, input *sms.CreateAppInput) *SmsCreateAppResult

    CreateReplicationJob(ctx workflow.Context, input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error)
    CreateReplicationJobAsync(ctx workflow.Context, input *sms.CreateReplicationJobInput) *SmsCreateReplicationJobResult

    DeleteApp(ctx workflow.Context, input *sms.DeleteAppInput) (*sms.DeleteAppOutput, error)
    DeleteAppAsync(ctx workflow.Context, input *sms.DeleteAppInput) *SmsDeleteAppResult

    DeleteAppLaunchConfiguration(ctx workflow.Context, input *sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error)
    DeleteAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppLaunchConfigurationInput) *SmsDeleteAppLaunchConfigurationResult

    DeleteAppReplicationConfiguration(ctx workflow.Context, input *sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error)
    DeleteAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppReplicationConfigurationInput) *SmsDeleteAppReplicationConfigurationResult

    DeleteAppValidationConfiguration(ctx workflow.Context, input *sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error)
    DeleteAppValidationConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppValidationConfigurationInput) *SmsDeleteAppValidationConfigurationResult

    DeleteReplicationJob(ctx workflow.Context, input *sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error)
    DeleteReplicationJobAsync(ctx workflow.Context, input *sms.DeleteReplicationJobInput) *SmsDeleteReplicationJobResult

    DeleteServerCatalog(ctx workflow.Context, input *sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error)
    DeleteServerCatalogAsync(ctx workflow.Context, input *sms.DeleteServerCatalogInput) *SmsDeleteServerCatalogResult

    DisassociateConnector(ctx workflow.Context, input *sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error)
    DisassociateConnectorAsync(ctx workflow.Context, input *sms.DisassociateConnectorInput) *SmsDisassociateConnectorResult

    GenerateChangeSet(ctx workflow.Context, input *sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error)
    GenerateChangeSetAsync(ctx workflow.Context, input *sms.GenerateChangeSetInput) *SmsGenerateChangeSetResult

    GenerateTemplate(ctx workflow.Context, input *sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error)
    GenerateTemplateAsync(ctx workflow.Context, input *sms.GenerateTemplateInput) *SmsGenerateTemplateResult

    GetApp(ctx workflow.Context, input *sms.GetAppInput) (*sms.GetAppOutput, error)
    GetAppAsync(ctx workflow.Context, input *sms.GetAppInput) *SmsGetAppResult

    GetAppLaunchConfiguration(ctx workflow.Context, input *sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error)
    GetAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.GetAppLaunchConfigurationInput) *SmsGetAppLaunchConfigurationResult

    GetAppReplicationConfiguration(ctx workflow.Context, input *sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error)
    GetAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.GetAppReplicationConfigurationInput) *SmsGetAppReplicationConfigurationResult

    GetAppValidationConfiguration(ctx workflow.Context, input *sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error)
    GetAppValidationConfigurationAsync(ctx workflow.Context, input *sms.GetAppValidationConfigurationInput) *SmsGetAppValidationConfigurationResult

    GetAppValidationOutput(ctx workflow.Context, input *sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error)
    GetAppValidationOutputAsync(ctx workflow.Context, input *sms.GetAppValidationOutputInput) *SmsGetAppValidationOutputResult

    GetConnectors(ctx workflow.Context, input *sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error)
    GetConnectorsAsync(ctx workflow.Context, input *sms.GetConnectorsInput) *SmsGetConnectorsResult

    GetReplicationJobs(ctx workflow.Context, input *sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error)
    GetReplicationJobsAsync(ctx workflow.Context, input *sms.GetReplicationJobsInput) *SmsGetReplicationJobsResult

    GetReplicationRuns(ctx workflow.Context, input *sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error)
    GetReplicationRunsAsync(ctx workflow.Context, input *sms.GetReplicationRunsInput) *SmsGetReplicationRunsResult

    GetServers(ctx workflow.Context, input *sms.GetServersInput) (*sms.GetServersOutput, error)
    GetServersAsync(ctx workflow.Context, input *sms.GetServersInput) *SmsGetServersResult

    ImportAppCatalog(ctx workflow.Context, input *sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error)
    ImportAppCatalogAsync(ctx workflow.Context, input *sms.ImportAppCatalogInput) *SmsImportAppCatalogResult

    ImportServerCatalog(ctx workflow.Context, input *sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error)
    ImportServerCatalogAsync(ctx workflow.Context, input *sms.ImportServerCatalogInput) *SmsImportServerCatalogResult

    LaunchApp(ctx workflow.Context, input *sms.LaunchAppInput) (*sms.LaunchAppOutput, error)
    LaunchAppAsync(ctx workflow.Context, input *sms.LaunchAppInput) *SmsLaunchAppResult

    ListApps(ctx workflow.Context, input *sms.ListAppsInput) (*sms.ListAppsOutput, error)
    ListAppsAsync(ctx workflow.Context, input *sms.ListAppsInput) *SmsListAppsResult

    NotifyAppValidationOutput(ctx workflow.Context, input *sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error)
    NotifyAppValidationOutputAsync(ctx workflow.Context, input *sms.NotifyAppValidationOutputInput) *SmsNotifyAppValidationOutputResult

    PutAppLaunchConfiguration(ctx workflow.Context, input *sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error)
    PutAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.PutAppLaunchConfigurationInput) *SmsPutAppLaunchConfigurationResult

    PutAppReplicationConfiguration(ctx workflow.Context, input *sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error)
    PutAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.PutAppReplicationConfigurationInput) *SmsPutAppReplicationConfigurationResult

    PutAppValidationConfiguration(ctx workflow.Context, input *sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error)
    PutAppValidationConfigurationAsync(ctx workflow.Context, input *sms.PutAppValidationConfigurationInput) *SmsPutAppValidationConfigurationResult

    StartAppReplication(ctx workflow.Context, input *sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error)
    StartAppReplicationAsync(ctx workflow.Context, input *sms.StartAppReplicationInput) *SmsStartAppReplicationResult

    StartOnDemandAppReplication(ctx workflow.Context, input *sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error)
    StartOnDemandAppReplicationAsync(ctx workflow.Context, input *sms.StartOnDemandAppReplicationInput) *SmsStartOnDemandAppReplicationResult

    StartOnDemandReplicationRun(ctx workflow.Context, input *sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error)
    StartOnDemandReplicationRunAsync(ctx workflow.Context, input *sms.StartOnDemandReplicationRunInput) *SmsStartOnDemandReplicationRunResult

    StopAppReplication(ctx workflow.Context, input *sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error)
    StopAppReplicationAsync(ctx workflow.Context, input *sms.StopAppReplicationInput) *SmsStopAppReplicationResult

    TerminateApp(ctx workflow.Context, input *sms.TerminateAppInput) (*sms.TerminateAppOutput, error)
    TerminateAppAsync(ctx workflow.Context, input *sms.TerminateAppInput) *SmsTerminateAppResult

    UpdateApp(ctx workflow.Context, input *sms.UpdateAppInput) (*sms.UpdateAppOutput, error)
    UpdateAppAsync(ctx workflow.Context, input *sms.UpdateAppInput) *SmsUpdateAppResult

    UpdateReplicationJob(ctx workflow.Context, input *sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error)
    UpdateReplicationJobAsync(ctx workflow.Context, input *sms.UpdateReplicationJobInput) *SmsUpdateReplicationJobResult
}

type SmsCreateAppResult struct {
	Result workflow.Future
}

func (r *SmsCreateAppResult) Get(ctx workflow.Context) (*sms.CreateAppOutput, error) {
    var output sms.CreateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsCreateReplicationJobResult struct {
	Result workflow.Future
}

func (r *SmsCreateReplicationJobResult) Get(ctx workflow.Context) (*sms.CreateReplicationJobOutput, error) {
    var output sms.CreateReplicationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDeleteAppResult struct {
	Result workflow.Future
}

func (r *SmsDeleteAppResult) Get(ctx workflow.Context) (*sms.DeleteAppOutput, error) {
    var output sms.DeleteAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDeleteAppLaunchConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsDeleteAppLaunchConfigurationResult) Get(ctx workflow.Context) (*sms.DeleteAppLaunchConfigurationOutput, error) {
    var output sms.DeleteAppLaunchConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDeleteAppReplicationConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsDeleteAppReplicationConfigurationResult) Get(ctx workflow.Context) (*sms.DeleteAppReplicationConfigurationOutput, error) {
    var output sms.DeleteAppReplicationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDeleteAppValidationConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsDeleteAppValidationConfigurationResult) Get(ctx workflow.Context) (*sms.DeleteAppValidationConfigurationOutput, error) {
    var output sms.DeleteAppValidationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDeleteReplicationJobResult struct {
	Result workflow.Future
}

func (r *SmsDeleteReplicationJobResult) Get(ctx workflow.Context) (*sms.DeleteReplicationJobOutput, error) {
    var output sms.DeleteReplicationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDeleteServerCatalogResult struct {
	Result workflow.Future
}

func (r *SmsDeleteServerCatalogResult) Get(ctx workflow.Context) (*sms.DeleteServerCatalogOutput, error) {
    var output sms.DeleteServerCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsDisassociateConnectorResult struct {
	Result workflow.Future
}

func (r *SmsDisassociateConnectorResult) Get(ctx workflow.Context) (*sms.DisassociateConnectorOutput, error) {
    var output sms.DisassociateConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGenerateChangeSetResult struct {
	Result workflow.Future
}

func (r *SmsGenerateChangeSetResult) Get(ctx workflow.Context) (*sms.GenerateChangeSetOutput, error) {
    var output sms.GenerateChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGenerateTemplateResult struct {
	Result workflow.Future
}

func (r *SmsGenerateTemplateResult) Get(ctx workflow.Context) (*sms.GenerateTemplateOutput, error) {
    var output sms.GenerateTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetAppResult struct {
	Result workflow.Future
}

func (r *SmsGetAppResult) Get(ctx workflow.Context) (*sms.GetAppOutput, error) {
    var output sms.GetAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetAppLaunchConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsGetAppLaunchConfigurationResult) Get(ctx workflow.Context) (*sms.GetAppLaunchConfigurationOutput, error) {
    var output sms.GetAppLaunchConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetAppReplicationConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsGetAppReplicationConfigurationResult) Get(ctx workflow.Context) (*sms.GetAppReplicationConfigurationOutput, error) {
    var output sms.GetAppReplicationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetAppValidationConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsGetAppValidationConfigurationResult) Get(ctx workflow.Context) (*sms.GetAppValidationConfigurationOutput, error) {
    var output sms.GetAppValidationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetAppValidationOutputResult struct {
	Result workflow.Future
}

func (r *SmsGetAppValidationOutputResult) Get(ctx workflow.Context) (*sms.GetAppValidationOutputOutput, error) {
    var output sms.GetAppValidationOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetConnectorsResult struct {
	Result workflow.Future
}

func (r *SmsGetConnectorsResult) Get(ctx workflow.Context) (*sms.GetConnectorsOutput, error) {
    var output sms.GetConnectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetReplicationJobsResult struct {
	Result workflow.Future
}

func (r *SmsGetReplicationJobsResult) Get(ctx workflow.Context) (*sms.GetReplicationJobsOutput, error) {
    var output sms.GetReplicationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetReplicationRunsResult struct {
	Result workflow.Future
}

func (r *SmsGetReplicationRunsResult) Get(ctx workflow.Context) (*sms.GetReplicationRunsOutput, error) {
    var output sms.GetReplicationRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsGetServersResult struct {
	Result workflow.Future
}

func (r *SmsGetServersResult) Get(ctx workflow.Context) (*sms.GetServersOutput, error) {
    var output sms.GetServersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsImportAppCatalogResult struct {
	Result workflow.Future
}

func (r *SmsImportAppCatalogResult) Get(ctx workflow.Context) (*sms.ImportAppCatalogOutput, error) {
    var output sms.ImportAppCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsImportServerCatalogResult struct {
	Result workflow.Future
}

func (r *SmsImportServerCatalogResult) Get(ctx workflow.Context) (*sms.ImportServerCatalogOutput, error) {
    var output sms.ImportServerCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsLaunchAppResult struct {
	Result workflow.Future
}

func (r *SmsLaunchAppResult) Get(ctx workflow.Context) (*sms.LaunchAppOutput, error) {
    var output sms.LaunchAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsListAppsResult struct {
	Result workflow.Future
}

func (r *SmsListAppsResult) Get(ctx workflow.Context) (*sms.ListAppsOutput, error) {
    var output sms.ListAppsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsNotifyAppValidationOutputResult struct {
	Result workflow.Future
}

func (r *SmsNotifyAppValidationOutputResult) Get(ctx workflow.Context) (*sms.NotifyAppValidationOutputOutput, error) {
    var output sms.NotifyAppValidationOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsPutAppLaunchConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsPutAppLaunchConfigurationResult) Get(ctx workflow.Context) (*sms.PutAppLaunchConfigurationOutput, error) {
    var output sms.PutAppLaunchConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsPutAppReplicationConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsPutAppReplicationConfigurationResult) Get(ctx workflow.Context) (*sms.PutAppReplicationConfigurationOutput, error) {
    var output sms.PutAppReplicationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsPutAppValidationConfigurationResult struct {
	Result workflow.Future
}

func (r *SmsPutAppValidationConfigurationResult) Get(ctx workflow.Context) (*sms.PutAppValidationConfigurationOutput, error) {
    var output sms.PutAppValidationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsStartAppReplicationResult struct {
	Result workflow.Future
}

func (r *SmsStartAppReplicationResult) Get(ctx workflow.Context) (*sms.StartAppReplicationOutput, error) {
    var output sms.StartAppReplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsStartOnDemandAppReplicationResult struct {
	Result workflow.Future
}

func (r *SmsStartOnDemandAppReplicationResult) Get(ctx workflow.Context) (*sms.StartOnDemandAppReplicationOutput, error) {
    var output sms.StartOnDemandAppReplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsStartOnDemandReplicationRunResult struct {
	Result workflow.Future
}

func (r *SmsStartOnDemandReplicationRunResult) Get(ctx workflow.Context) (*sms.StartOnDemandReplicationRunOutput, error) {
    var output sms.StartOnDemandReplicationRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsStopAppReplicationResult struct {
	Result workflow.Future
}

func (r *SmsStopAppReplicationResult) Get(ctx workflow.Context) (*sms.StopAppReplicationOutput, error) {
    var output sms.StopAppReplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsTerminateAppResult struct {
	Result workflow.Future
}

func (r *SmsTerminateAppResult) Get(ctx workflow.Context) (*sms.TerminateAppOutput, error) {
    var output sms.TerminateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsUpdateAppResult struct {
	Result workflow.Future
}

func (r *SmsUpdateAppResult) Get(ctx workflow.Context) (*sms.UpdateAppOutput, error) {
    var output sms.UpdateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SmsUpdateReplicationJobResult struct {
	Result workflow.Future
}

func (r *SmsUpdateReplicationJobResult) Get(ctx workflow.Context) (*sms.UpdateReplicationJobOutput, error) {
    var output sms.UpdateReplicationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SMSStub struct {
    activities awsactivities.SMSActivities
}

func NewSMSStub() SMSClient {
    return &SMSStub{}
}

func (a *SMSStub) CreateApp(ctx workflow.Context, input *sms.CreateAppInput) (*sms.CreateAppOutput, error) {
    var output sms.CreateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) CreateAppAsync(ctx workflow.Context, input *sms.CreateAppInput) *SmsCreateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input)
    return &SmsCreateAppResult{Result: future}
}

func (a *SMSStub) CreateReplicationJob(ctx workflow.Context, input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
    var output sms.CreateReplicationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) CreateReplicationJobAsync(ctx workflow.Context, input *sms.CreateReplicationJobInput) *SmsCreateReplicationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationJob, input)
    return &SmsCreateReplicationJobResult{Result: future}
}

func (a *SMSStub) DeleteApp(ctx workflow.Context, input *sms.DeleteAppInput) (*sms.DeleteAppOutput, error) {
    var output sms.DeleteAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DeleteAppAsync(ctx workflow.Context, input *sms.DeleteAppInput) *SmsDeleteAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input)
    return &SmsDeleteAppResult{Result: future}
}

func (a *SMSStub) DeleteAppLaunchConfiguration(ctx workflow.Context, input *sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error) {
    var output sms.DeleteAppLaunchConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAppLaunchConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DeleteAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppLaunchConfigurationInput) *SmsDeleteAppLaunchConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAppLaunchConfiguration, input)
    return &SmsDeleteAppLaunchConfigurationResult{Result: future}
}

func (a *SMSStub) DeleteAppReplicationConfiguration(ctx workflow.Context, input *sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error) {
    var output sms.DeleteAppReplicationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAppReplicationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DeleteAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppReplicationConfigurationInput) *SmsDeleteAppReplicationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAppReplicationConfiguration, input)
    return &SmsDeleteAppReplicationConfigurationResult{Result: future}
}

func (a *SMSStub) DeleteAppValidationConfiguration(ctx workflow.Context, input *sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error) {
    var output sms.DeleteAppValidationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAppValidationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DeleteAppValidationConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppValidationConfigurationInput) *SmsDeleteAppValidationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAppValidationConfiguration, input)
    return &SmsDeleteAppValidationConfigurationResult{Result: future}
}

func (a *SMSStub) DeleteReplicationJob(ctx workflow.Context, input *sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error) {
    var output sms.DeleteReplicationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DeleteReplicationJobAsync(ctx workflow.Context, input *sms.DeleteReplicationJobInput) *SmsDeleteReplicationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationJob, input)
    return &SmsDeleteReplicationJobResult{Result: future}
}

func (a *SMSStub) DeleteServerCatalog(ctx workflow.Context, input *sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error) {
    var output sms.DeleteServerCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServerCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DeleteServerCatalogAsync(ctx workflow.Context, input *sms.DeleteServerCatalogInput) *SmsDeleteServerCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteServerCatalog, input)
    return &SmsDeleteServerCatalogResult{Result: future}
}

func (a *SMSStub) DisassociateConnector(ctx workflow.Context, input *sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error) {
    var output sms.DisassociateConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) DisassociateConnectorAsync(ctx workflow.Context, input *sms.DisassociateConnectorInput) *SmsDisassociateConnectorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateConnector, input)
    return &SmsDisassociateConnectorResult{Result: future}
}

func (a *SMSStub) GenerateChangeSet(ctx workflow.Context, input *sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error) {
    var output sms.GenerateChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GenerateChangeSetAsync(ctx workflow.Context, input *sms.GenerateChangeSetInput) *SmsGenerateChangeSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GenerateChangeSet, input)
    return &SmsGenerateChangeSetResult{Result: future}
}

func (a *SMSStub) GenerateTemplate(ctx workflow.Context, input *sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error) {
    var output sms.GenerateTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GenerateTemplateAsync(ctx workflow.Context, input *sms.GenerateTemplateInput) *SmsGenerateTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GenerateTemplate, input)
    return &SmsGenerateTemplateResult{Result: future}
}

func (a *SMSStub) GetApp(ctx workflow.Context, input *sms.GetAppInput) (*sms.GetAppOutput, error) {
    var output sms.GetAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetAppAsync(ctx workflow.Context, input *sms.GetAppInput) *SmsGetAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApp, input)
    return &SmsGetAppResult{Result: future}
}

func (a *SMSStub) GetAppLaunchConfiguration(ctx workflow.Context, input *sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error) {
    var output sms.GetAppLaunchConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAppLaunchConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.GetAppLaunchConfigurationInput) *SmsGetAppLaunchConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAppLaunchConfiguration, input)
    return &SmsGetAppLaunchConfigurationResult{Result: future}
}

func (a *SMSStub) GetAppReplicationConfiguration(ctx workflow.Context, input *sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error) {
    var output sms.GetAppReplicationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAppReplicationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.GetAppReplicationConfigurationInput) *SmsGetAppReplicationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAppReplicationConfiguration, input)
    return &SmsGetAppReplicationConfigurationResult{Result: future}
}

func (a *SMSStub) GetAppValidationConfiguration(ctx workflow.Context, input *sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error) {
    var output sms.GetAppValidationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAppValidationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetAppValidationConfigurationAsync(ctx workflow.Context, input *sms.GetAppValidationConfigurationInput) *SmsGetAppValidationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAppValidationConfiguration, input)
    return &SmsGetAppValidationConfigurationResult{Result: future}
}

func (a *SMSStub) GetAppValidationOutput(ctx workflow.Context, input *sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error) {
    var output sms.GetAppValidationOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAppValidationOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetAppValidationOutputAsync(ctx workflow.Context, input *sms.GetAppValidationOutputInput) *SmsGetAppValidationOutputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAppValidationOutput, input)
    return &SmsGetAppValidationOutputResult{Result: future}
}

func (a *SMSStub) GetConnectors(ctx workflow.Context, input *sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error) {
    var output sms.GetConnectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnectors, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetConnectorsAsync(ctx workflow.Context, input *sms.GetConnectorsInput) *SmsGetConnectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConnectors, input)
    return &SmsGetConnectorsResult{Result: future}
}

func (a *SMSStub) GetReplicationJobs(ctx workflow.Context, input *sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error) {
    var output sms.GetReplicationJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetReplicationJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetReplicationJobsAsync(ctx workflow.Context, input *sms.GetReplicationJobsInput) *SmsGetReplicationJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetReplicationJobs, input)
    return &SmsGetReplicationJobsResult{Result: future}
}

func (a *SMSStub) GetReplicationRuns(ctx workflow.Context, input *sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error) {
    var output sms.GetReplicationRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetReplicationRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetReplicationRunsAsync(ctx workflow.Context, input *sms.GetReplicationRunsInput) *SmsGetReplicationRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetReplicationRuns, input)
    return &SmsGetReplicationRunsResult{Result: future}
}

func (a *SMSStub) GetServers(ctx workflow.Context, input *sms.GetServersInput) (*sms.GetServersOutput, error) {
    var output sms.GetServersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServers, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) GetServersAsync(ctx workflow.Context, input *sms.GetServersInput) *SmsGetServersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetServers, input)
    return &SmsGetServersResult{Result: future}
}

func (a *SMSStub) ImportAppCatalog(ctx workflow.Context, input *sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error) {
    var output sms.ImportAppCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportAppCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) ImportAppCatalogAsync(ctx workflow.Context, input *sms.ImportAppCatalogInput) *SmsImportAppCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportAppCatalog, input)
    return &SmsImportAppCatalogResult{Result: future}
}

func (a *SMSStub) ImportServerCatalog(ctx workflow.Context, input *sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error) {
    var output sms.ImportServerCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportServerCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) ImportServerCatalogAsync(ctx workflow.Context, input *sms.ImportServerCatalogInput) *SmsImportServerCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportServerCatalog, input)
    return &SmsImportServerCatalogResult{Result: future}
}

func (a *SMSStub) LaunchApp(ctx workflow.Context, input *sms.LaunchAppInput) (*sms.LaunchAppOutput, error) {
    var output sms.LaunchAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.LaunchApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) LaunchAppAsync(ctx workflow.Context, input *sms.LaunchAppInput) *SmsLaunchAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.LaunchApp, input)
    return &SmsLaunchAppResult{Result: future}
}

func (a *SMSStub) ListApps(ctx workflow.Context, input *sms.ListAppsInput) (*sms.ListAppsOutput, error) {
    var output sms.ListAppsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApps, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) ListAppsAsync(ctx workflow.Context, input *sms.ListAppsInput) *SmsListAppsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListApps, input)
    return &SmsListAppsResult{Result: future}
}

func (a *SMSStub) NotifyAppValidationOutput(ctx workflow.Context, input *sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error) {
    var output sms.NotifyAppValidationOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.NotifyAppValidationOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) NotifyAppValidationOutputAsync(ctx workflow.Context, input *sms.NotifyAppValidationOutputInput) *SmsNotifyAppValidationOutputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.NotifyAppValidationOutput, input)
    return &SmsNotifyAppValidationOutputResult{Result: future}
}

func (a *SMSStub) PutAppLaunchConfiguration(ctx workflow.Context, input *sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error) {
    var output sms.PutAppLaunchConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAppLaunchConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) PutAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.PutAppLaunchConfigurationInput) *SmsPutAppLaunchConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAppLaunchConfiguration, input)
    return &SmsPutAppLaunchConfigurationResult{Result: future}
}

func (a *SMSStub) PutAppReplicationConfiguration(ctx workflow.Context, input *sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error) {
    var output sms.PutAppReplicationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAppReplicationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) PutAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.PutAppReplicationConfigurationInput) *SmsPutAppReplicationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAppReplicationConfiguration, input)
    return &SmsPutAppReplicationConfigurationResult{Result: future}
}

func (a *SMSStub) PutAppValidationConfiguration(ctx workflow.Context, input *sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error) {
    var output sms.PutAppValidationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAppValidationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) PutAppValidationConfigurationAsync(ctx workflow.Context, input *sms.PutAppValidationConfigurationInput) *SmsPutAppValidationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAppValidationConfiguration, input)
    return &SmsPutAppValidationConfigurationResult{Result: future}
}

func (a *SMSStub) StartAppReplication(ctx workflow.Context, input *sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error) {
    var output sms.StartAppReplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartAppReplication, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) StartAppReplicationAsync(ctx workflow.Context, input *sms.StartAppReplicationInput) *SmsStartAppReplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartAppReplication, input)
    return &SmsStartAppReplicationResult{Result: future}
}

func (a *SMSStub) StartOnDemandAppReplication(ctx workflow.Context, input *sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error) {
    var output sms.StartOnDemandAppReplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartOnDemandAppReplication, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) StartOnDemandAppReplicationAsync(ctx workflow.Context, input *sms.StartOnDemandAppReplicationInput) *SmsStartOnDemandAppReplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartOnDemandAppReplication, input)
    return &SmsStartOnDemandAppReplicationResult{Result: future}
}

func (a *SMSStub) StartOnDemandReplicationRun(ctx workflow.Context, input *sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error) {
    var output sms.StartOnDemandReplicationRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartOnDemandReplicationRun, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) StartOnDemandReplicationRunAsync(ctx workflow.Context, input *sms.StartOnDemandReplicationRunInput) *SmsStartOnDemandReplicationRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartOnDemandReplicationRun, input)
    return &SmsStartOnDemandReplicationRunResult{Result: future}
}

func (a *SMSStub) StopAppReplication(ctx workflow.Context, input *sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error) {
    var output sms.StopAppReplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopAppReplication, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) StopAppReplicationAsync(ctx workflow.Context, input *sms.StopAppReplicationInput) *SmsStopAppReplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopAppReplication, input)
    return &SmsStopAppReplicationResult{Result: future}
}

func (a *SMSStub) TerminateApp(ctx workflow.Context, input *sms.TerminateAppInput) (*sms.TerminateAppOutput, error) {
    var output sms.TerminateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) TerminateAppAsync(ctx workflow.Context, input *sms.TerminateAppInput) *SmsTerminateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateApp, input)
    return &SmsTerminateAppResult{Result: future}
}

func (a *SMSStub) UpdateApp(ctx workflow.Context, input *sms.UpdateAppInput) (*sms.UpdateAppOutput, error) {
    var output sms.UpdateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) UpdateAppAsync(ctx workflow.Context, input *sms.UpdateAppInput) *SmsUpdateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApp, input)
    return &SmsUpdateAppResult{Result: future}
}

func (a *SMSStub) UpdateReplicationJob(ctx workflow.Context, input *sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error) {
    var output sms.UpdateReplicationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateReplicationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SMSStub) UpdateReplicationJobAsync(ctx workflow.Context, input *sms.UpdateReplicationJobInput) *SmsUpdateReplicationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateReplicationJob, input)
    return &SmsUpdateReplicationJobResult{Result: future}
}