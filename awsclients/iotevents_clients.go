package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iotevents"
	"go.temporal.io/sdk/workflow"
)

type IoTEventsClient interface {
       CreateDetectorModel(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error)
       CreateDetectorModelAsync(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) *IoteventsCreateDetectorModelResult

       CreateInput(ctx workflow.Context, input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error)
       CreateInputAsync(ctx workflow.Context, input *iotevents.CreateInputInput) *IoteventsCreateInputResult

       DeleteDetectorModel(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error)
       DeleteDetectorModelAsync(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) *IoteventsDeleteDetectorModelResult

       DeleteInput(ctx workflow.Context, input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error)
       DeleteInputAsync(ctx workflow.Context, input *iotevents.DeleteInputInput) *IoteventsDeleteInputResult

       DescribeDetectorModel(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error)
       DescribeDetectorModelAsync(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) *IoteventsDescribeDetectorModelResult

       DescribeInput(ctx workflow.Context, input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error)
       DescribeInputAsync(ctx workflow.Context, input *iotevents.DescribeInputInput) *IoteventsDescribeInputResult

       DescribeLoggingOptions(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error)
       DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) *IoteventsDescribeLoggingOptionsResult

       ListDetectorModelVersions(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error)
       ListDetectorModelVersionsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) *IoteventsListDetectorModelVersionsResult

       ListDetectorModels(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error)
       ListDetectorModelsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) *IoteventsListDetectorModelsResult

       ListInputs(ctx workflow.Context, input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error)
       ListInputsAsync(ctx workflow.Context, input *iotevents.ListInputsInput) *IoteventsListInputsResult

       ListTagsForResource(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) *IoteventsListTagsForResourceResult

       PutLoggingOptions(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error)
       PutLoggingOptionsAsync(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) *IoteventsPutLoggingOptionsResult

       TagResource(ctx workflow.Context, input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *iotevents.TagResourceInput) *IoteventsTagResourceResult

       UntagResource(ctx workflow.Context, input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *iotevents.UntagResourceInput) *IoteventsUntagResourceResult

       UpdateDetectorModel(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error)
       UpdateDetectorModelAsync(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) *IoteventsUpdateDetectorModelResult

       UpdateInput(ctx workflow.Context, input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error)
       UpdateInputAsync(ctx workflow.Context, input *iotevents.UpdateInputInput) *IoteventsUpdateInputResult
}

type IoTEventsStub struct {
}

func NewIoTEventsStub() IoTEventsClient {
    return &IoTEventsStub{}
}

type IoteventsCreateDetectorModelResult struct {
	Result workflow.Future
}

func (r *IoteventsCreateDetectorModelResult) Get(ctx workflow.Context) (*iotevents.CreateDetectorModelOutput, error) {
    var output iotevents.CreateDetectorModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsCreateInputResult struct {
	Result workflow.Future
}

func (r *IoteventsCreateInputResult) Get(ctx workflow.Context) (*iotevents.CreateInputOutput, error) {
    var output iotevents.CreateInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsDeleteDetectorModelResult struct {
	Result workflow.Future
}

func (r *IoteventsDeleteDetectorModelResult) Get(ctx workflow.Context) (*iotevents.DeleteDetectorModelOutput, error) {
    var output iotevents.DeleteDetectorModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsDeleteInputResult struct {
	Result workflow.Future
}

func (r *IoteventsDeleteInputResult) Get(ctx workflow.Context) (*iotevents.DeleteInputOutput, error) {
    var output iotevents.DeleteInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsDescribeDetectorModelResult struct {
	Result workflow.Future
}

func (r *IoteventsDescribeDetectorModelResult) Get(ctx workflow.Context) (*iotevents.DescribeDetectorModelOutput, error) {
    var output iotevents.DescribeDetectorModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsDescribeInputResult struct {
	Result workflow.Future
}

func (r *IoteventsDescribeInputResult) Get(ctx workflow.Context) (*iotevents.DescribeInputOutput, error) {
    var output iotevents.DescribeInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsDescribeLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IoteventsDescribeLoggingOptionsResult) Get(ctx workflow.Context) (*iotevents.DescribeLoggingOptionsOutput, error) {
    var output iotevents.DescribeLoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsListDetectorModelVersionsResult struct {
	Result workflow.Future
}

func (r *IoteventsListDetectorModelVersionsResult) Get(ctx workflow.Context) (*iotevents.ListDetectorModelVersionsOutput, error) {
    var output iotevents.ListDetectorModelVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsListDetectorModelsResult struct {
	Result workflow.Future
}

func (r *IoteventsListDetectorModelsResult) Get(ctx workflow.Context) (*iotevents.ListDetectorModelsOutput, error) {
    var output iotevents.ListDetectorModelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsListInputsResult struct {
	Result workflow.Future
}

func (r *IoteventsListInputsResult) Get(ctx workflow.Context) (*iotevents.ListInputsOutput, error) {
    var output iotevents.ListInputsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IoteventsListTagsForResourceResult) Get(ctx workflow.Context) (*iotevents.ListTagsForResourceOutput, error) {
    var output iotevents.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsPutLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IoteventsPutLoggingOptionsResult) Get(ctx workflow.Context) (*iotevents.PutLoggingOptionsOutput, error) {
    var output iotevents.PutLoggingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsTagResourceResult struct {
	Result workflow.Future
}

func (r *IoteventsTagResourceResult) Get(ctx workflow.Context) (*iotevents.TagResourceOutput, error) {
    var output iotevents.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsUntagResourceResult struct {
	Result workflow.Future
}

func (r *IoteventsUntagResourceResult) Get(ctx workflow.Context) (*iotevents.UntagResourceOutput, error) {
    var output iotevents.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsUpdateDetectorModelResult struct {
	Result workflow.Future
}

func (r *IoteventsUpdateDetectorModelResult) Get(ctx workflow.Context) (*iotevents.UpdateDetectorModelOutput, error) {
    var output iotevents.UpdateDetectorModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsUpdateInputResult struct {
	Result workflow.Future
}

func (r *IoteventsUpdateInputResult) Get(ctx workflow.Context) (*iotevents.UpdateInputOutput, error) {
    var output iotevents.UpdateInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) CreateDetectorModel(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error) {
    var output iotevents.CreateDetectorModelOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.CreateDetectorModel", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) CreateDetectorModelAsync(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) *IoteventsCreateDetectorModelResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.CreateDetectorModel", input)
    return &IoteventsCreateDetectorModelResult{Result: future}
}

func (a *IoTEventsStub) CreateInput(ctx workflow.Context, input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error) {
    var output iotevents.CreateInputOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.CreateInput", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) CreateInputAsync(ctx workflow.Context, input *iotevents.CreateInputInput) *IoteventsCreateInputResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.CreateInput", input)
    return &IoteventsCreateInputResult{Result: future}
}

func (a *IoTEventsStub) DeleteDetectorModel(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error) {
    var output iotevents.DeleteDetectorModelOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.DeleteDetectorModel", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) DeleteDetectorModelAsync(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) *IoteventsDeleteDetectorModelResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.DeleteDetectorModel", input)
    return &IoteventsDeleteDetectorModelResult{Result: future}
}

func (a *IoTEventsStub) DeleteInput(ctx workflow.Context, input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error) {
    var output iotevents.DeleteInputOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.DeleteInput", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) DeleteInputAsync(ctx workflow.Context, input *iotevents.DeleteInputInput) *IoteventsDeleteInputResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.DeleteInput", input)
    return &IoteventsDeleteInputResult{Result: future}
}

func (a *IoTEventsStub) DescribeDetectorModel(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error) {
    var output iotevents.DescribeDetectorModelOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.DescribeDetectorModel", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) DescribeDetectorModelAsync(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) *IoteventsDescribeDetectorModelResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.DescribeDetectorModel", input)
    return &IoteventsDescribeDetectorModelResult{Result: future}
}

func (a *IoTEventsStub) DescribeInput(ctx workflow.Context, input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error) {
    var output iotevents.DescribeInputOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.DescribeInput", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) DescribeInputAsync(ctx workflow.Context, input *iotevents.DescribeInputInput) *IoteventsDescribeInputResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.DescribeInput", input)
    return &IoteventsDescribeInputResult{Result: future}
}

func (a *IoTEventsStub) DescribeLoggingOptions(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error) {
    var output iotevents.DescribeLoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.DescribeLoggingOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) *IoteventsDescribeLoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.DescribeLoggingOptions", input)
    return &IoteventsDescribeLoggingOptionsResult{Result: future}
}

func (a *IoTEventsStub) ListDetectorModelVersions(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error) {
    var output iotevents.ListDetectorModelVersionsOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.ListDetectorModelVersions", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) ListDetectorModelVersionsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) *IoteventsListDetectorModelVersionsResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.ListDetectorModelVersions", input)
    return &IoteventsListDetectorModelVersionsResult{Result: future}
}

func (a *IoTEventsStub) ListDetectorModels(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error) {
    var output iotevents.ListDetectorModelsOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.ListDetectorModels", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) ListDetectorModelsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) *IoteventsListDetectorModelsResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.ListDetectorModels", input)
    return &IoteventsListDetectorModelsResult{Result: future}
}

func (a *IoTEventsStub) ListInputs(ctx workflow.Context, input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error) {
    var output iotevents.ListInputsOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.ListInputs", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) ListInputsAsync(ctx workflow.Context, input *iotevents.ListInputsInput) *IoteventsListInputsResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.ListInputs", input)
    return &IoteventsListInputsResult{Result: future}
}

func (a *IoTEventsStub) ListTagsForResource(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error) {
    var output iotevents.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) ListTagsForResourceAsync(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) *IoteventsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.ListTagsForResource", input)
    return &IoteventsListTagsForResourceResult{Result: future}
}

func (a *IoTEventsStub) PutLoggingOptions(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error) {
    var output iotevents.PutLoggingOptionsOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.PutLoggingOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) PutLoggingOptionsAsync(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) *IoteventsPutLoggingOptionsResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.PutLoggingOptions", input)
    return &IoteventsPutLoggingOptionsResult{Result: future}
}

func (a *IoTEventsStub) TagResource(ctx workflow.Context, input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error) {
    var output iotevents.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) TagResourceAsync(ctx workflow.Context, input *iotevents.TagResourceInput) *IoteventsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.TagResource", input)
    return &IoteventsTagResourceResult{Result: future}
}

func (a *IoTEventsStub) UntagResource(ctx workflow.Context, input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error) {
    var output iotevents.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) UntagResourceAsync(ctx workflow.Context, input *iotevents.UntagResourceInput) *IoteventsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.UntagResource", input)
    return &IoteventsUntagResourceResult{Result: future}
}

func (a *IoTEventsStub) UpdateDetectorModel(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error) {
    var output iotevents.UpdateDetectorModelOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.UpdateDetectorModel", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) UpdateDetectorModelAsync(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) *IoteventsUpdateDetectorModelResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.UpdateDetectorModel", input)
    return &IoteventsUpdateDetectorModelResult{Result: future}
}

func (a *IoTEventsStub) UpdateInput(ctx workflow.Context, input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error) {
    var output iotevents.UpdateInputOutput
    err := workflow.ExecuteActivity(ctx, "IoTEvents.UpdateInput", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsStub) UpdateInputAsync(ctx workflow.Context, input *iotevents.UpdateInputInput) *IoteventsUpdateInputResult {
    future := workflow.ExecuteActivity(ctx, "IoTEvents.UpdateInput", input)
    return &IoteventsUpdateInputResult{Result: future}
}
