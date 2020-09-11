
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DynamoDBStreamsActivities struct {
    client dynamodbstreamsiface.DynamoDBStreamsAPI
}

func NewDynamoDBStreamsActivities(session *session.Session, config ...*aws.Config) *DynamoDBStreamsActivities {
    client := dynamodbstreams.New(session, config...)
    return &DynamoDBStreamsActivities{client: client}
}

func (a *DynamoDBStreamsActivities) DescribeStream(ctx context.Context, input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
    return a.client.DescribeStreamWithContext(ctx, input)
}

func (a *DynamoDBStreamsActivities) GetRecords(ctx context.Context, input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
    return a.client.GetRecordsWithContext(ctx, input)
}

func (a *DynamoDBStreamsActivities) GetShardIterator(ctx context.Context, input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
    return a.client.GetShardIteratorWithContext(ctx, input)
}

func (a *DynamoDBStreamsActivities) ListStreams(ctx context.Context, input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
    return a.client.ListStreamsWithContext(ctx, input)
}
