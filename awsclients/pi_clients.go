package awsclients

import (
	"github.com/aws/aws-sdk-go/service/pi"
	"go.temporal.io/sdk/workflow"
)

type PIClient interface {
       DescribeDimensionKeys(ctx workflow.Context, input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error)
       DescribeDimensionKeysAsync(ctx workflow.Context, input *pi.DescribeDimensionKeysInput) *PiDescribeDimensionKeysResult

       GetResourceMetrics(ctx workflow.Context, input *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error)
       GetResourceMetricsAsync(ctx workflow.Context, input *pi.GetResourceMetricsInput) *PiGetResourceMetricsResult
}

type PIStub struct {
}

func NewPIStub() PIClient {
    return &PIStub{}
}

type PiDescribeDimensionKeysResult struct {
	Result workflow.Future
}

func (r *PiDescribeDimensionKeysResult) Get(ctx workflow.Context) (*pi.DescribeDimensionKeysOutput, error) {
    var output pi.DescribeDimensionKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type PiGetResourceMetricsResult struct {
	Result workflow.Future
}

func (r *PiGetResourceMetricsResult) Get(ctx workflow.Context) (*pi.GetResourceMetricsOutput, error) {
    var output pi.GetResourceMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *PIStub) DescribeDimensionKeys(ctx workflow.Context, input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
    var output pi.DescribeDimensionKeysOutput
    err := workflow.ExecuteActivity(ctx, "PI.DescribeDimensionKeys", input).Get(ctx, &output)
    return &output, err
}

func (a *PIStub) DescribeDimensionKeysAsync(ctx workflow.Context, input *pi.DescribeDimensionKeysInput) *PiDescribeDimensionKeysResult {
    future := workflow.ExecuteActivity(ctx, "PI.DescribeDimensionKeys", input)
    return &PiDescribeDimensionKeysResult{Result: future}
}

func (a *PIStub) GetResourceMetrics(ctx workflow.Context, input *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error) {
    var output pi.GetResourceMetricsOutput
    err := workflow.ExecuteActivity(ctx, "PI.GetResourceMetrics", input).Get(ctx, &output)
    return &output, err
}

func (a *PIStub) GetResourceMetricsAsync(ctx workflow.Context, input *pi.GetResourceMetricsInput) *PiGetResourceMetricsResult {
    future := workflow.ExecuteActivity(ctx, "PI.GetResourceMetrics", input)
    return &PiGetResourceMetricsResult{Result: future}
}
