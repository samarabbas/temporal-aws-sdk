
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/aws/aws-sdk-go/service/ssoadmin/ssoadminiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SSOAdminActivities struct {
    client ssoadminiface.SSOAdminAPI
}

func NewSSOAdminActivities(session *session.Session, config ...*aws.Config) *SSOAdminActivities {
    client := ssoadmin.New(session, config...)
    return &SSOAdminActivities{client: client}
}

func (a *SSOAdminActivities) AttachManagedPolicyToPermissionSet(ctx context.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error) {
    return a.client.AttachManagedPolicyToPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) CreateAccountAssignment(ctx context.Context, input *ssoadmin.CreateAccountAssignmentInput) (*ssoadmin.CreateAccountAssignmentOutput, error) {
    return a.client.CreateAccountAssignmentWithContext(ctx, input)
}

func (a *SSOAdminActivities) CreatePermissionSet(ctx context.Context, input *ssoadmin.CreatePermissionSetInput) (*ssoadmin.CreatePermissionSetOutput, error) {
    return a.client.CreatePermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) DeleteAccountAssignment(ctx context.Context, input *ssoadmin.DeleteAccountAssignmentInput) (*ssoadmin.DeleteAccountAssignmentOutput, error) {
    return a.client.DeleteAccountAssignmentWithContext(ctx, input)
}

func (a *SSOAdminActivities) DeleteInlinePolicyFromPermissionSet(ctx context.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error) {
    return a.client.DeleteInlinePolicyFromPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) DeletePermissionSet(ctx context.Context, input *ssoadmin.DeletePermissionSetInput) (*ssoadmin.DeletePermissionSetOutput, error) {
    return a.client.DeletePermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) DescribeAccountAssignmentCreationStatus(ctx context.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
    return a.client.DescribeAccountAssignmentCreationStatusWithContext(ctx, input)
}

func (a *SSOAdminActivities) DescribeAccountAssignmentDeletionStatus(ctx context.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
    return a.client.DescribeAccountAssignmentDeletionStatusWithContext(ctx, input)
}

func (a *SSOAdminActivities) DescribePermissionSet(ctx context.Context, input *ssoadmin.DescribePermissionSetInput) (*ssoadmin.DescribePermissionSetOutput, error) {
    return a.client.DescribePermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) DescribePermissionSetProvisioningStatus(ctx context.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
    return a.client.DescribePermissionSetProvisioningStatusWithContext(ctx, input)
}

func (a *SSOAdminActivities) DetachManagedPolicyFromPermissionSet(ctx context.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error) {
    return a.client.DetachManagedPolicyFromPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) GetInlinePolicyForPermissionSet(ctx context.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
    return a.client.GetInlinePolicyForPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListAccountAssignmentCreationStatus(ctx context.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
    return a.client.ListAccountAssignmentCreationStatusWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListAccountAssignmentDeletionStatus(ctx context.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
    return a.client.ListAccountAssignmentDeletionStatusWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListAccountAssignments(ctx context.Context, input *ssoadmin.ListAccountAssignmentsInput) (*ssoadmin.ListAccountAssignmentsOutput, error) {
    return a.client.ListAccountAssignmentsWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListAccountsForProvisionedPermissionSet(ctx context.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
    return a.client.ListAccountsForProvisionedPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListInstances(ctx context.Context, input *ssoadmin.ListInstancesInput) (*ssoadmin.ListInstancesOutput, error) {
    return a.client.ListInstancesWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListManagedPoliciesInPermissionSet(ctx context.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
    return a.client.ListManagedPoliciesInPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListPermissionSetProvisioningStatus(ctx context.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
    return a.client.ListPermissionSetProvisioningStatusWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListPermissionSets(ctx context.Context, input *ssoadmin.ListPermissionSetsInput) (*ssoadmin.ListPermissionSetsOutput, error) {
    return a.client.ListPermissionSetsWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListPermissionSetsProvisionedToAccount(ctx context.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
    return a.client.ListPermissionSetsProvisionedToAccountWithContext(ctx, input)
}

func (a *SSOAdminActivities) ListTagsForResource(ctx context.Context, input *ssoadmin.ListTagsForResourceInput) (*ssoadmin.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SSOAdminActivities) ProvisionPermissionSet(ctx context.Context, input *ssoadmin.ProvisionPermissionSetInput) (*ssoadmin.ProvisionPermissionSetOutput, error) {
    return a.client.ProvisionPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) PutInlinePolicyToPermissionSet(ctx context.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error) {
    return a.client.PutInlinePolicyToPermissionSetWithContext(ctx, input)
}

func (a *SSOAdminActivities) TagResource(ctx context.Context, input *ssoadmin.TagResourceInput) (*ssoadmin.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *SSOAdminActivities) UntagResource(ctx context.Context, input *ssoadmin.UntagResourceInput) (*ssoadmin.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SSOAdminActivities) UpdatePermissionSet(ctx context.Context, input *ssoadmin.UpdatePermissionSetInput) (*ssoadmin.UpdatePermissionSetOutput, error) {
    return a.client.UpdatePermissionSetWithContext(ctx, input)
}
