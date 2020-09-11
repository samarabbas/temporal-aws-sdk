
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SNSActivities struct {
    client snsiface.SNSAPI
}

func NewSNSActivities(session *session.Session, config ...*aws.Config) *SNSActivities {
    client := sns.New(session, config...)
    return &SNSActivities{client: client}
}

func (a *SNSActivities) AddPermission(ctx context.Context, input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
    return a.client.AddPermissionWithContext(ctx, input)
}

func (a *SNSActivities) CheckIfPhoneNumberIsOptedOut(ctx context.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
    return a.client.CheckIfPhoneNumberIsOptedOutWithContext(ctx, input)
}

func (a *SNSActivities) ConfirmSubscription(ctx context.Context, input *sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error) {
    return a.client.ConfirmSubscriptionWithContext(ctx, input)
}

func (a *SNSActivities) CreatePlatformApplication(ctx context.Context, input *sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error) {
    return a.client.CreatePlatformApplicationWithContext(ctx, input)
}

func (a *SNSActivities) CreatePlatformEndpoint(ctx context.Context, input *sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error) {
    return a.client.CreatePlatformEndpointWithContext(ctx, input)
}

func (a *SNSActivities) CreateTopic(ctx context.Context, input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
    return a.client.CreateTopicWithContext(ctx, input)
}

func (a *SNSActivities) DeleteEndpoint(ctx context.Context, input *sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error) {
    return a.client.DeleteEndpointWithContext(ctx, input)
}

func (a *SNSActivities) DeletePlatformApplication(ctx context.Context, input *sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error) {
    return a.client.DeletePlatformApplicationWithContext(ctx, input)
}

func (a *SNSActivities) DeleteTopic(ctx context.Context, input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error) {
    return a.client.DeleteTopicWithContext(ctx, input)
}

func (a *SNSActivities) GetEndpointAttributes(ctx context.Context, input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error) {
    return a.client.GetEndpointAttributesWithContext(ctx, input)
}

func (a *SNSActivities) GetPlatformApplicationAttributes(ctx context.Context, input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error) {
    return a.client.GetPlatformApplicationAttributesWithContext(ctx, input)
}

func (a *SNSActivities) GetSMSAttributes(ctx context.Context, input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error) {
    return a.client.GetSMSAttributesWithContext(ctx, input)
}

func (a *SNSActivities) GetSubscriptionAttributes(ctx context.Context, input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error) {
    return a.client.GetSubscriptionAttributesWithContext(ctx, input)
}

func (a *SNSActivities) GetTopicAttributes(ctx context.Context, input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error) {
    return a.client.GetTopicAttributesWithContext(ctx, input)
}

func (a *SNSActivities) ListEndpointsByPlatformApplication(ctx context.Context, input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
    return a.client.ListEndpointsByPlatformApplicationWithContext(ctx, input)
}

func (a *SNSActivities) ListPhoneNumbersOptedOut(ctx context.Context, input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error) {
    return a.client.ListPhoneNumbersOptedOutWithContext(ctx, input)
}

func (a *SNSActivities) ListPlatformApplications(ctx context.Context, input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error) {
    return a.client.ListPlatformApplicationsWithContext(ctx, input)
}

func (a *SNSActivities) ListSubscriptions(ctx context.Context, input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error) {
    return a.client.ListSubscriptionsWithContext(ctx, input)
}

func (a *SNSActivities) ListSubscriptionsByTopic(ctx context.Context, input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
    return a.client.ListSubscriptionsByTopicWithContext(ctx, input)
}

func (a *SNSActivities) ListTagsForResource(ctx context.Context, input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SNSActivities) ListTopics(ctx context.Context, input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
    return a.client.ListTopicsWithContext(ctx, input)
}

func (a *SNSActivities) OptInPhoneNumber(ctx context.Context, input *sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error) {
    return a.client.OptInPhoneNumberWithContext(ctx, input)
}

func (a *SNSActivities) Publish(ctx context.Context, input *sns.PublishInput) (*sns.PublishOutput, error) {
    return a.client.PublishWithContext(ctx, input)
}

func (a *SNSActivities) RemovePermission(ctx context.Context, input *sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error) {
    return a.client.RemovePermissionWithContext(ctx, input)
}

func (a *SNSActivities) SetEndpointAttributes(ctx context.Context, input *sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error) {
    return a.client.SetEndpointAttributesWithContext(ctx, input)
}

func (a *SNSActivities) SetPlatformApplicationAttributes(ctx context.Context, input *sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error) {
    return a.client.SetPlatformApplicationAttributesWithContext(ctx, input)
}

func (a *SNSActivities) SetSMSAttributes(ctx context.Context, input *sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error) {
    return a.client.SetSMSAttributesWithContext(ctx, input)
}

func (a *SNSActivities) SetSubscriptionAttributes(ctx context.Context, input *sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error) {
    return a.client.SetSubscriptionAttributesWithContext(ctx, input)
}

func (a *SNSActivities) SetTopicAttributes(ctx context.Context, input *sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error) {
    return a.client.SetTopicAttributesWithContext(ctx, input)
}

func (a *SNSActivities) Subscribe(ctx context.Context, input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
    return a.client.SubscribeWithContext(ctx, input)
}

func (a *SNSActivities) TagResource(ctx context.Context, input *sns.TagResourceInput) (*sns.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *SNSActivities) Unsubscribe(ctx context.Context, input *sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error) {
    return a.client.UnsubscribeWithContext(ctx, input)
}

func (a *SNSActivities) UntagResource(ctx context.Context, input *sns.UntagResourceInput) (*sns.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}
