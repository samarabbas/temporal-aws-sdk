// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"go.temporal.io/sdk/workflow"
)

type ELBV2Client interface {
	AddListenerCertificates(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error)
	AddListenerCertificatesAsync(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) *Elbv2AddListenerCertificatesResult

	AddTags(ctx workflow.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *elbv2.AddTagsInput) *Elbv2AddTagsResult

	CreateListener(ctx workflow.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error)
	CreateListenerAsync(ctx workflow.Context, input *elbv2.CreateListenerInput) *Elbv2CreateListenerResult

	CreateLoadBalancer(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) *Elbv2CreateLoadBalancerResult

	CreateRule(ctx workflow.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *elbv2.CreateRuleInput) *Elbv2CreateRuleResult

	CreateTargetGroup(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error)
	CreateTargetGroupAsync(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) *Elbv2CreateTargetGroupResult

	DeleteListener(ctx workflow.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error)
	DeleteListenerAsync(ctx workflow.Context, input *elbv2.DeleteListenerInput) *Elbv2DeleteListenerResult

	DeleteLoadBalancer(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) *Elbv2DeleteLoadBalancerResult

	DeleteRule(ctx workflow.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *elbv2.DeleteRuleInput) *Elbv2DeleteRuleResult

	DeleteTargetGroup(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error)
	DeleteTargetGroupAsync(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) *Elbv2DeleteTargetGroupResult

	DeregisterTargets(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error)
	DeregisterTargetsAsync(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) *Elbv2DeregisterTargetsResult

	DescribeAccountLimits(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) *Elbv2DescribeAccountLimitsResult

	DescribeListenerCertificates(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error)
	DescribeListenerCertificatesAsync(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) *Elbv2DescribeListenerCertificatesResult

	DescribeListeners(ctx workflow.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error)
	DescribeListenersAsync(ctx workflow.Context, input *elbv2.DescribeListenersInput) *Elbv2DescribeListenersResult

	DescribeLoadBalancerAttributes(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) *Elbv2DescribeLoadBalancerAttributesResult

	DescribeLoadBalancers(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *Elbv2DescribeLoadBalancersResult

	DescribeRules(ctx workflow.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error)
	DescribeRulesAsync(ctx workflow.Context, input *elbv2.DescribeRulesInput) *Elbv2DescribeRulesResult

	DescribeSSLPolicies(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error)
	DescribeSSLPoliciesAsync(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) *Elbv2DescribeSSLPoliciesResult

	DescribeTags(ctx workflow.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *elbv2.DescribeTagsInput) *Elbv2DescribeTagsResult

	DescribeTargetGroupAttributes(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error)
	DescribeTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) *Elbv2DescribeTargetGroupAttributesResult

	DescribeTargetGroups(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTargetGroupsAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) *Elbv2DescribeTargetGroupsResult

	DescribeTargetHealth(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTargetHealthAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *Elbv2DescribeTargetHealthResult

	ModifyListener(ctx workflow.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error)
	ModifyListenerAsync(ctx workflow.Context, input *elbv2.ModifyListenerInput) *Elbv2ModifyListenerResult

	ModifyLoadBalancerAttributes(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) *Elbv2ModifyLoadBalancerAttributesResult

	ModifyRule(ctx workflow.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error)
	ModifyRuleAsync(ctx workflow.Context, input *elbv2.ModifyRuleInput) *Elbv2ModifyRuleResult

	ModifyTargetGroup(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error)
	ModifyTargetGroupAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) *Elbv2ModifyTargetGroupResult

	ModifyTargetGroupAttributes(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error)
	ModifyTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) *Elbv2ModifyTargetGroupAttributesResult

	RegisterTargets(ctx workflow.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error)
	RegisterTargetsAsync(ctx workflow.Context, input *elbv2.RegisterTargetsInput) *Elbv2RegisterTargetsResult

	RemoveListenerCertificates(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error)
	RemoveListenerCertificatesAsync(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) *Elbv2RemoveListenerCertificatesResult

	RemoveTags(ctx workflow.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *elbv2.RemoveTagsInput) *Elbv2RemoveTagsResult

	SetIpAddressType(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error)
	SetIpAddressTypeAsync(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) *Elbv2SetIpAddressTypeResult

	SetRulePriorities(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error)
	SetRulePrioritiesAsync(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) *Elbv2SetRulePrioritiesResult

	SetSecurityGroups(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error)
	SetSecurityGroupsAsync(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) *Elbv2SetSecurityGroupsResult

	SetSubnets(ctx workflow.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error)
	SetSubnetsAsync(ctx workflow.Context, input *elbv2.SetSubnetsInput) *Elbv2SetSubnetsResult

	WaitUntilLoadBalancerAvailable(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerExists(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancersDeleted(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilTargetDeregistered(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error
	WaitUntilTargetInService(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error}

type ELBV2Stub struct{}

func NewELBV2Stub() ELBV2Client {
	return &ELBV2Stub{}
}

type Elbv2AddListenerCertificatesResult struct {
	Result workflow.Future
}

func (r *Elbv2AddListenerCertificatesResult) Get(ctx workflow.Context) (*elbv2.AddListenerCertificatesOutput, error) {
	var output elbv2.AddListenerCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2AddTagsResult struct {
	Result workflow.Future
}

func (r *Elbv2AddTagsResult) Get(ctx workflow.Context) (*elbv2.AddTagsOutput, error) {
	var output elbv2.AddTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2CreateListenerResult struct {
	Result workflow.Future
}

func (r *Elbv2CreateListenerResult) Get(ctx workflow.Context) (*elbv2.CreateListenerOutput, error) {
	var output elbv2.CreateListenerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2CreateLoadBalancerResult struct {
	Result workflow.Future
}

func (r *Elbv2CreateLoadBalancerResult) Get(ctx workflow.Context) (*elbv2.CreateLoadBalancerOutput, error) {
	var output elbv2.CreateLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2CreateRuleResult struct {
	Result workflow.Future
}

func (r *Elbv2CreateRuleResult) Get(ctx workflow.Context) (*elbv2.CreateRuleOutput, error) {
	var output elbv2.CreateRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2CreateTargetGroupResult struct {
	Result workflow.Future
}

func (r *Elbv2CreateTargetGroupResult) Get(ctx workflow.Context) (*elbv2.CreateTargetGroupOutput, error) {
	var output elbv2.CreateTargetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DeleteListenerResult struct {
	Result workflow.Future
}

func (r *Elbv2DeleteListenerResult) Get(ctx workflow.Context) (*elbv2.DeleteListenerOutput, error) {
	var output elbv2.DeleteListenerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DeleteLoadBalancerResult struct {
	Result workflow.Future
}

func (r *Elbv2DeleteLoadBalancerResult) Get(ctx workflow.Context) (*elbv2.DeleteLoadBalancerOutput, error) {
	var output elbv2.DeleteLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DeleteRuleResult struct {
	Result workflow.Future
}

func (r *Elbv2DeleteRuleResult) Get(ctx workflow.Context) (*elbv2.DeleteRuleOutput, error) {
	var output elbv2.DeleteRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DeleteTargetGroupResult struct {
	Result workflow.Future
}

func (r *Elbv2DeleteTargetGroupResult) Get(ctx workflow.Context) (*elbv2.DeleteTargetGroupOutput, error) {
	var output elbv2.DeleteTargetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DeregisterTargetsResult struct {
	Result workflow.Future
}

func (r *Elbv2DeregisterTargetsResult) Get(ctx workflow.Context) (*elbv2.DeregisterTargetsOutput, error) {
	var output elbv2.DeregisterTargetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeAccountLimitsResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeAccountLimitsResult) Get(ctx workflow.Context) (*elbv2.DescribeAccountLimitsOutput, error) {
	var output elbv2.DescribeAccountLimitsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeListenerCertificatesResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeListenerCertificatesResult) Get(ctx workflow.Context) (*elbv2.DescribeListenerCertificatesOutput, error) {
	var output elbv2.DescribeListenerCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeListenersResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeListenersResult) Get(ctx workflow.Context) (*elbv2.DescribeListenersOutput, error) {
	var output elbv2.DescribeListenersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeLoadBalancerAttributesResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeLoadBalancerAttributesResult) Get(ctx workflow.Context) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	var output elbv2.DescribeLoadBalancerAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeLoadBalancersResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeLoadBalancersResult) Get(ctx workflow.Context) (*elbv2.DescribeLoadBalancersOutput, error) {
	var output elbv2.DescribeLoadBalancersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeRulesResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeRulesResult) Get(ctx workflow.Context) (*elbv2.DescribeRulesOutput, error) {
	var output elbv2.DescribeRulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeSSLPoliciesResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeSSLPoliciesResult) Get(ctx workflow.Context) (*elbv2.DescribeSSLPoliciesOutput, error) {
	var output elbv2.DescribeSSLPoliciesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeTagsResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeTagsResult) Get(ctx workflow.Context) (*elbv2.DescribeTagsOutput, error) {
	var output elbv2.DescribeTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeTargetGroupAttributesResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeTargetGroupAttributesResult) Get(ctx workflow.Context) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	var output elbv2.DescribeTargetGroupAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeTargetGroupsResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeTargetGroupsResult) Get(ctx workflow.Context) (*elbv2.DescribeTargetGroupsOutput, error) {
	var output elbv2.DescribeTargetGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2DescribeTargetHealthResult struct {
	Result workflow.Future
}

func (r *Elbv2DescribeTargetHealthResult) Get(ctx workflow.Context) (*elbv2.DescribeTargetHealthOutput, error) {
	var output elbv2.DescribeTargetHealthOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2ModifyListenerResult struct {
	Result workflow.Future
}

func (r *Elbv2ModifyListenerResult) Get(ctx workflow.Context) (*elbv2.ModifyListenerOutput, error) {
	var output elbv2.ModifyListenerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2ModifyLoadBalancerAttributesResult struct {
	Result workflow.Future
}

func (r *Elbv2ModifyLoadBalancerAttributesResult) Get(ctx workflow.Context) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	var output elbv2.ModifyLoadBalancerAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2ModifyRuleResult struct {
	Result workflow.Future
}

func (r *Elbv2ModifyRuleResult) Get(ctx workflow.Context) (*elbv2.ModifyRuleOutput, error) {
	var output elbv2.ModifyRuleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2ModifyTargetGroupResult struct {
	Result workflow.Future
}

func (r *Elbv2ModifyTargetGroupResult) Get(ctx workflow.Context) (*elbv2.ModifyTargetGroupOutput, error) {
	var output elbv2.ModifyTargetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2ModifyTargetGroupAttributesResult struct {
	Result workflow.Future
}

func (r *Elbv2ModifyTargetGroupAttributesResult) Get(ctx workflow.Context) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	var output elbv2.ModifyTargetGroupAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2RegisterTargetsResult struct {
	Result workflow.Future
}

func (r *Elbv2RegisterTargetsResult) Get(ctx workflow.Context) (*elbv2.RegisterTargetsOutput, error) {
	var output elbv2.RegisterTargetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2RemoveListenerCertificatesResult struct {
	Result workflow.Future
}

func (r *Elbv2RemoveListenerCertificatesResult) Get(ctx workflow.Context) (*elbv2.RemoveListenerCertificatesOutput, error) {
	var output elbv2.RemoveListenerCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2RemoveTagsResult struct {
	Result workflow.Future
}

func (r *Elbv2RemoveTagsResult) Get(ctx workflow.Context) (*elbv2.RemoveTagsOutput, error) {
	var output elbv2.RemoveTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2SetIpAddressTypeResult struct {
	Result workflow.Future
}

func (r *Elbv2SetIpAddressTypeResult) Get(ctx workflow.Context) (*elbv2.SetIpAddressTypeOutput, error) {
	var output elbv2.SetIpAddressTypeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2SetRulePrioritiesResult struct {
	Result workflow.Future
}

func (r *Elbv2SetRulePrioritiesResult) Get(ctx workflow.Context) (*elbv2.SetRulePrioritiesOutput, error) {
	var output elbv2.SetRulePrioritiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2SetSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *Elbv2SetSecurityGroupsResult) Get(ctx workflow.Context) (*elbv2.SetSecurityGroupsOutput, error) {
	var output elbv2.SetSecurityGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Elbv2SetSubnetsResult struct {
	Result workflow.Future
}

func (r *Elbv2SetSubnetsResult) Get(ctx workflow.Context) (*elbv2.SetSubnetsOutput, error) {
	var output elbv2.SetSubnetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}






func (a *ELBV2Stub) AddListenerCertificates(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error) {
	var output elbv2.AddListenerCertificatesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.AddListenerCertificates", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) AddListenerCertificatesAsync(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) *Elbv2AddListenerCertificatesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.AddListenerCertificates", input)
	return &Elbv2AddListenerCertificatesResult{Result: future}
}

func (a *ELBV2Stub) AddTags(ctx workflow.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
	var output elbv2.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) AddTagsAsync(ctx workflow.Context, input *elbv2.AddTagsInput) *Elbv2AddTagsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.AddTags", input)
	return &Elbv2AddTagsResult{Result: future}
}

func (a *ELBV2Stub) CreateListener(ctx workflow.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
	var output elbv2.CreateListenerOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.CreateListener", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) CreateListenerAsync(ctx workflow.Context, input *elbv2.CreateListenerInput) *Elbv2CreateListenerResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.CreateListener", input)
	return &Elbv2CreateListenerResult{Result: future}
}

func (a *ELBV2Stub) CreateLoadBalancer(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
	var output elbv2.CreateLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.CreateLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) CreateLoadBalancerAsync(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) *Elbv2CreateLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.CreateLoadBalancer", input)
	return &Elbv2CreateLoadBalancerResult{Result: future}
}

func (a *ELBV2Stub) CreateRule(ctx workflow.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error) {
	var output elbv2.CreateRuleOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.CreateRule", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) CreateRuleAsync(ctx workflow.Context, input *elbv2.CreateRuleInput) *Elbv2CreateRuleResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.CreateRule", input)
	return &Elbv2CreateRuleResult{Result: future}
}

func (a *ELBV2Stub) CreateTargetGroup(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
	var output elbv2.CreateTargetGroupOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.CreateTargetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) CreateTargetGroupAsync(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) *Elbv2CreateTargetGroupResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.CreateTargetGroup", input)
	return &Elbv2CreateTargetGroupResult{Result: future}
}

func (a *ELBV2Stub) DeleteListener(ctx workflow.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error) {
	var output elbv2.DeleteListenerOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DeleteListener", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DeleteListenerAsync(ctx workflow.Context, input *elbv2.DeleteListenerInput) *Elbv2DeleteListenerResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DeleteListener", input)
	return &Elbv2DeleteListenerResult{Result: future}
}

func (a *ELBV2Stub) DeleteLoadBalancer(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
	var output elbv2.DeleteLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DeleteLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DeleteLoadBalancerAsync(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) *Elbv2DeleteLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DeleteLoadBalancer", input)
	return &Elbv2DeleteLoadBalancerResult{Result: future}
}

func (a *ELBV2Stub) DeleteRule(ctx workflow.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error) {
	var output elbv2.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DeleteRule", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DeleteRuleAsync(ctx workflow.Context, input *elbv2.DeleteRuleInput) *Elbv2DeleteRuleResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DeleteRule", input)
	return &Elbv2DeleteRuleResult{Result: future}
}

func (a *ELBV2Stub) DeleteTargetGroup(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error) {
	var output elbv2.DeleteTargetGroupOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DeleteTargetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DeleteTargetGroupAsync(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) *Elbv2DeleteTargetGroupResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DeleteTargetGroup", input)
	return &Elbv2DeleteTargetGroupResult{Result: future}
}

func (a *ELBV2Stub) DeregisterTargets(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	var output elbv2.DeregisterTargetsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DeregisterTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DeregisterTargetsAsync(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) *Elbv2DeregisterTargetsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DeregisterTargets", input)
	return &Elbv2DeregisterTargetsResult{Result: future}
}

func (a *ELBV2Stub) DescribeAccountLimits(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error) {
	var output elbv2.DescribeAccountLimitsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeAccountLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeAccountLimitsAsync(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) *Elbv2DescribeAccountLimitsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeAccountLimits", input)
	return &Elbv2DescribeAccountLimitsResult{Result: future}
}

func (a *ELBV2Stub) DescribeListenerCertificates(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error) {
	var output elbv2.DescribeListenerCertificatesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeListenerCertificates", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeListenerCertificatesAsync(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) *Elbv2DescribeListenerCertificatesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeListenerCertificates", input)
	return &Elbv2DescribeListenerCertificatesResult{Result: future}
}

func (a *ELBV2Stub) DescribeListeners(ctx workflow.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
	var output elbv2.DescribeListenersOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeListenersAsync(ctx workflow.Context, input *elbv2.DescribeListenersInput) *Elbv2DescribeListenersResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeListeners", input)
	return &Elbv2DescribeListenersResult{Result: future}
}

func (a *ELBV2Stub) DescribeLoadBalancerAttributes(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	var output elbv2.DescribeLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) *Elbv2DescribeLoadBalancerAttributesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeLoadBalancerAttributes", input)
	return &Elbv2DescribeLoadBalancerAttributesResult{Result: future}
}

func (a *ELBV2Stub) DescribeLoadBalancers(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	var output elbv2.DescribeLoadBalancersOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeLoadBalancers", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeLoadBalancersAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *Elbv2DescribeLoadBalancersResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeLoadBalancers", input)
	return &Elbv2DescribeLoadBalancersResult{Result: future}
}

func (a *ELBV2Stub) DescribeRules(ctx workflow.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	var output elbv2.DescribeRulesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeRules", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeRulesAsync(ctx workflow.Context, input *elbv2.DescribeRulesInput) *Elbv2DescribeRulesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeRules", input)
	return &Elbv2DescribeRulesResult{Result: future}
}

func (a *ELBV2Stub) DescribeSSLPolicies(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error) {
	var output elbv2.DescribeSSLPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeSSLPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeSSLPoliciesAsync(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) *Elbv2DescribeSSLPoliciesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeSSLPolicies", input)
	return &Elbv2DescribeSSLPoliciesResult{Result: future}
}

func (a *ELBV2Stub) DescribeTags(ctx workflow.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	var output elbv2.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeTagsAsync(ctx workflow.Context, input *elbv2.DescribeTagsInput) *Elbv2DescribeTagsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTags", input)
	return &Elbv2DescribeTagsResult{Result: future}
}

func (a *ELBV2Stub) DescribeTargetGroupAttributes(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	var output elbv2.DescribeTargetGroupAttributesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTargetGroupAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) *Elbv2DescribeTargetGroupAttributesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTargetGroupAttributes", input)
	return &Elbv2DescribeTargetGroupAttributesResult{Result: future}
}

func (a *ELBV2Stub) DescribeTargetGroups(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	var output elbv2.DescribeTargetGroupsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTargetGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeTargetGroupsAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) *Elbv2DescribeTargetGroupsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTargetGroups", input)
	return &Elbv2DescribeTargetGroupsResult{Result: future}
}

func (a *ELBV2Stub) DescribeTargetHealth(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	var output elbv2.DescribeTargetHealthOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTargetHealth", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) DescribeTargetHealthAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *Elbv2DescribeTargetHealthResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.DescribeTargetHealth", input)
	return &Elbv2DescribeTargetHealthResult{Result: future}
}

func (a *ELBV2Stub) ModifyListener(ctx workflow.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error) {
	var output elbv2.ModifyListenerOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.ModifyListener", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) ModifyListenerAsync(ctx workflow.Context, input *elbv2.ModifyListenerInput) *Elbv2ModifyListenerResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.ModifyListener", input)
	return &Elbv2ModifyListenerResult{Result: future}
}

func (a *ELBV2Stub) ModifyLoadBalancerAttributes(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	var output elbv2.ModifyLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.ModifyLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) *Elbv2ModifyLoadBalancerAttributesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.ModifyLoadBalancerAttributes", input)
	return &Elbv2ModifyLoadBalancerAttributesResult{Result: future}
}

func (a *ELBV2Stub) ModifyRule(ctx workflow.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error) {
	var output elbv2.ModifyRuleOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.ModifyRule", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) ModifyRuleAsync(ctx workflow.Context, input *elbv2.ModifyRuleInput) *Elbv2ModifyRuleResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.ModifyRule", input)
	return &Elbv2ModifyRuleResult{Result: future}
}

func (a *ELBV2Stub) ModifyTargetGroup(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error) {
	var output elbv2.ModifyTargetGroupOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.ModifyTargetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) ModifyTargetGroupAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) *Elbv2ModifyTargetGroupResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.ModifyTargetGroup", input)
	return &Elbv2ModifyTargetGroupResult{Result: future}
}

func (a *ELBV2Stub) ModifyTargetGroupAttributes(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	var output elbv2.ModifyTargetGroupAttributesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.ModifyTargetGroupAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) ModifyTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) *Elbv2ModifyTargetGroupAttributesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.ModifyTargetGroupAttributes", input)
	return &Elbv2ModifyTargetGroupAttributesResult{Result: future}
}

func (a *ELBV2Stub) RegisterTargets(ctx workflow.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	var output elbv2.RegisterTargetsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.RegisterTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) RegisterTargetsAsync(ctx workflow.Context, input *elbv2.RegisterTargetsInput) *Elbv2RegisterTargetsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.RegisterTargets", input)
	return &Elbv2RegisterTargetsResult{Result: future}
}

func (a *ELBV2Stub) RemoveListenerCertificates(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error) {
	var output elbv2.RemoveListenerCertificatesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.RemoveListenerCertificates", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) RemoveListenerCertificatesAsync(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) *Elbv2RemoveListenerCertificatesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.RemoveListenerCertificates", input)
	return &Elbv2RemoveListenerCertificatesResult{Result: future}
}

func (a *ELBV2Stub) RemoveTags(ctx workflow.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error) {
	var output elbv2.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) RemoveTagsAsync(ctx workflow.Context, input *elbv2.RemoveTagsInput) *Elbv2RemoveTagsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.RemoveTags", input)
	return &Elbv2RemoveTagsResult{Result: future}
}

func (a *ELBV2Stub) SetIpAddressType(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error) {
	var output elbv2.SetIpAddressTypeOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.SetIpAddressType", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) SetIpAddressTypeAsync(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) *Elbv2SetIpAddressTypeResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.SetIpAddressType", input)
	return &Elbv2SetIpAddressTypeResult{Result: future}
}

func (a *ELBV2Stub) SetRulePriorities(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error) {
	var output elbv2.SetRulePrioritiesOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.SetRulePriorities", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) SetRulePrioritiesAsync(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) *Elbv2SetRulePrioritiesResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.SetRulePriorities", input)
	return &Elbv2SetRulePrioritiesResult{Result: future}
}

func (a *ELBV2Stub) SetSecurityGroups(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error) {
	var output elbv2.SetSecurityGroupsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.SetSecurityGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) SetSecurityGroupsAsync(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) *Elbv2SetSecurityGroupsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.SetSecurityGroups", input)
	return &Elbv2SetSecurityGroupsResult{Result: future}
}

func (a *ELBV2Stub) SetSubnets(ctx workflow.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error) {
	var output elbv2.SetSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "ELBV2.SetSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBV2Stub) SetSubnetsAsync(ctx workflow.Context, input *elbv2.SetSubnetsInput) *Elbv2SetSubnetsResult {
	future := workflow.ExecuteActivity(ctx, "ELBV2.SetSubnets", input)
	return &Elbv2SetSubnetsResult{Result: future}
}

func (a *ELBV2Stub) WaitUntilLoadBalancerAvailable(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilLoadBalancerAvailable", input).Get(ctx, nil)
}

func (a *ELBV2Stub) WaitUntilLoadBalancerAvailableAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilLoadBalancerAvailable", input)
}

func (a *ELBV2Stub) WaitUntilLoadBalancerExists(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilLoadBalancerExists", input).Get(ctx, nil)
}

func (a *ELBV2Stub) WaitUntilLoadBalancerExistsAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilLoadBalancerExists", input)
}

func (a *ELBV2Stub) WaitUntilLoadBalancersDeleted(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilLoadBalancersDeleted", input).Get(ctx, nil)
}

func (a *ELBV2Stub) WaitUntilLoadBalancersDeletedAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilLoadBalancersDeleted", input)
}

func (a *ELBV2Stub) WaitUntilTargetDeregistered(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilTargetDeregistered", input).Get(ctx, nil)
}

func (a *ELBV2Stub) WaitUntilTargetDeregisteredAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilTargetDeregistered", input)
}

func (a *ELBV2Stub) WaitUntilTargetInService(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilTargetInService", input).Get(ctx, nil)
}

func (a *ELBV2Stub) WaitUntilTargetInServiceAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELBV2.WaitUntilTargetInService", input)
}
