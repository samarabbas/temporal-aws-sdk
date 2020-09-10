package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ServiceDiscoveryActivities struct {
	client servicediscoveryiface.ServiceDiscoveryAPI
}

func NewServiceDiscoveryActivities(session *session.Session, config ...*aws.Config) *ServiceDiscoveryActivities {
	client := servicediscovery.New(session, config...)
	return &ServiceDiscoveryActivities{client: client}
}

func (a *ServiceDiscoveryActivities) CreateHttpNamespace(ctx context.Context, input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error) {
	return a.client.CreateHttpNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) CreatePrivateDnsNamespace(ctx context.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error) {
	return a.client.CreatePrivateDnsNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) CreatePublicDnsNamespace(ctx context.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error) {
	return a.client.CreatePublicDnsNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) CreateService(ctx context.Context, input *servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error) {
	return a.client.CreateServiceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DeleteNamespace(ctx context.Context, input *servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error) {
	return a.client.DeleteNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DeleteService(ctx context.Context, input *servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error) {
	return a.client.DeleteServiceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DeregisterInstance(ctx context.Context, input *servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error) {
	return a.client.DeregisterInstanceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DiscoverInstances(ctx context.Context, input *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error) {
	return a.client.DiscoverInstancesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetInstance(ctx context.Context, input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error) {
	return a.client.GetInstanceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetInstancesHealthStatus(ctx context.Context, input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
	return a.client.GetInstancesHealthStatusWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetNamespace(ctx context.Context, input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error) {
	return a.client.GetNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetOperation(ctx context.Context, input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error) {
	return a.client.GetOperationWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetService(ctx context.Context, input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error) {
	return a.client.GetServiceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListInstances(ctx context.Context, input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error) {
	return a.client.ListInstancesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListNamespaces(ctx context.Context, input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error) {
	return a.client.ListNamespacesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListOperations(ctx context.Context, input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error) {
	return a.client.ListOperationsWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListServices(ctx context.Context, input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error) {
	return a.client.ListServicesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListTagsForResource(ctx context.Context, input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) RegisterInstance(ctx context.Context, input *servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error) {
	return a.client.RegisterInstanceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) TagResource(ctx context.Context, input *servicediscovery.TagResourceInput) (*servicediscovery.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) UntagResource(ctx context.Context, input *servicediscovery.UntagResourceInput) (*servicediscovery.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) UpdateInstanceCustomHealthStatus(ctx context.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error) {
	return a.client.UpdateInstanceCustomHealthStatusWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) UpdateService(ctx context.Context, input *servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error) {
	return a.client.UpdateServiceWithContext(ctx, input)
}
