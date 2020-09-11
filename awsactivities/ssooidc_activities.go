
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"github.com/aws/aws-sdk-go/service/ssooidc/ssooidciface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SSOOIDCActivities struct {
    client ssooidciface.SSOOIDCAPI
}

func NewSSOOIDCActivities(session *session.Session, config ...*aws.Config) *SSOOIDCActivities {
    client := ssooidc.New(session, config...)
    return &SSOOIDCActivities{client: client}
}

func (a *SSOOIDCActivities) CreateToken(ctx context.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error) {
    return a.client.CreateTokenWithContext(ctx, input)
}

func (a *SSOOIDCActivities) RegisterClient(ctx context.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error) {
    return a.client.RegisterClientWithContext(ctx, input)
}

func (a *SSOOIDCActivities) StartDeviceAuthorization(ctx context.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error) {
    return a.client.StartDeviceAuthorizationWithContext(ctx, input)
}
