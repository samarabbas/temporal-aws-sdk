
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/redshift/redshiftiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type RedshiftActivities struct {
    client redshiftiface.RedshiftAPI
}

func NewRedshiftActivities(session *session.Session, config ...*aws.Config) *RedshiftActivities {
    client := redshift.New(session, config...)
    return &RedshiftActivities{client: client}
}

func (a *RedshiftActivities) AcceptReservedNodeExchange(ctx context.Context, input *redshift.AcceptReservedNodeExchangeInput) (*redshift.AcceptReservedNodeExchangeOutput, error) {
    return a.client.AcceptReservedNodeExchangeWithContext(ctx, input)
}

func (a *RedshiftActivities) AuthorizeClusterSecurityGroupIngress(ctx context.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error) {
    return a.client.AuthorizeClusterSecurityGroupIngressWithContext(ctx, input)
}

func (a *RedshiftActivities) AuthorizeSnapshotAccess(ctx context.Context, input *redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error) {
    return a.client.AuthorizeSnapshotAccessWithContext(ctx, input)
}

func (a *RedshiftActivities) BatchDeleteClusterSnapshots(ctx context.Context, input *redshift.BatchDeleteClusterSnapshotsInput) (*redshift.BatchDeleteClusterSnapshotsOutput, error) {
    return a.client.BatchDeleteClusterSnapshotsWithContext(ctx, input)
}

func (a *RedshiftActivities) BatchModifyClusterSnapshots(ctx context.Context, input *redshift.BatchModifyClusterSnapshotsInput) (*redshift.BatchModifyClusterSnapshotsOutput, error) {
    return a.client.BatchModifyClusterSnapshotsWithContext(ctx, input)
}

func (a *RedshiftActivities) CancelResize(ctx context.Context, input *redshift.CancelResizeInput) (*redshift.CancelResizeOutput, error) {
    return a.client.CancelResizeWithContext(ctx, input)
}

func (a *RedshiftActivities) CopyClusterSnapshot(ctx context.Context, input *redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error) {
    return a.client.CopyClusterSnapshotWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateCluster(ctx context.Context, input *redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error) {
    return a.client.CreateClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateClusterParameterGroup(ctx context.Context, input *redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error) {
    return a.client.CreateClusterParameterGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateClusterSecurityGroup(ctx context.Context, input *redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error) {
    return a.client.CreateClusterSecurityGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateClusterSnapshot(ctx context.Context, input *redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error) {
    return a.client.CreateClusterSnapshotWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateClusterSubnetGroup(ctx context.Context, input *redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error) {
    return a.client.CreateClusterSubnetGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateEventSubscription(ctx context.Context, input *redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error) {
    return a.client.CreateEventSubscriptionWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateHsmClientCertificate(ctx context.Context, input *redshift.CreateHsmClientCertificateInput) (*redshift.CreateHsmClientCertificateOutput, error) {
    return a.client.CreateHsmClientCertificateWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateHsmConfiguration(ctx context.Context, input *redshift.CreateHsmConfigurationInput) (*redshift.CreateHsmConfigurationOutput, error) {
    return a.client.CreateHsmConfigurationWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateScheduledAction(ctx context.Context, input *redshift.CreateScheduledActionInput) (*redshift.CreateScheduledActionOutput, error) {
    return a.client.CreateScheduledActionWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateSnapshotCopyGrant(ctx context.Context, input *redshift.CreateSnapshotCopyGrantInput) (*redshift.CreateSnapshotCopyGrantOutput, error) {
    return a.client.CreateSnapshotCopyGrantWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateSnapshotSchedule(ctx context.Context, input *redshift.CreateSnapshotScheduleInput) (*redshift.CreateSnapshotScheduleOutput, error) {
    return a.client.CreateSnapshotScheduleWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateTags(ctx context.Context, input *redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error) {
    return a.client.CreateTagsWithContext(ctx, input)
}

func (a *RedshiftActivities) CreateUsageLimit(ctx context.Context, input *redshift.CreateUsageLimitInput) (*redshift.CreateUsageLimitOutput, error) {
    return a.client.CreateUsageLimitWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteCluster(ctx context.Context, input *redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error) {
    return a.client.DeleteClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteClusterParameterGroup(ctx context.Context, input *redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error) {
    return a.client.DeleteClusterParameterGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteClusterSecurityGroup(ctx context.Context, input *redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error) {
    return a.client.DeleteClusterSecurityGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteClusterSnapshot(ctx context.Context, input *redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error) {
    return a.client.DeleteClusterSnapshotWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteClusterSubnetGroup(ctx context.Context, input *redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error) {
    return a.client.DeleteClusterSubnetGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteEventSubscription(ctx context.Context, input *redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error) {
    return a.client.DeleteEventSubscriptionWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteHsmClientCertificate(ctx context.Context, input *redshift.DeleteHsmClientCertificateInput) (*redshift.DeleteHsmClientCertificateOutput, error) {
    return a.client.DeleteHsmClientCertificateWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteHsmConfiguration(ctx context.Context, input *redshift.DeleteHsmConfigurationInput) (*redshift.DeleteHsmConfigurationOutput, error) {
    return a.client.DeleteHsmConfigurationWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteScheduledAction(ctx context.Context, input *redshift.DeleteScheduledActionInput) (*redshift.DeleteScheduledActionOutput, error) {
    return a.client.DeleteScheduledActionWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteSnapshotCopyGrant(ctx context.Context, input *redshift.DeleteSnapshotCopyGrantInput) (*redshift.DeleteSnapshotCopyGrantOutput, error) {
    return a.client.DeleteSnapshotCopyGrantWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteSnapshotSchedule(ctx context.Context, input *redshift.DeleteSnapshotScheduleInput) (*redshift.DeleteSnapshotScheduleOutput, error) {
    return a.client.DeleteSnapshotScheduleWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteTags(ctx context.Context, input *redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error) {
    return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *RedshiftActivities) DeleteUsageLimit(ctx context.Context, input *redshift.DeleteUsageLimitInput) (*redshift.DeleteUsageLimitOutput, error) {
    return a.client.DeleteUsageLimitWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeAccountAttributes(ctx context.Context, input *redshift.DescribeAccountAttributesInput) (*redshift.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterDbRevisions(ctx context.Context, input *redshift.DescribeClusterDbRevisionsInput) (*redshift.DescribeClusterDbRevisionsOutput, error) {
    return a.client.DescribeClusterDbRevisionsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterParameterGroups(ctx context.Context, input *redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error) {
    return a.client.DescribeClusterParameterGroupsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterParameters(ctx context.Context, input *redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error) {
    return a.client.DescribeClusterParametersWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterSecurityGroups(ctx context.Context, input *redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
    return a.client.DescribeClusterSecurityGroupsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterSnapshots(ctx context.Context, input *redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error) {
    return a.client.DescribeClusterSnapshotsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterSubnetGroups(ctx context.Context, input *redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
    return a.client.DescribeClusterSubnetGroupsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterTracks(ctx context.Context, input *redshift.DescribeClusterTracksInput) (*redshift.DescribeClusterTracksOutput, error) {
    return a.client.DescribeClusterTracksWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusterVersions(ctx context.Context, input *redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error) {
    return a.client.DescribeClusterVersionsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeClusters(ctx context.Context, input *redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error) {
    return a.client.DescribeClustersWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeDefaultClusterParameters(ctx context.Context, input *redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error) {
    return a.client.DescribeDefaultClusterParametersWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeEventCategories(ctx context.Context, input *redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error) {
    return a.client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeEventSubscriptions(ctx context.Context, input *redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error) {
    return a.client.DescribeEventSubscriptionsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeEvents(ctx context.Context, input *redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error) {
    return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeHsmClientCertificates(ctx context.Context, input *redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error) {
    return a.client.DescribeHsmClientCertificatesWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeHsmConfigurations(ctx context.Context, input *redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error) {
    return a.client.DescribeHsmConfigurationsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeLoggingStatus(ctx context.Context, input *redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error) {
    return a.client.DescribeLoggingStatusWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeNodeConfigurationOptions(ctx context.Context, input *redshift.DescribeNodeConfigurationOptionsInput) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
    return a.client.DescribeNodeConfigurationOptionsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeOrderableClusterOptions(ctx context.Context, input *redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
    return a.client.DescribeOrderableClusterOptionsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeReservedNodeOfferings(ctx context.Context, input *redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
    return a.client.DescribeReservedNodeOfferingsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeReservedNodes(ctx context.Context, input *redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error) {
    return a.client.DescribeReservedNodesWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeResize(ctx context.Context, input *redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error) {
    return a.client.DescribeResizeWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeScheduledActions(ctx context.Context, input *redshift.DescribeScheduledActionsInput) (*redshift.DescribeScheduledActionsOutput, error) {
    return a.client.DescribeScheduledActionsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeSnapshotCopyGrants(ctx context.Context, input *redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
    return a.client.DescribeSnapshotCopyGrantsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeSnapshotSchedules(ctx context.Context, input *redshift.DescribeSnapshotSchedulesInput) (*redshift.DescribeSnapshotSchedulesOutput, error) {
    return a.client.DescribeSnapshotSchedulesWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeStorage(ctx context.Context, input *redshift.DescribeStorageInput) (*redshift.DescribeStorageOutput, error) {
    return a.client.DescribeStorageWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeTableRestoreStatus(ctx context.Context, input *redshift.DescribeTableRestoreStatusInput) (*redshift.DescribeTableRestoreStatusOutput, error) {
    return a.client.DescribeTableRestoreStatusWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeTags(ctx context.Context, input *redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error) {
    return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *RedshiftActivities) DescribeUsageLimits(ctx context.Context, input *redshift.DescribeUsageLimitsInput) (*redshift.DescribeUsageLimitsOutput, error) {
    return a.client.DescribeUsageLimitsWithContext(ctx, input)
}

func (a *RedshiftActivities) DisableLogging(ctx context.Context, input *redshift.DisableLoggingInput) (*redshift.LoggingStatus, error) {
    return a.client.DisableLoggingWithContext(ctx, input)
}

func (a *RedshiftActivities) DisableSnapshotCopy(ctx context.Context, input *redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error) {
    return a.client.DisableSnapshotCopyWithContext(ctx, input)
}

func (a *RedshiftActivities) EnableLogging(ctx context.Context, input *redshift.EnableLoggingInput) (*redshift.LoggingStatus, error) {
    return a.client.EnableLoggingWithContext(ctx, input)
}

func (a *RedshiftActivities) EnableSnapshotCopy(ctx context.Context, input *redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error) {
    return a.client.EnableSnapshotCopyWithContext(ctx, input)
}

func (a *RedshiftActivities) GetClusterCredentials(ctx context.Context, input *redshift.GetClusterCredentialsInput) (*redshift.GetClusterCredentialsOutput, error) {
    return a.client.GetClusterCredentialsWithContext(ctx, input)
}

func (a *RedshiftActivities) GetReservedNodeExchangeOfferings(ctx context.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
    return a.client.GetReservedNodeExchangeOfferingsWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyCluster(ctx context.Context, input *redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error) {
    return a.client.ModifyClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterDbRevision(ctx context.Context, input *redshift.ModifyClusterDbRevisionInput) (*redshift.ModifyClusterDbRevisionOutput, error) {
    return a.client.ModifyClusterDbRevisionWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterIamRoles(ctx context.Context, input *redshift.ModifyClusterIamRolesInput) (*redshift.ModifyClusterIamRolesOutput, error) {
    return a.client.ModifyClusterIamRolesWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterMaintenance(ctx context.Context, input *redshift.ModifyClusterMaintenanceInput) (*redshift.ModifyClusterMaintenanceOutput, error) {
    return a.client.ModifyClusterMaintenanceWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterParameterGroup(ctx context.Context, input *redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error) {
    return a.client.ModifyClusterParameterGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterSnapshot(ctx context.Context, input *redshift.ModifyClusterSnapshotInput) (*redshift.ModifyClusterSnapshotOutput, error) {
    return a.client.ModifyClusterSnapshotWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterSnapshotSchedule(ctx context.Context, input *redshift.ModifyClusterSnapshotScheduleInput) (*redshift.ModifyClusterSnapshotScheduleOutput, error) {
    return a.client.ModifyClusterSnapshotScheduleWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyClusterSubnetGroup(ctx context.Context, input *redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error) {
    return a.client.ModifyClusterSubnetGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyEventSubscription(ctx context.Context, input *redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error) {
    return a.client.ModifyEventSubscriptionWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyScheduledAction(ctx context.Context, input *redshift.ModifyScheduledActionInput) (*redshift.ModifyScheduledActionOutput, error) {
    return a.client.ModifyScheduledActionWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifySnapshotCopyRetentionPeriod(ctx context.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error) {
    return a.client.ModifySnapshotCopyRetentionPeriodWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifySnapshotSchedule(ctx context.Context, input *redshift.ModifySnapshotScheduleInput) (*redshift.ModifySnapshotScheduleOutput, error) {
    return a.client.ModifySnapshotScheduleWithContext(ctx, input)
}

func (a *RedshiftActivities) ModifyUsageLimit(ctx context.Context, input *redshift.ModifyUsageLimitInput) (*redshift.ModifyUsageLimitOutput, error) {
    return a.client.ModifyUsageLimitWithContext(ctx, input)
}

func (a *RedshiftActivities) PauseCluster(ctx context.Context, input *redshift.PauseClusterInput) (*redshift.PauseClusterOutput, error) {
    return a.client.PauseClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) PurchaseReservedNodeOffering(ctx context.Context, input *redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error) {
    return a.client.PurchaseReservedNodeOfferingWithContext(ctx, input)
}

func (a *RedshiftActivities) RebootCluster(ctx context.Context, input *redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error) {
    return a.client.RebootClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) ResetClusterParameterGroup(ctx context.Context, input *redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error) {
    return a.client.ResetClusterParameterGroupWithContext(ctx, input)
}

func (a *RedshiftActivities) ResizeCluster(ctx context.Context, input *redshift.ResizeClusterInput) (*redshift.ResizeClusterOutput, error) {
    return a.client.ResizeClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) RestoreFromClusterSnapshot(ctx context.Context, input *redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error) {
    return a.client.RestoreFromClusterSnapshotWithContext(ctx, input)
}

func (a *RedshiftActivities) RestoreTableFromClusterSnapshot(ctx context.Context, input *redshift.RestoreTableFromClusterSnapshotInput) (*redshift.RestoreTableFromClusterSnapshotOutput, error) {
    return a.client.RestoreTableFromClusterSnapshotWithContext(ctx, input)
}

func (a *RedshiftActivities) ResumeCluster(ctx context.Context, input *redshift.ResumeClusterInput) (*redshift.ResumeClusterOutput, error) {
    return a.client.ResumeClusterWithContext(ctx, input)
}

func (a *RedshiftActivities) RevokeClusterSecurityGroupIngress(ctx context.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error) {
    return a.client.RevokeClusterSecurityGroupIngressWithContext(ctx, input)
}

func (a *RedshiftActivities) RevokeSnapshotAccess(ctx context.Context, input *redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error) {
    return a.client.RevokeSnapshotAccessWithContext(ctx, input)
}

func (a *RedshiftActivities) RotateEncryptionKey(ctx context.Context, input *redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error) {
    return a.client.RotateEncryptionKeyWithContext(ctx, input)
}

func (a *RedshiftActivities) WaitUntilClusterAvailable(ctx context.Context, input *redshift.DescribeClustersInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilClusterAvailableWithContext(ctx, input, options...)
	})
}

func (a *RedshiftActivities) WaitUntilClusterDeleted(ctx context.Context, input *redshift.DescribeClustersInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilClusterDeletedWithContext(ctx, input, options...)
	})
}

func (a *RedshiftActivities) WaitUntilClusterRestored(ctx context.Context, input *redshift.DescribeClustersInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilClusterRestoredWithContext(ctx, input, options...)
	})
}

func (a *RedshiftActivities) WaitUntilSnapshotAvailable(ctx context.Context, input *redshift.DescribeClusterSnapshotsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSnapshotAvailableWithContext(ctx, input, options...)
	})
}
