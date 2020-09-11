
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type LexRuntimeServiceActivities struct {
    client lexruntimeserviceiface.LexRuntimeServiceAPI
}

func NewLexRuntimeServiceActivities(session *session.Session, config ...*aws.Config) *LexRuntimeServiceActivities {
    client := lexruntimeservice.New(session, config...)
    return &LexRuntimeServiceActivities{client: client}
}

func (a *LexRuntimeServiceActivities) DeleteSession(ctx context.Context, input *lexruntimeservice.DeleteSessionInput) (*lexruntimeservice.DeleteSessionOutput, error) {
    return a.client.DeleteSessionWithContext(ctx, input)
}

func (a *LexRuntimeServiceActivities) GetSession(ctx context.Context, input *lexruntimeservice.GetSessionInput) (*lexruntimeservice.GetSessionOutput, error) {
    return a.client.GetSessionWithContext(ctx, input)
}

func (a *LexRuntimeServiceActivities) PostContent(ctx context.Context, input *lexruntimeservice.PostContentInput) (*lexruntimeservice.PostContentOutput, error) {
    return a.client.PostContentWithContext(ctx, input)
}

func (a *LexRuntimeServiceActivities) PostText(ctx context.Context, input *lexruntimeservice.PostTextInput) (*lexruntimeservice.PostTextOutput, error) {
    return a.client.PostTextWithContext(ctx, input)
}

func (a *LexRuntimeServiceActivities) PutSession(ctx context.Context, input *lexruntimeservice.PutSessionInput) (*lexruntimeservice.PutSessionOutput, error) {
    return a.client.PutSessionWithContext(ctx, input)
}
