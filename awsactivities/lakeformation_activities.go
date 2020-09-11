
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lakeformation/lakeformationiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type LakeFormationActivities struct {
    client lakeformationiface.LakeFormationAPI
}

func NewLakeFormationActivities(session *session.Session, config ...*aws.Config) *LakeFormationActivities {
    client := lakeformation.New(session, config...)
    return &LakeFormationActivities{client: client}
}

func (a *LakeFormationActivities) BatchGrantPermissions(ctx context.Context, input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
    return a.client.BatchGrantPermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) BatchRevokePermissions(ctx context.Context, input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error) {
    return a.client.BatchRevokePermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) DeregisterResource(ctx context.Context, input *lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error) {
    return a.client.DeregisterResourceWithContext(ctx, input)
}

func (a *LakeFormationActivities) DescribeResource(ctx context.Context, input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error) {
    return a.client.DescribeResourceWithContext(ctx, input)
}

func (a *LakeFormationActivities) GetDataLakeSettings(ctx context.Context, input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error) {
    return a.client.GetDataLakeSettingsWithContext(ctx, input)
}

func (a *LakeFormationActivities) GetEffectivePermissionsForPath(ctx context.Context, input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
    return a.client.GetEffectivePermissionsForPathWithContext(ctx, input)
}

func (a *LakeFormationActivities) GrantPermissions(ctx context.Context, input *lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error) {
    return a.client.GrantPermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) ListPermissions(ctx context.Context, input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error) {
    return a.client.ListPermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) ListResources(ctx context.Context, input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error) {
    return a.client.ListResourcesWithContext(ctx, input)
}

func (a *LakeFormationActivities) PutDataLakeSettings(ctx context.Context, input *lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error) {
    return a.client.PutDataLakeSettingsWithContext(ctx, input)
}

func (a *LakeFormationActivities) RegisterResource(ctx context.Context, input *lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error) {
    return a.client.RegisterResourceWithContext(ctx, input)
}

func (a *LakeFormationActivities) RevokePermissions(ctx context.Context, input *lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error) {
    return a.client.RevokePermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) UpdateResource(ctx context.Context, input *lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error) {
    return a.client.UpdateResourceWithContext(ctx, input)
}
