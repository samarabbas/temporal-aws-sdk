
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"github.com/aws/aws-sdk-go/service/personalizeevents/personalizeeventsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type PersonalizeEventsActivities struct {
    client personalizeeventsiface.PersonalizeEventsAPI
}

func NewPersonalizeEventsActivities(session *session.Session, config ...*aws.Config) *PersonalizeEventsActivities {
    client := personalizeevents.New(session, config...)
    return &PersonalizeEventsActivities{client: client}
}

func (a *PersonalizeEventsActivities) PutEvents(ctx context.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
    return a.client.PutEventsWithContext(ctx, input)
}
