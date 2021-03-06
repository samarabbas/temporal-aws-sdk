// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/eks"
	"go.temporal.io/sdk/workflow"
)

type EKSClient interface {
	CreateCluster(ctx workflow.Context, input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *eks.CreateClusterInput) *EksCreateClusterResult

	CreateFargateProfile(ctx workflow.Context, input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error)
	CreateFargateProfileAsync(ctx workflow.Context, input *eks.CreateFargateProfileInput) *EksCreateFargateProfileResult

	CreateNodegroup(ctx workflow.Context, input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error)
	CreateNodegroupAsync(ctx workflow.Context, input *eks.CreateNodegroupInput) *EksCreateNodegroupResult

	DeleteCluster(ctx workflow.Context, input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *eks.DeleteClusterInput) *EksDeleteClusterResult

	DeleteFargateProfile(ctx workflow.Context, input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error)
	DeleteFargateProfileAsync(ctx workflow.Context, input *eks.DeleteFargateProfileInput) *EksDeleteFargateProfileResult

	DeleteNodegroup(ctx workflow.Context, input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error)
	DeleteNodegroupAsync(ctx workflow.Context, input *eks.DeleteNodegroupInput) *EksDeleteNodegroupResult

	DescribeCluster(ctx workflow.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error)
	DescribeClusterAsync(ctx workflow.Context, input *eks.DescribeClusterInput) *EksDescribeClusterResult

	DescribeFargateProfile(ctx workflow.Context, input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error)
	DescribeFargateProfileAsync(ctx workflow.Context, input *eks.DescribeFargateProfileInput) *EksDescribeFargateProfileResult

	DescribeNodegroup(ctx workflow.Context, input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error)
	DescribeNodegroupAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) *EksDescribeNodegroupResult

	DescribeUpdate(ctx workflow.Context, input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error)
	DescribeUpdateAsync(ctx workflow.Context, input *eks.DescribeUpdateInput) *EksDescribeUpdateResult

	ListClusters(ctx workflow.Context, input *eks.ListClustersInput) (*eks.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *eks.ListClustersInput) *EksListClustersResult

	ListFargateProfiles(ctx workflow.Context, input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error)
	ListFargateProfilesAsync(ctx workflow.Context, input *eks.ListFargateProfilesInput) *EksListFargateProfilesResult

	ListNodegroups(ctx workflow.Context, input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error)
	ListNodegroupsAsync(ctx workflow.Context, input *eks.ListNodegroupsInput) *EksListNodegroupsResult

	ListTagsForResource(ctx workflow.Context, input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *eks.ListTagsForResourceInput) *EksListTagsForResourceResult

	ListUpdates(ctx workflow.Context, input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error)
	ListUpdatesAsync(ctx workflow.Context, input *eks.ListUpdatesInput) *EksListUpdatesResult

	TagResource(ctx workflow.Context, input *eks.TagResourceInput) (*eks.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *eks.TagResourceInput) *EksTagResourceResult

	UntagResource(ctx workflow.Context, input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *eks.UntagResourceInput) *EksUntagResourceResult

	UpdateClusterConfig(ctx workflow.Context, input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error)
	UpdateClusterConfigAsync(ctx workflow.Context, input *eks.UpdateClusterConfigInput) *EksUpdateClusterConfigResult

	UpdateClusterVersion(ctx workflow.Context, input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error)
	UpdateClusterVersionAsync(ctx workflow.Context, input *eks.UpdateClusterVersionInput) *EksUpdateClusterVersionResult

	UpdateNodegroupConfig(ctx workflow.Context, input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error)
	UpdateNodegroupConfigAsync(ctx workflow.Context, input *eks.UpdateNodegroupConfigInput) *EksUpdateNodegroupConfigResult

	UpdateNodegroupVersion(ctx workflow.Context, input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error)
	UpdateNodegroupVersionAsync(ctx workflow.Context, input *eks.UpdateNodegroupVersionInput) *EksUpdateNodegroupVersionResult

	WaitUntilClusterActive(ctx workflow.Context, input *eks.DescribeClusterInput) error
	WaitUntilClusterDeleted(ctx workflow.Context, input *eks.DescribeClusterInput) error
	WaitUntilNodegroupActive(ctx workflow.Context, input *eks.DescribeNodegroupInput) error
	WaitUntilNodegroupDeleted(ctx workflow.Context, input *eks.DescribeNodegroupInput) error}

type EKSStub struct{}

func NewEKSStub() EKSClient {
	return &EKSStub{}
}

type EksCreateClusterResult struct {
	Result workflow.Future
}

func (r *EksCreateClusterResult) Get(ctx workflow.Context) (*eks.CreateClusterOutput, error) {
	var output eks.CreateClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksCreateFargateProfileResult struct {
	Result workflow.Future
}

func (r *EksCreateFargateProfileResult) Get(ctx workflow.Context) (*eks.CreateFargateProfileOutput, error) {
	var output eks.CreateFargateProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksCreateNodegroupResult struct {
	Result workflow.Future
}

func (r *EksCreateNodegroupResult) Get(ctx workflow.Context) (*eks.CreateNodegroupOutput, error) {
	var output eks.CreateNodegroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDeleteClusterResult struct {
	Result workflow.Future
}

func (r *EksDeleteClusterResult) Get(ctx workflow.Context) (*eks.DeleteClusterOutput, error) {
	var output eks.DeleteClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDeleteFargateProfileResult struct {
	Result workflow.Future
}

func (r *EksDeleteFargateProfileResult) Get(ctx workflow.Context) (*eks.DeleteFargateProfileOutput, error) {
	var output eks.DeleteFargateProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDeleteNodegroupResult struct {
	Result workflow.Future
}

func (r *EksDeleteNodegroupResult) Get(ctx workflow.Context) (*eks.DeleteNodegroupOutput, error) {
	var output eks.DeleteNodegroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDescribeClusterResult struct {
	Result workflow.Future
}

func (r *EksDescribeClusterResult) Get(ctx workflow.Context) (*eks.DescribeClusterOutput, error) {
	var output eks.DescribeClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDescribeFargateProfileResult struct {
	Result workflow.Future
}

func (r *EksDescribeFargateProfileResult) Get(ctx workflow.Context) (*eks.DescribeFargateProfileOutput, error) {
	var output eks.DescribeFargateProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDescribeNodegroupResult struct {
	Result workflow.Future
}

func (r *EksDescribeNodegroupResult) Get(ctx workflow.Context) (*eks.DescribeNodegroupOutput, error) {
	var output eks.DescribeNodegroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksDescribeUpdateResult struct {
	Result workflow.Future
}

func (r *EksDescribeUpdateResult) Get(ctx workflow.Context) (*eks.DescribeUpdateOutput, error) {
	var output eks.DescribeUpdateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksListClustersResult struct {
	Result workflow.Future
}

func (r *EksListClustersResult) Get(ctx workflow.Context) (*eks.ListClustersOutput, error) {
	var output eks.ListClustersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksListFargateProfilesResult struct {
	Result workflow.Future
}

func (r *EksListFargateProfilesResult) Get(ctx workflow.Context) (*eks.ListFargateProfilesOutput, error) {
	var output eks.ListFargateProfilesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksListNodegroupsResult struct {
	Result workflow.Future
}

func (r *EksListNodegroupsResult) Get(ctx workflow.Context) (*eks.ListNodegroupsOutput, error) {
	var output eks.ListNodegroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *EksListTagsForResourceResult) Get(ctx workflow.Context) (*eks.ListTagsForResourceOutput, error) {
	var output eks.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksListUpdatesResult struct {
	Result workflow.Future
}

func (r *EksListUpdatesResult) Get(ctx workflow.Context) (*eks.ListUpdatesOutput, error) {
	var output eks.ListUpdatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksTagResourceResult struct {
	Result workflow.Future
}

func (r *EksTagResourceResult) Get(ctx workflow.Context) (*eks.TagResourceOutput, error) {
	var output eks.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksUntagResourceResult struct {
	Result workflow.Future
}

func (r *EksUntagResourceResult) Get(ctx workflow.Context) (*eks.UntagResourceOutput, error) {
	var output eks.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksUpdateClusterConfigResult struct {
	Result workflow.Future
}

func (r *EksUpdateClusterConfigResult) Get(ctx workflow.Context) (*eks.UpdateClusterConfigOutput, error) {
	var output eks.UpdateClusterConfigOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksUpdateClusterVersionResult struct {
	Result workflow.Future
}

func (r *EksUpdateClusterVersionResult) Get(ctx workflow.Context) (*eks.UpdateClusterVersionOutput, error) {
	var output eks.UpdateClusterVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksUpdateNodegroupConfigResult struct {
	Result workflow.Future
}

func (r *EksUpdateNodegroupConfigResult) Get(ctx workflow.Context) (*eks.UpdateNodegroupConfigOutput, error) {
	var output eks.UpdateNodegroupConfigOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type EksUpdateNodegroupVersionResult struct {
	Result workflow.Future
}

func (r *EksUpdateNodegroupVersionResult) Get(ctx workflow.Context) (*eks.UpdateNodegroupVersionOutput, error) {
	var output eks.UpdateNodegroupVersionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}





func (a *EKSStub) CreateCluster(ctx workflow.Context, input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error) {
	var output eks.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "EKS.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) CreateClusterAsync(ctx workflow.Context, input *eks.CreateClusterInput) *EksCreateClusterResult {
	future := workflow.ExecuteActivity(ctx, "EKS.CreateCluster", input)
	return &EksCreateClusterResult{Result: future}
}

func (a *EKSStub) CreateFargateProfile(ctx workflow.Context, input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error) {
	var output eks.CreateFargateProfileOutput
	err := workflow.ExecuteActivity(ctx, "EKS.CreateFargateProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) CreateFargateProfileAsync(ctx workflow.Context, input *eks.CreateFargateProfileInput) *EksCreateFargateProfileResult {
	future := workflow.ExecuteActivity(ctx, "EKS.CreateFargateProfile", input)
	return &EksCreateFargateProfileResult{Result: future}
}

func (a *EKSStub) CreateNodegroup(ctx workflow.Context, input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error) {
	var output eks.CreateNodegroupOutput
	err := workflow.ExecuteActivity(ctx, "EKS.CreateNodegroup", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) CreateNodegroupAsync(ctx workflow.Context, input *eks.CreateNodegroupInput) *EksCreateNodegroupResult {
	future := workflow.ExecuteActivity(ctx, "EKS.CreateNodegroup", input)
	return &EksCreateNodegroupResult{Result: future}
}

func (a *EKSStub) DeleteCluster(ctx workflow.Context, input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error) {
	var output eks.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DeleteCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DeleteClusterAsync(ctx workflow.Context, input *eks.DeleteClusterInput) *EksDeleteClusterResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DeleteCluster", input)
	return &EksDeleteClusterResult{Result: future}
}

func (a *EKSStub) DeleteFargateProfile(ctx workflow.Context, input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error) {
	var output eks.DeleteFargateProfileOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DeleteFargateProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DeleteFargateProfileAsync(ctx workflow.Context, input *eks.DeleteFargateProfileInput) *EksDeleteFargateProfileResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DeleteFargateProfile", input)
	return &EksDeleteFargateProfileResult{Result: future}
}

func (a *EKSStub) DeleteNodegroup(ctx workflow.Context, input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error) {
	var output eks.DeleteNodegroupOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DeleteNodegroup", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DeleteNodegroupAsync(ctx workflow.Context, input *eks.DeleteNodegroupInput) *EksDeleteNodegroupResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DeleteNodegroup", input)
	return &EksDeleteNodegroupResult{Result: future}
}

func (a *EKSStub) DescribeCluster(ctx workflow.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	var output eks.DescribeClusterOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DescribeCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DescribeClusterAsync(ctx workflow.Context, input *eks.DescribeClusterInput) *EksDescribeClusterResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DescribeCluster", input)
	return &EksDescribeClusterResult{Result: future}
}

func (a *EKSStub) DescribeFargateProfile(ctx workflow.Context, input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error) {
	var output eks.DescribeFargateProfileOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DescribeFargateProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DescribeFargateProfileAsync(ctx workflow.Context, input *eks.DescribeFargateProfileInput) *EksDescribeFargateProfileResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DescribeFargateProfile", input)
	return &EksDescribeFargateProfileResult{Result: future}
}

func (a *EKSStub) DescribeNodegroup(ctx workflow.Context, input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
	var output eks.DescribeNodegroupOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DescribeNodegroup", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DescribeNodegroupAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) *EksDescribeNodegroupResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DescribeNodegroup", input)
	return &EksDescribeNodegroupResult{Result: future}
}

func (a *EKSStub) DescribeUpdate(ctx workflow.Context, input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {
	var output eks.DescribeUpdateOutput
	err := workflow.ExecuteActivity(ctx, "EKS.DescribeUpdate", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) DescribeUpdateAsync(ctx workflow.Context, input *eks.DescribeUpdateInput) *EksDescribeUpdateResult {
	future := workflow.ExecuteActivity(ctx, "EKS.DescribeUpdate", input)
	return &EksDescribeUpdateResult{Result: future}
}

func (a *EKSStub) ListClusters(ctx workflow.Context, input *eks.ListClustersInput) (*eks.ListClustersOutput, error) {
	var output eks.ListClustersOutput
	err := workflow.ExecuteActivity(ctx, "EKS.ListClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) ListClustersAsync(ctx workflow.Context, input *eks.ListClustersInput) *EksListClustersResult {
	future := workflow.ExecuteActivity(ctx, "EKS.ListClusters", input)
	return &EksListClustersResult{Result: future}
}

func (a *EKSStub) ListFargateProfiles(ctx workflow.Context, input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error) {
	var output eks.ListFargateProfilesOutput
	err := workflow.ExecuteActivity(ctx, "EKS.ListFargateProfiles", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) ListFargateProfilesAsync(ctx workflow.Context, input *eks.ListFargateProfilesInput) *EksListFargateProfilesResult {
	future := workflow.ExecuteActivity(ctx, "EKS.ListFargateProfiles", input)
	return &EksListFargateProfilesResult{Result: future}
}

func (a *EKSStub) ListNodegroups(ctx workflow.Context, input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error) {
	var output eks.ListNodegroupsOutput
	err := workflow.ExecuteActivity(ctx, "EKS.ListNodegroups", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) ListNodegroupsAsync(ctx workflow.Context, input *eks.ListNodegroupsInput) *EksListNodegroupsResult {
	future := workflow.ExecuteActivity(ctx, "EKS.ListNodegroups", input)
	return &EksListNodegroupsResult{Result: future}
}

func (a *EKSStub) ListTagsForResource(ctx workflow.Context, input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {
	var output eks.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "EKS.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) ListTagsForResourceAsync(ctx workflow.Context, input *eks.ListTagsForResourceInput) *EksListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, "EKS.ListTagsForResource", input)
	return &EksListTagsForResourceResult{Result: future}
}

func (a *EKSStub) ListUpdates(ctx workflow.Context, input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {
	var output eks.ListUpdatesOutput
	err := workflow.ExecuteActivity(ctx, "EKS.ListUpdates", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) ListUpdatesAsync(ctx workflow.Context, input *eks.ListUpdatesInput) *EksListUpdatesResult {
	future := workflow.ExecuteActivity(ctx, "EKS.ListUpdates", input)
	return &EksListUpdatesResult{Result: future}
}

func (a *EKSStub) TagResource(ctx workflow.Context, input *eks.TagResourceInput) (*eks.TagResourceOutput, error) {
	var output eks.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "EKS.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) TagResourceAsync(ctx workflow.Context, input *eks.TagResourceInput) *EksTagResourceResult {
	future := workflow.ExecuteActivity(ctx, "EKS.TagResource", input)
	return &EksTagResourceResult{Result: future}
}

func (a *EKSStub) UntagResource(ctx workflow.Context, input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {
	var output eks.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "EKS.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) UntagResourceAsync(ctx workflow.Context, input *eks.UntagResourceInput) *EksUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, "EKS.UntagResource", input)
	return &EksUntagResourceResult{Result: future}
}

func (a *EKSStub) UpdateClusterConfig(ctx workflow.Context, input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {
	var output eks.UpdateClusterConfigOutput
	err := workflow.ExecuteActivity(ctx, "EKS.UpdateClusterConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) UpdateClusterConfigAsync(ctx workflow.Context, input *eks.UpdateClusterConfigInput) *EksUpdateClusterConfigResult {
	future := workflow.ExecuteActivity(ctx, "EKS.UpdateClusterConfig", input)
	return &EksUpdateClusterConfigResult{Result: future}
}

func (a *EKSStub) UpdateClusterVersion(ctx workflow.Context, input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error) {
	var output eks.UpdateClusterVersionOutput
	err := workflow.ExecuteActivity(ctx, "EKS.UpdateClusterVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) UpdateClusterVersionAsync(ctx workflow.Context, input *eks.UpdateClusterVersionInput) *EksUpdateClusterVersionResult {
	future := workflow.ExecuteActivity(ctx, "EKS.UpdateClusterVersion", input)
	return &EksUpdateClusterVersionResult{Result: future}
}

func (a *EKSStub) UpdateNodegroupConfig(ctx workflow.Context, input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {
	var output eks.UpdateNodegroupConfigOutput
	err := workflow.ExecuteActivity(ctx, "EKS.UpdateNodegroupConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) UpdateNodegroupConfigAsync(ctx workflow.Context, input *eks.UpdateNodegroupConfigInput) *EksUpdateNodegroupConfigResult {
	future := workflow.ExecuteActivity(ctx, "EKS.UpdateNodegroupConfig", input)
	return &EksUpdateNodegroupConfigResult{Result: future}
}

func (a *EKSStub) UpdateNodegroupVersion(ctx workflow.Context, input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error) {
	var output eks.UpdateNodegroupVersionOutput
	err := workflow.ExecuteActivity(ctx, "EKS.UpdateNodegroupVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *EKSStub) UpdateNodegroupVersionAsync(ctx workflow.Context, input *eks.UpdateNodegroupVersionInput) *EksUpdateNodegroupVersionResult {
	future := workflow.ExecuteActivity(ctx, "EKS.UpdateNodegroupVersion", input)
	return &EksUpdateNodegroupVersionResult{Result: future}
}

func (a *EKSStub) WaitUntilClusterActive(ctx workflow.Context, input *eks.DescribeClusterInput) error {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilClusterActive", input).Get(ctx, nil)
}

func (a *EKSStub) WaitUntilClusterActiveAsync(ctx workflow.Context, input *eks.DescribeClusterInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilClusterActive", input)
}

func (a *EKSStub) WaitUntilClusterDeleted(ctx workflow.Context, input *eks.DescribeClusterInput) error {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilClusterDeleted", input).Get(ctx, nil)
}

func (a *EKSStub) WaitUntilClusterDeletedAsync(ctx workflow.Context, input *eks.DescribeClusterInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilClusterDeleted", input)
}

func (a *EKSStub) WaitUntilNodegroupActive(ctx workflow.Context, input *eks.DescribeNodegroupInput) error {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilNodegroupActive", input).Get(ctx, nil)
}

func (a *EKSStub) WaitUntilNodegroupActiveAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilNodegroupActive", input)
}

func (a *EKSStub) WaitUntilNodegroupDeleted(ctx workflow.Context, input *eks.DescribeNodegroupInput) error {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilNodegroupDeleted", input).Get(ctx, nil)
}

func (a *EKSStub) WaitUntilNodegroupDeletedAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "EKS.WaitUntilNodegroupDeleted", input)
}
