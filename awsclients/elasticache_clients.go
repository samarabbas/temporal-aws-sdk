package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elasticache"
	"go.temporal.io/sdk/workflow"
)

type ElastiCacheClient interface {
       AddTagsToResource(ctx workflow.Context, input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error)
       AddTagsToResourceAsync(ctx workflow.Context, input *elasticache.AddTagsToResourceInput) *ElasticacheAddTagsToResourceResult

       AuthorizeCacheSecurityGroupIngress(ctx workflow.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error)
       AuthorizeCacheSecurityGroupIngressAsync(ctx workflow.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) *ElasticacheAuthorizeCacheSecurityGroupIngressResult

       BatchApplyUpdateAction(ctx workflow.Context, input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error)
       BatchApplyUpdateActionAsync(ctx workflow.Context, input *elasticache.BatchApplyUpdateActionInput) *ElasticacheBatchApplyUpdateActionResult

       BatchStopUpdateAction(ctx workflow.Context, input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error)
       BatchStopUpdateActionAsync(ctx workflow.Context, input *elasticache.BatchStopUpdateActionInput) *ElasticacheBatchStopUpdateActionResult

       CompleteMigration(ctx workflow.Context, input *elasticache.CompleteMigrationInput) (*elasticache.CompleteMigrationOutput, error)
       CompleteMigrationAsync(ctx workflow.Context, input *elasticache.CompleteMigrationInput) *ElasticacheCompleteMigrationResult

       CopySnapshot(ctx workflow.Context, input *elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error)
       CopySnapshotAsync(ctx workflow.Context, input *elasticache.CopySnapshotInput) *ElasticacheCopySnapshotResult

       CreateCacheCluster(ctx workflow.Context, input *elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error)
       CreateCacheClusterAsync(ctx workflow.Context, input *elasticache.CreateCacheClusterInput) *ElasticacheCreateCacheClusterResult

       CreateCacheParameterGroup(ctx workflow.Context, input *elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error)
       CreateCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheParameterGroupInput) *ElasticacheCreateCacheParameterGroupResult

       CreateCacheSecurityGroup(ctx workflow.Context, input *elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error)
       CreateCacheSecurityGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheSecurityGroupInput) *ElasticacheCreateCacheSecurityGroupResult

       CreateCacheSubnetGroup(ctx workflow.Context, input *elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error)
       CreateCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheSubnetGroupInput) *ElasticacheCreateCacheSubnetGroupResult

       CreateGlobalReplicationGroup(ctx workflow.Context, input *elasticache.CreateGlobalReplicationGroupInput) (*elasticache.CreateGlobalReplicationGroupOutput, error)
       CreateGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.CreateGlobalReplicationGroupInput) *ElasticacheCreateGlobalReplicationGroupResult

       CreateReplicationGroup(ctx workflow.Context, input *elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error)
       CreateReplicationGroupAsync(ctx workflow.Context, input *elasticache.CreateReplicationGroupInput) *ElasticacheCreateReplicationGroupResult

       CreateSnapshot(ctx workflow.Context, input *elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error)
       CreateSnapshotAsync(ctx workflow.Context, input *elasticache.CreateSnapshotInput) *ElasticacheCreateSnapshotResult

       DecreaseNodeGroupsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error)
       DecreaseNodeGroupsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) *ElasticacheDecreaseNodeGroupsInGlobalReplicationGroupResult

       DecreaseReplicaCount(ctx workflow.Context, input *elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error)
       DecreaseReplicaCountAsync(ctx workflow.Context, input *elasticache.DecreaseReplicaCountInput) *ElasticacheDecreaseReplicaCountResult

       DeleteCacheCluster(ctx workflow.Context, input *elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error)
       DeleteCacheClusterAsync(ctx workflow.Context, input *elasticache.DeleteCacheClusterInput) *ElasticacheDeleteCacheClusterResult

       DeleteCacheParameterGroup(ctx workflow.Context, input *elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error)
       DeleteCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheParameterGroupInput) *ElasticacheDeleteCacheParameterGroupResult

       DeleteCacheSecurityGroup(ctx workflow.Context, input *elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error)
       DeleteCacheSecurityGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheSecurityGroupInput) *ElasticacheDeleteCacheSecurityGroupResult

       DeleteCacheSubnetGroup(ctx workflow.Context, input *elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error)
       DeleteCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheSubnetGroupInput) *ElasticacheDeleteCacheSubnetGroupResult

       DeleteGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DeleteGlobalReplicationGroupInput) (*elasticache.DeleteGlobalReplicationGroupOutput, error)
       DeleteGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DeleteGlobalReplicationGroupInput) *ElasticacheDeleteGlobalReplicationGroupResult

       DeleteReplicationGroup(ctx workflow.Context, input *elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error)
       DeleteReplicationGroupAsync(ctx workflow.Context, input *elasticache.DeleteReplicationGroupInput) *ElasticacheDeleteReplicationGroupResult

       DeleteSnapshot(ctx workflow.Context, input *elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error)
       DeleteSnapshotAsync(ctx workflow.Context, input *elasticache.DeleteSnapshotInput) *ElasticacheDeleteSnapshotResult

       DescribeCacheClusters(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error)
       DescribeCacheClustersAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) *ElasticacheDescribeCacheClustersResult

       DescribeCacheEngineVersions(ctx workflow.Context, input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error)
       DescribeCacheEngineVersionsAsync(ctx workflow.Context, input *elasticache.DescribeCacheEngineVersionsInput) *ElasticacheDescribeCacheEngineVersionsResult

       DescribeCacheParameterGroups(ctx workflow.Context, input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error)
       DescribeCacheParameterGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheParameterGroupsInput) *ElasticacheDescribeCacheParameterGroupsResult

       DescribeCacheParameters(ctx workflow.Context, input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error)
       DescribeCacheParametersAsync(ctx workflow.Context, input *elasticache.DescribeCacheParametersInput) *ElasticacheDescribeCacheParametersResult

       DescribeCacheSecurityGroups(ctx workflow.Context, input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error)
       DescribeCacheSecurityGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheSecurityGroupsInput) *ElasticacheDescribeCacheSecurityGroupsResult

       DescribeCacheSubnetGroups(ctx workflow.Context, input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
       DescribeCacheSubnetGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheSubnetGroupsInput) *ElasticacheDescribeCacheSubnetGroupsResult

       DescribeEngineDefaultParameters(ctx workflow.Context, input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error)
       DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *elasticache.DescribeEngineDefaultParametersInput) *ElasticacheDescribeEngineDefaultParametersResult

       DescribeEvents(ctx workflow.Context, input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error)
       DescribeEventsAsync(ctx workflow.Context, input *elasticache.DescribeEventsInput) *ElasticacheDescribeEventsResult

       DescribeGlobalReplicationGroups(ctx workflow.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error)
       DescribeGlobalReplicationGroupsAsync(ctx workflow.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) *ElasticacheDescribeGlobalReplicationGroupsResult

       DescribeReplicationGroups(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error)
       DescribeReplicationGroupsAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) *ElasticacheDescribeReplicationGroupsResult

       DescribeReservedCacheNodes(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error)
       DescribeReservedCacheNodesAsync(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesInput) *ElasticacheDescribeReservedCacheNodesResult

       DescribeReservedCacheNodesOfferings(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
       DescribeReservedCacheNodesOfferingsAsync(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) *ElasticacheDescribeReservedCacheNodesOfferingsResult

       DescribeServiceUpdates(ctx workflow.Context, input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error)
       DescribeServiceUpdatesAsync(ctx workflow.Context, input *elasticache.DescribeServiceUpdatesInput) *ElasticacheDescribeServiceUpdatesResult

       DescribeSnapshots(ctx workflow.Context, input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error)
       DescribeSnapshotsAsync(ctx workflow.Context, input *elasticache.DescribeSnapshotsInput) *ElasticacheDescribeSnapshotsResult

       DescribeUpdateActions(ctx workflow.Context, input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error)
       DescribeUpdateActionsAsync(ctx workflow.Context, input *elasticache.DescribeUpdateActionsInput) *ElasticacheDescribeUpdateActionsResult

       DisassociateGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) (*elasticache.DisassociateGlobalReplicationGroupOutput, error)
       DisassociateGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) *ElasticacheDisassociateGlobalReplicationGroupResult

       FailoverGlobalReplicationGroup(ctx workflow.Context, input *elasticache.FailoverGlobalReplicationGroupInput) (*elasticache.FailoverGlobalReplicationGroupOutput, error)
       FailoverGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.FailoverGlobalReplicationGroupInput) *ElasticacheFailoverGlobalReplicationGroupResult

       IncreaseNodeGroupsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error)
       IncreaseNodeGroupsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) *ElasticacheIncreaseNodeGroupsInGlobalReplicationGroupResult

       IncreaseReplicaCount(ctx workflow.Context, input *elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error)
       IncreaseReplicaCountAsync(ctx workflow.Context, input *elasticache.IncreaseReplicaCountInput) *ElasticacheIncreaseReplicaCountResult

       ListAllowedNodeTypeModifications(ctx workflow.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error)
       ListAllowedNodeTypeModificationsAsync(ctx workflow.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) *ElasticacheListAllowedNodeTypeModificationsResult

       ListTagsForResource(ctx workflow.Context, input *elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *elasticache.ListTagsForResourceInput) *ElasticacheListTagsForResourceResult

       ModifyCacheCluster(ctx workflow.Context, input *elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error)
       ModifyCacheClusterAsync(ctx workflow.Context, input *elasticache.ModifyCacheClusterInput) *ElasticacheModifyCacheClusterResult

       ModifyCacheParameterGroup(ctx workflow.Context, input *elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)
       ModifyCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.ModifyCacheParameterGroupInput) *ElasticacheModifyCacheParameterGroupResult

       ModifyCacheSubnetGroup(ctx workflow.Context, input *elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error)
       ModifyCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.ModifyCacheSubnetGroupInput) *ElasticacheModifyCacheSubnetGroupResult

       ModifyGlobalReplicationGroup(ctx workflow.Context, input *elasticache.ModifyGlobalReplicationGroupInput) (*elasticache.ModifyGlobalReplicationGroupOutput, error)
       ModifyGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.ModifyGlobalReplicationGroupInput) *ElasticacheModifyGlobalReplicationGroupResult

       ModifyReplicationGroup(ctx workflow.Context, input *elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error)
       ModifyReplicationGroupAsync(ctx workflow.Context, input *elasticache.ModifyReplicationGroupInput) *ElasticacheModifyReplicationGroupResult

       ModifyReplicationGroupShardConfiguration(ctx workflow.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error)
       ModifyReplicationGroupShardConfigurationAsync(ctx workflow.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) *ElasticacheModifyReplicationGroupShardConfigurationResult

       PurchaseReservedCacheNodesOffering(ctx workflow.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error)
       PurchaseReservedCacheNodesOfferingAsync(ctx workflow.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) *ElasticachePurchaseReservedCacheNodesOfferingResult

       RebalanceSlotsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error)
       RebalanceSlotsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) *ElasticacheRebalanceSlotsInGlobalReplicationGroupResult

       RebootCacheCluster(ctx workflow.Context, input *elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error)
       RebootCacheClusterAsync(ctx workflow.Context, input *elasticache.RebootCacheClusterInput) *ElasticacheRebootCacheClusterResult

       RemoveTagsFromResource(ctx workflow.Context, input *elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error)
       RemoveTagsFromResourceAsync(ctx workflow.Context, input *elasticache.RemoveTagsFromResourceInput) *ElasticacheRemoveTagsFromResourceResult

       ResetCacheParameterGroup(ctx workflow.Context, input *elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)
       ResetCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.ResetCacheParameterGroupInput) *ElasticacheResetCacheParameterGroupResult

       RevokeCacheSecurityGroupIngress(ctx workflow.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error)
       RevokeCacheSecurityGroupIngressAsync(ctx workflow.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) *ElasticacheRevokeCacheSecurityGroupIngressResult

       StartMigration(ctx workflow.Context, input *elasticache.StartMigrationInput) (*elasticache.StartMigrationOutput, error)
       StartMigrationAsync(ctx workflow.Context, input *elasticache.StartMigrationInput) *ElasticacheStartMigrationResult

       TestFailover(ctx workflow.Context, input *elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error)
       TestFailoverAsync(ctx workflow.Context, input *elasticache.TestFailoverInput) *ElasticacheTestFailoverResult

       WaitUntilCacheClusterAvailable(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) error
       WaitUntilCacheClusterDeleted(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) error
       WaitUntilReplicationGroupAvailable(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) error
       WaitUntilReplicationGroupDeleted(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) error}

type ElastiCacheStub struct {
}

func NewElastiCacheStub() ElastiCacheClient {
    return &ElastiCacheStub{}
}

type ElasticacheAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *ElasticacheAddTagsToResourceResult) Get(ctx workflow.Context) (*elasticache.TagListMessage, error) {
    var output elasticache.TagListMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheAuthorizeCacheSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *ElasticacheAuthorizeCacheSecurityGroupIngressResult) Get(ctx workflow.Context) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error) {
    var output elasticache.AuthorizeCacheSecurityGroupIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheBatchApplyUpdateActionResult struct {
	Result workflow.Future
}

func (r *ElasticacheBatchApplyUpdateActionResult) Get(ctx workflow.Context) (*elasticache.BatchApplyUpdateActionOutput, error) {
    var output elasticache.BatchApplyUpdateActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheBatchStopUpdateActionResult struct {
	Result workflow.Future
}

func (r *ElasticacheBatchStopUpdateActionResult) Get(ctx workflow.Context) (*elasticache.BatchStopUpdateActionOutput, error) {
    var output elasticache.BatchStopUpdateActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCompleteMigrationResult struct {
	Result workflow.Future
}

func (r *ElasticacheCompleteMigrationResult) Get(ctx workflow.Context) (*elasticache.CompleteMigrationOutput, error) {
    var output elasticache.CompleteMigrationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCopySnapshotResult struct {
	Result workflow.Future
}

func (r *ElasticacheCopySnapshotResult) Get(ctx workflow.Context) (*elasticache.CopySnapshotOutput, error) {
    var output elasticache.CopySnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateCacheClusterResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateCacheClusterResult) Get(ctx workflow.Context) (*elasticache.CreateCacheClusterOutput, error) {
    var output elasticache.CreateCacheClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateCacheParameterGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateCacheParameterGroupResult) Get(ctx workflow.Context) (*elasticache.CreateCacheParameterGroupOutput, error) {
    var output elasticache.CreateCacheParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateCacheSecurityGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateCacheSecurityGroupResult) Get(ctx workflow.Context) (*elasticache.CreateCacheSecurityGroupOutput, error) {
    var output elasticache.CreateCacheSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateCacheSubnetGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateCacheSubnetGroupResult) Get(ctx workflow.Context) (*elasticache.CreateCacheSubnetGroupOutput, error) {
    var output elasticache.CreateCacheSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.CreateGlobalReplicationGroupOutput, error) {
    var output elasticache.CreateGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.CreateReplicationGroupOutput, error) {
    var output elasticache.CreateReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheCreateSnapshotResult struct {
	Result workflow.Future
}

func (r *ElasticacheCreateSnapshotResult) Get(ctx workflow.Context) (*elasticache.CreateSnapshotOutput, error) {
    var output elasticache.CreateSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDecreaseNodeGroupsInGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDecreaseNodeGroupsInGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
    var output elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDecreaseReplicaCountResult struct {
	Result workflow.Future
}

func (r *ElasticacheDecreaseReplicaCountResult) Get(ctx workflow.Context) (*elasticache.DecreaseReplicaCountOutput, error) {
    var output elasticache.DecreaseReplicaCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteCacheClusterResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteCacheClusterResult) Get(ctx workflow.Context) (*elasticache.DeleteCacheClusterOutput, error) {
    var output elasticache.DeleteCacheClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteCacheParameterGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteCacheParameterGroupResult) Get(ctx workflow.Context) (*elasticache.DeleteCacheParameterGroupOutput, error) {
    var output elasticache.DeleteCacheParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteCacheSecurityGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteCacheSecurityGroupResult) Get(ctx workflow.Context) (*elasticache.DeleteCacheSecurityGroupOutput, error) {
    var output elasticache.DeleteCacheSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteCacheSubnetGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteCacheSubnetGroupResult) Get(ctx workflow.Context) (*elasticache.DeleteCacheSubnetGroupOutput, error) {
    var output elasticache.DeleteCacheSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.DeleteGlobalReplicationGroupOutput, error) {
    var output elasticache.DeleteGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.DeleteReplicationGroupOutput, error) {
    var output elasticache.DeleteReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDeleteSnapshotResult struct {
	Result workflow.Future
}

func (r *ElasticacheDeleteSnapshotResult) Get(ctx workflow.Context) (*elasticache.DeleteSnapshotOutput, error) {
    var output elasticache.DeleteSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeCacheClustersResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeCacheClustersResult) Get(ctx workflow.Context) (*elasticache.DescribeCacheClustersOutput, error) {
    var output elasticache.DescribeCacheClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeCacheEngineVersionsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeCacheEngineVersionsResult) Get(ctx workflow.Context) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
    var output elasticache.DescribeCacheEngineVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeCacheParameterGroupsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeCacheParameterGroupsResult) Get(ctx workflow.Context) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
    var output elasticache.DescribeCacheParameterGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeCacheParametersResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeCacheParametersResult) Get(ctx workflow.Context) (*elasticache.DescribeCacheParametersOutput, error) {
    var output elasticache.DescribeCacheParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeCacheSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeCacheSecurityGroupsResult) Get(ctx workflow.Context) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
    var output elasticache.DescribeCacheSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeCacheSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeCacheSubnetGroupsResult) Get(ctx workflow.Context) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
    var output elasticache.DescribeCacheSubnetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeEngineDefaultParametersResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeEngineDefaultParametersResult) Get(ctx workflow.Context) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
    var output elasticache.DescribeEngineDefaultParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeEventsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeEventsResult) Get(ctx workflow.Context) (*elasticache.DescribeEventsOutput, error) {
    var output elasticache.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeGlobalReplicationGroupsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeGlobalReplicationGroupsResult) Get(ctx workflow.Context) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
    var output elasticache.DescribeGlobalReplicationGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeReplicationGroupsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeReplicationGroupsResult) Get(ctx workflow.Context) (*elasticache.DescribeReplicationGroupsOutput, error) {
    var output elasticache.DescribeReplicationGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeReservedCacheNodesResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeReservedCacheNodesResult) Get(ctx workflow.Context) (*elasticache.DescribeReservedCacheNodesOutput, error) {
    var output elasticache.DescribeReservedCacheNodesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeReservedCacheNodesOfferingsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeReservedCacheNodesOfferingsResult) Get(ctx workflow.Context) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
    var output elasticache.DescribeReservedCacheNodesOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeServiceUpdatesResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeServiceUpdatesResult) Get(ctx workflow.Context) (*elasticache.DescribeServiceUpdatesOutput, error) {
    var output elasticache.DescribeServiceUpdatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeSnapshotsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeSnapshotsResult) Get(ctx workflow.Context) (*elasticache.DescribeSnapshotsOutput, error) {
    var output elasticache.DescribeSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDescribeUpdateActionsResult struct {
	Result workflow.Future
}

func (r *ElasticacheDescribeUpdateActionsResult) Get(ctx workflow.Context) (*elasticache.DescribeUpdateActionsOutput, error) {
    var output elasticache.DescribeUpdateActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheDisassociateGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheDisassociateGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.DisassociateGlobalReplicationGroupOutput, error) {
    var output elasticache.DisassociateGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheFailoverGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheFailoverGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.FailoverGlobalReplicationGroupOutput, error) {
    var output elasticache.FailoverGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheIncreaseNodeGroupsInGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheIncreaseNodeGroupsInGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
    var output elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheIncreaseReplicaCountResult struct {
	Result workflow.Future
}

func (r *ElasticacheIncreaseReplicaCountResult) Get(ctx workflow.Context) (*elasticache.IncreaseReplicaCountOutput, error) {
    var output elasticache.IncreaseReplicaCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheListAllowedNodeTypeModificationsResult struct {
	Result workflow.Future
}

func (r *ElasticacheListAllowedNodeTypeModificationsResult) Get(ctx workflow.Context) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
    var output elasticache.ListAllowedNodeTypeModificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ElasticacheListTagsForResourceResult) Get(ctx workflow.Context) (*elasticache.TagListMessage, error) {
    var output elasticache.TagListMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheModifyCacheClusterResult struct {
	Result workflow.Future
}

func (r *ElasticacheModifyCacheClusterResult) Get(ctx workflow.Context) (*elasticache.ModifyCacheClusterOutput, error) {
    var output elasticache.ModifyCacheClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheModifyCacheParameterGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheModifyCacheParameterGroupResult) Get(ctx workflow.Context) (*elasticache.CacheParameterGroupNameMessage, error) {
    var output elasticache.CacheParameterGroupNameMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheModifyCacheSubnetGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheModifyCacheSubnetGroupResult) Get(ctx workflow.Context) (*elasticache.ModifyCacheSubnetGroupOutput, error) {
    var output elasticache.ModifyCacheSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheModifyGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheModifyGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.ModifyGlobalReplicationGroupOutput, error) {
    var output elasticache.ModifyGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheModifyReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheModifyReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.ModifyReplicationGroupOutput, error) {
    var output elasticache.ModifyReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheModifyReplicationGroupShardConfigurationResult struct {
	Result workflow.Future
}

func (r *ElasticacheModifyReplicationGroupShardConfigurationResult) Get(ctx workflow.Context) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error) {
    var output elasticache.ModifyReplicationGroupShardConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticachePurchaseReservedCacheNodesOfferingResult struct {
	Result workflow.Future
}

func (r *ElasticachePurchaseReservedCacheNodesOfferingResult) Get(ctx workflow.Context) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error) {
    var output elasticache.PurchaseReservedCacheNodesOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheRebalanceSlotsInGlobalReplicationGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheRebalanceSlotsInGlobalReplicationGroupResult) Get(ctx workflow.Context) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error) {
    var output elasticache.RebalanceSlotsInGlobalReplicationGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheRebootCacheClusterResult struct {
	Result workflow.Future
}

func (r *ElasticacheRebootCacheClusterResult) Get(ctx workflow.Context) (*elasticache.RebootCacheClusterOutput, error) {
    var output elasticache.RebootCacheClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *ElasticacheRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*elasticache.TagListMessage, error) {
    var output elasticache.TagListMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheResetCacheParameterGroupResult struct {
	Result workflow.Future
}

func (r *ElasticacheResetCacheParameterGroupResult) Get(ctx workflow.Context) (*elasticache.CacheParameterGroupNameMessage, error) {
    var output elasticache.CacheParameterGroupNameMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheRevokeCacheSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *ElasticacheRevokeCacheSecurityGroupIngressResult) Get(ctx workflow.Context) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error) {
    var output elasticache.RevokeCacheSecurityGroupIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheStartMigrationResult struct {
	Result workflow.Future
}

func (r *ElasticacheStartMigrationResult) Get(ctx workflow.Context) (*elasticache.StartMigrationOutput, error) {
    var output elasticache.StartMigrationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticacheTestFailoverResult struct {
	Result workflow.Future
}

func (r *ElasticacheTestFailoverResult) Get(ctx workflow.Context) (*elasticache.TestFailoverOutput, error) {
    var output elasticache.TestFailoverOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) AddTagsToResource(ctx workflow.Context, input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error) {
    var output elasticache.TagListMessage
    err := workflow.ExecuteActivity(ctx, "ElastiCache.AddTagsToResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) AddTagsToResourceAsync(ctx workflow.Context, input *elasticache.AddTagsToResourceInput) *ElasticacheAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.AddTagsToResource", input)
    return &ElasticacheAddTagsToResourceResult{Result: future}
}

func (a *ElastiCacheStub) AuthorizeCacheSecurityGroupIngress(ctx workflow.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error) {
    var output elasticache.AuthorizeCacheSecurityGroupIngressOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.AuthorizeCacheSecurityGroupIngress", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) AuthorizeCacheSecurityGroupIngressAsync(ctx workflow.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) *ElasticacheAuthorizeCacheSecurityGroupIngressResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.AuthorizeCacheSecurityGroupIngress", input)
    return &ElasticacheAuthorizeCacheSecurityGroupIngressResult{Result: future}
}

func (a *ElastiCacheStub) BatchApplyUpdateAction(ctx workflow.Context, input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error) {
    var output elasticache.BatchApplyUpdateActionOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.BatchApplyUpdateAction", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) BatchApplyUpdateActionAsync(ctx workflow.Context, input *elasticache.BatchApplyUpdateActionInput) *ElasticacheBatchApplyUpdateActionResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.BatchApplyUpdateAction", input)
    return &ElasticacheBatchApplyUpdateActionResult{Result: future}
}

func (a *ElastiCacheStub) BatchStopUpdateAction(ctx workflow.Context, input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error) {
    var output elasticache.BatchStopUpdateActionOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.BatchStopUpdateAction", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) BatchStopUpdateActionAsync(ctx workflow.Context, input *elasticache.BatchStopUpdateActionInput) *ElasticacheBatchStopUpdateActionResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.BatchStopUpdateAction", input)
    return &ElasticacheBatchStopUpdateActionResult{Result: future}
}

func (a *ElastiCacheStub) CompleteMigration(ctx workflow.Context, input *elasticache.CompleteMigrationInput) (*elasticache.CompleteMigrationOutput, error) {
    var output elasticache.CompleteMigrationOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CompleteMigration", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CompleteMigrationAsync(ctx workflow.Context, input *elasticache.CompleteMigrationInput) *ElasticacheCompleteMigrationResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CompleteMigration", input)
    return &ElasticacheCompleteMigrationResult{Result: future}
}

func (a *ElastiCacheStub) CopySnapshot(ctx workflow.Context, input *elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error) {
    var output elasticache.CopySnapshotOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CopySnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CopySnapshotAsync(ctx workflow.Context, input *elasticache.CopySnapshotInput) *ElasticacheCopySnapshotResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CopySnapshot", input)
    return &ElasticacheCopySnapshotResult{Result: future}
}

func (a *ElastiCacheStub) CreateCacheCluster(ctx workflow.Context, input *elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error) {
    var output elasticache.CreateCacheClusterOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateCacheClusterAsync(ctx workflow.Context, input *elasticache.CreateCacheClusterInput) *ElasticacheCreateCacheClusterResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheCluster", input)
    return &ElasticacheCreateCacheClusterResult{Result: future}
}

func (a *ElastiCacheStub) CreateCacheParameterGroup(ctx workflow.Context, input *elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error) {
    var output elasticache.CreateCacheParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheParameterGroupInput) *ElasticacheCreateCacheParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheParameterGroup", input)
    return &ElasticacheCreateCacheParameterGroupResult{Result: future}
}

func (a *ElastiCacheStub) CreateCacheSecurityGroup(ctx workflow.Context, input *elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error) {
    var output elasticache.CreateCacheSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheSecurityGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateCacheSecurityGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheSecurityGroupInput) *ElasticacheCreateCacheSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheSecurityGroup", input)
    return &ElasticacheCreateCacheSecurityGroupResult{Result: future}
}

func (a *ElastiCacheStub) CreateCacheSubnetGroup(ctx workflow.Context, input *elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error) {
    var output elasticache.CreateCacheSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheSubnetGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheSubnetGroupInput) *ElasticacheCreateCacheSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateCacheSubnetGroup", input)
    return &ElasticacheCreateCacheSubnetGroupResult{Result: future}
}

func (a *ElastiCacheStub) CreateGlobalReplicationGroup(ctx workflow.Context, input *elasticache.CreateGlobalReplicationGroupInput) (*elasticache.CreateGlobalReplicationGroupOutput, error) {
    var output elasticache.CreateGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.CreateGlobalReplicationGroupInput) *ElasticacheCreateGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateGlobalReplicationGroup", input)
    return &ElasticacheCreateGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) CreateReplicationGroup(ctx workflow.Context, input *elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error) {
    var output elasticache.CreateReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateReplicationGroupAsync(ctx workflow.Context, input *elasticache.CreateReplicationGroupInput) *ElasticacheCreateReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateReplicationGroup", input)
    return &ElasticacheCreateReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) CreateSnapshot(ctx workflow.Context, input *elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error) {
    var output elasticache.CreateSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.CreateSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) CreateSnapshotAsync(ctx workflow.Context, input *elasticache.CreateSnapshotInput) *ElasticacheCreateSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.CreateSnapshot", input)
    return &ElasticacheCreateSnapshotResult{Result: future}
}

func (a *ElastiCacheStub) DecreaseNodeGroupsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
    var output elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DecreaseNodeGroupsInGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DecreaseNodeGroupsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) *ElasticacheDecreaseNodeGroupsInGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DecreaseNodeGroupsInGlobalReplicationGroup", input)
    return &ElasticacheDecreaseNodeGroupsInGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) DecreaseReplicaCount(ctx workflow.Context, input *elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error) {
    var output elasticache.DecreaseReplicaCountOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DecreaseReplicaCount", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DecreaseReplicaCountAsync(ctx workflow.Context, input *elasticache.DecreaseReplicaCountInput) *ElasticacheDecreaseReplicaCountResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DecreaseReplicaCount", input)
    return &ElasticacheDecreaseReplicaCountResult{Result: future}
}

func (a *ElastiCacheStub) DeleteCacheCluster(ctx workflow.Context, input *elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error) {
    var output elasticache.DeleteCacheClusterOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteCacheClusterAsync(ctx workflow.Context, input *elasticache.DeleteCacheClusterInput) *ElasticacheDeleteCacheClusterResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheCluster", input)
    return &ElasticacheDeleteCacheClusterResult{Result: future}
}

func (a *ElastiCacheStub) DeleteCacheParameterGroup(ctx workflow.Context, input *elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error) {
    var output elasticache.DeleteCacheParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheParameterGroupInput) *ElasticacheDeleteCacheParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheParameterGroup", input)
    return &ElasticacheDeleteCacheParameterGroupResult{Result: future}
}

func (a *ElastiCacheStub) DeleteCacheSecurityGroup(ctx workflow.Context, input *elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error) {
    var output elasticache.DeleteCacheSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheSecurityGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteCacheSecurityGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheSecurityGroupInput) *ElasticacheDeleteCacheSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheSecurityGroup", input)
    return &ElasticacheDeleteCacheSecurityGroupResult{Result: future}
}

func (a *ElastiCacheStub) DeleteCacheSubnetGroup(ctx workflow.Context, input *elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error) {
    var output elasticache.DeleteCacheSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheSubnetGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheSubnetGroupInput) *ElasticacheDeleteCacheSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteCacheSubnetGroup", input)
    return &ElasticacheDeleteCacheSubnetGroupResult{Result: future}
}

func (a *ElastiCacheStub) DeleteGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DeleteGlobalReplicationGroupInput) (*elasticache.DeleteGlobalReplicationGroupOutput, error) {
    var output elasticache.DeleteGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DeleteGlobalReplicationGroupInput) *ElasticacheDeleteGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteGlobalReplicationGroup", input)
    return &ElasticacheDeleteGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) DeleteReplicationGroup(ctx workflow.Context, input *elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error) {
    var output elasticache.DeleteReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteReplicationGroupAsync(ctx workflow.Context, input *elasticache.DeleteReplicationGroupInput) *ElasticacheDeleteReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteReplicationGroup", input)
    return &ElasticacheDeleteReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) DeleteSnapshot(ctx workflow.Context, input *elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error) {
    var output elasticache.DeleteSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DeleteSnapshotAsync(ctx workflow.Context, input *elasticache.DeleteSnapshotInput) *ElasticacheDeleteSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DeleteSnapshot", input)
    return &ElasticacheDeleteSnapshotResult{Result: future}
}

func (a *ElastiCacheStub) DescribeCacheClusters(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error) {
    var output elasticache.DescribeCacheClustersOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheClusters", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeCacheClustersAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) *ElasticacheDescribeCacheClustersResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheClusters", input)
    return &ElasticacheDescribeCacheClustersResult{Result: future}
}

func (a *ElastiCacheStub) DescribeCacheEngineVersions(ctx workflow.Context, input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
    var output elasticache.DescribeCacheEngineVersionsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheEngineVersions", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeCacheEngineVersionsAsync(ctx workflow.Context, input *elasticache.DescribeCacheEngineVersionsInput) *ElasticacheDescribeCacheEngineVersionsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheEngineVersions", input)
    return &ElasticacheDescribeCacheEngineVersionsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeCacheParameterGroups(ctx workflow.Context, input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
    var output elasticache.DescribeCacheParameterGroupsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheParameterGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeCacheParameterGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheParameterGroupsInput) *ElasticacheDescribeCacheParameterGroupsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheParameterGroups", input)
    return &ElasticacheDescribeCacheParameterGroupsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeCacheParameters(ctx workflow.Context, input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error) {
    var output elasticache.DescribeCacheParametersOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeCacheParametersAsync(ctx workflow.Context, input *elasticache.DescribeCacheParametersInput) *ElasticacheDescribeCacheParametersResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheParameters", input)
    return &ElasticacheDescribeCacheParametersResult{Result: future}
}

func (a *ElastiCacheStub) DescribeCacheSecurityGroups(ctx workflow.Context, input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
    var output elasticache.DescribeCacheSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheSecurityGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeCacheSecurityGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheSecurityGroupsInput) *ElasticacheDescribeCacheSecurityGroupsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheSecurityGroups", input)
    return &ElasticacheDescribeCacheSecurityGroupsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeCacheSubnetGroups(ctx workflow.Context, input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
    var output elasticache.DescribeCacheSubnetGroupsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheSubnetGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeCacheSubnetGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheSubnetGroupsInput) *ElasticacheDescribeCacheSubnetGroupsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeCacheSubnetGroups", input)
    return &ElasticacheDescribeCacheSubnetGroupsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeEngineDefaultParameters(ctx workflow.Context, input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
    var output elasticache.DescribeEngineDefaultParametersOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeEngineDefaultParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *elasticache.DescribeEngineDefaultParametersInput) *ElasticacheDescribeEngineDefaultParametersResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeEngineDefaultParameters", input)
    return &ElasticacheDescribeEngineDefaultParametersResult{Result: future}
}

func (a *ElastiCacheStub) DescribeEvents(ctx workflow.Context, input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error) {
    var output elasticache.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeEvents", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeEventsAsync(ctx workflow.Context, input *elasticache.DescribeEventsInput) *ElasticacheDescribeEventsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeEvents", input)
    return &ElasticacheDescribeEventsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeGlobalReplicationGroups(ctx workflow.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
    var output elasticache.DescribeGlobalReplicationGroupsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeGlobalReplicationGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeGlobalReplicationGroupsAsync(ctx workflow.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) *ElasticacheDescribeGlobalReplicationGroupsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeGlobalReplicationGroups", input)
    return &ElasticacheDescribeGlobalReplicationGroupsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeReplicationGroups(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error) {
    var output elasticache.DescribeReplicationGroupsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeReplicationGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeReplicationGroupsAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) *ElasticacheDescribeReplicationGroupsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeReplicationGroups", input)
    return &ElasticacheDescribeReplicationGroupsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeReservedCacheNodes(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error) {
    var output elasticache.DescribeReservedCacheNodesOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeReservedCacheNodes", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeReservedCacheNodesAsync(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesInput) *ElasticacheDescribeReservedCacheNodesResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeReservedCacheNodes", input)
    return &ElasticacheDescribeReservedCacheNodesResult{Result: future}
}

func (a *ElastiCacheStub) DescribeReservedCacheNodesOfferings(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
    var output elasticache.DescribeReservedCacheNodesOfferingsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeReservedCacheNodesOfferings", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeReservedCacheNodesOfferingsAsync(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) *ElasticacheDescribeReservedCacheNodesOfferingsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeReservedCacheNodesOfferings", input)
    return &ElasticacheDescribeReservedCacheNodesOfferingsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeServiceUpdates(ctx workflow.Context, input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error) {
    var output elasticache.DescribeServiceUpdatesOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeServiceUpdates", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeServiceUpdatesAsync(ctx workflow.Context, input *elasticache.DescribeServiceUpdatesInput) *ElasticacheDescribeServiceUpdatesResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeServiceUpdates", input)
    return &ElasticacheDescribeServiceUpdatesResult{Result: future}
}

func (a *ElastiCacheStub) DescribeSnapshots(ctx workflow.Context, input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error) {
    var output elasticache.DescribeSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeSnapshots", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeSnapshotsAsync(ctx workflow.Context, input *elasticache.DescribeSnapshotsInput) *ElasticacheDescribeSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeSnapshots", input)
    return &ElasticacheDescribeSnapshotsResult{Result: future}
}

func (a *ElastiCacheStub) DescribeUpdateActions(ctx workflow.Context, input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error) {
    var output elasticache.DescribeUpdateActionsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeUpdateActions", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DescribeUpdateActionsAsync(ctx workflow.Context, input *elasticache.DescribeUpdateActionsInput) *ElasticacheDescribeUpdateActionsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DescribeUpdateActions", input)
    return &ElasticacheDescribeUpdateActionsResult{Result: future}
}

func (a *ElastiCacheStub) DisassociateGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) (*elasticache.DisassociateGlobalReplicationGroupOutput, error) {
    var output elasticache.DisassociateGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.DisassociateGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) DisassociateGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) *ElasticacheDisassociateGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.DisassociateGlobalReplicationGroup", input)
    return &ElasticacheDisassociateGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) FailoverGlobalReplicationGroup(ctx workflow.Context, input *elasticache.FailoverGlobalReplicationGroupInput) (*elasticache.FailoverGlobalReplicationGroupOutput, error) {
    var output elasticache.FailoverGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.FailoverGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) FailoverGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.FailoverGlobalReplicationGroupInput) *ElasticacheFailoverGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.FailoverGlobalReplicationGroup", input)
    return &ElasticacheFailoverGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) IncreaseNodeGroupsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
    var output elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.IncreaseNodeGroupsInGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) IncreaseNodeGroupsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) *ElasticacheIncreaseNodeGroupsInGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.IncreaseNodeGroupsInGlobalReplicationGroup", input)
    return &ElasticacheIncreaseNodeGroupsInGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) IncreaseReplicaCount(ctx workflow.Context, input *elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error) {
    var output elasticache.IncreaseReplicaCountOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.IncreaseReplicaCount", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) IncreaseReplicaCountAsync(ctx workflow.Context, input *elasticache.IncreaseReplicaCountInput) *ElasticacheIncreaseReplicaCountResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.IncreaseReplicaCount", input)
    return &ElasticacheIncreaseReplicaCountResult{Result: future}
}

func (a *ElastiCacheStub) ListAllowedNodeTypeModifications(ctx workflow.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
    var output elasticache.ListAllowedNodeTypeModificationsOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ListAllowedNodeTypeModifications", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ListAllowedNodeTypeModificationsAsync(ctx workflow.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) *ElasticacheListAllowedNodeTypeModificationsResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ListAllowedNodeTypeModifications", input)
    return &ElasticacheListAllowedNodeTypeModificationsResult{Result: future}
}

func (a *ElastiCacheStub) ListTagsForResource(ctx workflow.Context, input *elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error) {
    var output elasticache.TagListMessage
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ListTagsForResourceAsync(ctx workflow.Context, input *elasticache.ListTagsForResourceInput) *ElasticacheListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ListTagsForResource", input)
    return &ElasticacheListTagsForResourceResult{Result: future}
}

func (a *ElastiCacheStub) ModifyCacheCluster(ctx workflow.Context, input *elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error) {
    var output elasticache.ModifyCacheClusterOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyCacheCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ModifyCacheClusterAsync(ctx workflow.Context, input *elasticache.ModifyCacheClusterInput) *ElasticacheModifyCacheClusterResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyCacheCluster", input)
    return &ElasticacheModifyCacheClusterResult{Result: future}
}

func (a *ElastiCacheStub) ModifyCacheParameterGroup(ctx workflow.Context, input *elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
    var output elasticache.CacheParameterGroupNameMessage
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyCacheParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ModifyCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.ModifyCacheParameterGroupInput) *ElasticacheModifyCacheParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyCacheParameterGroup", input)
    return &ElasticacheModifyCacheParameterGroupResult{Result: future}
}

func (a *ElastiCacheStub) ModifyCacheSubnetGroup(ctx workflow.Context, input *elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error) {
    var output elasticache.ModifyCacheSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyCacheSubnetGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ModifyCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.ModifyCacheSubnetGroupInput) *ElasticacheModifyCacheSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyCacheSubnetGroup", input)
    return &ElasticacheModifyCacheSubnetGroupResult{Result: future}
}

func (a *ElastiCacheStub) ModifyGlobalReplicationGroup(ctx workflow.Context, input *elasticache.ModifyGlobalReplicationGroupInput) (*elasticache.ModifyGlobalReplicationGroupOutput, error) {
    var output elasticache.ModifyGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ModifyGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.ModifyGlobalReplicationGroupInput) *ElasticacheModifyGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyGlobalReplicationGroup", input)
    return &ElasticacheModifyGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) ModifyReplicationGroup(ctx workflow.Context, input *elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error) {
    var output elasticache.ModifyReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ModifyReplicationGroupAsync(ctx workflow.Context, input *elasticache.ModifyReplicationGroupInput) *ElasticacheModifyReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyReplicationGroup", input)
    return &ElasticacheModifyReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) ModifyReplicationGroupShardConfiguration(ctx workflow.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error) {
    var output elasticache.ModifyReplicationGroupShardConfigurationOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyReplicationGroupShardConfiguration", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ModifyReplicationGroupShardConfigurationAsync(ctx workflow.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) *ElasticacheModifyReplicationGroupShardConfigurationResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ModifyReplicationGroupShardConfiguration", input)
    return &ElasticacheModifyReplicationGroupShardConfigurationResult{Result: future}
}

func (a *ElastiCacheStub) PurchaseReservedCacheNodesOffering(ctx workflow.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error) {
    var output elasticache.PurchaseReservedCacheNodesOfferingOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.PurchaseReservedCacheNodesOffering", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) PurchaseReservedCacheNodesOfferingAsync(ctx workflow.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) *ElasticachePurchaseReservedCacheNodesOfferingResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.PurchaseReservedCacheNodesOffering", input)
    return &ElasticachePurchaseReservedCacheNodesOfferingResult{Result: future}
}

func (a *ElastiCacheStub) RebalanceSlotsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error) {
    var output elasticache.RebalanceSlotsInGlobalReplicationGroupOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.RebalanceSlotsInGlobalReplicationGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) RebalanceSlotsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) *ElasticacheRebalanceSlotsInGlobalReplicationGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.RebalanceSlotsInGlobalReplicationGroup", input)
    return &ElasticacheRebalanceSlotsInGlobalReplicationGroupResult{Result: future}
}

func (a *ElastiCacheStub) RebootCacheCluster(ctx workflow.Context, input *elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error) {
    var output elasticache.RebootCacheClusterOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.RebootCacheCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) RebootCacheClusterAsync(ctx workflow.Context, input *elasticache.RebootCacheClusterInput) *ElasticacheRebootCacheClusterResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.RebootCacheCluster", input)
    return &ElasticacheRebootCacheClusterResult{Result: future}
}

func (a *ElastiCacheStub) RemoveTagsFromResource(ctx workflow.Context, input *elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error) {
    var output elasticache.TagListMessage
    err := workflow.ExecuteActivity(ctx, "ElastiCache.RemoveTagsFromResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *elasticache.RemoveTagsFromResourceInput) *ElasticacheRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.RemoveTagsFromResource", input)
    return &ElasticacheRemoveTagsFromResourceResult{Result: future}
}

func (a *ElastiCacheStub) ResetCacheParameterGroup(ctx workflow.Context, input *elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
    var output elasticache.CacheParameterGroupNameMessage
    err := workflow.ExecuteActivity(ctx, "ElastiCache.ResetCacheParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) ResetCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.ResetCacheParameterGroupInput) *ElasticacheResetCacheParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.ResetCacheParameterGroup", input)
    return &ElasticacheResetCacheParameterGroupResult{Result: future}
}

func (a *ElastiCacheStub) RevokeCacheSecurityGroupIngress(ctx workflow.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error) {
    var output elasticache.RevokeCacheSecurityGroupIngressOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.RevokeCacheSecurityGroupIngress", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) RevokeCacheSecurityGroupIngressAsync(ctx workflow.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) *ElasticacheRevokeCacheSecurityGroupIngressResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.RevokeCacheSecurityGroupIngress", input)
    return &ElasticacheRevokeCacheSecurityGroupIngressResult{Result: future}
}

func (a *ElastiCacheStub) StartMigration(ctx workflow.Context, input *elasticache.StartMigrationInput) (*elasticache.StartMigrationOutput, error) {
    var output elasticache.StartMigrationOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.StartMigration", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) StartMigrationAsync(ctx workflow.Context, input *elasticache.StartMigrationInput) *ElasticacheStartMigrationResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.StartMigration", input)
    return &ElasticacheStartMigrationResult{Result: future}
}

func (a *ElastiCacheStub) TestFailover(ctx workflow.Context, input *elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error) {
    var output elasticache.TestFailoverOutput
    err := workflow.ExecuteActivity(ctx, "ElastiCache.TestFailover", input).Get(ctx, &output)
    return &output, err
}

func (a *ElastiCacheStub) TestFailoverAsync(ctx workflow.Context, input *elasticache.TestFailoverInput) *ElasticacheTestFailoverResult {
    future := workflow.ExecuteActivity(ctx, "ElastiCache.TestFailover", input)
    return &ElasticacheTestFailoverResult{Result: future}
}

func (a *ElastiCacheStub) WaitUntilCacheClusterAvailable(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) error {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilCacheClusterAvailable", input).Get(ctx, nil)
}

func (a *ElastiCacheStub) WaitUntilCacheClusterAvailableAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilCacheClusterAvailable", input)
}


func (a *ElastiCacheStub) WaitUntilCacheClusterDeleted(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) error {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilCacheClusterDeleted", input).Get(ctx, nil)
}

func (a *ElastiCacheStub) WaitUntilCacheClusterDeletedAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilCacheClusterDeleted", input)
}


func (a *ElastiCacheStub) WaitUntilReplicationGroupAvailable(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) error {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilReplicationGroupAvailable", input).Get(ctx, nil)
}

func (a *ElastiCacheStub) WaitUntilReplicationGroupAvailableAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilReplicationGroupAvailable", input)
}


func (a *ElastiCacheStub) WaitUntilReplicationGroupDeleted(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) error {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilReplicationGroupDeleted", input).Get(ctx, nil)
}

func (a *ElastiCacheStub) WaitUntilReplicationGroupDeletedAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "ElastiCache.WaitUntilReplicationGroupDeleted", input)
}

