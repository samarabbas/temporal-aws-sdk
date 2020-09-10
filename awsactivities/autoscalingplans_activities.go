package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type AutoScalingPlansActivities struct {
	client autoscalingplansiface.AutoScalingPlansAPI
}

func NewAutoScalingPlansActivities(session *session.Session, config ...*aws.Config) *AutoScalingPlansActivities {
	client := autoscalingplans.New(session, config...)
	return &AutoScalingPlansActivities{client: client}
}

func (a *AutoScalingPlansActivities) CreateScalingPlan(ctx context.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
	return a.client.CreateScalingPlanWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) DeleteScalingPlan(ctx context.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	return a.client.DeleteScalingPlanWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) DescribeScalingPlanResources(ctx context.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	return a.client.DescribeScalingPlanResourcesWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) DescribeScalingPlans(ctx context.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	return a.client.DescribeScalingPlansWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) GetScalingPlanResourceForecastData(ctx context.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	return a.client.GetScalingPlanResourceForecastDataWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) UpdateScalingPlan(ctx context.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	return a.client.UpdateScalingPlanWithContext(ctx, input)
}
