// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"go.temporal.io/sdk/workflow"
)

type CloudWatchEventsClient interface {
	ActivateEventSource(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error)
	ActivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) *CloudwatcheventsActivateEventSourceResult

	CreateEventBus(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error)
	CreateEventBusAsync(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) *CloudwatcheventsCreateEventBusResult

	CreatePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error)
	CreatePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) *CloudwatcheventsCreatePartnerEventSourceResult

	DeactivateEventSource(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error)
	DeactivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) *CloudwatcheventsDeactivateEventSourceResult

	DeleteEventBus(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error)
	DeleteEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) *CloudwatcheventsDeleteEventBusResult

	DeletePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error)
	DeletePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) *CloudwatcheventsDeletePartnerEventSourceResult

	DeleteRule(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) *CloudwatcheventsDeleteRuleResult

	DescribeEventBus(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error)
	DescribeEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) *CloudwatcheventsDescribeEventBusResult

	DescribeEventSource(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error)
	DescribeEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) *CloudwatcheventsDescribeEventSourceResult

	DescribePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error)
	DescribePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) *CloudwatcheventsDescribePartnerEventSourceResult

	DescribeRule(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error)
	DescribeRuleAsync(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) *CloudwatcheventsDescribeRuleResult

	DisableRule(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error)
	DisableRuleAsync(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) *CloudwatcheventsDisableRuleResult

	EnableRule(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error)
	EnableRuleAsync(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) *CloudwatcheventsEnableRuleResult

	ListEventBuses(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error)
	ListEventBusesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) *CloudwatcheventsListEventBusesResult

	ListEventSources(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error)
	ListEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) *CloudwatcheventsListEventSourcesResult

	ListPartnerEventSourceAccounts(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) *CloudwatcheventsListPartnerEventSourceAccountsResult

	ListPartnerEventSources(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error)
	ListPartnerEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) *CloudwatcheventsListPartnerEventSourcesResult

	ListRuleNamesByTarget(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error)
	ListRuleNamesByTargetAsync(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) *CloudwatcheventsListRuleNamesByTargetResult

	ListRules(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error)
	ListRulesAsync(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) *CloudwatcheventsListRulesResult

	ListTagsForResource(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) *CloudwatcheventsListTagsForResourceResult

	ListTargetsByRule(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error)
	ListTargetsByRuleAsync(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) *CloudwatcheventsListTargetsByRuleResult

	PutEvents(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error)
	PutEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) *CloudwatcheventsPutEventsResult

	PutPartnerEvents(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error)
	PutPartnerEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) *CloudwatcheventsPutPartnerEventsResult

	PutPermission(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error)
	PutPermissionAsync(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) *CloudwatcheventsPutPermissionResult

	PutRule(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error)
	PutRuleAsync(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) *CloudwatcheventsPutRuleResult

	PutTargets(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error)
	PutTargetsAsync(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) *CloudwatcheventsPutTargetsResult

	RemovePermission(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) *CloudwatcheventsRemovePermissionResult

	RemoveTargets(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error)
	RemoveTargetsAsync(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) *CloudwatcheventsRemoveTargetsResult

	TagResource(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) *CloudwatcheventsTagResourceResult

	TestEventPattern(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error)
	TestEventPatternAsync(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) *CloudwatcheventsTestEventPatternResult

	UntagResource(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) *CloudwatcheventsUntagResourceResult
}

type CloudWatchEventsStub struct{}

func NewCloudWatchEventsStub() CloudWatchEventsClient {
	return &CloudWatchEventsStub{}
}

type CloudwatcheventsActivateEventSourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsActivateEventSourceResult) Get(ctx workflow.Context) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	var output cloudwatchevents.ActivateEventSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsCreateEventBusResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsCreateEventBusResult) Get(ctx workflow.Context) (*cloudwatchevents.CreateEventBusOutput, error) {
	var output cloudwatchevents.CreateEventBusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsCreatePartnerEventSourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsCreatePartnerEventSourceResult) Get(ctx workflow.Context) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	var output cloudwatchevents.CreatePartnerEventSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDeactivateEventSourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDeactivateEventSourceResult) Get(ctx workflow.Context) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	var output cloudwatchevents.DeactivateEventSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDeleteEventBusResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDeleteEventBusResult) Get(ctx workflow.Context) (*cloudwatchevents.DeleteEventBusOutput, error) {
	var output cloudwatchevents.DeleteEventBusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDeletePartnerEventSourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDeletePartnerEventSourceResult) Get(ctx workflow.Context) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DeletePartnerEventSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDeleteRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDeleteRuleResult) Get(ctx workflow.Context) (*cloudwatchevents.DeleteRuleOutput, error) {
	var output cloudwatchevents.DeleteRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDescribeEventBusResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDescribeEventBusResult) Get(ctx workflow.Context) (*cloudwatchevents.DescribeEventBusOutput, error) {
	var output cloudwatchevents.DescribeEventBusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDescribeEventSourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDescribeEventSourceResult) Get(ctx workflow.Context) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	var output cloudwatchevents.DescribeEventSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDescribePartnerEventSourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDescribePartnerEventSourceResult) Get(ctx workflow.Context) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DescribePartnerEventSourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDescribeRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDescribeRuleResult) Get(ctx workflow.Context) (*cloudwatchevents.DescribeRuleOutput, error) {
	var output cloudwatchevents.DescribeRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsDisableRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsDisableRuleResult) Get(ctx workflow.Context) (*cloudwatchevents.DisableRuleOutput, error) {
	var output cloudwatchevents.DisableRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsEnableRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsEnableRuleResult) Get(ctx workflow.Context) (*cloudwatchevents.EnableRuleOutput, error) {
	var output cloudwatchevents.EnableRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListEventBusesResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListEventBusesResult) Get(ctx workflow.Context) (*cloudwatchevents.ListEventBusesOutput, error) {
	var output cloudwatchevents.ListEventBusesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListEventSourcesResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListEventSourcesResult) Get(ctx workflow.Context) (*cloudwatchevents.ListEventSourcesOutput, error) {
	var output cloudwatchevents.ListEventSourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListPartnerEventSourceAccountsResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListPartnerEventSourceAccountsResult) Get(ctx workflow.Context) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourceAccountsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListPartnerEventSourcesResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListPartnerEventSourcesResult) Get(ctx workflow.Context) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListRuleNamesByTargetResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListRuleNamesByTargetResult) Get(ctx workflow.Context) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	var output cloudwatchevents.ListRuleNamesByTargetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListRulesResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListRulesResult) Get(ctx workflow.Context) (*cloudwatchevents.ListRulesOutput, error) {
	var output cloudwatchevents.ListRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListTagsForResourceResult) Get(ctx workflow.Context) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	var output cloudwatchevents.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsListTargetsByRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsListTargetsByRuleResult) Get(ctx workflow.Context) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	var output cloudwatchevents.ListTargetsByRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsPutEventsResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsPutEventsResult) Get(ctx workflow.Context) (*cloudwatchevents.PutEventsOutput, error) {
	var output cloudwatchevents.PutEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsPutPartnerEventsResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsPutPartnerEventsResult) Get(ctx workflow.Context) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	var output cloudwatchevents.PutPartnerEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsPutPermissionResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsPutPermissionResult) Get(ctx workflow.Context) (*cloudwatchevents.PutPermissionOutput, error) {
	var output cloudwatchevents.PutPermissionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsPutRuleResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsPutRuleResult) Get(ctx workflow.Context) (*cloudwatchevents.PutRuleOutput, error) {
	var output cloudwatchevents.PutRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsPutTargetsResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsPutTargetsResult) Get(ctx workflow.Context) (*cloudwatchevents.PutTargetsOutput, error) {
	var output cloudwatchevents.PutTargetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsRemovePermissionResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsRemovePermissionResult) Get(ctx workflow.Context) (*cloudwatchevents.RemovePermissionOutput, error) {
	var output cloudwatchevents.RemovePermissionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsRemoveTargetsResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsRemoveTargetsResult) Get(ctx workflow.Context) (*cloudwatchevents.RemoveTargetsOutput, error) {
	var output cloudwatchevents.RemoveTargetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsTagResourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsTagResourceResult) Get(ctx workflow.Context) (*cloudwatchevents.TagResourceOutput, error) {
	var output cloudwatchevents.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsTestEventPatternResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsTestEventPatternResult) Get(ctx workflow.Context) (*cloudwatchevents.TestEventPatternOutput, error) {
	var output cloudwatchevents.TestEventPatternOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudwatcheventsUntagResourceResult struct {
	Result workflow.Future
}

func (r *CloudwatcheventsUntagResourceResult) Get(ctx workflow.Context) (*cloudwatchevents.UntagResourceOutput, error) {
	var output cloudwatchevents.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ActivateEventSource(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	var output cloudwatchevents.ActivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ActivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ActivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) *CloudwatcheventsActivateEventSourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ActivateEventSource", input)
	return &CloudwatcheventsActivateEventSourceResult{Result: future}
}

func (a *CloudWatchEventsStub) CreateEventBus(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error) {
	var output cloudwatchevents.CreateEventBusOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.CreateEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) CreateEventBusAsync(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) *CloudwatcheventsCreateEventBusResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.CreateEventBus", input)
	return &CloudwatcheventsCreateEventBusResult{Result: future}
}

func (a *CloudWatchEventsStub) CreatePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	var output cloudwatchevents.CreatePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.CreatePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) CreatePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) *CloudwatcheventsCreatePartnerEventSourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.CreatePartnerEventSource", input)
	return &CloudwatcheventsCreatePartnerEventSourceResult{Result: future}
}

func (a *CloudWatchEventsStub) DeactivateEventSource(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	var output cloudwatchevents.DeactivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeactivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DeactivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) *CloudwatcheventsDeactivateEventSourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeactivateEventSource", input)
	return &CloudwatcheventsDeactivateEventSourceResult{Result: future}
}

func (a *CloudWatchEventsStub) DeleteEventBus(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error) {
	var output cloudwatchevents.DeleteEventBusOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeleteEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DeleteEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) *CloudwatcheventsDeleteEventBusResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeleteEventBus", input)
	return &CloudwatcheventsDeleteEventBusResult{Result: future}
}

func (a *CloudWatchEventsStub) DeletePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DeletePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeletePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DeletePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) *CloudwatcheventsDeletePartnerEventSourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeletePartnerEventSource", input)
	return &CloudwatcheventsDeletePartnerEventSourceResult{Result: future}
}

func (a *CloudWatchEventsStub) DeleteRule(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error) {
	var output cloudwatchevents.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeleteRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DeleteRuleAsync(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) *CloudwatcheventsDeleteRuleResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DeleteRule", input)
	return &CloudwatcheventsDeleteRuleResult{Result: future}
}

func (a *CloudWatchEventsStub) DescribeEventBus(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	var output cloudwatchevents.DescribeEventBusOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribeEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DescribeEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) *CloudwatcheventsDescribeEventBusResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribeEventBus", input)
	return &CloudwatcheventsDescribeEventBusResult{Result: future}
}

func (a *CloudWatchEventsStub) DescribeEventSource(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	var output cloudwatchevents.DescribeEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribeEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DescribeEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) *CloudwatcheventsDescribeEventSourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribeEventSource", input)
	return &CloudwatcheventsDescribeEventSourceResult{Result: future}
}

func (a *CloudWatchEventsStub) DescribePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	var output cloudwatchevents.DescribePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DescribePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) *CloudwatcheventsDescribePartnerEventSourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribePartnerEventSource", input)
	return &CloudwatcheventsDescribePartnerEventSourceResult{Result: future}
}

func (a *CloudWatchEventsStub) DescribeRule(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	var output cloudwatchevents.DescribeRuleOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribeRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DescribeRuleAsync(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) *CloudwatcheventsDescribeRuleResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DescribeRule", input)
	return &CloudwatcheventsDescribeRuleResult{Result: future}
}

func (a *CloudWatchEventsStub) DisableRule(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error) {
	var output cloudwatchevents.DisableRuleOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DisableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) DisableRuleAsync(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) *CloudwatcheventsDisableRuleResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.DisableRule", input)
	return &CloudwatcheventsDisableRuleResult{Result: future}
}

func (a *CloudWatchEventsStub) EnableRule(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error) {
	var output cloudwatchevents.EnableRuleOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.EnableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) EnableRuleAsync(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) *CloudwatcheventsEnableRuleResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.EnableRule", input)
	return &CloudwatcheventsEnableRuleResult{Result: future}
}

func (a *CloudWatchEventsStub) ListEventBuses(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error) {
	var output cloudwatchevents.ListEventBusesOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListEventBuses", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListEventBusesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) *CloudwatcheventsListEventBusesResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListEventBuses", input)
	return &CloudwatcheventsListEventBusesResult{Result: future}
}

func (a *CloudWatchEventsStub) ListEventSources(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error) {
	var output cloudwatchevents.ListEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) *CloudwatcheventsListEventSourcesResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListEventSources", input)
	return &CloudwatcheventsListEventSourcesResult{Result: future}
}

func (a *CloudWatchEventsStub) ListPartnerEventSourceAccounts(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourceAccountsOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListPartnerEventSourceAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) *CloudwatcheventsListPartnerEventSourceAccountsResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListPartnerEventSourceAccounts", input)
	return &CloudwatcheventsListPartnerEventSourceAccountsResult{Result: future}
}

func (a *CloudWatchEventsStub) ListPartnerEventSources(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	var output cloudwatchevents.ListPartnerEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListPartnerEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListPartnerEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) *CloudwatcheventsListPartnerEventSourcesResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListPartnerEventSources", input)
	return &CloudwatcheventsListPartnerEventSourcesResult{Result: future}
}

func (a *CloudWatchEventsStub) ListRuleNamesByTarget(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	var output cloudwatchevents.ListRuleNamesByTargetOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListRuleNamesByTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListRuleNamesByTargetAsync(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) *CloudwatcheventsListRuleNamesByTargetResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListRuleNamesByTarget", input)
	return &CloudwatcheventsListRuleNamesByTargetResult{Result: future}
}

func (a *CloudWatchEventsStub) ListRules(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	var output cloudwatchevents.ListRulesOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListRules", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListRulesAsync(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) *CloudwatcheventsListRulesResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListRules", input)
	return &CloudwatcheventsListRulesResult{Result: future}
}

func (a *CloudWatchEventsStub) ListTagsForResource(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	var output cloudwatchevents.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) *CloudwatcheventsListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListTagsForResource", input)
	return &CloudwatcheventsListTagsForResourceResult{Result: future}
}

func (a *CloudWatchEventsStub) ListTargetsByRule(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	var output cloudwatchevents.ListTargetsByRuleOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListTargetsByRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) ListTargetsByRuleAsync(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) *CloudwatcheventsListTargetsByRuleResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.ListTargetsByRule", input)
	return &CloudwatcheventsListTargetsByRuleResult{Result: future}
}

func (a *CloudWatchEventsStub) PutEvents(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error) {
	var output cloudwatchevents.PutEventsOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) PutEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) *CloudwatcheventsPutEventsResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutEvents", input)
	return &CloudwatcheventsPutEventsResult{Result: future}
}

func (a *CloudWatchEventsStub) PutPartnerEvents(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	var output cloudwatchevents.PutPartnerEventsOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutPartnerEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) PutPartnerEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) *CloudwatcheventsPutPartnerEventsResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutPartnerEvents", input)
	return &CloudwatcheventsPutPartnerEventsResult{Result: future}
}

func (a *CloudWatchEventsStub) PutPermission(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error) {
	var output cloudwatchevents.PutPermissionOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) PutPermissionAsync(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) *CloudwatcheventsPutPermissionResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutPermission", input)
	return &CloudwatcheventsPutPermissionResult{Result: future}
}

func (a *CloudWatchEventsStub) PutRule(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error) {
	var output cloudwatchevents.PutRuleOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) PutRuleAsync(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) *CloudwatcheventsPutRuleResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutRule", input)
	return &CloudwatcheventsPutRuleResult{Result: future}
}

func (a *CloudWatchEventsStub) PutTargets(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error) {
	var output cloudwatchevents.PutTargetsOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) PutTargetsAsync(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) *CloudwatcheventsPutTargetsResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.PutTargets", input)
	return &CloudwatcheventsPutTargetsResult{Result: future}
}

func (a *CloudWatchEventsStub) RemovePermission(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error) {
	var output cloudwatchevents.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) RemovePermissionAsync(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) *CloudwatcheventsRemovePermissionResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.RemovePermission", input)
	return &CloudwatcheventsRemovePermissionResult{Result: future}
}

func (a *CloudWatchEventsStub) RemoveTargets(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error) {
	var output cloudwatchevents.RemoveTargetsOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.RemoveTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) RemoveTargetsAsync(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) *CloudwatcheventsRemoveTargetsResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.RemoveTargets", input)
	return &CloudwatcheventsRemoveTargetsResult{Result: future}
}

func (a *CloudWatchEventsStub) TagResource(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error) {
	var output cloudwatchevents.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) TagResourceAsync(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) *CloudwatcheventsTagResourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.TagResource", input)
	return &CloudwatcheventsTagResourceResult{Result: future}
}

func (a *CloudWatchEventsStub) TestEventPattern(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error) {
	var output cloudwatchevents.TestEventPatternOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.TestEventPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) TestEventPatternAsync(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) *CloudwatcheventsTestEventPatternResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.TestEventPattern", input)
	return &CloudwatcheventsTestEventPatternResult{Result: future}
}

func (a *CloudWatchEventsStub) UntagResource(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error) {
	var output cloudwatchevents.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "CloudWatchEvents.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudWatchEventsStub) UntagResourceAsync(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) *CloudwatcheventsUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, "CloudWatchEvents.UntagResource", input)
	return &CloudwatcheventsUntagResourceResult{Result: future}
}
