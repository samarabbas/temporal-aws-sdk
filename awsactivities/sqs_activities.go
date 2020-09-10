package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SQSActivities struct {
	client sqsiface.SQSAPI
}

func NewSQSActivities(session *session.Session, config ...*aws.Config) *SQSActivities {
	client := sqs.New(session, config...)
	return &SQSActivities{client: client}
}

func (a *SQSActivities) AddPermission(ctx context.Context, input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	return a.client.AddPermissionWithContext(ctx, input)
}

func (a *SQSActivities) ChangeMessageVisibility(ctx context.Context, input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	return a.client.ChangeMessageVisibilityWithContext(ctx, input)
}

func (a *SQSActivities) ChangeMessageVisibilityBatch(ctx context.Context, input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	return a.client.ChangeMessageVisibilityBatchWithContext(ctx, input)
}

func (a *SQSActivities) CreateQueue(ctx context.Context, input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	return a.client.CreateQueueWithContext(ctx, input)
}

func (a *SQSActivities) DeleteMessage(ctx context.Context, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return a.client.DeleteMessageWithContext(ctx, input)
}

func (a *SQSActivities) DeleteMessageBatch(ctx context.Context, input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	return a.client.DeleteMessageBatchWithContext(ctx, input)
}

func (a *SQSActivities) DeleteQueue(ctx context.Context, input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	return a.client.DeleteQueueWithContext(ctx, input)
}

func (a *SQSActivities) GetQueueAttributes(ctx context.Context, input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	return a.client.GetQueueAttributesWithContext(ctx, input)
}

func (a *SQSActivities) GetQueueUrl(ctx context.Context, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return a.client.GetQueueUrlWithContext(ctx, input)
}

func (a *SQSActivities) ListDeadLetterSourceQueues(ctx context.Context, input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	return a.client.ListDeadLetterSourceQueuesWithContext(ctx, input)
}

func (a *SQSActivities) ListQueueTags(ctx context.Context, input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	return a.client.ListQueueTagsWithContext(ctx, input)
}

func (a *SQSActivities) ListQueues(ctx context.Context, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	return a.client.ListQueuesWithContext(ctx, input)
}

func (a *SQSActivities) PurgeQueue(ctx context.Context, input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	return a.client.PurgeQueueWithContext(ctx, input)
}

func (a *SQSActivities) ReceiveMessage(ctx context.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return a.client.ReceiveMessageWithContext(ctx, input)
}

func (a *SQSActivities) RemovePermission(ctx context.Context, input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	return a.client.RemovePermissionWithContext(ctx, input)
}

func (a *SQSActivities) SendMessage(ctx context.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return a.client.SendMessageWithContext(ctx, input)
}

func (a *SQSActivities) SendMessageBatch(ctx context.Context, input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	return a.client.SendMessageBatchWithContext(ctx, input)
}

func (a *SQSActivities) SetQueueAttributes(ctx context.Context, input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	return a.client.SetQueueAttributesWithContext(ctx, input)
}

func (a *SQSActivities) TagQueue(ctx context.Context, input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
	return a.client.TagQueueWithContext(ctx, input)
}

func (a *SQSActivities) UntagQueue(ctx context.Context, input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
	return a.client.UntagQueueWithContext(ctx, input)
}
