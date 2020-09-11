package awsclients

import (
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"go.temporal.io/sdk/workflow"
)

type TranscribeStreamingServiceClient interface {
       StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error)
       StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribestreamingserviceStartStreamTranscriptionResult
}

type TranscribeStreamingServiceStub struct {
}

func NewTranscribeStreamingServiceStub() TranscribeStreamingServiceClient {
    return &TranscribeStreamingServiceStub{}
}

type TranscribestreamingserviceStartStreamTranscriptionResult struct {
	Result workflow.Future
}

func (r *TranscribestreamingserviceStartStreamTranscriptionResult) Get(ctx workflow.Context) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
    var output transcribestreamingservice.StartStreamTranscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *TranscribeStreamingServiceStub) StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
    var output transcribestreamingservice.StartStreamTranscriptionOutput
    err := workflow.ExecuteActivity(ctx, "TranscribeStreamingService.StartStreamTranscription", input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeStreamingServiceStub) StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribestreamingserviceStartStreamTranscriptionResult {
    future := workflow.ExecuteActivity(ctx, "TranscribeStreamingService.StartStreamTranscription", input)
    return &TranscribestreamingserviceStartStreamTranscriptionResult{Result: future}
}
