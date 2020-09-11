
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideo/kinesisvideoiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type KinesisVideoActivities struct {
    client kinesisvideoiface.KinesisVideoAPI
}

func NewKinesisVideoActivities(session *session.Session, config ...*aws.Config) *KinesisVideoActivities {
    client := kinesisvideo.New(session, config...)
    return &KinesisVideoActivities{client: client}
}

func (a *KinesisVideoActivities) CreateSignalingChannel(ctx context.Context, input *kinesisvideo.CreateSignalingChannelInput) (*kinesisvideo.CreateSignalingChannelOutput, error) {
    return a.client.CreateSignalingChannelWithContext(ctx, input)
}

func (a *KinesisVideoActivities) CreateStream(ctx context.Context, input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error) {
    return a.client.CreateStreamWithContext(ctx, input)
}

func (a *KinesisVideoActivities) DeleteSignalingChannel(ctx context.Context, input *kinesisvideo.DeleteSignalingChannelInput) (*kinesisvideo.DeleteSignalingChannelOutput, error) {
    return a.client.DeleteSignalingChannelWithContext(ctx, input)
}

func (a *KinesisVideoActivities) DeleteStream(ctx context.Context, input *kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error) {
    return a.client.DeleteStreamWithContext(ctx, input)
}

func (a *KinesisVideoActivities) DescribeSignalingChannel(ctx context.Context, input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
    return a.client.DescribeSignalingChannelWithContext(ctx, input)
}

func (a *KinesisVideoActivities) DescribeStream(ctx context.Context, input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error) {
    return a.client.DescribeStreamWithContext(ctx, input)
}

func (a *KinesisVideoActivities) GetDataEndpoint(ctx context.Context, input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error) {
    return a.client.GetDataEndpointWithContext(ctx, input)
}

func (a *KinesisVideoActivities) GetSignalingChannelEndpoint(ctx context.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
    return a.client.GetSignalingChannelEndpointWithContext(ctx, input)
}

func (a *KinesisVideoActivities) ListSignalingChannels(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error) {
    return a.client.ListSignalingChannelsWithContext(ctx, input)
}

func (a *KinesisVideoActivities) ListStreams(ctx context.Context, input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error) {
    return a.client.ListStreamsWithContext(ctx, input)
}

func (a *KinesisVideoActivities) ListTagsForResource(ctx context.Context, input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KinesisVideoActivities) ListTagsForStream(ctx context.Context, input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error) {
    return a.client.ListTagsForStreamWithContext(ctx, input)
}

func (a *KinesisVideoActivities) TagResource(ctx context.Context, input *kinesisvideo.TagResourceInput) (*kinesisvideo.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *KinesisVideoActivities) TagStream(ctx context.Context, input *kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error) {
    return a.client.TagStreamWithContext(ctx, input)
}

func (a *KinesisVideoActivities) UntagResource(ctx context.Context, input *kinesisvideo.UntagResourceInput) (*kinesisvideo.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *KinesisVideoActivities) UntagStream(ctx context.Context, input *kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error) {
    return a.client.UntagStreamWithContext(ctx, input)
}

func (a *KinesisVideoActivities) UpdateDataRetention(ctx context.Context, input *kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error) {
    return a.client.UpdateDataRetentionWithContext(ctx, input)
}

func (a *KinesisVideoActivities) UpdateSignalingChannel(ctx context.Context, input *kinesisvideo.UpdateSignalingChannelInput) (*kinesisvideo.UpdateSignalingChannelOutput, error) {
    return a.client.UpdateSignalingChannelWithContext(ctx, input)
}

func (a *KinesisVideoActivities) UpdateStream(ctx context.Context, input *kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error) {
    return a.client.UpdateStreamWithContext(ctx, input)
}
