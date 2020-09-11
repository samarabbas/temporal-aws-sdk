
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DeviceFarmActivities struct {
    client devicefarmiface.DeviceFarmAPI
}

func NewDeviceFarmActivities(session *session.Session, config ...*aws.Config) *DeviceFarmActivities {
    client := devicefarm.New(session, config...)
    return &DeviceFarmActivities{client: client}
}

func (a *DeviceFarmActivities) CreateDevicePool(ctx context.Context, input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
    return a.client.CreateDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateInstanceProfile(ctx context.Context, input *devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error) {
    return a.client.CreateInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateNetworkProfile(ctx context.Context, input *devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
    return a.client.CreateNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateProject(ctx context.Context, input *devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
    return a.client.CreateProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateRemoteAccessSession(ctx context.Context, input *devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
    return a.client.CreateRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateTestGridProject(ctx context.Context, input *devicefarm.CreateTestGridProjectInput) (*devicefarm.CreateTestGridProjectOutput, error) {
    return a.client.CreateTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateTestGridUrl(ctx context.Context, input *devicefarm.CreateTestGridUrlInput) (*devicefarm.CreateTestGridUrlOutput, error) {
    return a.client.CreateTestGridUrlWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateUpload(ctx context.Context, input *devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
    return a.client.CreateUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateVPCEConfiguration(ctx context.Context, input *devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error) {
    return a.client.CreateVPCEConfigurationWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteDevicePool(ctx context.Context, input *devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
    return a.client.DeleteDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteInstanceProfile(ctx context.Context, input *devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error) {
    return a.client.DeleteInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteNetworkProfile(ctx context.Context, input *devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
    return a.client.DeleteNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteProject(ctx context.Context, input *devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
    return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteRemoteAccessSession(ctx context.Context, input *devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
    return a.client.DeleteRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteRun(ctx context.Context, input *devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
    return a.client.DeleteRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteTestGridProject(ctx context.Context, input *devicefarm.DeleteTestGridProjectInput) (*devicefarm.DeleteTestGridProjectOutput, error) {
    return a.client.DeleteTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteUpload(ctx context.Context, input *devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
    return a.client.DeleteUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteVPCEConfiguration(ctx context.Context, input *devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error) {
    return a.client.DeleteVPCEConfigurationWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetAccountSettings(ctx context.Context, input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
    return a.client.GetAccountSettingsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDevice(ctx context.Context, input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
    return a.client.GetDeviceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDeviceInstance(ctx context.Context, input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error) {
    return a.client.GetDeviceInstanceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDevicePool(ctx context.Context, input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
    return a.client.GetDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDevicePoolCompatibility(ctx context.Context, input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
    return a.client.GetDevicePoolCompatibilityWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetInstanceProfile(ctx context.Context, input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error) {
    return a.client.GetInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetJob(ctx context.Context, input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
    return a.client.GetJobWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetNetworkProfile(ctx context.Context, input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
    return a.client.GetNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetOfferingStatus(ctx context.Context, input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
    return a.client.GetOfferingStatusWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetProject(ctx context.Context, input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
    return a.client.GetProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetRemoteAccessSession(ctx context.Context, input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
    return a.client.GetRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetRun(ctx context.Context, input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
    return a.client.GetRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetSuite(ctx context.Context, input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
    return a.client.GetSuiteWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetTest(ctx context.Context, input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
    return a.client.GetTestWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetTestGridProject(ctx context.Context, input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error) {
    return a.client.GetTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetTestGridSession(ctx context.Context, input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error) {
    return a.client.GetTestGridSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetUpload(ctx context.Context, input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
    return a.client.GetUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetVPCEConfiguration(ctx context.Context, input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error) {
    return a.client.GetVPCEConfigurationWithContext(ctx, input)
}

func (a *DeviceFarmActivities) InstallToRemoteAccessSession(ctx context.Context, input *devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
    return a.client.InstallToRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListArtifacts(ctx context.Context, input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
    return a.client.ListArtifactsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListDeviceInstances(ctx context.Context, input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error) {
    return a.client.ListDeviceInstancesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListDevicePools(ctx context.Context, input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
    return a.client.ListDevicePoolsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListDevices(ctx context.Context, input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
    return a.client.ListDevicesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListInstanceProfiles(ctx context.Context, input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error) {
    return a.client.ListInstanceProfilesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListJobs(ctx context.Context, input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
    return a.client.ListJobsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListNetworkProfiles(ctx context.Context, input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
    return a.client.ListNetworkProfilesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListOfferingPromotions(ctx context.Context, input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
    return a.client.ListOfferingPromotionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListOfferingTransactions(ctx context.Context, input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
    return a.client.ListOfferingTransactionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListOfferings(ctx context.Context, input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
    return a.client.ListOfferingsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListProjects(ctx context.Context, input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
    return a.client.ListProjectsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListRemoteAccessSessions(ctx context.Context, input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
    return a.client.ListRemoteAccessSessionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListRuns(ctx context.Context, input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
    return a.client.ListRunsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListSamples(ctx context.Context, input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
    return a.client.ListSamplesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListSuites(ctx context.Context, input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
    return a.client.ListSuitesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTagsForResource(ctx context.Context, input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridProjects(ctx context.Context, input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error) {
    return a.client.ListTestGridProjectsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridSessionActions(ctx context.Context, input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error) {
    return a.client.ListTestGridSessionActionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridSessionArtifacts(ctx context.Context, input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
    return a.client.ListTestGridSessionArtifactsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridSessions(ctx context.Context, input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error) {
    return a.client.ListTestGridSessionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTests(ctx context.Context, input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
    return a.client.ListTestsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListUniqueProblems(ctx context.Context, input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
    return a.client.ListUniqueProblemsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListUploads(ctx context.Context, input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
    return a.client.ListUploadsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListVPCEConfigurations(ctx context.Context, input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error) {
    return a.client.ListVPCEConfigurationsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) PurchaseOffering(ctx context.Context, input *devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
    return a.client.PurchaseOfferingWithContext(ctx, input)
}

func (a *DeviceFarmActivities) RenewOffering(ctx context.Context, input *devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
    return a.client.RenewOfferingWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ScheduleRun(ctx context.Context, input *devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
    return a.client.ScheduleRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) StopJob(ctx context.Context, input *devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error) {
    return a.client.StopJobWithContext(ctx, input)
}

func (a *DeviceFarmActivities) StopRemoteAccessSession(ctx context.Context, input *devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
    return a.client.StopRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) StopRun(ctx context.Context, input *devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
    return a.client.StopRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) TagResource(ctx context.Context, input *devicefarm.TagResourceInput) (*devicefarm.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UntagResource(ctx context.Context, input *devicefarm.UntagResourceInput) (*devicefarm.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateDeviceInstance(ctx context.Context, input *devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error) {
    return a.client.UpdateDeviceInstanceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateDevicePool(ctx context.Context, input *devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
    return a.client.UpdateDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateInstanceProfile(ctx context.Context, input *devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error) {
    return a.client.UpdateInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateNetworkProfile(ctx context.Context, input *devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
    return a.client.UpdateNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateProject(ctx context.Context, input *devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
    return a.client.UpdateProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateTestGridProject(ctx context.Context, input *devicefarm.UpdateTestGridProjectInput) (*devicefarm.UpdateTestGridProjectOutput, error) {
    return a.client.UpdateTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateUpload(ctx context.Context, input *devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error) {
    return a.client.UpdateUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateVPCEConfiguration(ctx context.Context, input *devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error) {
    return a.client.UpdateVPCEConfigurationWithContext(ctx, input)
}
