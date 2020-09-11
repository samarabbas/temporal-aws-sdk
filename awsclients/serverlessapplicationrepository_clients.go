package awsclients

import (
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"go.temporal.io/sdk/workflow"
)

type ServerlessApplicationRepositoryClient interface {
       CreateApplication(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error)
       CreateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) *ServerlessapplicationrepositoryCreateApplicationResult

       CreateApplicationVersion(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
       CreateApplicationVersionAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) *ServerlessapplicationrepositoryCreateApplicationVersionResult

       CreateCloudFormationChangeSet(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
       CreateCloudFormationChangeSetAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) *ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult

       CreateCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
       CreateCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) *ServerlessapplicationrepositoryCreateCloudFormationTemplateResult

       DeleteApplication(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
       DeleteApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) *ServerlessapplicationrepositoryDeleteApplicationResult

       GetApplication(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error)
       GetApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) *ServerlessapplicationrepositoryGetApplicationResult

       GetApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
       GetApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) *ServerlessapplicationrepositoryGetApplicationPolicyResult

       GetCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
       GetCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) *ServerlessapplicationrepositoryGetCloudFormationTemplateResult

       ListApplicationDependencies(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
       ListApplicationDependenciesAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) *ServerlessapplicationrepositoryListApplicationDependenciesResult

       ListApplicationVersions(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
       ListApplicationVersionsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) *ServerlessapplicationrepositoryListApplicationVersionsResult

       ListApplications(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error)
       ListApplicationsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) *ServerlessapplicationrepositoryListApplicationsResult

       PutApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
       PutApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) *ServerlessapplicationrepositoryPutApplicationPolicyResult

       UnshareApplication(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error)
       UnshareApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) *ServerlessapplicationrepositoryUnshareApplicationResult

       UpdateApplication(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
       UpdateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) *ServerlessapplicationrepositoryUpdateApplicationResult
}

type ServerlessApplicationRepositoryStub struct {
}

func NewServerlessApplicationRepositoryStub() ServerlessApplicationRepositoryClient {
    return &ServerlessApplicationRepositoryStub{}
}

type ServerlessapplicationrepositoryCreateApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryCreateApplicationVersionResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateApplicationVersionResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryCreateCloudFormationTemplateResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateCloudFormationTemplateResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryDeleteApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
    var output serverlessapplicationrepository.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryGetApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryGetApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetApplicationOutput, error) {
    var output serverlessapplicationrepository.GetApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryGetApplicationPolicyResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryGetApplicationPolicyResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.GetApplicationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryGetCloudFormationTemplateResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryGetCloudFormationTemplateResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.GetCloudFormationTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryListApplicationDependenciesResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryListApplicationDependenciesResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
    var output serverlessapplicationrepository.ListApplicationDependenciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryListApplicationVersionsResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryListApplicationVersionsResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryListApplicationsResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryListApplicationsResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryPutApplicationPolicyResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryPutApplicationPolicyResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.PutApplicationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryUnshareApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryUnshareApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
    var output serverlessapplicationrepository.UnshareApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ServerlessapplicationrepositoryUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryUpdateApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
    var output serverlessapplicationrepository.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateApplication(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) *ServerlessapplicationrepositoryCreateApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateApplication", input)
    return &ServerlessapplicationrepositoryCreateApplicationResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) CreateApplicationVersion(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationVersionOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateApplicationVersion", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateApplicationVersionAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) *ServerlessapplicationrepositoryCreateApplicationVersionResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateApplicationVersion", input)
    return &ServerlessapplicationrepositoryCreateApplicationVersionResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) CreateCloudFormationChangeSet(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationChangeSetOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateCloudFormationChangeSet", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateCloudFormationChangeSetAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) *ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateCloudFormationChangeSet", input)
    return &ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) CreateCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationTemplateOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateCloudFormationTemplate", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) *ServerlessapplicationrepositoryCreateCloudFormationTemplateResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.CreateCloudFormationTemplate", input)
    return &ServerlessapplicationrepositoryCreateCloudFormationTemplateResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) DeleteApplication(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
    var output serverlessapplicationrepository.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.DeleteApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) DeleteApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) *ServerlessapplicationrepositoryDeleteApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.DeleteApplication", input)
    return &ServerlessapplicationrepositoryDeleteApplicationResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) GetApplication(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
    var output serverlessapplicationrepository.GetApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.GetApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) GetApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) *ServerlessapplicationrepositoryGetApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.GetApplication", input)
    return &ServerlessapplicationrepositoryGetApplicationResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) GetApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.GetApplicationPolicyOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.GetApplicationPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) GetApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) *ServerlessapplicationrepositoryGetApplicationPolicyResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.GetApplicationPolicy", input)
    return &ServerlessapplicationrepositoryGetApplicationPolicyResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) GetCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.GetCloudFormationTemplateOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.GetCloudFormationTemplate", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) GetCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) *ServerlessapplicationrepositoryGetCloudFormationTemplateResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.GetCloudFormationTemplate", input)
    return &ServerlessapplicationrepositoryGetCloudFormationTemplateResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationDependencies(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
    var output serverlessapplicationrepository.ListApplicationDependenciesOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.ListApplicationDependencies", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationDependenciesAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) *ServerlessapplicationrepositoryListApplicationDependenciesResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.ListApplicationDependencies", input)
    return &ServerlessapplicationrepositoryListApplicationDependenciesResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationVersions(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationVersionsOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.ListApplicationVersions", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationVersionsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) *ServerlessapplicationrepositoryListApplicationVersionsResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.ListApplicationVersions", input)
    return &ServerlessapplicationrepositoryListApplicationVersionsResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) ListApplications(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationsOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.ListApplications", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) *ServerlessapplicationrepositoryListApplicationsResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.ListApplications", input)
    return &ServerlessapplicationrepositoryListApplicationsResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) PutApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.PutApplicationPolicyOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.PutApplicationPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) PutApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) *ServerlessapplicationrepositoryPutApplicationPolicyResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.PutApplicationPolicy", input)
    return &ServerlessapplicationrepositoryPutApplicationPolicyResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) UnshareApplication(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
    var output serverlessapplicationrepository.UnshareApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.UnshareApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) UnshareApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) *ServerlessapplicationrepositoryUnshareApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.UnshareApplication", input)
    return &ServerlessapplicationrepositoryUnshareApplicationResult{Result: future}
}

func (a *ServerlessApplicationRepositoryStub) UpdateApplication(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
    var output serverlessapplicationrepository.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.UpdateApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) UpdateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) *ServerlessapplicationrepositoryUpdateApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ServerlessApplicationRepository.UpdateApplication", input)
    return &ServerlessapplicationrepositoryUpdateApplicationResult{Result: future}
}
