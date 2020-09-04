package awsclients

import (
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ResourceGroupsClient interface {
    CreateGroup(ctx workflow.Context, input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error)
    CreateGroupAsync(ctx workflow.Context, input *resourcegroups.CreateGroupInput) *ResourcegroupsCreateGroupResult

    DeleteGroup(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error)
    DeleteGroupAsync(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) *ResourcegroupsDeleteGroupResult

    GetGroup(ctx workflow.Context, input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error)
    GetGroupAsync(ctx workflow.Context, input *resourcegroups.GetGroupInput) *ResourcegroupsGetGroupResult

    GetGroupConfiguration(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error)
    GetGroupConfigurationAsync(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) *ResourcegroupsGetGroupConfigurationResult

    GetGroupQuery(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error)
    GetGroupQueryAsync(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) *ResourcegroupsGetGroupQueryResult

    GetTags(ctx workflow.Context, input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error)
    GetTagsAsync(ctx workflow.Context, input *resourcegroups.GetTagsInput) *ResourcegroupsGetTagsResult

    GroupResources(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error)
    GroupResourcesAsync(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) *ResourcegroupsGroupResourcesResult

    ListGroupResources(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error)
    ListGroupResourcesAsync(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) *ResourcegroupsListGroupResourcesResult

    ListGroups(ctx workflow.Context, input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error)
    ListGroupsAsync(ctx workflow.Context, input *resourcegroups.ListGroupsInput) *ResourcegroupsListGroupsResult

    SearchResources(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error)
    SearchResourcesAsync(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) *ResourcegroupsSearchResourcesResult

    Tag(ctx workflow.Context, input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error)
    TagAsync(ctx workflow.Context, input *resourcegroups.TagInput) *ResourcegroupsTagResult

    UngroupResources(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error)
    UngroupResourcesAsync(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) *ResourcegroupsUngroupResourcesResult

    Untag(ctx workflow.Context, input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error)
    UntagAsync(ctx workflow.Context, input *resourcegroups.UntagInput) *ResourcegroupsUntagResult

    UpdateGroup(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error)
    UpdateGroupAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) *ResourcegroupsUpdateGroupResult

    UpdateGroupQuery(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error)
    UpdateGroupQueryAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) *ResourcegroupsUpdateGroupQueryResult
}

type ResourcegroupsCreateGroupResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsCreateGroupResult) Get(ctx workflow.Context) (*resourcegroups.CreateGroupOutput, error) {
    var output resourcegroups.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsDeleteGroupResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsDeleteGroupResult) Get(ctx workflow.Context) (*resourcegroups.DeleteGroupOutput, error) {
    var output resourcegroups.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsGetGroupResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsGetGroupResult) Get(ctx workflow.Context) (*resourcegroups.GetGroupOutput, error) {
    var output resourcegroups.GetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsGetGroupConfigurationResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsGetGroupConfigurationResult) Get(ctx workflow.Context) (*resourcegroups.GetGroupConfigurationOutput, error) {
    var output resourcegroups.GetGroupConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsGetGroupQueryResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsGetGroupQueryResult) Get(ctx workflow.Context) (*resourcegroups.GetGroupQueryOutput, error) {
    var output resourcegroups.GetGroupQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsGetTagsResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsGetTagsResult) Get(ctx workflow.Context) (*resourcegroups.GetTagsOutput, error) {
    var output resourcegroups.GetTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsGroupResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsGroupResourcesResult) Get(ctx workflow.Context) (*resourcegroups.GroupResourcesOutput, error) {
    var output resourcegroups.GroupResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsListGroupResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsListGroupResourcesResult) Get(ctx workflow.Context) (*resourcegroups.ListGroupResourcesOutput, error) {
    var output resourcegroups.ListGroupResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsListGroupsResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsListGroupsResult) Get(ctx workflow.Context) (*resourcegroups.ListGroupsOutput, error) {
    var output resourcegroups.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsSearchResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsSearchResourcesResult) Get(ctx workflow.Context) (*resourcegroups.SearchResourcesOutput, error) {
    var output resourcegroups.SearchResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsTagResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsTagResult) Get(ctx workflow.Context) (*resourcegroups.TagOutput, error) {
    var output resourcegroups.TagOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsUngroupResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsUngroupResourcesResult) Get(ctx workflow.Context) (*resourcegroups.UngroupResourcesOutput, error) {
    var output resourcegroups.UngroupResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsUntagResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsUntagResult) Get(ctx workflow.Context) (*resourcegroups.UntagOutput, error) {
    var output resourcegroups.UntagOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsUpdateGroupResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsUpdateGroupResult) Get(ctx workflow.Context) (*resourcegroups.UpdateGroupOutput, error) {
    var output resourcegroups.UpdateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupsUpdateGroupQueryResult struct {
	Result workflow.Future
}

func (r *ResourcegroupsUpdateGroupQueryResult) Get(ctx workflow.Context) (*resourcegroups.UpdateGroupQueryOutput, error) {
    var output resourcegroups.UpdateGroupQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourceGroupsStub struct {
    activities awsactivities.ResourceGroupsActivities
}

func NewResourceGroupsStub() ResourceGroupsClient {
    return &ResourceGroupsStub{}
}

func (a *ResourceGroupsStub) CreateGroup(ctx workflow.Context, input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error) {
    var output resourcegroups.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) CreateGroupAsync(ctx workflow.Context, input *resourcegroups.CreateGroupInput) *ResourcegroupsCreateGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input)
    return &ResourcegroupsCreateGroupResult{Result: future}
}

func (a *ResourceGroupsStub) DeleteGroup(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error) {
    var output resourcegroups.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) DeleteGroupAsync(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) *ResourcegroupsDeleteGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input)
    return &ResourcegroupsDeleteGroupResult{Result: future}
}

func (a *ResourceGroupsStub) GetGroup(ctx workflow.Context, input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error) {
    var output resourcegroups.GetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) GetGroupAsync(ctx workflow.Context, input *resourcegroups.GetGroupInput) *ResourcegroupsGetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input)
    return &ResourcegroupsGetGroupResult{Result: future}
}

func (a *ResourceGroupsStub) GetGroupConfiguration(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error) {
    var output resourcegroups.GetGroupConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) GetGroupConfigurationAsync(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) *ResourcegroupsGetGroupConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGroupConfiguration, input)
    return &ResourcegroupsGetGroupConfigurationResult{Result: future}
}

func (a *ResourceGroupsStub) GetGroupQuery(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error) {
    var output resourcegroups.GetGroupQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) GetGroupQueryAsync(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) *ResourcegroupsGetGroupQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGroupQuery, input)
    return &ResourcegroupsGetGroupQueryResult{Result: future}
}

func (a *ResourceGroupsStub) GetTags(ctx workflow.Context, input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error) {
    var output resourcegroups.GetTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) GetTagsAsync(ctx workflow.Context, input *resourcegroups.GetTagsInput) *ResourcegroupsGetTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTags, input)
    return &ResourcegroupsGetTagsResult{Result: future}
}

func (a *ResourceGroupsStub) GroupResources(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error) {
    var output resourcegroups.GroupResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GroupResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) GroupResourcesAsync(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) *ResourcegroupsGroupResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GroupResources, input)
    return &ResourcegroupsGroupResourcesResult{Result: future}
}

func (a *ResourceGroupsStub) ListGroupResources(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error) {
    var output resourcegroups.ListGroupResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroupResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) ListGroupResourcesAsync(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) *ResourcegroupsListGroupResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGroupResources, input)
    return &ResourcegroupsListGroupResourcesResult{Result: future}
}

func (a *ResourceGroupsStub) ListGroups(ctx workflow.Context, input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error) {
    var output resourcegroups.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) ListGroupsAsync(ctx workflow.Context, input *resourcegroups.ListGroupsInput) *ResourcegroupsListGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input)
    return &ResourcegroupsListGroupsResult{Result: future}
}

func (a *ResourceGroupsStub) SearchResources(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error) {
    var output resourcegroups.SearchResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) SearchResourcesAsync(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) *ResourcegroupsSearchResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchResources, input)
    return &ResourcegroupsSearchResourcesResult{Result: future}
}

func (a *ResourceGroupsStub) Tag(ctx workflow.Context, input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error) {
    var output resourcegroups.TagOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Tag, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) TagAsync(ctx workflow.Context, input *resourcegroups.TagInput) *ResourcegroupsTagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Tag, input)
    return &ResourcegroupsTagResult{Result: future}
}

func (a *ResourceGroupsStub) UngroupResources(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error) {
    var output resourcegroups.UngroupResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UngroupResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) UngroupResourcesAsync(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) *ResourcegroupsUngroupResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UngroupResources, input)
    return &ResourcegroupsUngroupResourcesResult{Result: future}
}

func (a *ResourceGroupsStub) Untag(ctx workflow.Context, input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error) {
    var output resourcegroups.UntagOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Untag, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) UntagAsync(ctx workflow.Context, input *resourcegroups.UntagInput) *ResourcegroupsUntagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Untag, input)
    return &ResourcegroupsUntagResult{Result: future}
}

func (a *ResourceGroupsStub) UpdateGroup(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error) {
    var output resourcegroups.UpdateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) UpdateGroupAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) *ResourcegroupsUpdateGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input)
    return &ResourcegroupsUpdateGroupResult{Result: future}
}

func (a *ResourceGroupsStub) UpdateGroupQuery(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error) {
    var output resourcegroups.UpdateGroupQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroupQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsStub) UpdateGroupQueryAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) *ResourcegroupsUpdateGroupQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGroupQuery, input)
    return &ResourcegroupsUpdateGroupQueryResult{Result: future}
}