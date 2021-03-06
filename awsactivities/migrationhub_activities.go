// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/migrationhub/migrationhubiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MigrationHubActivities struct {
	client migrationhubiface.MigrationHubAPI
}

func NewMigrationHubActivities(session *session.Session, config ...*aws.Config) *MigrationHubActivities {
	client := migrationhub.New(session, config...)
	return &MigrationHubActivities{client: client}
}

func (a *MigrationHubActivities) AssociateCreatedArtifact(ctx context.Context, input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error) {
	return a.client.AssociateCreatedArtifactWithContext(ctx, input)
}

func (a *MigrationHubActivities) AssociateDiscoveredResource(ctx context.Context, input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
	return a.client.AssociateDiscoveredResourceWithContext(ctx, input)
}

func (a *MigrationHubActivities) CreateProgressUpdateStream(ctx context.Context, input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
	return a.client.CreateProgressUpdateStreamWithContext(ctx, input)
}

func (a *MigrationHubActivities) DeleteProgressUpdateStream(ctx context.Context, input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
	return a.client.DeleteProgressUpdateStreamWithContext(ctx, input)
}

func (a *MigrationHubActivities) DescribeApplicationState(ctx context.Context, input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error) {
	return a.client.DescribeApplicationStateWithContext(ctx, input)
}

func (a *MigrationHubActivities) DescribeMigrationTask(ctx context.Context, input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error) {
	return a.client.DescribeMigrationTaskWithContext(ctx, input)
}

func (a *MigrationHubActivities) DisassociateCreatedArtifact(ctx context.Context, input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
	return a.client.DisassociateCreatedArtifactWithContext(ctx, input)
}

func (a *MigrationHubActivities) DisassociateDiscoveredResource(ctx context.Context, input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
	return a.client.DisassociateDiscoveredResourceWithContext(ctx, input)
}

func (a *MigrationHubActivities) ImportMigrationTask(ctx context.Context, input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error) {
	return a.client.ImportMigrationTaskWithContext(ctx, input)
}

func (a *MigrationHubActivities) ListApplicationStates(ctx context.Context, input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error) {
	return a.client.ListApplicationStatesWithContext(ctx, input)
}

func (a *MigrationHubActivities) ListCreatedArtifacts(ctx context.Context, input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error) {
	return a.client.ListCreatedArtifactsWithContext(ctx, input)
}

func (a *MigrationHubActivities) ListDiscoveredResources(ctx context.Context, input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	return a.client.ListDiscoveredResourcesWithContext(ctx, input)
}

func (a *MigrationHubActivities) ListMigrationTasks(ctx context.Context, input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error) {
	return a.client.ListMigrationTasksWithContext(ctx, input)
}

func (a *MigrationHubActivities) ListProgressUpdateStreams(ctx context.Context, input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	return a.client.ListProgressUpdateStreamsWithContext(ctx, input)
}

func (a *MigrationHubActivities) NotifyApplicationState(ctx context.Context, input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error) {
	return a.client.NotifyApplicationStateWithContext(ctx, input)
}

func (a *MigrationHubActivities) NotifyMigrationTaskState(ctx context.Context, input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
	return a.client.NotifyMigrationTaskStateWithContext(ctx, input)
}

func (a *MigrationHubActivities) PutResourceAttributes(ctx context.Context, input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error) {
	return a.client.PutResourceAttributesWithContext(ctx, input)
}
