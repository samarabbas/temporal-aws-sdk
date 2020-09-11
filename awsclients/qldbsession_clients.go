package awsclients

import (
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"go.temporal.io/sdk/workflow"
)

type QLDBSessionClient interface {
       SendCommand(ctx workflow.Context, input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error)
       SendCommandAsync(ctx workflow.Context, input *qldbsession.SendCommandInput) *QldbsessionSendCommandResult
}

type QLDBSessionStub struct {
}

func NewQLDBSessionStub() QLDBSessionClient {
    return &QLDBSessionStub{}
}

type QldbsessionSendCommandResult struct {
	Result workflow.Future
}

func (r *QldbsessionSendCommandResult) Get(ctx workflow.Context) (*qldbsession.SendCommandOutput, error) {
    var output qldbsession.SendCommandOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *QLDBSessionStub) SendCommand(ctx workflow.Context, input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error) {
    var output qldbsession.SendCommandOutput
    err := workflow.ExecuteActivity(ctx, "QLDBSession.SendCommand", input).Get(ctx, &output)
    return &output, err
}

func (a *QLDBSessionStub) SendCommandAsync(ctx workflow.Context, input *qldbsession.SendCommandInput) *QldbsessionSendCommandResult {
    future := workflow.ExecuteActivity(ctx, "QLDBSession.SendCommand", input)
    return &QldbsessionSendCommandResult{Result: future}
}
