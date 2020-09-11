
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot/iotiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTActivities struct {
    client iotiface.IoTAPI
}

func NewIoTActivities(session *session.Session, config ...*aws.Config) *IoTActivities {
    client := iot.New(session, config...)
    return &IoTActivities{client: client}
}

func (a *IoTActivities) AcceptCertificateTransfer(ctx context.Context, input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error) {
    return a.client.AcceptCertificateTransferWithContext(ctx, input)
}

func (a *IoTActivities) AddThingToBillingGroup(ctx context.Context, input *iot.AddThingToBillingGroupInput) (*iot.AddThingToBillingGroupOutput, error) {
    return a.client.AddThingToBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) AddThingToThingGroup(ctx context.Context, input *iot.AddThingToThingGroupInput) (*iot.AddThingToThingGroupOutput, error) {
    return a.client.AddThingToThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) AssociateTargetsWithJob(ctx context.Context, input *iot.AssociateTargetsWithJobInput) (*iot.AssociateTargetsWithJobOutput, error) {
    return a.client.AssociateTargetsWithJobWithContext(ctx, input)
}

func (a *IoTActivities) AttachPolicy(ctx context.Context, input *iot.AttachPolicyInput) (*iot.AttachPolicyOutput, error) {
    return a.client.AttachPolicyWithContext(ctx, input)
}

func (a *IoTActivities) AttachPrincipalPolicy(ctx context.Context, input *iot.AttachPrincipalPolicyInput) (*iot.AttachPrincipalPolicyOutput, error) {
    return a.client.AttachPrincipalPolicyWithContext(ctx, input)
}

func (a *IoTActivities) AttachSecurityProfile(ctx context.Context, input *iot.AttachSecurityProfileInput) (*iot.AttachSecurityProfileOutput, error) {
    return a.client.AttachSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) AttachThingPrincipal(ctx context.Context, input *iot.AttachThingPrincipalInput) (*iot.AttachThingPrincipalOutput, error) {
    return a.client.AttachThingPrincipalWithContext(ctx, input)
}

func (a *IoTActivities) CancelAuditMitigationActionsTask(ctx context.Context, input *iot.CancelAuditMitigationActionsTaskInput) (*iot.CancelAuditMitigationActionsTaskOutput, error) {
    return a.client.CancelAuditMitigationActionsTaskWithContext(ctx, input)
}

func (a *IoTActivities) CancelAuditTask(ctx context.Context, input *iot.CancelAuditTaskInput) (*iot.CancelAuditTaskOutput, error) {
    return a.client.CancelAuditTaskWithContext(ctx, input)
}

func (a *IoTActivities) CancelCertificateTransfer(ctx context.Context, input *iot.CancelCertificateTransferInput) (*iot.CancelCertificateTransferOutput, error) {
    return a.client.CancelCertificateTransferWithContext(ctx, input)
}

func (a *IoTActivities) CancelJob(ctx context.Context, input *iot.CancelJobInput) (*iot.CancelJobOutput, error) {
    return a.client.CancelJobWithContext(ctx, input)
}

func (a *IoTActivities) CancelJobExecution(ctx context.Context, input *iot.CancelJobExecutionInput) (*iot.CancelJobExecutionOutput, error) {
    return a.client.CancelJobExecutionWithContext(ctx, input)
}

func (a *IoTActivities) ClearDefaultAuthorizer(ctx context.Context, input *iot.ClearDefaultAuthorizerInput) (*iot.ClearDefaultAuthorizerOutput, error) {
    return a.client.ClearDefaultAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) ConfirmTopicRuleDestination(ctx context.Context, input *iot.ConfirmTopicRuleDestinationInput) (*iot.ConfirmTopicRuleDestinationOutput, error) {
    return a.client.ConfirmTopicRuleDestinationWithContext(ctx, input)
}

func (a *IoTActivities) CreateAuditSuppression(ctx context.Context, input *iot.CreateAuditSuppressionInput) (*iot.CreateAuditSuppressionOutput, error) {
    return a.client.CreateAuditSuppressionWithContext(ctx, input)
}

func (a *IoTActivities) CreateAuthorizer(ctx context.Context, input *iot.CreateAuthorizerInput) (*iot.CreateAuthorizerOutput, error) {
    return a.client.CreateAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) CreateBillingGroup(ctx context.Context, input *iot.CreateBillingGroupInput) (*iot.CreateBillingGroupOutput, error) {
    return a.client.CreateBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) CreateCertificateFromCsr(ctx context.Context, input *iot.CreateCertificateFromCsrInput) (*iot.CreateCertificateFromCsrOutput, error) {
    return a.client.CreateCertificateFromCsrWithContext(ctx, input)
}

func (a *IoTActivities) CreateDimension(ctx context.Context, input *iot.CreateDimensionInput) (*iot.CreateDimensionOutput, error) {
    return a.client.CreateDimensionWithContext(ctx, input)
}

func (a *IoTActivities) CreateDomainConfiguration(ctx context.Context, input *iot.CreateDomainConfigurationInput) (*iot.CreateDomainConfigurationOutput, error) {
    return a.client.CreateDomainConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) CreateDynamicThingGroup(ctx context.Context, input *iot.CreateDynamicThingGroupInput) (*iot.CreateDynamicThingGroupOutput, error) {
    return a.client.CreateDynamicThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) CreateJob(ctx context.Context, input *iot.CreateJobInput) (*iot.CreateJobOutput, error) {
    return a.client.CreateJobWithContext(ctx, input)
}

func (a *IoTActivities) CreateKeysAndCertificate(ctx context.Context, input *iot.CreateKeysAndCertificateInput) (*iot.CreateKeysAndCertificateOutput, error) {
    return a.client.CreateKeysAndCertificateWithContext(ctx, input)
}

func (a *IoTActivities) CreateMitigationAction(ctx context.Context, input *iot.CreateMitigationActionInput) (*iot.CreateMitigationActionOutput, error) {
    return a.client.CreateMitigationActionWithContext(ctx, input)
}

func (a *IoTActivities) CreateOTAUpdate(ctx context.Context, input *iot.CreateOTAUpdateInput) (*iot.CreateOTAUpdateOutput, error) {
    return a.client.CreateOTAUpdateWithContext(ctx, input)
}

func (a *IoTActivities) CreatePolicy(ctx context.Context, input *iot.CreatePolicyInput) (*iot.CreatePolicyOutput, error) {
    return a.client.CreatePolicyWithContext(ctx, input)
}

func (a *IoTActivities) CreatePolicyVersion(ctx context.Context, input *iot.CreatePolicyVersionInput) (*iot.CreatePolicyVersionOutput, error) {
    return a.client.CreatePolicyVersionWithContext(ctx, input)
}

func (a *IoTActivities) CreateProvisioningClaim(ctx context.Context, input *iot.CreateProvisioningClaimInput) (*iot.CreateProvisioningClaimOutput, error) {
    return a.client.CreateProvisioningClaimWithContext(ctx, input)
}

func (a *IoTActivities) CreateProvisioningTemplate(ctx context.Context, input *iot.CreateProvisioningTemplateInput) (*iot.CreateProvisioningTemplateOutput, error) {
    return a.client.CreateProvisioningTemplateWithContext(ctx, input)
}

func (a *IoTActivities) CreateProvisioningTemplateVersion(ctx context.Context, input *iot.CreateProvisioningTemplateVersionInput) (*iot.CreateProvisioningTemplateVersionOutput, error) {
    return a.client.CreateProvisioningTemplateVersionWithContext(ctx, input)
}

func (a *IoTActivities) CreateRoleAlias(ctx context.Context, input *iot.CreateRoleAliasInput) (*iot.CreateRoleAliasOutput, error) {
    return a.client.CreateRoleAliasWithContext(ctx, input)
}

func (a *IoTActivities) CreateScheduledAudit(ctx context.Context, input *iot.CreateScheduledAuditInput) (*iot.CreateScheduledAuditOutput, error) {
    return a.client.CreateScheduledAuditWithContext(ctx, input)
}

func (a *IoTActivities) CreateSecurityProfile(ctx context.Context, input *iot.CreateSecurityProfileInput) (*iot.CreateSecurityProfileOutput, error) {
    return a.client.CreateSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) CreateStream(ctx context.Context, input *iot.CreateStreamInput) (*iot.CreateStreamOutput, error) {
    return a.client.CreateStreamWithContext(ctx, input)
}

func (a *IoTActivities) CreateThing(ctx context.Context, input *iot.CreateThingInput) (*iot.CreateThingOutput, error) {
    return a.client.CreateThingWithContext(ctx, input)
}

func (a *IoTActivities) CreateThingGroup(ctx context.Context, input *iot.CreateThingGroupInput) (*iot.CreateThingGroupOutput, error) {
    return a.client.CreateThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) CreateThingType(ctx context.Context, input *iot.CreateThingTypeInput) (*iot.CreateThingTypeOutput, error) {
    return a.client.CreateThingTypeWithContext(ctx, input)
}

func (a *IoTActivities) CreateTopicRule(ctx context.Context, input *iot.CreateTopicRuleInput) (*iot.CreateTopicRuleOutput, error) {
    return a.client.CreateTopicRuleWithContext(ctx, input)
}

func (a *IoTActivities) CreateTopicRuleDestination(ctx context.Context, input *iot.CreateTopicRuleDestinationInput) (*iot.CreateTopicRuleDestinationOutput, error) {
    return a.client.CreateTopicRuleDestinationWithContext(ctx, input)
}

func (a *IoTActivities) DeleteAccountAuditConfiguration(ctx context.Context, input *iot.DeleteAccountAuditConfigurationInput) (*iot.DeleteAccountAuditConfigurationOutput, error) {
    return a.client.DeleteAccountAuditConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) DeleteAuditSuppression(ctx context.Context, input *iot.DeleteAuditSuppressionInput) (*iot.DeleteAuditSuppressionOutput, error) {
    return a.client.DeleteAuditSuppressionWithContext(ctx, input)
}

func (a *IoTActivities) DeleteAuthorizer(ctx context.Context, input *iot.DeleteAuthorizerInput) (*iot.DeleteAuthorizerOutput, error) {
    return a.client.DeleteAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) DeleteBillingGroup(ctx context.Context, input *iot.DeleteBillingGroupInput) (*iot.DeleteBillingGroupOutput, error) {
    return a.client.DeleteBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) DeleteCACertificate(ctx context.Context, input *iot.DeleteCACertificateInput) (*iot.DeleteCACertificateOutput, error) {
    return a.client.DeleteCACertificateWithContext(ctx, input)
}

func (a *IoTActivities) DeleteCertificate(ctx context.Context, input *iot.DeleteCertificateInput) (*iot.DeleteCertificateOutput, error) {
    return a.client.DeleteCertificateWithContext(ctx, input)
}

func (a *IoTActivities) DeleteDimension(ctx context.Context, input *iot.DeleteDimensionInput) (*iot.DeleteDimensionOutput, error) {
    return a.client.DeleteDimensionWithContext(ctx, input)
}

func (a *IoTActivities) DeleteDomainConfiguration(ctx context.Context, input *iot.DeleteDomainConfigurationInput) (*iot.DeleteDomainConfigurationOutput, error) {
    return a.client.DeleteDomainConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) DeleteDynamicThingGroup(ctx context.Context, input *iot.DeleteDynamicThingGroupInput) (*iot.DeleteDynamicThingGroupOutput, error) {
    return a.client.DeleteDynamicThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) DeleteJob(ctx context.Context, input *iot.DeleteJobInput) (*iot.DeleteJobOutput, error) {
    return a.client.DeleteJobWithContext(ctx, input)
}

func (a *IoTActivities) DeleteJobExecution(ctx context.Context, input *iot.DeleteJobExecutionInput) (*iot.DeleteJobExecutionOutput, error) {
    return a.client.DeleteJobExecutionWithContext(ctx, input)
}

func (a *IoTActivities) DeleteMitigationAction(ctx context.Context, input *iot.DeleteMitigationActionInput) (*iot.DeleteMitigationActionOutput, error) {
    return a.client.DeleteMitigationActionWithContext(ctx, input)
}

func (a *IoTActivities) DeleteOTAUpdate(ctx context.Context, input *iot.DeleteOTAUpdateInput) (*iot.DeleteOTAUpdateOutput, error) {
    return a.client.DeleteOTAUpdateWithContext(ctx, input)
}

func (a *IoTActivities) DeletePolicy(ctx context.Context, input *iot.DeletePolicyInput) (*iot.DeletePolicyOutput, error) {
    return a.client.DeletePolicyWithContext(ctx, input)
}

func (a *IoTActivities) DeletePolicyVersion(ctx context.Context, input *iot.DeletePolicyVersionInput) (*iot.DeletePolicyVersionOutput, error) {
    return a.client.DeletePolicyVersionWithContext(ctx, input)
}

func (a *IoTActivities) DeleteProvisioningTemplate(ctx context.Context, input *iot.DeleteProvisioningTemplateInput) (*iot.DeleteProvisioningTemplateOutput, error) {
    return a.client.DeleteProvisioningTemplateWithContext(ctx, input)
}

func (a *IoTActivities) DeleteProvisioningTemplateVersion(ctx context.Context, input *iot.DeleteProvisioningTemplateVersionInput) (*iot.DeleteProvisioningTemplateVersionOutput, error) {
    return a.client.DeleteProvisioningTemplateVersionWithContext(ctx, input)
}

func (a *IoTActivities) DeleteRegistrationCode(ctx context.Context, input *iot.DeleteRegistrationCodeInput) (*iot.DeleteRegistrationCodeOutput, error) {
    return a.client.DeleteRegistrationCodeWithContext(ctx, input)
}

func (a *IoTActivities) DeleteRoleAlias(ctx context.Context, input *iot.DeleteRoleAliasInput) (*iot.DeleteRoleAliasOutput, error) {
    return a.client.DeleteRoleAliasWithContext(ctx, input)
}

func (a *IoTActivities) DeleteScheduledAudit(ctx context.Context, input *iot.DeleteScheduledAuditInput) (*iot.DeleteScheduledAuditOutput, error) {
    return a.client.DeleteScheduledAuditWithContext(ctx, input)
}

func (a *IoTActivities) DeleteSecurityProfile(ctx context.Context, input *iot.DeleteSecurityProfileInput) (*iot.DeleteSecurityProfileOutput, error) {
    return a.client.DeleteSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) DeleteStream(ctx context.Context, input *iot.DeleteStreamInput) (*iot.DeleteStreamOutput, error) {
    return a.client.DeleteStreamWithContext(ctx, input)
}

func (a *IoTActivities) DeleteThing(ctx context.Context, input *iot.DeleteThingInput) (*iot.DeleteThingOutput, error) {
    return a.client.DeleteThingWithContext(ctx, input)
}

func (a *IoTActivities) DeleteThingGroup(ctx context.Context, input *iot.DeleteThingGroupInput) (*iot.DeleteThingGroupOutput, error) {
    return a.client.DeleteThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) DeleteThingType(ctx context.Context, input *iot.DeleteThingTypeInput) (*iot.DeleteThingTypeOutput, error) {
    return a.client.DeleteThingTypeWithContext(ctx, input)
}

func (a *IoTActivities) DeleteTopicRule(ctx context.Context, input *iot.DeleteTopicRuleInput) (*iot.DeleteTopicRuleOutput, error) {
    return a.client.DeleteTopicRuleWithContext(ctx, input)
}

func (a *IoTActivities) DeleteTopicRuleDestination(ctx context.Context, input *iot.DeleteTopicRuleDestinationInput) (*iot.DeleteTopicRuleDestinationOutput, error) {
    return a.client.DeleteTopicRuleDestinationWithContext(ctx, input)
}

func (a *IoTActivities) DeleteV2LoggingLevel(ctx context.Context, input *iot.DeleteV2LoggingLevelInput) (*iot.DeleteV2LoggingLevelOutput, error) {
    return a.client.DeleteV2LoggingLevelWithContext(ctx, input)
}

func (a *IoTActivities) DeprecateThingType(ctx context.Context, input *iot.DeprecateThingTypeInput) (*iot.DeprecateThingTypeOutput, error) {
    return a.client.DeprecateThingTypeWithContext(ctx, input)
}

func (a *IoTActivities) DescribeAccountAuditConfiguration(ctx context.Context, input *iot.DescribeAccountAuditConfigurationInput) (*iot.DescribeAccountAuditConfigurationOutput, error) {
    return a.client.DescribeAccountAuditConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) DescribeAuditFinding(ctx context.Context, input *iot.DescribeAuditFindingInput) (*iot.DescribeAuditFindingOutput, error) {
    return a.client.DescribeAuditFindingWithContext(ctx, input)
}

func (a *IoTActivities) DescribeAuditMitigationActionsTask(ctx context.Context, input *iot.DescribeAuditMitigationActionsTaskInput) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
    return a.client.DescribeAuditMitigationActionsTaskWithContext(ctx, input)
}

func (a *IoTActivities) DescribeAuditSuppression(ctx context.Context, input *iot.DescribeAuditSuppressionInput) (*iot.DescribeAuditSuppressionOutput, error) {
    return a.client.DescribeAuditSuppressionWithContext(ctx, input)
}

func (a *IoTActivities) DescribeAuditTask(ctx context.Context, input *iot.DescribeAuditTaskInput) (*iot.DescribeAuditTaskOutput, error) {
    return a.client.DescribeAuditTaskWithContext(ctx, input)
}

func (a *IoTActivities) DescribeAuthorizer(ctx context.Context, input *iot.DescribeAuthorizerInput) (*iot.DescribeAuthorizerOutput, error) {
    return a.client.DescribeAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) DescribeBillingGroup(ctx context.Context, input *iot.DescribeBillingGroupInput) (*iot.DescribeBillingGroupOutput, error) {
    return a.client.DescribeBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) DescribeCACertificate(ctx context.Context, input *iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error) {
    return a.client.DescribeCACertificateWithContext(ctx, input)
}

func (a *IoTActivities) DescribeCertificate(ctx context.Context, input *iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error) {
    return a.client.DescribeCertificateWithContext(ctx, input)
}

func (a *IoTActivities) DescribeDefaultAuthorizer(ctx context.Context, input *iot.DescribeDefaultAuthorizerInput) (*iot.DescribeDefaultAuthorizerOutput, error) {
    return a.client.DescribeDefaultAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) DescribeDimension(ctx context.Context, input *iot.DescribeDimensionInput) (*iot.DescribeDimensionOutput, error) {
    return a.client.DescribeDimensionWithContext(ctx, input)
}

func (a *IoTActivities) DescribeDomainConfiguration(ctx context.Context, input *iot.DescribeDomainConfigurationInput) (*iot.DescribeDomainConfigurationOutput, error) {
    return a.client.DescribeDomainConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) DescribeEndpoint(ctx context.Context, input *iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error) {
    return a.client.DescribeEndpointWithContext(ctx, input)
}

func (a *IoTActivities) DescribeEventConfigurations(ctx context.Context, input *iot.DescribeEventConfigurationsInput) (*iot.DescribeEventConfigurationsOutput, error) {
    return a.client.DescribeEventConfigurationsWithContext(ctx, input)
}

func (a *IoTActivities) DescribeIndex(ctx context.Context, input *iot.DescribeIndexInput) (*iot.DescribeIndexOutput, error) {
    return a.client.DescribeIndexWithContext(ctx, input)
}

func (a *IoTActivities) DescribeJob(ctx context.Context, input *iot.DescribeJobInput) (*iot.DescribeJobOutput, error) {
    return a.client.DescribeJobWithContext(ctx, input)
}

func (a *IoTActivities) DescribeJobExecution(ctx context.Context, input *iot.DescribeJobExecutionInput) (*iot.DescribeJobExecutionOutput, error) {
    return a.client.DescribeJobExecutionWithContext(ctx, input)
}

func (a *IoTActivities) DescribeMitigationAction(ctx context.Context, input *iot.DescribeMitigationActionInput) (*iot.DescribeMitigationActionOutput, error) {
    return a.client.DescribeMitigationActionWithContext(ctx, input)
}

func (a *IoTActivities) DescribeProvisioningTemplate(ctx context.Context, input *iot.DescribeProvisioningTemplateInput) (*iot.DescribeProvisioningTemplateOutput, error) {
    return a.client.DescribeProvisioningTemplateWithContext(ctx, input)
}

func (a *IoTActivities) DescribeProvisioningTemplateVersion(ctx context.Context, input *iot.DescribeProvisioningTemplateVersionInput) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
    return a.client.DescribeProvisioningTemplateVersionWithContext(ctx, input)
}

func (a *IoTActivities) DescribeRoleAlias(ctx context.Context, input *iot.DescribeRoleAliasInput) (*iot.DescribeRoleAliasOutput, error) {
    return a.client.DescribeRoleAliasWithContext(ctx, input)
}

func (a *IoTActivities) DescribeScheduledAudit(ctx context.Context, input *iot.DescribeScheduledAuditInput) (*iot.DescribeScheduledAuditOutput, error) {
    return a.client.DescribeScheduledAuditWithContext(ctx, input)
}

func (a *IoTActivities) DescribeSecurityProfile(ctx context.Context, input *iot.DescribeSecurityProfileInput) (*iot.DescribeSecurityProfileOutput, error) {
    return a.client.DescribeSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) DescribeStream(ctx context.Context, input *iot.DescribeStreamInput) (*iot.DescribeStreamOutput, error) {
    return a.client.DescribeStreamWithContext(ctx, input)
}

func (a *IoTActivities) DescribeThing(ctx context.Context, input *iot.DescribeThingInput) (*iot.DescribeThingOutput, error) {
    return a.client.DescribeThingWithContext(ctx, input)
}

func (a *IoTActivities) DescribeThingGroup(ctx context.Context, input *iot.DescribeThingGroupInput) (*iot.DescribeThingGroupOutput, error) {
    return a.client.DescribeThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) DescribeThingRegistrationTask(ctx context.Context, input *iot.DescribeThingRegistrationTaskInput) (*iot.DescribeThingRegistrationTaskOutput, error) {
    return a.client.DescribeThingRegistrationTaskWithContext(ctx, input)
}

func (a *IoTActivities) DescribeThingType(ctx context.Context, input *iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error) {
    return a.client.DescribeThingTypeWithContext(ctx, input)
}

func (a *IoTActivities) DetachPolicy(ctx context.Context, input *iot.DetachPolicyInput) (*iot.DetachPolicyOutput, error) {
    return a.client.DetachPolicyWithContext(ctx, input)
}

func (a *IoTActivities) DetachPrincipalPolicy(ctx context.Context, input *iot.DetachPrincipalPolicyInput) (*iot.DetachPrincipalPolicyOutput, error) {
    return a.client.DetachPrincipalPolicyWithContext(ctx, input)
}

func (a *IoTActivities) DetachSecurityProfile(ctx context.Context, input *iot.DetachSecurityProfileInput) (*iot.DetachSecurityProfileOutput, error) {
    return a.client.DetachSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) DetachThingPrincipal(ctx context.Context, input *iot.DetachThingPrincipalInput) (*iot.DetachThingPrincipalOutput, error) {
    return a.client.DetachThingPrincipalWithContext(ctx, input)
}

func (a *IoTActivities) DisableTopicRule(ctx context.Context, input *iot.DisableTopicRuleInput) (*iot.DisableTopicRuleOutput, error) {
    return a.client.DisableTopicRuleWithContext(ctx, input)
}

func (a *IoTActivities) EnableTopicRule(ctx context.Context, input *iot.EnableTopicRuleInput) (*iot.EnableTopicRuleOutput, error) {
    return a.client.EnableTopicRuleWithContext(ctx, input)
}

func (a *IoTActivities) GetCardinality(ctx context.Context, input *iot.GetCardinalityInput) (*iot.GetCardinalityOutput, error) {
    return a.client.GetCardinalityWithContext(ctx, input)
}

func (a *IoTActivities) GetEffectivePolicies(ctx context.Context, input *iot.GetEffectivePoliciesInput) (*iot.GetEffectivePoliciesOutput, error) {
    return a.client.GetEffectivePoliciesWithContext(ctx, input)
}

func (a *IoTActivities) GetIndexingConfiguration(ctx context.Context, input *iot.GetIndexingConfigurationInput) (*iot.GetIndexingConfigurationOutput, error) {
    return a.client.GetIndexingConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) GetJobDocument(ctx context.Context, input *iot.GetJobDocumentInput) (*iot.GetJobDocumentOutput, error) {
    return a.client.GetJobDocumentWithContext(ctx, input)
}

func (a *IoTActivities) GetLoggingOptions(ctx context.Context, input *iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error) {
    return a.client.GetLoggingOptionsWithContext(ctx, input)
}

func (a *IoTActivities) GetOTAUpdate(ctx context.Context, input *iot.GetOTAUpdateInput) (*iot.GetOTAUpdateOutput, error) {
    return a.client.GetOTAUpdateWithContext(ctx, input)
}

func (a *IoTActivities) GetPercentiles(ctx context.Context, input *iot.GetPercentilesInput) (*iot.GetPercentilesOutput, error) {
    return a.client.GetPercentilesWithContext(ctx, input)
}

func (a *IoTActivities) GetPolicy(ctx context.Context, input *iot.GetPolicyInput) (*iot.GetPolicyOutput, error) {
    return a.client.GetPolicyWithContext(ctx, input)
}

func (a *IoTActivities) GetPolicyVersion(ctx context.Context, input *iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error) {
    return a.client.GetPolicyVersionWithContext(ctx, input)
}

func (a *IoTActivities) GetRegistrationCode(ctx context.Context, input *iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error) {
    return a.client.GetRegistrationCodeWithContext(ctx, input)
}

func (a *IoTActivities) GetStatistics(ctx context.Context, input *iot.GetStatisticsInput) (*iot.GetStatisticsOutput, error) {
    return a.client.GetStatisticsWithContext(ctx, input)
}

func (a *IoTActivities) GetTopicRule(ctx context.Context, input *iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error) {
    return a.client.GetTopicRuleWithContext(ctx, input)
}

func (a *IoTActivities) GetTopicRuleDestination(ctx context.Context, input *iot.GetTopicRuleDestinationInput) (*iot.GetTopicRuleDestinationOutput, error) {
    return a.client.GetTopicRuleDestinationWithContext(ctx, input)
}

func (a *IoTActivities) GetV2LoggingOptions(ctx context.Context, input *iot.GetV2LoggingOptionsInput) (*iot.GetV2LoggingOptionsOutput, error) {
    return a.client.GetV2LoggingOptionsWithContext(ctx, input)
}

func (a *IoTActivities) ListActiveViolations(ctx context.Context, input *iot.ListActiveViolationsInput) (*iot.ListActiveViolationsOutput, error) {
    return a.client.ListActiveViolationsWithContext(ctx, input)
}

func (a *IoTActivities) ListAttachedPolicies(ctx context.Context, input *iot.ListAttachedPoliciesInput) (*iot.ListAttachedPoliciesOutput, error) {
    return a.client.ListAttachedPoliciesWithContext(ctx, input)
}

func (a *IoTActivities) ListAuditFindings(ctx context.Context, input *iot.ListAuditFindingsInput) (*iot.ListAuditFindingsOutput, error) {
    return a.client.ListAuditFindingsWithContext(ctx, input)
}

func (a *IoTActivities) ListAuditMitigationActionsExecutions(ctx context.Context, input *iot.ListAuditMitigationActionsExecutionsInput) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
    return a.client.ListAuditMitigationActionsExecutionsWithContext(ctx, input)
}

func (a *IoTActivities) ListAuditMitigationActionsTasks(ctx context.Context, input *iot.ListAuditMitigationActionsTasksInput) (*iot.ListAuditMitigationActionsTasksOutput, error) {
    return a.client.ListAuditMitigationActionsTasksWithContext(ctx, input)
}

func (a *IoTActivities) ListAuditSuppressions(ctx context.Context, input *iot.ListAuditSuppressionsInput) (*iot.ListAuditSuppressionsOutput, error) {
    return a.client.ListAuditSuppressionsWithContext(ctx, input)
}

func (a *IoTActivities) ListAuditTasks(ctx context.Context, input *iot.ListAuditTasksInput) (*iot.ListAuditTasksOutput, error) {
    return a.client.ListAuditTasksWithContext(ctx, input)
}

func (a *IoTActivities) ListAuthorizers(ctx context.Context, input *iot.ListAuthorizersInput) (*iot.ListAuthorizersOutput, error) {
    return a.client.ListAuthorizersWithContext(ctx, input)
}

func (a *IoTActivities) ListBillingGroups(ctx context.Context, input *iot.ListBillingGroupsInput) (*iot.ListBillingGroupsOutput, error) {
    return a.client.ListBillingGroupsWithContext(ctx, input)
}

func (a *IoTActivities) ListCACertificates(ctx context.Context, input *iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error) {
    return a.client.ListCACertificatesWithContext(ctx, input)
}

func (a *IoTActivities) ListCertificates(ctx context.Context, input *iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error) {
    return a.client.ListCertificatesWithContext(ctx, input)
}

func (a *IoTActivities) ListCertificatesByCA(ctx context.Context, input *iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error) {
    return a.client.ListCertificatesByCAWithContext(ctx, input)
}

func (a *IoTActivities) ListDimensions(ctx context.Context, input *iot.ListDimensionsInput) (*iot.ListDimensionsOutput, error) {
    return a.client.ListDimensionsWithContext(ctx, input)
}

func (a *IoTActivities) ListDomainConfigurations(ctx context.Context, input *iot.ListDomainConfigurationsInput) (*iot.ListDomainConfigurationsOutput, error) {
    return a.client.ListDomainConfigurationsWithContext(ctx, input)
}

func (a *IoTActivities) ListIndices(ctx context.Context, input *iot.ListIndicesInput) (*iot.ListIndicesOutput, error) {
    return a.client.ListIndicesWithContext(ctx, input)
}

func (a *IoTActivities) ListJobExecutionsForJob(ctx context.Context, input *iot.ListJobExecutionsForJobInput) (*iot.ListJobExecutionsForJobOutput, error) {
    return a.client.ListJobExecutionsForJobWithContext(ctx, input)
}

func (a *IoTActivities) ListJobExecutionsForThing(ctx context.Context, input *iot.ListJobExecutionsForThingInput) (*iot.ListJobExecutionsForThingOutput, error) {
    return a.client.ListJobExecutionsForThingWithContext(ctx, input)
}

func (a *IoTActivities) ListJobs(ctx context.Context, input *iot.ListJobsInput) (*iot.ListJobsOutput, error) {
    return a.client.ListJobsWithContext(ctx, input)
}

func (a *IoTActivities) ListMitigationActions(ctx context.Context, input *iot.ListMitigationActionsInput) (*iot.ListMitigationActionsOutput, error) {
    return a.client.ListMitigationActionsWithContext(ctx, input)
}

func (a *IoTActivities) ListOTAUpdates(ctx context.Context, input *iot.ListOTAUpdatesInput) (*iot.ListOTAUpdatesOutput, error) {
    return a.client.ListOTAUpdatesWithContext(ctx, input)
}

func (a *IoTActivities) ListOutgoingCertificates(ctx context.Context, input *iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error) {
    return a.client.ListOutgoingCertificatesWithContext(ctx, input)
}

func (a *IoTActivities) ListPolicies(ctx context.Context, input *iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error) {
    return a.client.ListPoliciesWithContext(ctx, input)
}

func (a *IoTActivities) ListPolicyPrincipals(ctx context.Context, input *iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error) {
    return a.client.ListPolicyPrincipalsWithContext(ctx, input)
}

func (a *IoTActivities) ListPolicyVersions(ctx context.Context, input *iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error) {
    return a.client.ListPolicyVersionsWithContext(ctx, input)
}

func (a *IoTActivities) ListPrincipalPolicies(ctx context.Context, input *iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error) {
    return a.client.ListPrincipalPoliciesWithContext(ctx, input)
}

func (a *IoTActivities) ListPrincipalThings(ctx context.Context, input *iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error) {
    return a.client.ListPrincipalThingsWithContext(ctx, input)
}

func (a *IoTActivities) ListProvisioningTemplateVersions(ctx context.Context, input *iot.ListProvisioningTemplateVersionsInput) (*iot.ListProvisioningTemplateVersionsOutput, error) {
    return a.client.ListProvisioningTemplateVersionsWithContext(ctx, input)
}

func (a *IoTActivities) ListProvisioningTemplates(ctx context.Context, input *iot.ListProvisioningTemplatesInput) (*iot.ListProvisioningTemplatesOutput, error) {
    return a.client.ListProvisioningTemplatesWithContext(ctx, input)
}

func (a *IoTActivities) ListRoleAliases(ctx context.Context, input *iot.ListRoleAliasesInput) (*iot.ListRoleAliasesOutput, error) {
    return a.client.ListRoleAliasesWithContext(ctx, input)
}

func (a *IoTActivities) ListScheduledAudits(ctx context.Context, input *iot.ListScheduledAuditsInput) (*iot.ListScheduledAuditsOutput, error) {
    return a.client.ListScheduledAuditsWithContext(ctx, input)
}

func (a *IoTActivities) ListSecurityProfiles(ctx context.Context, input *iot.ListSecurityProfilesInput) (*iot.ListSecurityProfilesOutput, error) {
    return a.client.ListSecurityProfilesWithContext(ctx, input)
}

func (a *IoTActivities) ListSecurityProfilesForTarget(ctx context.Context, input *iot.ListSecurityProfilesForTargetInput) (*iot.ListSecurityProfilesForTargetOutput, error) {
    return a.client.ListSecurityProfilesForTargetWithContext(ctx, input)
}

func (a *IoTActivities) ListStreams(ctx context.Context, input *iot.ListStreamsInput) (*iot.ListStreamsOutput, error) {
    return a.client.ListStreamsWithContext(ctx, input)
}

func (a *IoTActivities) ListTagsForResource(ctx context.Context, input *iot.ListTagsForResourceInput) (*iot.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTActivities) ListTargetsForPolicy(ctx context.Context, input *iot.ListTargetsForPolicyInput) (*iot.ListTargetsForPolicyOutput, error) {
    return a.client.ListTargetsForPolicyWithContext(ctx, input)
}

func (a *IoTActivities) ListTargetsForSecurityProfile(ctx context.Context, input *iot.ListTargetsForSecurityProfileInput) (*iot.ListTargetsForSecurityProfileOutput, error) {
    return a.client.ListTargetsForSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) ListThingGroups(ctx context.Context, input *iot.ListThingGroupsInput) (*iot.ListThingGroupsOutput, error) {
    return a.client.ListThingGroupsWithContext(ctx, input)
}

func (a *IoTActivities) ListThingGroupsForThing(ctx context.Context, input *iot.ListThingGroupsForThingInput) (*iot.ListThingGroupsForThingOutput, error) {
    return a.client.ListThingGroupsForThingWithContext(ctx, input)
}

func (a *IoTActivities) ListThingPrincipals(ctx context.Context, input *iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error) {
    return a.client.ListThingPrincipalsWithContext(ctx, input)
}

func (a *IoTActivities) ListThingRegistrationTaskReports(ctx context.Context, input *iot.ListThingRegistrationTaskReportsInput) (*iot.ListThingRegistrationTaskReportsOutput, error) {
    return a.client.ListThingRegistrationTaskReportsWithContext(ctx, input)
}

func (a *IoTActivities) ListThingRegistrationTasks(ctx context.Context, input *iot.ListThingRegistrationTasksInput) (*iot.ListThingRegistrationTasksOutput, error) {
    return a.client.ListThingRegistrationTasksWithContext(ctx, input)
}

func (a *IoTActivities) ListThingTypes(ctx context.Context, input *iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error) {
    return a.client.ListThingTypesWithContext(ctx, input)
}

func (a *IoTActivities) ListThings(ctx context.Context, input *iot.ListThingsInput) (*iot.ListThingsOutput, error) {
    return a.client.ListThingsWithContext(ctx, input)
}

func (a *IoTActivities) ListThingsInBillingGroup(ctx context.Context, input *iot.ListThingsInBillingGroupInput) (*iot.ListThingsInBillingGroupOutput, error) {
    return a.client.ListThingsInBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) ListThingsInThingGroup(ctx context.Context, input *iot.ListThingsInThingGroupInput) (*iot.ListThingsInThingGroupOutput, error) {
    return a.client.ListThingsInThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) ListTopicRuleDestinations(ctx context.Context, input *iot.ListTopicRuleDestinationsInput) (*iot.ListTopicRuleDestinationsOutput, error) {
    return a.client.ListTopicRuleDestinationsWithContext(ctx, input)
}

func (a *IoTActivities) ListTopicRules(ctx context.Context, input *iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error) {
    return a.client.ListTopicRulesWithContext(ctx, input)
}

func (a *IoTActivities) ListV2LoggingLevels(ctx context.Context, input *iot.ListV2LoggingLevelsInput) (*iot.ListV2LoggingLevelsOutput, error) {
    return a.client.ListV2LoggingLevelsWithContext(ctx, input)
}

func (a *IoTActivities) ListViolationEvents(ctx context.Context, input *iot.ListViolationEventsInput) (*iot.ListViolationEventsOutput, error) {
    return a.client.ListViolationEventsWithContext(ctx, input)
}

func (a *IoTActivities) RegisterCACertificate(ctx context.Context, input *iot.RegisterCACertificateInput) (*iot.RegisterCACertificateOutput, error) {
    return a.client.RegisterCACertificateWithContext(ctx, input)
}

func (a *IoTActivities) RegisterCertificate(ctx context.Context, input *iot.RegisterCertificateInput) (*iot.RegisterCertificateOutput, error) {
    return a.client.RegisterCertificateWithContext(ctx, input)
}

func (a *IoTActivities) RegisterCertificateWithoutCA(ctx context.Context, input *iot.RegisterCertificateWithoutCAInput) (*iot.RegisterCertificateWithoutCAOutput, error) {
    return a.client.RegisterCertificateWithoutCAWithContext(ctx, input)
}

func (a *IoTActivities) RegisterThing(ctx context.Context, input *iot.RegisterThingInput) (*iot.RegisterThingOutput, error) {
    return a.client.RegisterThingWithContext(ctx, input)
}

func (a *IoTActivities) RejectCertificateTransfer(ctx context.Context, input *iot.RejectCertificateTransferInput) (*iot.RejectCertificateTransferOutput, error) {
    return a.client.RejectCertificateTransferWithContext(ctx, input)
}

func (a *IoTActivities) RemoveThingFromBillingGroup(ctx context.Context, input *iot.RemoveThingFromBillingGroupInput) (*iot.RemoveThingFromBillingGroupOutput, error) {
    return a.client.RemoveThingFromBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) RemoveThingFromThingGroup(ctx context.Context, input *iot.RemoveThingFromThingGroupInput) (*iot.RemoveThingFromThingGroupOutput, error) {
    return a.client.RemoveThingFromThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) ReplaceTopicRule(ctx context.Context, input *iot.ReplaceTopicRuleInput) (*iot.ReplaceTopicRuleOutput, error) {
    return a.client.ReplaceTopicRuleWithContext(ctx, input)
}

func (a *IoTActivities) SearchIndex(ctx context.Context, input *iot.SearchIndexInput) (*iot.SearchIndexOutput, error) {
    return a.client.SearchIndexWithContext(ctx, input)
}

func (a *IoTActivities) SetDefaultAuthorizer(ctx context.Context, input *iot.SetDefaultAuthorizerInput) (*iot.SetDefaultAuthorizerOutput, error) {
    return a.client.SetDefaultAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) SetDefaultPolicyVersion(ctx context.Context, input *iot.SetDefaultPolicyVersionInput) (*iot.SetDefaultPolicyVersionOutput, error) {
    return a.client.SetDefaultPolicyVersionWithContext(ctx, input)
}

func (a *IoTActivities) SetLoggingOptions(ctx context.Context, input *iot.SetLoggingOptionsInput) (*iot.SetLoggingOptionsOutput, error) {
    return a.client.SetLoggingOptionsWithContext(ctx, input)
}

func (a *IoTActivities) SetV2LoggingLevel(ctx context.Context, input *iot.SetV2LoggingLevelInput) (*iot.SetV2LoggingLevelOutput, error) {
    return a.client.SetV2LoggingLevelWithContext(ctx, input)
}

func (a *IoTActivities) SetV2LoggingOptions(ctx context.Context, input *iot.SetV2LoggingOptionsInput) (*iot.SetV2LoggingOptionsOutput, error) {
    return a.client.SetV2LoggingOptionsWithContext(ctx, input)
}

func (a *IoTActivities) StartAuditMitigationActionsTask(ctx context.Context, input *iot.StartAuditMitigationActionsTaskInput) (*iot.StartAuditMitigationActionsTaskOutput, error) {
    return a.client.StartAuditMitigationActionsTaskWithContext(ctx, input)
}

func (a *IoTActivities) StartOnDemandAuditTask(ctx context.Context, input *iot.StartOnDemandAuditTaskInput) (*iot.StartOnDemandAuditTaskOutput, error) {
    return a.client.StartOnDemandAuditTaskWithContext(ctx, input)
}

func (a *IoTActivities) StartThingRegistrationTask(ctx context.Context, input *iot.StartThingRegistrationTaskInput) (*iot.StartThingRegistrationTaskOutput, error) {
    return a.client.StartThingRegistrationTaskWithContext(ctx, input)
}

func (a *IoTActivities) StopThingRegistrationTask(ctx context.Context, input *iot.StopThingRegistrationTaskInput) (*iot.StopThingRegistrationTaskOutput, error) {
    return a.client.StopThingRegistrationTaskWithContext(ctx, input)
}

func (a *IoTActivities) TagResource(ctx context.Context, input *iot.TagResourceInput) (*iot.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoTActivities) TestAuthorization(ctx context.Context, input *iot.TestAuthorizationInput) (*iot.TestAuthorizationOutput, error) {
    return a.client.TestAuthorizationWithContext(ctx, input)
}

func (a *IoTActivities) TestInvokeAuthorizer(ctx context.Context, input *iot.TestInvokeAuthorizerInput) (*iot.TestInvokeAuthorizerOutput, error) {
    return a.client.TestInvokeAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) TransferCertificate(ctx context.Context, input *iot.TransferCertificateInput) (*iot.TransferCertificateOutput, error) {
    return a.client.TransferCertificateWithContext(ctx, input)
}

func (a *IoTActivities) UntagResource(ctx context.Context, input *iot.UntagResourceInput) (*iot.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoTActivities) UpdateAccountAuditConfiguration(ctx context.Context, input *iot.UpdateAccountAuditConfigurationInput) (*iot.UpdateAccountAuditConfigurationOutput, error) {
    return a.client.UpdateAccountAuditConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) UpdateAuditSuppression(ctx context.Context, input *iot.UpdateAuditSuppressionInput) (*iot.UpdateAuditSuppressionOutput, error) {
    return a.client.UpdateAuditSuppressionWithContext(ctx, input)
}

func (a *IoTActivities) UpdateAuthorizer(ctx context.Context, input *iot.UpdateAuthorizerInput) (*iot.UpdateAuthorizerOutput, error) {
    return a.client.UpdateAuthorizerWithContext(ctx, input)
}

func (a *IoTActivities) UpdateBillingGroup(ctx context.Context, input *iot.UpdateBillingGroupInput) (*iot.UpdateBillingGroupOutput, error) {
    return a.client.UpdateBillingGroupWithContext(ctx, input)
}

func (a *IoTActivities) UpdateCACertificate(ctx context.Context, input *iot.UpdateCACertificateInput) (*iot.UpdateCACertificateOutput, error) {
    return a.client.UpdateCACertificateWithContext(ctx, input)
}

func (a *IoTActivities) UpdateCertificate(ctx context.Context, input *iot.UpdateCertificateInput) (*iot.UpdateCertificateOutput, error) {
    return a.client.UpdateCertificateWithContext(ctx, input)
}

func (a *IoTActivities) UpdateDimension(ctx context.Context, input *iot.UpdateDimensionInput) (*iot.UpdateDimensionOutput, error) {
    return a.client.UpdateDimensionWithContext(ctx, input)
}

func (a *IoTActivities) UpdateDomainConfiguration(ctx context.Context, input *iot.UpdateDomainConfigurationInput) (*iot.UpdateDomainConfigurationOutput, error) {
    return a.client.UpdateDomainConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) UpdateDynamicThingGroup(ctx context.Context, input *iot.UpdateDynamicThingGroupInput) (*iot.UpdateDynamicThingGroupOutput, error) {
    return a.client.UpdateDynamicThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) UpdateEventConfigurations(ctx context.Context, input *iot.UpdateEventConfigurationsInput) (*iot.UpdateEventConfigurationsOutput, error) {
    return a.client.UpdateEventConfigurationsWithContext(ctx, input)
}

func (a *IoTActivities) UpdateIndexingConfiguration(ctx context.Context, input *iot.UpdateIndexingConfigurationInput) (*iot.UpdateIndexingConfigurationOutput, error) {
    return a.client.UpdateIndexingConfigurationWithContext(ctx, input)
}

func (a *IoTActivities) UpdateJob(ctx context.Context, input *iot.UpdateJobInput) (*iot.UpdateJobOutput, error) {
    return a.client.UpdateJobWithContext(ctx, input)
}

func (a *IoTActivities) UpdateMitigationAction(ctx context.Context, input *iot.UpdateMitigationActionInput) (*iot.UpdateMitigationActionOutput, error) {
    return a.client.UpdateMitigationActionWithContext(ctx, input)
}

func (a *IoTActivities) UpdateProvisioningTemplate(ctx context.Context, input *iot.UpdateProvisioningTemplateInput) (*iot.UpdateProvisioningTemplateOutput, error) {
    return a.client.UpdateProvisioningTemplateWithContext(ctx, input)
}

func (a *IoTActivities) UpdateRoleAlias(ctx context.Context, input *iot.UpdateRoleAliasInput) (*iot.UpdateRoleAliasOutput, error) {
    return a.client.UpdateRoleAliasWithContext(ctx, input)
}

func (a *IoTActivities) UpdateScheduledAudit(ctx context.Context, input *iot.UpdateScheduledAuditInput) (*iot.UpdateScheduledAuditOutput, error) {
    return a.client.UpdateScheduledAuditWithContext(ctx, input)
}

func (a *IoTActivities) UpdateSecurityProfile(ctx context.Context, input *iot.UpdateSecurityProfileInput) (*iot.UpdateSecurityProfileOutput, error) {
    return a.client.UpdateSecurityProfileWithContext(ctx, input)
}

func (a *IoTActivities) UpdateStream(ctx context.Context, input *iot.UpdateStreamInput) (*iot.UpdateStreamOutput, error) {
    return a.client.UpdateStreamWithContext(ctx, input)
}

func (a *IoTActivities) UpdateThing(ctx context.Context, input *iot.UpdateThingInput) (*iot.UpdateThingOutput, error) {
    return a.client.UpdateThingWithContext(ctx, input)
}

func (a *IoTActivities) UpdateThingGroup(ctx context.Context, input *iot.UpdateThingGroupInput) (*iot.UpdateThingGroupOutput, error) {
    return a.client.UpdateThingGroupWithContext(ctx, input)
}

func (a *IoTActivities) UpdateThingGroupsForThing(ctx context.Context, input *iot.UpdateThingGroupsForThingInput) (*iot.UpdateThingGroupsForThingOutput, error) {
    return a.client.UpdateThingGroupsForThingWithContext(ctx, input)
}

func (a *IoTActivities) UpdateTopicRuleDestination(ctx context.Context, input *iot.UpdateTopicRuleDestinationInput) (*iot.UpdateTopicRuleDestinationOutput, error) {
    return a.client.UpdateTopicRuleDestinationWithContext(ctx, input)
}

func (a *IoTActivities) ValidateSecurityProfileBehaviors(ctx context.Context, input *iot.ValidateSecurityProfileBehaviorsInput) (*iot.ValidateSecurityProfileBehaviorsOutput, error) {
    return a.client.ValidateSecurityProfileBehaviorsWithContext(ctx, input)
}
