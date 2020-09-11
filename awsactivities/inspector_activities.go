
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/inspector/inspectoriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type InspectorActivities struct {
    client inspectoriface.InspectorAPI
}

func NewInspectorActivities(session *session.Session, config ...*aws.Config) *InspectorActivities {
    client := inspector.New(session, config...)
    return &InspectorActivities{client: client}
}

func (a *InspectorActivities) AddAttributesToFindings(ctx context.Context, input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
    return a.client.AddAttributesToFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) CreateAssessmentTarget(ctx context.Context, input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error) {
    return a.client.CreateAssessmentTargetWithContext(ctx, input)
}

func (a *InspectorActivities) CreateAssessmentTemplate(ctx context.Context, input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error) {
    return a.client.CreateAssessmentTemplateWithContext(ctx, input)
}

func (a *InspectorActivities) CreateExclusionsPreview(ctx context.Context, input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error) {
    return a.client.CreateExclusionsPreviewWithContext(ctx, input)
}

func (a *InspectorActivities) CreateResourceGroup(ctx context.Context, input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error) {
    return a.client.CreateResourceGroupWithContext(ctx, input)
}

func (a *InspectorActivities) DeleteAssessmentRun(ctx context.Context, input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error) {
    return a.client.DeleteAssessmentRunWithContext(ctx, input)
}

func (a *InspectorActivities) DeleteAssessmentTarget(ctx context.Context, input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error) {
    return a.client.DeleteAssessmentTargetWithContext(ctx, input)
}

func (a *InspectorActivities) DeleteAssessmentTemplate(ctx context.Context, input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error) {
    return a.client.DeleteAssessmentTemplateWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeAssessmentRuns(ctx context.Context, input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error) {
    return a.client.DescribeAssessmentRunsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeAssessmentTargets(ctx context.Context, input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error) {
    return a.client.DescribeAssessmentTargetsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeAssessmentTemplates(ctx context.Context, input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error) {
    return a.client.DescribeAssessmentTemplatesWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeCrossAccountAccessRole(ctx context.Context, input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
    return a.client.DescribeCrossAccountAccessRoleWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeExclusions(ctx context.Context, input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error) {
    return a.client.DescribeExclusionsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeFindings(ctx context.Context, input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error) {
    return a.client.DescribeFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeResourceGroups(ctx context.Context, input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error) {
    return a.client.DescribeResourceGroupsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeRulesPackages(ctx context.Context, input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error) {
    return a.client.DescribeRulesPackagesWithContext(ctx, input)
}

func (a *InspectorActivities) GetAssessmentReport(ctx context.Context, input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error) {
    return a.client.GetAssessmentReportWithContext(ctx, input)
}

func (a *InspectorActivities) GetExclusionsPreview(ctx context.Context, input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error) {
    return a.client.GetExclusionsPreviewWithContext(ctx, input)
}

func (a *InspectorActivities) GetTelemetryMetadata(ctx context.Context, input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error) {
    return a.client.GetTelemetryMetadataWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentRunAgents(ctx context.Context, input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error) {
    return a.client.ListAssessmentRunAgentsWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentRuns(ctx context.Context, input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error) {
    return a.client.ListAssessmentRunsWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentTargets(ctx context.Context, input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error) {
    return a.client.ListAssessmentTargetsWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentTemplates(ctx context.Context, input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error) {
    return a.client.ListAssessmentTemplatesWithContext(ctx, input)
}

func (a *InspectorActivities) ListEventSubscriptions(ctx context.Context, input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error) {
    return a.client.ListEventSubscriptionsWithContext(ctx, input)
}

func (a *InspectorActivities) ListExclusions(ctx context.Context, input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error) {
    return a.client.ListExclusionsWithContext(ctx, input)
}

func (a *InspectorActivities) ListFindings(ctx context.Context, input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error) {
    return a.client.ListFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) ListRulesPackages(ctx context.Context, input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error) {
    return a.client.ListRulesPackagesWithContext(ctx, input)
}

func (a *InspectorActivities) ListTagsForResource(ctx context.Context, input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *InspectorActivities) PreviewAgents(ctx context.Context, input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error) {
    return a.client.PreviewAgentsWithContext(ctx, input)
}

func (a *InspectorActivities) RegisterCrossAccountAccessRole(ctx context.Context, input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
    return a.client.RegisterCrossAccountAccessRoleWithContext(ctx, input)
}

func (a *InspectorActivities) RemoveAttributesFromFindings(ctx context.Context, input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error) {
    return a.client.RemoveAttributesFromFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) SetTagsForResource(ctx context.Context, input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error) {
    return a.client.SetTagsForResourceWithContext(ctx, input)
}

func (a *InspectorActivities) StartAssessmentRun(ctx context.Context, input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error) {
    return a.client.StartAssessmentRunWithContext(ctx, input)
}

func (a *InspectorActivities) StopAssessmentRun(ctx context.Context, input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error) {
    return a.client.StopAssessmentRunWithContext(ctx, input)
}

func (a *InspectorActivities) SubscribeToEvent(ctx context.Context, input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error) {
    return a.client.SubscribeToEventWithContext(ctx, input)
}

func (a *InspectorActivities) UnsubscribeFromEvent(ctx context.Context, input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error) {
    return a.client.UnsubscribeFromEventWithContext(ctx, input)
}

func (a *InspectorActivities) UpdateAssessmentTarget(ctx context.Context, input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error) {
    return a.client.UpdateAssessmentTargetWithContext(ctx, input)
}
