package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ebs"
	"go.temporal.io/sdk/workflow"
)

type EBSClient interface {
       CompleteSnapshot(ctx workflow.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error)
       CompleteSnapshotAsync(ctx workflow.Context, input *ebs.CompleteSnapshotInput) *EbsCompleteSnapshotResult

       GetSnapshotBlock(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error)
       GetSnapshotBlockAsync(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) *EbsGetSnapshotBlockResult

       ListChangedBlocks(ctx workflow.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error)
       ListChangedBlocksAsync(ctx workflow.Context, input *ebs.ListChangedBlocksInput) *EbsListChangedBlocksResult

       ListSnapshotBlocks(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error)
       ListSnapshotBlocksAsync(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) *EbsListSnapshotBlocksResult

       PutSnapshotBlock(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error)
       PutSnapshotBlockAsync(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) *EbsPutSnapshotBlockResult

       StartSnapshot(ctx workflow.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error)
       StartSnapshotAsync(ctx workflow.Context, input *ebs.StartSnapshotInput) *EbsStartSnapshotResult
}

type EBSStub struct {
}

func NewEBSStub() EBSClient {
    return &EBSStub{}
}

type EbsCompleteSnapshotResult struct {
	Result workflow.Future
}

func (r *EbsCompleteSnapshotResult) Get(ctx workflow.Context) (*ebs.CompleteSnapshotOutput, error) {
    var output ebs.CompleteSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type EbsGetSnapshotBlockResult struct {
	Result workflow.Future
}

func (r *EbsGetSnapshotBlockResult) Get(ctx workflow.Context) (*ebs.GetSnapshotBlockOutput, error) {
    var output ebs.GetSnapshotBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type EbsListChangedBlocksResult struct {
	Result workflow.Future
}

func (r *EbsListChangedBlocksResult) Get(ctx workflow.Context) (*ebs.ListChangedBlocksOutput, error) {
    var output ebs.ListChangedBlocksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type EbsListSnapshotBlocksResult struct {
	Result workflow.Future
}

func (r *EbsListSnapshotBlocksResult) Get(ctx workflow.Context) (*ebs.ListSnapshotBlocksOutput, error) {
    var output ebs.ListSnapshotBlocksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type EbsPutSnapshotBlockResult struct {
	Result workflow.Future
}

func (r *EbsPutSnapshotBlockResult) Get(ctx workflow.Context) (*ebs.PutSnapshotBlockOutput, error) {
    var output ebs.PutSnapshotBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type EbsStartSnapshotResult struct {
	Result workflow.Future
}

func (r *EbsStartSnapshotResult) Get(ctx workflow.Context) (*ebs.StartSnapshotOutput, error) {
    var output ebs.StartSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) CompleteSnapshot(ctx workflow.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error) {
    var output ebs.CompleteSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "EBS.CompleteSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) CompleteSnapshotAsync(ctx workflow.Context, input *ebs.CompleteSnapshotInput) *EbsCompleteSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "EBS.CompleteSnapshot", input)
    return &EbsCompleteSnapshotResult{Result: future}
}

func (a *EBSStub) GetSnapshotBlock(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error) {
    var output ebs.GetSnapshotBlockOutput
    err := workflow.ExecuteActivity(ctx, "EBS.GetSnapshotBlock", input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) GetSnapshotBlockAsync(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) *EbsGetSnapshotBlockResult {
    future := workflow.ExecuteActivity(ctx, "EBS.GetSnapshotBlock", input)
    return &EbsGetSnapshotBlockResult{Result: future}
}

func (a *EBSStub) ListChangedBlocks(ctx workflow.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error) {
    var output ebs.ListChangedBlocksOutput
    err := workflow.ExecuteActivity(ctx, "EBS.ListChangedBlocks", input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) ListChangedBlocksAsync(ctx workflow.Context, input *ebs.ListChangedBlocksInput) *EbsListChangedBlocksResult {
    future := workflow.ExecuteActivity(ctx, "EBS.ListChangedBlocks", input)
    return &EbsListChangedBlocksResult{Result: future}
}

func (a *EBSStub) ListSnapshotBlocks(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error) {
    var output ebs.ListSnapshotBlocksOutput
    err := workflow.ExecuteActivity(ctx, "EBS.ListSnapshotBlocks", input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) ListSnapshotBlocksAsync(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) *EbsListSnapshotBlocksResult {
    future := workflow.ExecuteActivity(ctx, "EBS.ListSnapshotBlocks", input)
    return &EbsListSnapshotBlocksResult{Result: future}
}

func (a *EBSStub) PutSnapshotBlock(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error) {
    var output ebs.PutSnapshotBlockOutput
    err := workflow.ExecuteActivity(ctx, "EBS.PutSnapshotBlock", input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) PutSnapshotBlockAsync(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) *EbsPutSnapshotBlockResult {
    future := workflow.ExecuteActivity(ctx, "EBS.PutSnapshotBlock", input)
    return &EbsPutSnapshotBlockResult{Result: future}
}

func (a *EBSStub) StartSnapshot(ctx workflow.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error) {
    var output ebs.StartSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "EBS.StartSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) StartSnapshotAsync(ctx workflow.Context, input *ebs.StartSnapshotInput) *EbsStartSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "EBS.StartSnapshot", input)
    return &EbsStartSnapshotResult{Result: future}
}
