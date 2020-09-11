
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice/iot1clickdevicesserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoT1ClickDevicesServiceActivities struct {
    client iot1clickdevicesserviceiface.IoT1ClickDevicesServiceAPI
}

func NewIoT1ClickDevicesServiceActivities(session *session.Session, config ...*aws.Config) *IoT1ClickDevicesServiceActivities {
    client := iot1clickdevicesservice.New(session, config...)
    return &IoT1ClickDevicesServiceActivities{client: client}
}

func (a *IoT1ClickDevicesServiceActivities) ClaimDevicesByClaimCode(ctx context.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
    return a.client.ClaimDevicesByClaimCodeWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) DescribeDevice(ctx context.Context, input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
    return a.client.DescribeDeviceWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) FinalizeDeviceClaim(ctx context.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
    return a.client.FinalizeDeviceClaimWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) GetDeviceMethods(ctx context.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
    return a.client.GetDeviceMethodsWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) InitiateDeviceClaim(ctx context.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
    return a.client.InitiateDeviceClaimWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) InvokeDeviceMethod(ctx context.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
    return a.client.InvokeDeviceMethodWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) ListDeviceEvents(ctx context.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
    return a.client.ListDeviceEventsWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) ListDevices(ctx context.Context, input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error) {
    return a.client.ListDevicesWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) ListTagsForResource(ctx context.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) TagResource(ctx context.Context, input *iot1clickdevicesservice.TagResourceInput) (*iot1clickdevicesservice.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) UnclaimDevice(ctx context.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
    return a.client.UnclaimDeviceWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) UntagResource(ctx context.Context, input *iot1clickdevicesservice.UntagResourceInput) (*iot1clickdevicesservice.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoT1ClickDevicesServiceActivities) UpdateDeviceState(ctx context.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
    return a.client.UpdateDeviceStateWithContext(ctx, input)
}
