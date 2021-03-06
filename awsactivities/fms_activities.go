// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/fms/fmsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type FMSActivities struct {
	client fmsiface.FMSAPI
}

func NewFMSActivities(session *session.Session, config ...*aws.Config) *FMSActivities {
	client := fms.New(session, config...)
	return &FMSActivities{client: client}
}

func (a *FMSActivities) AssociateAdminAccount(ctx context.Context, input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
	return a.client.AssociateAdminAccountWithContext(ctx, input)
}

func (a *FMSActivities) DeleteAppsList(ctx context.Context, input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error) {
	return a.client.DeleteAppsListWithContext(ctx, input)
}

func (a *FMSActivities) DeleteNotificationChannel(ctx context.Context, input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error) {
	return a.client.DeleteNotificationChannelWithContext(ctx, input)
}

func (a *FMSActivities) DeletePolicy(ctx context.Context, input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error) {
	return a.client.DeletePolicyWithContext(ctx, input)
}

func (a *FMSActivities) DeleteProtocolsList(ctx context.Context, input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error) {
	return a.client.DeleteProtocolsListWithContext(ctx, input)
}

func (a *FMSActivities) DisassociateAdminAccount(ctx context.Context, input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error) {
	return a.client.DisassociateAdminAccountWithContext(ctx, input)
}

func (a *FMSActivities) GetAdminAccount(ctx context.Context, input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error) {
	return a.client.GetAdminAccountWithContext(ctx, input)
}

func (a *FMSActivities) GetAppsList(ctx context.Context, input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error) {
	return a.client.GetAppsListWithContext(ctx, input)
}

func (a *FMSActivities) GetComplianceDetail(ctx context.Context, input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error) {
	return a.client.GetComplianceDetailWithContext(ctx, input)
}

func (a *FMSActivities) GetNotificationChannel(ctx context.Context, input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error) {
	return a.client.GetNotificationChannelWithContext(ctx, input)
}

func (a *FMSActivities) GetPolicy(ctx context.Context, input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error) {
	return a.client.GetPolicyWithContext(ctx, input)
}

func (a *FMSActivities) GetProtectionStatus(ctx context.Context, input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error) {
	return a.client.GetProtectionStatusWithContext(ctx, input)
}

func (a *FMSActivities) GetProtocolsList(ctx context.Context, input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error) {
	return a.client.GetProtocolsListWithContext(ctx, input)
}

func (a *FMSActivities) GetViolationDetails(ctx context.Context, input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error) {
	return a.client.GetViolationDetailsWithContext(ctx, input)
}

func (a *FMSActivities) ListAppsLists(ctx context.Context, input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error) {
	return a.client.ListAppsListsWithContext(ctx, input)
}

func (a *FMSActivities) ListComplianceStatus(ctx context.Context, input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error) {
	return a.client.ListComplianceStatusWithContext(ctx, input)
}

func (a *FMSActivities) ListMemberAccounts(ctx context.Context, input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error) {
	return a.client.ListMemberAccountsWithContext(ctx, input)
}

func (a *FMSActivities) ListPolicies(ctx context.Context, input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error) {
	return a.client.ListPoliciesWithContext(ctx, input)
}

func (a *FMSActivities) ListProtocolsLists(ctx context.Context, input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error) {
	return a.client.ListProtocolsListsWithContext(ctx, input)
}

func (a *FMSActivities) ListTagsForResource(ctx context.Context, input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *FMSActivities) PutAppsList(ctx context.Context, input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error) {
	return a.client.PutAppsListWithContext(ctx, input)
}

func (a *FMSActivities) PutNotificationChannel(ctx context.Context, input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error) {
	return a.client.PutNotificationChannelWithContext(ctx, input)
}

func (a *FMSActivities) PutPolicy(ctx context.Context, input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error) {
	return a.client.PutPolicyWithContext(ctx, input)
}

func (a *FMSActivities) PutProtocolsList(ctx context.Context, input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error) {
	return a.client.PutProtocolsListWithContext(ctx, input)
}

func (a *FMSActivities) TagResource(ctx context.Context, input *fms.TagResourceInput) (*fms.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *FMSActivities) UntagResource(ctx context.Context, input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
