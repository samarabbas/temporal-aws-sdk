package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type IoT1ClickProjectsActivities struct {
	client iot1clickprojectsiface.IoT1ClickProjectsAPI
}

func NewIoT1ClickProjectsActivities(session *session.Session, config ...*aws.Config) *IoT1ClickProjectsActivities {
	client := iot1clickprojects.New(session, config...)
	return &IoT1ClickProjectsActivities{client: client}
}

func (a *IoT1ClickProjectsActivities) AssociateDeviceWithPlacement(ctx context.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	return a.client.AssociateDeviceWithPlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) CreatePlacement(ctx context.Context, input *iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error) {
	return a.client.CreatePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) CreateProject(ctx context.Context, input *iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error) {
	return a.client.CreateProjectWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DeletePlacement(ctx context.Context, input *iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error) {
	return a.client.DeletePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DeleteProject(ctx context.Context, input *iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error) {
	return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DescribePlacement(ctx context.Context, input *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error) {
	return a.client.DescribePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DescribeProject(ctx context.Context, input *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error) {
	return a.client.DescribeProjectWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DisassociateDeviceFromPlacement(ctx context.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	return a.client.DisassociateDeviceFromPlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) GetDevicesInPlacement(ctx context.Context, input *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	return a.client.GetDevicesInPlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) ListPlacements(ctx context.Context, input *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error) {
	return a.client.ListPlacementsWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) ListProjects(ctx context.Context, input *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error) {
	return a.client.ListProjectsWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) ListTagsForResource(ctx context.Context, input *iot1clickprojects.ListTagsForResourceInput) (*iot1clickprojects.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) TagResource(ctx context.Context, input *iot1clickprojects.TagResourceInput) (*iot1clickprojects.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) UntagResource(ctx context.Context, input *iot1clickprojects.UntagResourceInput) (*iot1clickprojects.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) UpdatePlacement(ctx context.Context, input *iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error) {
	return a.client.UpdatePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) UpdateProject(ctx context.Context, input *iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error) {
	return a.client.UpdateProjectWithContext(ctx, input)
}
