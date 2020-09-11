package awsclients

import (
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"go.temporal.io/sdk/workflow"
)

type LakeFormationClient interface {
       BatchGrantPermissions(ctx workflow.Context, input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error)
       BatchGrantPermissionsAsync(ctx workflow.Context, input *lakeformation.BatchGrantPermissionsInput) *LakeformationBatchGrantPermissionsResult

       BatchRevokePermissions(ctx workflow.Context, input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error)
       BatchRevokePermissionsAsync(ctx workflow.Context, input *lakeformation.BatchRevokePermissionsInput) *LakeformationBatchRevokePermissionsResult

       DeregisterResource(ctx workflow.Context, input *lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error)
       DeregisterResourceAsync(ctx workflow.Context, input *lakeformation.DeregisterResourceInput) *LakeformationDeregisterResourceResult

       DescribeResource(ctx workflow.Context, input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error)
       DescribeResourceAsync(ctx workflow.Context, input *lakeformation.DescribeResourceInput) *LakeformationDescribeResourceResult

       GetDataLakeSettings(ctx workflow.Context, input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error)
       GetDataLakeSettingsAsync(ctx workflow.Context, input *lakeformation.GetDataLakeSettingsInput) *LakeformationGetDataLakeSettingsResult

       GetEffectivePermissionsForPath(ctx workflow.Context, input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error)
       GetEffectivePermissionsForPathAsync(ctx workflow.Context, input *lakeformation.GetEffectivePermissionsForPathInput) *LakeformationGetEffectivePermissionsForPathResult

       GrantPermissions(ctx workflow.Context, input *lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error)
       GrantPermissionsAsync(ctx workflow.Context, input *lakeformation.GrantPermissionsInput) *LakeformationGrantPermissionsResult

       ListPermissions(ctx workflow.Context, input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error)
       ListPermissionsAsync(ctx workflow.Context, input *lakeformation.ListPermissionsInput) *LakeformationListPermissionsResult

       ListResources(ctx workflow.Context, input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error)
       ListResourcesAsync(ctx workflow.Context, input *lakeformation.ListResourcesInput) *LakeformationListResourcesResult

       PutDataLakeSettings(ctx workflow.Context, input *lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error)
       PutDataLakeSettingsAsync(ctx workflow.Context, input *lakeformation.PutDataLakeSettingsInput) *LakeformationPutDataLakeSettingsResult

       RegisterResource(ctx workflow.Context, input *lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error)
       RegisterResourceAsync(ctx workflow.Context, input *lakeformation.RegisterResourceInput) *LakeformationRegisterResourceResult

       RevokePermissions(ctx workflow.Context, input *lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error)
       RevokePermissionsAsync(ctx workflow.Context, input *lakeformation.RevokePermissionsInput) *LakeformationRevokePermissionsResult

       UpdateResource(ctx workflow.Context, input *lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error)
       UpdateResourceAsync(ctx workflow.Context, input *lakeformation.UpdateResourceInput) *LakeformationUpdateResourceResult
}

type LakeFormationStub struct {
}

func NewLakeFormationStub() LakeFormationClient {
    return &LakeFormationStub{}
}

type LakeformationBatchGrantPermissionsResult struct {
	Result workflow.Future
}

func (r *LakeformationBatchGrantPermissionsResult) Get(ctx workflow.Context) (*lakeformation.BatchGrantPermissionsOutput, error) {
    var output lakeformation.BatchGrantPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationBatchRevokePermissionsResult struct {
	Result workflow.Future
}

func (r *LakeformationBatchRevokePermissionsResult) Get(ctx workflow.Context) (*lakeformation.BatchRevokePermissionsOutput, error) {
    var output lakeformation.BatchRevokePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationDeregisterResourceResult struct {
	Result workflow.Future
}

func (r *LakeformationDeregisterResourceResult) Get(ctx workflow.Context) (*lakeformation.DeregisterResourceOutput, error) {
    var output lakeformation.DeregisterResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationDescribeResourceResult struct {
	Result workflow.Future
}

func (r *LakeformationDescribeResourceResult) Get(ctx workflow.Context) (*lakeformation.DescribeResourceOutput, error) {
    var output lakeformation.DescribeResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationGetDataLakeSettingsResult struct {
	Result workflow.Future
}

func (r *LakeformationGetDataLakeSettingsResult) Get(ctx workflow.Context) (*lakeformation.GetDataLakeSettingsOutput, error) {
    var output lakeformation.GetDataLakeSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationGetEffectivePermissionsForPathResult struct {
	Result workflow.Future
}

func (r *LakeformationGetEffectivePermissionsForPathResult) Get(ctx workflow.Context) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
    var output lakeformation.GetEffectivePermissionsForPathOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationGrantPermissionsResult struct {
	Result workflow.Future
}

func (r *LakeformationGrantPermissionsResult) Get(ctx workflow.Context) (*lakeformation.GrantPermissionsOutput, error) {
    var output lakeformation.GrantPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationListPermissionsResult struct {
	Result workflow.Future
}

func (r *LakeformationListPermissionsResult) Get(ctx workflow.Context) (*lakeformation.ListPermissionsOutput, error) {
    var output lakeformation.ListPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationListResourcesResult struct {
	Result workflow.Future
}

func (r *LakeformationListResourcesResult) Get(ctx workflow.Context) (*lakeformation.ListResourcesOutput, error) {
    var output lakeformation.ListResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationPutDataLakeSettingsResult struct {
	Result workflow.Future
}

func (r *LakeformationPutDataLakeSettingsResult) Get(ctx workflow.Context) (*lakeformation.PutDataLakeSettingsOutput, error) {
    var output lakeformation.PutDataLakeSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationRegisterResourceResult struct {
	Result workflow.Future
}

func (r *LakeformationRegisterResourceResult) Get(ctx workflow.Context) (*lakeformation.RegisterResourceOutput, error) {
    var output lakeformation.RegisterResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationRevokePermissionsResult struct {
	Result workflow.Future
}

func (r *LakeformationRevokePermissionsResult) Get(ctx workflow.Context) (*lakeformation.RevokePermissionsOutput, error) {
    var output lakeformation.RevokePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type LakeformationUpdateResourceResult struct {
	Result workflow.Future
}

func (r *LakeformationUpdateResourceResult) Get(ctx workflow.Context) (*lakeformation.UpdateResourceOutput, error) {
    var output lakeformation.UpdateResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) BatchGrantPermissions(ctx workflow.Context, input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
    var output lakeformation.BatchGrantPermissionsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.BatchGrantPermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) BatchGrantPermissionsAsync(ctx workflow.Context, input *lakeformation.BatchGrantPermissionsInput) *LakeformationBatchGrantPermissionsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.BatchGrantPermissions", input)
    return &LakeformationBatchGrantPermissionsResult{Result: future}
}

func (a *LakeFormationStub) BatchRevokePermissions(ctx workflow.Context, input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error) {
    var output lakeformation.BatchRevokePermissionsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.BatchRevokePermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) BatchRevokePermissionsAsync(ctx workflow.Context, input *lakeformation.BatchRevokePermissionsInput) *LakeformationBatchRevokePermissionsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.BatchRevokePermissions", input)
    return &LakeformationBatchRevokePermissionsResult{Result: future}
}

func (a *LakeFormationStub) DeregisterResource(ctx workflow.Context, input *lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error) {
    var output lakeformation.DeregisterResourceOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.DeregisterResource", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) DeregisterResourceAsync(ctx workflow.Context, input *lakeformation.DeregisterResourceInput) *LakeformationDeregisterResourceResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.DeregisterResource", input)
    return &LakeformationDeregisterResourceResult{Result: future}
}

func (a *LakeFormationStub) DescribeResource(ctx workflow.Context, input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error) {
    var output lakeformation.DescribeResourceOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.DescribeResource", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) DescribeResourceAsync(ctx workflow.Context, input *lakeformation.DescribeResourceInput) *LakeformationDescribeResourceResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.DescribeResource", input)
    return &LakeformationDescribeResourceResult{Result: future}
}

func (a *LakeFormationStub) GetDataLakeSettings(ctx workflow.Context, input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error) {
    var output lakeformation.GetDataLakeSettingsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.GetDataLakeSettings", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) GetDataLakeSettingsAsync(ctx workflow.Context, input *lakeformation.GetDataLakeSettingsInput) *LakeformationGetDataLakeSettingsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.GetDataLakeSettings", input)
    return &LakeformationGetDataLakeSettingsResult{Result: future}
}

func (a *LakeFormationStub) GetEffectivePermissionsForPath(ctx workflow.Context, input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
    var output lakeformation.GetEffectivePermissionsForPathOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.GetEffectivePermissionsForPath", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) GetEffectivePermissionsForPathAsync(ctx workflow.Context, input *lakeformation.GetEffectivePermissionsForPathInput) *LakeformationGetEffectivePermissionsForPathResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.GetEffectivePermissionsForPath", input)
    return &LakeformationGetEffectivePermissionsForPathResult{Result: future}
}

func (a *LakeFormationStub) GrantPermissions(ctx workflow.Context, input *lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error) {
    var output lakeformation.GrantPermissionsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.GrantPermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) GrantPermissionsAsync(ctx workflow.Context, input *lakeformation.GrantPermissionsInput) *LakeformationGrantPermissionsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.GrantPermissions", input)
    return &LakeformationGrantPermissionsResult{Result: future}
}

func (a *LakeFormationStub) ListPermissions(ctx workflow.Context, input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error) {
    var output lakeformation.ListPermissionsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.ListPermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) ListPermissionsAsync(ctx workflow.Context, input *lakeformation.ListPermissionsInput) *LakeformationListPermissionsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.ListPermissions", input)
    return &LakeformationListPermissionsResult{Result: future}
}

func (a *LakeFormationStub) ListResources(ctx workflow.Context, input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error) {
    var output lakeformation.ListResourcesOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.ListResources", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) ListResourcesAsync(ctx workflow.Context, input *lakeformation.ListResourcesInput) *LakeformationListResourcesResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.ListResources", input)
    return &LakeformationListResourcesResult{Result: future}
}

func (a *LakeFormationStub) PutDataLakeSettings(ctx workflow.Context, input *lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error) {
    var output lakeformation.PutDataLakeSettingsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.PutDataLakeSettings", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) PutDataLakeSettingsAsync(ctx workflow.Context, input *lakeformation.PutDataLakeSettingsInput) *LakeformationPutDataLakeSettingsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.PutDataLakeSettings", input)
    return &LakeformationPutDataLakeSettingsResult{Result: future}
}

func (a *LakeFormationStub) RegisterResource(ctx workflow.Context, input *lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error) {
    var output lakeformation.RegisterResourceOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.RegisterResource", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) RegisterResourceAsync(ctx workflow.Context, input *lakeformation.RegisterResourceInput) *LakeformationRegisterResourceResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.RegisterResource", input)
    return &LakeformationRegisterResourceResult{Result: future}
}

func (a *LakeFormationStub) RevokePermissions(ctx workflow.Context, input *lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error) {
    var output lakeformation.RevokePermissionsOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.RevokePermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) RevokePermissionsAsync(ctx workflow.Context, input *lakeformation.RevokePermissionsInput) *LakeformationRevokePermissionsResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.RevokePermissions", input)
    return &LakeformationRevokePermissionsResult{Result: future}
}

func (a *LakeFormationStub) UpdateResource(ctx workflow.Context, input *lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error) {
    var output lakeformation.UpdateResourceOutput
    err := workflow.ExecuteActivity(ctx, "LakeFormation.UpdateResource", input).Get(ctx, &output)
    return &output, err
}

func (a *LakeFormationStub) UpdateResourceAsync(ctx workflow.Context, input *lakeformation.UpdateResourceInput) *LakeformationUpdateResourceResult {
    future := workflow.ExecuteActivity(ctx, "LakeFormation.UpdateResource", input)
    return &LakeformationUpdateResourceResult{Result: future}
}
