// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"go.temporal.io/sdk/workflow"
)

type ApplicationAutoScalingClient interface {
	DeleteScalingPolicy(ctx workflow.Context, input *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyAsync(ctx workflow.Context, input *applicationautoscaling.DeleteScalingPolicyInput) *ApplicationautoscalingDeleteScalingPolicyResult

	DeleteScheduledAction(ctx workflow.Context, input *applicationautoscaling.DeleteScheduledActionInput) (*applicationautoscaling.DeleteScheduledActionOutput, error)
	DeleteScheduledActionAsync(ctx workflow.Context, input *applicationautoscaling.DeleteScheduledActionInput) *ApplicationautoscalingDeleteScheduledActionResult

	DeregisterScalableTarget(ctx workflow.Context, input *applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error)
	DeregisterScalableTargetAsync(ctx workflow.Context, input *applicationautoscaling.DeregisterScalableTargetInput) *ApplicationautoscalingDeregisterScalableTargetResult

	DescribeScalableTargets(ctx workflow.Context, input *applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error)
	DescribeScalableTargetsAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScalableTargetsInput) *ApplicationautoscalingDescribeScalableTargetsResult

	DescribeScalingActivities(ctx workflow.Context, input *applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScalingActivitiesInput) *ApplicationautoscalingDescribeScalingActivitiesResult

	DescribeScalingPolicies(ctx workflow.Context, input *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScalingPoliciesInput) *ApplicationautoscalingDescribeScalingPoliciesResult

	DescribeScheduledActions(ctx workflow.Context, input *applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScheduledActionsInput) *ApplicationautoscalingDescribeScheduledActionsResult

	PutScalingPolicy(ctx workflow.Context, input *applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error)
	PutScalingPolicyAsync(ctx workflow.Context, input *applicationautoscaling.PutScalingPolicyInput) *ApplicationautoscalingPutScalingPolicyResult

	PutScheduledAction(ctx workflow.Context, input *applicationautoscaling.PutScheduledActionInput) (*applicationautoscaling.PutScheduledActionOutput, error)
	PutScheduledActionAsync(ctx workflow.Context, input *applicationautoscaling.PutScheduledActionInput) *ApplicationautoscalingPutScheduledActionResult

	RegisterScalableTarget(ctx workflow.Context, input *applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error)
	RegisterScalableTargetAsync(ctx workflow.Context, input *applicationautoscaling.RegisterScalableTargetInput) *ApplicationautoscalingRegisterScalableTargetResult
}

type ApplicationAutoScalingStub struct{}

func NewApplicationAutoScalingStub() ApplicationAutoScalingClient {
	return &ApplicationAutoScalingStub{}
}

type ApplicationautoscalingDeleteScalingPolicyResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDeleteScalingPolicyResult) Get(ctx workflow.Context) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	var output applicationautoscaling.DeleteScalingPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingDeleteScheduledActionResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDeleteScheduledActionResult) Get(ctx workflow.Context) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	var output applicationautoscaling.DeleteScheduledActionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingDeregisterScalableTargetResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDeregisterScalableTargetResult) Get(ctx workflow.Context) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	var output applicationautoscaling.DeregisterScalableTargetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingDescribeScalableTargetsResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDescribeScalableTargetsResult) Get(ctx workflow.Context) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	var output applicationautoscaling.DescribeScalableTargetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingDescribeScalingActivitiesResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDescribeScalingActivitiesResult) Get(ctx workflow.Context) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	var output applicationautoscaling.DescribeScalingActivitiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingDescribeScalingPoliciesResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDescribeScalingPoliciesResult) Get(ctx workflow.Context) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	var output applicationautoscaling.DescribeScalingPoliciesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingDescribeScheduledActionsResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingDescribeScheduledActionsResult) Get(ctx workflow.Context) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	var output applicationautoscaling.DescribeScheduledActionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingPutScalingPolicyResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingPutScalingPolicyResult) Get(ctx workflow.Context) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	var output applicationautoscaling.PutScalingPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingPutScheduledActionResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingPutScheduledActionResult) Get(ctx workflow.Context) (*applicationautoscaling.PutScheduledActionOutput, error) {
	var output applicationautoscaling.PutScheduledActionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ApplicationautoscalingRegisterScalableTargetResult struct {
	Result workflow.Future
}

func (r *ApplicationautoscalingRegisterScalableTargetResult) Get(ctx workflow.Context) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	var output applicationautoscaling.RegisterScalableTargetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DeleteScalingPolicy(ctx workflow.Context, input *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	var output applicationautoscaling.DeleteScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DeleteScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DeleteScalingPolicyAsync(ctx workflow.Context, input *applicationautoscaling.DeleteScalingPolicyInput) *ApplicationautoscalingDeleteScalingPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DeleteScalingPolicy", input)
	return &ApplicationautoscalingDeleteScalingPolicyResult{Result: future}
}

func (a *ApplicationAutoScalingStub) DeleteScheduledAction(ctx workflow.Context, input *applicationautoscaling.DeleteScheduledActionInput) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	var output applicationautoscaling.DeleteScheduledActionOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DeleteScheduledAction", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DeleteScheduledActionAsync(ctx workflow.Context, input *applicationautoscaling.DeleteScheduledActionInput) *ApplicationautoscalingDeleteScheduledActionResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DeleteScheduledAction", input)
	return &ApplicationautoscalingDeleteScheduledActionResult{Result: future}
}

func (a *ApplicationAutoScalingStub) DeregisterScalableTarget(ctx workflow.Context, input *applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	var output applicationautoscaling.DeregisterScalableTargetOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DeregisterScalableTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DeregisterScalableTargetAsync(ctx workflow.Context, input *applicationautoscaling.DeregisterScalableTargetInput) *ApplicationautoscalingDeregisterScalableTargetResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DeregisterScalableTarget", input)
	return &ApplicationautoscalingDeregisterScalableTargetResult{Result: future}
}

func (a *ApplicationAutoScalingStub) DescribeScalableTargets(ctx workflow.Context, input *applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	var output applicationautoscaling.DescribeScalableTargetsOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScalableTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DescribeScalableTargetsAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScalableTargetsInput) *ApplicationautoscalingDescribeScalableTargetsResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScalableTargets", input)
	return &ApplicationautoscalingDescribeScalableTargetsResult{Result: future}
}

func (a *ApplicationAutoScalingStub) DescribeScalingActivities(ctx workflow.Context, input *applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	var output applicationautoscaling.DescribeScalingActivitiesOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScalingActivities", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DescribeScalingActivitiesAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScalingActivitiesInput) *ApplicationautoscalingDescribeScalingActivitiesResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScalingActivities", input)
	return &ApplicationautoscalingDescribeScalingActivitiesResult{Result: future}
}

func (a *ApplicationAutoScalingStub) DescribeScalingPolicies(ctx workflow.Context, input *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	var output applicationautoscaling.DescribeScalingPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScalingPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DescribeScalingPoliciesAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScalingPoliciesInput) *ApplicationautoscalingDescribeScalingPoliciesResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScalingPolicies", input)
	return &ApplicationautoscalingDescribeScalingPoliciesResult{Result: future}
}

func (a *ApplicationAutoScalingStub) DescribeScheduledActions(ctx workflow.Context, input *applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	var output applicationautoscaling.DescribeScheduledActionsOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScheduledActions", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) DescribeScheduledActionsAsync(ctx workflow.Context, input *applicationautoscaling.DescribeScheduledActionsInput) *ApplicationautoscalingDescribeScheduledActionsResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.DescribeScheduledActions", input)
	return &ApplicationautoscalingDescribeScheduledActionsResult{Result: future}
}

func (a *ApplicationAutoScalingStub) PutScalingPolicy(ctx workflow.Context, input *applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	var output applicationautoscaling.PutScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.PutScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) PutScalingPolicyAsync(ctx workflow.Context, input *applicationautoscaling.PutScalingPolicyInput) *ApplicationautoscalingPutScalingPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.PutScalingPolicy", input)
	return &ApplicationautoscalingPutScalingPolicyResult{Result: future}
}

func (a *ApplicationAutoScalingStub) PutScheduledAction(ctx workflow.Context, input *applicationautoscaling.PutScheduledActionInput) (*applicationautoscaling.PutScheduledActionOutput, error) {
	var output applicationautoscaling.PutScheduledActionOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.PutScheduledAction", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) PutScheduledActionAsync(ctx workflow.Context, input *applicationautoscaling.PutScheduledActionInput) *ApplicationautoscalingPutScheduledActionResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.PutScheduledAction", input)
	return &ApplicationautoscalingPutScheduledActionResult{Result: future}
}

func (a *ApplicationAutoScalingStub) RegisterScalableTarget(ctx workflow.Context, input *applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	var output applicationautoscaling.RegisterScalableTargetOutput
	err := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.RegisterScalableTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *ApplicationAutoScalingStub) RegisterScalableTargetAsync(ctx workflow.Context, input *applicationautoscaling.RegisterScalableTargetInput) *ApplicationautoscalingRegisterScalableTargetResult {
	future := workflow.ExecuteActivity(ctx, "ApplicationAutoScaling.RegisterScalableTarget", input)
	return &ApplicationautoscalingRegisterScalableTargetResult{Result: future}
}
