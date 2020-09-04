package awsclients

import (
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DeviceFarmClient interface {
    CreateDevicePool(ctx workflow.Context, input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error)
    CreateDevicePoolAsync(ctx workflow.Context, input *devicefarm.CreateDevicePoolInput) *DevicefarmCreateDevicePoolResult

    CreateInstanceProfile(ctx workflow.Context, input *devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error)
    CreateInstanceProfileAsync(ctx workflow.Context, input *devicefarm.CreateInstanceProfileInput) *DevicefarmCreateInstanceProfileResult

    CreateNetworkProfile(ctx workflow.Context, input *devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error)
    CreateNetworkProfileAsync(ctx workflow.Context, input *devicefarm.CreateNetworkProfileInput) *DevicefarmCreateNetworkProfileResult

    CreateProject(ctx workflow.Context, input *devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error)
    CreateProjectAsync(ctx workflow.Context, input *devicefarm.CreateProjectInput) *DevicefarmCreateProjectResult

    CreateRemoteAccessSession(ctx workflow.Context, input *devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error)
    CreateRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.CreateRemoteAccessSessionInput) *DevicefarmCreateRemoteAccessSessionResult

    CreateTestGridProject(ctx workflow.Context, input *devicefarm.CreateTestGridProjectInput) (*devicefarm.CreateTestGridProjectOutput, error)
    CreateTestGridProjectAsync(ctx workflow.Context, input *devicefarm.CreateTestGridProjectInput) *DevicefarmCreateTestGridProjectResult

    CreateTestGridUrl(ctx workflow.Context, input *devicefarm.CreateTestGridUrlInput) (*devicefarm.CreateTestGridUrlOutput, error)
    CreateTestGridUrlAsync(ctx workflow.Context, input *devicefarm.CreateTestGridUrlInput) *DevicefarmCreateTestGridUrlResult

    CreateUpload(ctx workflow.Context, input *devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error)
    CreateUploadAsync(ctx workflow.Context, input *devicefarm.CreateUploadInput) *DevicefarmCreateUploadResult

    CreateVPCEConfiguration(ctx workflow.Context, input *devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error)
    CreateVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.CreateVPCEConfigurationInput) *DevicefarmCreateVPCEConfigurationResult

    DeleteDevicePool(ctx workflow.Context, input *devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error)
    DeleteDevicePoolAsync(ctx workflow.Context, input *devicefarm.DeleteDevicePoolInput) *DevicefarmDeleteDevicePoolResult

    DeleteInstanceProfile(ctx workflow.Context, input *devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error)
    DeleteInstanceProfileAsync(ctx workflow.Context, input *devicefarm.DeleteInstanceProfileInput) *DevicefarmDeleteInstanceProfileResult

    DeleteNetworkProfile(ctx workflow.Context, input *devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error)
    DeleteNetworkProfileAsync(ctx workflow.Context, input *devicefarm.DeleteNetworkProfileInput) *DevicefarmDeleteNetworkProfileResult

    DeleteProject(ctx workflow.Context, input *devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error)
    DeleteProjectAsync(ctx workflow.Context, input *devicefarm.DeleteProjectInput) *DevicefarmDeleteProjectResult

    DeleteRemoteAccessSession(ctx workflow.Context, input *devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error)
    DeleteRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.DeleteRemoteAccessSessionInput) *DevicefarmDeleteRemoteAccessSessionResult

    DeleteRun(ctx workflow.Context, input *devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error)
    DeleteRunAsync(ctx workflow.Context, input *devicefarm.DeleteRunInput) *DevicefarmDeleteRunResult

    DeleteTestGridProject(ctx workflow.Context, input *devicefarm.DeleteTestGridProjectInput) (*devicefarm.DeleteTestGridProjectOutput, error)
    DeleteTestGridProjectAsync(ctx workflow.Context, input *devicefarm.DeleteTestGridProjectInput) *DevicefarmDeleteTestGridProjectResult

    DeleteUpload(ctx workflow.Context, input *devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error)
    DeleteUploadAsync(ctx workflow.Context, input *devicefarm.DeleteUploadInput) *DevicefarmDeleteUploadResult

    DeleteVPCEConfiguration(ctx workflow.Context, input *devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error)
    DeleteVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.DeleteVPCEConfigurationInput) *DevicefarmDeleteVPCEConfigurationResult

    GetAccountSettings(ctx workflow.Context, input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error)
    GetAccountSettingsAsync(ctx workflow.Context, input *devicefarm.GetAccountSettingsInput) *DevicefarmGetAccountSettingsResult

    GetDevice(ctx workflow.Context, input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error)
    GetDeviceAsync(ctx workflow.Context, input *devicefarm.GetDeviceInput) *DevicefarmGetDeviceResult

    GetDeviceInstance(ctx workflow.Context, input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error)
    GetDeviceInstanceAsync(ctx workflow.Context, input *devicefarm.GetDeviceInstanceInput) *DevicefarmGetDeviceInstanceResult

    GetDevicePool(ctx workflow.Context, input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error)
    GetDevicePoolAsync(ctx workflow.Context, input *devicefarm.GetDevicePoolInput) *DevicefarmGetDevicePoolResult

    GetDevicePoolCompatibility(ctx workflow.Context, input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error)
    GetDevicePoolCompatibilityAsync(ctx workflow.Context, input *devicefarm.GetDevicePoolCompatibilityInput) *DevicefarmGetDevicePoolCompatibilityResult

    GetInstanceProfile(ctx workflow.Context, input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error)
    GetInstanceProfileAsync(ctx workflow.Context, input *devicefarm.GetInstanceProfileInput) *DevicefarmGetInstanceProfileResult

    GetJob(ctx workflow.Context, input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error)
    GetJobAsync(ctx workflow.Context, input *devicefarm.GetJobInput) *DevicefarmGetJobResult

    GetNetworkProfile(ctx workflow.Context, input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error)
    GetNetworkProfileAsync(ctx workflow.Context, input *devicefarm.GetNetworkProfileInput) *DevicefarmGetNetworkProfileResult

    GetOfferingStatus(ctx workflow.Context, input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error)
    GetOfferingStatusAsync(ctx workflow.Context, input *devicefarm.GetOfferingStatusInput) *DevicefarmGetOfferingStatusResult

    GetProject(ctx workflow.Context, input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error)
    GetProjectAsync(ctx workflow.Context, input *devicefarm.GetProjectInput) *DevicefarmGetProjectResult

    GetRemoteAccessSession(ctx workflow.Context, input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error)
    GetRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.GetRemoteAccessSessionInput) *DevicefarmGetRemoteAccessSessionResult

    GetRun(ctx workflow.Context, input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error)
    GetRunAsync(ctx workflow.Context, input *devicefarm.GetRunInput) *DevicefarmGetRunResult

    GetSuite(ctx workflow.Context, input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error)
    GetSuiteAsync(ctx workflow.Context, input *devicefarm.GetSuiteInput) *DevicefarmGetSuiteResult

    GetTest(ctx workflow.Context, input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error)
    GetTestAsync(ctx workflow.Context, input *devicefarm.GetTestInput) *DevicefarmGetTestResult

    GetTestGridProject(ctx workflow.Context, input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error)
    GetTestGridProjectAsync(ctx workflow.Context, input *devicefarm.GetTestGridProjectInput) *DevicefarmGetTestGridProjectResult

    GetTestGridSession(ctx workflow.Context, input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error)
    GetTestGridSessionAsync(ctx workflow.Context, input *devicefarm.GetTestGridSessionInput) *DevicefarmGetTestGridSessionResult

    GetUpload(ctx workflow.Context, input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error)
    GetUploadAsync(ctx workflow.Context, input *devicefarm.GetUploadInput) *DevicefarmGetUploadResult

    GetVPCEConfiguration(ctx workflow.Context, input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error)
    GetVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.GetVPCEConfigurationInput) *DevicefarmGetVPCEConfigurationResult

    InstallToRemoteAccessSession(ctx workflow.Context, input *devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error)
    InstallToRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.InstallToRemoteAccessSessionInput) *DevicefarmInstallToRemoteAccessSessionResult

    ListArtifacts(ctx workflow.Context, input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error)
    ListArtifactsAsync(ctx workflow.Context, input *devicefarm.ListArtifactsInput) *DevicefarmListArtifactsResult

    ListDeviceInstances(ctx workflow.Context, input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error)
    ListDeviceInstancesAsync(ctx workflow.Context, input *devicefarm.ListDeviceInstancesInput) *DevicefarmListDeviceInstancesResult

    ListDevicePools(ctx workflow.Context, input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error)
    ListDevicePoolsAsync(ctx workflow.Context, input *devicefarm.ListDevicePoolsInput) *DevicefarmListDevicePoolsResult

    ListDevices(ctx workflow.Context, input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error)
    ListDevicesAsync(ctx workflow.Context, input *devicefarm.ListDevicesInput) *DevicefarmListDevicesResult

    ListInstanceProfiles(ctx workflow.Context, input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error)
    ListInstanceProfilesAsync(ctx workflow.Context, input *devicefarm.ListInstanceProfilesInput) *DevicefarmListInstanceProfilesResult

    ListJobs(ctx workflow.Context, input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *devicefarm.ListJobsInput) *DevicefarmListJobsResult

    ListNetworkProfiles(ctx workflow.Context, input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error)
    ListNetworkProfilesAsync(ctx workflow.Context, input *devicefarm.ListNetworkProfilesInput) *DevicefarmListNetworkProfilesResult

    ListOfferingPromotions(ctx workflow.Context, input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error)
    ListOfferingPromotionsAsync(ctx workflow.Context, input *devicefarm.ListOfferingPromotionsInput) *DevicefarmListOfferingPromotionsResult

    ListOfferingTransactions(ctx workflow.Context, input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error)
    ListOfferingTransactionsAsync(ctx workflow.Context, input *devicefarm.ListOfferingTransactionsInput) *DevicefarmListOfferingTransactionsResult

    ListOfferings(ctx workflow.Context, input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error)
    ListOfferingsAsync(ctx workflow.Context, input *devicefarm.ListOfferingsInput) *DevicefarmListOfferingsResult

    ListProjects(ctx workflow.Context, input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error)
    ListProjectsAsync(ctx workflow.Context, input *devicefarm.ListProjectsInput) *DevicefarmListProjectsResult

    ListRemoteAccessSessions(ctx workflow.Context, input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error)
    ListRemoteAccessSessionsAsync(ctx workflow.Context, input *devicefarm.ListRemoteAccessSessionsInput) *DevicefarmListRemoteAccessSessionsResult

    ListRuns(ctx workflow.Context, input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error)
    ListRunsAsync(ctx workflow.Context, input *devicefarm.ListRunsInput) *DevicefarmListRunsResult

    ListSamples(ctx workflow.Context, input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error)
    ListSamplesAsync(ctx workflow.Context, input *devicefarm.ListSamplesInput) *DevicefarmListSamplesResult

    ListSuites(ctx workflow.Context, input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error)
    ListSuitesAsync(ctx workflow.Context, input *devicefarm.ListSuitesInput) *DevicefarmListSuitesResult

    ListTagsForResource(ctx workflow.Context, input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *devicefarm.ListTagsForResourceInput) *DevicefarmListTagsForResourceResult

    ListTestGridProjects(ctx workflow.Context, input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error)
    ListTestGridProjectsAsync(ctx workflow.Context, input *devicefarm.ListTestGridProjectsInput) *DevicefarmListTestGridProjectsResult

    ListTestGridSessionActions(ctx workflow.Context, input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error)
    ListTestGridSessionActionsAsync(ctx workflow.Context, input *devicefarm.ListTestGridSessionActionsInput) *DevicefarmListTestGridSessionActionsResult

    ListTestGridSessionArtifacts(ctx workflow.Context, input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error)
    ListTestGridSessionArtifactsAsync(ctx workflow.Context, input *devicefarm.ListTestGridSessionArtifactsInput) *DevicefarmListTestGridSessionArtifactsResult

    ListTestGridSessions(ctx workflow.Context, input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error)
    ListTestGridSessionsAsync(ctx workflow.Context, input *devicefarm.ListTestGridSessionsInput) *DevicefarmListTestGridSessionsResult

    ListTests(ctx workflow.Context, input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error)
    ListTestsAsync(ctx workflow.Context, input *devicefarm.ListTestsInput) *DevicefarmListTestsResult

    ListUniqueProblems(ctx workflow.Context, input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error)
    ListUniqueProblemsAsync(ctx workflow.Context, input *devicefarm.ListUniqueProblemsInput) *DevicefarmListUniqueProblemsResult

    ListUploads(ctx workflow.Context, input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error)
    ListUploadsAsync(ctx workflow.Context, input *devicefarm.ListUploadsInput) *DevicefarmListUploadsResult

    ListVPCEConfigurations(ctx workflow.Context, input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error)
    ListVPCEConfigurationsAsync(ctx workflow.Context, input *devicefarm.ListVPCEConfigurationsInput) *DevicefarmListVPCEConfigurationsResult

    PurchaseOffering(ctx workflow.Context, input *devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error)
    PurchaseOfferingAsync(ctx workflow.Context, input *devicefarm.PurchaseOfferingInput) *DevicefarmPurchaseOfferingResult

    RenewOffering(ctx workflow.Context, input *devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error)
    RenewOfferingAsync(ctx workflow.Context, input *devicefarm.RenewOfferingInput) *DevicefarmRenewOfferingResult

    ScheduleRun(ctx workflow.Context, input *devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error)
    ScheduleRunAsync(ctx workflow.Context, input *devicefarm.ScheduleRunInput) *DevicefarmScheduleRunResult

    StopJob(ctx workflow.Context, input *devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error)
    StopJobAsync(ctx workflow.Context, input *devicefarm.StopJobInput) *DevicefarmStopJobResult

    StopRemoteAccessSession(ctx workflow.Context, input *devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error)
    StopRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.StopRemoteAccessSessionInput) *DevicefarmStopRemoteAccessSessionResult

    StopRun(ctx workflow.Context, input *devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error)
    StopRunAsync(ctx workflow.Context, input *devicefarm.StopRunInput) *DevicefarmStopRunResult

    TagResource(ctx workflow.Context, input *devicefarm.TagResourceInput) (*devicefarm.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *devicefarm.TagResourceInput) *DevicefarmTagResourceResult

    UntagResource(ctx workflow.Context, input *devicefarm.UntagResourceInput) (*devicefarm.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *devicefarm.UntagResourceInput) *DevicefarmUntagResourceResult

    UpdateDeviceInstance(ctx workflow.Context, input *devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error)
    UpdateDeviceInstanceAsync(ctx workflow.Context, input *devicefarm.UpdateDeviceInstanceInput) *DevicefarmUpdateDeviceInstanceResult

    UpdateDevicePool(ctx workflow.Context, input *devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error)
    UpdateDevicePoolAsync(ctx workflow.Context, input *devicefarm.UpdateDevicePoolInput) *DevicefarmUpdateDevicePoolResult

    UpdateInstanceProfile(ctx workflow.Context, input *devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error)
    UpdateInstanceProfileAsync(ctx workflow.Context, input *devicefarm.UpdateInstanceProfileInput) *DevicefarmUpdateInstanceProfileResult

    UpdateNetworkProfile(ctx workflow.Context, input *devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error)
    UpdateNetworkProfileAsync(ctx workflow.Context, input *devicefarm.UpdateNetworkProfileInput) *DevicefarmUpdateNetworkProfileResult

    UpdateProject(ctx workflow.Context, input *devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error)
    UpdateProjectAsync(ctx workflow.Context, input *devicefarm.UpdateProjectInput) *DevicefarmUpdateProjectResult

    UpdateTestGridProject(ctx workflow.Context, input *devicefarm.UpdateTestGridProjectInput) (*devicefarm.UpdateTestGridProjectOutput, error)
    UpdateTestGridProjectAsync(ctx workflow.Context, input *devicefarm.UpdateTestGridProjectInput) *DevicefarmUpdateTestGridProjectResult

    UpdateUpload(ctx workflow.Context, input *devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error)
    UpdateUploadAsync(ctx workflow.Context, input *devicefarm.UpdateUploadInput) *DevicefarmUpdateUploadResult

    UpdateVPCEConfiguration(ctx workflow.Context, input *devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error)
    UpdateVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.UpdateVPCEConfigurationInput) *DevicefarmUpdateVPCEConfigurationResult
}

type DevicefarmCreateDevicePoolResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateDevicePoolResult) Get(ctx workflow.Context) (*devicefarm.CreateDevicePoolOutput, error) {
    var output devicefarm.CreateDevicePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateInstanceProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateInstanceProfileResult) Get(ctx workflow.Context) (*devicefarm.CreateInstanceProfileOutput, error) {
    var output devicefarm.CreateInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateNetworkProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateNetworkProfileResult) Get(ctx workflow.Context) (*devicefarm.CreateNetworkProfileOutput, error) {
    var output devicefarm.CreateNetworkProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateProjectResult) Get(ctx workflow.Context) (*devicefarm.CreateProjectOutput, error) {
    var output devicefarm.CreateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateRemoteAccessSessionResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateRemoteAccessSessionResult) Get(ctx workflow.Context) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
    var output devicefarm.CreateRemoteAccessSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateTestGridProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateTestGridProjectResult) Get(ctx workflow.Context) (*devicefarm.CreateTestGridProjectOutput, error) {
    var output devicefarm.CreateTestGridProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateTestGridUrlResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateTestGridUrlResult) Get(ctx workflow.Context) (*devicefarm.CreateTestGridUrlOutput, error) {
    var output devicefarm.CreateTestGridUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateUploadResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateUploadResult) Get(ctx workflow.Context) (*devicefarm.CreateUploadOutput, error) {
    var output devicefarm.CreateUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmCreateVPCEConfigurationResult struct {
	Result workflow.Future
}

func (r *DevicefarmCreateVPCEConfigurationResult) Get(ctx workflow.Context) (*devicefarm.CreateVPCEConfigurationOutput, error) {
    var output devicefarm.CreateVPCEConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteDevicePoolResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteDevicePoolResult) Get(ctx workflow.Context) (*devicefarm.DeleteDevicePoolOutput, error) {
    var output devicefarm.DeleteDevicePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteInstanceProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteInstanceProfileResult) Get(ctx workflow.Context) (*devicefarm.DeleteInstanceProfileOutput, error) {
    var output devicefarm.DeleteInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteNetworkProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteNetworkProfileResult) Get(ctx workflow.Context) (*devicefarm.DeleteNetworkProfileOutput, error) {
    var output devicefarm.DeleteNetworkProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteProjectResult) Get(ctx workflow.Context) (*devicefarm.DeleteProjectOutput, error) {
    var output devicefarm.DeleteProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteRemoteAccessSessionResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteRemoteAccessSessionResult) Get(ctx workflow.Context) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
    var output devicefarm.DeleteRemoteAccessSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteRunResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteRunResult) Get(ctx workflow.Context) (*devicefarm.DeleteRunOutput, error) {
    var output devicefarm.DeleteRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteTestGridProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteTestGridProjectResult) Get(ctx workflow.Context) (*devicefarm.DeleteTestGridProjectOutput, error) {
    var output devicefarm.DeleteTestGridProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteUploadResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteUploadResult) Get(ctx workflow.Context) (*devicefarm.DeleteUploadOutput, error) {
    var output devicefarm.DeleteUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmDeleteVPCEConfigurationResult struct {
	Result workflow.Future
}

func (r *DevicefarmDeleteVPCEConfigurationResult) Get(ctx workflow.Context) (*devicefarm.DeleteVPCEConfigurationOutput, error) {
    var output devicefarm.DeleteVPCEConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetAccountSettingsResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetAccountSettingsResult) Get(ctx workflow.Context) (*devicefarm.GetAccountSettingsOutput, error) {
    var output devicefarm.GetAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetDeviceResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetDeviceResult) Get(ctx workflow.Context) (*devicefarm.GetDeviceOutput, error) {
    var output devicefarm.GetDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetDeviceInstanceResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetDeviceInstanceResult) Get(ctx workflow.Context) (*devicefarm.GetDeviceInstanceOutput, error) {
    var output devicefarm.GetDeviceInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetDevicePoolResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetDevicePoolResult) Get(ctx workflow.Context) (*devicefarm.GetDevicePoolOutput, error) {
    var output devicefarm.GetDevicePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetDevicePoolCompatibilityResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetDevicePoolCompatibilityResult) Get(ctx workflow.Context) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
    var output devicefarm.GetDevicePoolCompatibilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetInstanceProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetInstanceProfileResult) Get(ctx workflow.Context) (*devicefarm.GetInstanceProfileOutput, error) {
    var output devicefarm.GetInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetJobResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetJobResult) Get(ctx workflow.Context) (*devicefarm.GetJobOutput, error) {
    var output devicefarm.GetJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetNetworkProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetNetworkProfileResult) Get(ctx workflow.Context) (*devicefarm.GetNetworkProfileOutput, error) {
    var output devicefarm.GetNetworkProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetOfferingStatusResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetOfferingStatusResult) Get(ctx workflow.Context) (*devicefarm.GetOfferingStatusOutput, error) {
    var output devicefarm.GetOfferingStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetProjectResult) Get(ctx workflow.Context) (*devicefarm.GetProjectOutput, error) {
    var output devicefarm.GetProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetRemoteAccessSessionResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetRemoteAccessSessionResult) Get(ctx workflow.Context) (*devicefarm.GetRemoteAccessSessionOutput, error) {
    var output devicefarm.GetRemoteAccessSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetRunResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetRunResult) Get(ctx workflow.Context) (*devicefarm.GetRunOutput, error) {
    var output devicefarm.GetRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetSuiteResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetSuiteResult) Get(ctx workflow.Context) (*devicefarm.GetSuiteOutput, error) {
    var output devicefarm.GetSuiteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetTestResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetTestResult) Get(ctx workflow.Context) (*devicefarm.GetTestOutput, error) {
    var output devicefarm.GetTestOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetTestGridProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetTestGridProjectResult) Get(ctx workflow.Context) (*devicefarm.GetTestGridProjectOutput, error) {
    var output devicefarm.GetTestGridProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetTestGridSessionResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetTestGridSessionResult) Get(ctx workflow.Context) (*devicefarm.GetTestGridSessionOutput, error) {
    var output devicefarm.GetTestGridSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetUploadResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetUploadResult) Get(ctx workflow.Context) (*devicefarm.GetUploadOutput, error) {
    var output devicefarm.GetUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmGetVPCEConfigurationResult struct {
	Result workflow.Future
}

func (r *DevicefarmGetVPCEConfigurationResult) Get(ctx workflow.Context) (*devicefarm.GetVPCEConfigurationOutput, error) {
    var output devicefarm.GetVPCEConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmInstallToRemoteAccessSessionResult struct {
	Result workflow.Future
}

func (r *DevicefarmInstallToRemoteAccessSessionResult) Get(ctx workflow.Context) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
    var output devicefarm.InstallToRemoteAccessSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListArtifactsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListArtifactsResult) Get(ctx workflow.Context) (*devicefarm.ListArtifactsOutput, error) {
    var output devicefarm.ListArtifactsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListDeviceInstancesResult struct {
	Result workflow.Future
}

func (r *DevicefarmListDeviceInstancesResult) Get(ctx workflow.Context) (*devicefarm.ListDeviceInstancesOutput, error) {
    var output devicefarm.ListDeviceInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListDevicePoolsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListDevicePoolsResult) Get(ctx workflow.Context) (*devicefarm.ListDevicePoolsOutput, error) {
    var output devicefarm.ListDevicePoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListDevicesResult struct {
	Result workflow.Future
}

func (r *DevicefarmListDevicesResult) Get(ctx workflow.Context) (*devicefarm.ListDevicesOutput, error) {
    var output devicefarm.ListDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListInstanceProfilesResult struct {
	Result workflow.Future
}

func (r *DevicefarmListInstanceProfilesResult) Get(ctx workflow.Context) (*devicefarm.ListInstanceProfilesOutput, error) {
    var output devicefarm.ListInstanceProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListJobsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListJobsResult) Get(ctx workflow.Context) (*devicefarm.ListJobsOutput, error) {
    var output devicefarm.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListNetworkProfilesResult struct {
	Result workflow.Future
}

func (r *DevicefarmListNetworkProfilesResult) Get(ctx workflow.Context) (*devicefarm.ListNetworkProfilesOutput, error) {
    var output devicefarm.ListNetworkProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListOfferingPromotionsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListOfferingPromotionsResult) Get(ctx workflow.Context) (*devicefarm.ListOfferingPromotionsOutput, error) {
    var output devicefarm.ListOfferingPromotionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListOfferingTransactionsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListOfferingTransactionsResult) Get(ctx workflow.Context) (*devicefarm.ListOfferingTransactionsOutput, error) {
    var output devicefarm.ListOfferingTransactionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListOfferingsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListOfferingsResult) Get(ctx workflow.Context) (*devicefarm.ListOfferingsOutput, error) {
    var output devicefarm.ListOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListProjectsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListProjectsResult) Get(ctx workflow.Context) (*devicefarm.ListProjectsOutput, error) {
    var output devicefarm.ListProjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListRemoteAccessSessionsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListRemoteAccessSessionsResult) Get(ctx workflow.Context) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
    var output devicefarm.ListRemoteAccessSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListRunsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListRunsResult) Get(ctx workflow.Context) (*devicefarm.ListRunsOutput, error) {
    var output devicefarm.ListRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListSamplesResult struct {
	Result workflow.Future
}

func (r *DevicefarmListSamplesResult) Get(ctx workflow.Context) (*devicefarm.ListSamplesOutput, error) {
    var output devicefarm.ListSamplesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListSuitesResult struct {
	Result workflow.Future
}

func (r *DevicefarmListSuitesResult) Get(ctx workflow.Context) (*devicefarm.ListSuitesOutput, error) {
    var output devicefarm.ListSuitesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DevicefarmListTagsForResourceResult) Get(ctx workflow.Context) (*devicefarm.ListTagsForResourceOutput, error) {
    var output devicefarm.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListTestGridProjectsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListTestGridProjectsResult) Get(ctx workflow.Context) (*devicefarm.ListTestGridProjectsOutput, error) {
    var output devicefarm.ListTestGridProjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListTestGridSessionActionsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListTestGridSessionActionsResult) Get(ctx workflow.Context) (*devicefarm.ListTestGridSessionActionsOutput, error) {
    var output devicefarm.ListTestGridSessionActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListTestGridSessionArtifactsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListTestGridSessionArtifactsResult) Get(ctx workflow.Context) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
    var output devicefarm.ListTestGridSessionArtifactsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListTestGridSessionsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListTestGridSessionsResult) Get(ctx workflow.Context) (*devicefarm.ListTestGridSessionsOutput, error) {
    var output devicefarm.ListTestGridSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListTestsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListTestsResult) Get(ctx workflow.Context) (*devicefarm.ListTestsOutput, error) {
    var output devicefarm.ListTestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListUniqueProblemsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListUniqueProblemsResult) Get(ctx workflow.Context) (*devicefarm.ListUniqueProblemsOutput, error) {
    var output devicefarm.ListUniqueProblemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListUploadsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListUploadsResult) Get(ctx workflow.Context) (*devicefarm.ListUploadsOutput, error) {
    var output devicefarm.ListUploadsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmListVPCEConfigurationsResult struct {
	Result workflow.Future
}

func (r *DevicefarmListVPCEConfigurationsResult) Get(ctx workflow.Context) (*devicefarm.ListVPCEConfigurationsOutput, error) {
    var output devicefarm.ListVPCEConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmPurchaseOfferingResult struct {
	Result workflow.Future
}

func (r *DevicefarmPurchaseOfferingResult) Get(ctx workflow.Context) (*devicefarm.PurchaseOfferingOutput, error) {
    var output devicefarm.PurchaseOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmRenewOfferingResult struct {
	Result workflow.Future
}

func (r *DevicefarmRenewOfferingResult) Get(ctx workflow.Context) (*devicefarm.RenewOfferingOutput, error) {
    var output devicefarm.RenewOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmScheduleRunResult struct {
	Result workflow.Future
}

func (r *DevicefarmScheduleRunResult) Get(ctx workflow.Context) (*devicefarm.ScheduleRunOutput, error) {
    var output devicefarm.ScheduleRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmStopJobResult struct {
	Result workflow.Future
}

func (r *DevicefarmStopJobResult) Get(ctx workflow.Context) (*devicefarm.StopJobOutput, error) {
    var output devicefarm.StopJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmStopRemoteAccessSessionResult struct {
	Result workflow.Future
}

func (r *DevicefarmStopRemoteAccessSessionResult) Get(ctx workflow.Context) (*devicefarm.StopRemoteAccessSessionOutput, error) {
    var output devicefarm.StopRemoteAccessSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmStopRunResult struct {
	Result workflow.Future
}

func (r *DevicefarmStopRunResult) Get(ctx workflow.Context) (*devicefarm.StopRunOutput, error) {
    var output devicefarm.StopRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmTagResourceResult struct {
	Result workflow.Future
}

func (r *DevicefarmTagResourceResult) Get(ctx workflow.Context) (*devicefarm.TagResourceOutput, error) {
    var output devicefarm.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUntagResourceResult struct {
	Result workflow.Future
}

func (r *DevicefarmUntagResourceResult) Get(ctx workflow.Context) (*devicefarm.UntagResourceOutput, error) {
    var output devicefarm.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateDeviceInstanceResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateDeviceInstanceResult) Get(ctx workflow.Context) (*devicefarm.UpdateDeviceInstanceOutput, error) {
    var output devicefarm.UpdateDeviceInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateDevicePoolResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateDevicePoolResult) Get(ctx workflow.Context) (*devicefarm.UpdateDevicePoolOutput, error) {
    var output devicefarm.UpdateDevicePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateInstanceProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateInstanceProfileResult) Get(ctx workflow.Context) (*devicefarm.UpdateInstanceProfileOutput, error) {
    var output devicefarm.UpdateInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateNetworkProfileResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateNetworkProfileResult) Get(ctx workflow.Context) (*devicefarm.UpdateNetworkProfileOutput, error) {
    var output devicefarm.UpdateNetworkProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateProjectResult) Get(ctx workflow.Context) (*devicefarm.UpdateProjectOutput, error) {
    var output devicefarm.UpdateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateTestGridProjectResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateTestGridProjectResult) Get(ctx workflow.Context) (*devicefarm.UpdateTestGridProjectOutput, error) {
    var output devicefarm.UpdateTestGridProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateUploadResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateUploadResult) Get(ctx workflow.Context) (*devicefarm.UpdateUploadOutput, error) {
    var output devicefarm.UpdateUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DevicefarmUpdateVPCEConfigurationResult struct {
	Result workflow.Future
}

func (r *DevicefarmUpdateVPCEConfigurationResult) Get(ctx workflow.Context) (*devicefarm.UpdateVPCEConfigurationOutput, error) {
    var output devicefarm.UpdateVPCEConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DeviceFarmStub struct {
    activities awsactivities.DeviceFarmActivities
}

func NewDeviceFarmStub() DeviceFarmClient {
    return &DeviceFarmStub{}
}

func (a *DeviceFarmStub) CreateDevicePool(ctx workflow.Context, input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
    var output devicefarm.CreateDevicePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDevicePool, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateDevicePoolAsync(ctx workflow.Context, input *devicefarm.CreateDevicePoolInput) *DevicefarmCreateDevicePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDevicePool, input)
    return &DevicefarmCreateDevicePoolResult{Result: future}
}

func (a *DeviceFarmStub) CreateInstanceProfile(ctx workflow.Context, input *devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error) {
    var output devicefarm.CreateInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateInstanceProfileAsync(ctx workflow.Context, input *devicefarm.CreateInstanceProfileInput) *DevicefarmCreateInstanceProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceProfile, input)
    return &DevicefarmCreateInstanceProfileResult{Result: future}
}

func (a *DeviceFarmStub) CreateNetworkProfile(ctx workflow.Context, input *devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
    var output devicefarm.CreateNetworkProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateNetworkProfileAsync(ctx workflow.Context, input *devicefarm.CreateNetworkProfileInput) *DevicefarmCreateNetworkProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkProfile, input)
    return &DevicefarmCreateNetworkProfileResult{Result: future}
}

func (a *DeviceFarmStub) CreateProject(ctx workflow.Context, input *devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
    var output devicefarm.CreateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateProjectAsync(ctx workflow.Context, input *devicefarm.CreateProjectInput) *DevicefarmCreateProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input)
    return &DevicefarmCreateProjectResult{Result: future}
}

func (a *DeviceFarmStub) CreateRemoteAccessSession(ctx workflow.Context, input *devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
    var output devicefarm.CreateRemoteAccessSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRemoteAccessSession, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.CreateRemoteAccessSessionInput) *DevicefarmCreateRemoteAccessSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRemoteAccessSession, input)
    return &DevicefarmCreateRemoteAccessSessionResult{Result: future}
}

func (a *DeviceFarmStub) CreateTestGridProject(ctx workflow.Context, input *devicefarm.CreateTestGridProjectInput) (*devicefarm.CreateTestGridProjectOutput, error) {
    var output devicefarm.CreateTestGridProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTestGridProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateTestGridProjectAsync(ctx workflow.Context, input *devicefarm.CreateTestGridProjectInput) *DevicefarmCreateTestGridProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTestGridProject, input)
    return &DevicefarmCreateTestGridProjectResult{Result: future}
}

func (a *DeviceFarmStub) CreateTestGridUrl(ctx workflow.Context, input *devicefarm.CreateTestGridUrlInput) (*devicefarm.CreateTestGridUrlOutput, error) {
    var output devicefarm.CreateTestGridUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTestGridUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateTestGridUrlAsync(ctx workflow.Context, input *devicefarm.CreateTestGridUrlInput) *DevicefarmCreateTestGridUrlResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTestGridUrl, input)
    return &DevicefarmCreateTestGridUrlResult{Result: future}
}

func (a *DeviceFarmStub) CreateUpload(ctx workflow.Context, input *devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
    var output devicefarm.CreateUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateUploadAsync(ctx workflow.Context, input *devicefarm.CreateUploadInput) *DevicefarmCreateUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUpload, input)
    return &DevicefarmCreateUploadResult{Result: future}
}

func (a *DeviceFarmStub) CreateVPCEConfiguration(ctx workflow.Context, input *devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error) {
    var output devicefarm.CreateVPCEConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVPCEConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) CreateVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.CreateVPCEConfigurationInput) *DevicefarmCreateVPCEConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVPCEConfiguration, input)
    return &DevicefarmCreateVPCEConfigurationResult{Result: future}
}

func (a *DeviceFarmStub) DeleteDevicePool(ctx workflow.Context, input *devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
    var output devicefarm.DeleteDevicePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDevicePool, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteDevicePoolAsync(ctx workflow.Context, input *devicefarm.DeleteDevicePoolInput) *DevicefarmDeleteDevicePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDevicePool, input)
    return &DevicefarmDeleteDevicePoolResult{Result: future}
}

func (a *DeviceFarmStub) DeleteInstanceProfile(ctx workflow.Context, input *devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error) {
    var output devicefarm.DeleteInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteInstanceProfileAsync(ctx workflow.Context, input *devicefarm.DeleteInstanceProfileInput) *DevicefarmDeleteInstanceProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInstanceProfile, input)
    return &DevicefarmDeleteInstanceProfileResult{Result: future}
}

func (a *DeviceFarmStub) DeleteNetworkProfile(ctx workflow.Context, input *devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
    var output devicefarm.DeleteNetworkProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteNetworkProfileAsync(ctx workflow.Context, input *devicefarm.DeleteNetworkProfileInput) *DevicefarmDeleteNetworkProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkProfile, input)
    return &DevicefarmDeleteNetworkProfileResult{Result: future}
}

func (a *DeviceFarmStub) DeleteProject(ctx workflow.Context, input *devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
    var output devicefarm.DeleteProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteProjectAsync(ctx workflow.Context, input *devicefarm.DeleteProjectInput) *DevicefarmDeleteProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input)
    return &DevicefarmDeleteProjectResult{Result: future}
}

func (a *DeviceFarmStub) DeleteRemoteAccessSession(ctx workflow.Context, input *devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
    var output devicefarm.DeleteRemoteAccessSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRemoteAccessSession, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.DeleteRemoteAccessSessionInput) *DevicefarmDeleteRemoteAccessSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRemoteAccessSession, input)
    return &DevicefarmDeleteRemoteAccessSessionResult{Result: future}
}

func (a *DeviceFarmStub) DeleteRun(ctx workflow.Context, input *devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
    var output devicefarm.DeleteRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteRunAsync(ctx workflow.Context, input *devicefarm.DeleteRunInput) *DevicefarmDeleteRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRun, input)
    return &DevicefarmDeleteRunResult{Result: future}
}

func (a *DeviceFarmStub) DeleteTestGridProject(ctx workflow.Context, input *devicefarm.DeleteTestGridProjectInput) (*devicefarm.DeleteTestGridProjectOutput, error) {
    var output devicefarm.DeleteTestGridProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTestGridProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteTestGridProjectAsync(ctx workflow.Context, input *devicefarm.DeleteTestGridProjectInput) *DevicefarmDeleteTestGridProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTestGridProject, input)
    return &DevicefarmDeleteTestGridProjectResult{Result: future}
}

func (a *DeviceFarmStub) DeleteUpload(ctx workflow.Context, input *devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
    var output devicefarm.DeleteUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteUploadAsync(ctx workflow.Context, input *devicefarm.DeleteUploadInput) *DevicefarmDeleteUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUpload, input)
    return &DevicefarmDeleteUploadResult{Result: future}
}

func (a *DeviceFarmStub) DeleteVPCEConfiguration(ctx workflow.Context, input *devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error) {
    var output devicefarm.DeleteVPCEConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVPCEConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) DeleteVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.DeleteVPCEConfigurationInput) *DevicefarmDeleteVPCEConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVPCEConfiguration, input)
    return &DevicefarmDeleteVPCEConfigurationResult{Result: future}
}

func (a *DeviceFarmStub) GetAccountSettings(ctx workflow.Context, input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
    var output devicefarm.GetAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetAccountSettingsAsync(ctx workflow.Context, input *devicefarm.GetAccountSettingsInput) *DevicefarmGetAccountSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccountSettings, input)
    return &DevicefarmGetAccountSettingsResult{Result: future}
}

func (a *DeviceFarmStub) GetDevice(ctx workflow.Context, input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
    var output devicefarm.GetDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetDeviceAsync(ctx workflow.Context, input *devicefarm.GetDeviceInput) *DevicefarmGetDeviceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input)
    return &DevicefarmGetDeviceResult{Result: future}
}

func (a *DeviceFarmStub) GetDeviceInstance(ctx workflow.Context, input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error) {
    var output devicefarm.GetDeviceInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeviceInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetDeviceInstanceAsync(ctx workflow.Context, input *devicefarm.GetDeviceInstanceInput) *DevicefarmGetDeviceInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeviceInstance, input)
    return &DevicefarmGetDeviceInstanceResult{Result: future}
}

func (a *DeviceFarmStub) GetDevicePool(ctx workflow.Context, input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
    var output devicefarm.GetDevicePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevicePool, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetDevicePoolAsync(ctx workflow.Context, input *devicefarm.GetDevicePoolInput) *DevicefarmGetDevicePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevicePool, input)
    return &DevicefarmGetDevicePoolResult{Result: future}
}

func (a *DeviceFarmStub) GetDevicePoolCompatibility(ctx workflow.Context, input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
    var output devicefarm.GetDevicePoolCompatibilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevicePoolCompatibility, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetDevicePoolCompatibilityAsync(ctx workflow.Context, input *devicefarm.GetDevicePoolCompatibilityInput) *DevicefarmGetDevicePoolCompatibilityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevicePoolCompatibility, input)
    return &DevicefarmGetDevicePoolCompatibilityResult{Result: future}
}

func (a *DeviceFarmStub) GetInstanceProfile(ctx workflow.Context, input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error) {
    var output devicefarm.GetInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetInstanceProfileAsync(ctx workflow.Context, input *devicefarm.GetInstanceProfileInput) *DevicefarmGetInstanceProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceProfile, input)
    return &DevicefarmGetInstanceProfileResult{Result: future}
}

func (a *DeviceFarmStub) GetJob(ctx workflow.Context, input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
    var output devicefarm.GetJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJob, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetJobAsync(ctx workflow.Context, input *devicefarm.GetJobInput) *DevicefarmGetJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJob, input)
    return &DevicefarmGetJobResult{Result: future}
}

func (a *DeviceFarmStub) GetNetworkProfile(ctx workflow.Context, input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
    var output devicefarm.GetNetworkProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNetworkProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetNetworkProfileAsync(ctx workflow.Context, input *devicefarm.GetNetworkProfileInput) *DevicefarmGetNetworkProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetNetworkProfile, input)
    return &DevicefarmGetNetworkProfileResult{Result: future}
}

func (a *DeviceFarmStub) GetOfferingStatus(ctx workflow.Context, input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
    var output devicefarm.GetOfferingStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOfferingStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetOfferingStatusAsync(ctx workflow.Context, input *devicefarm.GetOfferingStatusInput) *DevicefarmGetOfferingStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOfferingStatus, input)
    return &DevicefarmGetOfferingStatusResult{Result: future}
}

func (a *DeviceFarmStub) GetProject(ctx workflow.Context, input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
    var output devicefarm.GetProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetProjectAsync(ctx workflow.Context, input *devicefarm.GetProjectInput) *DevicefarmGetProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetProject, input)
    return &DevicefarmGetProjectResult{Result: future}
}

func (a *DeviceFarmStub) GetRemoteAccessSession(ctx workflow.Context, input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
    var output devicefarm.GetRemoteAccessSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRemoteAccessSession, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.GetRemoteAccessSessionInput) *DevicefarmGetRemoteAccessSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRemoteAccessSession, input)
    return &DevicefarmGetRemoteAccessSessionResult{Result: future}
}

func (a *DeviceFarmStub) GetRun(ctx workflow.Context, input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
    var output devicefarm.GetRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetRunAsync(ctx workflow.Context, input *devicefarm.GetRunInput) *DevicefarmGetRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRun, input)
    return &DevicefarmGetRunResult{Result: future}
}

func (a *DeviceFarmStub) GetSuite(ctx workflow.Context, input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
    var output devicefarm.GetSuiteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSuite, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetSuiteAsync(ctx workflow.Context, input *devicefarm.GetSuiteInput) *DevicefarmGetSuiteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSuite, input)
    return &DevicefarmGetSuiteResult{Result: future}
}

func (a *DeviceFarmStub) GetTest(ctx workflow.Context, input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
    var output devicefarm.GetTestOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTest, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetTestAsync(ctx workflow.Context, input *devicefarm.GetTestInput) *DevicefarmGetTestResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTest, input)
    return &DevicefarmGetTestResult{Result: future}
}

func (a *DeviceFarmStub) GetTestGridProject(ctx workflow.Context, input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error) {
    var output devicefarm.GetTestGridProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTestGridProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetTestGridProjectAsync(ctx workflow.Context, input *devicefarm.GetTestGridProjectInput) *DevicefarmGetTestGridProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTestGridProject, input)
    return &DevicefarmGetTestGridProjectResult{Result: future}
}

func (a *DeviceFarmStub) GetTestGridSession(ctx workflow.Context, input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error) {
    var output devicefarm.GetTestGridSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTestGridSession, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetTestGridSessionAsync(ctx workflow.Context, input *devicefarm.GetTestGridSessionInput) *DevicefarmGetTestGridSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTestGridSession, input)
    return &DevicefarmGetTestGridSessionResult{Result: future}
}

func (a *DeviceFarmStub) GetUpload(ctx workflow.Context, input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
    var output devicefarm.GetUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetUploadAsync(ctx workflow.Context, input *devicefarm.GetUploadInput) *DevicefarmGetUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUpload, input)
    return &DevicefarmGetUploadResult{Result: future}
}

func (a *DeviceFarmStub) GetVPCEConfiguration(ctx workflow.Context, input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error) {
    var output devicefarm.GetVPCEConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVPCEConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) GetVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.GetVPCEConfigurationInput) *DevicefarmGetVPCEConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVPCEConfiguration, input)
    return &DevicefarmGetVPCEConfigurationResult{Result: future}
}

func (a *DeviceFarmStub) InstallToRemoteAccessSession(ctx workflow.Context, input *devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
    var output devicefarm.InstallToRemoteAccessSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InstallToRemoteAccessSession, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) InstallToRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.InstallToRemoteAccessSessionInput) *DevicefarmInstallToRemoteAccessSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InstallToRemoteAccessSession, input)
    return &DevicefarmInstallToRemoteAccessSessionResult{Result: future}
}

func (a *DeviceFarmStub) ListArtifacts(ctx workflow.Context, input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
    var output devicefarm.ListArtifactsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListArtifacts, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListArtifactsAsync(ctx workflow.Context, input *devicefarm.ListArtifactsInput) *DevicefarmListArtifactsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListArtifacts, input)
    return &DevicefarmListArtifactsResult{Result: future}
}

func (a *DeviceFarmStub) ListDeviceInstances(ctx workflow.Context, input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error) {
    var output devicefarm.ListDeviceInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeviceInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListDeviceInstancesAsync(ctx workflow.Context, input *devicefarm.ListDeviceInstancesInput) *DevicefarmListDeviceInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeviceInstances, input)
    return &DevicefarmListDeviceInstancesResult{Result: future}
}

func (a *DeviceFarmStub) ListDevicePools(ctx workflow.Context, input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
    var output devicefarm.ListDevicePoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevicePools, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListDevicePoolsAsync(ctx workflow.Context, input *devicefarm.ListDevicePoolsInput) *DevicefarmListDevicePoolsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDevicePools, input)
    return &DevicefarmListDevicePoolsResult{Result: future}
}

func (a *DeviceFarmStub) ListDevices(ctx workflow.Context, input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
    var output devicefarm.ListDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListDevicesAsync(ctx workflow.Context, input *devicefarm.ListDevicesInput) *DevicefarmListDevicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDevices, input)
    return &DevicefarmListDevicesResult{Result: future}
}

func (a *DeviceFarmStub) ListInstanceProfiles(ctx workflow.Context, input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error) {
    var output devicefarm.ListInstanceProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstanceProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListInstanceProfilesAsync(ctx workflow.Context, input *devicefarm.ListInstanceProfilesInput) *DevicefarmListInstanceProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListInstanceProfiles, input)
    return &DevicefarmListInstanceProfilesResult{Result: future}
}

func (a *DeviceFarmStub) ListJobs(ctx workflow.Context, input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
    var output devicefarm.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListJobsAsync(ctx workflow.Context, input *devicefarm.ListJobsInput) *DevicefarmListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &DevicefarmListJobsResult{Result: future}
}

func (a *DeviceFarmStub) ListNetworkProfiles(ctx workflow.Context, input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
    var output devicefarm.ListNetworkProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNetworkProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListNetworkProfilesAsync(ctx workflow.Context, input *devicefarm.ListNetworkProfilesInput) *DevicefarmListNetworkProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListNetworkProfiles, input)
    return &DevicefarmListNetworkProfilesResult{Result: future}
}

func (a *DeviceFarmStub) ListOfferingPromotions(ctx workflow.Context, input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
    var output devicefarm.ListOfferingPromotionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOfferingPromotions, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListOfferingPromotionsAsync(ctx workflow.Context, input *devicefarm.ListOfferingPromotionsInput) *DevicefarmListOfferingPromotionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOfferingPromotions, input)
    return &DevicefarmListOfferingPromotionsResult{Result: future}
}

func (a *DeviceFarmStub) ListOfferingTransactions(ctx workflow.Context, input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
    var output devicefarm.ListOfferingTransactionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOfferingTransactions, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListOfferingTransactionsAsync(ctx workflow.Context, input *devicefarm.ListOfferingTransactionsInput) *DevicefarmListOfferingTransactionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOfferingTransactions, input)
    return &DevicefarmListOfferingTransactionsResult{Result: future}
}

func (a *DeviceFarmStub) ListOfferings(ctx workflow.Context, input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
    var output devicefarm.ListOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListOfferingsAsync(ctx workflow.Context, input *devicefarm.ListOfferingsInput) *DevicefarmListOfferingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOfferings, input)
    return &DevicefarmListOfferingsResult{Result: future}
}

func (a *DeviceFarmStub) ListProjects(ctx workflow.Context, input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
    var output devicefarm.ListProjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListProjectsAsync(ctx workflow.Context, input *devicefarm.ListProjectsInput) *DevicefarmListProjectsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input)
    return &DevicefarmListProjectsResult{Result: future}
}

func (a *DeviceFarmStub) ListRemoteAccessSessions(ctx workflow.Context, input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
    var output devicefarm.ListRemoteAccessSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRemoteAccessSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListRemoteAccessSessionsAsync(ctx workflow.Context, input *devicefarm.ListRemoteAccessSessionsInput) *DevicefarmListRemoteAccessSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRemoteAccessSessions, input)
    return &DevicefarmListRemoteAccessSessionsResult{Result: future}
}

func (a *DeviceFarmStub) ListRuns(ctx workflow.Context, input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
    var output devicefarm.ListRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListRunsAsync(ctx workflow.Context, input *devicefarm.ListRunsInput) *DevicefarmListRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRuns, input)
    return &DevicefarmListRunsResult{Result: future}
}

func (a *DeviceFarmStub) ListSamples(ctx workflow.Context, input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
    var output devicefarm.ListSamplesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSamples, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListSamplesAsync(ctx workflow.Context, input *devicefarm.ListSamplesInput) *DevicefarmListSamplesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSamples, input)
    return &DevicefarmListSamplesResult{Result: future}
}

func (a *DeviceFarmStub) ListSuites(ctx workflow.Context, input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
    var output devicefarm.ListSuitesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSuites, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListSuitesAsync(ctx workflow.Context, input *devicefarm.ListSuitesInput) *DevicefarmListSuitesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSuites, input)
    return &DevicefarmListSuitesResult{Result: future}
}

func (a *DeviceFarmStub) ListTagsForResource(ctx workflow.Context, input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error) {
    var output devicefarm.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListTagsForResourceAsync(ctx workflow.Context, input *devicefarm.ListTagsForResourceInput) *DevicefarmListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &DevicefarmListTagsForResourceResult{Result: future}
}

func (a *DeviceFarmStub) ListTestGridProjects(ctx workflow.Context, input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error) {
    var output devicefarm.ListTestGridProjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTestGridProjects, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListTestGridProjectsAsync(ctx workflow.Context, input *devicefarm.ListTestGridProjectsInput) *DevicefarmListTestGridProjectsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTestGridProjects, input)
    return &DevicefarmListTestGridProjectsResult{Result: future}
}

func (a *DeviceFarmStub) ListTestGridSessionActions(ctx workflow.Context, input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error) {
    var output devicefarm.ListTestGridSessionActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTestGridSessionActions, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListTestGridSessionActionsAsync(ctx workflow.Context, input *devicefarm.ListTestGridSessionActionsInput) *DevicefarmListTestGridSessionActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTestGridSessionActions, input)
    return &DevicefarmListTestGridSessionActionsResult{Result: future}
}

func (a *DeviceFarmStub) ListTestGridSessionArtifacts(ctx workflow.Context, input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
    var output devicefarm.ListTestGridSessionArtifactsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTestGridSessionArtifacts, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListTestGridSessionArtifactsAsync(ctx workflow.Context, input *devicefarm.ListTestGridSessionArtifactsInput) *DevicefarmListTestGridSessionArtifactsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTestGridSessionArtifacts, input)
    return &DevicefarmListTestGridSessionArtifactsResult{Result: future}
}

func (a *DeviceFarmStub) ListTestGridSessions(ctx workflow.Context, input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error) {
    var output devicefarm.ListTestGridSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTestGridSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListTestGridSessionsAsync(ctx workflow.Context, input *devicefarm.ListTestGridSessionsInput) *DevicefarmListTestGridSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTestGridSessions, input)
    return &DevicefarmListTestGridSessionsResult{Result: future}
}

func (a *DeviceFarmStub) ListTests(ctx workflow.Context, input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
    var output devicefarm.ListTestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTests, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListTestsAsync(ctx workflow.Context, input *devicefarm.ListTestsInput) *DevicefarmListTestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTests, input)
    return &DevicefarmListTestsResult{Result: future}
}

func (a *DeviceFarmStub) ListUniqueProblems(ctx workflow.Context, input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
    var output devicefarm.ListUniqueProblemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUniqueProblems, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListUniqueProblemsAsync(ctx workflow.Context, input *devicefarm.ListUniqueProblemsInput) *DevicefarmListUniqueProblemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUniqueProblems, input)
    return &DevicefarmListUniqueProblemsResult{Result: future}
}

func (a *DeviceFarmStub) ListUploads(ctx workflow.Context, input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
    var output devicefarm.ListUploadsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUploads, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListUploadsAsync(ctx workflow.Context, input *devicefarm.ListUploadsInput) *DevicefarmListUploadsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUploads, input)
    return &DevicefarmListUploadsResult{Result: future}
}

func (a *DeviceFarmStub) ListVPCEConfigurations(ctx workflow.Context, input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error) {
    var output devicefarm.ListVPCEConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVPCEConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ListVPCEConfigurationsAsync(ctx workflow.Context, input *devicefarm.ListVPCEConfigurationsInput) *DevicefarmListVPCEConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVPCEConfigurations, input)
    return &DevicefarmListVPCEConfigurationsResult{Result: future}
}

func (a *DeviceFarmStub) PurchaseOffering(ctx workflow.Context, input *devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
    var output devicefarm.PurchaseOfferingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseOffering, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) PurchaseOfferingAsync(ctx workflow.Context, input *devicefarm.PurchaseOfferingInput) *DevicefarmPurchaseOfferingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurchaseOffering, input)
    return &DevicefarmPurchaseOfferingResult{Result: future}
}

func (a *DeviceFarmStub) RenewOffering(ctx workflow.Context, input *devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
    var output devicefarm.RenewOfferingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RenewOffering, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) RenewOfferingAsync(ctx workflow.Context, input *devicefarm.RenewOfferingInput) *DevicefarmRenewOfferingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RenewOffering, input)
    return &DevicefarmRenewOfferingResult{Result: future}
}

func (a *DeviceFarmStub) ScheduleRun(ctx workflow.Context, input *devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
    var output devicefarm.ScheduleRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ScheduleRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) ScheduleRunAsync(ctx workflow.Context, input *devicefarm.ScheduleRunInput) *DevicefarmScheduleRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ScheduleRun, input)
    return &DevicefarmScheduleRunResult{Result: future}
}

func (a *DeviceFarmStub) StopJob(ctx workflow.Context, input *devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error) {
    var output devicefarm.StopJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopJob, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) StopJobAsync(ctx workflow.Context, input *devicefarm.StopJobInput) *DevicefarmStopJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopJob, input)
    return &DevicefarmStopJobResult{Result: future}
}

func (a *DeviceFarmStub) StopRemoteAccessSession(ctx workflow.Context, input *devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
    var output devicefarm.StopRemoteAccessSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopRemoteAccessSession, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) StopRemoteAccessSessionAsync(ctx workflow.Context, input *devicefarm.StopRemoteAccessSessionInput) *DevicefarmStopRemoteAccessSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopRemoteAccessSession, input)
    return &DevicefarmStopRemoteAccessSessionResult{Result: future}
}

func (a *DeviceFarmStub) StopRun(ctx workflow.Context, input *devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
    var output devicefarm.StopRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) StopRunAsync(ctx workflow.Context, input *devicefarm.StopRunInput) *DevicefarmStopRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopRun, input)
    return &DevicefarmStopRunResult{Result: future}
}

func (a *DeviceFarmStub) TagResource(ctx workflow.Context, input *devicefarm.TagResourceInput) (*devicefarm.TagResourceOutput, error) {
    var output devicefarm.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) TagResourceAsync(ctx workflow.Context, input *devicefarm.TagResourceInput) *DevicefarmTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &DevicefarmTagResourceResult{Result: future}
}

func (a *DeviceFarmStub) UntagResource(ctx workflow.Context, input *devicefarm.UntagResourceInput) (*devicefarm.UntagResourceOutput, error) {
    var output devicefarm.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UntagResourceAsync(ctx workflow.Context, input *devicefarm.UntagResourceInput) *DevicefarmUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &DevicefarmUntagResourceResult{Result: future}
}

func (a *DeviceFarmStub) UpdateDeviceInstance(ctx workflow.Context, input *devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error) {
    var output devicefarm.UpdateDeviceInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeviceInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateDeviceInstanceAsync(ctx workflow.Context, input *devicefarm.UpdateDeviceInstanceInput) *DevicefarmUpdateDeviceInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDeviceInstance, input)
    return &DevicefarmUpdateDeviceInstanceResult{Result: future}
}

func (a *DeviceFarmStub) UpdateDevicePool(ctx workflow.Context, input *devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
    var output devicefarm.UpdateDevicePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDevicePool, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateDevicePoolAsync(ctx workflow.Context, input *devicefarm.UpdateDevicePoolInput) *DevicefarmUpdateDevicePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDevicePool, input)
    return &DevicefarmUpdateDevicePoolResult{Result: future}
}

func (a *DeviceFarmStub) UpdateInstanceProfile(ctx workflow.Context, input *devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error) {
    var output devicefarm.UpdateInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateInstanceProfileAsync(ctx workflow.Context, input *devicefarm.UpdateInstanceProfileInput) *DevicefarmUpdateInstanceProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateInstanceProfile, input)
    return &DevicefarmUpdateInstanceProfileResult{Result: future}
}

func (a *DeviceFarmStub) UpdateNetworkProfile(ctx workflow.Context, input *devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
    var output devicefarm.UpdateNetworkProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNetworkProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateNetworkProfileAsync(ctx workflow.Context, input *devicefarm.UpdateNetworkProfileInput) *DevicefarmUpdateNetworkProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateNetworkProfile, input)
    return &DevicefarmUpdateNetworkProfileResult{Result: future}
}

func (a *DeviceFarmStub) UpdateProject(ctx workflow.Context, input *devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
    var output devicefarm.UpdateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateProjectAsync(ctx workflow.Context, input *devicefarm.UpdateProjectInput) *DevicefarmUpdateProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input)
    return &DevicefarmUpdateProjectResult{Result: future}
}

func (a *DeviceFarmStub) UpdateTestGridProject(ctx workflow.Context, input *devicefarm.UpdateTestGridProjectInput) (*devicefarm.UpdateTestGridProjectOutput, error) {
    var output devicefarm.UpdateTestGridProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTestGridProject, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateTestGridProjectAsync(ctx workflow.Context, input *devicefarm.UpdateTestGridProjectInput) *DevicefarmUpdateTestGridProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTestGridProject, input)
    return &DevicefarmUpdateTestGridProjectResult{Result: future}
}

func (a *DeviceFarmStub) UpdateUpload(ctx workflow.Context, input *devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error) {
    var output devicefarm.UpdateUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateUploadAsync(ctx workflow.Context, input *devicefarm.UpdateUploadInput) *DevicefarmUpdateUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUpload, input)
    return &DevicefarmUpdateUploadResult{Result: future}
}

func (a *DeviceFarmStub) UpdateVPCEConfiguration(ctx workflow.Context, input *devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error) {
    var output devicefarm.UpdateVPCEConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVPCEConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *DeviceFarmStub) UpdateVPCEConfigurationAsync(ctx workflow.Context, input *devicefarm.UpdateVPCEConfigurationInput) *DevicefarmUpdateVPCEConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVPCEConfiguration, input)
    return &DevicefarmUpdateVPCEConfigurationResult{Result: future}
}