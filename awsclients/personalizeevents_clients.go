package awsclients

import (
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"go.temporal.io/sdk/workflow"
)

type PersonalizeEventsClient interface {
       PutEvents(ctx workflow.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error)
       PutEventsAsync(ctx workflow.Context, input *personalizeevents.PutEventsInput) *PersonalizeeventsPutEventsResult
}

type PersonalizeEventsStub struct {
}

func NewPersonalizeEventsStub() PersonalizeEventsClient {
    return &PersonalizeEventsStub{}
}

type PersonalizeeventsPutEventsResult struct {
	Result workflow.Future
}

func (r *PersonalizeeventsPutEventsResult) Get(ctx workflow.Context) (*personalizeevents.PutEventsOutput, error) {
    var output personalizeevents.PutEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeEventsStub) PutEvents(ctx workflow.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
    var output personalizeevents.PutEventsOutput
    err := workflow.ExecuteActivity(ctx, "PersonalizeEvents.PutEvents", input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeEventsStub) PutEventsAsync(ctx workflow.Context, input *personalizeevents.PutEventsInput) *PersonalizeeventsPutEventsResult {
    future := workflow.ExecuteActivity(ctx, "PersonalizeEvents.PutEvents", input)
    return &PersonalizeeventsPutEventsResult{Result: future}
}
