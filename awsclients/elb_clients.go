// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elb"
	"go.temporal.io/sdk/workflow"
)

type ELBClient interface {
	AddTags(ctx workflow.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *elb.AddTagsInput) *ElbAddTagsResult

	ApplySecurityGroupsToLoadBalancer(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)
	ApplySecurityGroupsToLoadBalancerAsync(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) *ElbApplySecurityGroupsToLoadBalancerResult

	AttachLoadBalancerToSubnets(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error)
	AttachLoadBalancerToSubnetsAsync(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) *ElbAttachLoadBalancerToSubnetsResult

	ConfigureHealthCheck(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error)
	ConfigureHealthCheckAsync(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) *ElbConfigureHealthCheckResult

	CreateAppCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error)
	CreateAppCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) *ElbCreateAppCookieStickinessPolicyResult

	CreateLBCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error)
	CreateLBCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) *ElbCreateLBCookieStickinessPolicyResult

	CreateLoadBalancer(ctx workflow.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *elb.CreateLoadBalancerInput) *ElbCreateLoadBalancerResult

	CreateLoadBalancerListeners(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error)
	CreateLoadBalancerListenersAsync(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) *ElbCreateLoadBalancerListenersResult

	CreateLoadBalancerPolicy(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error)
	CreateLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) *ElbCreateLoadBalancerPolicyResult

	DeleteLoadBalancer(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) *ElbDeleteLoadBalancerResult

	DeleteLoadBalancerListeners(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error)
	DeleteLoadBalancerListenersAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) *ElbDeleteLoadBalancerListenersResult

	DeleteLoadBalancerPolicy(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error)
	DeleteLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) *ElbDeleteLoadBalancerPolicyResult

	DeregisterInstancesFromLoadBalancer(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)
	DeregisterInstancesFromLoadBalancerAsync(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) *ElbDeregisterInstancesFromLoadBalancerResult

	DescribeAccountLimits(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) *ElbDescribeAccountLimitsResult

	DescribeInstanceHealth(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error)
	DescribeInstanceHealthAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *ElbDescribeInstanceHealthResult

	DescribeLoadBalancerAttributes(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) *ElbDescribeLoadBalancerAttributesResult

	DescribeLoadBalancerPolicies(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error)
	DescribeLoadBalancerPoliciesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) *ElbDescribeLoadBalancerPoliciesResult

	DescribeLoadBalancerPolicyTypes(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)
	DescribeLoadBalancerPolicyTypesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) *ElbDescribeLoadBalancerPolicyTypesResult

	DescribeLoadBalancers(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersAsync(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) *ElbDescribeLoadBalancersResult

	DescribeTags(ctx workflow.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *elb.DescribeTagsInput) *ElbDescribeTagsResult

	DetachLoadBalancerFromSubnets(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error)
	DetachLoadBalancerFromSubnetsAsync(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) *ElbDetachLoadBalancerFromSubnetsResult

	DisableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)
	DisableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) *ElbDisableAvailabilityZonesForLoadBalancerResult

	EnableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)
	EnableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) *ElbEnableAvailabilityZonesForLoadBalancerResult

	ModifyLoadBalancerAttributes(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) *ElbModifyLoadBalancerAttributesResult

	RegisterInstancesWithLoadBalancer(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	RegisterInstancesWithLoadBalancerAsync(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) *ElbRegisterInstancesWithLoadBalancerResult

	RemoveTags(ctx workflow.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *elb.RemoveTagsInput) *ElbRemoveTagsResult

	SetLoadBalancerListenerSSLCertificate(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)
	SetLoadBalancerListenerSSLCertificateAsync(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) *ElbSetLoadBalancerListenerSSLCertificateResult

	SetLoadBalancerPoliciesForBackendServer(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)
	SetLoadBalancerPoliciesForBackendServerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) *ElbSetLoadBalancerPoliciesForBackendServerResult

	SetLoadBalancerPoliciesOfListener(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
	SetLoadBalancerPoliciesOfListenerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) *ElbSetLoadBalancerPoliciesOfListenerResult

	WaitUntilAnyInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceDeregistered(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error}

type ELBStub struct{}

func NewELBStub() ELBClient {
	return &ELBStub{}
}

type ElbAddTagsResult struct {
	Result workflow.Future
}

func (r *ElbAddTagsResult) Get(ctx workflow.Context) (*elb.AddTagsOutput, error) {
	var output elb.AddTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbApplySecurityGroupsToLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbApplySecurityGroupsToLoadBalancerResult) Get(ctx workflow.Context) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	var output elb.ApplySecurityGroupsToLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbAttachLoadBalancerToSubnetsResult struct {
	Result workflow.Future
}

func (r *ElbAttachLoadBalancerToSubnetsResult) Get(ctx workflow.Context) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	var output elb.AttachLoadBalancerToSubnetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbConfigureHealthCheckResult struct {
	Result workflow.Future
}

func (r *ElbConfigureHealthCheckResult) Get(ctx workflow.Context) (*elb.ConfigureHealthCheckOutput, error) {
	var output elb.ConfigureHealthCheckOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbCreateAppCookieStickinessPolicyResult struct {
	Result workflow.Future
}

func (r *ElbCreateAppCookieStickinessPolicyResult) Get(ctx workflow.Context) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	var output elb.CreateAppCookieStickinessPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbCreateLBCookieStickinessPolicyResult struct {
	Result workflow.Future
}

func (r *ElbCreateLBCookieStickinessPolicyResult) Get(ctx workflow.Context) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	var output elb.CreateLBCookieStickinessPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbCreateLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbCreateLoadBalancerResult) Get(ctx workflow.Context) (*elb.CreateLoadBalancerOutput, error) {
	var output elb.CreateLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbCreateLoadBalancerListenersResult struct {
	Result workflow.Future
}

func (r *ElbCreateLoadBalancerListenersResult) Get(ctx workflow.Context) (*elb.CreateLoadBalancerListenersOutput, error) {
	var output elb.CreateLoadBalancerListenersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbCreateLoadBalancerPolicyResult struct {
	Result workflow.Future
}

func (r *ElbCreateLoadBalancerPolicyResult) Get(ctx workflow.Context) (*elb.CreateLoadBalancerPolicyOutput, error) {
	var output elb.CreateLoadBalancerPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDeleteLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbDeleteLoadBalancerResult) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerOutput, error) {
	var output elb.DeleteLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDeleteLoadBalancerListenersResult struct {
	Result workflow.Future
}

func (r *ElbDeleteLoadBalancerListenersResult) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerListenersOutput, error) {
	var output elb.DeleteLoadBalancerListenersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDeleteLoadBalancerPolicyResult struct {
	Result workflow.Future
}

func (r *ElbDeleteLoadBalancerPolicyResult) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	var output elb.DeleteLoadBalancerPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDeregisterInstancesFromLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbDeregisterInstancesFromLoadBalancerResult) Get(ctx workflow.Context) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	var output elb.DeregisterInstancesFromLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeAccountLimitsResult struct {
	Result workflow.Future
}

func (r *ElbDescribeAccountLimitsResult) Get(ctx workflow.Context) (*elb.DescribeAccountLimitsOutput, error) {
	var output elb.DescribeAccountLimitsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeInstanceHealthResult struct {
	Result workflow.Future
}

func (r *ElbDescribeInstanceHealthResult) Get(ctx workflow.Context) (*elb.DescribeInstanceHealthOutput, error) {
	var output elb.DescribeInstanceHealthOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeLoadBalancerAttributesResult struct {
	Result workflow.Future
}

func (r *ElbDescribeLoadBalancerAttributesResult) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	var output elb.DescribeLoadBalancerAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeLoadBalancerPoliciesResult struct {
	Result workflow.Future
}

func (r *ElbDescribeLoadBalancerPoliciesResult) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	var output elb.DescribeLoadBalancerPoliciesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeLoadBalancerPolicyTypesResult struct {
	Result workflow.Future
}

func (r *ElbDescribeLoadBalancerPolicyTypesResult) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	var output elb.DescribeLoadBalancerPolicyTypesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeLoadBalancersResult struct {
	Result workflow.Future
}

func (r *ElbDescribeLoadBalancersResult) Get(ctx workflow.Context) (*elb.DescribeLoadBalancersOutput, error) {
	var output elb.DescribeLoadBalancersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDescribeTagsResult struct {
	Result workflow.Future
}

func (r *ElbDescribeTagsResult) Get(ctx workflow.Context) (*elb.DescribeTagsOutput, error) {
	var output elb.DescribeTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDetachLoadBalancerFromSubnetsResult struct {
	Result workflow.Future
}

func (r *ElbDetachLoadBalancerFromSubnetsResult) Get(ctx workflow.Context) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	var output elb.DetachLoadBalancerFromSubnetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbDisableAvailabilityZonesForLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbDisableAvailabilityZonesForLoadBalancerResult) Get(ctx workflow.Context) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.DisableAvailabilityZonesForLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbEnableAvailabilityZonesForLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbEnableAvailabilityZonesForLoadBalancerResult) Get(ctx workflow.Context) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.EnableAvailabilityZonesForLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbModifyLoadBalancerAttributesResult struct {
	Result workflow.Future
}

func (r *ElbModifyLoadBalancerAttributesResult) Get(ctx workflow.Context) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	var output elb.ModifyLoadBalancerAttributesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbRegisterInstancesWithLoadBalancerResult struct {
	Result workflow.Future
}

func (r *ElbRegisterInstancesWithLoadBalancerResult) Get(ctx workflow.Context) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	var output elb.RegisterInstancesWithLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbRemoveTagsResult struct {
	Result workflow.Future
}

func (r *ElbRemoveTagsResult) Get(ctx workflow.Context) (*elb.RemoveTagsOutput, error) {
	var output elb.RemoveTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbSetLoadBalancerListenerSSLCertificateResult struct {
	Result workflow.Future
}

func (r *ElbSetLoadBalancerListenerSSLCertificateResult) Get(ctx workflow.Context) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	var output elb.SetLoadBalancerListenerSSLCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbSetLoadBalancerPoliciesForBackendServerResult struct {
	Result workflow.Future
}

func (r *ElbSetLoadBalancerPoliciesForBackendServerResult) Get(ctx workflow.Context) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	var output elb.SetLoadBalancerPoliciesForBackendServerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ElbSetLoadBalancerPoliciesOfListenerResult struct {
	Result workflow.Future
}

func (r *ElbSetLoadBalancerPoliciesOfListenerResult) Get(ctx workflow.Context) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	var output elb.SetLoadBalancerPoliciesOfListenerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}




func (a *ELBStub) AddTags(ctx workflow.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
	var output elb.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "ELB.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) AddTagsAsync(ctx workflow.Context, input *elb.AddTagsInput) *ElbAddTagsResult {
	future := workflow.ExecuteActivity(ctx, "ELB.AddTags", input)
	return &ElbAddTagsResult{Result: future}
}

func (a *ELBStub) ApplySecurityGroupsToLoadBalancer(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	var output elb.ApplySecurityGroupsToLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.ApplySecurityGroupsToLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) ApplySecurityGroupsToLoadBalancerAsync(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) *ElbApplySecurityGroupsToLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.ApplySecurityGroupsToLoadBalancer", input)
	return &ElbApplySecurityGroupsToLoadBalancerResult{Result: future}
}

func (a *ELBStub) AttachLoadBalancerToSubnets(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	var output elb.AttachLoadBalancerToSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "ELB.AttachLoadBalancerToSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) AttachLoadBalancerToSubnetsAsync(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) *ElbAttachLoadBalancerToSubnetsResult {
	future := workflow.ExecuteActivity(ctx, "ELB.AttachLoadBalancerToSubnets", input)
	return &ElbAttachLoadBalancerToSubnetsResult{Result: future}
}

func (a *ELBStub) ConfigureHealthCheck(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
	var output elb.ConfigureHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "ELB.ConfigureHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) ConfigureHealthCheckAsync(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) *ElbConfigureHealthCheckResult {
	future := workflow.ExecuteActivity(ctx, "ELB.ConfigureHealthCheck", input)
	return &ElbConfigureHealthCheckResult{Result: future}
}

func (a *ELBStub) CreateAppCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	var output elb.CreateAppCookieStickinessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ELB.CreateAppCookieStickinessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) CreateAppCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) *ElbCreateAppCookieStickinessPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ELB.CreateAppCookieStickinessPolicy", input)
	return &ElbCreateAppCookieStickinessPolicyResult{Result: future}
}

func (a *ELBStub) CreateLBCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	var output elb.CreateLBCookieStickinessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ELB.CreateLBCookieStickinessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) CreateLBCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) *ElbCreateLBCookieStickinessPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ELB.CreateLBCookieStickinessPolicy", input)
	return &ElbCreateLBCookieStickinessPolicyResult{Result: future}
}

func (a *ELBStub) CreateLoadBalancer(ctx workflow.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	var output elb.CreateLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.CreateLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) CreateLoadBalancerAsync(ctx workflow.Context, input *elb.CreateLoadBalancerInput) *ElbCreateLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.CreateLoadBalancer", input)
	return &ElbCreateLoadBalancerResult{Result: future}
}

func (a *ELBStub) CreateLoadBalancerListeners(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	var output elb.CreateLoadBalancerListenersOutput
	err := workflow.ExecuteActivity(ctx, "ELB.CreateLoadBalancerListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) CreateLoadBalancerListenersAsync(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) *ElbCreateLoadBalancerListenersResult {
	future := workflow.ExecuteActivity(ctx, "ELB.CreateLoadBalancerListeners", input)
	return &ElbCreateLoadBalancerListenersResult{Result: future}
}

func (a *ELBStub) CreateLoadBalancerPolicy(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
	var output elb.CreateLoadBalancerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ELB.CreateLoadBalancerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) CreateLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) *ElbCreateLoadBalancerPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ELB.CreateLoadBalancerPolicy", input)
	return &ElbCreateLoadBalancerPolicyResult{Result: future}
}

func (a *ELBStub) DeleteLoadBalancer(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	var output elb.DeleteLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DeleteLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DeleteLoadBalancerAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) *ElbDeleteLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DeleteLoadBalancer", input)
	return &ElbDeleteLoadBalancerResult{Result: future}
}

func (a *ELBStub) DeleteLoadBalancerListeners(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	var output elb.DeleteLoadBalancerListenersOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DeleteLoadBalancerListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DeleteLoadBalancerListenersAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) *ElbDeleteLoadBalancerListenersResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DeleteLoadBalancerListeners", input)
	return &ElbDeleteLoadBalancerListenersResult{Result: future}
}

func (a *ELBStub) DeleteLoadBalancerPolicy(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	var output elb.DeleteLoadBalancerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DeleteLoadBalancerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DeleteLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) *ElbDeleteLoadBalancerPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DeleteLoadBalancerPolicy", input)
	return &ElbDeleteLoadBalancerPolicyResult{Result: future}
}

func (a *ELBStub) DeregisterInstancesFromLoadBalancer(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	var output elb.DeregisterInstancesFromLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DeregisterInstancesFromLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DeregisterInstancesFromLoadBalancerAsync(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) *ElbDeregisterInstancesFromLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DeregisterInstancesFromLoadBalancer", input)
	return &ElbDeregisterInstancesFromLoadBalancerResult{Result: future}
}

func (a *ELBStub) DescribeAccountLimits(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
	var output elb.DescribeAccountLimitsOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeAccountLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeAccountLimitsAsync(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) *ElbDescribeAccountLimitsResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeAccountLimits", input)
	return &ElbDescribeAccountLimitsResult{Result: future}
}

func (a *ELBStub) DescribeInstanceHealth(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
	var output elb.DescribeInstanceHealthOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeInstanceHealth", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeInstanceHealthAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *ElbDescribeInstanceHealthResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeInstanceHealth", input)
	return &ElbDescribeInstanceHealthResult{Result: future}
}

func (a *ELBStub) DescribeLoadBalancerAttributes(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	var output elb.DescribeLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) *ElbDescribeLoadBalancerAttributesResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancerAttributes", input)
	return &ElbDescribeLoadBalancerAttributesResult{Result: future}
}

func (a *ELBStub) DescribeLoadBalancerPolicies(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	var output elb.DescribeLoadBalancerPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancerPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeLoadBalancerPoliciesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) *ElbDescribeLoadBalancerPoliciesResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancerPolicies", input)
	return &ElbDescribeLoadBalancerPoliciesResult{Result: future}
}

func (a *ELBStub) DescribeLoadBalancerPolicyTypes(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	var output elb.DescribeLoadBalancerPolicyTypesOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancerPolicyTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeLoadBalancerPolicyTypesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) *ElbDescribeLoadBalancerPolicyTypesResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancerPolicyTypes", input)
	return &ElbDescribeLoadBalancerPolicyTypesResult{Result: future}
}

func (a *ELBStub) DescribeLoadBalancers(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	var output elb.DescribeLoadBalancersOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancers", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeLoadBalancersAsync(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) *ElbDescribeLoadBalancersResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeLoadBalancers", input)
	return &ElbDescribeLoadBalancersResult{Result: future}
}

func (a *ELBStub) DescribeTags(ctx workflow.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	var output elb.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DescribeTagsAsync(ctx workflow.Context, input *elb.DescribeTagsInput) *ElbDescribeTagsResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DescribeTags", input)
	return &ElbDescribeTagsResult{Result: future}
}

func (a *ELBStub) DetachLoadBalancerFromSubnets(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	var output elb.DetachLoadBalancerFromSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DetachLoadBalancerFromSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DetachLoadBalancerFromSubnetsAsync(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) *ElbDetachLoadBalancerFromSubnetsResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DetachLoadBalancerFromSubnets", input)
	return &ElbDetachLoadBalancerFromSubnetsResult{Result: future}
}

func (a *ELBStub) DisableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.DisableAvailabilityZonesForLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.DisableAvailabilityZonesForLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) DisableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) *ElbDisableAvailabilityZonesForLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.DisableAvailabilityZonesForLoadBalancer", input)
	return &ElbDisableAvailabilityZonesForLoadBalancerResult{Result: future}
}

func (a *ELBStub) EnableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.EnableAvailabilityZonesForLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.EnableAvailabilityZonesForLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) EnableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) *ElbEnableAvailabilityZonesForLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.EnableAvailabilityZonesForLoadBalancer", input)
	return &ElbEnableAvailabilityZonesForLoadBalancerResult{Result: future}
}

func (a *ELBStub) ModifyLoadBalancerAttributes(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	var output elb.ModifyLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "ELB.ModifyLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) *ElbModifyLoadBalancerAttributesResult {
	future := workflow.ExecuteActivity(ctx, "ELB.ModifyLoadBalancerAttributes", input)
	return &ElbModifyLoadBalancerAttributesResult{Result: future}
}

func (a *ELBStub) RegisterInstancesWithLoadBalancer(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	var output elb.RegisterInstancesWithLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.RegisterInstancesWithLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) RegisterInstancesWithLoadBalancerAsync(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) *ElbRegisterInstancesWithLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.RegisterInstancesWithLoadBalancer", input)
	return &ElbRegisterInstancesWithLoadBalancerResult{Result: future}
}

func (a *ELBStub) RemoveTags(ctx workflow.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error) {
	var output elb.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "ELB.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) RemoveTagsAsync(ctx workflow.Context, input *elb.RemoveTagsInput) *ElbRemoveTagsResult {
	future := workflow.ExecuteActivity(ctx, "ELB.RemoveTags", input)
	return &ElbRemoveTagsResult{Result: future}
}

func (a *ELBStub) SetLoadBalancerListenerSSLCertificate(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	var output elb.SetLoadBalancerListenerSSLCertificateOutput
	err := workflow.ExecuteActivity(ctx, "ELB.SetLoadBalancerListenerSSLCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) SetLoadBalancerListenerSSLCertificateAsync(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) *ElbSetLoadBalancerListenerSSLCertificateResult {
	future := workflow.ExecuteActivity(ctx, "ELB.SetLoadBalancerListenerSSLCertificate", input)
	return &ElbSetLoadBalancerListenerSSLCertificateResult{Result: future}
}

func (a *ELBStub) SetLoadBalancerPoliciesForBackendServer(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	var output elb.SetLoadBalancerPoliciesForBackendServerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.SetLoadBalancerPoliciesForBackendServer", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) SetLoadBalancerPoliciesForBackendServerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) *ElbSetLoadBalancerPoliciesForBackendServerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.SetLoadBalancerPoliciesForBackendServer", input)
	return &ElbSetLoadBalancerPoliciesForBackendServerResult{Result: future}
}

func (a *ELBStub) SetLoadBalancerPoliciesOfListener(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	var output elb.SetLoadBalancerPoliciesOfListenerOutput
	err := workflow.ExecuteActivity(ctx, "ELB.SetLoadBalancerPoliciesOfListener", input).Get(ctx, &output)
	return &output, err
}

func (a *ELBStub) SetLoadBalancerPoliciesOfListenerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) *ElbSetLoadBalancerPoliciesOfListenerResult {
	future := workflow.ExecuteActivity(ctx, "ELB.SetLoadBalancerPoliciesOfListener", input)
	return &ElbSetLoadBalancerPoliciesOfListenerResult{Result: future}
}

func (a *ELBStub) WaitUntilAnyInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "ELB.WaitUntilAnyInstanceInService", input).Get(ctx, nil)
}

func (a *ELBStub) WaitUntilAnyInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELB.WaitUntilAnyInstanceInService", input)
}

func (a *ELBStub) WaitUntilInstanceDeregistered(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "ELB.WaitUntilInstanceDeregistered", input).Get(ctx, nil)
}

func (a *ELBStub) WaitUntilInstanceDeregisteredAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELB.WaitUntilInstanceDeregistered", input)
}

func (a *ELBStub) WaitUntilInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "ELB.WaitUntilInstanceInService", input).Get(ctx, nil)
}

func (a *ELBStub) WaitUntilInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ELB.WaitUntilInstanceInService", input)
}
