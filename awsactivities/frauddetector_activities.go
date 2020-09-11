
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/frauddetector/frauddetectoriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type FraudDetectorActivities struct {
    client frauddetectoriface.FraudDetectorAPI
}

func NewFraudDetectorActivities(session *session.Session, config ...*aws.Config) *FraudDetectorActivities {
    client := frauddetector.New(session, config...)
    return &FraudDetectorActivities{client: client}
}

func (a *FraudDetectorActivities) BatchCreateVariable(ctx context.Context, input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
    return a.client.BatchCreateVariableWithContext(ctx, input)
}

func (a *FraudDetectorActivities) BatchGetVariable(ctx context.Context, input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error) {
    return a.client.BatchGetVariableWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateDetectorVersion(ctx context.Context, input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error) {
    return a.client.CreateDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateModel(ctx context.Context, input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error) {
    return a.client.CreateModelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateModelVersion(ctx context.Context, input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error) {
    return a.client.CreateModelVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateRule(ctx context.Context, input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error) {
    return a.client.CreateRuleWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateVariable(ctx context.Context, input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error) {
    return a.client.CreateVariableWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteDetector(ctx context.Context, input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error) {
    return a.client.DeleteDetectorWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteDetectorVersion(ctx context.Context, input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error) {
    return a.client.DeleteDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteEvent(ctx context.Context, input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error) {
    return a.client.DeleteEventWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteRule(ctx context.Context, input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error) {
    return a.client.DeleteRuleWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DescribeDetector(ctx context.Context, input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error) {
    return a.client.DescribeDetectorWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DescribeModelVersions(ctx context.Context, input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error) {
    return a.client.DescribeModelVersionsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetDetectorVersion(ctx context.Context, input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error) {
    return a.client.GetDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetDetectors(ctx context.Context, input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error) {
    return a.client.GetDetectorsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetEntityTypes(ctx context.Context, input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error) {
    return a.client.GetEntityTypesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetEventPrediction(ctx context.Context, input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error) {
    return a.client.GetEventPredictionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetEventTypes(ctx context.Context, input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error) {
    return a.client.GetEventTypesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetExternalModels(ctx context.Context, input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error) {
    return a.client.GetExternalModelsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetKMSEncryptionKey(ctx context.Context, input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
    return a.client.GetKMSEncryptionKeyWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetLabels(ctx context.Context, input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error) {
    return a.client.GetLabelsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetModelVersion(ctx context.Context, input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error) {
    return a.client.GetModelVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetModels(ctx context.Context, input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error) {
    return a.client.GetModelsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetOutcomes(ctx context.Context, input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error) {
    return a.client.GetOutcomesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetRules(ctx context.Context, input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error) {
    return a.client.GetRulesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetVariables(ctx context.Context, input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error) {
    return a.client.GetVariablesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) ListTagsForResource(ctx context.Context, input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutDetector(ctx context.Context, input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error) {
    return a.client.PutDetectorWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutEntityType(ctx context.Context, input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error) {
    return a.client.PutEntityTypeWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutEventType(ctx context.Context, input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error) {
    return a.client.PutEventTypeWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutExternalModel(ctx context.Context, input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error) {
    return a.client.PutExternalModelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutKMSEncryptionKey(ctx context.Context, input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error) {
    return a.client.PutKMSEncryptionKeyWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutLabel(ctx context.Context, input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error) {
    return a.client.PutLabelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutOutcome(ctx context.Context, input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error) {
    return a.client.PutOutcomeWithContext(ctx, input)
}

func (a *FraudDetectorActivities) TagResource(ctx context.Context, input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UntagResource(ctx context.Context, input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersion(ctx context.Context, input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error) {
    return a.client.UpdateDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersionMetadata(ctx context.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error) {
    return a.client.UpdateDetectorVersionMetadataWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersionStatus(ctx context.Context, input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error) {
    return a.client.UpdateDetectorVersionStatusWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateModel(ctx context.Context, input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error) {
    return a.client.UpdateModelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateModelVersion(ctx context.Context, input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error) {
    return a.client.UpdateModelVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateModelVersionStatus(ctx context.Context, input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error) {
    return a.client.UpdateModelVersionStatusWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateRuleMetadata(ctx context.Context, input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error) {
    return a.client.UpdateRuleMetadataWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateRuleVersion(ctx context.Context, input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error) {
    return a.client.UpdateRuleVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateVariable(ctx context.Context, input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error) {
    return a.client.UpdateVariableWithContext(ctx, input)
}
