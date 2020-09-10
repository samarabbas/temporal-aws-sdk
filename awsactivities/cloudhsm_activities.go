package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CloudHSMActivities struct {
	client cloudhsmiface.CloudHSMAPI
}

func NewCloudHSMActivities(session *session.Session, config ...*aws.Config) *CloudHSMActivities {
	client := cloudhsm.New(session, config...)
	return &CloudHSMActivities{client: client}
}

func (a *CloudHSMActivities) AddTagsToResource(ctx context.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *CloudHSMActivities) CreateHapg(ctx context.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
	return a.client.CreateHapgWithContext(ctx, input)
}

func (a *CloudHSMActivities) CreateHsm(ctx context.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateHsmWithContext(ctx, input)
}

func (a *CloudHSMActivities) CreateLunaClient(ctx context.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
	return a.client.CreateLunaClientWithContext(ctx, input)
}

func (a *CloudHSMActivities) DeleteHapg(ctx context.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
	return a.client.DeleteHapgWithContext(ctx, input)
}

func (a *CloudHSMActivities) DeleteHsm(ctx context.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
	return a.client.DeleteHsmWithContext(ctx, input)
}

func (a *CloudHSMActivities) DeleteLunaClient(ctx context.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
	return a.client.DeleteLunaClientWithContext(ctx, input)
}

func (a *CloudHSMActivities) DescribeHapg(ctx context.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
	return a.client.DescribeHapgWithContext(ctx, input)
}

func (a *CloudHSMActivities) DescribeHsm(ctx context.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
	return a.client.DescribeHsmWithContext(ctx, input)
}

func (a *CloudHSMActivities) DescribeLunaClient(ctx context.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
	return a.client.DescribeLunaClientWithContext(ctx, input)
}

func (a *CloudHSMActivities) GetConfig(ctx context.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
	return a.client.GetConfigWithContext(ctx, input)
}

func (a *CloudHSMActivities) ListAvailableZones(ctx context.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
	return a.client.ListAvailableZonesWithContext(ctx, input)
}

func (a *CloudHSMActivities) ListHapgs(ctx context.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
	return a.client.ListHapgsWithContext(ctx, input)
}

func (a *CloudHSMActivities) ListHsms(ctx context.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
	return a.client.ListHsmsWithContext(ctx, input)
}

func (a *CloudHSMActivities) ListLunaClients(ctx context.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
	return a.client.ListLunaClientsWithContext(ctx, input)
}

func (a *CloudHSMActivities) ListTagsForResource(ctx context.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CloudHSMActivities) ModifyHapg(ctx context.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
	return a.client.ModifyHapgWithContext(ctx, input)
}

func (a *CloudHSMActivities) ModifyHsm(ctx context.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
	return a.client.ModifyHsmWithContext(ctx, input)
}

func (a *CloudHSMActivities) ModifyLunaClient(ctx context.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
	return a.client.ModifyLunaClientWithContext(ctx, input)
}

func (a *CloudHSMActivities) RemoveTagsFromResource(ctx context.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}
