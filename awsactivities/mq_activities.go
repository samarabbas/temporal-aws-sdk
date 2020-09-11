
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mq/mqiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MQActivities struct {
    client mqiface.MQAPI
}

func NewMQActivities(session *session.Session, config ...*aws.Config) *MQActivities {
    client := mq.New(session, config...)
    return &MQActivities{client: client}
}

func (a *MQActivities) CreateBroker(ctx context.Context, input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error) {
    return a.client.CreateBrokerWithContext(ctx, input)
}

func (a *MQActivities) CreateConfiguration(ctx context.Context, input *mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error) {
    return a.client.CreateConfigurationWithContext(ctx, input)
}

func (a *MQActivities) CreateTags(ctx context.Context, input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error) {
    return a.client.CreateTagsWithContext(ctx, input)
}

func (a *MQActivities) CreateUser(ctx context.Context, input *mq.CreateUserRequest) (*mq.CreateUserOutput, error) {
    return a.client.CreateUserWithContext(ctx, input)
}

func (a *MQActivities) DeleteBroker(ctx context.Context, input *mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error) {
    return a.client.DeleteBrokerWithContext(ctx, input)
}

func (a *MQActivities) DeleteTags(ctx context.Context, input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error) {
    return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *MQActivities) DeleteUser(ctx context.Context, input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error) {
    return a.client.DeleteUserWithContext(ctx, input)
}

func (a *MQActivities) DescribeBroker(ctx context.Context, input *mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error) {
    return a.client.DescribeBrokerWithContext(ctx, input)
}

func (a *MQActivities) DescribeBrokerEngineTypes(ctx context.Context, input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error) {
    return a.client.DescribeBrokerEngineTypesWithContext(ctx, input)
}

func (a *MQActivities) DescribeBrokerInstanceOptions(ctx context.Context, input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
    return a.client.DescribeBrokerInstanceOptionsWithContext(ctx, input)
}

func (a *MQActivities) DescribeConfiguration(ctx context.Context, input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error) {
    return a.client.DescribeConfigurationWithContext(ctx, input)
}

func (a *MQActivities) DescribeConfigurationRevision(ctx context.Context, input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error) {
    return a.client.DescribeConfigurationRevisionWithContext(ctx, input)
}

func (a *MQActivities) DescribeUser(ctx context.Context, input *mq.DescribeUserInput) (*mq.DescribeUserResponse, error) {
    return a.client.DescribeUserWithContext(ctx, input)
}

func (a *MQActivities) ListBrokers(ctx context.Context, input *mq.ListBrokersInput) (*mq.ListBrokersResponse, error) {
    return a.client.ListBrokersWithContext(ctx, input)
}

func (a *MQActivities) ListConfigurationRevisions(ctx context.Context, input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error) {
    return a.client.ListConfigurationRevisionsWithContext(ctx, input)
}

func (a *MQActivities) ListConfigurations(ctx context.Context, input *mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error) {
    return a.client.ListConfigurationsWithContext(ctx, input)
}

func (a *MQActivities) ListTags(ctx context.Context, input *mq.ListTagsInput) (*mq.ListTagsOutput, error) {
    return a.client.ListTagsWithContext(ctx, input)
}

func (a *MQActivities) ListUsers(ctx context.Context, input *mq.ListUsersInput) (*mq.ListUsersResponse, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *MQActivities) RebootBroker(ctx context.Context, input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error) {
    return a.client.RebootBrokerWithContext(ctx, input)
}

func (a *MQActivities) UpdateBroker(ctx context.Context, input *mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error) {
    return a.client.UpdateBrokerWithContext(ctx, input)
}

func (a *MQActivities) UpdateConfiguration(ctx context.Context, input *mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error) {
    return a.client.UpdateConfigurationWithContext(ctx, input)
}

func (a *MQActivities) UpdateUser(ctx context.Context, input *mq.UpdateUserRequest) (*mq.UpdateUserOutput, error) {
    return a.client.UpdateUserWithContext(ctx, input)
}
