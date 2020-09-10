package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/batch/batchiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type BatchActivities struct {
	client batchiface.BatchAPI
}

func NewBatchActivities(session *session.Session, config ...*aws.Config) *BatchActivities {
	client := batch.New(session, config...)
	return &BatchActivities{client: client}
}

func (a *BatchActivities) CancelJob(ctx context.Context, input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
	return a.client.CancelJobWithContext(ctx, input)
}

func (a *BatchActivities) CreateComputeEnvironment(ctx context.Context, input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error) {
	return a.client.CreateComputeEnvironmentWithContext(ctx, input)
}

func (a *BatchActivities) CreateJobQueue(ctx context.Context, input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error) {
	return a.client.CreateJobQueueWithContext(ctx, input)
}

func (a *BatchActivities) DeleteComputeEnvironment(ctx context.Context, input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error) {
	return a.client.DeleteComputeEnvironmentWithContext(ctx, input)
}

func (a *BatchActivities) DeleteJobQueue(ctx context.Context, input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error) {
	return a.client.DeleteJobQueueWithContext(ctx, input)
}

func (a *BatchActivities) DeregisterJobDefinition(ctx context.Context, input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error) {
	return a.client.DeregisterJobDefinitionWithContext(ctx, input)
}

func (a *BatchActivities) DescribeComputeEnvironments(ctx context.Context, input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error) {
	return a.client.DescribeComputeEnvironmentsWithContext(ctx, input)
}

func (a *BatchActivities) DescribeJobDefinitions(ctx context.Context, input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error) {
	return a.client.DescribeJobDefinitionsWithContext(ctx, input)
}

func (a *BatchActivities) DescribeJobQueues(ctx context.Context, input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error) {
	return a.client.DescribeJobQueuesWithContext(ctx, input)
}

func (a *BatchActivities) DescribeJobs(ctx context.Context, input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error) {
	return a.client.DescribeJobsWithContext(ctx, input)
}

func (a *BatchActivities) ListJobs(ctx context.Context, input *batch.ListJobsInput) (*batch.ListJobsOutput, error) {
	return a.client.ListJobsWithContext(ctx, input)
}

func (a *BatchActivities) RegisterJobDefinition(ctx context.Context, input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error) {
	return a.client.RegisterJobDefinitionWithContext(ctx, input)
}

func (a *BatchActivities) SubmitJob(ctx context.Context, input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error) {
	return a.client.SubmitJobWithContext(ctx, input)
}

func (a *BatchActivities) TerminateJob(ctx context.Context, input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error) {
	return a.client.TerminateJobWithContext(ctx, input)
}

func (a *BatchActivities) UpdateComputeEnvironment(ctx context.Context, input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error) {
	return a.client.UpdateComputeEnvironmentWithContext(ctx, input)
}

func (a *BatchActivities) UpdateJobQueue(ctx context.Context, input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error) {
	return a.client.UpdateJobQueueWithContext(ctx, input)
}
