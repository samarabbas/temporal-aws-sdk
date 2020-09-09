package awsclients

import (
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ForecastServiceClient interface {
    CreateDataset(ctx workflow.Context, input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error)
    CreateDatasetAsync(ctx workflow.Context, input *forecastservice.CreateDatasetInput) *ForecastserviceCreateDatasetResult

    CreateDatasetGroup(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error)
    CreateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) *ForecastserviceCreateDatasetGroupResult

    CreateDatasetImportJob(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error)
    CreateDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) *ForecastserviceCreateDatasetImportJobResult

    CreateForecast(ctx workflow.Context, input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error)
    CreateForecastAsync(ctx workflow.Context, input *forecastservice.CreateForecastInput) *ForecastserviceCreateForecastResult

    CreateForecastExportJob(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error)
    CreateForecastExportJobAsync(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) *ForecastserviceCreateForecastExportJobResult

    CreatePredictor(ctx workflow.Context, input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error)
    CreatePredictorAsync(ctx workflow.Context, input *forecastservice.CreatePredictorInput) *ForecastserviceCreatePredictorResult

    DeleteDataset(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error)
    DeleteDatasetAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) *ForecastserviceDeleteDatasetResult

    DeleteDatasetGroup(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error)
    DeleteDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) *ForecastserviceDeleteDatasetGroupResult

    DeleteDatasetImportJob(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error)
    DeleteDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) *ForecastserviceDeleteDatasetImportJobResult

    DeleteForecast(ctx workflow.Context, input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error)
    DeleteForecastAsync(ctx workflow.Context, input *forecastservice.DeleteForecastInput) *ForecastserviceDeleteForecastResult

    DeleteForecastExportJob(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error)
    DeleteForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) *ForecastserviceDeleteForecastExportJobResult

    DeletePredictor(ctx workflow.Context, input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error)
    DeletePredictorAsync(ctx workflow.Context, input *forecastservice.DeletePredictorInput) *ForecastserviceDeletePredictorResult

    DescribeDataset(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error)
    DescribeDatasetAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) *ForecastserviceDescribeDatasetResult

    DescribeDatasetGroup(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error)
    DescribeDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) *ForecastserviceDescribeDatasetGroupResult

    DescribeDatasetImportJob(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error)
    DescribeDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) *ForecastserviceDescribeDatasetImportJobResult

    DescribeForecast(ctx workflow.Context, input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error)
    DescribeForecastAsync(ctx workflow.Context, input *forecastservice.DescribeForecastInput) *ForecastserviceDescribeForecastResult

    DescribeForecastExportJob(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error)
    DescribeForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) *ForecastserviceDescribeForecastExportJobResult

    DescribePredictor(ctx workflow.Context, input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error)
    DescribePredictorAsync(ctx workflow.Context, input *forecastservice.DescribePredictorInput) *ForecastserviceDescribePredictorResult

    GetAccuracyMetrics(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error)
    GetAccuracyMetricsAsync(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) *ForecastserviceGetAccuracyMetricsResult

    ListDatasetGroups(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error)
    ListDatasetGroupsAsync(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) *ForecastserviceListDatasetGroupsResult

    ListDatasetImportJobs(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error)
    ListDatasetImportJobsAsync(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) *ForecastserviceListDatasetImportJobsResult

    ListDatasets(ctx workflow.Context, input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error)
    ListDatasetsAsync(ctx workflow.Context, input *forecastservice.ListDatasetsInput) *ForecastserviceListDatasetsResult

    ListForecastExportJobs(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error)
    ListForecastExportJobsAsync(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) *ForecastserviceListForecastExportJobsResult

    ListForecasts(ctx workflow.Context, input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error)
    ListForecastsAsync(ctx workflow.Context, input *forecastservice.ListForecastsInput) *ForecastserviceListForecastsResult

    ListPredictors(ctx workflow.Context, input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error)
    ListPredictorsAsync(ctx workflow.Context, input *forecastservice.ListPredictorsInput) *ForecastserviceListPredictorsResult

    ListTagsForResource(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) *ForecastserviceListTagsForResourceResult

    TagResource(ctx workflow.Context, input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *forecastservice.TagResourceInput) *ForecastserviceTagResourceResult

    UntagResource(ctx workflow.Context, input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *forecastservice.UntagResourceInput) *ForecastserviceUntagResourceResult

    UpdateDatasetGroup(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error)
    UpdateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) *ForecastserviceUpdateDatasetGroupResult
}

type ForecastserviceCreateDatasetResult struct {
	Result workflow.Future
}

func (r *ForecastserviceCreateDatasetResult) Get(ctx workflow.Context) (*forecastservice.CreateDatasetOutput, error) {
    var output forecastservice.CreateDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceCreateDatasetGroupResult struct {
	Result workflow.Future
}

func (r *ForecastserviceCreateDatasetGroupResult) Get(ctx workflow.Context) (*forecastservice.CreateDatasetGroupOutput, error) {
    var output forecastservice.CreateDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceCreateDatasetImportJobResult struct {
	Result workflow.Future
}

func (r *ForecastserviceCreateDatasetImportJobResult) Get(ctx workflow.Context) (*forecastservice.CreateDatasetImportJobOutput, error) {
    var output forecastservice.CreateDatasetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceCreateForecastResult struct {
	Result workflow.Future
}

func (r *ForecastserviceCreateForecastResult) Get(ctx workflow.Context) (*forecastservice.CreateForecastOutput, error) {
    var output forecastservice.CreateForecastOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceCreateForecastExportJobResult struct {
	Result workflow.Future
}

func (r *ForecastserviceCreateForecastExportJobResult) Get(ctx workflow.Context) (*forecastservice.CreateForecastExportJobOutput, error) {
    var output forecastservice.CreateForecastExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceCreatePredictorResult struct {
	Result workflow.Future
}

func (r *ForecastserviceCreatePredictorResult) Get(ctx workflow.Context) (*forecastservice.CreatePredictorOutput, error) {
    var output forecastservice.CreatePredictorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDeleteDatasetResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDeleteDatasetResult) Get(ctx workflow.Context) (*forecastservice.DeleteDatasetOutput, error) {
    var output forecastservice.DeleteDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDeleteDatasetGroupResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDeleteDatasetGroupResult) Get(ctx workflow.Context) (*forecastservice.DeleteDatasetGroupOutput, error) {
    var output forecastservice.DeleteDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDeleteDatasetImportJobResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDeleteDatasetImportJobResult) Get(ctx workflow.Context) (*forecastservice.DeleteDatasetImportJobOutput, error) {
    var output forecastservice.DeleteDatasetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDeleteForecastResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDeleteForecastResult) Get(ctx workflow.Context) (*forecastservice.DeleteForecastOutput, error) {
    var output forecastservice.DeleteForecastOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDeleteForecastExportJobResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDeleteForecastExportJobResult) Get(ctx workflow.Context) (*forecastservice.DeleteForecastExportJobOutput, error) {
    var output forecastservice.DeleteForecastExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDeletePredictorResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDeletePredictorResult) Get(ctx workflow.Context) (*forecastservice.DeletePredictorOutput, error) {
    var output forecastservice.DeletePredictorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDescribeDatasetResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDescribeDatasetResult) Get(ctx workflow.Context) (*forecastservice.DescribeDatasetOutput, error) {
    var output forecastservice.DescribeDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDescribeDatasetGroupResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDescribeDatasetGroupResult) Get(ctx workflow.Context) (*forecastservice.DescribeDatasetGroupOutput, error) {
    var output forecastservice.DescribeDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDescribeDatasetImportJobResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDescribeDatasetImportJobResult) Get(ctx workflow.Context) (*forecastservice.DescribeDatasetImportJobOutput, error) {
    var output forecastservice.DescribeDatasetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDescribeForecastResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDescribeForecastResult) Get(ctx workflow.Context) (*forecastservice.DescribeForecastOutput, error) {
    var output forecastservice.DescribeForecastOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDescribeForecastExportJobResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDescribeForecastExportJobResult) Get(ctx workflow.Context) (*forecastservice.DescribeForecastExportJobOutput, error) {
    var output forecastservice.DescribeForecastExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceDescribePredictorResult struct {
	Result workflow.Future
}

func (r *ForecastserviceDescribePredictorResult) Get(ctx workflow.Context) (*forecastservice.DescribePredictorOutput, error) {
    var output forecastservice.DescribePredictorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceGetAccuracyMetricsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceGetAccuracyMetricsResult) Get(ctx workflow.Context) (*forecastservice.GetAccuracyMetricsOutput, error) {
    var output forecastservice.GetAccuracyMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListDatasetGroupsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListDatasetGroupsResult) Get(ctx workflow.Context) (*forecastservice.ListDatasetGroupsOutput, error) {
    var output forecastservice.ListDatasetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListDatasetImportJobsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListDatasetImportJobsResult) Get(ctx workflow.Context) (*forecastservice.ListDatasetImportJobsOutput, error) {
    var output forecastservice.ListDatasetImportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListDatasetsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListDatasetsResult) Get(ctx workflow.Context) (*forecastservice.ListDatasetsOutput, error) {
    var output forecastservice.ListDatasetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListForecastExportJobsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListForecastExportJobsResult) Get(ctx workflow.Context) (*forecastservice.ListForecastExportJobsOutput, error) {
    var output forecastservice.ListForecastExportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListForecastsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListForecastsResult) Get(ctx workflow.Context) (*forecastservice.ListForecastsOutput, error) {
    var output forecastservice.ListForecastsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListPredictorsResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListPredictorsResult) Get(ctx workflow.Context) (*forecastservice.ListPredictorsOutput, error) {
    var output forecastservice.ListPredictorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ForecastserviceListTagsForResourceResult) Get(ctx workflow.Context) (*forecastservice.ListTagsForResourceOutput, error) {
    var output forecastservice.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceTagResourceResult struct {
	Result workflow.Future
}

func (r *ForecastserviceTagResourceResult) Get(ctx workflow.Context) (*forecastservice.TagResourceOutput, error) {
    var output forecastservice.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceUntagResourceResult struct {
	Result workflow.Future
}

func (r *ForecastserviceUntagResourceResult) Get(ctx workflow.Context) (*forecastservice.UntagResourceOutput, error) {
    var output forecastservice.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastserviceUpdateDatasetGroupResult struct {
	Result workflow.Future
}

func (r *ForecastserviceUpdateDatasetGroupResult) Get(ctx workflow.Context) (*forecastservice.UpdateDatasetGroupOutput, error) {
    var output forecastservice.UpdateDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastServiceStub struct {
    activities awsactivities.ForecastServiceActivities
}

func NewForecastServiceStub() ForecastServiceClient {
    return &ForecastServiceStub{}
}

func (a *ForecastServiceStub) CreateDataset(ctx workflow.Context, input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error) {
    var output forecastservice.CreateDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) CreateDatasetAsync(ctx workflow.Context, input *forecastservice.CreateDatasetInput) *ForecastserviceCreateDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataset, input)
    return &ForecastserviceCreateDatasetResult{Result: future}
}

func (a *ForecastServiceStub) CreateDatasetGroup(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error) {
    var output forecastservice.CreateDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) CreateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) *ForecastserviceCreateDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetGroup, input)
    return &ForecastserviceCreateDatasetGroupResult{Result: future}
}

func (a *ForecastServiceStub) CreateDatasetImportJob(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error) {
    var output forecastservice.CreateDatasetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) CreateDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) *ForecastserviceCreateDatasetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetImportJob, input)
    return &ForecastserviceCreateDatasetImportJobResult{Result: future}
}

func (a *ForecastServiceStub) CreateForecast(ctx workflow.Context, input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error) {
    var output forecastservice.CreateForecastOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateForecast, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) CreateForecastAsync(ctx workflow.Context, input *forecastservice.CreateForecastInput) *ForecastserviceCreateForecastResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateForecast, input)
    return &ForecastserviceCreateForecastResult{Result: future}
}

func (a *ForecastServiceStub) CreateForecastExportJob(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error) {
    var output forecastservice.CreateForecastExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateForecastExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) CreateForecastExportJobAsync(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) *ForecastserviceCreateForecastExportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateForecastExportJob, input)
    return &ForecastserviceCreateForecastExportJobResult{Result: future}
}

func (a *ForecastServiceStub) CreatePredictor(ctx workflow.Context, input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error) {
    var output forecastservice.CreatePredictorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePredictor, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) CreatePredictorAsync(ctx workflow.Context, input *forecastservice.CreatePredictorInput) *ForecastserviceCreatePredictorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePredictor, input)
    return &ForecastserviceCreatePredictorResult{Result: future}
}

func (a *ForecastServiceStub) DeleteDataset(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error) {
    var output forecastservice.DeleteDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DeleteDatasetAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) *ForecastserviceDeleteDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input)
    return &ForecastserviceDeleteDatasetResult{Result: future}
}

func (a *ForecastServiceStub) DeleteDatasetGroup(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error) {
    var output forecastservice.DeleteDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DeleteDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) *ForecastserviceDeleteDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetGroup, input)
    return &ForecastserviceDeleteDatasetGroupResult{Result: future}
}

func (a *ForecastServiceStub) DeleteDatasetImportJob(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error) {
    var output forecastservice.DeleteDatasetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DeleteDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) *ForecastserviceDeleteDatasetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetImportJob, input)
    return &ForecastserviceDeleteDatasetImportJobResult{Result: future}
}

func (a *ForecastServiceStub) DeleteForecast(ctx workflow.Context, input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error) {
    var output forecastservice.DeleteForecastOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteForecast, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DeleteForecastAsync(ctx workflow.Context, input *forecastservice.DeleteForecastInput) *ForecastserviceDeleteForecastResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteForecast, input)
    return &ForecastserviceDeleteForecastResult{Result: future}
}

func (a *ForecastServiceStub) DeleteForecastExportJob(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error) {
    var output forecastservice.DeleteForecastExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteForecastExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DeleteForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) *ForecastserviceDeleteForecastExportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteForecastExportJob, input)
    return &ForecastserviceDeleteForecastExportJobResult{Result: future}
}

func (a *ForecastServiceStub) DeletePredictor(ctx workflow.Context, input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error) {
    var output forecastservice.DeletePredictorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePredictor, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DeletePredictorAsync(ctx workflow.Context, input *forecastservice.DeletePredictorInput) *ForecastserviceDeletePredictorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePredictor, input)
    return &ForecastserviceDeletePredictorResult{Result: future}
}

func (a *ForecastServiceStub) DescribeDataset(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error) {
    var output forecastservice.DescribeDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DescribeDatasetAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) *ForecastserviceDescribeDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input)
    return &ForecastserviceDescribeDatasetResult{Result: future}
}

func (a *ForecastServiceStub) DescribeDatasetGroup(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error) {
    var output forecastservice.DescribeDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DescribeDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) *ForecastserviceDescribeDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetGroup, input)
    return &ForecastserviceDescribeDatasetGroupResult{Result: future}
}

func (a *ForecastServiceStub) DescribeDatasetImportJob(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error) {
    var output forecastservice.DescribeDatasetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DescribeDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) *ForecastserviceDescribeDatasetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetImportJob, input)
    return &ForecastserviceDescribeDatasetImportJobResult{Result: future}
}

func (a *ForecastServiceStub) DescribeForecast(ctx workflow.Context, input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error) {
    var output forecastservice.DescribeForecastOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeForecast, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DescribeForecastAsync(ctx workflow.Context, input *forecastservice.DescribeForecastInput) *ForecastserviceDescribeForecastResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeForecast, input)
    return &ForecastserviceDescribeForecastResult{Result: future}
}

func (a *ForecastServiceStub) DescribeForecastExportJob(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error) {
    var output forecastservice.DescribeForecastExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeForecastExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DescribeForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) *ForecastserviceDescribeForecastExportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeForecastExportJob, input)
    return &ForecastserviceDescribeForecastExportJobResult{Result: future}
}

func (a *ForecastServiceStub) DescribePredictor(ctx workflow.Context, input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error) {
    var output forecastservice.DescribePredictorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePredictor, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) DescribePredictorAsync(ctx workflow.Context, input *forecastservice.DescribePredictorInput) *ForecastserviceDescribePredictorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePredictor, input)
    return &ForecastserviceDescribePredictorResult{Result: future}
}

func (a *ForecastServiceStub) GetAccuracyMetrics(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error) {
    var output forecastservice.GetAccuracyMetricsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccuracyMetrics, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) GetAccuracyMetricsAsync(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) *ForecastserviceGetAccuracyMetricsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccuracyMetrics, input)
    return &ForecastserviceGetAccuracyMetricsResult{Result: future}
}

func (a *ForecastServiceStub) ListDatasetGroups(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error) {
    var output forecastservice.ListDatasetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListDatasetGroupsAsync(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) *ForecastserviceListDatasetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasetGroups, input)
    return &ForecastserviceListDatasetGroupsResult{Result: future}
}

func (a *ForecastServiceStub) ListDatasetImportJobs(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error) {
    var output forecastservice.ListDatasetImportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasetImportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListDatasetImportJobsAsync(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) *ForecastserviceListDatasetImportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasetImportJobs, input)
    return &ForecastserviceListDatasetImportJobsResult{Result: future}
}

func (a *ForecastServiceStub) ListDatasets(ctx workflow.Context, input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error) {
    var output forecastservice.ListDatasetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListDatasetsAsync(ctx workflow.Context, input *forecastservice.ListDatasetsInput) *ForecastserviceListDatasetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input)
    return &ForecastserviceListDatasetsResult{Result: future}
}

func (a *ForecastServiceStub) ListForecastExportJobs(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error) {
    var output forecastservice.ListForecastExportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListForecastExportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListForecastExportJobsAsync(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) *ForecastserviceListForecastExportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListForecastExportJobs, input)
    return &ForecastserviceListForecastExportJobsResult{Result: future}
}

func (a *ForecastServiceStub) ListForecasts(ctx workflow.Context, input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error) {
    var output forecastservice.ListForecastsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListForecasts, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListForecastsAsync(ctx workflow.Context, input *forecastservice.ListForecastsInput) *ForecastserviceListForecastsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListForecasts, input)
    return &ForecastserviceListForecastsResult{Result: future}
}

func (a *ForecastServiceStub) ListPredictors(ctx workflow.Context, input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error) {
    var output forecastservice.ListPredictorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPredictors, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListPredictorsAsync(ctx workflow.Context, input *forecastservice.ListPredictorsInput) *ForecastserviceListPredictorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPredictors, input)
    return &ForecastserviceListPredictorsResult{Result: future}
}

func (a *ForecastServiceStub) ListTagsForResource(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error) {
    var output forecastservice.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) *ForecastserviceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ForecastserviceListTagsForResourceResult{Result: future}
}

func (a *ForecastServiceStub) TagResource(ctx workflow.Context, input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error) {
    var output forecastservice.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) TagResourceAsync(ctx workflow.Context, input *forecastservice.TagResourceInput) *ForecastserviceTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &ForecastserviceTagResourceResult{Result: future}
}

func (a *ForecastServiceStub) UntagResource(ctx workflow.Context, input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error) {
    var output forecastservice.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) UntagResourceAsync(ctx workflow.Context, input *forecastservice.UntagResourceInput) *ForecastserviceUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &ForecastserviceUntagResourceResult{Result: future}
}

func (a *ForecastServiceStub) UpdateDatasetGroup(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error) {
    var output forecastservice.UpdateDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastServiceStub) UpdateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) *ForecastserviceUpdateDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDatasetGroup, input)
    return &ForecastserviceUpdateDatasetGroupResult{Result: future}
}