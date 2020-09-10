package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodeBuildActivities struct {
	client codebuildiface.CodeBuildAPI
}

func NewCodeBuildActivities(session *session.Session, config ...*aws.Config) *CodeBuildActivities {
	client := codebuild.New(session, config...)
	return &CodeBuildActivities{client: client}
}

func (a *CodeBuildActivities) BatchDeleteBuilds(ctx context.Context, input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
	return a.client.BatchDeleteBuildsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetBuildBatches(ctx context.Context, input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error) {
	return a.client.BatchGetBuildBatchesWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetBuilds(ctx context.Context, input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	return a.client.BatchGetBuildsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetProjects(ctx context.Context, input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
	return a.client.BatchGetProjectsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetReportGroups(ctx context.Context, input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error) {
	return a.client.BatchGetReportGroupsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetReports(ctx context.Context, input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error) {
	return a.client.BatchGetReportsWithContext(ctx, input)
}

func (a *CodeBuildActivities) CreateProject(ctx context.Context, input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error) {
	return a.client.CreateProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) CreateReportGroup(ctx context.Context, input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error) {
	return a.client.CreateReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) CreateWebhook(ctx context.Context, input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error) {
	return a.client.CreateWebhookWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteBuildBatch(ctx context.Context, input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error) {
	return a.client.DeleteBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteProject(ctx context.Context, input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error) {
	return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteReport(ctx context.Context, input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error) {
	return a.client.DeleteReportWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteReportGroup(ctx context.Context, input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error) {
	return a.client.DeleteReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteResourcePolicy(ctx context.Context, input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error) {
	return a.client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteSourceCredentials(ctx context.Context, input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error) {
	return a.client.DeleteSourceCredentialsWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteWebhook(ctx context.Context, input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error) {
	return a.client.DeleteWebhookWithContext(ctx, input)
}

func (a *CodeBuildActivities) DescribeCodeCoverages(ctx context.Context, input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error) {
	return a.client.DescribeCodeCoveragesWithContext(ctx, input)
}

func (a *CodeBuildActivities) DescribeTestCases(ctx context.Context, input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error) {
	return a.client.DescribeTestCasesWithContext(ctx, input)
}

func (a *CodeBuildActivities) GetResourcePolicy(ctx context.Context, input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error) {
	return a.client.GetResourcePolicyWithContext(ctx, input)
}

func (a *CodeBuildActivities) ImportSourceCredentials(ctx context.Context, input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error) {
	return a.client.ImportSourceCredentialsWithContext(ctx, input)
}

func (a *CodeBuildActivities) InvalidateProjectCache(ctx context.Context, input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error) {
	return a.client.InvalidateProjectCacheWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuildBatches(ctx context.Context, input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error) {
	return a.client.ListBuildBatchesWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuildBatchesForProject(ctx context.Context, input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	return a.client.ListBuildBatchesForProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuilds(ctx context.Context, input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	return a.client.ListBuildsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuildsForProject(ctx context.Context, input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
	return a.client.ListBuildsForProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListCuratedEnvironmentImages(ctx context.Context, input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	return a.client.ListCuratedEnvironmentImagesWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListProjects(ctx context.Context, input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
	return a.client.ListProjectsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListReportGroups(ctx context.Context, input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error) {
	return a.client.ListReportGroupsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListReports(ctx context.Context, input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error) {
	return a.client.ListReportsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListReportsForReportGroup(ctx context.Context, input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error) {
	return a.client.ListReportsForReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListSharedProjects(ctx context.Context, input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error) {
	return a.client.ListSharedProjectsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListSharedReportGroups(ctx context.Context, input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error) {
	return a.client.ListSharedReportGroupsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListSourceCredentials(ctx context.Context, input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error) {
	return a.client.ListSourceCredentialsWithContext(ctx, input)
}

func (a *CodeBuildActivities) PutResourcePolicy(ctx context.Context, input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error) {
	return a.client.PutResourcePolicyWithContext(ctx, input)
}

func (a *CodeBuildActivities) RetryBuild(ctx context.Context, input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error) {
	return a.client.RetryBuildWithContext(ctx, input)
}

func (a *CodeBuildActivities) RetryBuildBatch(ctx context.Context, input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error) {
	return a.client.RetryBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) StartBuild(ctx context.Context, input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error) {
	return a.client.StartBuildWithContext(ctx, input)
}

func (a *CodeBuildActivities) StartBuildBatch(ctx context.Context, input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error) {
	return a.client.StartBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) StopBuild(ctx context.Context, input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error) {
	return a.client.StopBuildWithContext(ctx, input)
}

func (a *CodeBuildActivities) StopBuildBatch(ctx context.Context, input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error) {
	return a.client.StopBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) UpdateProject(ctx context.Context, input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error) {
	return a.client.UpdateProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) UpdateReportGroup(ctx context.Context, input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error) {
	return a.client.UpdateReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) UpdateWebhook(ctx context.Context, input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error) {
	return a.client.UpdateWebhookWithContext(ctx, input)
}
