package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type IoTDataPlaneActivities struct {
	client iotdataplaneiface.IoTDataPlaneAPI
}

func NewIoTDataPlaneActivities(session *session.Session, config ...*aws.Config) *IoTDataPlaneActivities {
	client := iotdataplane.New(session, config...)
	return &IoTDataPlaneActivities{client: client}
}

func (a *IoTDataPlaneActivities) DeleteThingShadow(ctx context.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
	return a.client.DeleteThingShadowWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) GetThingShadow(ctx context.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	return a.client.GetThingShadowWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) ListNamedShadowsForThing(ctx context.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	return a.client.ListNamedShadowsForThingWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) Publish(ctx context.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
	return a.client.PublishWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) UpdateThingShadow(ctx context.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
	return a.client.UpdateThingShadowWithContext(ctx, input)
}
