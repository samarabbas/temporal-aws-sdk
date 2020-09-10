package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/detective/detectiveiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type DetectiveActivities struct {
	client detectiveiface.DetectiveAPI
}

func NewDetectiveActivities(session *session.Session, config ...*aws.Config) *DetectiveActivities {
	client := detective.New(session, config...)
	return &DetectiveActivities{client: client}
}

func (a *DetectiveActivities) AcceptInvitation(ctx context.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
	return a.client.AcceptInvitationWithContext(ctx, input)
}

func (a *DetectiveActivities) CreateGraph(ctx context.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error) {
	return a.client.CreateGraphWithContext(ctx, input)
}

func (a *DetectiveActivities) CreateMembers(ctx context.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error) {
	return a.client.CreateMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) DeleteGraph(ctx context.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error) {
	return a.client.DeleteGraphWithContext(ctx, input)
}

func (a *DetectiveActivities) DeleteMembers(ctx context.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error) {
	return a.client.DeleteMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) DisassociateMembership(ctx context.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error) {
	return a.client.DisassociateMembershipWithContext(ctx, input)
}

func (a *DetectiveActivities) GetMembers(ctx context.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
	return a.client.GetMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) ListGraphs(ctx context.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
	return a.client.ListGraphsWithContext(ctx, input)
}

func (a *DetectiveActivities) ListInvitations(ctx context.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
	return a.client.ListInvitationsWithContext(ctx, input)
}

func (a *DetectiveActivities) ListMembers(ctx context.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
	return a.client.ListMembersWithContext(ctx, input)
}

func (a *DetectiveActivities) RejectInvitation(ctx context.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error) {
	return a.client.RejectInvitationWithContext(ctx, input)
}

func (a *DetectiveActivities) StartMonitoringMember(ctx context.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error) {
	return a.client.StartMonitoringMemberWithContext(ctx, input)
}
