// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type GlobalAcceleratorActivities struct {
	client globalacceleratoriface.GlobalAcceleratorAPI
}

func NewGlobalAcceleratorActivities(session *session.Session, config ...*aws.Config) *GlobalAcceleratorActivities {
	client := globalaccelerator.New(session, config...)
	return &GlobalAcceleratorActivities{client: client}
}

func (a *GlobalAcceleratorActivities) AdvertiseByoipCidr(ctx context.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
	return a.client.AdvertiseByoipCidrWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) CreateAccelerator(ctx context.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error) {
	return a.client.CreateAcceleratorWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) CreateEndpointGroup(ctx context.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error) {
	return a.client.CreateEndpointGroupWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) CreateListener(ctx context.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error) {
	return a.client.CreateListenerWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DeleteAccelerator(ctx context.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error) {
	return a.client.DeleteAcceleratorWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DeleteEndpointGroup(ctx context.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
	return a.client.DeleteEndpointGroupWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DeleteListener(ctx context.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error) {
	return a.client.DeleteListenerWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DeprovisionByoipCidr(ctx context.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
	return a.client.DeprovisionByoipCidrWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DescribeAccelerator(ctx context.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	return a.client.DescribeAcceleratorWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DescribeAcceleratorAttributes(ctx context.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	return a.client.DescribeAcceleratorAttributesWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DescribeEndpointGroup(ctx context.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	return a.client.DescribeEndpointGroupWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) DescribeListener(ctx context.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error) {
	return a.client.DescribeListenerWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) ListAccelerators(ctx context.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error) {
	return a.client.ListAcceleratorsWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) ListByoipCidrs(ctx context.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error) {
	return a.client.ListByoipCidrsWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) ListEndpointGroups(ctx context.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	return a.client.ListEndpointGroupsWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) ListListeners(ctx context.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error) {
	return a.client.ListListenersWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) ListTagsForResource(ctx context.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) ProvisionByoipCidr(ctx context.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
	return a.client.ProvisionByoipCidrWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) TagResource(ctx context.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) UntagResource(ctx context.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) UpdateAccelerator(ctx context.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error) {
	return a.client.UpdateAcceleratorWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) UpdateAcceleratorAttributes(ctx context.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
	return a.client.UpdateAcceleratorAttributesWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) UpdateEndpointGroup(ctx context.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
	return a.client.UpdateEndpointGroupWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) UpdateListener(ctx context.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error) {
	return a.client.UpdateListenerWithContext(ctx, input)
}

func (a *GlobalAcceleratorActivities) WithdrawByoipCidr(ctx context.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
	return a.client.WithdrawByoipCidrWithContext(ctx, input)
}
