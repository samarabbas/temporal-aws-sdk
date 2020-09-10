package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ComprehendMedicalActivities struct {
	client comprehendmedicaliface.ComprehendMedicalAPI
}

func NewComprehendMedicalActivities(session *session.Session, config ...*aws.Config) *ComprehendMedicalActivities {
	client := comprehendmedical.New(session, config...)
	return &ComprehendMedicalActivities{client: client}
}

func (a *ComprehendMedicalActivities) DescribeEntitiesDetectionV2Job(ctx context.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	return a.client.DescribeEntitiesDetectionV2JobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) DescribeICD10CMInferenceJob(ctx context.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	return a.client.DescribeICD10CMInferenceJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) DescribePHIDetectionJob(ctx context.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	return a.client.DescribePHIDetectionJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) DescribeRxNormInferenceJob(ctx context.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	return a.client.DescribeRxNormInferenceJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) DetectEntities(ctx context.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	return a.client.DetectEntitiesWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) DetectEntitiesV2(ctx context.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error) {
	return a.client.DetectEntitiesV2WithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) DetectPHI(ctx context.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	return a.client.DetectPHIWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) InferICD10CM(ctx context.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error) {
	return a.client.InferICD10CMWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) InferRxNorm(ctx context.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error) {
	return a.client.InferRxNormWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) ListEntitiesDetectionV2Jobs(ctx context.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	return a.client.ListEntitiesDetectionV2JobsWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) ListICD10CMInferenceJobs(ctx context.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	return a.client.ListICD10CMInferenceJobsWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) ListPHIDetectionJobs(ctx context.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	return a.client.ListPHIDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) ListRxNormInferenceJobs(ctx context.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	return a.client.ListRxNormInferenceJobsWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StartEntitiesDetectionV2Job(ctx context.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	return a.client.StartEntitiesDetectionV2JobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StartICD10CMInferenceJob(ctx context.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	return a.client.StartICD10CMInferenceJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StartPHIDetectionJob(ctx context.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	return a.client.StartPHIDetectionJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StartRxNormInferenceJob(ctx context.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	return a.client.StartRxNormInferenceJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StopEntitiesDetectionV2Job(ctx context.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	return a.client.StopEntitiesDetectionV2JobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StopICD10CMInferenceJob(ctx context.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	return a.client.StopICD10CMInferenceJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StopPHIDetectionJob(ctx context.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	return a.client.StopPHIDetectionJobWithContext(ctx, input)
}

func (a *ComprehendMedicalActivities) StopRxNormInferenceJob(ctx context.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	return a.client.StopRxNormInferenceJobWithContext(ctx, input)
}
