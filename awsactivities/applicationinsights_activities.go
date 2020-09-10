package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"github.com/aws/aws-sdk-go/service/applicationinsights/applicationinsightsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ApplicationInsightsActivities struct {
	client applicationinsightsiface.ApplicationInsightsAPI
}

func NewApplicationInsightsActivities(session *session.Session, config ...*aws.Config) *ApplicationInsightsActivities {
	client := applicationinsights.New(session, config...)
	return &ApplicationInsightsActivities{client: client}
}

func (a *ApplicationInsightsActivities) CreateApplication(ctx context.Context, input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error) {
	return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) CreateComponent(ctx context.Context, input *applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error) {
	return a.client.CreateComponentWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) CreateLogPattern(ctx context.Context, input *applicationinsights.CreateLogPatternInput) (*applicationinsights.CreateLogPatternOutput, error) {
	return a.client.CreateLogPatternWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DeleteApplication(ctx context.Context, input *applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error) {
	return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DeleteComponent(ctx context.Context, input *applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error) {
	return a.client.DeleteComponentWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DeleteLogPattern(ctx context.Context, input *applicationinsights.DeleteLogPatternInput) (*applicationinsights.DeleteLogPatternOutput, error) {
	return a.client.DeleteLogPatternWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeApplication(ctx context.Context, input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error) {
	return a.client.DescribeApplicationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeComponent(ctx context.Context, input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error) {
	return a.client.DescribeComponentWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeComponentConfiguration(ctx context.Context, input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	return a.client.DescribeComponentConfigurationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeComponentConfigurationRecommendation(ctx context.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	return a.client.DescribeComponentConfigurationRecommendationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeLogPattern(ctx context.Context, input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error) {
	return a.client.DescribeLogPatternWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeObservation(ctx context.Context, input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error) {
	return a.client.DescribeObservationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeProblem(ctx context.Context, input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error) {
	return a.client.DescribeProblemWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) DescribeProblemObservations(ctx context.Context, input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	return a.client.DescribeProblemObservationsWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListApplications(ctx context.Context, input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error) {
	return a.client.ListApplicationsWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListComponents(ctx context.Context, input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error) {
	return a.client.ListComponentsWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListConfigurationHistory(ctx context.Context, input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	return a.client.ListConfigurationHistoryWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListLogPatternSets(ctx context.Context, input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error) {
	return a.client.ListLogPatternSetsWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListLogPatterns(ctx context.Context, input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error) {
	return a.client.ListLogPatternsWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListProblems(ctx context.Context, input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error) {
	return a.client.ListProblemsWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) ListTagsForResource(ctx context.Context, input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) TagResource(ctx context.Context, input *applicationinsights.TagResourceInput) (*applicationinsights.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) UntagResource(ctx context.Context, input *applicationinsights.UntagResourceInput) (*applicationinsights.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) UpdateApplication(ctx context.Context, input *applicationinsights.UpdateApplicationInput) (*applicationinsights.UpdateApplicationOutput, error) {
	return a.client.UpdateApplicationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) UpdateComponent(ctx context.Context, input *applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error) {
	return a.client.UpdateComponentWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) UpdateComponentConfiguration(ctx context.Context, input *applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
	return a.client.UpdateComponentConfigurationWithContext(ctx, input)
}

func (a *ApplicationInsightsActivities) UpdateLogPattern(ctx context.Context, input *applicationinsights.UpdateLogPatternInput) (*applicationinsights.UpdateLogPatternOutput, error) {
	return a.client.UpdateLogPatternWithContext(ctx, input)
}
