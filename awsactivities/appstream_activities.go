package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appstream/appstreamiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type AppStreamActivities struct {
	client appstreamiface.AppStreamAPI
}

func NewAppStreamActivities(session *session.Session, config ...*aws.Config) *AppStreamActivities {
	client := appstream.New(session, config...)
	return &AppStreamActivities{client: client}
}

func (a *AppStreamActivities) AssociateFleet(ctx context.Context, input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
	return a.client.AssociateFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) BatchAssociateUserStack(ctx context.Context, input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error) {
	return a.client.BatchAssociateUserStackWithContext(ctx, input)
}

func (a *AppStreamActivities) BatchDisassociateUserStack(ctx context.Context, input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error) {
	return a.client.BatchDisassociateUserStackWithContext(ctx, input)
}

func (a *AppStreamActivities) CopyImage(ctx context.Context, input *appstream.CopyImageInput) (*appstream.CopyImageOutput, error) {
	return a.client.CopyImageWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateDirectoryConfig(ctx context.Context, input *appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error) {
	return a.client.CreateDirectoryConfigWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateFleet(ctx context.Context, input *appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error) {
	return a.client.CreateFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateImageBuilder(ctx context.Context, input *appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error) {
	return a.client.CreateImageBuilderWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateImageBuilderStreamingURL(ctx context.Context, input *appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error) {
	return a.client.CreateImageBuilderStreamingURLWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateStack(ctx context.Context, input *appstream.CreateStackInput) (*appstream.CreateStackOutput, error) {
	return a.client.CreateStackWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateStreamingURL(ctx context.Context, input *appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error) {
	return a.client.CreateStreamingURLWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateUsageReportSubscription(ctx context.Context, input *appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error) {
	return a.client.CreateUsageReportSubscriptionWithContext(ctx, input)
}

func (a *AppStreamActivities) CreateUser(ctx context.Context, input *appstream.CreateUserInput) (*appstream.CreateUserOutput, error) {
	return a.client.CreateUserWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteDirectoryConfig(ctx context.Context, input *appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error) {
	return a.client.DeleteDirectoryConfigWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteFleet(ctx context.Context, input *appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error) {
	return a.client.DeleteFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteImage(ctx context.Context, input *appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error) {
	return a.client.DeleteImageWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteImageBuilder(ctx context.Context, input *appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error) {
	return a.client.DeleteImageBuilderWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteImagePermissions(ctx context.Context, input *appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error) {
	return a.client.DeleteImagePermissionsWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteStack(ctx context.Context, input *appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error) {
	return a.client.DeleteStackWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteUsageReportSubscription(ctx context.Context, input *appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error) {
	return a.client.DeleteUsageReportSubscriptionWithContext(ctx, input)
}

func (a *AppStreamActivities) DeleteUser(ctx context.Context, input *appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error) {
	return a.client.DeleteUserWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeDirectoryConfigs(ctx context.Context, input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error) {
	return a.client.DescribeDirectoryConfigsWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeFleets(ctx context.Context, input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error) {
	return a.client.DescribeFleetsWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeImageBuilders(ctx context.Context, input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error) {
	return a.client.DescribeImageBuildersWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeImagePermissions(ctx context.Context, input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error) {
	return a.client.DescribeImagePermissionsWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeImages(ctx context.Context, input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error) {
	return a.client.DescribeImagesWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeSessions(ctx context.Context, input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error) {
	return a.client.DescribeSessionsWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeStacks(ctx context.Context, input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error) {
	return a.client.DescribeStacksWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeUsageReportSubscriptions(ctx context.Context, input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
	return a.client.DescribeUsageReportSubscriptionsWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeUserStackAssociations(ctx context.Context, input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error) {
	return a.client.DescribeUserStackAssociationsWithContext(ctx, input)
}

func (a *AppStreamActivities) DescribeUsers(ctx context.Context, input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error) {
	return a.client.DescribeUsersWithContext(ctx, input)
}

func (a *AppStreamActivities) DisableUser(ctx context.Context, input *appstream.DisableUserInput) (*appstream.DisableUserOutput, error) {
	return a.client.DisableUserWithContext(ctx, input)
}

func (a *AppStreamActivities) DisassociateFleet(ctx context.Context, input *appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error) {
	return a.client.DisassociateFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) EnableUser(ctx context.Context, input *appstream.EnableUserInput) (*appstream.EnableUserOutput, error) {
	return a.client.EnableUserWithContext(ctx, input)
}

func (a *AppStreamActivities) ExpireSession(ctx context.Context, input *appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error) {
	return a.client.ExpireSessionWithContext(ctx, input)
}

func (a *AppStreamActivities) ListAssociatedFleets(ctx context.Context, input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error) {
	return a.client.ListAssociatedFleetsWithContext(ctx, input)
}

func (a *AppStreamActivities) ListAssociatedStacks(ctx context.Context, input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error) {
	return a.client.ListAssociatedStacksWithContext(ctx, input)
}

func (a *AppStreamActivities) ListTagsForResource(ctx context.Context, input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppStreamActivities) StartFleet(ctx context.Context, input *appstream.StartFleetInput) (*appstream.StartFleetOutput, error) {
	return a.client.StartFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) StartImageBuilder(ctx context.Context, input *appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error) {
	return a.client.StartImageBuilderWithContext(ctx, input)
}

func (a *AppStreamActivities) StopFleet(ctx context.Context, input *appstream.StopFleetInput) (*appstream.StopFleetOutput, error) {
	return a.client.StopFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) StopImageBuilder(ctx context.Context, input *appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error) {
	return a.client.StopImageBuilderWithContext(ctx, input)
}

func (a *AppStreamActivities) TagResource(ctx context.Context, input *appstream.TagResourceInput) (*appstream.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AppStreamActivities) UntagResource(ctx context.Context, input *appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AppStreamActivities) UpdateDirectoryConfig(ctx context.Context, input *appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error) {
	return a.client.UpdateDirectoryConfigWithContext(ctx, input)
}

func (a *AppStreamActivities) UpdateFleet(ctx context.Context, input *appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error) {
	return a.client.UpdateFleetWithContext(ctx, input)
}

func (a *AppStreamActivities) UpdateImagePermissions(ctx context.Context, input *appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error) {
	return a.client.UpdateImagePermissionsWithContext(ctx, input)
}

func (a *AppStreamActivities) UpdateStack(ctx context.Context, input *appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error) {
	return a.client.UpdateStackWithContext(ctx, input)
}

func (a *AppStreamActivities) WaitUntilFleetStarted(ctx context.Context, input *appstream.DescribeFleetsInput) error {
	return a.client.WaitUntilFleetStartedWithContext(ctx, input)

}

func (a *AppStreamActivities) WaitUntilFleetStopped(ctx context.Context, input *appstream.DescribeFleetsInput) error {
	return a.client.WaitUntilFleetStoppedWithContext(ctx, input)

}
