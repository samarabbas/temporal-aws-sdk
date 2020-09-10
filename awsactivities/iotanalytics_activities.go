package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotanalytics/iotanalyticsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type IoTAnalyticsActivities struct {
	client iotanalyticsiface.IoTAnalyticsAPI
}

func NewIoTAnalyticsActivities(session *session.Session, config ...*aws.Config) *IoTAnalyticsActivities {
	client := iotanalytics.New(session, config...)
	return &IoTAnalyticsActivities{client: client}
}

func (a *IoTAnalyticsActivities) BatchPutMessage(ctx context.Context, input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
	return a.client.BatchPutMessageWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) CancelPipelineReprocessing(ctx context.Context, input *iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error) {
	return a.client.CancelPipelineReprocessingWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) CreateChannel(ctx context.Context, input *iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error) {
	return a.client.CreateChannelWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) CreateDataset(ctx context.Context, input *iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error) {
	return a.client.CreateDatasetWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) CreateDatasetContent(ctx context.Context, input *iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error) {
	return a.client.CreateDatasetContentWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) CreateDatastore(ctx context.Context, input *iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error) {
	return a.client.CreateDatastoreWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) CreatePipeline(ctx context.Context, input *iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error) {
	return a.client.CreatePipelineWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DeleteChannel(ctx context.Context, input *iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error) {
	return a.client.DeleteChannelWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DeleteDataset(ctx context.Context, input *iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error) {
	return a.client.DeleteDatasetWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DeleteDatasetContent(ctx context.Context, input *iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error) {
	return a.client.DeleteDatasetContentWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DeleteDatastore(ctx context.Context, input *iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error) {
	return a.client.DeleteDatastoreWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DeletePipeline(ctx context.Context, input *iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error) {
	return a.client.DeletePipelineWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DescribeChannel(ctx context.Context, input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error) {
	return a.client.DescribeChannelWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DescribeDataset(ctx context.Context, input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error) {
	return a.client.DescribeDatasetWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DescribeDatastore(ctx context.Context, input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error) {
	return a.client.DescribeDatastoreWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DescribeLoggingOptions(ctx context.Context, input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
	return a.client.DescribeLoggingOptionsWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) DescribePipeline(ctx context.Context, input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error) {
	return a.client.DescribePipelineWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) GetDatasetContent(ctx context.Context, input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error) {
	return a.client.GetDatasetContentWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) ListChannels(ctx context.Context, input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error) {
	return a.client.ListChannelsWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) ListDatasetContents(ctx context.Context, input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error) {
	return a.client.ListDatasetContentsWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) ListDatasets(ctx context.Context, input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error) {
	return a.client.ListDatasetsWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) ListDatastores(ctx context.Context, input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error) {
	return a.client.ListDatastoresWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) ListPipelines(ctx context.Context, input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error) {
	return a.client.ListPipelinesWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) ListTagsForResource(ctx context.Context, input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) PutLoggingOptions(ctx context.Context, input *iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error) {
	return a.client.PutLoggingOptionsWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) RunPipelineActivity(ctx context.Context, input *iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error) {
	return a.client.RunPipelineActivityWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) SampleChannelData(ctx context.Context, input *iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error) {
	return a.client.SampleChannelDataWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) StartPipelineReprocessing(ctx context.Context, input *iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error) {
	return a.client.StartPipelineReprocessingWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) TagResource(ctx context.Context, input *iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) UntagResource(ctx context.Context, input *iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) UpdateChannel(ctx context.Context, input *iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error) {
	return a.client.UpdateChannelWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) UpdateDataset(ctx context.Context, input *iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error) {
	return a.client.UpdateDatasetWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) UpdateDatastore(ctx context.Context, input *iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error) {
	return a.client.UpdateDatastoreWithContext(ctx, input)
}

func (a *IoTAnalyticsActivities) UpdatePipeline(ctx context.Context, input *iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error) {
	return a.client.UpdatePipelineWithContext(ctx, input)
}
