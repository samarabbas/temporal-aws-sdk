package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"go.temporal.io/sdk/workflow"
)

type MobileAnalyticsClient interface {
       PutEvents(ctx workflow.Context, input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error)
       PutEventsAsync(ctx workflow.Context, input *mobileanalytics.PutEventsInput) *MobileanalyticsPutEventsResult
}

type MobileAnalyticsStub struct {
}

func NewMobileAnalyticsStub() MobileAnalyticsClient {
    return &MobileAnalyticsStub{}
}

type MobileanalyticsPutEventsResult struct {
	Result workflow.Future
}

func (r *MobileanalyticsPutEventsResult) Get(ctx workflow.Context) (*mobileanalytics.PutEventsOutput, error) {
    var output mobileanalytics.PutEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *MobileAnalyticsStub) PutEvents(ctx workflow.Context, input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error) {
    var output mobileanalytics.PutEventsOutput
    err := workflow.ExecuteActivity(ctx, "MobileAnalytics.PutEvents", input).Get(ctx, &output)
    return &output, err
}

func (a *MobileAnalyticsStub) PutEventsAsync(ctx workflow.Context, input *mobileanalytics.PutEventsInput) *MobileanalyticsPutEventsResult {
    future := workflow.ExecuteActivity(ctx, "MobileAnalytics.PutEvents", input)
    return &MobileanalyticsPutEventsResult{Result: future}
}
