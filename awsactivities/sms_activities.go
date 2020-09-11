
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/sms/smsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SMSActivities struct {
    client smsiface.SMSAPI
}

func NewSMSActivities(session *session.Session, config ...*aws.Config) *SMSActivities {
    client := sms.New(session, config...)
    return &SMSActivities{client: client}
}

func (a *SMSActivities) CreateApp(ctx context.Context, input *sms.CreateAppInput) (*sms.CreateAppOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateAppWithContext(ctx, input)
}

func (a *SMSActivities) CreateReplicationJob(ctx context.Context, input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
    return a.client.CreateReplicationJobWithContext(ctx, input)
}

func (a *SMSActivities) DeleteApp(ctx context.Context, input *sms.DeleteAppInput) (*sms.DeleteAppOutput, error) {
    return a.client.DeleteAppWithContext(ctx, input)
}

func (a *SMSActivities) DeleteAppLaunchConfiguration(ctx context.Context, input *sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error) {
    return a.client.DeleteAppLaunchConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) DeleteAppReplicationConfiguration(ctx context.Context, input *sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error) {
    return a.client.DeleteAppReplicationConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) DeleteAppValidationConfiguration(ctx context.Context, input *sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error) {
    return a.client.DeleteAppValidationConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) DeleteReplicationJob(ctx context.Context, input *sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error) {
    return a.client.DeleteReplicationJobWithContext(ctx, input)
}

func (a *SMSActivities) DeleteServerCatalog(ctx context.Context, input *sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error) {
    return a.client.DeleteServerCatalogWithContext(ctx, input)
}

func (a *SMSActivities) DisassociateConnector(ctx context.Context, input *sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error) {
    return a.client.DisassociateConnectorWithContext(ctx, input)
}

func (a *SMSActivities) GenerateChangeSet(ctx context.Context, input *sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error) {
    return a.client.GenerateChangeSetWithContext(ctx, input)
}

func (a *SMSActivities) GenerateTemplate(ctx context.Context, input *sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error) {
    return a.client.GenerateTemplateWithContext(ctx, input)
}

func (a *SMSActivities) GetApp(ctx context.Context, input *sms.GetAppInput) (*sms.GetAppOutput, error) {
    return a.client.GetAppWithContext(ctx, input)
}

func (a *SMSActivities) GetAppLaunchConfiguration(ctx context.Context, input *sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error) {
    return a.client.GetAppLaunchConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) GetAppReplicationConfiguration(ctx context.Context, input *sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error) {
    return a.client.GetAppReplicationConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) GetAppValidationConfiguration(ctx context.Context, input *sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error) {
    return a.client.GetAppValidationConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) GetAppValidationOutput(ctx context.Context, input *sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error) {
    return a.client.GetAppValidationOutputWithContext(ctx, input)
}

func (a *SMSActivities) GetConnectors(ctx context.Context, input *sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error) {
    return a.client.GetConnectorsWithContext(ctx, input)
}

func (a *SMSActivities) GetReplicationJobs(ctx context.Context, input *sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error) {
    return a.client.GetReplicationJobsWithContext(ctx, input)
}

func (a *SMSActivities) GetReplicationRuns(ctx context.Context, input *sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error) {
    return a.client.GetReplicationRunsWithContext(ctx, input)
}

func (a *SMSActivities) GetServers(ctx context.Context, input *sms.GetServersInput) (*sms.GetServersOutput, error) {
    return a.client.GetServersWithContext(ctx, input)
}

func (a *SMSActivities) ImportAppCatalog(ctx context.Context, input *sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error) {
    return a.client.ImportAppCatalogWithContext(ctx, input)
}

func (a *SMSActivities) ImportServerCatalog(ctx context.Context, input *sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error) {
    return a.client.ImportServerCatalogWithContext(ctx, input)
}

func (a *SMSActivities) LaunchApp(ctx context.Context, input *sms.LaunchAppInput) (*sms.LaunchAppOutput, error) {
    return a.client.LaunchAppWithContext(ctx, input)
}

func (a *SMSActivities) ListApps(ctx context.Context, input *sms.ListAppsInput) (*sms.ListAppsOutput, error) {
    return a.client.ListAppsWithContext(ctx, input)
}

func (a *SMSActivities) NotifyAppValidationOutput(ctx context.Context, input *sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error) {
    return a.client.NotifyAppValidationOutputWithContext(ctx, input)
}

func (a *SMSActivities) PutAppLaunchConfiguration(ctx context.Context, input *sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error) {
    return a.client.PutAppLaunchConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) PutAppReplicationConfiguration(ctx context.Context, input *sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error) {
    return a.client.PutAppReplicationConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) PutAppValidationConfiguration(ctx context.Context, input *sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error) {
    return a.client.PutAppValidationConfigurationWithContext(ctx, input)
}

func (a *SMSActivities) StartAppReplication(ctx context.Context, input *sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error) {
    return a.client.StartAppReplicationWithContext(ctx, input)
}

func (a *SMSActivities) StartOnDemandAppReplication(ctx context.Context, input *sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error) {
    return a.client.StartOnDemandAppReplicationWithContext(ctx, input)
}

func (a *SMSActivities) StartOnDemandReplicationRun(ctx context.Context, input *sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error) {
    return a.client.StartOnDemandReplicationRunWithContext(ctx, input)
}

func (a *SMSActivities) StopAppReplication(ctx context.Context, input *sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error) {
    return a.client.StopAppReplicationWithContext(ctx, input)
}

func (a *SMSActivities) TerminateApp(ctx context.Context, input *sms.TerminateAppInput) (*sms.TerminateAppOutput, error) {
    return a.client.TerminateAppWithContext(ctx, input)
}

func (a *SMSActivities) UpdateApp(ctx context.Context, input *sms.UpdateAppInput) (*sms.UpdateAppOutput, error) {
    return a.client.UpdateAppWithContext(ctx, input)
}

func (a *SMSActivities) UpdateReplicationJob(ctx context.Context, input *sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error) {
    return a.client.UpdateReplicationJobWithContext(ctx, input)
}
