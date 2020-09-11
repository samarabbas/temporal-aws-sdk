
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CloudHSMV2Activities struct {
    client cloudhsmv2iface.CloudHSMV2API
}

func NewCloudHSMV2Activities(session *session.Session, config ...*aws.Config) *CloudHSMV2Activities {
    client := cloudhsmv2.New(session, config...)
    return &CloudHSMV2Activities{client: client}
}

func (a *CloudHSMV2Activities) CopyBackupToRegion(ctx context.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
    return a.client.CopyBackupToRegionWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) CreateCluster(ctx context.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
    return a.client.CreateClusterWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) CreateHsm(ctx context.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
    return a.client.CreateHsmWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DeleteBackup(ctx context.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
    return a.client.DeleteBackupWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DeleteCluster(ctx context.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
    return a.client.DeleteClusterWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DeleteHsm(ctx context.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
    return a.client.DeleteHsmWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DescribeBackups(ctx context.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
    return a.client.DescribeBackupsWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DescribeClusters(ctx context.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
    return a.client.DescribeClustersWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) InitializeCluster(ctx context.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
    return a.client.InitializeClusterWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) ListTags(ctx context.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
    return a.client.ListTagsWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) RestoreBackup(ctx context.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
    return a.client.RestoreBackupWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) TagResource(ctx context.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) UntagResource(ctx context.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}
