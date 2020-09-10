package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emr/emriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type EMRActivities struct {
	client emriface.EMRAPI
}

func NewEMRActivities(session *session.Session, config ...*aws.Config) *EMRActivities {
	client := emr.New(session, config...)
	return &EMRActivities{client: client}
}

func (a *EMRActivities) AddInstanceFleet(ctx context.Context, input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
	return a.client.AddInstanceFleetWithContext(ctx, input)
}

func (a *EMRActivities) AddInstanceGroups(ctx context.Context, input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error) {
	return a.client.AddInstanceGroupsWithContext(ctx, input)
}

func (a *EMRActivities) AddJobFlowSteps(ctx context.Context, input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error) {
	return a.client.AddJobFlowStepsWithContext(ctx, input)
}

func (a *EMRActivities) AddTags(ctx context.Context, input *emr.AddTagsInput) (*emr.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *EMRActivities) CancelSteps(ctx context.Context, input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error) {
	return a.client.CancelStepsWithContext(ctx, input)
}

func (a *EMRActivities) CreateSecurityConfiguration(ctx context.Context, input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error) {
	return a.client.CreateSecurityConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) DeleteSecurityConfiguration(ctx context.Context, input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error) {
	return a.client.DeleteSecurityConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) DescribeCluster(ctx context.Context, input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error) {
	return a.client.DescribeClusterWithContext(ctx, input)
}

func (a *EMRActivities) DescribeJobFlows(ctx context.Context, input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error) {
	return a.client.DescribeJobFlowsWithContext(ctx, input)
}

func (a *EMRActivities) DescribeNotebookExecution(ctx context.Context, input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error) {
	return a.client.DescribeNotebookExecutionWithContext(ctx, input)
}

func (a *EMRActivities) DescribeSecurityConfiguration(ctx context.Context, input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error) {
	return a.client.DescribeSecurityConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) DescribeStep(ctx context.Context, input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error) {
	return a.client.DescribeStepWithContext(ctx, input)
}

func (a *EMRActivities) GetBlockPublicAccessConfiguration(ctx context.Context, input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	return a.client.GetBlockPublicAccessConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) GetManagedScalingPolicy(ctx context.Context, input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error) {
	return a.client.GetManagedScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) ListBootstrapActions(ctx context.Context, input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error) {
	return a.client.ListBootstrapActionsWithContext(ctx, input)
}

func (a *EMRActivities) ListClusters(ctx context.Context, input *emr.ListClustersInput) (*emr.ListClustersOutput, error) {
	return a.client.ListClustersWithContext(ctx, input)
}

func (a *EMRActivities) ListInstanceFleets(ctx context.Context, input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error) {
	return a.client.ListInstanceFleetsWithContext(ctx, input)
}

func (a *EMRActivities) ListInstanceGroups(ctx context.Context, input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error) {
	return a.client.ListInstanceGroupsWithContext(ctx, input)
}

func (a *EMRActivities) ListInstances(ctx context.Context, input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error) {
	return a.client.ListInstancesWithContext(ctx, input)
}

func (a *EMRActivities) ListNotebookExecutions(ctx context.Context, input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error) {
	return a.client.ListNotebookExecutionsWithContext(ctx, input)
}

func (a *EMRActivities) ListSecurityConfigurations(ctx context.Context, input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error) {
	return a.client.ListSecurityConfigurationsWithContext(ctx, input)
}

func (a *EMRActivities) ListSteps(ctx context.Context, input *emr.ListStepsInput) (*emr.ListStepsOutput, error) {
	return a.client.ListStepsWithContext(ctx, input)
}

func (a *EMRActivities) ModifyCluster(ctx context.Context, input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error) {
	return a.client.ModifyClusterWithContext(ctx, input)
}

func (a *EMRActivities) ModifyInstanceFleet(ctx context.Context, input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error) {
	return a.client.ModifyInstanceFleetWithContext(ctx, input)
}

func (a *EMRActivities) ModifyInstanceGroups(ctx context.Context, input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error) {
	return a.client.ModifyInstanceGroupsWithContext(ctx, input)
}

func (a *EMRActivities) PutAutoScalingPolicy(ctx context.Context, input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error) {
	return a.client.PutAutoScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) PutBlockPublicAccessConfiguration(ctx context.Context, input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
	return a.client.PutBlockPublicAccessConfigurationWithContext(ctx, input)
}

func (a *EMRActivities) PutManagedScalingPolicy(ctx context.Context, input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error) {
	return a.client.PutManagedScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) RemoveAutoScalingPolicy(ctx context.Context, input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error) {
	return a.client.RemoveAutoScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) RemoveManagedScalingPolicy(ctx context.Context, input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error) {
	return a.client.RemoveManagedScalingPolicyWithContext(ctx, input)
}

func (a *EMRActivities) RemoveTags(ctx context.Context, input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error) {
	return a.client.RemoveTagsWithContext(ctx, input)
}

func (a *EMRActivities) RunJobFlow(ctx context.Context, input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error) {
	return a.client.RunJobFlowWithContext(ctx, input)
}

func (a *EMRActivities) SetTerminationProtection(ctx context.Context, input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error) {
	return a.client.SetTerminationProtectionWithContext(ctx, input)
}

func (a *EMRActivities) SetVisibleToAllUsers(ctx context.Context, input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error) {
	return a.client.SetVisibleToAllUsersWithContext(ctx, input)
}

func (a *EMRActivities) StartNotebookExecution(ctx context.Context, input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error) {
	return a.client.StartNotebookExecutionWithContext(ctx, input)
}

func (a *EMRActivities) StopNotebookExecution(ctx context.Context, input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error) {
	return a.client.StopNotebookExecutionWithContext(ctx, input)
}

func (a *EMRActivities) TerminateJobFlows(ctx context.Context, input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error) {
	return a.client.TerminateJobFlowsWithContext(ctx, input)
}

func (a *EMRActivities) WaitUntilClusterRunning(ctx context.Context, input *emr.DescribeClusterInput) error {
	return a.client.WaitUntilClusterRunningWithContext(ctx, input)

}

func (a *EMRActivities) WaitUntilClusterTerminated(ctx context.Context, input *emr.DescribeClusterInput) error {
	return a.client.WaitUntilClusterTerminatedWithContext(ctx, input)

}

func (a *EMRActivities) WaitUntilStepComplete(ctx context.Context, input *emr.DescribeStepInput) error {
	return a.client.WaitUntilStepCompleteWithContext(ctx, input)

}
