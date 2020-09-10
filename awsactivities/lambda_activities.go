package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type LambdaActivities struct {
	client lambdaiface.LambdaAPI
}

func NewLambdaActivities(session *session.Session, config ...*aws.Config) *LambdaActivities {
	client := lambda.New(session, config...)
	return &LambdaActivities{client: client}
}

func (a *LambdaActivities) AddLayerVersionPermission(ctx context.Context, input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error) {
	return a.client.AddLayerVersionPermissionWithContext(ctx, input)
}

func (a *LambdaActivities) AddPermission(ctx context.Context, input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
	return a.client.AddPermissionWithContext(ctx, input)
}

func (a *LambdaActivities) CreateAlias(ctx context.Context, input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error) {
	return a.client.CreateAliasWithContext(ctx, input)
}

func (a *LambdaActivities) CreateEventSourceMapping(ctx context.Context, input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	return a.client.CreateEventSourceMappingWithContext(ctx, input)
}

func (a *LambdaActivities) CreateFunction(ctx context.Context, input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error) {
	return a.client.CreateFunctionWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteAlias(ctx context.Context, input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error) {
	return a.client.DeleteAliasWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteEventSourceMapping(ctx context.Context, input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	return a.client.DeleteEventSourceMappingWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteFunction(ctx context.Context, input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error) {
	return a.client.DeleteFunctionWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteFunctionConcurrency(ctx context.Context, input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error) {
	return a.client.DeleteFunctionConcurrencyWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteFunctionEventInvokeConfig(ctx context.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error) {
	return a.client.DeleteFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteLayerVersion(ctx context.Context, input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error) {
	return a.client.DeleteLayerVersionWithContext(ctx, input)
}

func (a *LambdaActivities) DeleteProvisionedConcurrencyConfig(ctx context.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error) {
	return a.client.DeleteProvisionedConcurrencyConfigWithContext(ctx, input)
}

func (a *LambdaActivities) GetAccountSettings(ctx context.Context, input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error) {
	return a.client.GetAccountSettingsWithContext(ctx, input)
}

func (a *LambdaActivities) GetAlias(ctx context.Context, input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error) {
	return a.client.GetAliasWithContext(ctx, input)
}

func (a *LambdaActivities) GetEventSourceMapping(ctx context.Context, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	return a.client.GetEventSourceMappingWithContext(ctx, input)
}

func (a *LambdaActivities) GetFunction(ctx context.Context, input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
	return a.client.GetFunctionWithContext(ctx, input)
}

func (a *LambdaActivities) GetFunctionConcurrency(ctx context.Context, input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error) {
	return a.client.GetFunctionConcurrencyWithContext(ctx, input)
}

func (a *LambdaActivities) GetFunctionConfiguration(ctx context.Context, input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
	return a.client.GetFunctionConfigurationWithContext(ctx, input)
}

func (a *LambdaActivities) GetFunctionEventInvokeConfig(ctx context.Context, input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
	return a.client.GetFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *LambdaActivities) GetLayerVersion(ctx context.Context, input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error) {
	return a.client.GetLayerVersionWithContext(ctx, input)
}

func (a *LambdaActivities) GetLayerVersionByArn(ctx context.Context, input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error) {
	return a.client.GetLayerVersionByArnWithContext(ctx, input)
}

func (a *LambdaActivities) GetLayerVersionPolicy(ctx context.Context, input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error) {
	return a.client.GetLayerVersionPolicyWithContext(ctx, input)
}

func (a *LambdaActivities) GetPolicy(ctx context.Context, input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error) {
	return a.client.GetPolicyWithContext(ctx, input)
}

func (a *LambdaActivities) GetProvisionedConcurrencyConfig(ctx context.Context, input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
	return a.client.GetProvisionedConcurrencyConfigWithContext(ctx, input)
}

func (a *LambdaActivities) Invoke(ctx context.Context, input *lambda.InvokeInput) (*lambda.InvokeOutput, error) {
	return a.client.InvokeWithContext(ctx, input)
}

func (a *LambdaActivities) ListAliases(ctx context.Context, input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error) {
	return a.client.ListAliasesWithContext(ctx, input)
}

func (a *LambdaActivities) ListEventSourceMappings(ctx context.Context, input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error) {
	return a.client.ListEventSourceMappingsWithContext(ctx, input)
}

func (a *LambdaActivities) ListFunctionEventInvokeConfigs(ctx context.Context, input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	return a.client.ListFunctionEventInvokeConfigsWithContext(ctx, input)
}

func (a *LambdaActivities) ListFunctions(ctx context.Context, input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
	return a.client.ListFunctionsWithContext(ctx, input)
}

func (a *LambdaActivities) ListLayerVersions(ctx context.Context, input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error) {
	return a.client.ListLayerVersionsWithContext(ctx, input)
}

func (a *LambdaActivities) ListLayers(ctx context.Context, input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error) {
	return a.client.ListLayersWithContext(ctx, input)
}

func (a *LambdaActivities) ListProvisionedConcurrencyConfigs(ctx context.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	return a.client.ListProvisionedConcurrencyConfigsWithContext(ctx, input)
}

func (a *LambdaActivities) ListTags(ctx context.Context, input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *LambdaActivities) ListVersionsByFunction(ctx context.Context, input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error) {
	return a.client.ListVersionsByFunctionWithContext(ctx, input)
}

func (a *LambdaActivities) PublishLayerVersion(ctx context.Context, input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error) {
	return a.client.PublishLayerVersionWithContext(ctx, input)
}

func (a *LambdaActivities) PublishVersion(ctx context.Context, input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error) {
	return a.client.PublishVersionWithContext(ctx, input)
}

func (a *LambdaActivities) PutFunctionConcurrency(ctx context.Context, input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error) {
	return a.client.PutFunctionConcurrencyWithContext(ctx, input)
}

func (a *LambdaActivities) PutFunctionEventInvokeConfig(ctx context.Context, input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error) {
	return a.client.PutFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *LambdaActivities) PutProvisionedConcurrencyConfig(ctx context.Context, input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error) {
	return a.client.PutProvisionedConcurrencyConfigWithContext(ctx, input)
}

func (a *LambdaActivities) RemoveLayerVersionPermission(ctx context.Context, input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error) {
	return a.client.RemoveLayerVersionPermissionWithContext(ctx, input)
}

func (a *LambdaActivities) RemovePermission(ctx context.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error) {
	return a.client.RemovePermissionWithContext(ctx, input)
}

func (a *LambdaActivities) TagResource(ctx context.Context, input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *LambdaActivities) UntagResource(ctx context.Context, input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *LambdaActivities) UpdateAlias(ctx context.Context, input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error) {
	return a.client.UpdateAliasWithContext(ctx, input)
}

func (a *LambdaActivities) UpdateEventSourceMapping(ctx context.Context, input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	return a.client.UpdateEventSourceMappingWithContext(ctx, input)
}

func (a *LambdaActivities) UpdateFunctionCode(ctx context.Context, input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error) {
	return a.client.UpdateFunctionCodeWithContext(ctx, input)
}

func (a *LambdaActivities) UpdateFunctionConfiguration(ctx context.Context, input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
	return a.client.UpdateFunctionConfigurationWithContext(ctx, input)
}

func (a *LambdaActivities) UpdateFunctionEventInvokeConfig(ctx context.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error) {
	return a.client.UpdateFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *LambdaActivities) WaitUntilFunctionActive(ctx context.Context, input *lambda.GetFunctionConfigurationInput) error {
	return a.client.WaitUntilFunctionActiveWithContext(ctx, input)

}

func (a *LambdaActivities) WaitUntilFunctionExists(ctx context.Context, input *lambda.GetFunctionInput) error {
	return a.client.WaitUntilFunctionExistsWithContext(ctx, input)

}

func (a *LambdaActivities) WaitUntilFunctionUpdated(ctx context.Context, input *lambda.GetFunctionConfigurationInput) error {
	return a.client.WaitUntilFunctionUpdatedWithContext(ctx, input)

}
