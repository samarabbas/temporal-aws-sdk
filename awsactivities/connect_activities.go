
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ConnectActivities struct {
    client connectiface.ConnectAPI
}

func NewConnectActivities(session *session.Session, config ...*aws.Config) *ConnectActivities {
    client := connect.New(session, config...)
    return &ConnectActivities{client: client}
}

func (a *ConnectActivities) CreateUser(ctx context.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
    return a.client.CreateUserWithContext(ctx, input)
}

func (a *ConnectActivities) DeleteUser(ctx context.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error) {
    return a.client.DeleteUserWithContext(ctx, input)
}

func (a *ConnectActivities) DescribeUser(ctx context.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
    return a.client.DescribeUserWithContext(ctx, input)
}

func (a *ConnectActivities) DescribeUserHierarchyGroup(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
    return a.client.DescribeUserHierarchyGroupWithContext(ctx, input)
}

func (a *ConnectActivities) DescribeUserHierarchyStructure(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
    return a.client.DescribeUserHierarchyStructureWithContext(ctx, input)
}

func (a *ConnectActivities) GetContactAttributes(ctx context.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
    return a.client.GetContactAttributesWithContext(ctx, input)
}

func (a *ConnectActivities) GetCurrentMetricData(ctx context.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
    return a.client.GetCurrentMetricDataWithContext(ctx, input)
}

func (a *ConnectActivities) GetFederationToken(ctx context.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
    return a.client.GetFederationTokenWithContext(ctx, input)
}

func (a *ConnectActivities) GetMetricData(ctx context.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
    return a.client.GetMetricDataWithContext(ctx, input)
}

func (a *ConnectActivities) ListContactFlows(ctx context.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
    return a.client.ListContactFlowsWithContext(ctx, input)
}

func (a *ConnectActivities) ListHoursOfOperations(ctx context.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
    return a.client.ListHoursOfOperationsWithContext(ctx, input)
}

func (a *ConnectActivities) ListPhoneNumbers(ctx context.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
    return a.client.ListPhoneNumbersWithContext(ctx, input)
}

func (a *ConnectActivities) ListQueues(ctx context.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
    return a.client.ListQueuesWithContext(ctx, input)
}

func (a *ConnectActivities) ListRoutingProfiles(ctx context.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
    return a.client.ListRoutingProfilesWithContext(ctx, input)
}

func (a *ConnectActivities) ListSecurityProfiles(ctx context.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
    return a.client.ListSecurityProfilesWithContext(ctx, input)
}

func (a *ConnectActivities) ListTagsForResource(ctx context.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ConnectActivities) ListUserHierarchyGroups(ctx context.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
    return a.client.ListUserHierarchyGroupsWithContext(ctx, input)
}

func (a *ConnectActivities) ListUsers(ctx context.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *ConnectActivities) ResumeContactRecording(ctx context.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error) {
    return a.client.ResumeContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) StartChatContact(ctx context.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.StartChatContactWithContext(ctx, input)
}

func (a *ConnectActivities) StartContactRecording(ctx context.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error) {
    return a.client.StartContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) StartOutboundVoiceContact(ctx context.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.StartOutboundVoiceContactWithContext(ctx, input)
}

func (a *ConnectActivities) StopContact(ctx context.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error) {
    return a.client.StopContactWithContext(ctx, input)
}

func (a *ConnectActivities) StopContactRecording(ctx context.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error) {
    return a.client.StopContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) SuspendContactRecording(ctx context.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error) {
    return a.client.SuspendContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) TagResource(ctx context.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *ConnectActivities) UntagResource(ctx context.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateContactAttributes(ctx context.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error) {
    return a.client.UpdateContactAttributesWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserHierarchy(ctx context.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error) {
    return a.client.UpdateUserHierarchyWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserIdentityInfo(ctx context.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error) {
    return a.client.UpdateUserIdentityInfoWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserPhoneConfig(ctx context.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error) {
    return a.client.UpdateUserPhoneConfigWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserRoutingProfile(ctx context.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error) {
    return a.client.UpdateUserRoutingProfileWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserSecurityProfiles(ctx context.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error) {
    return a.client.UpdateUserSecurityProfilesWithContext(ctx, input)
}
