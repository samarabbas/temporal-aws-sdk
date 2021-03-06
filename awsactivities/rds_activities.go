// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type RDSActivities struct {
	client rdsiface.RDSAPI
}

func NewRDSActivities(session *session.Session, config ...*aws.Config) *RDSActivities {
	client := rds.New(session, config...)
	return &RDSActivities{client: client}
}

func (a *RDSActivities) AddRoleToDBCluster(ctx context.Context, input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error) {
	return a.client.AddRoleToDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) AddRoleToDBInstance(ctx context.Context, input *rds.AddRoleToDBInstanceInput) (*rds.AddRoleToDBInstanceOutput, error) {
	return a.client.AddRoleToDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) AddSourceIdentifierToSubscription(ctx context.Context, input *rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error) {
	return a.client.AddSourceIdentifierToSubscriptionWithContext(ctx, input)
}

func (a *RDSActivities) AddTagsToResource(ctx context.Context, input *rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *RDSActivities) ApplyPendingMaintenanceAction(ctx context.Context, input *rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error) {
	return a.client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *RDSActivities) AuthorizeDBSecurityGroupIngress(ctx context.Context, input *rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error) {
	return a.client.AuthorizeDBSecurityGroupIngressWithContext(ctx, input)
}

func (a *RDSActivities) BacktrackDBCluster(ctx context.Context, input *rds.BacktrackDBClusterInput) (*rds.BacktrackDBClusterOutput, error) {
	return a.client.BacktrackDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) CancelExportTask(ctx context.Context, input *rds.CancelExportTaskInput) (*rds.CancelExportTaskOutput, error) {
	return a.client.CancelExportTaskWithContext(ctx, input)
}

func (a *RDSActivities) CopyDBClusterParameterGroup(ctx context.Context, input *rds.CopyDBClusterParameterGroupInput) (*rds.CopyDBClusterParameterGroupOutput, error) {
	return a.client.CopyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) CopyDBClusterSnapshot(ctx context.Context, input *rds.CopyDBClusterSnapshotInput) (*rds.CopyDBClusterSnapshotOutput, error) {
	return a.client.CopyDBClusterSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) CopyDBParameterGroup(ctx context.Context, input *rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error) {
	return a.client.CopyDBParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) CopyDBSnapshot(ctx context.Context, input *rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error) {
	return a.client.CopyDBSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) CopyOptionGroup(ctx context.Context, input *rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error) {
	return a.client.CopyOptionGroupWithContext(ctx, input)
}

func (a *RDSActivities) CreateCustomAvailabilityZone(ctx context.Context, input *rds.CreateCustomAvailabilityZoneInput) (*rds.CreateCustomAvailabilityZoneOutput, error) {
	return a.client.CreateCustomAvailabilityZoneWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBCluster(ctx context.Context, input *rds.CreateDBClusterInput) (*rds.CreateDBClusterOutput, error) {
	return a.client.CreateDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBClusterEndpoint(ctx context.Context, input *rds.CreateDBClusterEndpointInput) (*rds.CreateDBClusterEndpointOutput, error) {
	return a.client.CreateDBClusterEndpointWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBClusterParameterGroup(ctx context.Context, input *rds.CreateDBClusterParameterGroupInput) (*rds.CreateDBClusterParameterGroupOutput, error) {
	return a.client.CreateDBClusterParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBClusterSnapshot(ctx context.Context, input *rds.CreateDBClusterSnapshotInput) (*rds.CreateDBClusterSnapshotOutput, error) {
	return a.client.CreateDBClusterSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBInstance(ctx context.Context, input *rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error) {
	return a.client.CreateDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBInstanceReadReplica(ctx context.Context, input *rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error) {
	return a.client.CreateDBInstanceReadReplicaWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBParameterGroup(ctx context.Context, input *rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error) {
	return a.client.CreateDBParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBProxy(ctx context.Context, input *rds.CreateDBProxyInput) (*rds.CreateDBProxyOutput, error) {
	return a.client.CreateDBProxyWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBSecurityGroup(ctx context.Context, input *rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error) {
	return a.client.CreateDBSecurityGroupWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBSnapshot(ctx context.Context, input *rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error) {
	return a.client.CreateDBSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) CreateDBSubnetGroup(ctx context.Context, input *rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error) {
	return a.client.CreateDBSubnetGroupWithContext(ctx, input)
}

func (a *RDSActivities) CreateEventSubscription(ctx context.Context, input *rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error) {
	return a.client.CreateEventSubscriptionWithContext(ctx, input)
}

func (a *RDSActivities) CreateGlobalCluster(ctx context.Context, input *rds.CreateGlobalClusterInput) (*rds.CreateGlobalClusterOutput, error) {
	return a.client.CreateGlobalClusterWithContext(ctx, input)
}

func (a *RDSActivities) CreateOptionGroup(ctx context.Context, input *rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error) {
	return a.client.CreateOptionGroupWithContext(ctx, input)
}

func (a *RDSActivities) DeleteCustomAvailabilityZone(ctx context.Context, input *rds.DeleteCustomAvailabilityZoneInput) (*rds.DeleteCustomAvailabilityZoneOutput, error) {
	return a.client.DeleteCustomAvailabilityZoneWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBCluster(ctx context.Context, input *rds.DeleteDBClusterInput) (*rds.DeleteDBClusterOutput, error) {
	return a.client.DeleteDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBClusterEndpoint(ctx context.Context, input *rds.DeleteDBClusterEndpointInput) (*rds.DeleteDBClusterEndpointOutput, error) {
	return a.client.DeleteDBClusterEndpointWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBClusterParameterGroup(ctx context.Context, input *rds.DeleteDBClusterParameterGroupInput) (*rds.DeleteDBClusterParameterGroupOutput, error) {
	return a.client.DeleteDBClusterParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBClusterSnapshot(ctx context.Context, input *rds.DeleteDBClusterSnapshotInput) (*rds.DeleteDBClusterSnapshotOutput, error) {
	return a.client.DeleteDBClusterSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBInstance(ctx context.Context, input *rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error) {
	return a.client.DeleteDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBInstanceAutomatedBackup(ctx context.Context, input *rds.DeleteDBInstanceAutomatedBackupInput) (*rds.DeleteDBInstanceAutomatedBackupOutput, error) {
	return a.client.DeleteDBInstanceAutomatedBackupWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBParameterGroup(ctx context.Context, input *rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error) {
	return a.client.DeleteDBParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBProxy(ctx context.Context, input *rds.DeleteDBProxyInput) (*rds.DeleteDBProxyOutput, error) {
	return a.client.DeleteDBProxyWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBSecurityGroup(ctx context.Context, input *rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error) {
	return a.client.DeleteDBSecurityGroupWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBSnapshot(ctx context.Context, input *rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error) {
	return a.client.DeleteDBSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) DeleteDBSubnetGroup(ctx context.Context, input *rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error) {
	return a.client.DeleteDBSubnetGroupWithContext(ctx, input)
}

func (a *RDSActivities) DeleteEventSubscription(ctx context.Context, input *rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error) {
	return a.client.DeleteEventSubscriptionWithContext(ctx, input)
}

func (a *RDSActivities) DeleteGlobalCluster(ctx context.Context, input *rds.DeleteGlobalClusterInput) (*rds.DeleteGlobalClusterOutput, error) {
	return a.client.DeleteGlobalClusterWithContext(ctx, input)
}

func (a *RDSActivities) DeleteInstallationMedia(ctx context.Context, input *rds.DeleteInstallationMediaInput) (*rds.DeleteInstallationMediaOutput, error) {
	return a.client.DeleteInstallationMediaWithContext(ctx, input)
}

func (a *RDSActivities) DeleteOptionGroup(ctx context.Context, input *rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error) {
	return a.client.DeleteOptionGroupWithContext(ctx, input)
}

func (a *RDSActivities) DeregisterDBProxyTargets(ctx context.Context, input *rds.DeregisterDBProxyTargetsInput) (*rds.DeregisterDBProxyTargetsOutput, error) {
	return a.client.DeregisterDBProxyTargetsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeAccountAttributes(ctx context.Context, input *rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error) {
	return a.client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeCertificates(ctx context.Context, input *rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error) {
	return a.client.DescribeCertificatesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeCustomAvailabilityZones(ctx context.Context, input *rds.DescribeCustomAvailabilityZonesInput) (*rds.DescribeCustomAvailabilityZonesOutput, error) {
	return a.client.DescribeCustomAvailabilityZonesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusterBacktracks(ctx context.Context, input *rds.DescribeDBClusterBacktracksInput) (*rds.DescribeDBClusterBacktracksOutput, error) {
	return a.client.DescribeDBClusterBacktracksWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusterEndpoints(ctx context.Context, input *rds.DescribeDBClusterEndpointsInput) (*rds.DescribeDBClusterEndpointsOutput, error) {
	return a.client.DescribeDBClusterEndpointsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusterParameterGroups(ctx context.Context, input *rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
	return a.client.DescribeDBClusterParameterGroupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusterParameters(ctx context.Context, input *rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error) {
	return a.client.DescribeDBClusterParametersWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusterSnapshotAttributes(ctx context.Context, input *rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
	return a.client.DescribeDBClusterSnapshotAttributesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusterSnapshots(ctx context.Context, input *rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error) {
	return a.client.DescribeDBClusterSnapshotsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBClusters(ctx context.Context, input *rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error) {
	return a.client.DescribeDBClustersWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBEngineVersions(ctx context.Context, input *rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error) {
	return a.client.DescribeDBEngineVersionsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBInstanceAutomatedBackups(ctx context.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
	return a.client.DescribeDBInstanceAutomatedBackupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBInstances(ctx context.Context, input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	return a.client.DescribeDBInstancesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBLogFiles(ctx context.Context, input *rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error) {
	return a.client.DescribeDBLogFilesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBParameterGroups(ctx context.Context, input *rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error) {
	return a.client.DescribeDBParameterGroupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBParameters(ctx context.Context, input *rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error) {
	return a.client.DescribeDBParametersWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBProxies(ctx context.Context, input *rds.DescribeDBProxiesInput) (*rds.DescribeDBProxiesOutput, error) {
	return a.client.DescribeDBProxiesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBProxyTargetGroups(ctx context.Context, input *rds.DescribeDBProxyTargetGroupsInput) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
	return a.client.DescribeDBProxyTargetGroupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBProxyTargets(ctx context.Context, input *rds.DescribeDBProxyTargetsInput) (*rds.DescribeDBProxyTargetsOutput, error) {
	return a.client.DescribeDBProxyTargetsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBSecurityGroups(ctx context.Context, input *rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error) {
	return a.client.DescribeDBSecurityGroupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBSnapshotAttributes(ctx context.Context, input *rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error) {
	return a.client.DescribeDBSnapshotAttributesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBSnapshots(ctx context.Context, input *rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error) {
	return a.client.DescribeDBSnapshotsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeDBSubnetGroups(ctx context.Context, input *rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error) {
	return a.client.DescribeDBSubnetGroupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeEngineDefaultClusterParameters(ctx context.Context, input *rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
	return a.client.DescribeEngineDefaultClusterParametersWithContext(ctx, input)
}

func (a *RDSActivities) DescribeEngineDefaultParameters(ctx context.Context, input *rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error) {
	return a.client.DescribeEngineDefaultParametersWithContext(ctx, input)
}

func (a *RDSActivities) DescribeEventCategories(ctx context.Context, input *rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error) {
	return a.client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeEventSubscriptions(ctx context.Context, input *rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error) {
	return a.client.DescribeEventSubscriptionsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeEvents(ctx context.Context, input *rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeExportTasks(ctx context.Context, input *rds.DescribeExportTasksInput) (*rds.DescribeExportTasksOutput, error) {
	return a.client.DescribeExportTasksWithContext(ctx, input)
}

func (a *RDSActivities) DescribeGlobalClusters(ctx context.Context, input *rds.DescribeGlobalClustersInput) (*rds.DescribeGlobalClustersOutput, error) {
	return a.client.DescribeGlobalClustersWithContext(ctx, input)
}

func (a *RDSActivities) DescribeInstallationMedia(ctx context.Context, input *rds.DescribeInstallationMediaInput) (*rds.DescribeInstallationMediaOutput, error) {
	return a.client.DescribeInstallationMediaWithContext(ctx, input)
}

func (a *RDSActivities) DescribeOptionGroupOptions(ctx context.Context, input *rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error) {
	return a.client.DescribeOptionGroupOptionsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeOptionGroups(ctx context.Context, input *rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error) {
	return a.client.DescribeOptionGroupsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeOrderableDBInstanceOptions(ctx context.Context, input *rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
	return a.client.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)
}

func (a *RDSActivities) DescribePendingMaintenanceActions(ctx context.Context, input *rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error) {
	return a.client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeReservedDBInstances(ctx context.Context, input *rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error) {
	return a.client.DescribeReservedDBInstancesWithContext(ctx, input)
}

func (a *RDSActivities) DescribeReservedDBInstancesOfferings(ctx context.Context, input *rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
	return a.client.DescribeReservedDBInstancesOfferingsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeSourceRegions(ctx context.Context, input *rds.DescribeSourceRegionsInput) (*rds.DescribeSourceRegionsOutput, error) {
	return a.client.DescribeSourceRegionsWithContext(ctx, input)
}

func (a *RDSActivities) DescribeValidDBInstanceModifications(ctx context.Context, input *rds.DescribeValidDBInstanceModificationsInput) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
	return a.client.DescribeValidDBInstanceModificationsWithContext(ctx, input)
}

func (a *RDSActivities) DownloadDBLogFilePortion(ctx context.Context, input *rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error) {
	return a.client.DownloadDBLogFilePortionWithContext(ctx, input)
}

func (a *RDSActivities) FailoverDBCluster(ctx context.Context, input *rds.FailoverDBClusterInput) (*rds.FailoverDBClusterOutput, error) {
	return a.client.FailoverDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) ImportInstallationMedia(ctx context.Context, input *rds.ImportInstallationMediaInput) (*rds.ImportInstallationMediaOutput, error) {
	return a.client.ImportInstallationMediaWithContext(ctx, input)
}

func (a *RDSActivities) ListTagsForResource(ctx context.Context, input *rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *RDSActivities) ModifyCertificates(ctx context.Context, input *rds.ModifyCertificatesInput) (*rds.ModifyCertificatesOutput, error) {
	return a.client.ModifyCertificatesWithContext(ctx, input)
}

func (a *RDSActivities) ModifyCurrentDBClusterCapacity(ctx context.Context, input *rds.ModifyCurrentDBClusterCapacityInput) (*rds.ModifyCurrentDBClusterCapacityOutput, error) {
	return a.client.ModifyCurrentDBClusterCapacityWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBCluster(ctx context.Context, input *rds.ModifyDBClusterInput) (*rds.ModifyDBClusterOutput, error) {
	return a.client.ModifyDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBClusterEndpoint(ctx context.Context, input *rds.ModifyDBClusterEndpointInput) (*rds.ModifyDBClusterEndpointOutput, error) {
	return a.client.ModifyDBClusterEndpointWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBClusterParameterGroup(ctx context.Context, input *rds.ModifyDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error) {
	return a.client.ModifyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBClusterSnapshotAttribute(ctx context.Context, input *rds.ModifyDBClusterSnapshotAttributeInput) (*rds.ModifyDBClusterSnapshotAttributeOutput, error) {
	return a.client.ModifyDBClusterSnapshotAttributeWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBInstance(ctx context.Context, input *rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error) {
	return a.client.ModifyDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBParameterGroup(ctx context.Context, input *rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error) {
	return a.client.ModifyDBParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBProxy(ctx context.Context, input *rds.ModifyDBProxyInput) (*rds.ModifyDBProxyOutput, error) {
	return a.client.ModifyDBProxyWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBProxyTargetGroup(ctx context.Context, input *rds.ModifyDBProxyTargetGroupInput) (*rds.ModifyDBProxyTargetGroupOutput, error) {
	return a.client.ModifyDBProxyTargetGroupWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBSnapshot(ctx context.Context, input *rds.ModifyDBSnapshotInput) (*rds.ModifyDBSnapshotOutput, error) {
	return a.client.ModifyDBSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBSnapshotAttribute(ctx context.Context, input *rds.ModifyDBSnapshotAttributeInput) (*rds.ModifyDBSnapshotAttributeOutput, error) {
	return a.client.ModifyDBSnapshotAttributeWithContext(ctx, input)
}

func (a *RDSActivities) ModifyDBSubnetGroup(ctx context.Context, input *rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error) {
	return a.client.ModifyDBSubnetGroupWithContext(ctx, input)
}

func (a *RDSActivities) ModifyEventSubscription(ctx context.Context, input *rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error) {
	return a.client.ModifyEventSubscriptionWithContext(ctx, input)
}

func (a *RDSActivities) ModifyGlobalCluster(ctx context.Context, input *rds.ModifyGlobalClusterInput) (*rds.ModifyGlobalClusterOutput, error) {
	return a.client.ModifyGlobalClusterWithContext(ctx, input)
}

func (a *RDSActivities) ModifyOptionGroup(ctx context.Context, input *rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error) {
	return a.client.ModifyOptionGroupWithContext(ctx, input)
}

func (a *RDSActivities) PromoteReadReplica(ctx context.Context, input *rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error) {
	return a.client.PromoteReadReplicaWithContext(ctx, input)
}

func (a *RDSActivities) PromoteReadReplicaDBCluster(ctx context.Context, input *rds.PromoteReadReplicaDBClusterInput) (*rds.PromoteReadReplicaDBClusterOutput, error) {
	return a.client.PromoteReadReplicaDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) PurchaseReservedDBInstancesOffering(ctx context.Context, input *rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error) {
	return a.client.PurchaseReservedDBInstancesOfferingWithContext(ctx, input)
}

func (a *RDSActivities) RebootDBInstance(ctx context.Context, input *rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error) {
	return a.client.RebootDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) RegisterDBProxyTargets(ctx context.Context, input *rds.RegisterDBProxyTargetsInput) (*rds.RegisterDBProxyTargetsOutput, error) {
	return a.client.RegisterDBProxyTargetsWithContext(ctx, input)
}

func (a *RDSActivities) RemoveFromGlobalCluster(ctx context.Context, input *rds.RemoveFromGlobalClusterInput) (*rds.RemoveFromGlobalClusterOutput, error) {
	return a.client.RemoveFromGlobalClusterWithContext(ctx, input)
}

func (a *RDSActivities) RemoveRoleFromDBCluster(ctx context.Context, input *rds.RemoveRoleFromDBClusterInput) (*rds.RemoveRoleFromDBClusterOutput, error) {
	return a.client.RemoveRoleFromDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) RemoveRoleFromDBInstance(ctx context.Context, input *rds.RemoveRoleFromDBInstanceInput) (*rds.RemoveRoleFromDBInstanceOutput, error) {
	return a.client.RemoveRoleFromDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) RemoveSourceIdentifierFromSubscription(ctx context.Context, input *rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error) {
	return a.client.RemoveSourceIdentifierFromSubscriptionWithContext(ctx, input)
}

func (a *RDSActivities) RemoveTagsFromResource(ctx context.Context, input *rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *RDSActivities) ResetDBClusterParameterGroup(ctx context.Context, input *rds.ResetDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error) {
	return a.client.ResetDBClusterParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) ResetDBParameterGroup(ctx context.Context, input *rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error) {
	return a.client.ResetDBParameterGroupWithContext(ctx, input)
}

func (a *RDSActivities) RestoreDBClusterFromS3(ctx context.Context, input *rds.RestoreDBClusterFromS3Input) (*rds.RestoreDBClusterFromS3Output, error) {
	return a.client.RestoreDBClusterFromS3WithContext(ctx, input)
}

func (a *RDSActivities) RestoreDBClusterFromSnapshot(ctx context.Context, input *rds.RestoreDBClusterFromSnapshotInput) (*rds.RestoreDBClusterFromSnapshotOutput, error) {
	return a.client.RestoreDBClusterFromSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) RestoreDBClusterToPointInTime(ctx context.Context, input *rds.RestoreDBClusterToPointInTimeInput) (*rds.RestoreDBClusterToPointInTimeOutput, error) {
	return a.client.RestoreDBClusterToPointInTimeWithContext(ctx, input)
}

func (a *RDSActivities) RestoreDBInstanceFromDBSnapshot(ctx context.Context, input *rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error) {
	return a.client.RestoreDBInstanceFromDBSnapshotWithContext(ctx, input)
}

func (a *RDSActivities) RestoreDBInstanceFromS3(ctx context.Context, input *rds.RestoreDBInstanceFromS3Input) (*rds.RestoreDBInstanceFromS3Output, error) {
	return a.client.RestoreDBInstanceFromS3WithContext(ctx, input)
}

func (a *RDSActivities) RestoreDBInstanceToPointInTime(ctx context.Context, input *rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error) {
	return a.client.RestoreDBInstanceToPointInTimeWithContext(ctx, input)
}

func (a *RDSActivities) RevokeDBSecurityGroupIngress(ctx context.Context, input *rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error) {
	return a.client.RevokeDBSecurityGroupIngressWithContext(ctx, input)
}

func (a *RDSActivities) StartActivityStream(ctx context.Context, input *rds.StartActivityStreamInput) (*rds.StartActivityStreamOutput, error) {
	return a.client.StartActivityStreamWithContext(ctx, input)
}

func (a *RDSActivities) StartDBCluster(ctx context.Context, input *rds.StartDBClusterInput) (*rds.StartDBClusterOutput, error) {
	return a.client.StartDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) StartDBInstance(ctx context.Context, input *rds.StartDBInstanceInput) (*rds.StartDBInstanceOutput, error) {
	return a.client.StartDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) StartExportTask(ctx context.Context, input *rds.StartExportTaskInput) (*rds.StartExportTaskOutput, error) {
	return a.client.StartExportTaskWithContext(ctx, input)
}

func (a *RDSActivities) StopActivityStream(ctx context.Context, input *rds.StopActivityStreamInput) (*rds.StopActivityStreamOutput, error) {
	return a.client.StopActivityStreamWithContext(ctx, input)
}

func (a *RDSActivities) StopDBCluster(ctx context.Context, input *rds.StopDBClusterInput) (*rds.StopDBClusterOutput, error) {
	return a.client.StopDBClusterWithContext(ctx, input)
}

func (a *RDSActivities) StopDBInstance(ctx context.Context, input *rds.StopDBInstanceInput) (*rds.StopDBInstanceOutput, error) {
	return a.client.StopDBInstanceWithContext(ctx, input)
}

func (a *RDSActivities) WaitUntilDBClusterSnapshotAvailable(ctx context.Context, input *rds.DescribeDBClusterSnapshotsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBClusterSnapshotAvailableWithContext(ctx, input, options...)
	})
}

func (a *RDSActivities) WaitUntilDBClusterSnapshotDeleted(ctx context.Context, input *rds.DescribeDBClusterSnapshotsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBClusterSnapshotDeletedWithContext(ctx, input, options...)
	})
}

func (a *RDSActivities) WaitUntilDBInstanceAvailable(ctx context.Context, input *rds.DescribeDBInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *RDSActivities) WaitUntilDBInstanceDeleted(ctx context.Context, input *rds.DescribeDBInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBInstanceDeletedWithContext(ctx, input, options...)
	})
}

func (a *RDSActivities) WaitUntilDBSnapshotAvailable(ctx context.Context, input *rds.DescribeDBSnapshotsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBSnapshotAvailableWithContext(ctx, input, options...)
	})
}

func (a *RDSActivities) WaitUntilDBSnapshotDeleted(ctx context.Context, input *rds.DescribeDBSnapshotsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBSnapshotDeletedWithContext(ctx, input, options...)
	})
}
