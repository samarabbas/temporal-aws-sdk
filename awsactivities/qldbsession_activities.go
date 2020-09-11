
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/aws/aws-sdk-go/service/qldbsession/qldbsessioniface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type QLDBSessionActivities struct {
    client qldbsessioniface.QLDBSessionAPI
}

func NewQLDBSessionActivities(session *session.Session, config ...*aws.Config) *QLDBSessionActivities {
    client := qldbsession.New(session, config...)
    return &QLDBSessionActivities{client: client}
}

func (a *QLDBSessionActivities) SendCommand(ctx context.Context, input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error) {
    return a.client.SendCommandWithContext(ctx, input)
}
