package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ResourceGroupsTaggingAPIActivities struct {
	client resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
}

func NewResourceGroupsTaggingAPIActivities(session *session.Session, config ...*aws.Config) *ResourceGroupsTaggingAPIActivities {
	client := resourcegroupstaggingapi.New(session, config...)
	return &ResourceGroupsTaggingAPIActivities{client: client}
}

func (a *ResourceGroupsTaggingAPIActivities) DescribeReportCreation(ctx context.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	return a.client.DescribeReportCreationWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetComplianceSummary(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	return a.client.GetComplianceSummaryWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetResources(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	return a.client.GetResourcesWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetTagKeys(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	return a.client.GetTagKeysWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetTagValues(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	return a.client.GetTagValuesWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) StartReportCreation(ctx context.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	return a.client.StartReportCreationWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) TagResources(ctx context.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	return a.client.TagResourcesWithContext(ctx, input)
}

func (a *ResourceGroupsTaggingAPIActivities) UntagResources(ctx context.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	return a.client.UntagResourcesWithContext(ctx, input)
}
