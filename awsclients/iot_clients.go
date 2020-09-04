package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iot"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IoTClient interface {
    AcceptCertificateTransfer(ctx workflow.Context, input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error)
    AcceptCertificateTransferAsync(ctx workflow.Context, input *iot.AcceptCertificateTransferInput) *IotAcceptCertificateTransferResult

    AddThingToBillingGroup(ctx workflow.Context, input *iot.AddThingToBillingGroupInput) (*iot.AddThingToBillingGroupOutput, error)
    AddThingToBillingGroupAsync(ctx workflow.Context, input *iot.AddThingToBillingGroupInput) *IotAddThingToBillingGroupResult

    AddThingToThingGroup(ctx workflow.Context, input *iot.AddThingToThingGroupInput) (*iot.AddThingToThingGroupOutput, error)
    AddThingToThingGroupAsync(ctx workflow.Context, input *iot.AddThingToThingGroupInput) *IotAddThingToThingGroupResult

    AssociateTargetsWithJob(ctx workflow.Context, input *iot.AssociateTargetsWithJobInput) (*iot.AssociateTargetsWithJobOutput, error)
    AssociateTargetsWithJobAsync(ctx workflow.Context, input *iot.AssociateTargetsWithJobInput) *IotAssociateTargetsWithJobResult

    AttachPolicy(ctx workflow.Context, input *iot.AttachPolicyInput) (*iot.AttachPolicyOutput, error)
    AttachPolicyAsync(ctx workflow.Context, input *iot.AttachPolicyInput) *IotAttachPolicyResult

    AttachPrincipalPolicy(ctx workflow.Context, input *iot.AttachPrincipalPolicyInput) (*iot.AttachPrincipalPolicyOutput, error)
    AttachPrincipalPolicyAsync(ctx workflow.Context, input *iot.AttachPrincipalPolicyInput) *IotAttachPrincipalPolicyResult

    AttachSecurityProfile(ctx workflow.Context, input *iot.AttachSecurityProfileInput) (*iot.AttachSecurityProfileOutput, error)
    AttachSecurityProfileAsync(ctx workflow.Context, input *iot.AttachSecurityProfileInput) *IotAttachSecurityProfileResult

    AttachThingPrincipal(ctx workflow.Context, input *iot.AttachThingPrincipalInput) (*iot.AttachThingPrincipalOutput, error)
    AttachThingPrincipalAsync(ctx workflow.Context, input *iot.AttachThingPrincipalInput) *IotAttachThingPrincipalResult

    CancelAuditMitigationActionsTask(ctx workflow.Context, input *iot.CancelAuditMitigationActionsTaskInput) (*iot.CancelAuditMitigationActionsTaskOutput, error)
    CancelAuditMitigationActionsTaskAsync(ctx workflow.Context, input *iot.CancelAuditMitigationActionsTaskInput) *IotCancelAuditMitigationActionsTaskResult

    CancelAuditTask(ctx workflow.Context, input *iot.CancelAuditTaskInput) (*iot.CancelAuditTaskOutput, error)
    CancelAuditTaskAsync(ctx workflow.Context, input *iot.CancelAuditTaskInput) *IotCancelAuditTaskResult

    CancelCertificateTransfer(ctx workflow.Context, input *iot.CancelCertificateTransferInput) (*iot.CancelCertificateTransferOutput, error)
    CancelCertificateTransferAsync(ctx workflow.Context, input *iot.CancelCertificateTransferInput) *IotCancelCertificateTransferResult

    CancelJob(ctx workflow.Context, input *iot.CancelJobInput) (*iot.CancelJobOutput, error)
    CancelJobAsync(ctx workflow.Context, input *iot.CancelJobInput) *IotCancelJobResult

    CancelJobExecution(ctx workflow.Context, input *iot.CancelJobExecutionInput) (*iot.CancelJobExecutionOutput, error)
    CancelJobExecutionAsync(ctx workflow.Context, input *iot.CancelJobExecutionInput) *IotCancelJobExecutionResult

    ClearDefaultAuthorizer(ctx workflow.Context, input *iot.ClearDefaultAuthorizerInput) (*iot.ClearDefaultAuthorizerOutput, error)
    ClearDefaultAuthorizerAsync(ctx workflow.Context, input *iot.ClearDefaultAuthorizerInput) *IotClearDefaultAuthorizerResult

    ConfirmTopicRuleDestination(ctx workflow.Context, input *iot.ConfirmTopicRuleDestinationInput) (*iot.ConfirmTopicRuleDestinationOutput, error)
    ConfirmTopicRuleDestinationAsync(ctx workflow.Context, input *iot.ConfirmTopicRuleDestinationInput) *IotConfirmTopicRuleDestinationResult

    CreateAuditSuppression(ctx workflow.Context, input *iot.CreateAuditSuppressionInput) (*iot.CreateAuditSuppressionOutput, error)
    CreateAuditSuppressionAsync(ctx workflow.Context, input *iot.CreateAuditSuppressionInput) *IotCreateAuditSuppressionResult

    CreateAuthorizer(ctx workflow.Context, input *iot.CreateAuthorizerInput) (*iot.CreateAuthorizerOutput, error)
    CreateAuthorizerAsync(ctx workflow.Context, input *iot.CreateAuthorizerInput) *IotCreateAuthorizerResult

    CreateBillingGroup(ctx workflow.Context, input *iot.CreateBillingGroupInput) (*iot.CreateBillingGroupOutput, error)
    CreateBillingGroupAsync(ctx workflow.Context, input *iot.CreateBillingGroupInput) *IotCreateBillingGroupResult

    CreateCertificateFromCsr(ctx workflow.Context, input *iot.CreateCertificateFromCsrInput) (*iot.CreateCertificateFromCsrOutput, error)
    CreateCertificateFromCsrAsync(ctx workflow.Context, input *iot.CreateCertificateFromCsrInput) *IotCreateCertificateFromCsrResult

    CreateDimension(ctx workflow.Context, input *iot.CreateDimensionInput) (*iot.CreateDimensionOutput, error)
    CreateDimensionAsync(ctx workflow.Context, input *iot.CreateDimensionInput) *IotCreateDimensionResult

    CreateDomainConfiguration(ctx workflow.Context, input *iot.CreateDomainConfigurationInput) (*iot.CreateDomainConfigurationOutput, error)
    CreateDomainConfigurationAsync(ctx workflow.Context, input *iot.CreateDomainConfigurationInput) *IotCreateDomainConfigurationResult

    CreateDynamicThingGroup(ctx workflow.Context, input *iot.CreateDynamicThingGroupInput) (*iot.CreateDynamicThingGroupOutput, error)
    CreateDynamicThingGroupAsync(ctx workflow.Context, input *iot.CreateDynamicThingGroupInput) *IotCreateDynamicThingGroupResult

    CreateJob(ctx workflow.Context, input *iot.CreateJobInput) (*iot.CreateJobOutput, error)
    CreateJobAsync(ctx workflow.Context, input *iot.CreateJobInput) *IotCreateJobResult

    CreateKeysAndCertificate(ctx workflow.Context, input *iot.CreateKeysAndCertificateInput) (*iot.CreateKeysAndCertificateOutput, error)
    CreateKeysAndCertificateAsync(ctx workflow.Context, input *iot.CreateKeysAndCertificateInput) *IotCreateKeysAndCertificateResult

    CreateMitigationAction(ctx workflow.Context, input *iot.CreateMitigationActionInput) (*iot.CreateMitigationActionOutput, error)
    CreateMitigationActionAsync(ctx workflow.Context, input *iot.CreateMitigationActionInput) *IotCreateMitigationActionResult

    CreateOTAUpdate(ctx workflow.Context, input *iot.CreateOTAUpdateInput) (*iot.CreateOTAUpdateOutput, error)
    CreateOTAUpdateAsync(ctx workflow.Context, input *iot.CreateOTAUpdateInput) *IotCreateOTAUpdateResult

    CreatePolicy(ctx workflow.Context, input *iot.CreatePolicyInput) (*iot.CreatePolicyOutput, error)
    CreatePolicyAsync(ctx workflow.Context, input *iot.CreatePolicyInput) *IotCreatePolicyResult

    CreatePolicyVersion(ctx workflow.Context, input *iot.CreatePolicyVersionInput) (*iot.CreatePolicyVersionOutput, error)
    CreatePolicyVersionAsync(ctx workflow.Context, input *iot.CreatePolicyVersionInput) *IotCreatePolicyVersionResult

    CreateProvisioningClaim(ctx workflow.Context, input *iot.CreateProvisioningClaimInput) (*iot.CreateProvisioningClaimOutput, error)
    CreateProvisioningClaimAsync(ctx workflow.Context, input *iot.CreateProvisioningClaimInput) *IotCreateProvisioningClaimResult

    CreateProvisioningTemplate(ctx workflow.Context, input *iot.CreateProvisioningTemplateInput) (*iot.CreateProvisioningTemplateOutput, error)
    CreateProvisioningTemplateAsync(ctx workflow.Context, input *iot.CreateProvisioningTemplateInput) *IotCreateProvisioningTemplateResult

    CreateProvisioningTemplateVersion(ctx workflow.Context, input *iot.CreateProvisioningTemplateVersionInput) (*iot.CreateProvisioningTemplateVersionOutput, error)
    CreateProvisioningTemplateVersionAsync(ctx workflow.Context, input *iot.CreateProvisioningTemplateVersionInput) *IotCreateProvisioningTemplateVersionResult

    CreateRoleAlias(ctx workflow.Context, input *iot.CreateRoleAliasInput) (*iot.CreateRoleAliasOutput, error)
    CreateRoleAliasAsync(ctx workflow.Context, input *iot.CreateRoleAliasInput) *IotCreateRoleAliasResult

    CreateScheduledAudit(ctx workflow.Context, input *iot.CreateScheduledAuditInput) (*iot.CreateScheduledAuditOutput, error)
    CreateScheduledAuditAsync(ctx workflow.Context, input *iot.CreateScheduledAuditInput) *IotCreateScheduledAuditResult

    CreateSecurityProfile(ctx workflow.Context, input *iot.CreateSecurityProfileInput) (*iot.CreateSecurityProfileOutput, error)
    CreateSecurityProfileAsync(ctx workflow.Context, input *iot.CreateSecurityProfileInput) *IotCreateSecurityProfileResult

    CreateStream(ctx workflow.Context, input *iot.CreateStreamInput) (*iot.CreateStreamOutput, error)
    CreateStreamAsync(ctx workflow.Context, input *iot.CreateStreamInput) *IotCreateStreamResult

    CreateThing(ctx workflow.Context, input *iot.CreateThingInput) (*iot.CreateThingOutput, error)
    CreateThingAsync(ctx workflow.Context, input *iot.CreateThingInput) *IotCreateThingResult

    CreateThingGroup(ctx workflow.Context, input *iot.CreateThingGroupInput) (*iot.CreateThingGroupOutput, error)
    CreateThingGroupAsync(ctx workflow.Context, input *iot.CreateThingGroupInput) *IotCreateThingGroupResult

    CreateThingType(ctx workflow.Context, input *iot.CreateThingTypeInput) (*iot.CreateThingTypeOutput, error)
    CreateThingTypeAsync(ctx workflow.Context, input *iot.CreateThingTypeInput) *IotCreateThingTypeResult

    CreateTopicRule(ctx workflow.Context, input *iot.CreateTopicRuleInput) (*iot.CreateTopicRuleOutput, error)
    CreateTopicRuleAsync(ctx workflow.Context, input *iot.CreateTopicRuleInput) *IotCreateTopicRuleResult

    CreateTopicRuleDestination(ctx workflow.Context, input *iot.CreateTopicRuleDestinationInput) (*iot.CreateTopicRuleDestinationOutput, error)
    CreateTopicRuleDestinationAsync(ctx workflow.Context, input *iot.CreateTopicRuleDestinationInput) *IotCreateTopicRuleDestinationResult

    DeleteAccountAuditConfiguration(ctx workflow.Context, input *iot.DeleteAccountAuditConfigurationInput) (*iot.DeleteAccountAuditConfigurationOutput, error)
    DeleteAccountAuditConfigurationAsync(ctx workflow.Context, input *iot.DeleteAccountAuditConfigurationInput) *IotDeleteAccountAuditConfigurationResult

    DeleteAuditSuppression(ctx workflow.Context, input *iot.DeleteAuditSuppressionInput) (*iot.DeleteAuditSuppressionOutput, error)
    DeleteAuditSuppressionAsync(ctx workflow.Context, input *iot.DeleteAuditSuppressionInput) *IotDeleteAuditSuppressionResult

    DeleteAuthorizer(ctx workflow.Context, input *iot.DeleteAuthorizerInput) (*iot.DeleteAuthorizerOutput, error)
    DeleteAuthorizerAsync(ctx workflow.Context, input *iot.DeleteAuthorizerInput) *IotDeleteAuthorizerResult

    DeleteBillingGroup(ctx workflow.Context, input *iot.DeleteBillingGroupInput) (*iot.DeleteBillingGroupOutput, error)
    DeleteBillingGroupAsync(ctx workflow.Context, input *iot.DeleteBillingGroupInput) *IotDeleteBillingGroupResult

    DeleteCACertificate(ctx workflow.Context, input *iot.DeleteCACertificateInput) (*iot.DeleteCACertificateOutput, error)
    DeleteCACertificateAsync(ctx workflow.Context, input *iot.DeleteCACertificateInput) *IotDeleteCACertificateResult

    DeleteCertificate(ctx workflow.Context, input *iot.DeleteCertificateInput) (*iot.DeleteCertificateOutput, error)
    DeleteCertificateAsync(ctx workflow.Context, input *iot.DeleteCertificateInput) *IotDeleteCertificateResult

    DeleteDimension(ctx workflow.Context, input *iot.DeleteDimensionInput) (*iot.DeleteDimensionOutput, error)
    DeleteDimensionAsync(ctx workflow.Context, input *iot.DeleteDimensionInput) *IotDeleteDimensionResult

    DeleteDomainConfiguration(ctx workflow.Context, input *iot.DeleteDomainConfigurationInput) (*iot.DeleteDomainConfigurationOutput, error)
    DeleteDomainConfigurationAsync(ctx workflow.Context, input *iot.DeleteDomainConfigurationInput) *IotDeleteDomainConfigurationResult

    DeleteDynamicThingGroup(ctx workflow.Context, input *iot.DeleteDynamicThingGroupInput) (*iot.DeleteDynamicThingGroupOutput, error)
    DeleteDynamicThingGroupAsync(ctx workflow.Context, input *iot.DeleteDynamicThingGroupInput) *IotDeleteDynamicThingGroupResult

    DeleteJob(ctx workflow.Context, input *iot.DeleteJobInput) (*iot.DeleteJobOutput, error)
    DeleteJobAsync(ctx workflow.Context, input *iot.DeleteJobInput) *IotDeleteJobResult

    DeleteJobExecution(ctx workflow.Context, input *iot.DeleteJobExecutionInput) (*iot.DeleteJobExecutionOutput, error)
    DeleteJobExecutionAsync(ctx workflow.Context, input *iot.DeleteJobExecutionInput) *IotDeleteJobExecutionResult

    DeleteMitigationAction(ctx workflow.Context, input *iot.DeleteMitigationActionInput) (*iot.DeleteMitigationActionOutput, error)
    DeleteMitigationActionAsync(ctx workflow.Context, input *iot.DeleteMitigationActionInput) *IotDeleteMitigationActionResult

    DeleteOTAUpdate(ctx workflow.Context, input *iot.DeleteOTAUpdateInput) (*iot.DeleteOTAUpdateOutput, error)
    DeleteOTAUpdateAsync(ctx workflow.Context, input *iot.DeleteOTAUpdateInput) *IotDeleteOTAUpdateResult

    DeletePolicy(ctx workflow.Context, input *iot.DeletePolicyInput) (*iot.DeletePolicyOutput, error)
    DeletePolicyAsync(ctx workflow.Context, input *iot.DeletePolicyInput) *IotDeletePolicyResult

    DeletePolicyVersion(ctx workflow.Context, input *iot.DeletePolicyVersionInput) (*iot.DeletePolicyVersionOutput, error)
    DeletePolicyVersionAsync(ctx workflow.Context, input *iot.DeletePolicyVersionInput) *IotDeletePolicyVersionResult

    DeleteProvisioningTemplate(ctx workflow.Context, input *iot.DeleteProvisioningTemplateInput) (*iot.DeleteProvisioningTemplateOutput, error)
    DeleteProvisioningTemplateAsync(ctx workflow.Context, input *iot.DeleteProvisioningTemplateInput) *IotDeleteProvisioningTemplateResult

    DeleteProvisioningTemplateVersion(ctx workflow.Context, input *iot.DeleteProvisioningTemplateVersionInput) (*iot.DeleteProvisioningTemplateVersionOutput, error)
    DeleteProvisioningTemplateVersionAsync(ctx workflow.Context, input *iot.DeleteProvisioningTemplateVersionInput) *IotDeleteProvisioningTemplateVersionResult

    DeleteRegistrationCode(ctx workflow.Context, input *iot.DeleteRegistrationCodeInput) (*iot.DeleteRegistrationCodeOutput, error)
    DeleteRegistrationCodeAsync(ctx workflow.Context, input *iot.DeleteRegistrationCodeInput) *IotDeleteRegistrationCodeResult

    DeleteRoleAlias(ctx workflow.Context, input *iot.DeleteRoleAliasInput) (*iot.DeleteRoleAliasOutput, error)
    DeleteRoleAliasAsync(ctx workflow.Context, input *iot.DeleteRoleAliasInput) *IotDeleteRoleAliasResult

    DeleteScheduledAudit(ctx workflow.Context, input *iot.DeleteScheduledAuditInput) (*iot.DeleteScheduledAuditOutput, error)
    DeleteScheduledAuditAsync(ctx workflow.Context, input *iot.DeleteScheduledAuditInput) *IotDeleteScheduledAuditResult

    DeleteSecurityProfile(ctx workflow.Context, input *iot.DeleteSecurityProfileInput) (*iot.DeleteSecurityProfileOutput, error)
    DeleteSecurityProfileAsync(ctx workflow.Context, input *iot.DeleteSecurityProfileInput) *IotDeleteSecurityProfileResult

    DeleteStream(ctx workflow.Context, input *iot.DeleteStreamInput) (*iot.DeleteStreamOutput, error)
    DeleteStreamAsync(ctx workflow.Context, input *iot.DeleteStreamInput) *IotDeleteStreamResult

    DeleteThing(ctx workflow.Context, input *iot.DeleteThingInput) (*iot.DeleteThingOutput, error)
    DeleteThingAsync(ctx workflow.Context, input *iot.DeleteThingInput) *IotDeleteThingResult

    DeleteThingGroup(ctx workflow.Context, input *iot.DeleteThingGroupInput) (*iot.DeleteThingGroupOutput, error)
    DeleteThingGroupAsync(ctx workflow.Context, input *iot.DeleteThingGroupInput) *IotDeleteThingGroupResult

    DeleteThingType(ctx workflow.Context, input *iot.DeleteThingTypeInput) (*iot.DeleteThingTypeOutput, error)
    DeleteThingTypeAsync(ctx workflow.Context, input *iot.DeleteThingTypeInput) *IotDeleteThingTypeResult

    DeleteTopicRule(ctx workflow.Context, input *iot.DeleteTopicRuleInput) (*iot.DeleteTopicRuleOutput, error)
    DeleteTopicRuleAsync(ctx workflow.Context, input *iot.DeleteTopicRuleInput) *IotDeleteTopicRuleResult

    DeleteTopicRuleDestination(ctx workflow.Context, input *iot.DeleteTopicRuleDestinationInput) (*iot.DeleteTopicRuleDestinationOutput, error)
    DeleteTopicRuleDestinationAsync(ctx workflow.Context, input *iot.DeleteTopicRuleDestinationInput) *IotDeleteTopicRuleDestinationResult

    DeleteV2LoggingLevel(ctx workflow.Context, input *iot.DeleteV2LoggingLevelInput) (*iot.DeleteV2LoggingLevelOutput, error)
    DeleteV2LoggingLevelAsync(ctx workflow.Context, input *iot.DeleteV2LoggingLevelInput) *IotDeleteV2LoggingLevelResult

    DeprecateThingType(ctx workflow.Context, input *iot.DeprecateThingTypeInput) (*iot.DeprecateThingTypeOutput, error)
    DeprecateThingTypeAsync(ctx workflow.Context, input *iot.DeprecateThingTypeInput) *IotDeprecateThingTypeResult

    DescribeAccountAuditConfiguration(ctx workflow.Context, input *iot.DescribeAccountAuditConfigurationInput) (*iot.DescribeAccountAuditConfigurationOutput, error)
    DescribeAccountAuditConfigurationAsync(ctx workflow.Context, input *iot.DescribeAccountAuditConfigurationInput) *IotDescribeAccountAuditConfigurationResult

    DescribeAuditFinding(ctx workflow.Context, input *iot.DescribeAuditFindingInput) (*iot.DescribeAuditFindingOutput, error)
    DescribeAuditFindingAsync(ctx workflow.Context, input *iot.DescribeAuditFindingInput) *IotDescribeAuditFindingResult

    DescribeAuditMitigationActionsTask(ctx workflow.Context, input *iot.DescribeAuditMitigationActionsTaskInput) (*iot.DescribeAuditMitigationActionsTaskOutput, error)
    DescribeAuditMitigationActionsTaskAsync(ctx workflow.Context, input *iot.DescribeAuditMitigationActionsTaskInput) *IotDescribeAuditMitigationActionsTaskResult

    DescribeAuditSuppression(ctx workflow.Context, input *iot.DescribeAuditSuppressionInput) (*iot.DescribeAuditSuppressionOutput, error)
    DescribeAuditSuppressionAsync(ctx workflow.Context, input *iot.DescribeAuditSuppressionInput) *IotDescribeAuditSuppressionResult

    DescribeAuditTask(ctx workflow.Context, input *iot.DescribeAuditTaskInput) (*iot.DescribeAuditTaskOutput, error)
    DescribeAuditTaskAsync(ctx workflow.Context, input *iot.DescribeAuditTaskInput) *IotDescribeAuditTaskResult

    DescribeAuthorizer(ctx workflow.Context, input *iot.DescribeAuthorizerInput) (*iot.DescribeAuthorizerOutput, error)
    DescribeAuthorizerAsync(ctx workflow.Context, input *iot.DescribeAuthorizerInput) *IotDescribeAuthorizerResult

    DescribeBillingGroup(ctx workflow.Context, input *iot.DescribeBillingGroupInput) (*iot.DescribeBillingGroupOutput, error)
    DescribeBillingGroupAsync(ctx workflow.Context, input *iot.DescribeBillingGroupInput) *IotDescribeBillingGroupResult

    DescribeCACertificate(ctx workflow.Context, input *iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error)
    DescribeCACertificateAsync(ctx workflow.Context, input *iot.DescribeCACertificateInput) *IotDescribeCACertificateResult

    DescribeCertificate(ctx workflow.Context, input *iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error)
    DescribeCertificateAsync(ctx workflow.Context, input *iot.DescribeCertificateInput) *IotDescribeCertificateResult

    DescribeDefaultAuthorizer(ctx workflow.Context, input *iot.DescribeDefaultAuthorizerInput) (*iot.DescribeDefaultAuthorizerOutput, error)
    DescribeDefaultAuthorizerAsync(ctx workflow.Context, input *iot.DescribeDefaultAuthorizerInput) *IotDescribeDefaultAuthorizerResult

    DescribeDimension(ctx workflow.Context, input *iot.DescribeDimensionInput) (*iot.DescribeDimensionOutput, error)
    DescribeDimensionAsync(ctx workflow.Context, input *iot.DescribeDimensionInput) *IotDescribeDimensionResult

    DescribeDomainConfiguration(ctx workflow.Context, input *iot.DescribeDomainConfigurationInput) (*iot.DescribeDomainConfigurationOutput, error)
    DescribeDomainConfigurationAsync(ctx workflow.Context, input *iot.DescribeDomainConfigurationInput) *IotDescribeDomainConfigurationResult

    DescribeEndpoint(ctx workflow.Context, input *iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error)
    DescribeEndpointAsync(ctx workflow.Context, input *iot.DescribeEndpointInput) *IotDescribeEndpointResult

    DescribeEventConfigurations(ctx workflow.Context, input *iot.DescribeEventConfigurationsInput) (*iot.DescribeEventConfigurationsOutput, error)
    DescribeEventConfigurationsAsync(ctx workflow.Context, input *iot.DescribeEventConfigurationsInput) *IotDescribeEventConfigurationsResult

    DescribeIndex(ctx workflow.Context, input *iot.DescribeIndexInput) (*iot.DescribeIndexOutput, error)
    DescribeIndexAsync(ctx workflow.Context, input *iot.DescribeIndexInput) *IotDescribeIndexResult

    DescribeJob(ctx workflow.Context, input *iot.DescribeJobInput) (*iot.DescribeJobOutput, error)
    DescribeJobAsync(ctx workflow.Context, input *iot.DescribeJobInput) *IotDescribeJobResult

    DescribeJobExecution(ctx workflow.Context, input *iot.DescribeJobExecutionInput) (*iot.DescribeJobExecutionOutput, error)
    DescribeJobExecutionAsync(ctx workflow.Context, input *iot.DescribeJobExecutionInput) *IotDescribeJobExecutionResult

    DescribeMitigationAction(ctx workflow.Context, input *iot.DescribeMitigationActionInput) (*iot.DescribeMitigationActionOutput, error)
    DescribeMitigationActionAsync(ctx workflow.Context, input *iot.DescribeMitigationActionInput) *IotDescribeMitigationActionResult

    DescribeProvisioningTemplate(ctx workflow.Context, input *iot.DescribeProvisioningTemplateInput) (*iot.DescribeProvisioningTemplateOutput, error)
    DescribeProvisioningTemplateAsync(ctx workflow.Context, input *iot.DescribeProvisioningTemplateInput) *IotDescribeProvisioningTemplateResult

    DescribeProvisioningTemplateVersion(ctx workflow.Context, input *iot.DescribeProvisioningTemplateVersionInput) (*iot.DescribeProvisioningTemplateVersionOutput, error)
    DescribeProvisioningTemplateVersionAsync(ctx workflow.Context, input *iot.DescribeProvisioningTemplateVersionInput) *IotDescribeProvisioningTemplateVersionResult

    DescribeRoleAlias(ctx workflow.Context, input *iot.DescribeRoleAliasInput) (*iot.DescribeRoleAliasOutput, error)
    DescribeRoleAliasAsync(ctx workflow.Context, input *iot.DescribeRoleAliasInput) *IotDescribeRoleAliasResult

    DescribeScheduledAudit(ctx workflow.Context, input *iot.DescribeScheduledAuditInput) (*iot.DescribeScheduledAuditOutput, error)
    DescribeScheduledAuditAsync(ctx workflow.Context, input *iot.DescribeScheduledAuditInput) *IotDescribeScheduledAuditResult

    DescribeSecurityProfile(ctx workflow.Context, input *iot.DescribeSecurityProfileInput) (*iot.DescribeSecurityProfileOutput, error)
    DescribeSecurityProfileAsync(ctx workflow.Context, input *iot.DescribeSecurityProfileInput) *IotDescribeSecurityProfileResult

    DescribeStream(ctx workflow.Context, input *iot.DescribeStreamInput) (*iot.DescribeStreamOutput, error)
    DescribeStreamAsync(ctx workflow.Context, input *iot.DescribeStreamInput) *IotDescribeStreamResult

    DescribeThing(ctx workflow.Context, input *iot.DescribeThingInput) (*iot.DescribeThingOutput, error)
    DescribeThingAsync(ctx workflow.Context, input *iot.DescribeThingInput) *IotDescribeThingResult

    DescribeThingGroup(ctx workflow.Context, input *iot.DescribeThingGroupInput) (*iot.DescribeThingGroupOutput, error)
    DescribeThingGroupAsync(ctx workflow.Context, input *iot.DescribeThingGroupInput) *IotDescribeThingGroupResult

    DescribeThingRegistrationTask(ctx workflow.Context, input *iot.DescribeThingRegistrationTaskInput) (*iot.DescribeThingRegistrationTaskOutput, error)
    DescribeThingRegistrationTaskAsync(ctx workflow.Context, input *iot.DescribeThingRegistrationTaskInput) *IotDescribeThingRegistrationTaskResult

    DescribeThingType(ctx workflow.Context, input *iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error)
    DescribeThingTypeAsync(ctx workflow.Context, input *iot.DescribeThingTypeInput) *IotDescribeThingTypeResult

    DetachPolicy(ctx workflow.Context, input *iot.DetachPolicyInput) (*iot.DetachPolicyOutput, error)
    DetachPolicyAsync(ctx workflow.Context, input *iot.DetachPolicyInput) *IotDetachPolicyResult

    DetachPrincipalPolicy(ctx workflow.Context, input *iot.DetachPrincipalPolicyInput) (*iot.DetachPrincipalPolicyOutput, error)
    DetachPrincipalPolicyAsync(ctx workflow.Context, input *iot.DetachPrincipalPolicyInput) *IotDetachPrincipalPolicyResult

    DetachSecurityProfile(ctx workflow.Context, input *iot.DetachSecurityProfileInput) (*iot.DetachSecurityProfileOutput, error)
    DetachSecurityProfileAsync(ctx workflow.Context, input *iot.DetachSecurityProfileInput) *IotDetachSecurityProfileResult

    DetachThingPrincipal(ctx workflow.Context, input *iot.DetachThingPrincipalInput) (*iot.DetachThingPrincipalOutput, error)
    DetachThingPrincipalAsync(ctx workflow.Context, input *iot.DetachThingPrincipalInput) *IotDetachThingPrincipalResult

    DisableTopicRule(ctx workflow.Context, input *iot.DisableTopicRuleInput) (*iot.DisableTopicRuleOutput, error)
    DisableTopicRuleAsync(ctx workflow.Context, input *iot.DisableTopicRuleInput) *IotDisableTopicRuleResult

    EnableTopicRule(ctx workflow.Context, input *iot.EnableTopicRuleInput) (*iot.EnableTopicRuleOutput, error)
    EnableTopicRuleAsync(ctx workflow.Context, input *iot.EnableTopicRuleInput) *IotEnableTopicRuleResult

    GetCardinality(ctx workflow.Context, input *iot.GetCardinalityInput) (*iot.GetCardinalityOutput, error)
    GetCardinalityAsync(ctx workflow.Context, input *iot.GetCardinalityInput) *IotGetCardinalityResult

    GetEffectivePolicies(ctx workflow.Context, input *iot.GetEffectivePoliciesInput) (*iot.GetEffectivePoliciesOutput, error)
    GetEffectivePoliciesAsync(ctx workflow.Context, input *iot.GetEffectivePoliciesInput) *IotGetEffectivePoliciesResult

    GetIndexingConfiguration(ctx workflow.Context, input *iot.GetIndexingConfigurationInput) (*iot.GetIndexingConfigurationOutput, error)
    GetIndexingConfigurationAsync(ctx workflow.Context, input *iot.GetIndexingConfigurationInput) *IotGetIndexingConfigurationResult

    GetJobDocument(ctx workflow.Context, input *iot.GetJobDocumentInput) (*iot.GetJobDocumentOutput, error)
    GetJobDocumentAsync(ctx workflow.Context, input *iot.GetJobDocumentInput) *IotGetJobDocumentResult

    GetLoggingOptions(ctx workflow.Context, input *iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error)
    GetLoggingOptionsAsync(ctx workflow.Context, input *iot.GetLoggingOptionsInput) *IotGetLoggingOptionsResult

    GetOTAUpdate(ctx workflow.Context, input *iot.GetOTAUpdateInput) (*iot.GetOTAUpdateOutput, error)
    GetOTAUpdateAsync(ctx workflow.Context, input *iot.GetOTAUpdateInput) *IotGetOTAUpdateResult

    GetPercentiles(ctx workflow.Context, input *iot.GetPercentilesInput) (*iot.GetPercentilesOutput, error)
    GetPercentilesAsync(ctx workflow.Context, input *iot.GetPercentilesInput) *IotGetPercentilesResult

    GetPolicy(ctx workflow.Context, input *iot.GetPolicyInput) (*iot.GetPolicyOutput, error)
    GetPolicyAsync(ctx workflow.Context, input *iot.GetPolicyInput) *IotGetPolicyResult

    GetPolicyVersion(ctx workflow.Context, input *iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error)
    GetPolicyVersionAsync(ctx workflow.Context, input *iot.GetPolicyVersionInput) *IotGetPolicyVersionResult

    GetRegistrationCode(ctx workflow.Context, input *iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error)
    GetRegistrationCodeAsync(ctx workflow.Context, input *iot.GetRegistrationCodeInput) *IotGetRegistrationCodeResult

    GetStatistics(ctx workflow.Context, input *iot.GetStatisticsInput) (*iot.GetStatisticsOutput, error)
    GetStatisticsAsync(ctx workflow.Context, input *iot.GetStatisticsInput) *IotGetStatisticsResult

    GetTopicRule(ctx workflow.Context, input *iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error)
    GetTopicRuleAsync(ctx workflow.Context, input *iot.GetTopicRuleInput) *IotGetTopicRuleResult

    GetTopicRuleDestination(ctx workflow.Context, input *iot.GetTopicRuleDestinationInput) (*iot.GetTopicRuleDestinationOutput, error)
    GetTopicRuleDestinationAsync(ctx workflow.Context, input *iot.GetTopicRuleDestinationInput) *IotGetTopicRuleDestinationResult

    GetV2LoggingOptions(ctx workflow.Context, input *iot.GetV2LoggingOptionsInput) (*iot.GetV2LoggingOptionsOutput, error)
    GetV2LoggingOptionsAsync(ctx workflow.Context, input *iot.GetV2LoggingOptionsInput) *IotGetV2LoggingOptionsResult

    ListActiveViolations(ctx workflow.Context, input *iot.ListActiveViolationsInput) (*iot.ListActiveViolationsOutput, error)
    ListActiveViolationsAsync(ctx workflow.Context, input *iot.ListActiveViolationsInput) *IotListActiveViolationsResult

    ListAttachedPolicies(ctx workflow.Context, input *iot.ListAttachedPoliciesInput) (*iot.ListAttachedPoliciesOutput, error)
    ListAttachedPoliciesAsync(ctx workflow.Context, input *iot.ListAttachedPoliciesInput) *IotListAttachedPoliciesResult

    ListAuditFindings(ctx workflow.Context, input *iot.ListAuditFindingsInput) (*iot.ListAuditFindingsOutput, error)
    ListAuditFindingsAsync(ctx workflow.Context, input *iot.ListAuditFindingsInput) *IotListAuditFindingsResult

    ListAuditMitigationActionsExecutions(ctx workflow.Context, input *iot.ListAuditMitigationActionsExecutionsInput) (*iot.ListAuditMitigationActionsExecutionsOutput, error)
    ListAuditMitigationActionsExecutionsAsync(ctx workflow.Context, input *iot.ListAuditMitigationActionsExecutionsInput) *IotListAuditMitigationActionsExecutionsResult

    ListAuditMitigationActionsTasks(ctx workflow.Context, input *iot.ListAuditMitigationActionsTasksInput) (*iot.ListAuditMitigationActionsTasksOutput, error)
    ListAuditMitigationActionsTasksAsync(ctx workflow.Context, input *iot.ListAuditMitigationActionsTasksInput) *IotListAuditMitigationActionsTasksResult

    ListAuditSuppressions(ctx workflow.Context, input *iot.ListAuditSuppressionsInput) (*iot.ListAuditSuppressionsOutput, error)
    ListAuditSuppressionsAsync(ctx workflow.Context, input *iot.ListAuditSuppressionsInput) *IotListAuditSuppressionsResult

    ListAuditTasks(ctx workflow.Context, input *iot.ListAuditTasksInput) (*iot.ListAuditTasksOutput, error)
    ListAuditTasksAsync(ctx workflow.Context, input *iot.ListAuditTasksInput) *IotListAuditTasksResult

    ListAuthorizers(ctx workflow.Context, input *iot.ListAuthorizersInput) (*iot.ListAuthorizersOutput, error)
    ListAuthorizersAsync(ctx workflow.Context, input *iot.ListAuthorizersInput) *IotListAuthorizersResult

    ListBillingGroups(ctx workflow.Context, input *iot.ListBillingGroupsInput) (*iot.ListBillingGroupsOutput, error)
    ListBillingGroupsAsync(ctx workflow.Context, input *iot.ListBillingGroupsInput) *IotListBillingGroupsResult

    ListCACertificates(ctx workflow.Context, input *iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error)
    ListCACertificatesAsync(ctx workflow.Context, input *iot.ListCACertificatesInput) *IotListCACertificatesResult

    ListCertificates(ctx workflow.Context, input *iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error)
    ListCertificatesAsync(ctx workflow.Context, input *iot.ListCertificatesInput) *IotListCertificatesResult

    ListCertificatesByCA(ctx workflow.Context, input *iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error)
    ListCertificatesByCAAsync(ctx workflow.Context, input *iot.ListCertificatesByCAInput) *IotListCertificatesByCAResult

    ListDimensions(ctx workflow.Context, input *iot.ListDimensionsInput) (*iot.ListDimensionsOutput, error)
    ListDimensionsAsync(ctx workflow.Context, input *iot.ListDimensionsInput) *IotListDimensionsResult

    ListDomainConfigurations(ctx workflow.Context, input *iot.ListDomainConfigurationsInput) (*iot.ListDomainConfigurationsOutput, error)
    ListDomainConfigurationsAsync(ctx workflow.Context, input *iot.ListDomainConfigurationsInput) *IotListDomainConfigurationsResult

    ListIndices(ctx workflow.Context, input *iot.ListIndicesInput) (*iot.ListIndicesOutput, error)
    ListIndicesAsync(ctx workflow.Context, input *iot.ListIndicesInput) *IotListIndicesResult

    ListJobExecutionsForJob(ctx workflow.Context, input *iot.ListJobExecutionsForJobInput) (*iot.ListJobExecutionsForJobOutput, error)
    ListJobExecutionsForJobAsync(ctx workflow.Context, input *iot.ListJobExecutionsForJobInput) *IotListJobExecutionsForJobResult

    ListJobExecutionsForThing(ctx workflow.Context, input *iot.ListJobExecutionsForThingInput) (*iot.ListJobExecutionsForThingOutput, error)
    ListJobExecutionsForThingAsync(ctx workflow.Context, input *iot.ListJobExecutionsForThingInput) *IotListJobExecutionsForThingResult

    ListJobs(ctx workflow.Context, input *iot.ListJobsInput) (*iot.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *iot.ListJobsInput) *IotListJobsResult

    ListMitigationActions(ctx workflow.Context, input *iot.ListMitigationActionsInput) (*iot.ListMitigationActionsOutput, error)
    ListMitigationActionsAsync(ctx workflow.Context, input *iot.ListMitigationActionsInput) *IotListMitigationActionsResult

    ListOTAUpdates(ctx workflow.Context, input *iot.ListOTAUpdatesInput) (*iot.ListOTAUpdatesOutput, error)
    ListOTAUpdatesAsync(ctx workflow.Context, input *iot.ListOTAUpdatesInput) *IotListOTAUpdatesResult

    ListOutgoingCertificates(ctx workflow.Context, input *iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error)
    ListOutgoingCertificatesAsync(ctx workflow.Context, input *iot.ListOutgoingCertificatesInput) *IotListOutgoingCertificatesResult

    ListPolicies(ctx workflow.Context, input *iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error)
    ListPoliciesAsync(ctx workflow.Context, input *iot.ListPoliciesInput) *IotListPoliciesResult

    ListPolicyPrincipals(ctx workflow.Context, input *iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error)
    ListPolicyPrincipalsAsync(ctx workflow.Context, input *iot.ListPolicyPrincipalsInput) *IotListPolicyPrincipalsResult

    ListPolicyVersions(ctx workflow.Context, input *iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error)
    ListPolicyVersionsAsync(ctx workflow.Context, input *iot.ListPolicyVersionsInput) *IotListPolicyVersionsResult

    ListPrincipalPolicies(ctx workflow.Context, input *iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error)
    ListPrincipalPoliciesAsync(ctx workflow.Context, input *iot.ListPrincipalPoliciesInput) *IotListPrincipalPoliciesResult

    ListPrincipalThings(ctx workflow.Context, input *iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error)
    ListPrincipalThingsAsync(ctx workflow.Context, input *iot.ListPrincipalThingsInput) *IotListPrincipalThingsResult

    ListProvisioningTemplateVersions(ctx workflow.Context, input *iot.ListProvisioningTemplateVersionsInput) (*iot.ListProvisioningTemplateVersionsOutput, error)
    ListProvisioningTemplateVersionsAsync(ctx workflow.Context, input *iot.ListProvisioningTemplateVersionsInput) *IotListProvisioningTemplateVersionsResult

    ListProvisioningTemplates(ctx workflow.Context, input *iot.ListProvisioningTemplatesInput) (*iot.ListProvisioningTemplatesOutput, error)
    ListProvisioningTemplatesAsync(ctx workflow.Context, input *iot.ListProvisioningTemplatesInput) *IotListProvisioningTemplatesResult

    ListRoleAliases(ctx workflow.Context, input *iot.ListRoleAliasesInput) (*iot.ListRoleAliasesOutput, error)
    ListRoleAliasesAsync(ctx workflow.Context, input *iot.ListRoleAliasesInput) *IotListRoleAliasesResult

    ListScheduledAudits(ctx workflow.Context, input *iot.ListScheduledAuditsInput) (*iot.ListScheduledAuditsOutput, error)
    ListScheduledAuditsAsync(ctx workflow.Context, input *iot.ListScheduledAuditsInput) *IotListScheduledAuditsResult

    ListSecurityProfiles(ctx workflow.Context, input *iot.ListSecurityProfilesInput) (*iot.ListSecurityProfilesOutput, error)
    ListSecurityProfilesAsync(ctx workflow.Context, input *iot.ListSecurityProfilesInput) *IotListSecurityProfilesResult

    ListSecurityProfilesForTarget(ctx workflow.Context, input *iot.ListSecurityProfilesForTargetInput) (*iot.ListSecurityProfilesForTargetOutput, error)
    ListSecurityProfilesForTargetAsync(ctx workflow.Context, input *iot.ListSecurityProfilesForTargetInput) *IotListSecurityProfilesForTargetResult

    ListStreams(ctx workflow.Context, input *iot.ListStreamsInput) (*iot.ListStreamsOutput, error)
    ListStreamsAsync(ctx workflow.Context, input *iot.ListStreamsInput) *IotListStreamsResult

    ListTagsForResource(ctx workflow.Context, input *iot.ListTagsForResourceInput) (*iot.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *iot.ListTagsForResourceInput) *IotListTagsForResourceResult

    ListTargetsForPolicy(ctx workflow.Context, input *iot.ListTargetsForPolicyInput) (*iot.ListTargetsForPolicyOutput, error)
    ListTargetsForPolicyAsync(ctx workflow.Context, input *iot.ListTargetsForPolicyInput) *IotListTargetsForPolicyResult

    ListTargetsForSecurityProfile(ctx workflow.Context, input *iot.ListTargetsForSecurityProfileInput) (*iot.ListTargetsForSecurityProfileOutput, error)
    ListTargetsForSecurityProfileAsync(ctx workflow.Context, input *iot.ListTargetsForSecurityProfileInput) *IotListTargetsForSecurityProfileResult

    ListThingGroups(ctx workflow.Context, input *iot.ListThingGroupsInput) (*iot.ListThingGroupsOutput, error)
    ListThingGroupsAsync(ctx workflow.Context, input *iot.ListThingGroupsInput) *IotListThingGroupsResult

    ListThingGroupsForThing(ctx workflow.Context, input *iot.ListThingGroupsForThingInput) (*iot.ListThingGroupsForThingOutput, error)
    ListThingGroupsForThingAsync(ctx workflow.Context, input *iot.ListThingGroupsForThingInput) *IotListThingGroupsForThingResult

    ListThingPrincipals(ctx workflow.Context, input *iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error)
    ListThingPrincipalsAsync(ctx workflow.Context, input *iot.ListThingPrincipalsInput) *IotListThingPrincipalsResult

    ListThingRegistrationTaskReports(ctx workflow.Context, input *iot.ListThingRegistrationTaskReportsInput) (*iot.ListThingRegistrationTaskReportsOutput, error)
    ListThingRegistrationTaskReportsAsync(ctx workflow.Context, input *iot.ListThingRegistrationTaskReportsInput) *IotListThingRegistrationTaskReportsResult

    ListThingRegistrationTasks(ctx workflow.Context, input *iot.ListThingRegistrationTasksInput) (*iot.ListThingRegistrationTasksOutput, error)
    ListThingRegistrationTasksAsync(ctx workflow.Context, input *iot.ListThingRegistrationTasksInput) *IotListThingRegistrationTasksResult

    ListThingTypes(ctx workflow.Context, input *iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error)
    ListThingTypesAsync(ctx workflow.Context, input *iot.ListThingTypesInput) *IotListThingTypesResult

    ListThings(ctx workflow.Context, input *iot.ListThingsInput) (*iot.ListThingsOutput, error)
    ListThingsAsync(ctx workflow.Context, input *iot.ListThingsInput) *IotListThingsResult

    ListThingsInBillingGroup(ctx workflow.Context, input *iot.ListThingsInBillingGroupInput) (*iot.ListThingsInBillingGroupOutput, error)
    ListThingsInBillingGroupAsync(ctx workflow.Context, input *iot.ListThingsInBillingGroupInput) *IotListThingsInBillingGroupResult

    ListThingsInThingGroup(ctx workflow.Context, input *iot.ListThingsInThingGroupInput) (*iot.ListThingsInThingGroupOutput, error)
    ListThingsInThingGroupAsync(ctx workflow.Context, input *iot.ListThingsInThingGroupInput) *IotListThingsInThingGroupResult

    ListTopicRuleDestinations(ctx workflow.Context, input *iot.ListTopicRuleDestinationsInput) (*iot.ListTopicRuleDestinationsOutput, error)
    ListTopicRuleDestinationsAsync(ctx workflow.Context, input *iot.ListTopicRuleDestinationsInput) *IotListTopicRuleDestinationsResult

    ListTopicRules(ctx workflow.Context, input *iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error)
    ListTopicRulesAsync(ctx workflow.Context, input *iot.ListTopicRulesInput) *IotListTopicRulesResult

    ListV2LoggingLevels(ctx workflow.Context, input *iot.ListV2LoggingLevelsInput) (*iot.ListV2LoggingLevelsOutput, error)
    ListV2LoggingLevelsAsync(ctx workflow.Context, input *iot.ListV2LoggingLevelsInput) *IotListV2LoggingLevelsResult

    ListViolationEvents(ctx workflow.Context, input *iot.ListViolationEventsInput) (*iot.ListViolationEventsOutput, error)
    ListViolationEventsAsync(ctx workflow.Context, input *iot.ListViolationEventsInput) *IotListViolationEventsResult

    RegisterCACertificate(ctx workflow.Context, input *iot.RegisterCACertificateInput) (*iot.RegisterCACertificateOutput, error)
    RegisterCACertificateAsync(ctx workflow.Context, input *iot.RegisterCACertificateInput) *IotRegisterCACertificateResult

    RegisterCertificate(ctx workflow.Context, input *iot.RegisterCertificateInput) (*iot.RegisterCertificateOutput, error)
    RegisterCertificateAsync(ctx workflow.Context, input *iot.RegisterCertificateInput) *IotRegisterCertificateResult

    RegisterCertificateWithoutCA(ctx workflow.Context, input *iot.RegisterCertificateWithoutCAInput) (*iot.RegisterCertificateWithoutCAOutput, error)
    RegisterCertificateWithoutCAAsync(ctx workflow.Context, input *iot.RegisterCertificateWithoutCAInput) *IotRegisterCertificateWithoutCAResult

    RegisterThing(ctx workflow.Context, input *iot.RegisterThingInput) (*iot.RegisterThingOutput, error)
    RegisterThingAsync(ctx workflow.Context, input *iot.RegisterThingInput) *IotRegisterThingResult

    RejectCertificateTransfer(ctx workflow.Context, input *iot.RejectCertificateTransferInput) (*iot.RejectCertificateTransferOutput, error)
    RejectCertificateTransferAsync(ctx workflow.Context, input *iot.RejectCertificateTransferInput) *IotRejectCertificateTransferResult

    RemoveThingFromBillingGroup(ctx workflow.Context, input *iot.RemoveThingFromBillingGroupInput) (*iot.RemoveThingFromBillingGroupOutput, error)
    RemoveThingFromBillingGroupAsync(ctx workflow.Context, input *iot.RemoveThingFromBillingGroupInput) *IotRemoveThingFromBillingGroupResult

    RemoveThingFromThingGroup(ctx workflow.Context, input *iot.RemoveThingFromThingGroupInput) (*iot.RemoveThingFromThingGroupOutput, error)
    RemoveThingFromThingGroupAsync(ctx workflow.Context, input *iot.RemoveThingFromThingGroupInput) *IotRemoveThingFromThingGroupResult

    ReplaceTopicRule(ctx workflow.Context, input *iot.ReplaceTopicRuleInput) (*iot.ReplaceTopicRuleOutput, error)
    ReplaceTopicRuleAsync(ctx workflow.Context, input *iot.ReplaceTopicRuleInput) *IotReplaceTopicRuleResult

    SearchIndex(ctx workflow.Context, input *iot.SearchIndexInput) (*iot.SearchIndexOutput, error)
    SearchIndexAsync(ctx workflow.Context, input *iot.SearchIndexInput) *IotSearchIndexResult

    SetDefaultAuthorizer(ctx workflow.Context, input *iot.SetDefaultAuthorizerInput) (*iot.SetDefaultAuthorizerOutput, error)
    SetDefaultAuthorizerAsync(ctx workflow.Context, input *iot.SetDefaultAuthorizerInput) *IotSetDefaultAuthorizerResult

    SetDefaultPolicyVersion(ctx workflow.Context, input *iot.SetDefaultPolicyVersionInput) (*iot.SetDefaultPolicyVersionOutput, error)
    SetDefaultPolicyVersionAsync(ctx workflow.Context, input *iot.SetDefaultPolicyVersionInput) *IotSetDefaultPolicyVersionResult

    SetLoggingOptions(ctx workflow.Context, input *iot.SetLoggingOptionsInput) (*iot.SetLoggingOptionsOutput, error)
    SetLoggingOptionsAsync(ctx workflow.Context, input *iot.SetLoggingOptionsInput) *IotSetLoggingOptionsResult

    SetV2LoggingLevel(ctx workflow.Context, input *iot.SetV2LoggingLevelInput) (*iot.SetV2LoggingLevelOutput, error)
    SetV2LoggingLevelAsync(ctx workflow.Context, input *iot.SetV2LoggingLevelInput) *IotSetV2LoggingLevelResult

    SetV2LoggingOptions(ctx workflow.Context, input *iot.SetV2LoggingOptionsInput) (*iot.SetV2LoggingOptionsOutput, error)
    SetV2LoggingOptionsAsync(ctx workflow.Context, input *iot.SetV2LoggingOptionsInput) *IotSetV2LoggingOptionsResult

    StartAuditMitigationActionsTask(ctx workflow.Context, input *iot.StartAuditMitigationActionsTaskInput) (*iot.StartAuditMitigationActionsTaskOutput, error)
    StartAuditMitigationActionsTaskAsync(ctx workflow.Context, input *iot.StartAuditMitigationActionsTaskInput) *IotStartAuditMitigationActionsTaskResult

    StartOnDemandAuditTask(ctx workflow.Context, input *iot.StartOnDemandAuditTaskInput) (*iot.StartOnDemandAuditTaskOutput, error)
    StartOnDemandAuditTaskAsync(ctx workflow.Context, input *iot.StartOnDemandAuditTaskInput) *IotStartOnDemandAuditTaskResult

    StartThingRegistrationTask(ctx workflow.Context, input *iot.StartThingRegistrationTaskInput) (*iot.StartThingRegistrationTaskOutput, error)
    StartThingRegistrationTaskAsync(ctx workflow.Context, input *iot.StartThingRegistrationTaskInput) *IotStartThingRegistrationTaskResult

    StopThingRegistrationTask(ctx workflow.Context, input *iot.StopThingRegistrationTaskInput) (*iot.StopThingRegistrationTaskOutput, error)
    StopThingRegistrationTaskAsync(ctx workflow.Context, input *iot.StopThingRegistrationTaskInput) *IotStopThingRegistrationTaskResult

    TagResource(ctx workflow.Context, input *iot.TagResourceInput) (*iot.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *iot.TagResourceInput) *IotTagResourceResult

    TestAuthorization(ctx workflow.Context, input *iot.TestAuthorizationInput) (*iot.TestAuthorizationOutput, error)
    TestAuthorizationAsync(ctx workflow.Context, input *iot.TestAuthorizationInput) *IotTestAuthorizationResult

    TestInvokeAuthorizer(ctx workflow.Context, input *iot.TestInvokeAuthorizerInput) (*iot.TestInvokeAuthorizerOutput, error)
    TestInvokeAuthorizerAsync(ctx workflow.Context, input *iot.TestInvokeAuthorizerInput) *IotTestInvokeAuthorizerResult

    TransferCertificate(ctx workflow.Context, input *iot.TransferCertificateInput) (*iot.TransferCertificateOutput, error)
    TransferCertificateAsync(ctx workflow.Context, input *iot.TransferCertificateInput) *IotTransferCertificateResult

    UntagResource(ctx workflow.Context, input *iot.UntagResourceInput) (*iot.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *iot.UntagResourceInput) *IotUntagResourceResult

    UpdateAccountAuditConfiguration(ctx workflow.Context, input *iot.UpdateAccountAuditConfigurationInput) (*iot.UpdateAccountAuditConfigurationOutput, error)
    UpdateAccountAuditConfigurationAsync(ctx workflow.Context, input *iot.UpdateAccountAuditConfigurationInput) *IotUpdateAccountAuditConfigurationResult

    UpdateAuditSuppression(ctx workflow.Context, input *iot.UpdateAuditSuppressionInput) (*iot.UpdateAuditSuppressionOutput, error)
    UpdateAuditSuppressionAsync(ctx workflow.Context, input *iot.UpdateAuditSuppressionInput) *IotUpdateAuditSuppressionResult

    UpdateAuthorizer(ctx workflow.Context, input *iot.UpdateAuthorizerInput) (*iot.UpdateAuthorizerOutput, error)
    UpdateAuthorizerAsync(ctx workflow.Context, input *iot.UpdateAuthorizerInput) *IotUpdateAuthorizerResult

    UpdateBillingGroup(ctx workflow.Context, input *iot.UpdateBillingGroupInput) (*iot.UpdateBillingGroupOutput, error)
    UpdateBillingGroupAsync(ctx workflow.Context, input *iot.UpdateBillingGroupInput) *IotUpdateBillingGroupResult

    UpdateCACertificate(ctx workflow.Context, input *iot.UpdateCACertificateInput) (*iot.UpdateCACertificateOutput, error)
    UpdateCACertificateAsync(ctx workflow.Context, input *iot.UpdateCACertificateInput) *IotUpdateCACertificateResult

    UpdateCertificate(ctx workflow.Context, input *iot.UpdateCertificateInput) (*iot.UpdateCertificateOutput, error)
    UpdateCertificateAsync(ctx workflow.Context, input *iot.UpdateCertificateInput) *IotUpdateCertificateResult

    UpdateDimension(ctx workflow.Context, input *iot.UpdateDimensionInput) (*iot.UpdateDimensionOutput, error)
    UpdateDimensionAsync(ctx workflow.Context, input *iot.UpdateDimensionInput) *IotUpdateDimensionResult

    UpdateDomainConfiguration(ctx workflow.Context, input *iot.UpdateDomainConfigurationInput) (*iot.UpdateDomainConfigurationOutput, error)
    UpdateDomainConfigurationAsync(ctx workflow.Context, input *iot.UpdateDomainConfigurationInput) *IotUpdateDomainConfigurationResult

    UpdateDynamicThingGroup(ctx workflow.Context, input *iot.UpdateDynamicThingGroupInput) (*iot.UpdateDynamicThingGroupOutput, error)
    UpdateDynamicThingGroupAsync(ctx workflow.Context, input *iot.UpdateDynamicThingGroupInput) *IotUpdateDynamicThingGroupResult

    UpdateEventConfigurations(ctx workflow.Context, input *iot.UpdateEventConfigurationsInput) (*iot.UpdateEventConfigurationsOutput, error)
    UpdateEventConfigurationsAsync(ctx workflow.Context, input *iot.UpdateEventConfigurationsInput) *IotUpdateEventConfigurationsResult

    UpdateIndexingConfiguration(ctx workflow.Context, input *iot.UpdateIndexingConfigurationInput) (*iot.UpdateIndexingConfigurationOutput, error)
    UpdateIndexingConfigurationAsync(ctx workflow.Context, input *iot.UpdateIndexingConfigurationInput) *IotUpdateIndexingConfigurationResult

    UpdateJob(ctx workflow.Context, input *iot.UpdateJobInput) (*iot.UpdateJobOutput, error)
    UpdateJobAsync(ctx workflow.Context, input *iot.UpdateJobInput) *IotUpdateJobResult

    UpdateMitigationAction(ctx workflow.Context, input *iot.UpdateMitigationActionInput) (*iot.UpdateMitigationActionOutput, error)
    UpdateMitigationActionAsync(ctx workflow.Context, input *iot.UpdateMitigationActionInput) *IotUpdateMitigationActionResult

    UpdateProvisioningTemplate(ctx workflow.Context, input *iot.UpdateProvisioningTemplateInput) (*iot.UpdateProvisioningTemplateOutput, error)
    UpdateProvisioningTemplateAsync(ctx workflow.Context, input *iot.UpdateProvisioningTemplateInput) *IotUpdateProvisioningTemplateResult

    UpdateRoleAlias(ctx workflow.Context, input *iot.UpdateRoleAliasInput) (*iot.UpdateRoleAliasOutput, error)
    UpdateRoleAliasAsync(ctx workflow.Context, input *iot.UpdateRoleAliasInput) *IotUpdateRoleAliasResult

    UpdateScheduledAudit(ctx workflow.Context, input *iot.UpdateScheduledAuditInput) (*iot.UpdateScheduledAuditOutput, error)
    UpdateScheduledAuditAsync(ctx workflow.Context, input *iot.UpdateScheduledAuditInput) *IotUpdateScheduledAuditResult

    UpdateSecurityProfile(ctx workflow.Context, input *iot.UpdateSecurityProfileInput) (*iot.UpdateSecurityProfileOutput, error)
    UpdateSecurityProfileAsync(ctx workflow.Context, input *iot.UpdateSecurityProfileInput) *IotUpdateSecurityProfileResult

    UpdateStream(ctx workflow.Context, input *iot.UpdateStreamInput) (*iot.UpdateStreamOutput, error)
    UpdateStreamAsync(ctx workflow.Context, input *iot.UpdateStreamInput) *IotUpdateStreamResult

    UpdateThing(ctx workflow.Context, input *iot.UpdateThingInput) (*iot.UpdateThingOutput, error)
    UpdateThingAsync(ctx workflow.Context, input *iot.UpdateThingInput) *IotUpdateThingResult

    UpdateThingGroup(ctx workflow.Context, input *iot.UpdateThingGroupInput) (*iot.UpdateThingGroupOutput, error)
    UpdateThingGroupAsync(ctx workflow.Context, input *iot.UpdateThingGroupInput) *IotUpdateThingGroupResult

    UpdateThingGroupsForThing(ctx workflow.Context, input *iot.UpdateThingGroupsForThingInput) (*iot.UpdateThingGroupsForThingOutput, error)
    UpdateThingGroupsForThingAsync(ctx workflow.Context, input *iot.UpdateThingGroupsForThingInput) *IotUpdateThingGroupsForThingResult

    UpdateTopicRuleDestination(ctx workflow.Context, input *iot.UpdateTopicRuleDestinationInput) (*iot.UpdateTopicRuleDestinationOutput, error)
    UpdateTopicRuleDestinationAsync(ctx workflow.Context, input *iot.UpdateTopicRuleDestinationInput) *IotUpdateTopicRuleDestinationResult

    ValidateSecurityProfileBehaviors(ctx workflow.Context, input *iot.ValidateSecurityProfileBehaviorsInput) (*iot.ValidateSecurityProfileBehaviorsOutput, error)
    ValidateSecurityProfileBehaviorsAsync(ctx workflow.Context, input *iot.ValidateSecurityProfileBehaviorsInput) *IotValidateSecurityProfileBehaviorsResult
}

type IotAcceptCertificateTransferResult struct {
	Result workflow.Future
}

func (r *IotAcceptCertificateTransferResult) Get(ctx workflow.Context) (*iot.AcceptCertificateTransferOutput, error) {
    var output iot.AcceptCertificateTransferOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAddThingToBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotAddThingToBillingGroupResult) Get(ctx workflow.Context) (*iot.AddThingToBillingGroupOutput, error) {
    var output iot.AddThingToBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAddThingToThingGroupResult struct {
	Result workflow.Future
}

func (r *IotAddThingToThingGroupResult) Get(ctx workflow.Context) (*iot.AddThingToThingGroupOutput, error) {
    var output iot.AddThingToThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAssociateTargetsWithJobResult struct {
	Result workflow.Future
}

func (r *IotAssociateTargetsWithJobResult) Get(ctx workflow.Context) (*iot.AssociateTargetsWithJobOutput, error) {
    var output iot.AssociateTargetsWithJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAttachPolicyResult struct {
	Result workflow.Future
}

func (r *IotAttachPolicyResult) Get(ctx workflow.Context) (*iot.AttachPolicyOutput, error) {
    var output iot.AttachPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAttachPrincipalPolicyResult struct {
	Result workflow.Future
}

func (r *IotAttachPrincipalPolicyResult) Get(ctx workflow.Context) (*iot.AttachPrincipalPolicyOutput, error) {
    var output iot.AttachPrincipalPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAttachSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotAttachSecurityProfileResult) Get(ctx workflow.Context) (*iot.AttachSecurityProfileOutput, error) {
    var output iot.AttachSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotAttachThingPrincipalResult struct {
	Result workflow.Future
}

func (r *IotAttachThingPrincipalResult) Get(ctx workflow.Context) (*iot.AttachThingPrincipalOutput, error) {
    var output iot.AttachThingPrincipalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCancelAuditMitigationActionsTaskResult struct {
	Result workflow.Future
}

func (r *IotCancelAuditMitigationActionsTaskResult) Get(ctx workflow.Context) (*iot.CancelAuditMitigationActionsTaskOutput, error) {
    var output iot.CancelAuditMitigationActionsTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCancelAuditTaskResult struct {
	Result workflow.Future
}

func (r *IotCancelAuditTaskResult) Get(ctx workflow.Context) (*iot.CancelAuditTaskOutput, error) {
    var output iot.CancelAuditTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCancelCertificateTransferResult struct {
	Result workflow.Future
}

func (r *IotCancelCertificateTransferResult) Get(ctx workflow.Context) (*iot.CancelCertificateTransferOutput, error) {
    var output iot.CancelCertificateTransferOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCancelJobResult struct {
	Result workflow.Future
}

func (r *IotCancelJobResult) Get(ctx workflow.Context) (*iot.CancelJobOutput, error) {
    var output iot.CancelJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCancelJobExecutionResult struct {
	Result workflow.Future
}

func (r *IotCancelJobExecutionResult) Get(ctx workflow.Context) (*iot.CancelJobExecutionOutput, error) {
    var output iot.CancelJobExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotClearDefaultAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotClearDefaultAuthorizerResult) Get(ctx workflow.Context) (*iot.ClearDefaultAuthorizerOutput, error) {
    var output iot.ClearDefaultAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotConfirmTopicRuleDestinationResult struct {
	Result workflow.Future
}

func (r *IotConfirmTopicRuleDestinationResult) Get(ctx workflow.Context) (*iot.ConfirmTopicRuleDestinationOutput, error) {
    var output iot.ConfirmTopicRuleDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateAuditSuppressionResult struct {
	Result workflow.Future
}

func (r *IotCreateAuditSuppressionResult) Get(ctx workflow.Context) (*iot.CreateAuditSuppressionOutput, error) {
    var output iot.CreateAuditSuppressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotCreateAuthorizerResult) Get(ctx workflow.Context) (*iot.CreateAuthorizerOutput, error) {
    var output iot.CreateAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotCreateBillingGroupResult) Get(ctx workflow.Context) (*iot.CreateBillingGroupOutput, error) {
    var output iot.CreateBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateCertificateFromCsrResult struct {
	Result workflow.Future
}

func (r *IotCreateCertificateFromCsrResult) Get(ctx workflow.Context) (*iot.CreateCertificateFromCsrOutput, error) {
    var output iot.CreateCertificateFromCsrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateDimensionResult struct {
	Result workflow.Future
}

func (r *IotCreateDimensionResult) Get(ctx workflow.Context) (*iot.CreateDimensionOutput, error) {
    var output iot.CreateDimensionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateDomainConfigurationResult struct {
	Result workflow.Future
}

func (r *IotCreateDomainConfigurationResult) Get(ctx workflow.Context) (*iot.CreateDomainConfigurationOutput, error) {
    var output iot.CreateDomainConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateDynamicThingGroupResult struct {
	Result workflow.Future
}

func (r *IotCreateDynamicThingGroupResult) Get(ctx workflow.Context) (*iot.CreateDynamicThingGroupOutput, error) {
    var output iot.CreateDynamicThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateJobResult struct {
	Result workflow.Future
}

func (r *IotCreateJobResult) Get(ctx workflow.Context) (*iot.CreateJobOutput, error) {
    var output iot.CreateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateKeysAndCertificateResult struct {
	Result workflow.Future
}

func (r *IotCreateKeysAndCertificateResult) Get(ctx workflow.Context) (*iot.CreateKeysAndCertificateOutput, error) {
    var output iot.CreateKeysAndCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateMitigationActionResult struct {
	Result workflow.Future
}

func (r *IotCreateMitigationActionResult) Get(ctx workflow.Context) (*iot.CreateMitigationActionOutput, error) {
    var output iot.CreateMitigationActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateOTAUpdateResult struct {
	Result workflow.Future
}

func (r *IotCreateOTAUpdateResult) Get(ctx workflow.Context) (*iot.CreateOTAUpdateOutput, error) {
    var output iot.CreateOTAUpdateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreatePolicyResult struct {
	Result workflow.Future
}

func (r *IotCreatePolicyResult) Get(ctx workflow.Context) (*iot.CreatePolicyOutput, error) {
    var output iot.CreatePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreatePolicyVersionResult struct {
	Result workflow.Future
}

func (r *IotCreatePolicyVersionResult) Get(ctx workflow.Context) (*iot.CreatePolicyVersionOutput, error) {
    var output iot.CreatePolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateProvisioningClaimResult struct {
	Result workflow.Future
}

func (r *IotCreateProvisioningClaimResult) Get(ctx workflow.Context) (*iot.CreateProvisioningClaimOutput, error) {
    var output iot.CreateProvisioningClaimOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateProvisioningTemplateResult struct {
	Result workflow.Future
}

func (r *IotCreateProvisioningTemplateResult) Get(ctx workflow.Context) (*iot.CreateProvisioningTemplateOutput, error) {
    var output iot.CreateProvisioningTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateProvisioningTemplateVersionResult struct {
	Result workflow.Future
}

func (r *IotCreateProvisioningTemplateVersionResult) Get(ctx workflow.Context) (*iot.CreateProvisioningTemplateVersionOutput, error) {
    var output iot.CreateProvisioningTemplateVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateRoleAliasResult struct {
	Result workflow.Future
}

func (r *IotCreateRoleAliasResult) Get(ctx workflow.Context) (*iot.CreateRoleAliasOutput, error) {
    var output iot.CreateRoleAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateScheduledAuditResult struct {
	Result workflow.Future
}

func (r *IotCreateScheduledAuditResult) Get(ctx workflow.Context) (*iot.CreateScheduledAuditOutput, error) {
    var output iot.CreateScheduledAuditOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotCreateSecurityProfileResult) Get(ctx workflow.Context) (*iot.CreateSecurityProfileOutput, error) {
    var output iot.CreateSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateStreamResult struct {
	Result workflow.Future
}

func (r *IotCreateStreamResult) Get(ctx workflow.Context) (*iot.CreateStreamOutput, error) {
    var output iot.CreateStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateThingResult struct {
	Result workflow.Future
}

func (r *IotCreateThingResult) Get(ctx workflow.Context) (*iot.CreateThingOutput, error) {
    var output iot.CreateThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateThingGroupResult struct {
	Result workflow.Future
}

func (r *IotCreateThingGroupResult) Get(ctx workflow.Context) (*iot.CreateThingGroupOutput, error) {
    var output iot.CreateThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateThingTypeResult struct {
	Result workflow.Future
}

func (r *IotCreateThingTypeResult) Get(ctx workflow.Context) (*iot.CreateThingTypeOutput, error) {
    var output iot.CreateThingTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateTopicRuleResult struct {
	Result workflow.Future
}

func (r *IotCreateTopicRuleResult) Get(ctx workflow.Context) (*iot.CreateTopicRuleOutput, error) {
    var output iot.CreateTopicRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotCreateTopicRuleDestinationResult struct {
	Result workflow.Future
}

func (r *IotCreateTopicRuleDestinationResult) Get(ctx workflow.Context) (*iot.CreateTopicRuleDestinationOutput, error) {
    var output iot.CreateTopicRuleDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteAccountAuditConfigurationResult struct {
	Result workflow.Future
}

func (r *IotDeleteAccountAuditConfigurationResult) Get(ctx workflow.Context) (*iot.DeleteAccountAuditConfigurationOutput, error) {
    var output iot.DeleteAccountAuditConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteAuditSuppressionResult struct {
	Result workflow.Future
}

func (r *IotDeleteAuditSuppressionResult) Get(ctx workflow.Context) (*iot.DeleteAuditSuppressionOutput, error) {
    var output iot.DeleteAuditSuppressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotDeleteAuthorizerResult) Get(ctx workflow.Context) (*iot.DeleteAuthorizerOutput, error) {
    var output iot.DeleteAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotDeleteBillingGroupResult) Get(ctx workflow.Context) (*iot.DeleteBillingGroupOutput, error) {
    var output iot.DeleteBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteCACertificateResult struct {
	Result workflow.Future
}

func (r *IotDeleteCACertificateResult) Get(ctx workflow.Context) (*iot.DeleteCACertificateOutput, error) {
    var output iot.DeleteCACertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteCertificateResult struct {
	Result workflow.Future
}

func (r *IotDeleteCertificateResult) Get(ctx workflow.Context) (*iot.DeleteCertificateOutput, error) {
    var output iot.DeleteCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteDimensionResult struct {
	Result workflow.Future
}

func (r *IotDeleteDimensionResult) Get(ctx workflow.Context) (*iot.DeleteDimensionOutput, error) {
    var output iot.DeleteDimensionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteDomainConfigurationResult struct {
	Result workflow.Future
}

func (r *IotDeleteDomainConfigurationResult) Get(ctx workflow.Context) (*iot.DeleteDomainConfigurationOutput, error) {
    var output iot.DeleteDomainConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteDynamicThingGroupResult struct {
	Result workflow.Future
}

func (r *IotDeleteDynamicThingGroupResult) Get(ctx workflow.Context) (*iot.DeleteDynamicThingGroupOutput, error) {
    var output iot.DeleteDynamicThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteJobResult struct {
	Result workflow.Future
}

func (r *IotDeleteJobResult) Get(ctx workflow.Context) (*iot.DeleteJobOutput, error) {
    var output iot.DeleteJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteJobExecutionResult struct {
	Result workflow.Future
}

func (r *IotDeleteJobExecutionResult) Get(ctx workflow.Context) (*iot.DeleteJobExecutionOutput, error) {
    var output iot.DeleteJobExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteMitigationActionResult struct {
	Result workflow.Future
}

func (r *IotDeleteMitigationActionResult) Get(ctx workflow.Context) (*iot.DeleteMitigationActionOutput, error) {
    var output iot.DeleteMitigationActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteOTAUpdateResult struct {
	Result workflow.Future
}

func (r *IotDeleteOTAUpdateResult) Get(ctx workflow.Context) (*iot.DeleteOTAUpdateOutput, error) {
    var output iot.DeleteOTAUpdateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeletePolicyResult struct {
	Result workflow.Future
}

func (r *IotDeletePolicyResult) Get(ctx workflow.Context) (*iot.DeletePolicyOutput, error) {
    var output iot.DeletePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeletePolicyVersionResult struct {
	Result workflow.Future
}

func (r *IotDeletePolicyVersionResult) Get(ctx workflow.Context) (*iot.DeletePolicyVersionOutput, error) {
    var output iot.DeletePolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteProvisioningTemplateResult struct {
	Result workflow.Future
}

func (r *IotDeleteProvisioningTemplateResult) Get(ctx workflow.Context) (*iot.DeleteProvisioningTemplateOutput, error) {
    var output iot.DeleteProvisioningTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteProvisioningTemplateVersionResult struct {
	Result workflow.Future
}

func (r *IotDeleteProvisioningTemplateVersionResult) Get(ctx workflow.Context) (*iot.DeleteProvisioningTemplateVersionOutput, error) {
    var output iot.DeleteProvisioningTemplateVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteRegistrationCodeResult struct {
	Result workflow.Future
}

func (r *IotDeleteRegistrationCodeResult) Get(ctx workflow.Context) (*iot.DeleteRegistrationCodeOutput, error) {
    var output iot.DeleteRegistrationCodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteRoleAliasResult struct {
	Result workflow.Future
}

func (r *IotDeleteRoleAliasResult) Get(ctx workflow.Context) (*iot.DeleteRoleAliasOutput, error) {
    var output iot.DeleteRoleAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteScheduledAuditResult struct {
	Result workflow.Future
}

func (r *IotDeleteScheduledAuditResult) Get(ctx workflow.Context) (*iot.DeleteScheduledAuditOutput, error) {
    var output iot.DeleteScheduledAuditOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotDeleteSecurityProfileResult) Get(ctx workflow.Context) (*iot.DeleteSecurityProfileOutput, error) {
    var output iot.DeleteSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteStreamResult struct {
	Result workflow.Future
}

func (r *IotDeleteStreamResult) Get(ctx workflow.Context) (*iot.DeleteStreamOutput, error) {
    var output iot.DeleteStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteThingResult struct {
	Result workflow.Future
}

func (r *IotDeleteThingResult) Get(ctx workflow.Context) (*iot.DeleteThingOutput, error) {
    var output iot.DeleteThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteThingGroupResult struct {
	Result workflow.Future
}

func (r *IotDeleteThingGroupResult) Get(ctx workflow.Context) (*iot.DeleteThingGroupOutput, error) {
    var output iot.DeleteThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteThingTypeResult struct {
	Result workflow.Future
}

func (r *IotDeleteThingTypeResult) Get(ctx workflow.Context) (*iot.DeleteThingTypeOutput, error) {
    var output iot.DeleteThingTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteTopicRuleResult struct {
	Result workflow.Future
}

func (r *IotDeleteTopicRuleResult) Get(ctx workflow.Context) (*iot.DeleteTopicRuleOutput, error) {
    var output iot.DeleteTopicRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteTopicRuleDestinationResult struct {
	Result workflow.Future
}

func (r *IotDeleteTopicRuleDestinationResult) Get(ctx workflow.Context) (*iot.DeleteTopicRuleDestinationOutput, error) {
    var output iot.DeleteTopicRuleDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeleteV2LoggingLevelResult struct {
	Result workflow.Future
}

func (r *IotDeleteV2LoggingLevelResult) Get(ctx workflow.Context) (*iot.DeleteV2LoggingLevelOutput, error) {
    var output iot.DeleteV2LoggingLevelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDeprecateThingTypeResult struct {
	Result workflow.Future
}

func (r *IotDeprecateThingTypeResult) Get(ctx workflow.Context) (*iot.DeprecateThingTypeOutput, error) {
    var output iot.DeprecateThingTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeAccountAuditConfigurationResult struct {
	Result workflow.Future
}

func (r *IotDescribeAccountAuditConfigurationResult) Get(ctx workflow.Context) (*iot.DescribeAccountAuditConfigurationOutput, error) {
    var output iot.DescribeAccountAuditConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeAuditFindingResult struct {
	Result workflow.Future
}

func (r *IotDescribeAuditFindingResult) Get(ctx workflow.Context) (*iot.DescribeAuditFindingOutput, error) {
    var output iot.DescribeAuditFindingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeAuditMitigationActionsTaskResult struct {
	Result workflow.Future
}

func (r *IotDescribeAuditMitigationActionsTaskResult) Get(ctx workflow.Context) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
    var output iot.DescribeAuditMitigationActionsTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeAuditSuppressionResult struct {
	Result workflow.Future
}

func (r *IotDescribeAuditSuppressionResult) Get(ctx workflow.Context) (*iot.DescribeAuditSuppressionOutput, error) {
    var output iot.DescribeAuditSuppressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeAuditTaskResult struct {
	Result workflow.Future
}

func (r *IotDescribeAuditTaskResult) Get(ctx workflow.Context) (*iot.DescribeAuditTaskOutput, error) {
    var output iot.DescribeAuditTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotDescribeAuthorizerResult) Get(ctx workflow.Context) (*iot.DescribeAuthorizerOutput, error) {
    var output iot.DescribeAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotDescribeBillingGroupResult) Get(ctx workflow.Context) (*iot.DescribeBillingGroupOutput, error) {
    var output iot.DescribeBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeCACertificateResult struct {
	Result workflow.Future
}

func (r *IotDescribeCACertificateResult) Get(ctx workflow.Context) (*iot.DescribeCACertificateOutput, error) {
    var output iot.DescribeCACertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeCertificateResult struct {
	Result workflow.Future
}

func (r *IotDescribeCertificateResult) Get(ctx workflow.Context) (*iot.DescribeCertificateOutput, error) {
    var output iot.DescribeCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeDefaultAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotDescribeDefaultAuthorizerResult) Get(ctx workflow.Context) (*iot.DescribeDefaultAuthorizerOutput, error) {
    var output iot.DescribeDefaultAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeDimensionResult struct {
	Result workflow.Future
}

func (r *IotDescribeDimensionResult) Get(ctx workflow.Context) (*iot.DescribeDimensionOutput, error) {
    var output iot.DescribeDimensionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeDomainConfigurationResult struct {
	Result workflow.Future
}

func (r *IotDescribeDomainConfigurationResult) Get(ctx workflow.Context) (*iot.DescribeDomainConfigurationOutput, error) {
    var output iot.DescribeDomainConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeEndpointResult struct {
	Result workflow.Future
}

func (r *IotDescribeEndpointResult) Get(ctx workflow.Context) (*iot.DescribeEndpointOutput, error) {
    var output iot.DescribeEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeEventConfigurationsResult struct {
	Result workflow.Future
}

func (r *IotDescribeEventConfigurationsResult) Get(ctx workflow.Context) (*iot.DescribeEventConfigurationsOutput, error) {
    var output iot.DescribeEventConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeIndexResult struct {
	Result workflow.Future
}

func (r *IotDescribeIndexResult) Get(ctx workflow.Context) (*iot.DescribeIndexOutput, error) {
    var output iot.DescribeIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeJobResult struct {
	Result workflow.Future
}

func (r *IotDescribeJobResult) Get(ctx workflow.Context) (*iot.DescribeJobOutput, error) {
    var output iot.DescribeJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeJobExecutionResult struct {
	Result workflow.Future
}

func (r *IotDescribeJobExecutionResult) Get(ctx workflow.Context) (*iot.DescribeJobExecutionOutput, error) {
    var output iot.DescribeJobExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeMitigationActionResult struct {
	Result workflow.Future
}

func (r *IotDescribeMitigationActionResult) Get(ctx workflow.Context) (*iot.DescribeMitigationActionOutput, error) {
    var output iot.DescribeMitigationActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeProvisioningTemplateResult struct {
	Result workflow.Future
}

func (r *IotDescribeProvisioningTemplateResult) Get(ctx workflow.Context) (*iot.DescribeProvisioningTemplateOutput, error) {
    var output iot.DescribeProvisioningTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeProvisioningTemplateVersionResult struct {
	Result workflow.Future
}

func (r *IotDescribeProvisioningTemplateVersionResult) Get(ctx workflow.Context) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
    var output iot.DescribeProvisioningTemplateVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeRoleAliasResult struct {
	Result workflow.Future
}

func (r *IotDescribeRoleAliasResult) Get(ctx workflow.Context) (*iot.DescribeRoleAliasOutput, error) {
    var output iot.DescribeRoleAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeScheduledAuditResult struct {
	Result workflow.Future
}

func (r *IotDescribeScheduledAuditResult) Get(ctx workflow.Context) (*iot.DescribeScheduledAuditOutput, error) {
    var output iot.DescribeScheduledAuditOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotDescribeSecurityProfileResult) Get(ctx workflow.Context) (*iot.DescribeSecurityProfileOutput, error) {
    var output iot.DescribeSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeStreamResult struct {
	Result workflow.Future
}

func (r *IotDescribeStreamResult) Get(ctx workflow.Context) (*iot.DescribeStreamOutput, error) {
    var output iot.DescribeStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeThingResult struct {
	Result workflow.Future
}

func (r *IotDescribeThingResult) Get(ctx workflow.Context) (*iot.DescribeThingOutput, error) {
    var output iot.DescribeThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeThingGroupResult struct {
	Result workflow.Future
}

func (r *IotDescribeThingGroupResult) Get(ctx workflow.Context) (*iot.DescribeThingGroupOutput, error) {
    var output iot.DescribeThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeThingRegistrationTaskResult struct {
	Result workflow.Future
}

func (r *IotDescribeThingRegistrationTaskResult) Get(ctx workflow.Context) (*iot.DescribeThingRegistrationTaskOutput, error) {
    var output iot.DescribeThingRegistrationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDescribeThingTypeResult struct {
	Result workflow.Future
}

func (r *IotDescribeThingTypeResult) Get(ctx workflow.Context) (*iot.DescribeThingTypeOutput, error) {
    var output iot.DescribeThingTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDetachPolicyResult struct {
	Result workflow.Future
}

func (r *IotDetachPolicyResult) Get(ctx workflow.Context) (*iot.DetachPolicyOutput, error) {
    var output iot.DetachPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDetachPrincipalPolicyResult struct {
	Result workflow.Future
}

func (r *IotDetachPrincipalPolicyResult) Get(ctx workflow.Context) (*iot.DetachPrincipalPolicyOutput, error) {
    var output iot.DetachPrincipalPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDetachSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotDetachSecurityProfileResult) Get(ctx workflow.Context) (*iot.DetachSecurityProfileOutput, error) {
    var output iot.DetachSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDetachThingPrincipalResult struct {
	Result workflow.Future
}

func (r *IotDetachThingPrincipalResult) Get(ctx workflow.Context) (*iot.DetachThingPrincipalOutput, error) {
    var output iot.DetachThingPrincipalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotDisableTopicRuleResult struct {
	Result workflow.Future
}

func (r *IotDisableTopicRuleResult) Get(ctx workflow.Context) (*iot.DisableTopicRuleOutput, error) {
    var output iot.DisableTopicRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotEnableTopicRuleResult struct {
	Result workflow.Future
}

func (r *IotEnableTopicRuleResult) Get(ctx workflow.Context) (*iot.EnableTopicRuleOutput, error) {
    var output iot.EnableTopicRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetCardinalityResult struct {
	Result workflow.Future
}

func (r *IotGetCardinalityResult) Get(ctx workflow.Context) (*iot.GetCardinalityOutput, error) {
    var output iot.GetCardinalityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetEffectivePoliciesResult struct {
	Result workflow.Future
}

func (r *IotGetEffectivePoliciesResult) Get(ctx workflow.Context) (*iot.GetEffectivePoliciesOutput, error) {
    var output iot.GetEffectivePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetIndexingConfigurationResult struct {
	Result workflow.Future
}

func (r *IotGetIndexingConfigurationResult) Get(ctx workflow.Context) (*iot.GetIndexingConfigurationOutput, error) {
    var output iot.GetIndexingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetJobDocumentResult struct {
	Result workflow.Future
}

func (r *IotGetJobDocumentResult) Get(ctx workflow.Context) (*iot.GetJobDocumentOutput, error) {
    var output iot.GetJobDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotGetLoggingOptionsResult) Get(ctx workflow.Context) (*iot.GetLoggingOptionsOutput, error) {
    var output iot.GetLoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetOTAUpdateResult struct {
	Result workflow.Future
}

func (r *IotGetOTAUpdateResult) Get(ctx workflow.Context) (*iot.GetOTAUpdateOutput, error) {
    var output iot.GetOTAUpdateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetPercentilesResult struct {
	Result workflow.Future
}

func (r *IotGetPercentilesResult) Get(ctx workflow.Context) (*iot.GetPercentilesOutput, error) {
    var output iot.GetPercentilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetPolicyResult struct {
	Result workflow.Future
}

func (r *IotGetPolicyResult) Get(ctx workflow.Context) (*iot.GetPolicyOutput, error) {
    var output iot.GetPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetPolicyVersionResult struct {
	Result workflow.Future
}

func (r *IotGetPolicyVersionResult) Get(ctx workflow.Context) (*iot.GetPolicyVersionOutput, error) {
    var output iot.GetPolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetRegistrationCodeResult struct {
	Result workflow.Future
}

func (r *IotGetRegistrationCodeResult) Get(ctx workflow.Context) (*iot.GetRegistrationCodeOutput, error) {
    var output iot.GetRegistrationCodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetStatisticsResult struct {
	Result workflow.Future
}

func (r *IotGetStatisticsResult) Get(ctx workflow.Context) (*iot.GetStatisticsOutput, error) {
    var output iot.GetStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetTopicRuleResult struct {
	Result workflow.Future
}

func (r *IotGetTopicRuleResult) Get(ctx workflow.Context) (*iot.GetTopicRuleOutput, error) {
    var output iot.GetTopicRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetTopicRuleDestinationResult struct {
	Result workflow.Future
}

func (r *IotGetTopicRuleDestinationResult) Get(ctx workflow.Context) (*iot.GetTopicRuleDestinationOutput, error) {
    var output iot.GetTopicRuleDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotGetV2LoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotGetV2LoggingOptionsResult) Get(ctx workflow.Context) (*iot.GetV2LoggingOptionsOutput, error) {
    var output iot.GetV2LoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListActiveViolationsResult struct {
	Result workflow.Future
}

func (r *IotListActiveViolationsResult) Get(ctx workflow.Context) (*iot.ListActiveViolationsOutput, error) {
    var output iot.ListActiveViolationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAttachedPoliciesResult struct {
	Result workflow.Future
}

func (r *IotListAttachedPoliciesResult) Get(ctx workflow.Context) (*iot.ListAttachedPoliciesOutput, error) {
    var output iot.ListAttachedPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAuditFindingsResult struct {
	Result workflow.Future
}

func (r *IotListAuditFindingsResult) Get(ctx workflow.Context) (*iot.ListAuditFindingsOutput, error) {
    var output iot.ListAuditFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAuditMitigationActionsExecutionsResult struct {
	Result workflow.Future
}

func (r *IotListAuditMitigationActionsExecutionsResult) Get(ctx workflow.Context) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
    var output iot.ListAuditMitigationActionsExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAuditMitigationActionsTasksResult struct {
	Result workflow.Future
}

func (r *IotListAuditMitigationActionsTasksResult) Get(ctx workflow.Context) (*iot.ListAuditMitigationActionsTasksOutput, error) {
    var output iot.ListAuditMitigationActionsTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAuditSuppressionsResult struct {
	Result workflow.Future
}

func (r *IotListAuditSuppressionsResult) Get(ctx workflow.Context) (*iot.ListAuditSuppressionsOutput, error) {
    var output iot.ListAuditSuppressionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAuditTasksResult struct {
	Result workflow.Future
}

func (r *IotListAuditTasksResult) Get(ctx workflow.Context) (*iot.ListAuditTasksOutput, error) {
    var output iot.ListAuditTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListAuthorizersResult struct {
	Result workflow.Future
}

func (r *IotListAuthorizersResult) Get(ctx workflow.Context) (*iot.ListAuthorizersOutput, error) {
    var output iot.ListAuthorizersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListBillingGroupsResult struct {
	Result workflow.Future
}

func (r *IotListBillingGroupsResult) Get(ctx workflow.Context) (*iot.ListBillingGroupsOutput, error) {
    var output iot.ListBillingGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListCACertificatesResult struct {
	Result workflow.Future
}

func (r *IotListCACertificatesResult) Get(ctx workflow.Context) (*iot.ListCACertificatesOutput, error) {
    var output iot.ListCACertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListCertificatesResult struct {
	Result workflow.Future
}

func (r *IotListCertificatesResult) Get(ctx workflow.Context) (*iot.ListCertificatesOutput, error) {
    var output iot.ListCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListCertificatesByCAResult struct {
	Result workflow.Future
}

func (r *IotListCertificatesByCAResult) Get(ctx workflow.Context) (*iot.ListCertificatesByCAOutput, error) {
    var output iot.ListCertificatesByCAOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListDimensionsResult struct {
	Result workflow.Future
}

func (r *IotListDimensionsResult) Get(ctx workflow.Context) (*iot.ListDimensionsOutput, error) {
    var output iot.ListDimensionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListDomainConfigurationsResult struct {
	Result workflow.Future
}

func (r *IotListDomainConfigurationsResult) Get(ctx workflow.Context) (*iot.ListDomainConfigurationsOutput, error) {
    var output iot.ListDomainConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListIndicesResult struct {
	Result workflow.Future
}

func (r *IotListIndicesResult) Get(ctx workflow.Context) (*iot.ListIndicesOutput, error) {
    var output iot.ListIndicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListJobExecutionsForJobResult struct {
	Result workflow.Future
}

func (r *IotListJobExecutionsForJobResult) Get(ctx workflow.Context) (*iot.ListJobExecutionsForJobOutput, error) {
    var output iot.ListJobExecutionsForJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListJobExecutionsForThingResult struct {
	Result workflow.Future
}

func (r *IotListJobExecutionsForThingResult) Get(ctx workflow.Context) (*iot.ListJobExecutionsForThingOutput, error) {
    var output iot.ListJobExecutionsForThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListJobsResult struct {
	Result workflow.Future
}

func (r *IotListJobsResult) Get(ctx workflow.Context) (*iot.ListJobsOutput, error) {
    var output iot.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListMitigationActionsResult struct {
	Result workflow.Future
}

func (r *IotListMitigationActionsResult) Get(ctx workflow.Context) (*iot.ListMitigationActionsOutput, error) {
    var output iot.ListMitigationActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListOTAUpdatesResult struct {
	Result workflow.Future
}

func (r *IotListOTAUpdatesResult) Get(ctx workflow.Context) (*iot.ListOTAUpdatesOutput, error) {
    var output iot.ListOTAUpdatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListOutgoingCertificatesResult struct {
	Result workflow.Future
}

func (r *IotListOutgoingCertificatesResult) Get(ctx workflow.Context) (*iot.ListOutgoingCertificatesOutput, error) {
    var output iot.ListOutgoingCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListPoliciesResult struct {
	Result workflow.Future
}

func (r *IotListPoliciesResult) Get(ctx workflow.Context) (*iot.ListPoliciesOutput, error) {
    var output iot.ListPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListPolicyPrincipalsResult struct {
	Result workflow.Future
}

func (r *IotListPolicyPrincipalsResult) Get(ctx workflow.Context) (*iot.ListPolicyPrincipalsOutput, error) {
    var output iot.ListPolicyPrincipalsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListPolicyVersionsResult struct {
	Result workflow.Future
}

func (r *IotListPolicyVersionsResult) Get(ctx workflow.Context) (*iot.ListPolicyVersionsOutput, error) {
    var output iot.ListPolicyVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListPrincipalPoliciesResult struct {
	Result workflow.Future
}

func (r *IotListPrincipalPoliciesResult) Get(ctx workflow.Context) (*iot.ListPrincipalPoliciesOutput, error) {
    var output iot.ListPrincipalPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListPrincipalThingsResult struct {
	Result workflow.Future
}

func (r *IotListPrincipalThingsResult) Get(ctx workflow.Context) (*iot.ListPrincipalThingsOutput, error) {
    var output iot.ListPrincipalThingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListProvisioningTemplateVersionsResult struct {
	Result workflow.Future
}

func (r *IotListProvisioningTemplateVersionsResult) Get(ctx workflow.Context) (*iot.ListProvisioningTemplateVersionsOutput, error) {
    var output iot.ListProvisioningTemplateVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListProvisioningTemplatesResult struct {
	Result workflow.Future
}

func (r *IotListProvisioningTemplatesResult) Get(ctx workflow.Context) (*iot.ListProvisioningTemplatesOutput, error) {
    var output iot.ListProvisioningTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListRoleAliasesResult struct {
	Result workflow.Future
}

func (r *IotListRoleAliasesResult) Get(ctx workflow.Context) (*iot.ListRoleAliasesOutput, error) {
    var output iot.ListRoleAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListScheduledAuditsResult struct {
	Result workflow.Future
}

func (r *IotListScheduledAuditsResult) Get(ctx workflow.Context) (*iot.ListScheduledAuditsOutput, error) {
    var output iot.ListScheduledAuditsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListSecurityProfilesResult struct {
	Result workflow.Future
}

func (r *IotListSecurityProfilesResult) Get(ctx workflow.Context) (*iot.ListSecurityProfilesOutput, error) {
    var output iot.ListSecurityProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListSecurityProfilesForTargetResult struct {
	Result workflow.Future
}

func (r *IotListSecurityProfilesForTargetResult) Get(ctx workflow.Context) (*iot.ListSecurityProfilesForTargetOutput, error) {
    var output iot.ListSecurityProfilesForTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListStreamsResult struct {
	Result workflow.Future
}

func (r *IotListStreamsResult) Get(ctx workflow.Context) (*iot.ListStreamsOutput, error) {
    var output iot.ListStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IotListTagsForResourceResult) Get(ctx workflow.Context) (*iot.ListTagsForResourceOutput, error) {
    var output iot.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListTargetsForPolicyResult struct {
	Result workflow.Future
}

func (r *IotListTargetsForPolicyResult) Get(ctx workflow.Context) (*iot.ListTargetsForPolicyOutput, error) {
    var output iot.ListTargetsForPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListTargetsForSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotListTargetsForSecurityProfileResult) Get(ctx workflow.Context) (*iot.ListTargetsForSecurityProfileOutput, error) {
    var output iot.ListTargetsForSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingGroupsResult struct {
	Result workflow.Future
}

func (r *IotListThingGroupsResult) Get(ctx workflow.Context) (*iot.ListThingGroupsOutput, error) {
    var output iot.ListThingGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingGroupsForThingResult struct {
	Result workflow.Future
}

func (r *IotListThingGroupsForThingResult) Get(ctx workflow.Context) (*iot.ListThingGroupsForThingOutput, error) {
    var output iot.ListThingGroupsForThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingPrincipalsResult struct {
	Result workflow.Future
}

func (r *IotListThingPrincipalsResult) Get(ctx workflow.Context) (*iot.ListThingPrincipalsOutput, error) {
    var output iot.ListThingPrincipalsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingRegistrationTaskReportsResult struct {
	Result workflow.Future
}

func (r *IotListThingRegistrationTaskReportsResult) Get(ctx workflow.Context) (*iot.ListThingRegistrationTaskReportsOutput, error) {
    var output iot.ListThingRegistrationTaskReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingRegistrationTasksResult struct {
	Result workflow.Future
}

func (r *IotListThingRegistrationTasksResult) Get(ctx workflow.Context) (*iot.ListThingRegistrationTasksOutput, error) {
    var output iot.ListThingRegistrationTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingTypesResult struct {
	Result workflow.Future
}

func (r *IotListThingTypesResult) Get(ctx workflow.Context) (*iot.ListThingTypesOutput, error) {
    var output iot.ListThingTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingsResult struct {
	Result workflow.Future
}

func (r *IotListThingsResult) Get(ctx workflow.Context) (*iot.ListThingsOutput, error) {
    var output iot.ListThingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingsInBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotListThingsInBillingGroupResult) Get(ctx workflow.Context) (*iot.ListThingsInBillingGroupOutput, error) {
    var output iot.ListThingsInBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListThingsInThingGroupResult struct {
	Result workflow.Future
}

func (r *IotListThingsInThingGroupResult) Get(ctx workflow.Context) (*iot.ListThingsInThingGroupOutput, error) {
    var output iot.ListThingsInThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListTopicRuleDestinationsResult struct {
	Result workflow.Future
}

func (r *IotListTopicRuleDestinationsResult) Get(ctx workflow.Context) (*iot.ListTopicRuleDestinationsOutput, error) {
    var output iot.ListTopicRuleDestinationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListTopicRulesResult struct {
	Result workflow.Future
}

func (r *IotListTopicRulesResult) Get(ctx workflow.Context) (*iot.ListTopicRulesOutput, error) {
    var output iot.ListTopicRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListV2LoggingLevelsResult struct {
	Result workflow.Future
}

func (r *IotListV2LoggingLevelsResult) Get(ctx workflow.Context) (*iot.ListV2LoggingLevelsOutput, error) {
    var output iot.ListV2LoggingLevelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotListViolationEventsResult struct {
	Result workflow.Future
}

func (r *IotListViolationEventsResult) Get(ctx workflow.Context) (*iot.ListViolationEventsOutput, error) {
    var output iot.ListViolationEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRegisterCACertificateResult struct {
	Result workflow.Future
}

func (r *IotRegisterCACertificateResult) Get(ctx workflow.Context) (*iot.RegisterCACertificateOutput, error) {
    var output iot.RegisterCACertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRegisterCertificateResult struct {
	Result workflow.Future
}

func (r *IotRegisterCertificateResult) Get(ctx workflow.Context) (*iot.RegisterCertificateOutput, error) {
    var output iot.RegisterCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRegisterCertificateWithoutCAResult struct {
	Result workflow.Future
}

func (r *IotRegisterCertificateWithoutCAResult) Get(ctx workflow.Context) (*iot.RegisterCertificateWithoutCAOutput, error) {
    var output iot.RegisterCertificateWithoutCAOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRegisterThingResult struct {
	Result workflow.Future
}

func (r *IotRegisterThingResult) Get(ctx workflow.Context) (*iot.RegisterThingOutput, error) {
    var output iot.RegisterThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRejectCertificateTransferResult struct {
	Result workflow.Future
}

func (r *IotRejectCertificateTransferResult) Get(ctx workflow.Context) (*iot.RejectCertificateTransferOutput, error) {
    var output iot.RejectCertificateTransferOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRemoveThingFromBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotRemoveThingFromBillingGroupResult) Get(ctx workflow.Context) (*iot.RemoveThingFromBillingGroupOutput, error) {
    var output iot.RemoveThingFromBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotRemoveThingFromThingGroupResult struct {
	Result workflow.Future
}

func (r *IotRemoveThingFromThingGroupResult) Get(ctx workflow.Context) (*iot.RemoveThingFromThingGroupOutput, error) {
    var output iot.RemoveThingFromThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotReplaceTopicRuleResult struct {
	Result workflow.Future
}

func (r *IotReplaceTopicRuleResult) Get(ctx workflow.Context) (*iot.ReplaceTopicRuleOutput, error) {
    var output iot.ReplaceTopicRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotSearchIndexResult struct {
	Result workflow.Future
}

func (r *IotSearchIndexResult) Get(ctx workflow.Context) (*iot.SearchIndexOutput, error) {
    var output iot.SearchIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotSetDefaultAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotSetDefaultAuthorizerResult) Get(ctx workflow.Context) (*iot.SetDefaultAuthorizerOutput, error) {
    var output iot.SetDefaultAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotSetDefaultPolicyVersionResult struct {
	Result workflow.Future
}

func (r *IotSetDefaultPolicyVersionResult) Get(ctx workflow.Context) (*iot.SetDefaultPolicyVersionOutput, error) {
    var output iot.SetDefaultPolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotSetLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotSetLoggingOptionsResult) Get(ctx workflow.Context) (*iot.SetLoggingOptionsOutput, error) {
    var output iot.SetLoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotSetV2LoggingLevelResult struct {
	Result workflow.Future
}

func (r *IotSetV2LoggingLevelResult) Get(ctx workflow.Context) (*iot.SetV2LoggingLevelOutput, error) {
    var output iot.SetV2LoggingLevelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotSetV2LoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotSetV2LoggingOptionsResult) Get(ctx workflow.Context) (*iot.SetV2LoggingOptionsOutput, error) {
    var output iot.SetV2LoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotStartAuditMitigationActionsTaskResult struct {
	Result workflow.Future
}

func (r *IotStartAuditMitigationActionsTaskResult) Get(ctx workflow.Context) (*iot.StartAuditMitigationActionsTaskOutput, error) {
    var output iot.StartAuditMitigationActionsTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotStartOnDemandAuditTaskResult struct {
	Result workflow.Future
}

func (r *IotStartOnDemandAuditTaskResult) Get(ctx workflow.Context) (*iot.StartOnDemandAuditTaskOutput, error) {
    var output iot.StartOnDemandAuditTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotStartThingRegistrationTaskResult struct {
	Result workflow.Future
}

func (r *IotStartThingRegistrationTaskResult) Get(ctx workflow.Context) (*iot.StartThingRegistrationTaskOutput, error) {
    var output iot.StartThingRegistrationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotStopThingRegistrationTaskResult struct {
	Result workflow.Future
}

func (r *IotStopThingRegistrationTaskResult) Get(ctx workflow.Context) (*iot.StopThingRegistrationTaskOutput, error) {
    var output iot.StopThingRegistrationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotTagResourceResult struct {
	Result workflow.Future
}

func (r *IotTagResourceResult) Get(ctx workflow.Context) (*iot.TagResourceOutput, error) {
    var output iot.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotTestAuthorizationResult struct {
	Result workflow.Future
}

func (r *IotTestAuthorizationResult) Get(ctx workflow.Context) (*iot.TestAuthorizationOutput, error) {
    var output iot.TestAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotTestInvokeAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotTestInvokeAuthorizerResult) Get(ctx workflow.Context) (*iot.TestInvokeAuthorizerOutput, error) {
    var output iot.TestInvokeAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotTransferCertificateResult struct {
	Result workflow.Future
}

func (r *IotTransferCertificateResult) Get(ctx workflow.Context) (*iot.TransferCertificateOutput, error) {
    var output iot.TransferCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUntagResourceResult struct {
	Result workflow.Future
}

func (r *IotUntagResourceResult) Get(ctx workflow.Context) (*iot.UntagResourceOutput, error) {
    var output iot.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateAccountAuditConfigurationResult struct {
	Result workflow.Future
}

func (r *IotUpdateAccountAuditConfigurationResult) Get(ctx workflow.Context) (*iot.UpdateAccountAuditConfigurationOutput, error) {
    var output iot.UpdateAccountAuditConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateAuditSuppressionResult struct {
	Result workflow.Future
}

func (r *IotUpdateAuditSuppressionResult) Get(ctx workflow.Context) (*iot.UpdateAuditSuppressionOutput, error) {
    var output iot.UpdateAuditSuppressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateAuthorizerResult struct {
	Result workflow.Future
}

func (r *IotUpdateAuthorizerResult) Get(ctx workflow.Context) (*iot.UpdateAuthorizerOutput, error) {
    var output iot.UpdateAuthorizerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateBillingGroupResult struct {
	Result workflow.Future
}

func (r *IotUpdateBillingGroupResult) Get(ctx workflow.Context) (*iot.UpdateBillingGroupOutput, error) {
    var output iot.UpdateBillingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateCACertificateResult struct {
	Result workflow.Future
}

func (r *IotUpdateCACertificateResult) Get(ctx workflow.Context) (*iot.UpdateCACertificateOutput, error) {
    var output iot.UpdateCACertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateCertificateResult struct {
	Result workflow.Future
}

func (r *IotUpdateCertificateResult) Get(ctx workflow.Context) (*iot.UpdateCertificateOutput, error) {
    var output iot.UpdateCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateDimensionResult struct {
	Result workflow.Future
}

func (r *IotUpdateDimensionResult) Get(ctx workflow.Context) (*iot.UpdateDimensionOutput, error) {
    var output iot.UpdateDimensionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateDomainConfigurationResult struct {
	Result workflow.Future
}

func (r *IotUpdateDomainConfigurationResult) Get(ctx workflow.Context) (*iot.UpdateDomainConfigurationOutput, error) {
    var output iot.UpdateDomainConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateDynamicThingGroupResult struct {
	Result workflow.Future
}

func (r *IotUpdateDynamicThingGroupResult) Get(ctx workflow.Context) (*iot.UpdateDynamicThingGroupOutput, error) {
    var output iot.UpdateDynamicThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateEventConfigurationsResult struct {
	Result workflow.Future
}

func (r *IotUpdateEventConfigurationsResult) Get(ctx workflow.Context) (*iot.UpdateEventConfigurationsOutput, error) {
    var output iot.UpdateEventConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateIndexingConfigurationResult struct {
	Result workflow.Future
}

func (r *IotUpdateIndexingConfigurationResult) Get(ctx workflow.Context) (*iot.UpdateIndexingConfigurationOutput, error) {
    var output iot.UpdateIndexingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateJobResult struct {
	Result workflow.Future
}

func (r *IotUpdateJobResult) Get(ctx workflow.Context) (*iot.UpdateJobOutput, error) {
    var output iot.UpdateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateMitigationActionResult struct {
	Result workflow.Future
}

func (r *IotUpdateMitigationActionResult) Get(ctx workflow.Context) (*iot.UpdateMitigationActionOutput, error) {
    var output iot.UpdateMitigationActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateProvisioningTemplateResult struct {
	Result workflow.Future
}

func (r *IotUpdateProvisioningTemplateResult) Get(ctx workflow.Context) (*iot.UpdateProvisioningTemplateOutput, error) {
    var output iot.UpdateProvisioningTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateRoleAliasResult struct {
	Result workflow.Future
}

func (r *IotUpdateRoleAliasResult) Get(ctx workflow.Context) (*iot.UpdateRoleAliasOutput, error) {
    var output iot.UpdateRoleAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateScheduledAuditResult struct {
	Result workflow.Future
}

func (r *IotUpdateScheduledAuditResult) Get(ctx workflow.Context) (*iot.UpdateScheduledAuditOutput, error) {
    var output iot.UpdateScheduledAuditOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateSecurityProfileResult struct {
	Result workflow.Future
}

func (r *IotUpdateSecurityProfileResult) Get(ctx workflow.Context) (*iot.UpdateSecurityProfileOutput, error) {
    var output iot.UpdateSecurityProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateStreamResult struct {
	Result workflow.Future
}

func (r *IotUpdateStreamResult) Get(ctx workflow.Context) (*iot.UpdateStreamOutput, error) {
    var output iot.UpdateStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateThingResult struct {
	Result workflow.Future
}

func (r *IotUpdateThingResult) Get(ctx workflow.Context) (*iot.UpdateThingOutput, error) {
    var output iot.UpdateThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateThingGroupResult struct {
	Result workflow.Future
}

func (r *IotUpdateThingGroupResult) Get(ctx workflow.Context) (*iot.UpdateThingGroupOutput, error) {
    var output iot.UpdateThingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateThingGroupsForThingResult struct {
	Result workflow.Future
}

func (r *IotUpdateThingGroupsForThingResult) Get(ctx workflow.Context) (*iot.UpdateThingGroupsForThingOutput, error) {
    var output iot.UpdateThingGroupsForThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotUpdateTopicRuleDestinationResult struct {
	Result workflow.Future
}

func (r *IotUpdateTopicRuleDestinationResult) Get(ctx workflow.Context) (*iot.UpdateTopicRuleDestinationOutput, error) {
    var output iot.UpdateTopicRuleDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotValidateSecurityProfileBehaviorsResult struct {
	Result workflow.Future
}

func (r *IotValidateSecurityProfileBehaviorsResult) Get(ctx workflow.Context) (*iot.ValidateSecurityProfileBehaviorsOutput, error) {
    var output iot.ValidateSecurityProfileBehaviorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IoTStub struct {
    activities awsactivities.IoTActivities
}

func NewIoTStub() IoTClient {
    return &IoTStub{}
}

func (a *IoTStub) AcceptCertificateTransfer(ctx workflow.Context, input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error) {
    var output iot.AcceptCertificateTransferOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptCertificateTransfer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AcceptCertificateTransferAsync(ctx workflow.Context, input *iot.AcceptCertificateTransferInput) *IotAcceptCertificateTransferResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptCertificateTransfer, input)
    return &IotAcceptCertificateTransferResult{Result: future}
}

func (a *IoTStub) AddThingToBillingGroup(ctx workflow.Context, input *iot.AddThingToBillingGroupInput) (*iot.AddThingToBillingGroupOutput, error) {
    var output iot.AddThingToBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddThingToBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AddThingToBillingGroupAsync(ctx workflow.Context, input *iot.AddThingToBillingGroupInput) *IotAddThingToBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddThingToBillingGroup, input)
    return &IotAddThingToBillingGroupResult{Result: future}
}

func (a *IoTStub) AddThingToThingGroup(ctx workflow.Context, input *iot.AddThingToThingGroupInput) (*iot.AddThingToThingGroupOutput, error) {
    var output iot.AddThingToThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddThingToThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AddThingToThingGroupAsync(ctx workflow.Context, input *iot.AddThingToThingGroupInput) *IotAddThingToThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddThingToThingGroup, input)
    return &IotAddThingToThingGroupResult{Result: future}
}

func (a *IoTStub) AssociateTargetsWithJob(ctx workflow.Context, input *iot.AssociateTargetsWithJobInput) (*iot.AssociateTargetsWithJobOutput, error) {
    var output iot.AssociateTargetsWithJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateTargetsWithJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AssociateTargetsWithJobAsync(ctx workflow.Context, input *iot.AssociateTargetsWithJobInput) *IotAssociateTargetsWithJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateTargetsWithJob, input)
    return &IotAssociateTargetsWithJobResult{Result: future}
}

func (a *IoTStub) AttachPolicy(ctx workflow.Context, input *iot.AttachPolicyInput) (*iot.AttachPolicyOutput, error) {
    var output iot.AttachPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AttachPolicyAsync(ctx workflow.Context, input *iot.AttachPolicyInput) *IotAttachPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachPolicy, input)
    return &IotAttachPolicyResult{Result: future}
}

func (a *IoTStub) AttachPrincipalPolicy(ctx workflow.Context, input *iot.AttachPrincipalPolicyInput) (*iot.AttachPrincipalPolicyOutput, error) {
    var output iot.AttachPrincipalPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachPrincipalPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AttachPrincipalPolicyAsync(ctx workflow.Context, input *iot.AttachPrincipalPolicyInput) *IotAttachPrincipalPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachPrincipalPolicy, input)
    return &IotAttachPrincipalPolicyResult{Result: future}
}

func (a *IoTStub) AttachSecurityProfile(ctx workflow.Context, input *iot.AttachSecurityProfileInput) (*iot.AttachSecurityProfileOutput, error) {
    var output iot.AttachSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AttachSecurityProfileAsync(ctx workflow.Context, input *iot.AttachSecurityProfileInput) *IotAttachSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachSecurityProfile, input)
    return &IotAttachSecurityProfileResult{Result: future}
}

func (a *IoTStub) AttachThingPrincipal(ctx workflow.Context, input *iot.AttachThingPrincipalInput) (*iot.AttachThingPrincipalOutput, error) {
    var output iot.AttachThingPrincipalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachThingPrincipal, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) AttachThingPrincipalAsync(ctx workflow.Context, input *iot.AttachThingPrincipalInput) *IotAttachThingPrincipalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachThingPrincipal, input)
    return &IotAttachThingPrincipalResult{Result: future}
}

func (a *IoTStub) CancelAuditMitigationActionsTask(ctx workflow.Context, input *iot.CancelAuditMitigationActionsTaskInput) (*iot.CancelAuditMitigationActionsTaskOutput, error) {
    var output iot.CancelAuditMitigationActionsTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelAuditMitigationActionsTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CancelAuditMitigationActionsTaskAsync(ctx workflow.Context, input *iot.CancelAuditMitigationActionsTaskInput) *IotCancelAuditMitigationActionsTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelAuditMitigationActionsTask, input)
    return &IotCancelAuditMitigationActionsTaskResult{Result: future}
}

func (a *IoTStub) CancelAuditTask(ctx workflow.Context, input *iot.CancelAuditTaskInput) (*iot.CancelAuditTaskOutput, error) {
    var output iot.CancelAuditTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelAuditTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CancelAuditTaskAsync(ctx workflow.Context, input *iot.CancelAuditTaskInput) *IotCancelAuditTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelAuditTask, input)
    return &IotCancelAuditTaskResult{Result: future}
}

func (a *IoTStub) CancelCertificateTransfer(ctx workflow.Context, input *iot.CancelCertificateTransferInput) (*iot.CancelCertificateTransferOutput, error) {
    var output iot.CancelCertificateTransferOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelCertificateTransfer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CancelCertificateTransferAsync(ctx workflow.Context, input *iot.CancelCertificateTransferInput) *IotCancelCertificateTransferResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelCertificateTransfer, input)
    return &IotCancelCertificateTransferResult{Result: future}
}

func (a *IoTStub) CancelJob(ctx workflow.Context, input *iot.CancelJobInput) (*iot.CancelJobOutput, error) {
    var output iot.CancelJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CancelJobAsync(ctx workflow.Context, input *iot.CancelJobInput) *IotCancelJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input)
    return &IotCancelJobResult{Result: future}
}

func (a *IoTStub) CancelJobExecution(ctx workflow.Context, input *iot.CancelJobExecutionInput) (*iot.CancelJobExecutionOutput, error) {
    var output iot.CancelJobExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelJobExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CancelJobExecutionAsync(ctx workflow.Context, input *iot.CancelJobExecutionInput) *IotCancelJobExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelJobExecution, input)
    return &IotCancelJobExecutionResult{Result: future}
}

func (a *IoTStub) ClearDefaultAuthorizer(ctx workflow.Context, input *iot.ClearDefaultAuthorizerInput) (*iot.ClearDefaultAuthorizerOutput, error) {
    var output iot.ClearDefaultAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ClearDefaultAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ClearDefaultAuthorizerAsync(ctx workflow.Context, input *iot.ClearDefaultAuthorizerInput) *IotClearDefaultAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ClearDefaultAuthorizer, input)
    return &IotClearDefaultAuthorizerResult{Result: future}
}

func (a *IoTStub) ConfirmTopicRuleDestination(ctx workflow.Context, input *iot.ConfirmTopicRuleDestinationInput) (*iot.ConfirmTopicRuleDestinationOutput, error) {
    var output iot.ConfirmTopicRuleDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmTopicRuleDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ConfirmTopicRuleDestinationAsync(ctx workflow.Context, input *iot.ConfirmTopicRuleDestinationInput) *IotConfirmTopicRuleDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfirmTopicRuleDestination, input)
    return &IotConfirmTopicRuleDestinationResult{Result: future}
}

func (a *IoTStub) CreateAuditSuppression(ctx workflow.Context, input *iot.CreateAuditSuppressionInput) (*iot.CreateAuditSuppressionOutput, error) {
    var output iot.CreateAuditSuppressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAuditSuppression, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateAuditSuppressionAsync(ctx workflow.Context, input *iot.CreateAuditSuppressionInput) *IotCreateAuditSuppressionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAuditSuppression, input)
    return &IotCreateAuditSuppressionResult{Result: future}
}

func (a *IoTStub) CreateAuthorizer(ctx workflow.Context, input *iot.CreateAuthorizerInput) (*iot.CreateAuthorizerOutput, error) {
    var output iot.CreateAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateAuthorizerAsync(ctx workflow.Context, input *iot.CreateAuthorizerInput) *IotCreateAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAuthorizer, input)
    return &IotCreateAuthorizerResult{Result: future}
}

func (a *IoTStub) CreateBillingGroup(ctx workflow.Context, input *iot.CreateBillingGroupInput) (*iot.CreateBillingGroupOutput, error) {
    var output iot.CreateBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateBillingGroupAsync(ctx workflow.Context, input *iot.CreateBillingGroupInput) *IotCreateBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBillingGroup, input)
    return &IotCreateBillingGroupResult{Result: future}
}

func (a *IoTStub) CreateCertificateFromCsr(ctx workflow.Context, input *iot.CreateCertificateFromCsrInput) (*iot.CreateCertificateFromCsrOutput, error) {
    var output iot.CreateCertificateFromCsrOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCertificateFromCsr, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateCertificateFromCsrAsync(ctx workflow.Context, input *iot.CreateCertificateFromCsrInput) *IotCreateCertificateFromCsrResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCertificateFromCsr, input)
    return &IotCreateCertificateFromCsrResult{Result: future}
}

func (a *IoTStub) CreateDimension(ctx workflow.Context, input *iot.CreateDimensionInput) (*iot.CreateDimensionOutput, error) {
    var output iot.CreateDimensionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDimension, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateDimensionAsync(ctx workflow.Context, input *iot.CreateDimensionInput) *IotCreateDimensionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDimension, input)
    return &IotCreateDimensionResult{Result: future}
}

func (a *IoTStub) CreateDomainConfiguration(ctx workflow.Context, input *iot.CreateDomainConfigurationInput) (*iot.CreateDomainConfigurationOutput, error) {
    var output iot.CreateDomainConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDomainConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateDomainConfigurationAsync(ctx workflow.Context, input *iot.CreateDomainConfigurationInput) *IotCreateDomainConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDomainConfiguration, input)
    return &IotCreateDomainConfigurationResult{Result: future}
}

func (a *IoTStub) CreateDynamicThingGroup(ctx workflow.Context, input *iot.CreateDynamicThingGroupInput) (*iot.CreateDynamicThingGroupOutput, error) {
    var output iot.CreateDynamicThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDynamicThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateDynamicThingGroupAsync(ctx workflow.Context, input *iot.CreateDynamicThingGroupInput) *IotCreateDynamicThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDynamicThingGroup, input)
    return &IotCreateDynamicThingGroupResult{Result: future}
}

func (a *IoTStub) CreateJob(ctx workflow.Context, input *iot.CreateJobInput) (*iot.CreateJobOutput, error) {
    var output iot.CreateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateJobAsync(ctx workflow.Context, input *iot.CreateJobInput) *IotCreateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
    return &IotCreateJobResult{Result: future}
}

func (a *IoTStub) CreateKeysAndCertificate(ctx workflow.Context, input *iot.CreateKeysAndCertificateInput) (*iot.CreateKeysAndCertificateOutput, error) {
    var output iot.CreateKeysAndCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateKeysAndCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateKeysAndCertificateAsync(ctx workflow.Context, input *iot.CreateKeysAndCertificateInput) *IotCreateKeysAndCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateKeysAndCertificate, input)
    return &IotCreateKeysAndCertificateResult{Result: future}
}

func (a *IoTStub) CreateMitigationAction(ctx workflow.Context, input *iot.CreateMitigationActionInput) (*iot.CreateMitigationActionOutput, error) {
    var output iot.CreateMitigationActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMitigationAction, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateMitigationActionAsync(ctx workflow.Context, input *iot.CreateMitigationActionInput) *IotCreateMitigationActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMitigationAction, input)
    return &IotCreateMitigationActionResult{Result: future}
}

func (a *IoTStub) CreateOTAUpdate(ctx workflow.Context, input *iot.CreateOTAUpdateInput) (*iot.CreateOTAUpdateOutput, error) {
    var output iot.CreateOTAUpdateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOTAUpdate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateOTAUpdateAsync(ctx workflow.Context, input *iot.CreateOTAUpdateInput) *IotCreateOTAUpdateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateOTAUpdate, input)
    return &IotCreateOTAUpdateResult{Result: future}
}

func (a *IoTStub) CreatePolicy(ctx workflow.Context, input *iot.CreatePolicyInput) (*iot.CreatePolicyOutput, error) {
    var output iot.CreatePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreatePolicyAsync(ctx workflow.Context, input *iot.CreatePolicyInput) *IotCreatePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePolicy, input)
    return &IotCreatePolicyResult{Result: future}
}

func (a *IoTStub) CreatePolicyVersion(ctx workflow.Context, input *iot.CreatePolicyVersionInput) (*iot.CreatePolicyVersionOutput, error) {
    var output iot.CreatePolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreatePolicyVersionAsync(ctx workflow.Context, input *iot.CreatePolicyVersionInput) *IotCreatePolicyVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePolicyVersion, input)
    return &IotCreatePolicyVersionResult{Result: future}
}

func (a *IoTStub) CreateProvisioningClaim(ctx workflow.Context, input *iot.CreateProvisioningClaimInput) (*iot.CreateProvisioningClaimOutput, error) {
    var output iot.CreateProvisioningClaimOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningClaim, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateProvisioningClaimAsync(ctx workflow.Context, input *iot.CreateProvisioningClaimInput) *IotCreateProvisioningClaimResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningClaim, input)
    return &IotCreateProvisioningClaimResult{Result: future}
}

func (a *IoTStub) CreateProvisioningTemplate(ctx workflow.Context, input *iot.CreateProvisioningTemplateInput) (*iot.CreateProvisioningTemplateOutput, error) {
    var output iot.CreateProvisioningTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateProvisioningTemplateAsync(ctx workflow.Context, input *iot.CreateProvisioningTemplateInput) *IotCreateProvisioningTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningTemplate, input)
    return &IotCreateProvisioningTemplateResult{Result: future}
}

func (a *IoTStub) CreateProvisioningTemplateVersion(ctx workflow.Context, input *iot.CreateProvisioningTemplateVersionInput) (*iot.CreateProvisioningTemplateVersionOutput, error) {
    var output iot.CreateProvisioningTemplateVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningTemplateVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateProvisioningTemplateVersionAsync(ctx workflow.Context, input *iot.CreateProvisioningTemplateVersionInput) *IotCreateProvisioningTemplateVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningTemplateVersion, input)
    return &IotCreateProvisioningTemplateVersionResult{Result: future}
}

func (a *IoTStub) CreateRoleAlias(ctx workflow.Context, input *iot.CreateRoleAliasInput) (*iot.CreateRoleAliasOutput, error) {
    var output iot.CreateRoleAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRoleAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateRoleAliasAsync(ctx workflow.Context, input *iot.CreateRoleAliasInput) *IotCreateRoleAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRoleAlias, input)
    return &IotCreateRoleAliasResult{Result: future}
}

func (a *IoTStub) CreateScheduledAudit(ctx workflow.Context, input *iot.CreateScheduledAuditInput) (*iot.CreateScheduledAuditOutput, error) {
    var output iot.CreateScheduledAuditOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateScheduledAudit, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateScheduledAuditAsync(ctx workflow.Context, input *iot.CreateScheduledAuditInput) *IotCreateScheduledAuditResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateScheduledAudit, input)
    return &IotCreateScheduledAuditResult{Result: future}
}

func (a *IoTStub) CreateSecurityProfile(ctx workflow.Context, input *iot.CreateSecurityProfileInput) (*iot.CreateSecurityProfileOutput, error) {
    var output iot.CreateSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateSecurityProfileAsync(ctx workflow.Context, input *iot.CreateSecurityProfileInput) *IotCreateSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityProfile, input)
    return &IotCreateSecurityProfileResult{Result: future}
}

func (a *IoTStub) CreateStream(ctx workflow.Context, input *iot.CreateStreamInput) (*iot.CreateStreamOutput, error) {
    var output iot.CreateStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStream, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateStreamAsync(ctx workflow.Context, input *iot.CreateStreamInput) *IotCreateStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStream, input)
    return &IotCreateStreamResult{Result: future}
}

func (a *IoTStub) CreateThing(ctx workflow.Context, input *iot.CreateThingInput) (*iot.CreateThingOutput, error) {
    var output iot.CreateThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateThingAsync(ctx workflow.Context, input *iot.CreateThingInput) *IotCreateThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateThing, input)
    return &IotCreateThingResult{Result: future}
}

func (a *IoTStub) CreateThingGroup(ctx workflow.Context, input *iot.CreateThingGroupInput) (*iot.CreateThingGroupOutput, error) {
    var output iot.CreateThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateThingGroupAsync(ctx workflow.Context, input *iot.CreateThingGroupInput) *IotCreateThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateThingGroup, input)
    return &IotCreateThingGroupResult{Result: future}
}

func (a *IoTStub) CreateThingType(ctx workflow.Context, input *iot.CreateThingTypeInput) (*iot.CreateThingTypeOutput, error) {
    var output iot.CreateThingTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateThingType, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateThingTypeAsync(ctx workflow.Context, input *iot.CreateThingTypeInput) *IotCreateThingTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateThingType, input)
    return &IotCreateThingTypeResult{Result: future}
}

func (a *IoTStub) CreateTopicRule(ctx workflow.Context, input *iot.CreateTopicRuleInput) (*iot.CreateTopicRuleOutput, error) {
    var output iot.CreateTopicRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTopicRule, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateTopicRuleAsync(ctx workflow.Context, input *iot.CreateTopicRuleInput) *IotCreateTopicRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTopicRule, input)
    return &IotCreateTopicRuleResult{Result: future}
}

func (a *IoTStub) CreateTopicRuleDestination(ctx workflow.Context, input *iot.CreateTopicRuleDestinationInput) (*iot.CreateTopicRuleDestinationOutput, error) {
    var output iot.CreateTopicRuleDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTopicRuleDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) CreateTopicRuleDestinationAsync(ctx workflow.Context, input *iot.CreateTopicRuleDestinationInput) *IotCreateTopicRuleDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTopicRuleDestination, input)
    return &IotCreateTopicRuleDestinationResult{Result: future}
}

func (a *IoTStub) DeleteAccountAuditConfiguration(ctx workflow.Context, input *iot.DeleteAccountAuditConfigurationInput) (*iot.DeleteAccountAuditConfigurationOutput, error) {
    var output iot.DeleteAccountAuditConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountAuditConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteAccountAuditConfigurationAsync(ctx workflow.Context, input *iot.DeleteAccountAuditConfigurationInput) *IotDeleteAccountAuditConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountAuditConfiguration, input)
    return &IotDeleteAccountAuditConfigurationResult{Result: future}
}

func (a *IoTStub) DeleteAuditSuppression(ctx workflow.Context, input *iot.DeleteAuditSuppressionInput) (*iot.DeleteAuditSuppressionOutput, error) {
    var output iot.DeleteAuditSuppressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAuditSuppression, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteAuditSuppressionAsync(ctx workflow.Context, input *iot.DeleteAuditSuppressionInput) *IotDeleteAuditSuppressionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAuditSuppression, input)
    return &IotDeleteAuditSuppressionResult{Result: future}
}

func (a *IoTStub) DeleteAuthorizer(ctx workflow.Context, input *iot.DeleteAuthorizerInput) (*iot.DeleteAuthorizerOutput, error) {
    var output iot.DeleteAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteAuthorizerAsync(ctx workflow.Context, input *iot.DeleteAuthorizerInput) *IotDeleteAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAuthorizer, input)
    return &IotDeleteAuthorizerResult{Result: future}
}

func (a *IoTStub) DeleteBillingGroup(ctx workflow.Context, input *iot.DeleteBillingGroupInput) (*iot.DeleteBillingGroupOutput, error) {
    var output iot.DeleteBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteBillingGroupAsync(ctx workflow.Context, input *iot.DeleteBillingGroupInput) *IotDeleteBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBillingGroup, input)
    return &IotDeleteBillingGroupResult{Result: future}
}

func (a *IoTStub) DeleteCACertificate(ctx workflow.Context, input *iot.DeleteCACertificateInput) (*iot.DeleteCACertificateOutput, error) {
    var output iot.DeleteCACertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCACertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteCACertificateAsync(ctx workflow.Context, input *iot.DeleteCACertificateInput) *IotDeleteCACertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCACertificate, input)
    return &IotDeleteCACertificateResult{Result: future}
}

func (a *IoTStub) DeleteCertificate(ctx workflow.Context, input *iot.DeleteCertificateInput) (*iot.DeleteCertificateOutput, error) {
    var output iot.DeleteCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteCertificateAsync(ctx workflow.Context, input *iot.DeleteCertificateInput) *IotDeleteCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input)
    return &IotDeleteCertificateResult{Result: future}
}

func (a *IoTStub) DeleteDimension(ctx workflow.Context, input *iot.DeleteDimensionInput) (*iot.DeleteDimensionOutput, error) {
    var output iot.DeleteDimensionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDimension, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteDimensionAsync(ctx workflow.Context, input *iot.DeleteDimensionInput) *IotDeleteDimensionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDimension, input)
    return &IotDeleteDimensionResult{Result: future}
}

func (a *IoTStub) DeleteDomainConfiguration(ctx workflow.Context, input *iot.DeleteDomainConfigurationInput) (*iot.DeleteDomainConfigurationOutput, error) {
    var output iot.DeleteDomainConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteDomainConfigurationAsync(ctx workflow.Context, input *iot.DeleteDomainConfigurationInput) *IotDeleteDomainConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainConfiguration, input)
    return &IotDeleteDomainConfigurationResult{Result: future}
}

func (a *IoTStub) DeleteDynamicThingGroup(ctx workflow.Context, input *iot.DeleteDynamicThingGroupInput) (*iot.DeleteDynamicThingGroupOutput, error) {
    var output iot.DeleteDynamicThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDynamicThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteDynamicThingGroupAsync(ctx workflow.Context, input *iot.DeleteDynamicThingGroupInput) *IotDeleteDynamicThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDynamicThingGroup, input)
    return &IotDeleteDynamicThingGroupResult{Result: future}
}

func (a *IoTStub) DeleteJob(ctx workflow.Context, input *iot.DeleteJobInput) (*iot.DeleteJobOutput, error) {
    var output iot.DeleteJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteJobAsync(ctx workflow.Context, input *iot.DeleteJobInput) *IotDeleteJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJob, input)
    return &IotDeleteJobResult{Result: future}
}

func (a *IoTStub) DeleteJobExecution(ctx workflow.Context, input *iot.DeleteJobExecutionInput) (*iot.DeleteJobExecutionOutput, error) {
    var output iot.DeleteJobExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJobExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteJobExecutionAsync(ctx workflow.Context, input *iot.DeleteJobExecutionInput) *IotDeleteJobExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJobExecution, input)
    return &IotDeleteJobExecutionResult{Result: future}
}

func (a *IoTStub) DeleteMitigationAction(ctx workflow.Context, input *iot.DeleteMitigationActionInput) (*iot.DeleteMitigationActionOutput, error) {
    var output iot.DeleteMitigationActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMitigationAction, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteMitigationActionAsync(ctx workflow.Context, input *iot.DeleteMitigationActionInput) *IotDeleteMitigationActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMitigationAction, input)
    return &IotDeleteMitigationActionResult{Result: future}
}

func (a *IoTStub) DeleteOTAUpdate(ctx workflow.Context, input *iot.DeleteOTAUpdateInput) (*iot.DeleteOTAUpdateOutput, error) {
    var output iot.DeleteOTAUpdateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOTAUpdate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteOTAUpdateAsync(ctx workflow.Context, input *iot.DeleteOTAUpdateInput) *IotDeleteOTAUpdateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteOTAUpdate, input)
    return &IotDeleteOTAUpdateResult{Result: future}
}

func (a *IoTStub) DeletePolicy(ctx workflow.Context, input *iot.DeletePolicyInput) (*iot.DeletePolicyOutput, error) {
    var output iot.DeletePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeletePolicyAsync(ctx workflow.Context, input *iot.DeletePolicyInput) *IotDeletePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input)
    return &IotDeletePolicyResult{Result: future}
}

func (a *IoTStub) DeletePolicyVersion(ctx workflow.Context, input *iot.DeletePolicyVersionInput) (*iot.DeletePolicyVersionOutput, error) {
    var output iot.DeletePolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeletePolicyVersionAsync(ctx workflow.Context, input *iot.DeletePolicyVersionInput) *IotDeletePolicyVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePolicyVersion, input)
    return &IotDeletePolicyVersionResult{Result: future}
}

func (a *IoTStub) DeleteProvisioningTemplate(ctx workflow.Context, input *iot.DeleteProvisioningTemplateInput) (*iot.DeleteProvisioningTemplateOutput, error) {
    var output iot.DeleteProvisioningTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisioningTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteProvisioningTemplateAsync(ctx workflow.Context, input *iot.DeleteProvisioningTemplateInput) *IotDeleteProvisioningTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisioningTemplate, input)
    return &IotDeleteProvisioningTemplateResult{Result: future}
}

func (a *IoTStub) DeleteProvisioningTemplateVersion(ctx workflow.Context, input *iot.DeleteProvisioningTemplateVersionInput) (*iot.DeleteProvisioningTemplateVersionOutput, error) {
    var output iot.DeleteProvisioningTemplateVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisioningTemplateVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteProvisioningTemplateVersionAsync(ctx workflow.Context, input *iot.DeleteProvisioningTemplateVersionInput) *IotDeleteProvisioningTemplateVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisioningTemplateVersion, input)
    return &IotDeleteProvisioningTemplateVersionResult{Result: future}
}

func (a *IoTStub) DeleteRegistrationCode(ctx workflow.Context, input *iot.DeleteRegistrationCodeInput) (*iot.DeleteRegistrationCodeOutput, error) {
    var output iot.DeleteRegistrationCodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegistrationCode, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteRegistrationCodeAsync(ctx workflow.Context, input *iot.DeleteRegistrationCodeInput) *IotDeleteRegistrationCodeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRegistrationCode, input)
    return &IotDeleteRegistrationCodeResult{Result: future}
}

func (a *IoTStub) DeleteRoleAlias(ctx workflow.Context, input *iot.DeleteRoleAliasInput) (*iot.DeleteRoleAliasOutput, error) {
    var output iot.DeleteRoleAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoleAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteRoleAliasAsync(ctx workflow.Context, input *iot.DeleteRoleAliasInput) *IotDeleteRoleAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRoleAlias, input)
    return &IotDeleteRoleAliasResult{Result: future}
}

func (a *IoTStub) DeleteScheduledAudit(ctx workflow.Context, input *iot.DeleteScheduledAuditInput) (*iot.DeleteScheduledAuditOutput, error) {
    var output iot.DeleteScheduledAuditOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteScheduledAudit, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteScheduledAuditAsync(ctx workflow.Context, input *iot.DeleteScheduledAuditInput) *IotDeleteScheduledAuditResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteScheduledAudit, input)
    return &IotDeleteScheduledAuditResult{Result: future}
}

func (a *IoTStub) DeleteSecurityProfile(ctx workflow.Context, input *iot.DeleteSecurityProfileInput) (*iot.DeleteSecurityProfileOutput, error) {
    var output iot.DeleteSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteSecurityProfileAsync(ctx workflow.Context, input *iot.DeleteSecurityProfileInput) *IotDeleteSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityProfile, input)
    return &IotDeleteSecurityProfileResult{Result: future}
}

func (a *IoTStub) DeleteStream(ctx workflow.Context, input *iot.DeleteStreamInput) (*iot.DeleteStreamOutput, error) {
    var output iot.DeleteStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStream, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteStreamAsync(ctx workflow.Context, input *iot.DeleteStreamInput) *IotDeleteStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteStream, input)
    return &IotDeleteStreamResult{Result: future}
}

func (a *IoTStub) DeleteThing(ctx workflow.Context, input *iot.DeleteThingInput) (*iot.DeleteThingOutput, error) {
    var output iot.DeleteThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteThingAsync(ctx workflow.Context, input *iot.DeleteThingInput) *IotDeleteThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteThing, input)
    return &IotDeleteThingResult{Result: future}
}

func (a *IoTStub) DeleteThingGroup(ctx workflow.Context, input *iot.DeleteThingGroupInput) (*iot.DeleteThingGroupOutput, error) {
    var output iot.DeleteThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteThingGroupAsync(ctx workflow.Context, input *iot.DeleteThingGroupInput) *IotDeleteThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteThingGroup, input)
    return &IotDeleteThingGroupResult{Result: future}
}

func (a *IoTStub) DeleteThingType(ctx workflow.Context, input *iot.DeleteThingTypeInput) (*iot.DeleteThingTypeOutput, error) {
    var output iot.DeleteThingTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteThingType, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteThingTypeAsync(ctx workflow.Context, input *iot.DeleteThingTypeInput) *IotDeleteThingTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteThingType, input)
    return &IotDeleteThingTypeResult{Result: future}
}

func (a *IoTStub) DeleteTopicRule(ctx workflow.Context, input *iot.DeleteTopicRuleInput) (*iot.DeleteTopicRuleOutput, error) {
    var output iot.DeleteTopicRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTopicRule, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteTopicRuleAsync(ctx workflow.Context, input *iot.DeleteTopicRuleInput) *IotDeleteTopicRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTopicRule, input)
    return &IotDeleteTopicRuleResult{Result: future}
}

func (a *IoTStub) DeleteTopicRuleDestination(ctx workflow.Context, input *iot.DeleteTopicRuleDestinationInput) (*iot.DeleteTopicRuleDestinationOutput, error) {
    var output iot.DeleteTopicRuleDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTopicRuleDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteTopicRuleDestinationAsync(ctx workflow.Context, input *iot.DeleteTopicRuleDestinationInput) *IotDeleteTopicRuleDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTopicRuleDestination, input)
    return &IotDeleteTopicRuleDestinationResult{Result: future}
}

func (a *IoTStub) DeleteV2LoggingLevel(ctx workflow.Context, input *iot.DeleteV2LoggingLevelInput) (*iot.DeleteV2LoggingLevelOutput, error) {
    var output iot.DeleteV2LoggingLevelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteV2LoggingLevel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeleteV2LoggingLevelAsync(ctx workflow.Context, input *iot.DeleteV2LoggingLevelInput) *IotDeleteV2LoggingLevelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteV2LoggingLevel, input)
    return &IotDeleteV2LoggingLevelResult{Result: future}
}

func (a *IoTStub) DeprecateThingType(ctx workflow.Context, input *iot.DeprecateThingTypeInput) (*iot.DeprecateThingTypeOutput, error) {
    var output iot.DeprecateThingTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprecateThingType, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DeprecateThingTypeAsync(ctx workflow.Context, input *iot.DeprecateThingTypeInput) *IotDeprecateThingTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeprecateThingType, input)
    return &IotDeprecateThingTypeResult{Result: future}
}

func (a *IoTStub) DescribeAccountAuditConfiguration(ctx workflow.Context, input *iot.DescribeAccountAuditConfigurationInput) (*iot.DescribeAccountAuditConfigurationOutput, error) {
    var output iot.DescribeAccountAuditConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAuditConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeAccountAuditConfigurationAsync(ctx workflow.Context, input *iot.DescribeAccountAuditConfigurationInput) *IotDescribeAccountAuditConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAuditConfiguration, input)
    return &IotDescribeAccountAuditConfigurationResult{Result: future}
}

func (a *IoTStub) DescribeAuditFinding(ctx workflow.Context, input *iot.DescribeAuditFindingInput) (*iot.DescribeAuditFindingOutput, error) {
    var output iot.DescribeAuditFindingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditFinding, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeAuditFindingAsync(ctx workflow.Context, input *iot.DescribeAuditFindingInput) *IotDescribeAuditFindingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditFinding, input)
    return &IotDescribeAuditFindingResult{Result: future}
}

func (a *IoTStub) DescribeAuditMitigationActionsTask(ctx workflow.Context, input *iot.DescribeAuditMitigationActionsTaskInput) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
    var output iot.DescribeAuditMitigationActionsTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditMitigationActionsTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeAuditMitigationActionsTaskAsync(ctx workflow.Context, input *iot.DescribeAuditMitigationActionsTaskInput) *IotDescribeAuditMitigationActionsTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditMitigationActionsTask, input)
    return &IotDescribeAuditMitigationActionsTaskResult{Result: future}
}

func (a *IoTStub) DescribeAuditSuppression(ctx workflow.Context, input *iot.DescribeAuditSuppressionInput) (*iot.DescribeAuditSuppressionOutput, error) {
    var output iot.DescribeAuditSuppressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditSuppression, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeAuditSuppressionAsync(ctx workflow.Context, input *iot.DescribeAuditSuppressionInput) *IotDescribeAuditSuppressionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditSuppression, input)
    return &IotDescribeAuditSuppressionResult{Result: future}
}

func (a *IoTStub) DescribeAuditTask(ctx workflow.Context, input *iot.DescribeAuditTaskInput) (*iot.DescribeAuditTaskOutput, error) {
    var output iot.DescribeAuditTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeAuditTaskAsync(ctx workflow.Context, input *iot.DescribeAuditTaskInput) *IotDescribeAuditTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditTask, input)
    return &IotDescribeAuditTaskResult{Result: future}
}

func (a *IoTStub) DescribeAuthorizer(ctx workflow.Context, input *iot.DescribeAuthorizerInput) (*iot.DescribeAuthorizerOutput, error) {
    var output iot.DescribeAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeAuthorizerAsync(ctx workflow.Context, input *iot.DescribeAuthorizerInput) *IotDescribeAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAuthorizer, input)
    return &IotDescribeAuthorizerResult{Result: future}
}

func (a *IoTStub) DescribeBillingGroup(ctx workflow.Context, input *iot.DescribeBillingGroupInput) (*iot.DescribeBillingGroupOutput, error) {
    var output iot.DescribeBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeBillingGroupAsync(ctx workflow.Context, input *iot.DescribeBillingGroupInput) *IotDescribeBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBillingGroup, input)
    return &IotDescribeBillingGroupResult{Result: future}
}

func (a *IoTStub) DescribeCACertificate(ctx workflow.Context, input *iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error) {
    var output iot.DescribeCACertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCACertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeCACertificateAsync(ctx workflow.Context, input *iot.DescribeCACertificateInput) *IotDescribeCACertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCACertificate, input)
    return &IotDescribeCACertificateResult{Result: future}
}

func (a *IoTStub) DescribeCertificate(ctx workflow.Context, input *iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error) {
    var output iot.DescribeCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeCertificateAsync(ctx workflow.Context, input *iot.DescribeCertificateInput) *IotDescribeCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificate, input)
    return &IotDescribeCertificateResult{Result: future}
}

func (a *IoTStub) DescribeDefaultAuthorizer(ctx workflow.Context, input *iot.DescribeDefaultAuthorizerInput) (*iot.DescribeDefaultAuthorizerOutput, error) {
    var output iot.DescribeDefaultAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDefaultAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeDefaultAuthorizerAsync(ctx workflow.Context, input *iot.DescribeDefaultAuthorizerInput) *IotDescribeDefaultAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDefaultAuthorizer, input)
    return &IotDescribeDefaultAuthorizerResult{Result: future}
}

func (a *IoTStub) DescribeDimension(ctx workflow.Context, input *iot.DescribeDimensionInput) (*iot.DescribeDimensionOutput, error) {
    var output iot.DescribeDimensionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDimension, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeDimensionAsync(ctx workflow.Context, input *iot.DescribeDimensionInput) *IotDescribeDimensionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDimension, input)
    return &IotDescribeDimensionResult{Result: future}
}

func (a *IoTStub) DescribeDomainConfiguration(ctx workflow.Context, input *iot.DescribeDomainConfigurationInput) (*iot.DescribeDomainConfigurationOutput, error) {
    var output iot.DescribeDomainConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomainConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeDomainConfigurationAsync(ctx workflow.Context, input *iot.DescribeDomainConfigurationInput) *IotDescribeDomainConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDomainConfiguration, input)
    return &IotDescribeDomainConfigurationResult{Result: future}
}

func (a *IoTStub) DescribeEndpoint(ctx workflow.Context, input *iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error) {
    var output iot.DescribeEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeEndpointAsync(ctx workflow.Context, input *iot.DescribeEndpointInput) *IotDescribeEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoint, input)
    return &IotDescribeEndpointResult{Result: future}
}

func (a *IoTStub) DescribeEventConfigurations(ctx workflow.Context, input *iot.DescribeEventConfigurationsInput) (*iot.DescribeEventConfigurationsOutput, error) {
    var output iot.DescribeEventConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeEventConfigurationsAsync(ctx workflow.Context, input *iot.DescribeEventConfigurationsInput) *IotDescribeEventConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventConfigurations, input)
    return &IotDescribeEventConfigurationsResult{Result: future}
}

func (a *IoTStub) DescribeIndex(ctx workflow.Context, input *iot.DescribeIndexInput) (*iot.DescribeIndexOutput, error) {
    var output iot.DescribeIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeIndexAsync(ctx workflow.Context, input *iot.DescribeIndexInput) *IotDescribeIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIndex, input)
    return &IotDescribeIndexResult{Result: future}
}

func (a *IoTStub) DescribeJob(ctx workflow.Context, input *iot.DescribeJobInput) (*iot.DescribeJobOutput, error) {
    var output iot.DescribeJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeJobAsync(ctx workflow.Context, input *iot.DescribeJobInput) *IotDescribeJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input)
    return &IotDescribeJobResult{Result: future}
}

func (a *IoTStub) DescribeJobExecution(ctx workflow.Context, input *iot.DescribeJobExecutionInput) (*iot.DescribeJobExecutionOutput, error) {
    var output iot.DescribeJobExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJobExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeJobExecutionAsync(ctx workflow.Context, input *iot.DescribeJobExecutionInput) *IotDescribeJobExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJobExecution, input)
    return &IotDescribeJobExecutionResult{Result: future}
}

func (a *IoTStub) DescribeMitigationAction(ctx workflow.Context, input *iot.DescribeMitigationActionInput) (*iot.DescribeMitigationActionOutput, error) {
    var output iot.DescribeMitigationActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMitigationAction, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeMitigationActionAsync(ctx workflow.Context, input *iot.DescribeMitigationActionInput) *IotDescribeMitigationActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMitigationAction, input)
    return &IotDescribeMitigationActionResult{Result: future}
}

func (a *IoTStub) DescribeProvisioningTemplate(ctx workflow.Context, input *iot.DescribeProvisioningTemplateInput) (*iot.DescribeProvisioningTemplateOutput, error) {
    var output iot.DescribeProvisioningTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeProvisioningTemplateAsync(ctx workflow.Context, input *iot.DescribeProvisioningTemplateInput) *IotDescribeProvisioningTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningTemplate, input)
    return &IotDescribeProvisioningTemplateResult{Result: future}
}

func (a *IoTStub) DescribeProvisioningTemplateVersion(ctx workflow.Context, input *iot.DescribeProvisioningTemplateVersionInput) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
    var output iot.DescribeProvisioningTemplateVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningTemplateVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeProvisioningTemplateVersionAsync(ctx workflow.Context, input *iot.DescribeProvisioningTemplateVersionInput) *IotDescribeProvisioningTemplateVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningTemplateVersion, input)
    return &IotDescribeProvisioningTemplateVersionResult{Result: future}
}

func (a *IoTStub) DescribeRoleAlias(ctx workflow.Context, input *iot.DescribeRoleAliasInput) (*iot.DescribeRoleAliasOutput, error) {
    var output iot.DescribeRoleAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRoleAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeRoleAliasAsync(ctx workflow.Context, input *iot.DescribeRoleAliasInput) *IotDescribeRoleAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRoleAlias, input)
    return &IotDescribeRoleAliasResult{Result: future}
}

func (a *IoTStub) DescribeScheduledAudit(ctx workflow.Context, input *iot.DescribeScheduledAuditInput) (*iot.DescribeScheduledAuditOutput, error) {
    var output iot.DescribeScheduledAuditOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledAudit, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeScheduledAuditAsync(ctx workflow.Context, input *iot.DescribeScheduledAuditInput) *IotDescribeScheduledAuditResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledAudit, input)
    return &IotDescribeScheduledAuditResult{Result: future}
}

func (a *IoTStub) DescribeSecurityProfile(ctx workflow.Context, input *iot.DescribeSecurityProfileInput) (*iot.DescribeSecurityProfileOutput, error) {
    var output iot.DescribeSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeSecurityProfileAsync(ctx workflow.Context, input *iot.DescribeSecurityProfileInput) *IotDescribeSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityProfile, input)
    return &IotDescribeSecurityProfileResult{Result: future}
}

func (a *IoTStub) DescribeStream(ctx workflow.Context, input *iot.DescribeStreamInput) (*iot.DescribeStreamOutput, error) {
    var output iot.DescribeStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeStreamAsync(ctx workflow.Context, input *iot.DescribeStreamInput) *IotDescribeStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input)
    return &IotDescribeStreamResult{Result: future}
}

func (a *IoTStub) DescribeThing(ctx workflow.Context, input *iot.DescribeThingInput) (*iot.DescribeThingOutput, error) {
    var output iot.DescribeThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeThingAsync(ctx workflow.Context, input *iot.DescribeThingInput) *IotDescribeThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeThing, input)
    return &IotDescribeThingResult{Result: future}
}

func (a *IoTStub) DescribeThingGroup(ctx workflow.Context, input *iot.DescribeThingGroupInput) (*iot.DescribeThingGroupOutput, error) {
    var output iot.DescribeThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeThingGroupAsync(ctx workflow.Context, input *iot.DescribeThingGroupInput) *IotDescribeThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeThingGroup, input)
    return &IotDescribeThingGroupResult{Result: future}
}

func (a *IoTStub) DescribeThingRegistrationTask(ctx workflow.Context, input *iot.DescribeThingRegistrationTaskInput) (*iot.DescribeThingRegistrationTaskOutput, error) {
    var output iot.DescribeThingRegistrationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeThingRegistrationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeThingRegistrationTaskAsync(ctx workflow.Context, input *iot.DescribeThingRegistrationTaskInput) *IotDescribeThingRegistrationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeThingRegistrationTask, input)
    return &IotDescribeThingRegistrationTaskResult{Result: future}
}

func (a *IoTStub) DescribeThingType(ctx workflow.Context, input *iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error) {
    var output iot.DescribeThingTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeThingType, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DescribeThingTypeAsync(ctx workflow.Context, input *iot.DescribeThingTypeInput) *IotDescribeThingTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeThingType, input)
    return &IotDescribeThingTypeResult{Result: future}
}

func (a *IoTStub) DetachPolicy(ctx workflow.Context, input *iot.DetachPolicyInput) (*iot.DetachPolicyOutput, error) {
    var output iot.DetachPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DetachPolicyAsync(ctx workflow.Context, input *iot.DetachPolicyInput) *IotDetachPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachPolicy, input)
    return &IotDetachPolicyResult{Result: future}
}

func (a *IoTStub) DetachPrincipalPolicy(ctx workflow.Context, input *iot.DetachPrincipalPolicyInput) (*iot.DetachPrincipalPolicyOutput, error) {
    var output iot.DetachPrincipalPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachPrincipalPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DetachPrincipalPolicyAsync(ctx workflow.Context, input *iot.DetachPrincipalPolicyInput) *IotDetachPrincipalPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachPrincipalPolicy, input)
    return &IotDetachPrincipalPolicyResult{Result: future}
}

func (a *IoTStub) DetachSecurityProfile(ctx workflow.Context, input *iot.DetachSecurityProfileInput) (*iot.DetachSecurityProfileOutput, error) {
    var output iot.DetachSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DetachSecurityProfileAsync(ctx workflow.Context, input *iot.DetachSecurityProfileInput) *IotDetachSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachSecurityProfile, input)
    return &IotDetachSecurityProfileResult{Result: future}
}

func (a *IoTStub) DetachThingPrincipal(ctx workflow.Context, input *iot.DetachThingPrincipalInput) (*iot.DetachThingPrincipalOutput, error) {
    var output iot.DetachThingPrincipalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachThingPrincipal, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DetachThingPrincipalAsync(ctx workflow.Context, input *iot.DetachThingPrincipalInput) *IotDetachThingPrincipalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachThingPrincipal, input)
    return &IotDetachThingPrincipalResult{Result: future}
}

func (a *IoTStub) DisableTopicRule(ctx workflow.Context, input *iot.DisableTopicRuleInput) (*iot.DisableTopicRuleOutput, error) {
    var output iot.DisableTopicRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableTopicRule, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) DisableTopicRuleAsync(ctx workflow.Context, input *iot.DisableTopicRuleInput) *IotDisableTopicRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableTopicRule, input)
    return &IotDisableTopicRuleResult{Result: future}
}

func (a *IoTStub) EnableTopicRule(ctx workflow.Context, input *iot.EnableTopicRuleInput) (*iot.EnableTopicRuleOutput, error) {
    var output iot.EnableTopicRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableTopicRule, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) EnableTopicRuleAsync(ctx workflow.Context, input *iot.EnableTopicRuleInput) *IotEnableTopicRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableTopicRule, input)
    return &IotEnableTopicRuleResult{Result: future}
}

func (a *IoTStub) GetCardinality(ctx workflow.Context, input *iot.GetCardinalityInput) (*iot.GetCardinalityOutput, error) {
    var output iot.GetCardinalityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCardinality, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetCardinalityAsync(ctx workflow.Context, input *iot.GetCardinalityInput) *IotGetCardinalityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCardinality, input)
    return &IotGetCardinalityResult{Result: future}
}

func (a *IoTStub) GetEffectivePolicies(ctx workflow.Context, input *iot.GetEffectivePoliciesInput) (*iot.GetEffectivePoliciesOutput, error) {
    var output iot.GetEffectivePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEffectivePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetEffectivePoliciesAsync(ctx workflow.Context, input *iot.GetEffectivePoliciesInput) *IotGetEffectivePoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEffectivePolicies, input)
    return &IotGetEffectivePoliciesResult{Result: future}
}

func (a *IoTStub) GetIndexingConfiguration(ctx workflow.Context, input *iot.GetIndexingConfigurationInput) (*iot.GetIndexingConfigurationOutput, error) {
    var output iot.GetIndexingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIndexingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetIndexingConfigurationAsync(ctx workflow.Context, input *iot.GetIndexingConfigurationInput) *IotGetIndexingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIndexingConfiguration, input)
    return &IotGetIndexingConfigurationResult{Result: future}
}

func (a *IoTStub) GetJobDocument(ctx workflow.Context, input *iot.GetJobDocumentInput) (*iot.GetJobDocumentOutput, error) {
    var output iot.GetJobDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetJobDocumentAsync(ctx workflow.Context, input *iot.GetJobDocumentInput) *IotGetJobDocumentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobDocument, input)
    return &IotGetJobDocumentResult{Result: future}
}

func (a *IoTStub) GetLoggingOptions(ctx workflow.Context, input *iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error) {
    var output iot.GetLoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLoggingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetLoggingOptionsAsync(ctx workflow.Context, input *iot.GetLoggingOptionsInput) *IotGetLoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLoggingOptions, input)
    return &IotGetLoggingOptionsResult{Result: future}
}

func (a *IoTStub) GetOTAUpdate(ctx workflow.Context, input *iot.GetOTAUpdateInput) (*iot.GetOTAUpdateOutput, error) {
    var output iot.GetOTAUpdateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOTAUpdate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetOTAUpdateAsync(ctx workflow.Context, input *iot.GetOTAUpdateInput) *IotGetOTAUpdateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOTAUpdate, input)
    return &IotGetOTAUpdateResult{Result: future}
}

func (a *IoTStub) GetPercentiles(ctx workflow.Context, input *iot.GetPercentilesInput) (*iot.GetPercentilesOutput, error) {
    var output iot.GetPercentilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPercentiles, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetPercentilesAsync(ctx workflow.Context, input *iot.GetPercentilesInput) *IotGetPercentilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPercentiles, input)
    return &IotGetPercentilesResult{Result: future}
}

func (a *IoTStub) GetPolicy(ctx workflow.Context, input *iot.GetPolicyInput) (*iot.GetPolicyOutput, error) {
    var output iot.GetPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetPolicyAsync(ctx workflow.Context, input *iot.GetPolicyInput) *IotGetPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input)
    return &IotGetPolicyResult{Result: future}
}

func (a *IoTStub) GetPolicyVersion(ctx workflow.Context, input *iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error) {
    var output iot.GetPolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetPolicyVersionAsync(ctx workflow.Context, input *iot.GetPolicyVersionInput) *IotGetPolicyVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPolicyVersion, input)
    return &IotGetPolicyVersionResult{Result: future}
}

func (a *IoTStub) GetRegistrationCode(ctx workflow.Context, input *iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error) {
    var output iot.GetRegistrationCodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRegistrationCode, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetRegistrationCodeAsync(ctx workflow.Context, input *iot.GetRegistrationCodeInput) *IotGetRegistrationCodeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRegistrationCode, input)
    return &IotGetRegistrationCodeResult{Result: future}
}

func (a *IoTStub) GetStatistics(ctx workflow.Context, input *iot.GetStatisticsInput) (*iot.GetStatisticsOutput, error) {
    var output iot.GetStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetStatisticsAsync(ctx workflow.Context, input *iot.GetStatisticsInput) *IotGetStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetStatistics, input)
    return &IotGetStatisticsResult{Result: future}
}

func (a *IoTStub) GetTopicRule(ctx workflow.Context, input *iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error) {
    var output iot.GetTopicRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTopicRule, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetTopicRuleAsync(ctx workflow.Context, input *iot.GetTopicRuleInput) *IotGetTopicRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTopicRule, input)
    return &IotGetTopicRuleResult{Result: future}
}

func (a *IoTStub) GetTopicRuleDestination(ctx workflow.Context, input *iot.GetTopicRuleDestinationInput) (*iot.GetTopicRuleDestinationOutput, error) {
    var output iot.GetTopicRuleDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTopicRuleDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetTopicRuleDestinationAsync(ctx workflow.Context, input *iot.GetTopicRuleDestinationInput) *IotGetTopicRuleDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTopicRuleDestination, input)
    return &IotGetTopicRuleDestinationResult{Result: future}
}

func (a *IoTStub) GetV2LoggingOptions(ctx workflow.Context, input *iot.GetV2LoggingOptionsInput) (*iot.GetV2LoggingOptionsOutput, error) {
    var output iot.GetV2LoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetV2LoggingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) GetV2LoggingOptionsAsync(ctx workflow.Context, input *iot.GetV2LoggingOptionsInput) *IotGetV2LoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetV2LoggingOptions, input)
    return &IotGetV2LoggingOptionsResult{Result: future}
}

func (a *IoTStub) ListActiveViolations(ctx workflow.Context, input *iot.ListActiveViolationsInput) (*iot.ListActiveViolationsOutput, error) {
    var output iot.ListActiveViolationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListActiveViolations, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListActiveViolationsAsync(ctx workflow.Context, input *iot.ListActiveViolationsInput) *IotListActiveViolationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListActiveViolations, input)
    return &IotListActiveViolationsResult{Result: future}
}

func (a *IoTStub) ListAttachedPolicies(ctx workflow.Context, input *iot.ListAttachedPoliciesInput) (*iot.ListAttachedPoliciesOutput, error) {
    var output iot.ListAttachedPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttachedPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAttachedPoliciesAsync(ctx workflow.Context, input *iot.ListAttachedPoliciesInput) *IotListAttachedPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAttachedPolicies, input)
    return &IotListAttachedPoliciesResult{Result: future}
}

func (a *IoTStub) ListAuditFindings(ctx workflow.Context, input *iot.ListAuditFindingsInput) (*iot.ListAuditFindingsOutput, error) {
    var output iot.ListAuditFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAuditFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAuditFindingsAsync(ctx workflow.Context, input *iot.ListAuditFindingsInput) *IotListAuditFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAuditFindings, input)
    return &IotListAuditFindingsResult{Result: future}
}

func (a *IoTStub) ListAuditMitigationActionsExecutions(ctx workflow.Context, input *iot.ListAuditMitigationActionsExecutionsInput) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
    var output iot.ListAuditMitigationActionsExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAuditMitigationActionsExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAuditMitigationActionsExecutionsAsync(ctx workflow.Context, input *iot.ListAuditMitigationActionsExecutionsInput) *IotListAuditMitigationActionsExecutionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAuditMitigationActionsExecutions, input)
    return &IotListAuditMitigationActionsExecutionsResult{Result: future}
}

func (a *IoTStub) ListAuditMitigationActionsTasks(ctx workflow.Context, input *iot.ListAuditMitigationActionsTasksInput) (*iot.ListAuditMitigationActionsTasksOutput, error) {
    var output iot.ListAuditMitigationActionsTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAuditMitigationActionsTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAuditMitigationActionsTasksAsync(ctx workflow.Context, input *iot.ListAuditMitigationActionsTasksInput) *IotListAuditMitigationActionsTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAuditMitigationActionsTasks, input)
    return &IotListAuditMitigationActionsTasksResult{Result: future}
}

func (a *IoTStub) ListAuditSuppressions(ctx workflow.Context, input *iot.ListAuditSuppressionsInput) (*iot.ListAuditSuppressionsOutput, error) {
    var output iot.ListAuditSuppressionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAuditSuppressions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAuditSuppressionsAsync(ctx workflow.Context, input *iot.ListAuditSuppressionsInput) *IotListAuditSuppressionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAuditSuppressions, input)
    return &IotListAuditSuppressionsResult{Result: future}
}

func (a *IoTStub) ListAuditTasks(ctx workflow.Context, input *iot.ListAuditTasksInput) (*iot.ListAuditTasksOutput, error) {
    var output iot.ListAuditTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAuditTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAuditTasksAsync(ctx workflow.Context, input *iot.ListAuditTasksInput) *IotListAuditTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAuditTasks, input)
    return &IotListAuditTasksResult{Result: future}
}

func (a *IoTStub) ListAuthorizers(ctx workflow.Context, input *iot.ListAuthorizersInput) (*iot.ListAuthorizersOutput, error) {
    var output iot.ListAuthorizersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAuthorizers, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListAuthorizersAsync(ctx workflow.Context, input *iot.ListAuthorizersInput) *IotListAuthorizersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAuthorizers, input)
    return &IotListAuthorizersResult{Result: future}
}

func (a *IoTStub) ListBillingGroups(ctx workflow.Context, input *iot.ListBillingGroupsInput) (*iot.ListBillingGroupsOutput, error) {
    var output iot.ListBillingGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBillingGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListBillingGroupsAsync(ctx workflow.Context, input *iot.ListBillingGroupsInput) *IotListBillingGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBillingGroups, input)
    return &IotListBillingGroupsResult{Result: future}
}

func (a *IoTStub) ListCACertificates(ctx workflow.Context, input *iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error) {
    var output iot.ListCACertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCACertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListCACertificatesAsync(ctx workflow.Context, input *iot.ListCACertificatesInput) *IotListCACertificatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCACertificates, input)
    return &IotListCACertificatesResult{Result: future}
}

func (a *IoTStub) ListCertificates(ctx workflow.Context, input *iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error) {
    var output iot.ListCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListCertificatesAsync(ctx workflow.Context, input *iot.ListCertificatesInput) *IotListCertificatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCertificates, input)
    return &IotListCertificatesResult{Result: future}
}

func (a *IoTStub) ListCertificatesByCA(ctx workflow.Context, input *iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error) {
    var output iot.ListCertificatesByCAOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCertificatesByCA, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListCertificatesByCAAsync(ctx workflow.Context, input *iot.ListCertificatesByCAInput) *IotListCertificatesByCAResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCertificatesByCA, input)
    return &IotListCertificatesByCAResult{Result: future}
}

func (a *IoTStub) ListDimensions(ctx workflow.Context, input *iot.ListDimensionsInput) (*iot.ListDimensionsOutput, error) {
    var output iot.ListDimensionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDimensions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListDimensionsAsync(ctx workflow.Context, input *iot.ListDimensionsInput) *IotListDimensionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDimensions, input)
    return &IotListDimensionsResult{Result: future}
}

func (a *IoTStub) ListDomainConfigurations(ctx workflow.Context, input *iot.ListDomainConfigurationsInput) (*iot.ListDomainConfigurationsOutput, error) {
    var output iot.ListDomainConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListDomainConfigurationsAsync(ctx workflow.Context, input *iot.ListDomainConfigurationsInput) *IotListDomainConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDomainConfigurations, input)
    return &IotListDomainConfigurationsResult{Result: future}
}

func (a *IoTStub) ListIndices(ctx workflow.Context, input *iot.ListIndicesInput) (*iot.ListIndicesOutput, error) {
    var output iot.ListIndicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIndices, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListIndicesAsync(ctx workflow.Context, input *iot.ListIndicesInput) *IotListIndicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListIndices, input)
    return &IotListIndicesResult{Result: future}
}

func (a *IoTStub) ListJobExecutionsForJob(ctx workflow.Context, input *iot.ListJobExecutionsForJobInput) (*iot.ListJobExecutionsForJobOutput, error) {
    var output iot.ListJobExecutionsForJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobExecutionsForJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListJobExecutionsForJobAsync(ctx workflow.Context, input *iot.ListJobExecutionsForJobInput) *IotListJobExecutionsForJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobExecutionsForJob, input)
    return &IotListJobExecutionsForJobResult{Result: future}
}

func (a *IoTStub) ListJobExecutionsForThing(ctx workflow.Context, input *iot.ListJobExecutionsForThingInput) (*iot.ListJobExecutionsForThingOutput, error) {
    var output iot.ListJobExecutionsForThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobExecutionsForThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListJobExecutionsForThingAsync(ctx workflow.Context, input *iot.ListJobExecutionsForThingInput) *IotListJobExecutionsForThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobExecutionsForThing, input)
    return &IotListJobExecutionsForThingResult{Result: future}
}

func (a *IoTStub) ListJobs(ctx workflow.Context, input *iot.ListJobsInput) (*iot.ListJobsOutput, error) {
    var output iot.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListJobsAsync(ctx workflow.Context, input *iot.ListJobsInput) *IotListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &IotListJobsResult{Result: future}
}

func (a *IoTStub) ListMitigationActions(ctx workflow.Context, input *iot.ListMitigationActionsInput) (*iot.ListMitigationActionsOutput, error) {
    var output iot.ListMitigationActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMitigationActions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListMitigationActionsAsync(ctx workflow.Context, input *iot.ListMitigationActionsInput) *IotListMitigationActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMitigationActions, input)
    return &IotListMitigationActionsResult{Result: future}
}

func (a *IoTStub) ListOTAUpdates(ctx workflow.Context, input *iot.ListOTAUpdatesInput) (*iot.ListOTAUpdatesOutput, error) {
    var output iot.ListOTAUpdatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOTAUpdates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListOTAUpdatesAsync(ctx workflow.Context, input *iot.ListOTAUpdatesInput) *IotListOTAUpdatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOTAUpdates, input)
    return &IotListOTAUpdatesResult{Result: future}
}

func (a *IoTStub) ListOutgoingCertificates(ctx workflow.Context, input *iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error) {
    var output iot.ListOutgoingCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOutgoingCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListOutgoingCertificatesAsync(ctx workflow.Context, input *iot.ListOutgoingCertificatesInput) *IotListOutgoingCertificatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOutgoingCertificates, input)
    return &IotListOutgoingCertificatesResult{Result: future}
}

func (a *IoTStub) ListPolicies(ctx workflow.Context, input *iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error) {
    var output iot.ListPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListPoliciesAsync(ctx workflow.Context, input *iot.ListPoliciesInput) *IotListPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input)
    return &IotListPoliciesResult{Result: future}
}

func (a *IoTStub) ListPolicyPrincipals(ctx workflow.Context, input *iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error) {
    var output iot.ListPolicyPrincipalsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicyPrincipals, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListPolicyPrincipalsAsync(ctx workflow.Context, input *iot.ListPolicyPrincipalsInput) *IotListPolicyPrincipalsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPolicyPrincipals, input)
    return &IotListPolicyPrincipalsResult{Result: future}
}

func (a *IoTStub) ListPolicyVersions(ctx workflow.Context, input *iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error) {
    var output iot.ListPolicyVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicyVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListPolicyVersionsAsync(ctx workflow.Context, input *iot.ListPolicyVersionsInput) *IotListPolicyVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPolicyVersions, input)
    return &IotListPolicyVersionsResult{Result: future}
}

func (a *IoTStub) ListPrincipalPolicies(ctx workflow.Context, input *iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error) {
    var output iot.ListPrincipalPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPrincipalPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListPrincipalPoliciesAsync(ctx workflow.Context, input *iot.ListPrincipalPoliciesInput) *IotListPrincipalPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPrincipalPolicies, input)
    return &IotListPrincipalPoliciesResult{Result: future}
}

func (a *IoTStub) ListPrincipalThings(ctx workflow.Context, input *iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error) {
    var output iot.ListPrincipalThingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPrincipalThings, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListPrincipalThingsAsync(ctx workflow.Context, input *iot.ListPrincipalThingsInput) *IotListPrincipalThingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPrincipalThings, input)
    return &IotListPrincipalThingsResult{Result: future}
}

func (a *IoTStub) ListProvisioningTemplateVersions(ctx workflow.Context, input *iot.ListProvisioningTemplateVersionsInput) (*iot.ListProvisioningTemplateVersionsOutput, error) {
    var output iot.ListProvisioningTemplateVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningTemplateVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListProvisioningTemplateVersionsAsync(ctx workflow.Context, input *iot.ListProvisioningTemplateVersionsInput) *IotListProvisioningTemplateVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningTemplateVersions, input)
    return &IotListProvisioningTemplateVersionsResult{Result: future}
}

func (a *IoTStub) ListProvisioningTemplates(ctx workflow.Context, input *iot.ListProvisioningTemplatesInput) (*iot.ListProvisioningTemplatesOutput, error) {
    var output iot.ListProvisioningTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListProvisioningTemplatesAsync(ctx workflow.Context, input *iot.ListProvisioningTemplatesInput) *IotListProvisioningTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningTemplates, input)
    return &IotListProvisioningTemplatesResult{Result: future}
}

func (a *IoTStub) ListRoleAliases(ctx workflow.Context, input *iot.ListRoleAliasesInput) (*iot.ListRoleAliasesOutput, error) {
    var output iot.ListRoleAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRoleAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListRoleAliasesAsync(ctx workflow.Context, input *iot.ListRoleAliasesInput) *IotListRoleAliasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRoleAliases, input)
    return &IotListRoleAliasesResult{Result: future}
}

func (a *IoTStub) ListScheduledAudits(ctx workflow.Context, input *iot.ListScheduledAuditsInput) (*iot.ListScheduledAuditsOutput, error) {
    var output iot.ListScheduledAuditsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListScheduledAudits, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListScheduledAuditsAsync(ctx workflow.Context, input *iot.ListScheduledAuditsInput) *IotListScheduledAuditsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListScheduledAudits, input)
    return &IotListScheduledAuditsResult{Result: future}
}

func (a *IoTStub) ListSecurityProfiles(ctx workflow.Context, input *iot.ListSecurityProfilesInput) (*iot.ListSecurityProfilesOutput, error) {
    var output iot.ListSecurityProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecurityProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListSecurityProfilesAsync(ctx workflow.Context, input *iot.ListSecurityProfilesInput) *IotListSecurityProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSecurityProfiles, input)
    return &IotListSecurityProfilesResult{Result: future}
}

func (a *IoTStub) ListSecurityProfilesForTarget(ctx workflow.Context, input *iot.ListSecurityProfilesForTargetInput) (*iot.ListSecurityProfilesForTargetOutput, error) {
    var output iot.ListSecurityProfilesForTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecurityProfilesForTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListSecurityProfilesForTargetAsync(ctx workflow.Context, input *iot.ListSecurityProfilesForTargetInput) *IotListSecurityProfilesForTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSecurityProfilesForTarget, input)
    return &IotListSecurityProfilesForTargetResult{Result: future}
}

func (a *IoTStub) ListStreams(ctx workflow.Context, input *iot.ListStreamsInput) (*iot.ListStreamsOutput, error) {
    var output iot.ListStreamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListStreamsAsync(ctx workflow.Context, input *iot.ListStreamsInput) *IotListStreamsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input)
    return &IotListStreamsResult{Result: future}
}

func (a *IoTStub) ListTagsForResource(ctx workflow.Context, input *iot.ListTagsForResourceInput) (*iot.ListTagsForResourceOutput, error) {
    var output iot.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListTagsForResourceAsync(ctx workflow.Context, input *iot.ListTagsForResourceInput) *IotListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &IotListTagsForResourceResult{Result: future}
}

func (a *IoTStub) ListTargetsForPolicy(ctx workflow.Context, input *iot.ListTargetsForPolicyInput) (*iot.ListTargetsForPolicyOutput, error) {
    var output iot.ListTargetsForPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTargetsForPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListTargetsForPolicyAsync(ctx workflow.Context, input *iot.ListTargetsForPolicyInput) *IotListTargetsForPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTargetsForPolicy, input)
    return &IotListTargetsForPolicyResult{Result: future}
}

func (a *IoTStub) ListTargetsForSecurityProfile(ctx workflow.Context, input *iot.ListTargetsForSecurityProfileInput) (*iot.ListTargetsForSecurityProfileOutput, error) {
    var output iot.ListTargetsForSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTargetsForSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListTargetsForSecurityProfileAsync(ctx workflow.Context, input *iot.ListTargetsForSecurityProfileInput) *IotListTargetsForSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTargetsForSecurityProfile, input)
    return &IotListTargetsForSecurityProfileResult{Result: future}
}

func (a *IoTStub) ListThingGroups(ctx workflow.Context, input *iot.ListThingGroupsInput) (*iot.ListThingGroupsOutput, error) {
    var output iot.ListThingGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingGroupsAsync(ctx workflow.Context, input *iot.ListThingGroupsInput) *IotListThingGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingGroups, input)
    return &IotListThingGroupsResult{Result: future}
}

func (a *IoTStub) ListThingGroupsForThing(ctx workflow.Context, input *iot.ListThingGroupsForThingInput) (*iot.ListThingGroupsForThingOutput, error) {
    var output iot.ListThingGroupsForThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingGroupsForThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingGroupsForThingAsync(ctx workflow.Context, input *iot.ListThingGroupsForThingInput) *IotListThingGroupsForThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingGroupsForThing, input)
    return &IotListThingGroupsForThingResult{Result: future}
}

func (a *IoTStub) ListThingPrincipals(ctx workflow.Context, input *iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error) {
    var output iot.ListThingPrincipalsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingPrincipals, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingPrincipalsAsync(ctx workflow.Context, input *iot.ListThingPrincipalsInput) *IotListThingPrincipalsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingPrincipals, input)
    return &IotListThingPrincipalsResult{Result: future}
}

func (a *IoTStub) ListThingRegistrationTaskReports(ctx workflow.Context, input *iot.ListThingRegistrationTaskReportsInput) (*iot.ListThingRegistrationTaskReportsOutput, error) {
    var output iot.ListThingRegistrationTaskReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingRegistrationTaskReports, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingRegistrationTaskReportsAsync(ctx workflow.Context, input *iot.ListThingRegistrationTaskReportsInput) *IotListThingRegistrationTaskReportsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingRegistrationTaskReports, input)
    return &IotListThingRegistrationTaskReportsResult{Result: future}
}

func (a *IoTStub) ListThingRegistrationTasks(ctx workflow.Context, input *iot.ListThingRegistrationTasksInput) (*iot.ListThingRegistrationTasksOutput, error) {
    var output iot.ListThingRegistrationTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingRegistrationTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingRegistrationTasksAsync(ctx workflow.Context, input *iot.ListThingRegistrationTasksInput) *IotListThingRegistrationTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingRegistrationTasks, input)
    return &IotListThingRegistrationTasksResult{Result: future}
}

func (a *IoTStub) ListThingTypes(ctx workflow.Context, input *iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error) {
    var output iot.ListThingTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingTypesAsync(ctx workflow.Context, input *iot.ListThingTypesInput) *IotListThingTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingTypes, input)
    return &IotListThingTypesResult{Result: future}
}

func (a *IoTStub) ListThings(ctx workflow.Context, input *iot.ListThingsInput) (*iot.ListThingsOutput, error) {
    var output iot.ListThingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThings, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingsAsync(ctx workflow.Context, input *iot.ListThingsInput) *IotListThingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThings, input)
    return &IotListThingsResult{Result: future}
}

func (a *IoTStub) ListThingsInBillingGroup(ctx workflow.Context, input *iot.ListThingsInBillingGroupInput) (*iot.ListThingsInBillingGroupOutput, error) {
    var output iot.ListThingsInBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingsInBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingsInBillingGroupAsync(ctx workflow.Context, input *iot.ListThingsInBillingGroupInput) *IotListThingsInBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingsInBillingGroup, input)
    return &IotListThingsInBillingGroupResult{Result: future}
}

func (a *IoTStub) ListThingsInThingGroup(ctx workflow.Context, input *iot.ListThingsInThingGroupInput) (*iot.ListThingsInThingGroupOutput, error) {
    var output iot.ListThingsInThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThingsInThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListThingsInThingGroupAsync(ctx workflow.Context, input *iot.ListThingsInThingGroupInput) *IotListThingsInThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThingsInThingGroup, input)
    return &IotListThingsInThingGroupResult{Result: future}
}

func (a *IoTStub) ListTopicRuleDestinations(ctx workflow.Context, input *iot.ListTopicRuleDestinationsInput) (*iot.ListTopicRuleDestinationsOutput, error) {
    var output iot.ListTopicRuleDestinationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTopicRuleDestinations, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListTopicRuleDestinationsAsync(ctx workflow.Context, input *iot.ListTopicRuleDestinationsInput) *IotListTopicRuleDestinationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTopicRuleDestinations, input)
    return &IotListTopicRuleDestinationsResult{Result: future}
}

func (a *IoTStub) ListTopicRules(ctx workflow.Context, input *iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error) {
    var output iot.ListTopicRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTopicRules, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListTopicRulesAsync(ctx workflow.Context, input *iot.ListTopicRulesInput) *IotListTopicRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTopicRules, input)
    return &IotListTopicRulesResult{Result: future}
}

func (a *IoTStub) ListV2LoggingLevels(ctx workflow.Context, input *iot.ListV2LoggingLevelsInput) (*iot.ListV2LoggingLevelsOutput, error) {
    var output iot.ListV2LoggingLevelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListV2LoggingLevels, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListV2LoggingLevelsAsync(ctx workflow.Context, input *iot.ListV2LoggingLevelsInput) *IotListV2LoggingLevelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListV2LoggingLevels, input)
    return &IotListV2LoggingLevelsResult{Result: future}
}

func (a *IoTStub) ListViolationEvents(ctx workflow.Context, input *iot.ListViolationEventsInput) (*iot.ListViolationEventsOutput, error) {
    var output iot.ListViolationEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListViolationEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ListViolationEventsAsync(ctx workflow.Context, input *iot.ListViolationEventsInput) *IotListViolationEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListViolationEvents, input)
    return &IotListViolationEventsResult{Result: future}
}

func (a *IoTStub) RegisterCACertificate(ctx workflow.Context, input *iot.RegisterCACertificateInput) (*iot.RegisterCACertificateOutput, error) {
    var output iot.RegisterCACertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterCACertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RegisterCACertificateAsync(ctx workflow.Context, input *iot.RegisterCACertificateInput) *IotRegisterCACertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterCACertificate, input)
    return &IotRegisterCACertificateResult{Result: future}
}

func (a *IoTStub) RegisterCertificate(ctx workflow.Context, input *iot.RegisterCertificateInput) (*iot.RegisterCertificateOutput, error) {
    var output iot.RegisterCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RegisterCertificateAsync(ctx workflow.Context, input *iot.RegisterCertificateInput) *IotRegisterCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterCertificate, input)
    return &IotRegisterCertificateResult{Result: future}
}

func (a *IoTStub) RegisterCertificateWithoutCA(ctx workflow.Context, input *iot.RegisterCertificateWithoutCAInput) (*iot.RegisterCertificateWithoutCAOutput, error) {
    var output iot.RegisterCertificateWithoutCAOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterCertificateWithoutCA, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RegisterCertificateWithoutCAAsync(ctx workflow.Context, input *iot.RegisterCertificateWithoutCAInput) *IotRegisterCertificateWithoutCAResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterCertificateWithoutCA, input)
    return &IotRegisterCertificateWithoutCAResult{Result: future}
}

func (a *IoTStub) RegisterThing(ctx workflow.Context, input *iot.RegisterThingInput) (*iot.RegisterThingOutput, error) {
    var output iot.RegisterThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RegisterThingAsync(ctx workflow.Context, input *iot.RegisterThingInput) *IotRegisterThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterThing, input)
    return &IotRegisterThingResult{Result: future}
}

func (a *IoTStub) RejectCertificateTransfer(ctx workflow.Context, input *iot.RejectCertificateTransferInput) (*iot.RejectCertificateTransferOutput, error) {
    var output iot.RejectCertificateTransferOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectCertificateTransfer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RejectCertificateTransferAsync(ctx workflow.Context, input *iot.RejectCertificateTransferInput) *IotRejectCertificateTransferResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectCertificateTransfer, input)
    return &IotRejectCertificateTransferResult{Result: future}
}

func (a *IoTStub) RemoveThingFromBillingGroup(ctx workflow.Context, input *iot.RemoveThingFromBillingGroupInput) (*iot.RemoveThingFromBillingGroupOutput, error) {
    var output iot.RemoveThingFromBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveThingFromBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RemoveThingFromBillingGroupAsync(ctx workflow.Context, input *iot.RemoveThingFromBillingGroupInput) *IotRemoveThingFromBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveThingFromBillingGroup, input)
    return &IotRemoveThingFromBillingGroupResult{Result: future}
}

func (a *IoTStub) RemoveThingFromThingGroup(ctx workflow.Context, input *iot.RemoveThingFromThingGroupInput) (*iot.RemoveThingFromThingGroupOutput, error) {
    var output iot.RemoveThingFromThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveThingFromThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) RemoveThingFromThingGroupAsync(ctx workflow.Context, input *iot.RemoveThingFromThingGroupInput) *IotRemoveThingFromThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveThingFromThingGroup, input)
    return &IotRemoveThingFromThingGroupResult{Result: future}
}

func (a *IoTStub) ReplaceTopicRule(ctx workflow.Context, input *iot.ReplaceTopicRuleInput) (*iot.ReplaceTopicRuleOutput, error) {
    var output iot.ReplaceTopicRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceTopicRule, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ReplaceTopicRuleAsync(ctx workflow.Context, input *iot.ReplaceTopicRuleInput) *IotReplaceTopicRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceTopicRule, input)
    return &IotReplaceTopicRuleResult{Result: future}
}

func (a *IoTStub) SearchIndex(ctx workflow.Context, input *iot.SearchIndexInput) (*iot.SearchIndexOutput, error) {
    var output iot.SearchIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) SearchIndexAsync(ctx workflow.Context, input *iot.SearchIndexInput) *IotSearchIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchIndex, input)
    return &IotSearchIndexResult{Result: future}
}

func (a *IoTStub) SetDefaultAuthorizer(ctx workflow.Context, input *iot.SetDefaultAuthorizerInput) (*iot.SetDefaultAuthorizerOutput, error) {
    var output iot.SetDefaultAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetDefaultAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) SetDefaultAuthorizerAsync(ctx workflow.Context, input *iot.SetDefaultAuthorizerInput) *IotSetDefaultAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetDefaultAuthorizer, input)
    return &IotSetDefaultAuthorizerResult{Result: future}
}

func (a *IoTStub) SetDefaultPolicyVersion(ctx workflow.Context, input *iot.SetDefaultPolicyVersionInput) (*iot.SetDefaultPolicyVersionOutput, error) {
    var output iot.SetDefaultPolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetDefaultPolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) SetDefaultPolicyVersionAsync(ctx workflow.Context, input *iot.SetDefaultPolicyVersionInput) *IotSetDefaultPolicyVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetDefaultPolicyVersion, input)
    return &IotSetDefaultPolicyVersionResult{Result: future}
}

func (a *IoTStub) SetLoggingOptions(ctx workflow.Context, input *iot.SetLoggingOptionsInput) (*iot.SetLoggingOptionsOutput, error) {
    var output iot.SetLoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetLoggingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) SetLoggingOptionsAsync(ctx workflow.Context, input *iot.SetLoggingOptionsInput) *IotSetLoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetLoggingOptions, input)
    return &IotSetLoggingOptionsResult{Result: future}
}

func (a *IoTStub) SetV2LoggingLevel(ctx workflow.Context, input *iot.SetV2LoggingLevelInput) (*iot.SetV2LoggingLevelOutput, error) {
    var output iot.SetV2LoggingLevelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetV2LoggingLevel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) SetV2LoggingLevelAsync(ctx workflow.Context, input *iot.SetV2LoggingLevelInput) *IotSetV2LoggingLevelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetV2LoggingLevel, input)
    return &IotSetV2LoggingLevelResult{Result: future}
}

func (a *IoTStub) SetV2LoggingOptions(ctx workflow.Context, input *iot.SetV2LoggingOptionsInput) (*iot.SetV2LoggingOptionsOutput, error) {
    var output iot.SetV2LoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetV2LoggingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) SetV2LoggingOptionsAsync(ctx workflow.Context, input *iot.SetV2LoggingOptionsInput) *IotSetV2LoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetV2LoggingOptions, input)
    return &IotSetV2LoggingOptionsResult{Result: future}
}

func (a *IoTStub) StartAuditMitigationActionsTask(ctx workflow.Context, input *iot.StartAuditMitigationActionsTaskInput) (*iot.StartAuditMitigationActionsTaskOutput, error) {
    var output iot.StartAuditMitigationActionsTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartAuditMitigationActionsTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) StartAuditMitigationActionsTaskAsync(ctx workflow.Context, input *iot.StartAuditMitigationActionsTaskInput) *IotStartAuditMitigationActionsTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartAuditMitigationActionsTask, input)
    return &IotStartAuditMitigationActionsTaskResult{Result: future}
}

func (a *IoTStub) StartOnDemandAuditTask(ctx workflow.Context, input *iot.StartOnDemandAuditTaskInput) (*iot.StartOnDemandAuditTaskOutput, error) {
    var output iot.StartOnDemandAuditTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartOnDemandAuditTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) StartOnDemandAuditTaskAsync(ctx workflow.Context, input *iot.StartOnDemandAuditTaskInput) *IotStartOnDemandAuditTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartOnDemandAuditTask, input)
    return &IotStartOnDemandAuditTaskResult{Result: future}
}

func (a *IoTStub) StartThingRegistrationTask(ctx workflow.Context, input *iot.StartThingRegistrationTaskInput) (*iot.StartThingRegistrationTaskOutput, error) {
    var output iot.StartThingRegistrationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartThingRegistrationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) StartThingRegistrationTaskAsync(ctx workflow.Context, input *iot.StartThingRegistrationTaskInput) *IotStartThingRegistrationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartThingRegistrationTask, input)
    return &IotStartThingRegistrationTaskResult{Result: future}
}

func (a *IoTStub) StopThingRegistrationTask(ctx workflow.Context, input *iot.StopThingRegistrationTaskInput) (*iot.StopThingRegistrationTaskOutput, error) {
    var output iot.StopThingRegistrationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopThingRegistrationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) StopThingRegistrationTaskAsync(ctx workflow.Context, input *iot.StopThingRegistrationTaskInput) *IotStopThingRegistrationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopThingRegistrationTask, input)
    return &IotStopThingRegistrationTaskResult{Result: future}
}

func (a *IoTStub) TagResource(ctx workflow.Context, input *iot.TagResourceInput) (*iot.TagResourceOutput, error) {
    var output iot.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) TagResourceAsync(ctx workflow.Context, input *iot.TagResourceInput) *IotTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &IotTagResourceResult{Result: future}
}

func (a *IoTStub) TestAuthorization(ctx workflow.Context, input *iot.TestAuthorizationInput) (*iot.TestAuthorizationOutput, error) {
    var output iot.TestAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) TestAuthorizationAsync(ctx workflow.Context, input *iot.TestAuthorizationInput) *IotTestAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestAuthorization, input)
    return &IotTestAuthorizationResult{Result: future}
}

func (a *IoTStub) TestInvokeAuthorizer(ctx workflow.Context, input *iot.TestInvokeAuthorizerInput) (*iot.TestInvokeAuthorizerOutput, error) {
    var output iot.TestInvokeAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestInvokeAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) TestInvokeAuthorizerAsync(ctx workflow.Context, input *iot.TestInvokeAuthorizerInput) *IotTestInvokeAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestInvokeAuthorizer, input)
    return &IotTestInvokeAuthorizerResult{Result: future}
}

func (a *IoTStub) TransferCertificate(ctx workflow.Context, input *iot.TransferCertificateInput) (*iot.TransferCertificateOutput, error) {
    var output iot.TransferCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TransferCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) TransferCertificateAsync(ctx workflow.Context, input *iot.TransferCertificateInput) *IotTransferCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TransferCertificate, input)
    return &IotTransferCertificateResult{Result: future}
}

func (a *IoTStub) UntagResource(ctx workflow.Context, input *iot.UntagResourceInput) (*iot.UntagResourceOutput, error) {
    var output iot.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UntagResourceAsync(ctx workflow.Context, input *iot.UntagResourceInput) *IotUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &IotUntagResourceResult{Result: future}
}

func (a *IoTStub) UpdateAccountAuditConfiguration(ctx workflow.Context, input *iot.UpdateAccountAuditConfigurationInput) (*iot.UpdateAccountAuditConfigurationOutput, error) {
    var output iot.UpdateAccountAuditConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountAuditConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateAccountAuditConfigurationAsync(ctx workflow.Context, input *iot.UpdateAccountAuditConfigurationInput) *IotUpdateAccountAuditConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountAuditConfiguration, input)
    return &IotUpdateAccountAuditConfigurationResult{Result: future}
}

func (a *IoTStub) UpdateAuditSuppression(ctx workflow.Context, input *iot.UpdateAuditSuppressionInput) (*iot.UpdateAuditSuppressionOutput, error) {
    var output iot.UpdateAuditSuppressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAuditSuppression, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateAuditSuppressionAsync(ctx workflow.Context, input *iot.UpdateAuditSuppressionInput) *IotUpdateAuditSuppressionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAuditSuppression, input)
    return &IotUpdateAuditSuppressionResult{Result: future}
}

func (a *IoTStub) UpdateAuthorizer(ctx workflow.Context, input *iot.UpdateAuthorizerInput) (*iot.UpdateAuthorizerOutput, error) {
    var output iot.UpdateAuthorizerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthorizer, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateAuthorizerAsync(ctx workflow.Context, input *iot.UpdateAuthorizerInput) *IotUpdateAuthorizerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthorizer, input)
    return &IotUpdateAuthorizerResult{Result: future}
}

func (a *IoTStub) UpdateBillingGroup(ctx workflow.Context, input *iot.UpdateBillingGroupInput) (*iot.UpdateBillingGroupOutput, error) {
    var output iot.UpdateBillingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBillingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateBillingGroupAsync(ctx workflow.Context, input *iot.UpdateBillingGroupInput) *IotUpdateBillingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBillingGroup, input)
    return &IotUpdateBillingGroupResult{Result: future}
}

func (a *IoTStub) UpdateCACertificate(ctx workflow.Context, input *iot.UpdateCACertificateInput) (*iot.UpdateCACertificateOutput, error) {
    var output iot.UpdateCACertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCACertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateCACertificateAsync(ctx workflow.Context, input *iot.UpdateCACertificateInput) *IotUpdateCACertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCACertificate, input)
    return &IotUpdateCACertificateResult{Result: future}
}

func (a *IoTStub) UpdateCertificate(ctx workflow.Context, input *iot.UpdateCertificateInput) (*iot.UpdateCertificateOutput, error) {
    var output iot.UpdateCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateCertificateAsync(ctx workflow.Context, input *iot.UpdateCertificateInput) *IotUpdateCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCertificate, input)
    return &IotUpdateCertificateResult{Result: future}
}

func (a *IoTStub) UpdateDimension(ctx workflow.Context, input *iot.UpdateDimensionInput) (*iot.UpdateDimensionOutput, error) {
    var output iot.UpdateDimensionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDimension, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateDimensionAsync(ctx workflow.Context, input *iot.UpdateDimensionInput) *IotUpdateDimensionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDimension, input)
    return &IotUpdateDimensionResult{Result: future}
}

func (a *IoTStub) UpdateDomainConfiguration(ctx workflow.Context, input *iot.UpdateDomainConfigurationInput) (*iot.UpdateDomainConfigurationOutput, error) {
    var output iot.UpdateDomainConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateDomainConfigurationAsync(ctx workflow.Context, input *iot.UpdateDomainConfigurationInput) *IotUpdateDomainConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainConfiguration, input)
    return &IotUpdateDomainConfigurationResult{Result: future}
}

func (a *IoTStub) UpdateDynamicThingGroup(ctx workflow.Context, input *iot.UpdateDynamicThingGroupInput) (*iot.UpdateDynamicThingGroupOutput, error) {
    var output iot.UpdateDynamicThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDynamicThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateDynamicThingGroupAsync(ctx workflow.Context, input *iot.UpdateDynamicThingGroupInput) *IotUpdateDynamicThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDynamicThingGroup, input)
    return &IotUpdateDynamicThingGroupResult{Result: future}
}

func (a *IoTStub) UpdateEventConfigurations(ctx workflow.Context, input *iot.UpdateEventConfigurationsInput) (*iot.UpdateEventConfigurationsOutput, error) {
    var output iot.UpdateEventConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEventConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateEventConfigurationsAsync(ctx workflow.Context, input *iot.UpdateEventConfigurationsInput) *IotUpdateEventConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEventConfigurations, input)
    return &IotUpdateEventConfigurationsResult{Result: future}
}

func (a *IoTStub) UpdateIndexingConfiguration(ctx workflow.Context, input *iot.UpdateIndexingConfigurationInput) (*iot.UpdateIndexingConfigurationOutput, error) {
    var output iot.UpdateIndexingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIndexingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateIndexingConfigurationAsync(ctx workflow.Context, input *iot.UpdateIndexingConfigurationInput) *IotUpdateIndexingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateIndexingConfiguration, input)
    return &IotUpdateIndexingConfigurationResult{Result: future}
}

func (a *IoTStub) UpdateJob(ctx workflow.Context, input *iot.UpdateJobInput) (*iot.UpdateJobOutput, error) {
    var output iot.UpdateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateJobAsync(ctx workflow.Context, input *iot.UpdateJobInput) *IotUpdateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJob, input)
    return &IotUpdateJobResult{Result: future}
}

func (a *IoTStub) UpdateMitigationAction(ctx workflow.Context, input *iot.UpdateMitigationActionInput) (*iot.UpdateMitigationActionOutput, error) {
    var output iot.UpdateMitigationActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMitigationAction, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateMitigationActionAsync(ctx workflow.Context, input *iot.UpdateMitigationActionInput) *IotUpdateMitigationActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMitigationAction, input)
    return &IotUpdateMitigationActionResult{Result: future}
}

func (a *IoTStub) UpdateProvisioningTemplate(ctx workflow.Context, input *iot.UpdateProvisioningTemplateInput) (*iot.UpdateProvisioningTemplateOutput, error) {
    var output iot.UpdateProvisioningTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisioningTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateProvisioningTemplateAsync(ctx workflow.Context, input *iot.UpdateProvisioningTemplateInput) *IotUpdateProvisioningTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisioningTemplate, input)
    return &IotUpdateProvisioningTemplateResult{Result: future}
}

func (a *IoTStub) UpdateRoleAlias(ctx workflow.Context, input *iot.UpdateRoleAliasInput) (*iot.UpdateRoleAliasOutput, error) {
    var output iot.UpdateRoleAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRoleAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateRoleAliasAsync(ctx workflow.Context, input *iot.UpdateRoleAliasInput) *IotUpdateRoleAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRoleAlias, input)
    return &IotUpdateRoleAliasResult{Result: future}
}

func (a *IoTStub) UpdateScheduledAudit(ctx workflow.Context, input *iot.UpdateScheduledAuditInput) (*iot.UpdateScheduledAuditOutput, error) {
    var output iot.UpdateScheduledAuditOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateScheduledAudit, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateScheduledAuditAsync(ctx workflow.Context, input *iot.UpdateScheduledAuditInput) *IotUpdateScheduledAuditResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateScheduledAudit, input)
    return &IotUpdateScheduledAuditResult{Result: future}
}

func (a *IoTStub) UpdateSecurityProfile(ctx workflow.Context, input *iot.UpdateSecurityProfileInput) (*iot.UpdateSecurityProfileOutput, error) {
    var output iot.UpdateSecurityProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateSecurityProfileAsync(ctx workflow.Context, input *iot.UpdateSecurityProfileInput) *IotUpdateSecurityProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityProfile, input)
    return &IotUpdateSecurityProfileResult{Result: future}
}

func (a *IoTStub) UpdateStream(ctx workflow.Context, input *iot.UpdateStreamInput) (*iot.UpdateStreamOutput, error) {
    var output iot.UpdateStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStream, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateStreamAsync(ctx workflow.Context, input *iot.UpdateStreamInput) *IotUpdateStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateStream, input)
    return &IotUpdateStreamResult{Result: future}
}

func (a *IoTStub) UpdateThing(ctx workflow.Context, input *iot.UpdateThingInput) (*iot.UpdateThingOutput, error) {
    var output iot.UpdateThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateThingAsync(ctx workflow.Context, input *iot.UpdateThingInput) *IotUpdateThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateThing, input)
    return &IotUpdateThingResult{Result: future}
}

func (a *IoTStub) UpdateThingGroup(ctx workflow.Context, input *iot.UpdateThingGroupInput) (*iot.UpdateThingGroupOutput, error) {
    var output iot.UpdateThingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateThingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateThingGroupAsync(ctx workflow.Context, input *iot.UpdateThingGroupInput) *IotUpdateThingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateThingGroup, input)
    return &IotUpdateThingGroupResult{Result: future}
}

func (a *IoTStub) UpdateThingGroupsForThing(ctx workflow.Context, input *iot.UpdateThingGroupsForThingInput) (*iot.UpdateThingGroupsForThingOutput, error) {
    var output iot.UpdateThingGroupsForThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateThingGroupsForThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateThingGroupsForThingAsync(ctx workflow.Context, input *iot.UpdateThingGroupsForThingInput) *IotUpdateThingGroupsForThingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateThingGroupsForThing, input)
    return &IotUpdateThingGroupsForThingResult{Result: future}
}

func (a *IoTStub) UpdateTopicRuleDestination(ctx workflow.Context, input *iot.UpdateTopicRuleDestinationInput) (*iot.UpdateTopicRuleDestinationOutput, error) {
    var output iot.UpdateTopicRuleDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTopicRuleDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) UpdateTopicRuleDestinationAsync(ctx workflow.Context, input *iot.UpdateTopicRuleDestinationInput) *IotUpdateTopicRuleDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTopicRuleDestination, input)
    return &IotUpdateTopicRuleDestinationResult{Result: future}
}

func (a *IoTStub) ValidateSecurityProfileBehaviors(ctx workflow.Context, input *iot.ValidateSecurityProfileBehaviorsInput) (*iot.ValidateSecurityProfileBehaviorsOutput, error) {
    var output iot.ValidateSecurityProfileBehaviorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidateSecurityProfileBehaviors, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTStub) ValidateSecurityProfileBehaviorsAsync(ctx workflow.Context, input *iot.ValidateSecurityProfileBehaviorsInput) *IotValidateSecurityProfileBehaviorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ValidateSecurityProfileBehaviors, input)
    return &IotValidateSecurityProfileBehaviorsResult{Result: future}
}