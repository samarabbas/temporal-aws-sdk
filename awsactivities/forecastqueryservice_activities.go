
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice/forecastqueryserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ForecastQueryServiceActivities struct {
    client forecastqueryserviceiface.ForecastQueryServiceAPI
}

func NewForecastQueryServiceActivities(session *session.Session, config ...*aws.Config) *ForecastQueryServiceActivities {
    client := forecastqueryservice.New(session, config...)
    return &ForecastQueryServiceActivities{client: client}
}

func (a *ForecastQueryServiceActivities) QueryForecast(ctx context.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
    return a.client.QueryForecastWithContext(ctx, input)
}
