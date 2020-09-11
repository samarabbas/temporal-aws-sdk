
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MachineLearningActivities struct {
    client machinelearningiface.MachineLearningAPI
}

func NewMachineLearningActivities(session *session.Session, config ...*aws.Config) *MachineLearningActivities {
    client := machinelearning.New(session, config...)
    return &MachineLearningActivities{client: client}
}

func (a *MachineLearningActivities) AddTags(ctx context.Context, input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error) {
    return a.client.AddTagsWithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateBatchPrediction(ctx context.Context, input *machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error) {
    return a.client.CreateBatchPredictionWithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateDataSourceFromRDS(ctx context.Context, input *machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
    return a.client.CreateDataSourceFromRDSWithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateDataSourceFromRedshift(ctx context.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
    return a.client.CreateDataSourceFromRedshiftWithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateDataSourceFromS3(ctx context.Context, input *machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error) {
    return a.client.CreateDataSourceFromS3WithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateEvaluation(ctx context.Context, input *machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error) {
    return a.client.CreateEvaluationWithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateMLModel(ctx context.Context, input *machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error) {
    return a.client.CreateMLModelWithContext(ctx, input)
}

func (a *MachineLearningActivities) CreateRealtimeEndpoint(ctx context.Context, input *machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error) {
    return a.client.CreateRealtimeEndpointWithContext(ctx, input)
}

func (a *MachineLearningActivities) DeleteBatchPrediction(ctx context.Context, input *machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error) {
    return a.client.DeleteBatchPredictionWithContext(ctx, input)
}

func (a *MachineLearningActivities) DeleteDataSource(ctx context.Context, input *machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error) {
    return a.client.DeleteDataSourceWithContext(ctx, input)
}

func (a *MachineLearningActivities) DeleteEvaluation(ctx context.Context, input *machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error) {
    return a.client.DeleteEvaluationWithContext(ctx, input)
}

func (a *MachineLearningActivities) DeleteMLModel(ctx context.Context, input *machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error) {
    return a.client.DeleteMLModelWithContext(ctx, input)
}

func (a *MachineLearningActivities) DeleteRealtimeEndpoint(ctx context.Context, input *machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
    return a.client.DeleteRealtimeEndpointWithContext(ctx, input)
}

func (a *MachineLearningActivities) DeleteTags(ctx context.Context, input *machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error) {
    return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *MachineLearningActivities) DescribeBatchPredictions(ctx context.Context, input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error) {
    return a.client.DescribeBatchPredictionsWithContext(ctx, input)
}

func (a *MachineLearningActivities) DescribeDataSources(ctx context.Context, input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error) {
    return a.client.DescribeDataSourcesWithContext(ctx, input)
}

func (a *MachineLearningActivities) DescribeEvaluations(ctx context.Context, input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error) {
    return a.client.DescribeEvaluationsWithContext(ctx, input)
}

func (a *MachineLearningActivities) DescribeMLModels(ctx context.Context, input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error) {
    return a.client.DescribeMLModelsWithContext(ctx, input)
}

func (a *MachineLearningActivities) DescribeTags(ctx context.Context, input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error) {
    return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *MachineLearningActivities) GetBatchPrediction(ctx context.Context, input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error) {
    return a.client.GetBatchPredictionWithContext(ctx, input)
}

func (a *MachineLearningActivities) GetDataSource(ctx context.Context, input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error) {
    return a.client.GetDataSourceWithContext(ctx, input)
}

func (a *MachineLearningActivities) GetEvaluation(ctx context.Context, input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error) {
    return a.client.GetEvaluationWithContext(ctx, input)
}

func (a *MachineLearningActivities) GetMLModel(ctx context.Context, input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error) {
    return a.client.GetMLModelWithContext(ctx, input)
}

func (a *MachineLearningActivities) Predict(ctx context.Context, input *machinelearning.PredictInput) (*machinelearning.PredictOutput, error) {
    return a.client.PredictWithContext(ctx, input)
}

func (a *MachineLearningActivities) UpdateBatchPrediction(ctx context.Context, input *machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error) {
    return a.client.UpdateBatchPredictionWithContext(ctx, input)
}

func (a *MachineLearningActivities) UpdateDataSource(ctx context.Context, input *machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error) {
    return a.client.UpdateDataSourceWithContext(ctx, input)
}

func (a *MachineLearningActivities) UpdateEvaluation(ctx context.Context, input *machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error) {
    return a.client.UpdateEvaluationWithContext(ctx, input)
}

func (a *MachineLearningActivities) UpdateMLModel(ctx context.Context, input *machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error) {
    return a.client.UpdateMLModelWithContext(ctx, input)
}

func (a *MachineLearningActivities) WaitUntilBatchPredictionAvailable(ctx context.Context, input *machinelearning.DescribeBatchPredictionsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilBatchPredictionAvailableWithContext(ctx, input, options...)
	})
}

func (a *MachineLearningActivities) WaitUntilDataSourceAvailable(ctx context.Context, input *machinelearning.DescribeDataSourcesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDataSourceAvailableWithContext(ctx, input, options...)
	})
}

func (a *MachineLearningActivities) WaitUntilEvaluationAvailable(ctx context.Context, input *machinelearning.DescribeEvaluationsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilEvaluationAvailableWithContext(ctx, input, options...)
	})
}

func (a *MachineLearningActivities) WaitUntilMLModelAvailable(ctx context.Context, input *machinelearning.DescribeMLModelsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilMLModelAvailableWithContext(ctx, input, options...)
	})
}
