package awsclients

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DynamoDBClient interface {
    BatchGetItem(ctx workflow.Context, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
    BatchGetItemAsync(ctx workflow.Context, input *dynamodb.BatchGetItemInput) *DynamodbBatchGetItemResult

    BatchWriteItem(ctx workflow.Context, input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
    BatchWriteItemAsync(ctx workflow.Context, input *dynamodb.BatchWriteItemInput) *DynamodbBatchWriteItemResult

    CreateBackup(ctx workflow.Context, input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error)
    CreateBackupAsync(ctx workflow.Context, input *dynamodb.CreateBackupInput) *DynamodbCreateBackupResult

    CreateGlobalTable(ctx workflow.Context, input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error)
    CreateGlobalTableAsync(ctx workflow.Context, input *dynamodb.CreateGlobalTableInput) *DynamodbCreateGlobalTableResult

    CreateTable(ctx workflow.Context, input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
    CreateTableAsync(ctx workflow.Context, input *dynamodb.CreateTableInput) *DynamodbCreateTableResult

    DeleteBackup(ctx workflow.Context, input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error)
    DeleteBackupAsync(ctx workflow.Context, input *dynamodb.DeleteBackupInput) *DynamodbDeleteBackupResult

    DeleteItem(ctx workflow.Context, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
    DeleteItemAsync(ctx workflow.Context, input *dynamodb.DeleteItemInput) *DynamodbDeleteItemResult

    DeleteTable(ctx workflow.Context, input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)
    DeleteTableAsync(ctx workflow.Context, input *dynamodb.DeleteTableInput) *DynamodbDeleteTableResult

    DescribeBackup(ctx workflow.Context, input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error)
    DescribeBackupAsync(ctx workflow.Context, input *dynamodb.DescribeBackupInput) *DynamodbDescribeBackupResult

    DescribeContinuousBackups(ctx workflow.Context, input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error)
    DescribeContinuousBackupsAsync(ctx workflow.Context, input *dynamodb.DescribeContinuousBackupsInput) *DynamodbDescribeContinuousBackupsResult

    DescribeContributorInsights(ctx workflow.Context, input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error)
    DescribeContributorInsightsAsync(ctx workflow.Context, input *dynamodb.DescribeContributorInsightsInput) *DynamodbDescribeContributorInsightsResult

    DescribeEndpoints(ctx workflow.Context, input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error)
    DescribeEndpointsAsync(ctx workflow.Context, input *dynamodb.DescribeEndpointsInput) *DynamodbDescribeEndpointsResult

    DescribeGlobalTable(ctx workflow.Context, input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error)
    DescribeGlobalTableAsync(ctx workflow.Context, input *dynamodb.DescribeGlobalTableInput) *DynamodbDescribeGlobalTableResult

    DescribeGlobalTableSettings(ctx workflow.Context, input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
    DescribeGlobalTableSettingsAsync(ctx workflow.Context, input *dynamodb.DescribeGlobalTableSettingsInput) *DynamodbDescribeGlobalTableSettingsResult

    DescribeLimits(ctx workflow.Context, input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error)
    DescribeLimitsAsync(ctx workflow.Context, input *dynamodb.DescribeLimitsInput) *DynamodbDescribeLimitsResult

    DescribeTable(ctx workflow.Context, input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
    DescribeTableAsync(ctx workflow.Context, input *dynamodb.DescribeTableInput) *DynamodbDescribeTableResult

    DescribeTableReplicaAutoScaling(ctx workflow.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
    DescribeTableReplicaAutoScalingAsync(ctx workflow.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput) *DynamodbDescribeTableReplicaAutoScalingResult

    DescribeTimeToLive(ctx workflow.Context, input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error)
    DescribeTimeToLiveAsync(ctx workflow.Context, input *dynamodb.DescribeTimeToLiveInput) *DynamodbDescribeTimeToLiveResult

    GetItem(ctx workflow.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
    GetItemAsync(ctx workflow.Context, input *dynamodb.GetItemInput) *DynamodbGetItemResult

    ListBackups(ctx workflow.Context, input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error)
    ListBackupsAsync(ctx workflow.Context, input *dynamodb.ListBackupsInput) *DynamodbListBackupsResult

    ListContributorInsights(ctx workflow.Context, input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error)
    ListContributorInsightsAsync(ctx workflow.Context, input *dynamodb.ListContributorInsightsInput) *DynamodbListContributorInsightsResult

    ListGlobalTables(ctx workflow.Context, input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error)
    ListGlobalTablesAsync(ctx workflow.Context, input *dynamodb.ListGlobalTablesInput) *DynamodbListGlobalTablesResult

    ListTables(ctx workflow.Context, input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)
    ListTablesAsync(ctx workflow.Context, input *dynamodb.ListTablesInput) *DynamodbListTablesResult

    ListTagsOfResource(ctx workflow.Context, input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error)
    ListTagsOfResourceAsync(ctx workflow.Context, input *dynamodb.ListTagsOfResourceInput) *DynamodbListTagsOfResourceResult

    PutItem(ctx workflow.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
    PutItemAsync(ctx workflow.Context, input *dynamodb.PutItemInput) *DynamodbPutItemResult

    Query(ctx workflow.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
    QueryAsync(ctx workflow.Context, input *dynamodb.QueryInput) *DynamodbQueryResult

    RestoreTableFromBackup(ctx workflow.Context, input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error)
    RestoreTableFromBackupAsync(ctx workflow.Context, input *dynamodb.RestoreTableFromBackupInput) *DynamodbRestoreTableFromBackupResult

    RestoreTableToPointInTime(ctx workflow.Context, input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error)
    RestoreTableToPointInTimeAsync(ctx workflow.Context, input *dynamodb.RestoreTableToPointInTimeInput) *DynamodbRestoreTableToPointInTimeResult

    Scan(ctx workflow.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
    ScanAsync(ctx workflow.Context, input *dynamodb.ScanInput) *DynamodbScanResult

    TagResource(ctx workflow.Context, input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *dynamodb.TagResourceInput) *DynamodbTagResourceResult

    TransactGetItems(ctx workflow.Context, input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error)
    TransactGetItemsAsync(ctx workflow.Context, input *dynamodb.TransactGetItemsInput) *DynamodbTransactGetItemsResult

    TransactWriteItems(ctx workflow.Context, input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
    TransactWriteItemsAsync(ctx workflow.Context, input *dynamodb.TransactWriteItemsInput) *DynamodbTransactWriteItemsResult

    UntagResource(ctx workflow.Context, input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *dynamodb.UntagResourceInput) *DynamodbUntagResourceResult

    UpdateContinuousBackups(ctx workflow.Context, input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error)
    UpdateContinuousBackupsAsync(ctx workflow.Context, input *dynamodb.UpdateContinuousBackupsInput) *DynamodbUpdateContinuousBackupsResult

    UpdateContributorInsights(ctx workflow.Context, input *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error)
    UpdateContributorInsightsAsync(ctx workflow.Context, input *dynamodb.UpdateContributorInsightsInput) *DynamodbUpdateContributorInsightsResult

    UpdateGlobalTable(ctx workflow.Context, input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error)
    UpdateGlobalTableAsync(ctx workflow.Context, input *dynamodb.UpdateGlobalTableInput) *DynamodbUpdateGlobalTableResult

    UpdateGlobalTableSettings(ctx workflow.Context, input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
    UpdateGlobalTableSettingsAsync(ctx workflow.Context, input *dynamodb.UpdateGlobalTableSettingsInput) *DynamodbUpdateGlobalTableSettingsResult

    UpdateItem(ctx workflow.Context, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
    UpdateItemAsync(ctx workflow.Context, input *dynamodb.UpdateItemInput) *DynamodbUpdateItemResult

    UpdateTable(ctx workflow.Context, input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
    UpdateTableAsync(ctx workflow.Context, input *dynamodb.UpdateTableInput) *DynamodbUpdateTableResult

    UpdateTableReplicaAutoScaling(ctx workflow.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
    UpdateTableReplicaAutoScalingAsync(ctx workflow.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput) *DynamodbUpdateTableReplicaAutoScalingResult

    UpdateTimeToLive(ctx workflow.Context, input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error)
    UpdateTimeToLiveAsync(ctx workflow.Context, input *dynamodb.UpdateTimeToLiveInput) *DynamodbUpdateTimeToLiveResult

    WaitUntilTableExists(ctx workflow.Context, input *dynamodb.DescribeTableInput) error
    WaitUntilTableNotExists(ctx workflow.Context, input *dynamodb.DescribeTableInput) error}

type DynamodbBatchGetItemResult struct {
	Result workflow.Future
}

func (r *DynamodbBatchGetItemResult) Get(ctx workflow.Context) (*dynamodb.BatchGetItemOutput, error) {
    var output dynamodb.BatchGetItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbBatchWriteItemResult struct {
	Result workflow.Future
}

func (r *DynamodbBatchWriteItemResult) Get(ctx workflow.Context) (*dynamodb.BatchWriteItemOutput, error) {
    var output dynamodb.BatchWriteItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbCreateBackupResult struct {
	Result workflow.Future
}

func (r *DynamodbCreateBackupResult) Get(ctx workflow.Context) (*dynamodb.CreateBackupOutput, error) {
    var output dynamodb.CreateBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbCreateGlobalTableResult struct {
	Result workflow.Future
}

func (r *DynamodbCreateGlobalTableResult) Get(ctx workflow.Context) (*dynamodb.CreateGlobalTableOutput, error) {
    var output dynamodb.CreateGlobalTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbCreateTableResult struct {
	Result workflow.Future
}

func (r *DynamodbCreateTableResult) Get(ctx workflow.Context) (*dynamodb.CreateTableOutput, error) {
    var output dynamodb.CreateTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDeleteBackupResult struct {
	Result workflow.Future
}

func (r *DynamodbDeleteBackupResult) Get(ctx workflow.Context) (*dynamodb.DeleteBackupOutput, error) {
    var output dynamodb.DeleteBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDeleteItemResult struct {
	Result workflow.Future
}

func (r *DynamodbDeleteItemResult) Get(ctx workflow.Context) (*dynamodb.DeleteItemOutput, error) {
    var output dynamodb.DeleteItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDeleteTableResult struct {
	Result workflow.Future
}

func (r *DynamodbDeleteTableResult) Get(ctx workflow.Context) (*dynamodb.DeleteTableOutput, error) {
    var output dynamodb.DeleteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeBackupResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeBackupResult) Get(ctx workflow.Context) (*dynamodb.DescribeBackupOutput, error) {
    var output dynamodb.DescribeBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeContinuousBackupsResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeContinuousBackupsResult) Get(ctx workflow.Context) (*dynamodb.DescribeContinuousBackupsOutput, error) {
    var output dynamodb.DescribeContinuousBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeContributorInsightsResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeContributorInsightsResult) Get(ctx workflow.Context) (*dynamodb.DescribeContributorInsightsOutput, error) {
    var output dynamodb.DescribeContributorInsightsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeEndpointsResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeEndpointsResult) Get(ctx workflow.Context) (*dynamodb.DescribeEndpointsOutput, error) {
    var output dynamodb.DescribeEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeGlobalTableResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeGlobalTableResult) Get(ctx workflow.Context) (*dynamodb.DescribeGlobalTableOutput, error) {
    var output dynamodb.DescribeGlobalTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeGlobalTableSettingsResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeGlobalTableSettingsResult) Get(ctx workflow.Context) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
    var output dynamodb.DescribeGlobalTableSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeLimitsResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeLimitsResult) Get(ctx workflow.Context) (*dynamodb.DescribeLimitsOutput, error) {
    var output dynamodb.DescribeLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeTableResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeTableResult) Get(ctx workflow.Context) (*dynamodb.DescribeTableOutput, error) {
    var output dynamodb.DescribeTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeTableReplicaAutoScalingResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeTableReplicaAutoScalingResult) Get(ctx workflow.Context) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
    var output dynamodb.DescribeTableReplicaAutoScalingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbDescribeTimeToLiveResult struct {
	Result workflow.Future
}

func (r *DynamodbDescribeTimeToLiveResult) Get(ctx workflow.Context) (*dynamodb.DescribeTimeToLiveOutput, error) {
    var output dynamodb.DescribeTimeToLiveOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbGetItemResult struct {
	Result workflow.Future
}

func (r *DynamodbGetItemResult) Get(ctx workflow.Context) (*dynamodb.GetItemOutput, error) {
    var output dynamodb.GetItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbListBackupsResult struct {
	Result workflow.Future
}

func (r *DynamodbListBackupsResult) Get(ctx workflow.Context) (*dynamodb.ListBackupsOutput, error) {
    var output dynamodb.ListBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbListContributorInsightsResult struct {
	Result workflow.Future
}

func (r *DynamodbListContributorInsightsResult) Get(ctx workflow.Context) (*dynamodb.ListContributorInsightsOutput, error) {
    var output dynamodb.ListContributorInsightsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbListGlobalTablesResult struct {
	Result workflow.Future
}

func (r *DynamodbListGlobalTablesResult) Get(ctx workflow.Context) (*dynamodb.ListGlobalTablesOutput, error) {
    var output dynamodb.ListGlobalTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbListTablesResult struct {
	Result workflow.Future
}

func (r *DynamodbListTablesResult) Get(ctx workflow.Context) (*dynamodb.ListTablesOutput, error) {
    var output dynamodb.ListTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbListTagsOfResourceResult struct {
	Result workflow.Future
}

func (r *DynamodbListTagsOfResourceResult) Get(ctx workflow.Context) (*dynamodb.ListTagsOfResourceOutput, error) {
    var output dynamodb.ListTagsOfResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbPutItemResult struct {
	Result workflow.Future
}

func (r *DynamodbPutItemResult) Get(ctx workflow.Context) (*dynamodb.PutItemOutput, error) {
    var output dynamodb.PutItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbQueryResult struct {
	Result workflow.Future
}

func (r *DynamodbQueryResult) Get(ctx workflow.Context) (*dynamodb.QueryOutput, error) {
    var output dynamodb.QueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbRestoreTableFromBackupResult struct {
	Result workflow.Future
}

func (r *DynamodbRestoreTableFromBackupResult) Get(ctx workflow.Context) (*dynamodb.RestoreTableFromBackupOutput, error) {
    var output dynamodb.RestoreTableFromBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbRestoreTableToPointInTimeResult struct {
	Result workflow.Future
}

func (r *DynamodbRestoreTableToPointInTimeResult) Get(ctx workflow.Context) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
    var output dynamodb.RestoreTableToPointInTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbScanResult struct {
	Result workflow.Future
}

func (r *DynamodbScanResult) Get(ctx workflow.Context) (*dynamodb.ScanOutput, error) {
    var output dynamodb.ScanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbTagResourceResult struct {
	Result workflow.Future
}

func (r *DynamodbTagResourceResult) Get(ctx workflow.Context) (*dynamodb.TagResourceOutput, error) {
    var output dynamodb.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbTransactGetItemsResult struct {
	Result workflow.Future
}

func (r *DynamodbTransactGetItemsResult) Get(ctx workflow.Context) (*dynamodb.TransactGetItemsOutput, error) {
    var output dynamodb.TransactGetItemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbTransactWriteItemsResult struct {
	Result workflow.Future
}

func (r *DynamodbTransactWriteItemsResult) Get(ctx workflow.Context) (*dynamodb.TransactWriteItemsOutput, error) {
    var output dynamodb.TransactWriteItemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUntagResourceResult struct {
	Result workflow.Future
}

func (r *DynamodbUntagResourceResult) Get(ctx workflow.Context) (*dynamodb.UntagResourceOutput, error) {
    var output dynamodb.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateContinuousBackupsResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateContinuousBackupsResult) Get(ctx workflow.Context) (*dynamodb.UpdateContinuousBackupsOutput, error) {
    var output dynamodb.UpdateContinuousBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateContributorInsightsResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateContributorInsightsResult) Get(ctx workflow.Context) (*dynamodb.UpdateContributorInsightsOutput, error) {
    var output dynamodb.UpdateContributorInsightsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateGlobalTableResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateGlobalTableResult) Get(ctx workflow.Context) (*dynamodb.UpdateGlobalTableOutput, error) {
    var output dynamodb.UpdateGlobalTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateGlobalTableSettingsResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateGlobalTableSettingsResult) Get(ctx workflow.Context) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
    var output dynamodb.UpdateGlobalTableSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateItemResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateItemResult) Get(ctx workflow.Context) (*dynamodb.UpdateItemOutput, error) {
    var output dynamodb.UpdateItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateTableResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateTableResult) Get(ctx workflow.Context) (*dynamodb.UpdateTableOutput, error) {
    var output dynamodb.UpdateTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateTableReplicaAutoScalingResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateTableReplicaAutoScalingResult) Get(ctx workflow.Context) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
    var output dynamodb.UpdateTableReplicaAutoScalingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamodbUpdateTimeToLiveResult struct {
	Result workflow.Future
}

func (r *DynamodbUpdateTimeToLiveResult) Get(ctx workflow.Context) (*dynamodb.UpdateTimeToLiveOutput, error) {
    var output dynamodb.UpdateTimeToLiveOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DynamoDBStub struct {
    activities awsactivities.DynamoDBActivities
}

func NewDynamoDBStub() DynamoDBClient {
    return &DynamoDBStub{}
}

func (a *DynamoDBStub) BatchGetItem(ctx workflow.Context, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
    var output dynamodb.BatchGetItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetItem, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) BatchGetItemAsync(ctx workflow.Context, input *dynamodb.BatchGetItemInput) *DynamodbBatchGetItemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetItem, input)
    return &DynamodbBatchGetItemResult{Result: future}
}

func (a *DynamoDBStub) BatchWriteItem(ctx workflow.Context, input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
    var output dynamodb.BatchWriteItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchWriteItem, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) BatchWriteItemAsync(ctx workflow.Context, input *dynamodb.BatchWriteItemInput) *DynamodbBatchWriteItemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchWriteItem, input)
    return &DynamodbBatchWriteItemResult{Result: future}
}

func (a *DynamoDBStub) CreateBackup(ctx workflow.Context, input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
    var output dynamodb.CreateBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) CreateBackupAsync(ctx workflow.Context, input *dynamodb.CreateBackupInput) *DynamodbCreateBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBackup, input)
    return &DynamodbCreateBackupResult{Result: future}
}

func (a *DynamoDBStub) CreateGlobalTable(ctx workflow.Context, input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
    var output dynamodb.CreateGlobalTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGlobalTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) CreateGlobalTableAsync(ctx workflow.Context, input *dynamodb.CreateGlobalTableInput) *DynamodbCreateGlobalTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGlobalTable, input)
    return &DynamodbCreateGlobalTableResult{Result: future}
}

func (a *DynamoDBStub) CreateTable(ctx workflow.Context, input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
    var output dynamodb.CreateTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) CreateTableAsync(ctx workflow.Context, input *dynamodb.CreateTableInput) *DynamodbCreateTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTable, input)
    return &DynamodbCreateTableResult{Result: future}
}

func (a *DynamoDBStub) DeleteBackup(ctx workflow.Context, input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
    var output dynamodb.DeleteBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DeleteBackupAsync(ctx workflow.Context, input *dynamodb.DeleteBackupInput) *DynamodbDeleteBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackup, input)
    return &DynamodbDeleteBackupResult{Result: future}
}

func (a *DynamoDBStub) DeleteItem(ctx workflow.Context, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
    var output dynamodb.DeleteItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteItem, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DeleteItemAsync(ctx workflow.Context, input *dynamodb.DeleteItemInput) *DynamodbDeleteItemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteItem, input)
    return &DynamodbDeleteItemResult{Result: future}
}

func (a *DynamoDBStub) DeleteTable(ctx workflow.Context, input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
    var output dynamodb.DeleteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DeleteTableAsync(ctx workflow.Context, input *dynamodb.DeleteTableInput) *DynamodbDeleteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTable, input)
    return &DynamodbDeleteTableResult{Result: future}
}

func (a *DynamoDBStub) DescribeBackup(ctx workflow.Context, input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
    var output dynamodb.DescribeBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeBackupAsync(ctx workflow.Context, input *dynamodb.DescribeBackupInput) *DynamodbDescribeBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBackup, input)
    return &DynamodbDescribeBackupResult{Result: future}
}

func (a *DynamoDBStub) DescribeContinuousBackups(ctx workflow.Context, input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
    var output dynamodb.DescribeContinuousBackupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeContinuousBackups, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeContinuousBackupsAsync(ctx workflow.Context, input *dynamodb.DescribeContinuousBackupsInput) *DynamodbDescribeContinuousBackupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeContinuousBackups, input)
    return &DynamodbDescribeContinuousBackupsResult{Result: future}
}

func (a *DynamoDBStub) DescribeContributorInsights(ctx workflow.Context, input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
    var output dynamodb.DescribeContributorInsightsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeContributorInsights, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeContributorInsightsAsync(ctx workflow.Context, input *dynamodb.DescribeContributorInsightsInput) *DynamodbDescribeContributorInsightsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeContributorInsights, input)
    return &DynamodbDescribeContributorInsightsResult{Result: future}
}

func (a *DynamoDBStub) DescribeEndpoints(ctx workflow.Context, input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
    var output dynamodb.DescribeEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeEndpointsAsync(ctx workflow.Context, input *dynamodb.DescribeEndpointsInput) *DynamodbDescribeEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoints, input)
    return &DynamodbDescribeEndpointsResult{Result: future}
}

func (a *DynamoDBStub) DescribeGlobalTable(ctx workflow.Context, input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
    var output dynamodb.DescribeGlobalTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeGlobalTableAsync(ctx workflow.Context, input *dynamodb.DescribeGlobalTableInput) *DynamodbDescribeGlobalTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalTable, input)
    return &DynamodbDescribeGlobalTableResult{Result: future}
}

func (a *DynamoDBStub) DescribeGlobalTableSettings(ctx workflow.Context, input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
    var output dynamodb.DescribeGlobalTableSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalTableSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeGlobalTableSettingsAsync(ctx workflow.Context, input *dynamodb.DescribeGlobalTableSettingsInput) *DynamodbDescribeGlobalTableSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalTableSettings, input)
    return &DynamodbDescribeGlobalTableSettingsResult{Result: future}
}

func (a *DynamoDBStub) DescribeLimits(ctx workflow.Context, input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
    var output dynamodb.DescribeLimitsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLimits, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeLimitsAsync(ctx workflow.Context, input *dynamodb.DescribeLimitsInput) *DynamodbDescribeLimitsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLimits, input)
    return &DynamodbDescribeLimitsResult{Result: future}
}

func (a *DynamoDBStub) DescribeTable(ctx workflow.Context, input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
    var output dynamodb.DescribeTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeTableAsync(ctx workflow.Context, input *dynamodb.DescribeTableInput) *DynamodbDescribeTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTable, input)
    return &DynamodbDescribeTableResult{Result: future}
}

func (a *DynamoDBStub) DescribeTableReplicaAutoScaling(ctx workflow.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
    var output dynamodb.DescribeTableReplicaAutoScalingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTableReplicaAutoScaling, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeTableReplicaAutoScalingAsync(ctx workflow.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput) *DynamodbDescribeTableReplicaAutoScalingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTableReplicaAutoScaling, input)
    return &DynamodbDescribeTableReplicaAutoScalingResult{Result: future}
}

func (a *DynamoDBStub) DescribeTimeToLive(ctx workflow.Context, input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
    var output dynamodb.DescribeTimeToLiveOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTimeToLive, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) DescribeTimeToLiveAsync(ctx workflow.Context, input *dynamodb.DescribeTimeToLiveInput) *DynamodbDescribeTimeToLiveResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTimeToLive, input)
    return &DynamodbDescribeTimeToLiveResult{Result: future}
}

func (a *DynamoDBStub) GetItem(ctx workflow.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
    var output dynamodb.GetItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetItem, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) GetItemAsync(ctx workflow.Context, input *dynamodb.GetItemInput) *DynamodbGetItemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetItem, input)
    return &DynamodbGetItemResult{Result: future}
}

func (a *DynamoDBStub) ListBackups(ctx workflow.Context, input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
    var output dynamodb.ListBackupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackups, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) ListBackupsAsync(ctx workflow.Context, input *dynamodb.ListBackupsInput) *DynamodbListBackupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackups, input)
    return &DynamodbListBackupsResult{Result: future}
}

func (a *DynamoDBStub) ListContributorInsights(ctx workflow.Context, input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
    var output dynamodb.ListContributorInsightsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListContributorInsights, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) ListContributorInsightsAsync(ctx workflow.Context, input *dynamodb.ListContributorInsightsInput) *DynamodbListContributorInsightsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListContributorInsights, input)
    return &DynamodbListContributorInsightsResult{Result: future}
}

func (a *DynamoDBStub) ListGlobalTables(ctx workflow.Context, input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
    var output dynamodb.ListGlobalTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGlobalTables, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) ListGlobalTablesAsync(ctx workflow.Context, input *dynamodb.ListGlobalTablesInput) *DynamodbListGlobalTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGlobalTables, input)
    return &DynamodbListGlobalTablesResult{Result: future}
}

func (a *DynamoDBStub) ListTables(ctx workflow.Context, input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
    var output dynamodb.ListTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTables, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) ListTablesAsync(ctx workflow.Context, input *dynamodb.ListTablesInput) *DynamodbListTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTables, input)
    return &DynamodbListTablesResult{Result: future}
}

func (a *DynamoDBStub) ListTagsOfResource(ctx workflow.Context, input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
    var output dynamodb.ListTagsOfResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsOfResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) ListTagsOfResourceAsync(ctx workflow.Context, input *dynamodb.ListTagsOfResourceInput) *DynamodbListTagsOfResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsOfResource, input)
    return &DynamodbListTagsOfResourceResult{Result: future}
}

func (a *DynamoDBStub) PutItem(ctx workflow.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
    var output dynamodb.PutItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutItem, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) PutItemAsync(ctx workflow.Context, input *dynamodb.PutItemInput) *DynamodbPutItemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutItem, input)
    return &DynamodbPutItemResult{Result: future}
}

func (a *DynamoDBStub) Query(ctx workflow.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
    var output dynamodb.QueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Query, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) QueryAsync(ctx workflow.Context, input *dynamodb.QueryInput) *DynamodbQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Query, input)
    return &DynamodbQueryResult{Result: future}
}

func (a *DynamoDBStub) RestoreTableFromBackup(ctx workflow.Context, input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
    var output dynamodb.RestoreTableFromBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreTableFromBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) RestoreTableFromBackupAsync(ctx workflow.Context, input *dynamodb.RestoreTableFromBackupInput) *DynamodbRestoreTableFromBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreTableFromBackup, input)
    return &DynamodbRestoreTableFromBackupResult{Result: future}
}

func (a *DynamoDBStub) RestoreTableToPointInTime(ctx workflow.Context, input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
    var output dynamodb.RestoreTableToPointInTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreTableToPointInTime, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) RestoreTableToPointInTimeAsync(ctx workflow.Context, input *dynamodb.RestoreTableToPointInTimeInput) *DynamodbRestoreTableToPointInTimeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreTableToPointInTime, input)
    return &DynamodbRestoreTableToPointInTimeResult{Result: future}
}

func (a *DynamoDBStub) Scan(ctx workflow.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
    var output dynamodb.ScanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Scan, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) ScanAsync(ctx workflow.Context, input *dynamodb.ScanInput) *DynamodbScanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Scan, input)
    return &DynamodbScanResult{Result: future}
}

func (a *DynamoDBStub) TagResource(ctx workflow.Context, input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
    var output dynamodb.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) TagResourceAsync(ctx workflow.Context, input *dynamodb.TagResourceInput) *DynamodbTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &DynamodbTagResourceResult{Result: future}
}

func (a *DynamoDBStub) TransactGetItems(ctx workflow.Context, input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
    var output dynamodb.TransactGetItemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TransactGetItems, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) TransactGetItemsAsync(ctx workflow.Context, input *dynamodb.TransactGetItemsInput) *DynamodbTransactGetItemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TransactGetItems, input)
    return &DynamodbTransactGetItemsResult{Result: future}
}

func (a *DynamoDBStub) TransactWriteItems(ctx workflow.Context, input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
    var output dynamodb.TransactWriteItemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TransactWriteItems, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) TransactWriteItemsAsync(ctx workflow.Context, input *dynamodb.TransactWriteItemsInput) *DynamodbTransactWriteItemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TransactWriteItems, input)
    return &DynamodbTransactWriteItemsResult{Result: future}
}

func (a *DynamoDBStub) UntagResource(ctx workflow.Context, input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
    var output dynamodb.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UntagResourceAsync(ctx workflow.Context, input *dynamodb.UntagResourceInput) *DynamodbUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &DynamodbUntagResourceResult{Result: future}
}

func (a *DynamoDBStub) UpdateContinuousBackups(ctx workflow.Context, input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
    var output dynamodb.UpdateContinuousBackupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateContinuousBackups, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateContinuousBackupsAsync(ctx workflow.Context, input *dynamodb.UpdateContinuousBackupsInput) *DynamodbUpdateContinuousBackupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateContinuousBackups, input)
    return &DynamodbUpdateContinuousBackupsResult{Result: future}
}

func (a *DynamoDBStub) UpdateContributorInsights(ctx workflow.Context, input *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
    var output dynamodb.UpdateContributorInsightsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateContributorInsights, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateContributorInsightsAsync(ctx workflow.Context, input *dynamodb.UpdateContributorInsightsInput) *DynamodbUpdateContributorInsightsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateContributorInsights, input)
    return &DynamodbUpdateContributorInsightsResult{Result: future}
}

func (a *DynamoDBStub) UpdateGlobalTable(ctx workflow.Context, input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
    var output dynamodb.UpdateGlobalTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateGlobalTableAsync(ctx workflow.Context, input *dynamodb.UpdateGlobalTableInput) *DynamodbUpdateGlobalTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalTable, input)
    return &DynamodbUpdateGlobalTableResult{Result: future}
}

func (a *DynamoDBStub) UpdateGlobalTableSettings(ctx workflow.Context, input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
    var output dynamodb.UpdateGlobalTableSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalTableSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateGlobalTableSettingsAsync(ctx workflow.Context, input *dynamodb.UpdateGlobalTableSettingsInput) *DynamodbUpdateGlobalTableSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalTableSettings, input)
    return &DynamodbUpdateGlobalTableSettingsResult{Result: future}
}

func (a *DynamoDBStub) UpdateItem(ctx workflow.Context, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
    var output dynamodb.UpdateItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateItem, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateItemAsync(ctx workflow.Context, input *dynamodb.UpdateItemInput) *DynamodbUpdateItemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateItem, input)
    return &DynamodbUpdateItemResult{Result: future}
}

func (a *DynamoDBStub) UpdateTable(ctx workflow.Context, input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
    var output dynamodb.UpdateTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTable, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateTableAsync(ctx workflow.Context, input *dynamodb.UpdateTableInput) *DynamodbUpdateTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTable, input)
    return &DynamodbUpdateTableResult{Result: future}
}

func (a *DynamoDBStub) UpdateTableReplicaAutoScaling(ctx workflow.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
    var output dynamodb.UpdateTableReplicaAutoScalingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTableReplicaAutoScaling, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateTableReplicaAutoScalingAsync(ctx workflow.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput) *DynamodbUpdateTableReplicaAutoScalingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTableReplicaAutoScaling, input)
    return &DynamodbUpdateTableReplicaAutoScalingResult{Result: future}
}

func (a *DynamoDBStub) UpdateTimeToLive(ctx workflow.Context, input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
    var output dynamodb.UpdateTimeToLiveOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTimeToLive, input).Get(ctx, &output)
    return &output, err
}

func (a *DynamoDBStub) UpdateTimeToLiveAsync(ctx workflow.Context, input *dynamodb.UpdateTimeToLiveInput) *DynamodbUpdateTimeToLiveResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTimeToLive, input)
    return &DynamodbUpdateTimeToLiveResult{Result: future}
}

func (a *DynamoDBStub) WaitUntilTableExists(ctx workflow.Context, input *dynamodb.DescribeTableInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTableExists, input).Get(ctx, nil)
}

func (a *DynamoDBStub) WaitUntilTableExistsAsync(ctx workflow.Context, input *dynamodb.DescribeTableInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTableExists, input)
}


func (a *DynamoDBStub) WaitUntilTableNotExists(ctx workflow.Context, input *dynamodb.DescribeTableInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTableNotExists, input).Get(ctx, nil)
}

func (a *DynamoDBStub) WaitUntilTableNotExistsAsync(ctx workflow.Context, input *dynamodb.DescribeTableInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTableNotExists, input)
}
