package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ApiGatewayV2Activities struct {
	client apigatewayv2iface.ApiGatewayV2API
}

func NewApiGatewayV2Activities(session *session.Session, config ...*aws.Config) *ApiGatewayV2Activities {
	client := apigatewayv2.New(session, config...)
	return &ApiGatewayV2Activities{client: client}
}

func (a *ApiGatewayV2Activities) CreateApi(ctx context.Context, input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error) {
	return a.client.CreateApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateApiMapping(ctx context.Context, input *apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error) {
	return a.client.CreateApiMappingWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateAuthorizer(ctx context.Context, input *apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error) {
	return a.client.CreateAuthorizerWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateDeployment(ctx context.Context, input *apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error) {
	return a.client.CreateDeploymentWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateDomainName(ctx context.Context, input *apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error) {
	return a.client.CreateDomainNameWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateIntegration(ctx context.Context, input *apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error) {
	return a.client.CreateIntegrationWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateIntegrationResponse(ctx context.Context, input *apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error) {
	return a.client.CreateIntegrationResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateModel(ctx context.Context, input *apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error) {
	return a.client.CreateModelWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateRoute(ctx context.Context, input *apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error) {
	return a.client.CreateRouteWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateRouteResponse(ctx context.Context, input *apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error) {
	return a.client.CreateRouteResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateStage(ctx context.Context, input *apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error) {
	return a.client.CreateStageWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) CreateVpcLink(ctx context.Context, input *apigatewayv2.CreateVpcLinkInput) (*apigatewayv2.CreateVpcLinkOutput, error) {
	return a.client.CreateVpcLinkWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteAccessLogSettings(ctx context.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) (*apigatewayv2.DeleteAccessLogSettingsOutput, error) {
	return a.client.DeleteAccessLogSettingsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteApi(ctx context.Context, input *apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error) {
	return a.client.DeleteApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteApiMapping(ctx context.Context, input *apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error) {
	return a.client.DeleteApiMappingWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteAuthorizer(ctx context.Context, input *apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error) {
	return a.client.DeleteAuthorizerWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteCorsConfiguration(ctx context.Context, input *apigatewayv2.DeleteCorsConfigurationInput) (*apigatewayv2.DeleteCorsConfigurationOutput, error) {
	return a.client.DeleteCorsConfigurationWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteDeployment(ctx context.Context, input *apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error) {
	return a.client.DeleteDeploymentWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteDomainName(ctx context.Context, input *apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error) {
	return a.client.DeleteDomainNameWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteIntegration(ctx context.Context, input *apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error) {
	return a.client.DeleteIntegrationWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteIntegrationResponse(ctx context.Context, input *apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error) {
	return a.client.DeleteIntegrationResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteModel(ctx context.Context, input *apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error) {
	return a.client.DeleteModelWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteRoute(ctx context.Context, input *apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error) {
	return a.client.DeleteRouteWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteRouteRequestParameter(ctx context.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) (*apigatewayv2.DeleteRouteRequestParameterOutput, error) {
	return a.client.DeleteRouteRequestParameterWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteRouteResponse(ctx context.Context, input *apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error) {
	return a.client.DeleteRouteResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteRouteSettings(ctx context.Context, input *apigatewayv2.DeleteRouteSettingsInput) (*apigatewayv2.DeleteRouteSettingsOutput, error) {
	return a.client.DeleteRouteSettingsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteStage(ctx context.Context, input *apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error) {
	return a.client.DeleteStageWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) DeleteVpcLink(ctx context.Context, input *apigatewayv2.DeleteVpcLinkInput) (*apigatewayv2.DeleteVpcLinkOutput, error) {
	return a.client.DeleteVpcLinkWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) ExportApi(ctx context.Context, input *apigatewayv2.ExportApiInput) (*apigatewayv2.ExportApiOutput, error) {
	return a.client.ExportApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetApi(ctx context.Context, input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error) {
	return a.client.GetApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetApiMapping(ctx context.Context, input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error) {
	return a.client.GetApiMappingWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetApiMappings(ctx context.Context, input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error) {
	return a.client.GetApiMappingsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetApis(ctx context.Context, input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error) {
	return a.client.GetApisWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetAuthorizer(ctx context.Context, input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error) {
	return a.client.GetAuthorizerWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetAuthorizers(ctx context.Context, input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error) {
	return a.client.GetAuthorizersWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetDeployment(ctx context.Context, input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error) {
	return a.client.GetDeploymentWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetDeployments(ctx context.Context, input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error) {
	return a.client.GetDeploymentsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetDomainName(ctx context.Context, input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error) {
	return a.client.GetDomainNameWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetDomainNames(ctx context.Context, input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error) {
	return a.client.GetDomainNamesWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetIntegration(ctx context.Context, input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error) {
	return a.client.GetIntegrationWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetIntegrationResponse(ctx context.Context, input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	return a.client.GetIntegrationResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetIntegrationResponses(ctx context.Context, input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	return a.client.GetIntegrationResponsesWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetIntegrations(ctx context.Context, input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error) {
	return a.client.GetIntegrationsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetModel(ctx context.Context, input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error) {
	return a.client.GetModelWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetModelTemplate(ctx context.Context, input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error) {
	return a.client.GetModelTemplateWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetModels(ctx context.Context, input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error) {
	return a.client.GetModelsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetRoute(ctx context.Context, input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error) {
	return a.client.GetRouteWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetRouteResponse(ctx context.Context, input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error) {
	return a.client.GetRouteResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetRouteResponses(ctx context.Context, input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error) {
	return a.client.GetRouteResponsesWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetRoutes(ctx context.Context, input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error) {
	return a.client.GetRoutesWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetStage(ctx context.Context, input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error) {
	return a.client.GetStageWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetStages(ctx context.Context, input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error) {
	return a.client.GetStagesWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetTags(ctx context.Context, input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error) {
	return a.client.GetTagsWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetVpcLink(ctx context.Context, input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error) {
	return a.client.GetVpcLinkWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) GetVpcLinks(ctx context.Context, input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error) {
	return a.client.GetVpcLinksWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) ImportApi(ctx context.Context, input *apigatewayv2.ImportApiInput) (*apigatewayv2.ImportApiOutput, error) {
	return a.client.ImportApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) ReimportApi(ctx context.Context, input *apigatewayv2.ReimportApiInput) (*apigatewayv2.ReimportApiOutput, error) {
	return a.client.ReimportApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) TagResource(ctx context.Context, input *apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UntagResource(ctx context.Context, input *apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateApi(ctx context.Context, input *apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error) {
	return a.client.UpdateApiWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateApiMapping(ctx context.Context, input *apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error) {
	return a.client.UpdateApiMappingWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateAuthorizer(ctx context.Context, input *apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error) {
	return a.client.UpdateAuthorizerWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateDeployment(ctx context.Context, input *apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error) {
	return a.client.UpdateDeploymentWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateDomainName(ctx context.Context, input *apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error) {
	return a.client.UpdateDomainNameWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateIntegration(ctx context.Context, input *apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error) {
	return a.client.UpdateIntegrationWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateIntegrationResponse(ctx context.Context, input *apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error) {
	return a.client.UpdateIntegrationResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateModel(ctx context.Context, input *apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error) {
	return a.client.UpdateModelWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateRoute(ctx context.Context, input *apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error) {
	return a.client.UpdateRouteWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateRouteResponse(ctx context.Context, input *apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error) {
	return a.client.UpdateRouteResponseWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateStage(ctx context.Context, input *apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error) {
	return a.client.UpdateStageWithContext(ctx, input)
}

func (a *ApiGatewayV2Activities) UpdateVpcLink(ctx context.Context, input *apigatewayv2.UpdateVpcLinkInput) (*apigatewayv2.UpdateVpcLinkOutput, error) {
	return a.client.UpdateVpcLinkWithContext(ctx, input)
}
