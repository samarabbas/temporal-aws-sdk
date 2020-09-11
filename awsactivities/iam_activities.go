
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IAMActivities struct {
    client iamiface.IAMAPI
}

func NewIAMActivities(session *session.Session, config ...*aws.Config) *IAMActivities {
    client := iam.New(session, config...)
    return &IAMActivities{client: client}
}

func (a *IAMActivities) AddClientIDToOpenIDConnectProvider(ctx context.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
    return a.client.AddClientIDToOpenIDConnectProviderWithContext(ctx, input)
}

func (a *IAMActivities) AddRoleToInstanceProfile(ctx context.Context, input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error) {
    return a.client.AddRoleToInstanceProfileWithContext(ctx, input)
}

func (a *IAMActivities) AddUserToGroup(ctx context.Context, input *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error) {
    return a.client.AddUserToGroupWithContext(ctx, input)
}

func (a *IAMActivities) AttachGroupPolicy(ctx context.Context, input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {
    return a.client.AttachGroupPolicyWithContext(ctx, input)
}

func (a *IAMActivities) AttachRolePolicy(ctx context.Context, input *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {
    return a.client.AttachRolePolicyWithContext(ctx, input)
}

func (a *IAMActivities) AttachUserPolicy(ctx context.Context, input *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {
    return a.client.AttachUserPolicyWithContext(ctx, input)
}

func (a *IAMActivities) ChangePassword(ctx context.Context, input *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error) {
    return a.client.ChangePasswordWithContext(ctx, input)
}

func (a *IAMActivities) CreateAccessKey(ctx context.Context, input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
    return a.client.CreateAccessKeyWithContext(ctx, input)
}

func (a *IAMActivities) CreateAccountAlias(ctx context.Context, input *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error) {
    return a.client.CreateAccountAliasWithContext(ctx, input)
}

func (a *IAMActivities) CreateGroup(ctx context.Context, input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error) {
    return a.client.CreateGroupWithContext(ctx, input)
}

func (a *IAMActivities) CreateInstanceProfile(ctx context.Context, input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error) {
    return a.client.CreateInstanceProfileWithContext(ctx, input)
}

func (a *IAMActivities) CreateLoginProfile(ctx context.Context, input *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error) {
    return a.client.CreateLoginProfileWithContext(ctx, input)
}

func (a *IAMActivities) CreateOpenIDConnectProvider(ctx context.Context, input *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error) {
    return a.client.CreateOpenIDConnectProviderWithContext(ctx, input)
}

func (a *IAMActivities) CreatePolicy(ctx context.Context, input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
    return a.client.CreatePolicyWithContext(ctx, input)
}

func (a *IAMActivities) CreatePolicyVersion(ctx context.Context, input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {
    return a.client.CreatePolicyVersionWithContext(ctx, input)
}

func (a *IAMActivities) CreateRole(ctx context.Context, input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
    return a.client.CreateRoleWithContext(ctx, input)
}

func (a *IAMActivities) CreateSAMLProvider(ctx context.Context, input *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error) {
    return a.client.CreateSAMLProviderWithContext(ctx, input)
}

func (a *IAMActivities) CreateServiceLinkedRole(ctx context.Context, input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error) {
    return a.client.CreateServiceLinkedRoleWithContext(ctx, input)
}

func (a *IAMActivities) CreateServiceSpecificCredential(ctx context.Context, input *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error) {
    return a.client.CreateServiceSpecificCredentialWithContext(ctx, input)
}

func (a *IAMActivities) CreateUser(ctx context.Context, input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
    return a.client.CreateUserWithContext(ctx, input)
}

func (a *IAMActivities) CreateVirtualMFADevice(ctx context.Context, input *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error) {
    return a.client.CreateVirtualMFADeviceWithContext(ctx, input)
}

func (a *IAMActivities) DeactivateMFADevice(ctx context.Context, input *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error) {
    return a.client.DeactivateMFADeviceWithContext(ctx, input)
}

func (a *IAMActivities) DeleteAccessKey(ctx context.Context, input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
    return a.client.DeleteAccessKeyWithContext(ctx, input)
}

func (a *IAMActivities) DeleteAccountAlias(ctx context.Context, input *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error) {
    return a.client.DeleteAccountAliasWithContext(ctx, input)
}

func (a *IAMActivities) DeleteAccountPasswordPolicy(ctx context.Context, input *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error) {
    return a.client.DeleteAccountPasswordPolicyWithContext(ctx, input)
}

func (a *IAMActivities) DeleteGroup(ctx context.Context, input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error) {
    return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *IAMActivities) DeleteGroupPolicy(ctx context.Context, input *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error) {
    return a.client.DeleteGroupPolicyWithContext(ctx, input)
}

func (a *IAMActivities) DeleteInstanceProfile(ctx context.Context, input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
    return a.client.DeleteInstanceProfileWithContext(ctx, input)
}

func (a *IAMActivities) DeleteLoginProfile(ctx context.Context, input *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error) {
    return a.client.DeleteLoginProfileWithContext(ctx, input)
}

func (a *IAMActivities) DeleteOpenIDConnectProvider(ctx context.Context, input *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error) {
    return a.client.DeleteOpenIDConnectProviderWithContext(ctx, input)
}

func (a *IAMActivities) DeletePolicy(ctx context.Context, input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
    return a.client.DeletePolicyWithContext(ctx, input)
}

func (a *IAMActivities) DeletePolicyVersion(ctx context.Context, input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error) {
    return a.client.DeletePolicyVersionWithContext(ctx, input)
}

func (a *IAMActivities) DeleteRole(ctx context.Context, input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
    return a.client.DeleteRoleWithContext(ctx, input)
}

func (a *IAMActivities) DeleteRolePermissionsBoundary(ctx context.Context, input *iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error) {
    return a.client.DeleteRolePermissionsBoundaryWithContext(ctx, input)
}

func (a *IAMActivities) DeleteRolePolicy(ctx context.Context, input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
    return a.client.DeleteRolePolicyWithContext(ctx, input)
}

func (a *IAMActivities) DeleteSAMLProvider(ctx context.Context, input *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error) {
    return a.client.DeleteSAMLProviderWithContext(ctx, input)
}

func (a *IAMActivities) DeleteSSHPublicKey(ctx context.Context, input *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error) {
    return a.client.DeleteSSHPublicKeyWithContext(ctx, input)
}

func (a *IAMActivities) DeleteServerCertificate(ctx context.Context, input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
    return a.client.DeleteServerCertificateWithContext(ctx, input)
}

func (a *IAMActivities) DeleteServiceLinkedRole(ctx context.Context, input *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error) {
    return a.client.DeleteServiceLinkedRoleWithContext(ctx, input)
}

func (a *IAMActivities) DeleteServiceSpecificCredential(ctx context.Context, input *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error) {
    return a.client.DeleteServiceSpecificCredentialWithContext(ctx, input)
}

func (a *IAMActivities) DeleteSigningCertificate(ctx context.Context, input *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error) {
    return a.client.DeleteSigningCertificateWithContext(ctx, input)
}

func (a *IAMActivities) DeleteUser(ctx context.Context, input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
    return a.client.DeleteUserWithContext(ctx, input)
}

func (a *IAMActivities) DeleteUserPermissionsBoundary(ctx context.Context, input *iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error) {
    return a.client.DeleteUserPermissionsBoundaryWithContext(ctx, input)
}

func (a *IAMActivities) DeleteUserPolicy(ctx context.Context, input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
    return a.client.DeleteUserPolicyWithContext(ctx, input)
}

func (a *IAMActivities) DeleteVirtualMFADevice(ctx context.Context, input *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error) {
    return a.client.DeleteVirtualMFADeviceWithContext(ctx, input)
}

func (a *IAMActivities) DetachGroupPolicy(ctx context.Context, input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error) {
    return a.client.DetachGroupPolicyWithContext(ctx, input)
}

func (a *IAMActivities) DetachRolePolicy(ctx context.Context, input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
    return a.client.DetachRolePolicyWithContext(ctx, input)
}

func (a *IAMActivities) DetachUserPolicy(ctx context.Context, input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {
    return a.client.DetachUserPolicyWithContext(ctx, input)
}

func (a *IAMActivities) EnableMFADevice(ctx context.Context, input *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error) {
    return a.client.EnableMFADeviceWithContext(ctx, input)
}

func (a *IAMActivities) GenerateCredentialReport(ctx context.Context, input *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error) {
    return a.client.GenerateCredentialReportWithContext(ctx, input)
}

func (a *IAMActivities) GenerateOrganizationsAccessReport(ctx context.Context, input *iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error) {
    return a.client.GenerateOrganizationsAccessReportWithContext(ctx, input)
}

func (a *IAMActivities) GenerateServiceLastAccessedDetails(ctx context.Context, input *iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error) {
    return a.client.GenerateServiceLastAccessedDetailsWithContext(ctx, input)
}

func (a *IAMActivities) GetAccessKeyLastUsed(ctx context.Context, input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error) {
    return a.client.GetAccessKeyLastUsedWithContext(ctx, input)
}

func (a *IAMActivities) GetAccountAuthorizationDetails(ctx context.Context, input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {
    return a.client.GetAccountAuthorizationDetailsWithContext(ctx, input)
}

func (a *IAMActivities) GetAccountPasswordPolicy(ctx context.Context, input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error) {
    return a.client.GetAccountPasswordPolicyWithContext(ctx, input)
}

func (a *IAMActivities) GetAccountSummary(ctx context.Context, input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error) {
    return a.client.GetAccountSummaryWithContext(ctx, input)
}

func (a *IAMActivities) GetContextKeysForCustomPolicy(ctx context.Context, input *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
    return a.client.GetContextKeysForCustomPolicyWithContext(ctx, input)
}

func (a *IAMActivities) GetContextKeysForPrincipalPolicy(ctx context.Context, input *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
    return a.client.GetContextKeysForPrincipalPolicyWithContext(ctx, input)
}

func (a *IAMActivities) GetCredentialReport(ctx context.Context, input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error) {
    return a.client.GetCredentialReportWithContext(ctx, input)
}

func (a *IAMActivities) GetGroup(ctx context.Context, input *iam.GetGroupInput) (*iam.GetGroupOutput, error) {
    return a.client.GetGroupWithContext(ctx, input)
}

func (a *IAMActivities) GetGroupPolicy(ctx context.Context, input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error) {
    return a.client.GetGroupPolicyWithContext(ctx, input)
}

func (a *IAMActivities) GetInstanceProfile(ctx context.Context, input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error) {
    return a.client.GetInstanceProfileWithContext(ctx, input)
}

func (a *IAMActivities) GetLoginProfile(ctx context.Context, input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error) {
    return a.client.GetLoginProfileWithContext(ctx, input)
}

func (a *IAMActivities) GetOpenIDConnectProvider(ctx context.Context, input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error) {
    return a.client.GetOpenIDConnectProviderWithContext(ctx, input)
}

func (a *IAMActivities) GetOrganizationsAccessReport(ctx context.Context, input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error) {
    return a.client.GetOrganizationsAccessReportWithContext(ctx, input)
}

func (a *IAMActivities) GetPolicy(ctx context.Context, input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error) {
    return a.client.GetPolicyWithContext(ctx, input)
}

func (a *IAMActivities) GetPolicyVersion(ctx context.Context, input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {
    return a.client.GetPolicyVersionWithContext(ctx, input)
}

func (a *IAMActivities) GetRole(ctx context.Context, input *iam.GetRoleInput) (*iam.GetRoleOutput, error) {
    return a.client.GetRoleWithContext(ctx, input)
}

func (a *IAMActivities) GetRolePolicy(ctx context.Context, input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error) {
    return a.client.GetRolePolicyWithContext(ctx, input)
}

func (a *IAMActivities) GetSAMLProvider(ctx context.Context, input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error) {
    return a.client.GetSAMLProviderWithContext(ctx, input)
}

func (a *IAMActivities) GetSSHPublicKey(ctx context.Context, input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error) {
    return a.client.GetSSHPublicKeyWithContext(ctx, input)
}

func (a *IAMActivities) GetServerCertificate(ctx context.Context, input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error) {
    return a.client.GetServerCertificateWithContext(ctx, input)
}

func (a *IAMActivities) GetServiceLastAccessedDetails(ctx context.Context, input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error) {
    return a.client.GetServiceLastAccessedDetailsWithContext(ctx, input)
}

func (a *IAMActivities) GetServiceLastAccessedDetailsWithEntities(ctx context.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
    return a.client.GetServiceLastAccessedDetailsWithEntitiesWithContext(ctx, input)
}

func (a *IAMActivities) GetServiceLinkedRoleDeletionStatus(ctx context.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
    return a.client.GetServiceLinkedRoleDeletionStatusWithContext(ctx, input)
}

func (a *IAMActivities) GetUser(ctx context.Context, input *iam.GetUserInput) (*iam.GetUserOutput, error) {
    return a.client.GetUserWithContext(ctx, input)
}

func (a *IAMActivities) GetUserPolicy(ctx context.Context, input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {
    return a.client.GetUserPolicyWithContext(ctx, input)
}

func (a *IAMActivities) ListAccessKeys(ctx context.Context, input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
    return a.client.ListAccessKeysWithContext(ctx, input)
}

func (a *IAMActivities) ListAccountAliases(ctx context.Context, input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {
    return a.client.ListAccountAliasesWithContext(ctx, input)
}

func (a *IAMActivities) ListAttachedGroupPolicies(ctx context.Context, input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {
    return a.client.ListAttachedGroupPoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListAttachedRolePolicies(ctx context.Context, input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
    return a.client.ListAttachedRolePoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListAttachedUserPolicies(ctx context.Context, input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
    return a.client.ListAttachedUserPoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListEntitiesForPolicy(ctx context.Context, input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error) {
    return a.client.ListEntitiesForPolicyWithContext(ctx, input)
}

func (a *IAMActivities) ListGroupPolicies(ctx context.Context, input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error) {
    return a.client.ListGroupPoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListGroups(ctx context.Context, input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error) {
    return a.client.ListGroupsWithContext(ctx, input)
}

func (a *IAMActivities) ListGroupsForUser(ctx context.Context, input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error) {
    return a.client.ListGroupsForUserWithContext(ctx, input)
}

func (a *IAMActivities) ListInstanceProfiles(ctx context.Context, input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
    return a.client.ListInstanceProfilesWithContext(ctx, input)
}

func (a *IAMActivities) ListInstanceProfilesForRole(ctx context.Context, input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error) {
    return a.client.ListInstanceProfilesForRoleWithContext(ctx, input)
}

func (a *IAMActivities) ListMFADevices(ctx context.Context, input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error) {
    return a.client.ListMFADevicesWithContext(ctx, input)
}

func (a *IAMActivities) ListOpenIDConnectProviders(ctx context.Context, input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error) {
    return a.client.ListOpenIDConnectProvidersWithContext(ctx, input)
}

func (a *IAMActivities) ListPolicies(ctx context.Context, input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
    return a.client.ListPoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListPoliciesGrantingServiceAccess(ctx context.Context, input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
    return a.client.ListPoliciesGrantingServiceAccessWithContext(ctx, input)
}

func (a *IAMActivities) ListPolicyVersions(ctx context.Context, input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {
    return a.client.ListPolicyVersionsWithContext(ctx, input)
}

func (a *IAMActivities) ListRolePolicies(ctx context.Context, input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
    return a.client.ListRolePoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListRoleTags(ctx context.Context, input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error) {
    return a.client.ListRoleTagsWithContext(ctx, input)
}

func (a *IAMActivities) ListRoles(ctx context.Context, input *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
    return a.client.ListRolesWithContext(ctx, input)
}

func (a *IAMActivities) ListSAMLProviders(ctx context.Context, input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error) {
    return a.client.ListSAMLProvidersWithContext(ctx, input)
}

func (a *IAMActivities) ListSSHPublicKeys(ctx context.Context, input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error) {
    return a.client.ListSSHPublicKeysWithContext(ctx, input)
}

func (a *IAMActivities) ListServerCertificates(ctx context.Context, input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
    return a.client.ListServerCertificatesWithContext(ctx, input)
}

func (a *IAMActivities) ListServiceSpecificCredentials(ctx context.Context, input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error) {
    return a.client.ListServiceSpecificCredentialsWithContext(ctx, input)
}

func (a *IAMActivities) ListSigningCertificates(ctx context.Context, input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error) {
    return a.client.ListSigningCertificatesWithContext(ctx, input)
}

func (a *IAMActivities) ListUserPolicies(ctx context.Context, input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
    return a.client.ListUserPoliciesWithContext(ctx, input)
}

func (a *IAMActivities) ListUserTags(ctx context.Context, input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error) {
    return a.client.ListUserTagsWithContext(ctx, input)
}

func (a *IAMActivities) ListUsers(ctx context.Context, input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *IAMActivities) ListVirtualMFADevices(ctx context.Context, input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error) {
    return a.client.ListVirtualMFADevicesWithContext(ctx, input)
}

func (a *IAMActivities) PutGroupPolicy(ctx context.Context, input *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error) {
    return a.client.PutGroupPolicyWithContext(ctx, input)
}

func (a *IAMActivities) PutRolePermissionsBoundary(ctx context.Context, input *iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error) {
    return a.client.PutRolePermissionsBoundaryWithContext(ctx, input)
}

func (a *IAMActivities) PutRolePolicy(ctx context.Context, input *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error) {
    return a.client.PutRolePolicyWithContext(ctx, input)
}

func (a *IAMActivities) PutUserPermissionsBoundary(ctx context.Context, input *iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error) {
    return a.client.PutUserPermissionsBoundaryWithContext(ctx, input)
}

func (a *IAMActivities) PutUserPolicy(ctx context.Context, input *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
    return a.client.PutUserPolicyWithContext(ctx, input)
}

func (a *IAMActivities) RemoveClientIDFromOpenIDConnectProvider(ctx context.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {
    return a.client.RemoveClientIDFromOpenIDConnectProviderWithContext(ctx, input)
}

func (a *IAMActivities) RemoveRoleFromInstanceProfile(ctx context.Context, input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
    return a.client.RemoveRoleFromInstanceProfileWithContext(ctx, input)
}

func (a *IAMActivities) RemoveUserFromGroup(ctx context.Context, input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error) {
    return a.client.RemoveUserFromGroupWithContext(ctx, input)
}

func (a *IAMActivities) ResetServiceSpecificCredential(ctx context.Context, input *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error) {
    return a.client.ResetServiceSpecificCredentialWithContext(ctx, input)
}

func (a *IAMActivities) ResyncMFADevice(ctx context.Context, input *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error) {
    return a.client.ResyncMFADeviceWithContext(ctx, input)
}

func (a *IAMActivities) SetDefaultPolicyVersion(ctx context.Context, input *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error) {
    return a.client.SetDefaultPolicyVersionWithContext(ctx, input)
}

func (a *IAMActivities) SetSecurityTokenServicePreferences(ctx context.Context, input *iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error) {
    return a.client.SetSecurityTokenServicePreferencesWithContext(ctx, input)
}

func (a *IAMActivities) SimulateCustomPolicy(ctx context.Context, input *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error) {
    return a.client.SimulateCustomPolicyWithContext(ctx, input)
}

func (a *IAMActivities) SimulatePrincipalPolicy(ctx context.Context, input *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error) {
    return a.client.SimulatePrincipalPolicyWithContext(ctx, input)
}

func (a *IAMActivities) TagRole(ctx context.Context, input *iam.TagRoleInput) (*iam.TagRoleOutput, error) {
    return a.client.TagRoleWithContext(ctx, input)
}

func (a *IAMActivities) TagUser(ctx context.Context, input *iam.TagUserInput) (*iam.TagUserOutput, error) {
    return a.client.TagUserWithContext(ctx, input)
}

func (a *IAMActivities) UntagRole(ctx context.Context, input *iam.UntagRoleInput) (*iam.UntagRoleOutput, error) {
    return a.client.UntagRoleWithContext(ctx, input)
}

func (a *IAMActivities) UntagUser(ctx context.Context, input *iam.UntagUserInput) (*iam.UntagUserOutput, error) {
    return a.client.UntagUserWithContext(ctx, input)
}

func (a *IAMActivities) UpdateAccessKey(ctx context.Context, input *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error) {
    return a.client.UpdateAccessKeyWithContext(ctx, input)
}

func (a *IAMActivities) UpdateAccountPasswordPolicy(ctx context.Context, input *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error) {
    return a.client.UpdateAccountPasswordPolicyWithContext(ctx, input)
}

func (a *IAMActivities) UpdateAssumeRolePolicy(ctx context.Context, input *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error) {
    return a.client.UpdateAssumeRolePolicyWithContext(ctx, input)
}

func (a *IAMActivities) UpdateGroup(ctx context.Context, input *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error) {
    return a.client.UpdateGroupWithContext(ctx, input)
}

func (a *IAMActivities) UpdateLoginProfile(ctx context.Context, input *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error) {
    return a.client.UpdateLoginProfileWithContext(ctx, input)
}

func (a *IAMActivities) UpdateOpenIDConnectProviderThumbprint(ctx context.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {
    return a.client.UpdateOpenIDConnectProviderThumbprintWithContext(ctx, input)
}

func (a *IAMActivities) UpdateRole(ctx context.Context, input *iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error) {
    return a.client.UpdateRoleWithContext(ctx, input)
}

func (a *IAMActivities) UpdateRoleDescription(ctx context.Context, input *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error) {
    return a.client.UpdateRoleDescriptionWithContext(ctx, input)
}

func (a *IAMActivities) UpdateSAMLProvider(ctx context.Context, input *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error) {
    return a.client.UpdateSAMLProviderWithContext(ctx, input)
}

func (a *IAMActivities) UpdateSSHPublicKey(ctx context.Context, input *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error) {
    return a.client.UpdateSSHPublicKeyWithContext(ctx, input)
}

func (a *IAMActivities) UpdateServerCertificate(ctx context.Context, input *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error) {
    return a.client.UpdateServerCertificateWithContext(ctx, input)
}

func (a *IAMActivities) UpdateServiceSpecificCredential(ctx context.Context, input *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error) {
    return a.client.UpdateServiceSpecificCredentialWithContext(ctx, input)
}

func (a *IAMActivities) UpdateSigningCertificate(ctx context.Context, input *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error) {
    return a.client.UpdateSigningCertificateWithContext(ctx, input)
}

func (a *IAMActivities) UpdateUser(ctx context.Context, input *iam.UpdateUserInput) (*iam.UpdateUserOutput, error) {
    return a.client.UpdateUserWithContext(ctx, input)
}

func (a *IAMActivities) UploadSSHPublicKey(ctx context.Context, input *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error) {
    return a.client.UploadSSHPublicKeyWithContext(ctx, input)
}

func (a *IAMActivities) UploadServerCertificate(ctx context.Context, input *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error) {
    return a.client.UploadServerCertificateWithContext(ctx, input)
}

func (a *IAMActivities) UploadSigningCertificate(ctx context.Context, input *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error) {
    return a.client.UploadSigningCertificateWithContext(ctx, input)
}

func (a *IAMActivities) WaitUntilInstanceProfileExists(ctx context.Context, input *iam.GetInstanceProfileInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceProfileExistsWithContext(ctx, input, options...)
	})
}

func (a *IAMActivities) WaitUntilPolicyExists(ctx context.Context, input *iam.GetPolicyInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilPolicyExistsWithContext(ctx, input, options...)
	})
}

func (a *IAMActivities) WaitUntilRoleExists(ctx context.Context, input *iam.GetRoleInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilRoleExistsWithContext(ctx, input, options...)
	})
}

func (a *IAMActivities) WaitUntilUserExists(ctx context.Context, input *iam.GetUserInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilUserExistsWithContext(ctx, input, options...)
	})
}
