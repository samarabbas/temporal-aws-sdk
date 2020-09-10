package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler/codeguruprofileriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodeGuruProfilerActivities struct {
	client codeguruprofileriface.CodeGuruProfilerAPI
}

func NewCodeGuruProfilerActivities(session *session.Session, config ...*aws.Config) *CodeGuruProfilerActivities {
	client := codeguruprofiler.New(session, config...)
	return &CodeGuruProfilerActivities{client: client}
}

func (a *CodeGuruProfilerActivities) AddNotificationChannels(ctx context.Context, input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
	return a.client.AddNotificationChannelsWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) BatchGetFrameMetricData(ctx context.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
	return a.client.BatchGetFrameMetricDataWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) ConfigureAgent(ctx context.Context, input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error) {
	return a.client.ConfigureAgentWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) CreateProfilingGroup(ctx context.Context, input *codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateProfilingGroupWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) DeleteProfilingGroup(ctx context.Context, input *codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error) {
	return a.client.DeleteProfilingGroupWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) DescribeProfilingGroup(ctx context.Context, input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
	return a.client.DescribeProfilingGroupWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) GetFindingsReportAccountSummary(ctx context.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
	return a.client.GetFindingsReportAccountSummaryWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) GetNotificationConfiguration(ctx context.Context, input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
	return a.client.GetNotificationConfigurationWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) GetPolicy(ctx context.Context, input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error) {
	return a.client.GetPolicyWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) GetProfile(ctx context.Context, input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error) {
	return a.client.GetProfileWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) GetRecommendations(ctx context.Context, input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error) {
	return a.client.GetRecommendationsWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) ListFindingsReports(ctx context.Context, input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error) {
	return a.client.ListFindingsReportsWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) ListProfileTimes(ctx context.Context, input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error) {
	return a.client.ListProfileTimesWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) ListProfilingGroups(ctx context.Context, input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
	return a.client.ListProfilingGroupsWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) ListTagsForResource(ctx context.Context, input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) PostAgentProfile(ctx context.Context, input *codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error) {
	return a.client.PostAgentProfileWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) PutPermission(ctx context.Context, input *codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error) {
	return a.client.PutPermissionWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) RemoveNotificationChannel(ctx context.Context, input *codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error) {
	return a.client.RemoveNotificationChannelWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) RemovePermission(ctx context.Context, input *codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error) {
	return a.client.RemovePermissionWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) SubmitFeedback(ctx context.Context, input *codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error) {
	return a.client.SubmitFeedbackWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) TagResource(ctx context.Context, input *codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) UntagResource(ctx context.Context, input *codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CodeGuruProfilerActivities) UpdateProfilingGroup(ctx context.Context, input *codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error) {
	return a.client.UpdateProfilingGroupWithContext(ctx, input)
}
