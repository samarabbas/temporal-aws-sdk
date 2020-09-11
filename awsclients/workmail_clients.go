package awsclients

import (
	"github.com/aws/aws-sdk-go/service/workmail"
	"go.temporal.io/sdk/workflow"
)

type WorkMailClient interface {
       AssociateDelegateToResource(ctx workflow.Context, input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error)
       AssociateDelegateToResourceAsync(ctx workflow.Context, input *workmail.AssociateDelegateToResourceInput) *WorkmailAssociateDelegateToResourceResult

       AssociateMemberToGroup(ctx workflow.Context, input *workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error)
       AssociateMemberToGroupAsync(ctx workflow.Context, input *workmail.AssociateMemberToGroupInput) *WorkmailAssociateMemberToGroupResult

       CreateAlias(ctx workflow.Context, input *workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error)
       CreateAliasAsync(ctx workflow.Context, input *workmail.CreateAliasInput) *WorkmailCreateAliasResult

       CreateGroup(ctx workflow.Context, input *workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error)
       CreateGroupAsync(ctx workflow.Context, input *workmail.CreateGroupInput) *WorkmailCreateGroupResult

       CreateResource(ctx workflow.Context, input *workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error)
       CreateResourceAsync(ctx workflow.Context, input *workmail.CreateResourceInput) *WorkmailCreateResourceResult

       CreateUser(ctx workflow.Context, input *workmail.CreateUserInput) (*workmail.CreateUserOutput, error)
       CreateUserAsync(ctx workflow.Context, input *workmail.CreateUserInput) *WorkmailCreateUserResult

       DeleteAccessControlRule(ctx workflow.Context, input *workmail.DeleteAccessControlRuleInput) (*workmail.DeleteAccessControlRuleOutput, error)
       DeleteAccessControlRuleAsync(ctx workflow.Context, input *workmail.DeleteAccessControlRuleInput) *WorkmailDeleteAccessControlRuleResult

       DeleteAlias(ctx workflow.Context, input *workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error)
       DeleteAliasAsync(ctx workflow.Context, input *workmail.DeleteAliasInput) *WorkmailDeleteAliasResult

       DeleteGroup(ctx workflow.Context, input *workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error)
       DeleteGroupAsync(ctx workflow.Context, input *workmail.DeleteGroupInput) *WorkmailDeleteGroupResult

       DeleteMailboxPermissions(ctx workflow.Context, input *workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error)
       DeleteMailboxPermissionsAsync(ctx workflow.Context, input *workmail.DeleteMailboxPermissionsInput) *WorkmailDeleteMailboxPermissionsResult

       DeleteResource(ctx workflow.Context, input *workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error)
       DeleteResourceAsync(ctx workflow.Context, input *workmail.DeleteResourceInput) *WorkmailDeleteResourceResult

       DeleteRetentionPolicy(ctx workflow.Context, input *workmail.DeleteRetentionPolicyInput) (*workmail.DeleteRetentionPolicyOutput, error)
       DeleteRetentionPolicyAsync(ctx workflow.Context, input *workmail.DeleteRetentionPolicyInput) *WorkmailDeleteRetentionPolicyResult

       DeleteUser(ctx workflow.Context, input *workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error)
       DeleteUserAsync(ctx workflow.Context, input *workmail.DeleteUserInput) *WorkmailDeleteUserResult

       DeregisterFromWorkMail(ctx workflow.Context, input *workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error)
       DeregisterFromWorkMailAsync(ctx workflow.Context, input *workmail.DeregisterFromWorkMailInput) *WorkmailDeregisterFromWorkMailResult

       DescribeGroup(ctx workflow.Context, input *workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error)
       DescribeGroupAsync(ctx workflow.Context, input *workmail.DescribeGroupInput) *WorkmailDescribeGroupResult

       DescribeOrganization(ctx workflow.Context, input *workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error)
       DescribeOrganizationAsync(ctx workflow.Context, input *workmail.DescribeOrganizationInput) *WorkmailDescribeOrganizationResult

       DescribeResource(ctx workflow.Context, input *workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error)
       DescribeResourceAsync(ctx workflow.Context, input *workmail.DescribeResourceInput) *WorkmailDescribeResourceResult

       DescribeUser(ctx workflow.Context, input *workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error)
       DescribeUserAsync(ctx workflow.Context, input *workmail.DescribeUserInput) *WorkmailDescribeUserResult

       DisassociateDelegateFromResource(ctx workflow.Context, input *workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error)
       DisassociateDelegateFromResourceAsync(ctx workflow.Context, input *workmail.DisassociateDelegateFromResourceInput) *WorkmailDisassociateDelegateFromResourceResult

       DisassociateMemberFromGroup(ctx workflow.Context, input *workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error)
       DisassociateMemberFromGroupAsync(ctx workflow.Context, input *workmail.DisassociateMemberFromGroupInput) *WorkmailDisassociateMemberFromGroupResult

       GetAccessControlEffect(ctx workflow.Context, input *workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error)
       GetAccessControlEffectAsync(ctx workflow.Context, input *workmail.GetAccessControlEffectInput) *WorkmailGetAccessControlEffectResult

       GetDefaultRetentionPolicy(ctx workflow.Context, input *workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error)
       GetDefaultRetentionPolicyAsync(ctx workflow.Context, input *workmail.GetDefaultRetentionPolicyInput) *WorkmailGetDefaultRetentionPolicyResult

       GetMailboxDetails(ctx workflow.Context, input *workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error)
       GetMailboxDetailsAsync(ctx workflow.Context, input *workmail.GetMailboxDetailsInput) *WorkmailGetMailboxDetailsResult

       ListAccessControlRules(ctx workflow.Context, input *workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error)
       ListAccessControlRulesAsync(ctx workflow.Context, input *workmail.ListAccessControlRulesInput) *WorkmailListAccessControlRulesResult

       ListAliases(ctx workflow.Context, input *workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error)
       ListAliasesAsync(ctx workflow.Context, input *workmail.ListAliasesInput) *WorkmailListAliasesResult

       ListGroupMembers(ctx workflow.Context, input *workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error)
       ListGroupMembersAsync(ctx workflow.Context, input *workmail.ListGroupMembersInput) *WorkmailListGroupMembersResult

       ListGroups(ctx workflow.Context, input *workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error)
       ListGroupsAsync(ctx workflow.Context, input *workmail.ListGroupsInput) *WorkmailListGroupsResult

       ListMailboxPermissions(ctx workflow.Context, input *workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error)
       ListMailboxPermissionsAsync(ctx workflow.Context, input *workmail.ListMailboxPermissionsInput) *WorkmailListMailboxPermissionsResult

       ListOrganizations(ctx workflow.Context, input *workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error)
       ListOrganizationsAsync(ctx workflow.Context, input *workmail.ListOrganizationsInput) *WorkmailListOrganizationsResult

       ListResourceDelegates(ctx workflow.Context, input *workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error)
       ListResourceDelegatesAsync(ctx workflow.Context, input *workmail.ListResourceDelegatesInput) *WorkmailListResourceDelegatesResult

       ListResources(ctx workflow.Context, input *workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error)
       ListResourcesAsync(ctx workflow.Context, input *workmail.ListResourcesInput) *WorkmailListResourcesResult

       ListTagsForResource(ctx workflow.Context, input *workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *workmail.ListTagsForResourceInput) *WorkmailListTagsForResourceResult

       ListUsers(ctx workflow.Context, input *workmail.ListUsersInput) (*workmail.ListUsersOutput, error)
       ListUsersAsync(ctx workflow.Context, input *workmail.ListUsersInput) *WorkmailListUsersResult

       PutAccessControlRule(ctx workflow.Context, input *workmail.PutAccessControlRuleInput) (*workmail.PutAccessControlRuleOutput, error)
       PutAccessControlRuleAsync(ctx workflow.Context, input *workmail.PutAccessControlRuleInput) *WorkmailPutAccessControlRuleResult

       PutMailboxPermissions(ctx workflow.Context, input *workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error)
       PutMailboxPermissionsAsync(ctx workflow.Context, input *workmail.PutMailboxPermissionsInput) *WorkmailPutMailboxPermissionsResult

       PutRetentionPolicy(ctx workflow.Context, input *workmail.PutRetentionPolicyInput) (*workmail.PutRetentionPolicyOutput, error)
       PutRetentionPolicyAsync(ctx workflow.Context, input *workmail.PutRetentionPolicyInput) *WorkmailPutRetentionPolicyResult

       RegisterToWorkMail(ctx workflow.Context, input *workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error)
       RegisterToWorkMailAsync(ctx workflow.Context, input *workmail.RegisterToWorkMailInput) *WorkmailRegisterToWorkMailResult

       ResetPassword(ctx workflow.Context, input *workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error)
       ResetPasswordAsync(ctx workflow.Context, input *workmail.ResetPasswordInput) *WorkmailResetPasswordResult

       TagResource(ctx workflow.Context, input *workmail.TagResourceInput) (*workmail.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *workmail.TagResourceInput) *WorkmailTagResourceResult

       UntagResource(ctx workflow.Context, input *workmail.UntagResourceInput) (*workmail.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *workmail.UntagResourceInput) *WorkmailUntagResourceResult

       UpdateMailboxQuota(ctx workflow.Context, input *workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error)
       UpdateMailboxQuotaAsync(ctx workflow.Context, input *workmail.UpdateMailboxQuotaInput) *WorkmailUpdateMailboxQuotaResult

       UpdatePrimaryEmailAddress(ctx workflow.Context, input *workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error)
       UpdatePrimaryEmailAddressAsync(ctx workflow.Context, input *workmail.UpdatePrimaryEmailAddressInput) *WorkmailUpdatePrimaryEmailAddressResult

       UpdateResource(ctx workflow.Context, input *workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error)
       UpdateResourceAsync(ctx workflow.Context, input *workmail.UpdateResourceInput) *WorkmailUpdateResourceResult
}

type WorkMailStub struct {
}

func NewWorkMailStub() WorkMailClient {
    return &WorkMailStub{}
}

type WorkmailAssociateDelegateToResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailAssociateDelegateToResourceResult) Get(ctx workflow.Context) (*workmail.AssociateDelegateToResourceOutput, error) {
    var output workmail.AssociateDelegateToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailAssociateMemberToGroupResult struct {
	Result workflow.Future
}

func (r *WorkmailAssociateMemberToGroupResult) Get(ctx workflow.Context) (*workmail.AssociateMemberToGroupOutput, error) {
    var output workmail.AssociateMemberToGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailCreateAliasResult struct {
	Result workflow.Future
}

func (r *WorkmailCreateAliasResult) Get(ctx workflow.Context) (*workmail.CreateAliasOutput, error) {
    var output workmail.CreateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailCreateGroupResult struct {
	Result workflow.Future
}

func (r *WorkmailCreateGroupResult) Get(ctx workflow.Context) (*workmail.CreateGroupOutput, error) {
    var output workmail.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailCreateResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailCreateResourceResult) Get(ctx workflow.Context) (*workmail.CreateResourceOutput, error) {
    var output workmail.CreateResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailCreateUserResult struct {
	Result workflow.Future
}

func (r *WorkmailCreateUserResult) Get(ctx workflow.Context) (*workmail.CreateUserOutput, error) {
    var output workmail.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteAccessControlRuleResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteAccessControlRuleResult) Get(ctx workflow.Context) (*workmail.DeleteAccessControlRuleOutput, error) {
    var output workmail.DeleteAccessControlRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteAliasResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteAliasResult) Get(ctx workflow.Context) (*workmail.DeleteAliasOutput, error) {
    var output workmail.DeleteAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteGroupResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteGroupResult) Get(ctx workflow.Context) (*workmail.DeleteGroupOutput, error) {
    var output workmail.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteMailboxPermissionsResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteMailboxPermissionsResult) Get(ctx workflow.Context) (*workmail.DeleteMailboxPermissionsOutput, error) {
    var output workmail.DeleteMailboxPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteResourceResult) Get(ctx workflow.Context) (*workmail.DeleteResourceOutput, error) {
    var output workmail.DeleteResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteRetentionPolicyResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteRetentionPolicyResult) Get(ctx workflow.Context) (*workmail.DeleteRetentionPolicyOutput, error) {
    var output workmail.DeleteRetentionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeleteUserResult struct {
	Result workflow.Future
}

func (r *WorkmailDeleteUserResult) Get(ctx workflow.Context) (*workmail.DeleteUserOutput, error) {
    var output workmail.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDeregisterFromWorkMailResult struct {
	Result workflow.Future
}

func (r *WorkmailDeregisterFromWorkMailResult) Get(ctx workflow.Context) (*workmail.DeregisterFromWorkMailOutput, error) {
    var output workmail.DeregisterFromWorkMailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDescribeGroupResult struct {
	Result workflow.Future
}

func (r *WorkmailDescribeGroupResult) Get(ctx workflow.Context) (*workmail.DescribeGroupOutput, error) {
    var output workmail.DescribeGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDescribeOrganizationResult struct {
	Result workflow.Future
}

func (r *WorkmailDescribeOrganizationResult) Get(ctx workflow.Context) (*workmail.DescribeOrganizationOutput, error) {
    var output workmail.DescribeOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDescribeResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailDescribeResourceResult) Get(ctx workflow.Context) (*workmail.DescribeResourceOutput, error) {
    var output workmail.DescribeResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDescribeUserResult struct {
	Result workflow.Future
}

func (r *WorkmailDescribeUserResult) Get(ctx workflow.Context) (*workmail.DescribeUserOutput, error) {
    var output workmail.DescribeUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDisassociateDelegateFromResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailDisassociateDelegateFromResourceResult) Get(ctx workflow.Context) (*workmail.DisassociateDelegateFromResourceOutput, error) {
    var output workmail.DisassociateDelegateFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailDisassociateMemberFromGroupResult struct {
	Result workflow.Future
}

func (r *WorkmailDisassociateMemberFromGroupResult) Get(ctx workflow.Context) (*workmail.DisassociateMemberFromGroupOutput, error) {
    var output workmail.DisassociateMemberFromGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailGetAccessControlEffectResult struct {
	Result workflow.Future
}

func (r *WorkmailGetAccessControlEffectResult) Get(ctx workflow.Context) (*workmail.GetAccessControlEffectOutput, error) {
    var output workmail.GetAccessControlEffectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailGetDefaultRetentionPolicyResult struct {
	Result workflow.Future
}

func (r *WorkmailGetDefaultRetentionPolicyResult) Get(ctx workflow.Context) (*workmail.GetDefaultRetentionPolicyOutput, error) {
    var output workmail.GetDefaultRetentionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailGetMailboxDetailsResult struct {
	Result workflow.Future
}

func (r *WorkmailGetMailboxDetailsResult) Get(ctx workflow.Context) (*workmail.GetMailboxDetailsOutput, error) {
    var output workmail.GetMailboxDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListAccessControlRulesResult struct {
	Result workflow.Future
}

func (r *WorkmailListAccessControlRulesResult) Get(ctx workflow.Context) (*workmail.ListAccessControlRulesOutput, error) {
    var output workmail.ListAccessControlRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListAliasesResult struct {
	Result workflow.Future
}

func (r *WorkmailListAliasesResult) Get(ctx workflow.Context) (*workmail.ListAliasesOutput, error) {
    var output workmail.ListAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListGroupMembersResult struct {
	Result workflow.Future
}

func (r *WorkmailListGroupMembersResult) Get(ctx workflow.Context) (*workmail.ListGroupMembersOutput, error) {
    var output workmail.ListGroupMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListGroupsResult struct {
	Result workflow.Future
}

func (r *WorkmailListGroupsResult) Get(ctx workflow.Context) (*workmail.ListGroupsOutput, error) {
    var output workmail.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListMailboxPermissionsResult struct {
	Result workflow.Future
}

func (r *WorkmailListMailboxPermissionsResult) Get(ctx workflow.Context) (*workmail.ListMailboxPermissionsOutput, error) {
    var output workmail.ListMailboxPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListOrganizationsResult struct {
	Result workflow.Future
}

func (r *WorkmailListOrganizationsResult) Get(ctx workflow.Context) (*workmail.ListOrganizationsOutput, error) {
    var output workmail.ListOrganizationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListResourceDelegatesResult struct {
	Result workflow.Future
}

func (r *WorkmailListResourceDelegatesResult) Get(ctx workflow.Context) (*workmail.ListResourceDelegatesOutput, error) {
    var output workmail.ListResourceDelegatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListResourcesResult struct {
	Result workflow.Future
}

func (r *WorkmailListResourcesResult) Get(ctx workflow.Context) (*workmail.ListResourcesOutput, error) {
    var output workmail.ListResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailListTagsForResourceResult) Get(ctx workflow.Context) (*workmail.ListTagsForResourceOutput, error) {
    var output workmail.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailListUsersResult struct {
	Result workflow.Future
}

func (r *WorkmailListUsersResult) Get(ctx workflow.Context) (*workmail.ListUsersOutput, error) {
    var output workmail.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailPutAccessControlRuleResult struct {
	Result workflow.Future
}

func (r *WorkmailPutAccessControlRuleResult) Get(ctx workflow.Context) (*workmail.PutAccessControlRuleOutput, error) {
    var output workmail.PutAccessControlRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailPutMailboxPermissionsResult struct {
	Result workflow.Future
}

func (r *WorkmailPutMailboxPermissionsResult) Get(ctx workflow.Context) (*workmail.PutMailboxPermissionsOutput, error) {
    var output workmail.PutMailboxPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailPutRetentionPolicyResult struct {
	Result workflow.Future
}

func (r *WorkmailPutRetentionPolicyResult) Get(ctx workflow.Context) (*workmail.PutRetentionPolicyOutput, error) {
    var output workmail.PutRetentionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailRegisterToWorkMailResult struct {
	Result workflow.Future
}

func (r *WorkmailRegisterToWorkMailResult) Get(ctx workflow.Context) (*workmail.RegisterToWorkMailOutput, error) {
    var output workmail.RegisterToWorkMailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailResetPasswordResult struct {
	Result workflow.Future
}

func (r *WorkmailResetPasswordResult) Get(ctx workflow.Context) (*workmail.ResetPasswordOutput, error) {
    var output workmail.ResetPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailTagResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailTagResourceResult) Get(ctx workflow.Context) (*workmail.TagResourceOutput, error) {
    var output workmail.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailUntagResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailUntagResourceResult) Get(ctx workflow.Context) (*workmail.UntagResourceOutput, error) {
    var output workmail.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailUpdateMailboxQuotaResult struct {
	Result workflow.Future
}

func (r *WorkmailUpdateMailboxQuotaResult) Get(ctx workflow.Context) (*workmail.UpdateMailboxQuotaOutput, error) {
    var output workmail.UpdateMailboxQuotaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailUpdatePrimaryEmailAddressResult struct {
	Result workflow.Future
}

func (r *WorkmailUpdatePrimaryEmailAddressResult) Get(ctx workflow.Context) (*workmail.UpdatePrimaryEmailAddressOutput, error) {
    var output workmail.UpdatePrimaryEmailAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkmailUpdateResourceResult struct {
	Result workflow.Future
}

func (r *WorkmailUpdateResourceResult) Get(ctx workflow.Context) (*workmail.UpdateResourceOutput, error) {
    var output workmail.UpdateResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) AssociateDelegateToResource(ctx workflow.Context, input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error) {
    var output workmail.AssociateDelegateToResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.AssociateDelegateToResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) AssociateDelegateToResourceAsync(ctx workflow.Context, input *workmail.AssociateDelegateToResourceInput) *WorkmailAssociateDelegateToResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.AssociateDelegateToResource", input)
    return &WorkmailAssociateDelegateToResourceResult{Result: future}
}

func (a *WorkMailStub) AssociateMemberToGroup(ctx workflow.Context, input *workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error) {
    var output workmail.AssociateMemberToGroupOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.AssociateMemberToGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) AssociateMemberToGroupAsync(ctx workflow.Context, input *workmail.AssociateMemberToGroupInput) *WorkmailAssociateMemberToGroupResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.AssociateMemberToGroup", input)
    return &WorkmailAssociateMemberToGroupResult{Result: future}
}

func (a *WorkMailStub) CreateAlias(ctx workflow.Context, input *workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error) {
    var output workmail.CreateAliasOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.CreateAlias", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) CreateAliasAsync(ctx workflow.Context, input *workmail.CreateAliasInput) *WorkmailCreateAliasResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.CreateAlias", input)
    return &WorkmailCreateAliasResult{Result: future}
}

func (a *WorkMailStub) CreateGroup(ctx workflow.Context, input *workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error) {
    var output workmail.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.CreateGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) CreateGroupAsync(ctx workflow.Context, input *workmail.CreateGroupInput) *WorkmailCreateGroupResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.CreateGroup", input)
    return &WorkmailCreateGroupResult{Result: future}
}

func (a *WorkMailStub) CreateResource(ctx workflow.Context, input *workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error) {
    var output workmail.CreateResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.CreateResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) CreateResourceAsync(ctx workflow.Context, input *workmail.CreateResourceInput) *WorkmailCreateResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.CreateResource", input)
    return &WorkmailCreateResourceResult{Result: future}
}

func (a *WorkMailStub) CreateUser(ctx workflow.Context, input *workmail.CreateUserInput) (*workmail.CreateUserOutput, error) {
    var output workmail.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.CreateUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) CreateUserAsync(ctx workflow.Context, input *workmail.CreateUserInput) *WorkmailCreateUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.CreateUser", input)
    return &WorkmailCreateUserResult{Result: future}
}

func (a *WorkMailStub) DeleteAccessControlRule(ctx workflow.Context, input *workmail.DeleteAccessControlRuleInput) (*workmail.DeleteAccessControlRuleOutput, error) {
    var output workmail.DeleteAccessControlRuleOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteAccessControlRule", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteAccessControlRuleAsync(ctx workflow.Context, input *workmail.DeleteAccessControlRuleInput) *WorkmailDeleteAccessControlRuleResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteAccessControlRule", input)
    return &WorkmailDeleteAccessControlRuleResult{Result: future}
}

func (a *WorkMailStub) DeleteAlias(ctx workflow.Context, input *workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error) {
    var output workmail.DeleteAliasOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteAlias", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteAliasAsync(ctx workflow.Context, input *workmail.DeleteAliasInput) *WorkmailDeleteAliasResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteAlias", input)
    return &WorkmailDeleteAliasResult{Result: future}
}

func (a *WorkMailStub) DeleteGroup(ctx workflow.Context, input *workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error) {
    var output workmail.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteGroupAsync(ctx workflow.Context, input *workmail.DeleteGroupInput) *WorkmailDeleteGroupResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteGroup", input)
    return &WorkmailDeleteGroupResult{Result: future}
}

func (a *WorkMailStub) DeleteMailboxPermissions(ctx workflow.Context, input *workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error) {
    var output workmail.DeleteMailboxPermissionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteMailboxPermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteMailboxPermissionsAsync(ctx workflow.Context, input *workmail.DeleteMailboxPermissionsInput) *WorkmailDeleteMailboxPermissionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteMailboxPermissions", input)
    return &WorkmailDeleteMailboxPermissionsResult{Result: future}
}

func (a *WorkMailStub) DeleteResource(ctx workflow.Context, input *workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error) {
    var output workmail.DeleteResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteResourceAsync(ctx workflow.Context, input *workmail.DeleteResourceInput) *WorkmailDeleteResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteResource", input)
    return &WorkmailDeleteResourceResult{Result: future}
}

func (a *WorkMailStub) DeleteRetentionPolicy(ctx workflow.Context, input *workmail.DeleteRetentionPolicyInput) (*workmail.DeleteRetentionPolicyOutput, error) {
    var output workmail.DeleteRetentionPolicyOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteRetentionPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteRetentionPolicyAsync(ctx workflow.Context, input *workmail.DeleteRetentionPolicyInput) *WorkmailDeleteRetentionPolicyResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteRetentionPolicy", input)
    return &WorkmailDeleteRetentionPolicyResult{Result: future}
}

func (a *WorkMailStub) DeleteUser(ctx workflow.Context, input *workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error) {
    var output workmail.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeleteUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeleteUserAsync(ctx workflow.Context, input *workmail.DeleteUserInput) *WorkmailDeleteUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeleteUser", input)
    return &WorkmailDeleteUserResult{Result: future}
}

func (a *WorkMailStub) DeregisterFromWorkMail(ctx workflow.Context, input *workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error) {
    var output workmail.DeregisterFromWorkMailOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DeregisterFromWorkMail", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DeregisterFromWorkMailAsync(ctx workflow.Context, input *workmail.DeregisterFromWorkMailInput) *WorkmailDeregisterFromWorkMailResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DeregisterFromWorkMail", input)
    return &WorkmailDeregisterFromWorkMailResult{Result: future}
}

func (a *WorkMailStub) DescribeGroup(ctx workflow.Context, input *workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error) {
    var output workmail.DescribeGroupOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DescribeGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DescribeGroupAsync(ctx workflow.Context, input *workmail.DescribeGroupInput) *WorkmailDescribeGroupResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DescribeGroup", input)
    return &WorkmailDescribeGroupResult{Result: future}
}

func (a *WorkMailStub) DescribeOrganization(ctx workflow.Context, input *workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error) {
    var output workmail.DescribeOrganizationOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DescribeOrganization", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DescribeOrganizationAsync(ctx workflow.Context, input *workmail.DescribeOrganizationInput) *WorkmailDescribeOrganizationResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DescribeOrganization", input)
    return &WorkmailDescribeOrganizationResult{Result: future}
}

func (a *WorkMailStub) DescribeResource(ctx workflow.Context, input *workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error) {
    var output workmail.DescribeResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DescribeResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DescribeResourceAsync(ctx workflow.Context, input *workmail.DescribeResourceInput) *WorkmailDescribeResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DescribeResource", input)
    return &WorkmailDescribeResourceResult{Result: future}
}

func (a *WorkMailStub) DescribeUser(ctx workflow.Context, input *workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error) {
    var output workmail.DescribeUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DescribeUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DescribeUserAsync(ctx workflow.Context, input *workmail.DescribeUserInput) *WorkmailDescribeUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DescribeUser", input)
    return &WorkmailDescribeUserResult{Result: future}
}

func (a *WorkMailStub) DisassociateDelegateFromResource(ctx workflow.Context, input *workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error) {
    var output workmail.DisassociateDelegateFromResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DisassociateDelegateFromResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DisassociateDelegateFromResourceAsync(ctx workflow.Context, input *workmail.DisassociateDelegateFromResourceInput) *WorkmailDisassociateDelegateFromResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DisassociateDelegateFromResource", input)
    return &WorkmailDisassociateDelegateFromResourceResult{Result: future}
}

func (a *WorkMailStub) DisassociateMemberFromGroup(ctx workflow.Context, input *workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error) {
    var output workmail.DisassociateMemberFromGroupOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.DisassociateMemberFromGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) DisassociateMemberFromGroupAsync(ctx workflow.Context, input *workmail.DisassociateMemberFromGroupInput) *WorkmailDisassociateMemberFromGroupResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.DisassociateMemberFromGroup", input)
    return &WorkmailDisassociateMemberFromGroupResult{Result: future}
}

func (a *WorkMailStub) GetAccessControlEffect(ctx workflow.Context, input *workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error) {
    var output workmail.GetAccessControlEffectOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.GetAccessControlEffect", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) GetAccessControlEffectAsync(ctx workflow.Context, input *workmail.GetAccessControlEffectInput) *WorkmailGetAccessControlEffectResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.GetAccessControlEffect", input)
    return &WorkmailGetAccessControlEffectResult{Result: future}
}

func (a *WorkMailStub) GetDefaultRetentionPolicy(ctx workflow.Context, input *workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error) {
    var output workmail.GetDefaultRetentionPolicyOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.GetDefaultRetentionPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) GetDefaultRetentionPolicyAsync(ctx workflow.Context, input *workmail.GetDefaultRetentionPolicyInput) *WorkmailGetDefaultRetentionPolicyResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.GetDefaultRetentionPolicy", input)
    return &WorkmailGetDefaultRetentionPolicyResult{Result: future}
}

func (a *WorkMailStub) GetMailboxDetails(ctx workflow.Context, input *workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error) {
    var output workmail.GetMailboxDetailsOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.GetMailboxDetails", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) GetMailboxDetailsAsync(ctx workflow.Context, input *workmail.GetMailboxDetailsInput) *WorkmailGetMailboxDetailsResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.GetMailboxDetails", input)
    return &WorkmailGetMailboxDetailsResult{Result: future}
}

func (a *WorkMailStub) ListAccessControlRules(ctx workflow.Context, input *workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error) {
    var output workmail.ListAccessControlRulesOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListAccessControlRules", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListAccessControlRulesAsync(ctx workflow.Context, input *workmail.ListAccessControlRulesInput) *WorkmailListAccessControlRulesResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListAccessControlRules", input)
    return &WorkmailListAccessControlRulesResult{Result: future}
}

func (a *WorkMailStub) ListAliases(ctx workflow.Context, input *workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error) {
    var output workmail.ListAliasesOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListAliases", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListAliasesAsync(ctx workflow.Context, input *workmail.ListAliasesInput) *WorkmailListAliasesResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListAliases", input)
    return &WorkmailListAliasesResult{Result: future}
}

func (a *WorkMailStub) ListGroupMembers(ctx workflow.Context, input *workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error) {
    var output workmail.ListGroupMembersOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListGroupMembers", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListGroupMembersAsync(ctx workflow.Context, input *workmail.ListGroupMembersInput) *WorkmailListGroupMembersResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListGroupMembers", input)
    return &WorkmailListGroupMembersResult{Result: future}
}

func (a *WorkMailStub) ListGroups(ctx workflow.Context, input *workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error) {
    var output workmail.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListGroupsAsync(ctx workflow.Context, input *workmail.ListGroupsInput) *WorkmailListGroupsResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListGroups", input)
    return &WorkmailListGroupsResult{Result: future}
}

func (a *WorkMailStub) ListMailboxPermissions(ctx workflow.Context, input *workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error) {
    var output workmail.ListMailboxPermissionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListMailboxPermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListMailboxPermissionsAsync(ctx workflow.Context, input *workmail.ListMailboxPermissionsInput) *WorkmailListMailboxPermissionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListMailboxPermissions", input)
    return &WorkmailListMailboxPermissionsResult{Result: future}
}

func (a *WorkMailStub) ListOrganizations(ctx workflow.Context, input *workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error) {
    var output workmail.ListOrganizationsOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListOrganizations", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListOrganizationsAsync(ctx workflow.Context, input *workmail.ListOrganizationsInput) *WorkmailListOrganizationsResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListOrganizations", input)
    return &WorkmailListOrganizationsResult{Result: future}
}

func (a *WorkMailStub) ListResourceDelegates(ctx workflow.Context, input *workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error) {
    var output workmail.ListResourceDelegatesOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListResourceDelegates", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListResourceDelegatesAsync(ctx workflow.Context, input *workmail.ListResourceDelegatesInput) *WorkmailListResourceDelegatesResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListResourceDelegates", input)
    return &WorkmailListResourceDelegatesResult{Result: future}
}

func (a *WorkMailStub) ListResources(ctx workflow.Context, input *workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error) {
    var output workmail.ListResourcesOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListResources", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListResourcesAsync(ctx workflow.Context, input *workmail.ListResourcesInput) *WorkmailListResourcesResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListResources", input)
    return &WorkmailListResourcesResult{Result: future}
}

func (a *WorkMailStub) ListTagsForResource(ctx workflow.Context, input *workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error) {
    var output workmail.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListTagsForResourceAsync(ctx workflow.Context, input *workmail.ListTagsForResourceInput) *WorkmailListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListTagsForResource", input)
    return &WorkmailListTagsForResourceResult{Result: future}
}

func (a *WorkMailStub) ListUsers(ctx workflow.Context, input *workmail.ListUsersInput) (*workmail.ListUsersOutput, error) {
    var output workmail.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ListUsers", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ListUsersAsync(ctx workflow.Context, input *workmail.ListUsersInput) *WorkmailListUsersResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ListUsers", input)
    return &WorkmailListUsersResult{Result: future}
}

func (a *WorkMailStub) PutAccessControlRule(ctx workflow.Context, input *workmail.PutAccessControlRuleInput) (*workmail.PutAccessControlRuleOutput, error) {
    var output workmail.PutAccessControlRuleOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.PutAccessControlRule", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) PutAccessControlRuleAsync(ctx workflow.Context, input *workmail.PutAccessControlRuleInput) *WorkmailPutAccessControlRuleResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.PutAccessControlRule", input)
    return &WorkmailPutAccessControlRuleResult{Result: future}
}

func (a *WorkMailStub) PutMailboxPermissions(ctx workflow.Context, input *workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error) {
    var output workmail.PutMailboxPermissionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.PutMailboxPermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) PutMailboxPermissionsAsync(ctx workflow.Context, input *workmail.PutMailboxPermissionsInput) *WorkmailPutMailboxPermissionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.PutMailboxPermissions", input)
    return &WorkmailPutMailboxPermissionsResult{Result: future}
}

func (a *WorkMailStub) PutRetentionPolicy(ctx workflow.Context, input *workmail.PutRetentionPolicyInput) (*workmail.PutRetentionPolicyOutput, error) {
    var output workmail.PutRetentionPolicyOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.PutRetentionPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) PutRetentionPolicyAsync(ctx workflow.Context, input *workmail.PutRetentionPolicyInput) *WorkmailPutRetentionPolicyResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.PutRetentionPolicy", input)
    return &WorkmailPutRetentionPolicyResult{Result: future}
}

func (a *WorkMailStub) RegisterToWorkMail(ctx workflow.Context, input *workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error) {
    var output workmail.RegisterToWorkMailOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.RegisterToWorkMail", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) RegisterToWorkMailAsync(ctx workflow.Context, input *workmail.RegisterToWorkMailInput) *WorkmailRegisterToWorkMailResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.RegisterToWorkMail", input)
    return &WorkmailRegisterToWorkMailResult{Result: future}
}

func (a *WorkMailStub) ResetPassword(ctx workflow.Context, input *workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error) {
    var output workmail.ResetPasswordOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.ResetPassword", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) ResetPasswordAsync(ctx workflow.Context, input *workmail.ResetPasswordInput) *WorkmailResetPasswordResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.ResetPassword", input)
    return &WorkmailResetPasswordResult{Result: future}
}

func (a *WorkMailStub) TagResource(ctx workflow.Context, input *workmail.TagResourceInput) (*workmail.TagResourceOutput, error) {
    var output workmail.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) TagResourceAsync(ctx workflow.Context, input *workmail.TagResourceInput) *WorkmailTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.TagResource", input)
    return &WorkmailTagResourceResult{Result: future}
}

func (a *WorkMailStub) UntagResource(ctx workflow.Context, input *workmail.UntagResourceInput) (*workmail.UntagResourceOutput, error) {
    var output workmail.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) UntagResourceAsync(ctx workflow.Context, input *workmail.UntagResourceInput) *WorkmailUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.UntagResource", input)
    return &WorkmailUntagResourceResult{Result: future}
}

func (a *WorkMailStub) UpdateMailboxQuota(ctx workflow.Context, input *workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error) {
    var output workmail.UpdateMailboxQuotaOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.UpdateMailboxQuota", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) UpdateMailboxQuotaAsync(ctx workflow.Context, input *workmail.UpdateMailboxQuotaInput) *WorkmailUpdateMailboxQuotaResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.UpdateMailboxQuota", input)
    return &WorkmailUpdateMailboxQuotaResult{Result: future}
}

func (a *WorkMailStub) UpdatePrimaryEmailAddress(ctx workflow.Context, input *workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error) {
    var output workmail.UpdatePrimaryEmailAddressOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.UpdatePrimaryEmailAddress", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) UpdatePrimaryEmailAddressAsync(ctx workflow.Context, input *workmail.UpdatePrimaryEmailAddressInput) *WorkmailUpdatePrimaryEmailAddressResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.UpdatePrimaryEmailAddress", input)
    return &WorkmailUpdatePrimaryEmailAddressResult{Result: future}
}

func (a *WorkMailStub) UpdateResource(ctx workflow.Context, input *workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error) {
    var output workmail.UpdateResourceOutput
    err := workflow.ExecuteActivity(ctx, "WorkMail.UpdateResource", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailStub) UpdateResourceAsync(ctx workflow.Context, input *workmail.UpdateResourceInput) *WorkmailUpdateResourceResult {
    future := workflow.ExecuteActivity(ctx, "WorkMail.UpdateResource", input)
    return &WorkmailUpdateResourceResult{Result: future}
}
