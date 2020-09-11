
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type KinesisActivities struct {
    client kinesisiface.KinesisAPI
}

func NewKinesisActivities(session *session.Session, config ...*aws.Config) *KinesisActivities {
    client := kinesis.New(session, config...)
    return &KinesisActivities{client: client}
}

func (a *KinesisActivities) AddTagsToStream(ctx context.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
    return a.client.AddTagsToStreamWithContext(ctx, input)
}

func (a *KinesisActivities) CreateStream(ctx context.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
    return a.client.CreateStreamWithContext(ctx, input)
}

func (a *KinesisActivities) DecreaseStreamRetentionPeriod(ctx context.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
    return a.client.DecreaseStreamRetentionPeriodWithContext(ctx, input)
}

func (a *KinesisActivities) DeleteStream(ctx context.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
    return a.client.DeleteStreamWithContext(ctx, input)
}

func (a *KinesisActivities) DeregisterStreamConsumer(ctx context.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
    return a.client.DeregisterStreamConsumerWithContext(ctx, input)
}

func (a *KinesisActivities) DescribeLimits(ctx context.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
    return a.client.DescribeLimitsWithContext(ctx, input)
}

func (a *KinesisActivities) DescribeStream(ctx context.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
    return a.client.DescribeStreamWithContext(ctx, input)
}

func (a *KinesisActivities) DescribeStreamConsumer(ctx context.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
    return a.client.DescribeStreamConsumerWithContext(ctx, input)
}

func (a *KinesisActivities) DescribeStreamSummary(ctx context.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
    return a.client.DescribeStreamSummaryWithContext(ctx, input)
}

func (a *KinesisActivities) DisableEnhancedMonitoring(ctx context.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
    return a.client.DisableEnhancedMonitoringWithContext(ctx, input)
}

func (a *KinesisActivities) EnableEnhancedMonitoring(ctx context.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
    return a.client.EnableEnhancedMonitoringWithContext(ctx, input)
}

func (a *KinesisActivities) GetRecords(ctx context.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
    return a.client.GetRecordsWithContext(ctx, input)
}

func (a *KinesisActivities) GetShardIterator(ctx context.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
    return a.client.GetShardIteratorWithContext(ctx, input)
}

func (a *KinesisActivities) IncreaseStreamRetentionPeriod(ctx context.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
    return a.client.IncreaseStreamRetentionPeriodWithContext(ctx, input)
}

func (a *KinesisActivities) ListShards(ctx context.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
    return a.client.ListShardsWithContext(ctx, input)
}

func (a *KinesisActivities) ListStreamConsumers(ctx context.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
    return a.client.ListStreamConsumersWithContext(ctx, input)
}

func (a *KinesisActivities) ListStreams(ctx context.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
    return a.client.ListStreamsWithContext(ctx, input)
}

func (a *KinesisActivities) ListTagsForStream(ctx context.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
    return a.client.ListTagsForStreamWithContext(ctx, input)
}

func (a *KinesisActivities) MergeShards(ctx context.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
    return a.client.MergeShardsWithContext(ctx, input)
}

func (a *KinesisActivities) PutRecord(ctx context.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
    return a.client.PutRecordWithContext(ctx, input)
}

func (a *KinesisActivities) PutRecords(ctx context.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
    return a.client.PutRecordsWithContext(ctx, input)
}

func (a *KinesisActivities) RegisterStreamConsumer(ctx context.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
    return a.client.RegisterStreamConsumerWithContext(ctx, input)
}

func (a *KinesisActivities) RemoveTagsFromStream(ctx context.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
    return a.client.RemoveTagsFromStreamWithContext(ctx, input)
}

func (a *KinesisActivities) SplitShard(ctx context.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
    return a.client.SplitShardWithContext(ctx, input)
}

func (a *KinesisActivities) StartStreamEncryption(ctx context.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
    return a.client.StartStreamEncryptionWithContext(ctx, input)
}

func (a *KinesisActivities) StopStreamEncryption(ctx context.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
    return a.client.StopStreamEncryptionWithContext(ctx, input)
}

func (a *KinesisActivities) SubscribeToShard(ctx context.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
    return a.client.SubscribeToShardWithContext(ctx, input)
}

func (a *KinesisActivities) UpdateShardCount(ctx context.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
    return a.client.UpdateShardCountWithContext(ctx, input)
}

func (a *KinesisActivities) WaitUntilStreamExists(ctx context.Context, input *kinesis.DescribeStreamInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStreamExistsWithContext(ctx, input, options...)
	})
}

func (a *KinesisActivities) WaitUntilStreamNotExists(ctx context.Context, input *kinesis.DescribeStreamInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStreamNotExistsWithContext(ctx, input, options...)
	})
}
