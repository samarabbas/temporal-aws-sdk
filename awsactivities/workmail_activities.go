
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workmail/workmailiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type WorkMailActivities struct {
    client workmailiface.WorkMailAPI
}

func NewWorkMailActivities(session *session.Session, config ...*aws.Config) *WorkMailActivities {
    client := workmail.New(session, config...)
    return &WorkMailActivities{client: client}
}

func (a *WorkMailActivities) AssociateDelegateToResource(ctx context.Context, input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error) {
    return a.client.AssociateDelegateToResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) AssociateMemberToGroup(ctx context.Context, input *workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error) {
    return a.client.AssociateMemberToGroupWithContext(ctx, input)
}

func (a *WorkMailActivities) CreateAlias(ctx context.Context, input *workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error) {
    return a.client.CreateAliasWithContext(ctx, input)
}

func (a *WorkMailActivities) CreateGroup(ctx context.Context, input *workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error) {
    return a.client.CreateGroupWithContext(ctx, input)
}

func (a *WorkMailActivities) CreateResource(ctx context.Context, input *workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error) {
    return a.client.CreateResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) CreateUser(ctx context.Context, input *workmail.CreateUserInput) (*workmail.CreateUserOutput, error) {
    return a.client.CreateUserWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteAccessControlRule(ctx context.Context, input *workmail.DeleteAccessControlRuleInput) (*workmail.DeleteAccessControlRuleOutput, error) {
    return a.client.DeleteAccessControlRuleWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteAlias(ctx context.Context, input *workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error) {
    return a.client.DeleteAliasWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteGroup(ctx context.Context, input *workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error) {
    return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteMailboxPermissions(ctx context.Context, input *workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error) {
    return a.client.DeleteMailboxPermissionsWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteResource(ctx context.Context, input *workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error) {
    return a.client.DeleteResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteRetentionPolicy(ctx context.Context, input *workmail.DeleteRetentionPolicyInput) (*workmail.DeleteRetentionPolicyOutput, error) {
    return a.client.DeleteRetentionPolicyWithContext(ctx, input)
}

func (a *WorkMailActivities) DeleteUser(ctx context.Context, input *workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error) {
    return a.client.DeleteUserWithContext(ctx, input)
}

func (a *WorkMailActivities) DeregisterFromWorkMail(ctx context.Context, input *workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error) {
    return a.client.DeregisterFromWorkMailWithContext(ctx, input)
}

func (a *WorkMailActivities) DescribeGroup(ctx context.Context, input *workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error) {
    return a.client.DescribeGroupWithContext(ctx, input)
}

func (a *WorkMailActivities) DescribeOrganization(ctx context.Context, input *workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error) {
    return a.client.DescribeOrganizationWithContext(ctx, input)
}

func (a *WorkMailActivities) DescribeResource(ctx context.Context, input *workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error) {
    return a.client.DescribeResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) DescribeUser(ctx context.Context, input *workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error) {
    return a.client.DescribeUserWithContext(ctx, input)
}

func (a *WorkMailActivities) DisassociateDelegateFromResource(ctx context.Context, input *workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error) {
    return a.client.DisassociateDelegateFromResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) DisassociateMemberFromGroup(ctx context.Context, input *workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error) {
    return a.client.DisassociateMemberFromGroupWithContext(ctx, input)
}

func (a *WorkMailActivities) GetAccessControlEffect(ctx context.Context, input *workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error) {
    return a.client.GetAccessControlEffectWithContext(ctx, input)
}

func (a *WorkMailActivities) GetDefaultRetentionPolicy(ctx context.Context, input *workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error) {
    return a.client.GetDefaultRetentionPolicyWithContext(ctx, input)
}

func (a *WorkMailActivities) GetMailboxDetails(ctx context.Context, input *workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error) {
    return a.client.GetMailboxDetailsWithContext(ctx, input)
}

func (a *WorkMailActivities) ListAccessControlRules(ctx context.Context, input *workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error) {
    return a.client.ListAccessControlRulesWithContext(ctx, input)
}

func (a *WorkMailActivities) ListAliases(ctx context.Context, input *workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error) {
    return a.client.ListAliasesWithContext(ctx, input)
}

func (a *WorkMailActivities) ListGroupMembers(ctx context.Context, input *workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error) {
    return a.client.ListGroupMembersWithContext(ctx, input)
}

func (a *WorkMailActivities) ListGroups(ctx context.Context, input *workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error) {
    return a.client.ListGroupsWithContext(ctx, input)
}

func (a *WorkMailActivities) ListMailboxPermissions(ctx context.Context, input *workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error) {
    return a.client.ListMailboxPermissionsWithContext(ctx, input)
}

func (a *WorkMailActivities) ListOrganizations(ctx context.Context, input *workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error) {
    return a.client.ListOrganizationsWithContext(ctx, input)
}

func (a *WorkMailActivities) ListResourceDelegates(ctx context.Context, input *workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error) {
    return a.client.ListResourceDelegatesWithContext(ctx, input)
}

func (a *WorkMailActivities) ListResources(ctx context.Context, input *workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error) {
    return a.client.ListResourcesWithContext(ctx, input)
}

func (a *WorkMailActivities) ListTagsForResource(ctx context.Context, input *workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) ListUsers(ctx context.Context, input *workmail.ListUsersInput) (*workmail.ListUsersOutput, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *WorkMailActivities) PutAccessControlRule(ctx context.Context, input *workmail.PutAccessControlRuleInput) (*workmail.PutAccessControlRuleOutput, error) {
    return a.client.PutAccessControlRuleWithContext(ctx, input)
}

func (a *WorkMailActivities) PutMailboxPermissions(ctx context.Context, input *workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error) {
    return a.client.PutMailboxPermissionsWithContext(ctx, input)
}

func (a *WorkMailActivities) PutRetentionPolicy(ctx context.Context, input *workmail.PutRetentionPolicyInput) (*workmail.PutRetentionPolicyOutput, error) {
    return a.client.PutRetentionPolicyWithContext(ctx, input)
}

func (a *WorkMailActivities) RegisterToWorkMail(ctx context.Context, input *workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error) {
    return a.client.RegisterToWorkMailWithContext(ctx, input)
}

func (a *WorkMailActivities) ResetPassword(ctx context.Context, input *workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error) {
    return a.client.ResetPasswordWithContext(ctx, input)
}

func (a *WorkMailActivities) TagResource(ctx context.Context, input *workmail.TagResourceInput) (*workmail.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) UntagResource(ctx context.Context, input *workmail.UntagResourceInput) (*workmail.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *WorkMailActivities) UpdateMailboxQuota(ctx context.Context, input *workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error) {
    return a.client.UpdateMailboxQuotaWithContext(ctx, input)
}

func (a *WorkMailActivities) UpdatePrimaryEmailAddress(ctx context.Context, input *workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error) {
    return a.client.UpdatePrimaryEmailAddressWithContext(ctx, input)
}

func (a *WorkMailActivities) UpdateResource(ctx context.Context, input *workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error) {
    return a.client.UpdateResourceWithContext(ctx, input)
}
