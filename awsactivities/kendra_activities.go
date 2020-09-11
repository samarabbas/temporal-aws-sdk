
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kendra"
	"github.com/aws/aws-sdk-go/service/kendra/kendraiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type KendraActivities struct {
    client kendraiface.KendraAPI
}

func NewKendraActivities(session *session.Session, config ...*aws.Config) *KendraActivities {
    client := kendra.New(session, config...)
    return &KendraActivities{client: client}
}

func (a *KendraActivities) BatchDeleteDocument(ctx context.Context, input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error) {
    return a.client.BatchDeleteDocumentWithContext(ctx, input)
}

func (a *KendraActivities) BatchPutDocument(ctx context.Context, input *kendra.BatchPutDocumentInput) (*kendra.BatchPutDocumentOutput, error) {
    return a.client.BatchPutDocumentWithContext(ctx, input)
}

func (a *KendraActivities) CreateDataSource(ctx context.Context, input *kendra.CreateDataSourceInput) (*kendra.CreateDataSourceOutput, error) {
    return a.client.CreateDataSourceWithContext(ctx, input)
}

func (a *KendraActivities) CreateFaq(ctx context.Context, input *kendra.CreateFaqInput) (*kendra.CreateFaqOutput, error) {
    return a.client.CreateFaqWithContext(ctx, input)
}

func (a *KendraActivities) CreateIndex(ctx context.Context, input *kendra.CreateIndexInput) (*kendra.CreateIndexOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateIndexWithContext(ctx, input)
}

func (a *KendraActivities) DeleteDataSource(ctx context.Context, input *kendra.DeleteDataSourceInput) (*kendra.DeleteDataSourceOutput, error) {
    return a.client.DeleteDataSourceWithContext(ctx, input)
}

func (a *KendraActivities) DeleteFaq(ctx context.Context, input *kendra.DeleteFaqInput) (*kendra.DeleteFaqOutput, error) {
    return a.client.DeleteFaqWithContext(ctx, input)
}

func (a *KendraActivities) DeleteIndex(ctx context.Context, input *kendra.DeleteIndexInput) (*kendra.DeleteIndexOutput, error) {
    return a.client.DeleteIndexWithContext(ctx, input)
}

func (a *KendraActivities) DescribeDataSource(ctx context.Context, input *kendra.DescribeDataSourceInput) (*kendra.DescribeDataSourceOutput, error) {
    return a.client.DescribeDataSourceWithContext(ctx, input)
}

func (a *KendraActivities) DescribeFaq(ctx context.Context, input *kendra.DescribeFaqInput) (*kendra.DescribeFaqOutput, error) {
    return a.client.DescribeFaqWithContext(ctx, input)
}

func (a *KendraActivities) DescribeIndex(ctx context.Context, input *kendra.DescribeIndexInput) (*kendra.DescribeIndexOutput, error) {
    return a.client.DescribeIndexWithContext(ctx, input)
}

func (a *KendraActivities) ListDataSourceSyncJobs(ctx context.Context, input *kendra.ListDataSourceSyncJobsInput) (*kendra.ListDataSourceSyncJobsOutput, error) {
    return a.client.ListDataSourceSyncJobsWithContext(ctx, input)
}

func (a *KendraActivities) ListDataSources(ctx context.Context, input *kendra.ListDataSourcesInput) (*kendra.ListDataSourcesOutput, error) {
    return a.client.ListDataSourcesWithContext(ctx, input)
}

func (a *KendraActivities) ListFaqs(ctx context.Context, input *kendra.ListFaqsInput) (*kendra.ListFaqsOutput, error) {
    return a.client.ListFaqsWithContext(ctx, input)
}

func (a *KendraActivities) ListIndices(ctx context.Context, input *kendra.ListIndicesInput) (*kendra.ListIndicesOutput, error) {
    return a.client.ListIndicesWithContext(ctx, input)
}

func (a *KendraActivities) ListTagsForResource(ctx context.Context, input *kendra.ListTagsForResourceInput) (*kendra.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KendraActivities) Query(ctx context.Context, input *kendra.QueryInput) (*kendra.QueryOutput, error) {
    return a.client.QueryWithContext(ctx, input)
}

func (a *KendraActivities) StartDataSourceSyncJob(ctx context.Context, input *kendra.StartDataSourceSyncJobInput) (*kendra.StartDataSourceSyncJobOutput, error) {
    return a.client.StartDataSourceSyncJobWithContext(ctx, input)
}

func (a *KendraActivities) StopDataSourceSyncJob(ctx context.Context, input *kendra.StopDataSourceSyncJobInput) (*kendra.StopDataSourceSyncJobOutput, error) {
    return a.client.StopDataSourceSyncJobWithContext(ctx, input)
}

func (a *KendraActivities) SubmitFeedback(ctx context.Context, input *kendra.SubmitFeedbackInput) (*kendra.SubmitFeedbackOutput, error) {
    return a.client.SubmitFeedbackWithContext(ctx, input)
}

func (a *KendraActivities) TagResource(ctx context.Context, input *kendra.TagResourceInput) (*kendra.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *KendraActivities) UntagResource(ctx context.Context, input *kendra.UntagResourceInput) (*kendra.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *KendraActivities) UpdateDataSource(ctx context.Context, input *kendra.UpdateDataSourceInput) (*kendra.UpdateDataSourceOutput, error) {
    return a.client.UpdateDataSourceWithContext(ctx, input)
}

func (a *KendraActivities) UpdateIndex(ctx context.Context, input *kendra.UpdateIndexInput) (*kendra.UpdateIndexOutput, error) {
    return a.client.UpdateIndexWithContext(ctx, input)
}
