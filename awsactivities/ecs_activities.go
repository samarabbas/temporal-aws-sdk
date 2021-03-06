// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ECSActivities struct {
	client ecsiface.ECSAPI
}

func NewECSActivities(session *session.Session, config ...*aws.Config) *ECSActivities {
	client := ecs.New(session, config...)
	return &ECSActivities{client: client}
}

func (a *ECSActivities) CreateCapacityProvider(ctx context.Context, input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error) {
	return a.client.CreateCapacityProviderWithContext(ctx, input)
}

func (a *ECSActivities) CreateCluster(ctx context.Context, input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	return a.client.CreateClusterWithContext(ctx, input)
}

func (a *ECSActivities) CreateService(ctx context.Context, input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateServiceWithContext(ctx, input)
}

func (a *ECSActivities) CreateTaskSet(ctx context.Context, input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateTaskSetWithContext(ctx, input)
}

func (a *ECSActivities) DeleteAccountSetting(ctx context.Context, input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error) {
	return a.client.DeleteAccountSettingWithContext(ctx, input)
}

func (a *ECSActivities) DeleteAttributes(ctx context.Context, input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error) {
	return a.client.DeleteAttributesWithContext(ctx, input)
}

func (a *ECSActivities) DeleteCapacityProvider(ctx context.Context, input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error) {
	return a.client.DeleteCapacityProviderWithContext(ctx, input)
}

func (a *ECSActivities) DeleteCluster(ctx context.Context, input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
	return a.client.DeleteClusterWithContext(ctx, input)
}

func (a *ECSActivities) DeleteService(ctx context.Context, input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
	return a.client.DeleteServiceWithContext(ctx, input)
}

func (a *ECSActivities) DeleteTaskSet(ctx context.Context, input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error) {
	return a.client.DeleteTaskSetWithContext(ctx, input)
}

func (a *ECSActivities) DeregisterContainerInstance(ctx context.Context, input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
	return a.client.DeregisterContainerInstanceWithContext(ctx, input)
}

func (a *ECSActivities) DeregisterTaskDefinition(ctx context.Context, input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
	return a.client.DeregisterTaskDefinitionWithContext(ctx, input)
}

func (a *ECSActivities) DescribeCapacityProviders(ctx context.Context, input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error) {
	return a.client.DescribeCapacityProvidersWithContext(ctx, input)
}

func (a *ECSActivities) DescribeClusters(ctx context.Context, input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
	return a.client.DescribeClustersWithContext(ctx, input)
}

func (a *ECSActivities) DescribeContainerInstances(ctx context.Context, input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
	return a.client.DescribeContainerInstancesWithContext(ctx, input)
}

func (a *ECSActivities) DescribeServices(ctx context.Context, input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
	return a.client.DescribeServicesWithContext(ctx, input)
}

func (a *ECSActivities) DescribeTaskDefinition(ctx context.Context, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
	return a.client.DescribeTaskDefinitionWithContext(ctx, input)
}

func (a *ECSActivities) DescribeTaskSets(ctx context.Context, input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error) {
	return a.client.DescribeTaskSetsWithContext(ctx, input)
}

func (a *ECSActivities) DescribeTasks(ctx context.Context, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	return a.client.DescribeTasksWithContext(ctx, input)
}

func (a *ECSActivities) DiscoverPollEndpoint(ctx context.Context, input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
	return a.client.DiscoverPollEndpointWithContext(ctx, input)
}

func (a *ECSActivities) ListAccountSettings(ctx context.Context, input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error) {
	return a.client.ListAccountSettingsWithContext(ctx, input)
}

func (a *ECSActivities) ListAttributes(ctx context.Context, input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error) {
	return a.client.ListAttributesWithContext(ctx, input)
}

func (a *ECSActivities) ListClusters(ctx context.Context, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	return a.client.ListClustersWithContext(ctx, input)
}

func (a *ECSActivities) ListContainerInstances(ctx context.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
	return a.client.ListContainerInstancesWithContext(ctx, input)
}

func (a *ECSActivities) ListServices(ctx context.Context, input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
	return a.client.ListServicesWithContext(ctx, input)
}

func (a *ECSActivities) ListTagsForResource(ctx context.Context, input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ECSActivities) ListTaskDefinitionFamilies(ctx context.Context, input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	return a.client.ListTaskDefinitionFamiliesWithContext(ctx, input)
}

func (a *ECSActivities) ListTaskDefinitions(ctx context.Context, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return a.client.ListTaskDefinitionsWithContext(ctx, input)
}

func (a *ECSActivities) ListTasks(ctx context.Context, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
	return a.client.ListTasksWithContext(ctx, input)
}

func (a *ECSActivities) PutAccountSetting(ctx context.Context, input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error) {
	return a.client.PutAccountSettingWithContext(ctx, input)
}

func (a *ECSActivities) PutAccountSettingDefault(ctx context.Context, input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error) {
	return a.client.PutAccountSettingDefaultWithContext(ctx, input)
}

func (a *ECSActivities) PutAttributes(ctx context.Context, input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error) {
	return a.client.PutAttributesWithContext(ctx, input)
}

func (a *ECSActivities) PutClusterCapacityProviders(ctx context.Context, input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error) {
	return a.client.PutClusterCapacityProvidersWithContext(ctx, input)
}

func (a *ECSActivities) RegisterContainerInstance(ctx context.Context, input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
	return a.client.RegisterContainerInstanceWithContext(ctx, input)
}

func (a *ECSActivities) RegisterTaskDefinition(ctx context.Context, input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	return a.client.RegisterTaskDefinitionWithContext(ctx, input)
}

func (a *ECSActivities) RunTask(ctx context.Context, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
	return a.client.RunTaskWithContext(ctx, input)
}

func (a *ECSActivities) StartTask(ctx context.Context, input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
	return a.client.StartTaskWithContext(ctx, input)
}

func (a *ECSActivities) StopTask(ctx context.Context, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
	return a.client.StopTaskWithContext(ctx, input)
}

func (a *ECSActivities) SubmitAttachmentStateChanges(ctx context.Context, input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error) {
	return a.client.SubmitAttachmentStateChangesWithContext(ctx, input)
}

func (a *ECSActivities) SubmitContainerStateChange(ctx context.Context, input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
	return a.client.SubmitContainerStateChangeWithContext(ctx, input)
}

func (a *ECSActivities) SubmitTaskStateChange(ctx context.Context, input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
	return a.client.SubmitTaskStateChangeWithContext(ctx, input)
}

func (a *ECSActivities) TagResource(ctx context.Context, input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ECSActivities) UntagResource(ctx context.Context, input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ECSActivities) UpdateClusterSettings(ctx context.Context, input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error) {
	return a.client.UpdateClusterSettingsWithContext(ctx, input)
}

func (a *ECSActivities) UpdateContainerAgent(ctx context.Context, input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
	return a.client.UpdateContainerAgentWithContext(ctx, input)
}

func (a *ECSActivities) UpdateContainerInstancesState(ctx context.Context, input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error) {
	return a.client.UpdateContainerInstancesStateWithContext(ctx, input)
}

func (a *ECSActivities) UpdateService(ctx context.Context, input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	return a.client.UpdateServiceWithContext(ctx, input)
}

func (a *ECSActivities) UpdateServicePrimaryTaskSet(ctx context.Context, input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error) {
	return a.client.UpdateServicePrimaryTaskSetWithContext(ctx, input)
}

func (a *ECSActivities) UpdateTaskSet(ctx context.Context, input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error) {
	return a.client.UpdateTaskSetWithContext(ctx, input)
}

func (a *ECSActivities) WaitUntilServicesInactive(ctx context.Context, input *ecs.DescribeServicesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilServicesInactiveWithContext(ctx, input, options...)
	})
}

func (a *ECSActivities) WaitUntilServicesStable(ctx context.Context, input *ecs.DescribeServicesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilServicesStableWithContext(ctx, input, options...)
	})
}

func (a *ECSActivities) WaitUntilTasksRunning(ctx context.Context, input *ecs.DescribeTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilTasksRunningWithContext(ctx, input, options...)
	})
}

func (a *ECSActivities) WaitUntilTasksStopped(ctx context.Context, input *ecs.DescribeTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilTasksStoppedWithContext(ctx, input, options...)
	})
}
