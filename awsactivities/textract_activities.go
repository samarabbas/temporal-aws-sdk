
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/textract/textractiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TextractActivities struct {
    client textractiface.TextractAPI
}

func NewTextractActivities(session *session.Session, config ...*aws.Config) *TextractActivities {
    client := textract.New(session, config...)
    return &TextractActivities{client: client}
}

func (a *TextractActivities) AnalyzeDocument(ctx context.Context, input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
    return a.client.AnalyzeDocumentWithContext(ctx, input)
}

func (a *TextractActivities) DetectDocumentText(ctx context.Context, input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error) {
    return a.client.DetectDocumentTextWithContext(ctx, input)
}

func (a *TextractActivities) GetDocumentAnalysis(ctx context.Context, input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error) {
    return a.client.GetDocumentAnalysisWithContext(ctx, input)
}

func (a *TextractActivities) GetDocumentTextDetection(ctx context.Context, input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error) {
    return a.client.GetDocumentTextDetectionWithContext(ctx, input)
}

func (a *TextractActivities) StartDocumentAnalysis(ctx context.Context, input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error) {
    return a.client.StartDocumentAnalysisWithContext(ctx, input)
}

func (a *TextractActivities) StartDocumentTextDetection(ctx context.Context, input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error) {
    return a.client.StartDocumentTextDetectionWithContext(ctx, input)
}
