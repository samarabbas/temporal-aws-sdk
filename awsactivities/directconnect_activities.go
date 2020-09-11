
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directconnect/directconnectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DirectConnectActivities struct {
    client directconnectiface.DirectConnectAPI
}

func NewDirectConnectActivities(session *session.Session, config ...*aws.Config) *DirectConnectActivities {
    client := directconnect.New(session, config...)
    return &DirectConnectActivities{client: client}
}

func (a *DirectConnectActivities) AcceptDirectConnectGatewayAssociationProposal(ctx context.Context, input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error) {
    return a.client.AcceptDirectConnectGatewayAssociationProposalWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocateConnectionOnInterconnect(ctx context.Context, input *directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error) {
    return a.client.AllocateConnectionOnInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocateHostedConnection(ctx context.Context, input *directconnect.AllocateHostedConnectionInput) (*directconnect.Connection, error) {
    return a.client.AllocateHostedConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocatePrivateVirtualInterface(ctx context.Context, input *directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.AllocatePrivateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocatePublicVirtualInterface(ctx context.Context, input *directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.AllocatePublicVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocateTransitVirtualInterface(ctx context.Context, input *directconnect.AllocateTransitVirtualInterfaceInput) (*directconnect.AllocateTransitVirtualInterfaceOutput, error) {
    return a.client.AllocateTransitVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) AssociateConnectionWithLag(ctx context.Context, input *directconnect.AssociateConnectionWithLagInput) (*directconnect.Connection, error) {
    return a.client.AssociateConnectionWithLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) AssociateHostedConnection(ctx context.Context, input *directconnect.AssociateHostedConnectionInput) (*directconnect.Connection, error) {
    return a.client.AssociateHostedConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) AssociateVirtualInterface(ctx context.Context, input *directconnect.AssociateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.AssociateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmConnection(ctx context.Context, input *directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error) {
    return a.client.ConfirmConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmPrivateVirtualInterface(ctx context.Context, input *directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error) {
    return a.client.ConfirmPrivateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmPublicVirtualInterface(ctx context.Context, input *directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error) {
    return a.client.ConfirmPublicVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmTransitVirtualInterface(ctx context.Context, input *directconnect.ConfirmTransitVirtualInterfaceInput) (*directconnect.ConfirmTransitVirtualInterfaceOutput, error) {
    return a.client.ConfirmTransitVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateBGPPeer(ctx context.Context, input *directconnect.CreateBGPPeerInput) (*directconnect.CreateBGPPeerOutput, error) {
    return a.client.CreateBGPPeerWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateConnection(ctx context.Context, input *directconnect.CreateConnectionInput) (*directconnect.Connection, error) {
    return a.client.CreateConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateDirectConnectGateway(ctx context.Context, input *directconnect.CreateDirectConnectGatewayInput) (*directconnect.CreateDirectConnectGatewayOutput, error) {
    return a.client.CreateDirectConnectGatewayWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateDirectConnectGatewayAssociation(ctx context.Context, input *directconnect.CreateDirectConnectGatewayAssociationInput) (*directconnect.CreateDirectConnectGatewayAssociationOutput, error) {
    return a.client.CreateDirectConnectGatewayAssociationWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateDirectConnectGatewayAssociationProposal(ctx context.Context, input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) (*directconnect.CreateDirectConnectGatewayAssociationProposalOutput, error) {
    return a.client.CreateDirectConnectGatewayAssociationProposalWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateInterconnect(ctx context.Context, input *directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error) {
    return a.client.CreateInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateLag(ctx context.Context, input *directconnect.CreateLagInput) (*directconnect.Lag, error) {
    return a.client.CreateLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreatePrivateVirtualInterface(ctx context.Context, input *directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.CreatePrivateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreatePublicVirtualInterface(ctx context.Context, input *directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.CreatePublicVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateTransitVirtualInterface(ctx context.Context, input *directconnect.CreateTransitVirtualInterfaceInput) (*directconnect.CreateTransitVirtualInterfaceOutput, error) {
    return a.client.CreateTransitVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteBGPPeer(ctx context.Context, input *directconnect.DeleteBGPPeerInput) (*directconnect.DeleteBGPPeerOutput, error) {
    return a.client.DeleteBGPPeerWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteConnection(ctx context.Context, input *directconnect.DeleteConnectionInput) (*directconnect.Connection, error) {
    return a.client.DeleteConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGateway(ctx context.Context, input *directconnect.DeleteDirectConnectGatewayInput) (*directconnect.DeleteDirectConnectGatewayOutput, error) {
    return a.client.DeleteDirectConnectGatewayWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGatewayAssociation(ctx context.Context, input *directconnect.DeleteDirectConnectGatewayAssociationInput) (*directconnect.DeleteDirectConnectGatewayAssociationOutput, error) {
    return a.client.DeleteDirectConnectGatewayAssociationWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGatewayAssociationProposal(ctx context.Context, input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) (*directconnect.DeleteDirectConnectGatewayAssociationProposalOutput, error) {
    return a.client.DeleteDirectConnectGatewayAssociationProposalWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteInterconnect(ctx context.Context, input *directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error) {
    return a.client.DeleteInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteLag(ctx context.Context, input *directconnect.DeleteLagInput) (*directconnect.Lag, error) {
    return a.client.DeleteLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteVirtualInterface(ctx context.Context, input *directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error) {
    return a.client.DeleteVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeConnectionLoa(ctx context.Context, input *directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error) {
    return a.client.DescribeConnectionLoaWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeConnections(ctx context.Context, input *directconnect.DescribeConnectionsInput) (*directconnect.Connections, error) {
    return a.client.DescribeConnectionsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeConnectionsOnInterconnect(ctx context.Context, input *directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error) {
    return a.client.DescribeConnectionsOnInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAssociationProposals(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
    return a.client.DescribeDirectConnectGatewayAssociationProposalsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAssociations(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
    return a.client.DescribeDirectConnectGatewayAssociationsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAttachments(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
    return a.client.DescribeDirectConnectGatewayAttachmentsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGateways(ctx context.Context, input *directconnect.DescribeDirectConnectGatewaysInput) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
    return a.client.DescribeDirectConnectGatewaysWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeHostedConnections(ctx context.Context, input *directconnect.DescribeHostedConnectionsInput) (*directconnect.Connections, error) {
    return a.client.DescribeHostedConnectionsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeInterconnectLoa(ctx context.Context, input *directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error) {
    return a.client.DescribeInterconnectLoaWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeInterconnects(ctx context.Context, input *directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error) {
    return a.client.DescribeInterconnectsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeLags(ctx context.Context, input *directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error) {
    return a.client.DescribeLagsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeLoa(ctx context.Context, input *directconnect.DescribeLoaInput) (*directconnect.Loa, error) {
    return a.client.DescribeLoaWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeLocations(ctx context.Context, input *directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error) {
    return a.client.DescribeLocationsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeTags(ctx context.Context, input *directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error) {
    return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeVirtualGateways(ctx context.Context, input *directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error) {
    return a.client.DescribeVirtualGatewaysWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeVirtualInterfaces(ctx context.Context, input *directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error) {
    return a.client.DescribeVirtualInterfacesWithContext(ctx, input)
}

func (a *DirectConnectActivities) DisassociateConnectionFromLag(ctx context.Context, input *directconnect.DisassociateConnectionFromLagInput) (*directconnect.Connection, error) {
    return a.client.DisassociateConnectionFromLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) ListVirtualInterfaceTestHistory(ctx context.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
    return a.client.ListVirtualInterfaceTestHistoryWithContext(ctx, input)
}

func (a *DirectConnectActivities) StartBgpFailoverTest(ctx context.Context, input *directconnect.StartBgpFailoverTestInput) (*directconnect.StartBgpFailoverTestOutput, error) {
    return a.client.StartBgpFailoverTestWithContext(ctx, input)
}

func (a *DirectConnectActivities) StopBgpFailoverTest(ctx context.Context, input *directconnect.StopBgpFailoverTestInput) (*directconnect.StopBgpFailoverTestOutput, error) {
    return a.client.StopBgpFailoverTestWithContext(ctx, input)
}

func (a *DirectConnectActivities) TagResource(ctx context.Context, input *directconnect.TagResourceInput) (*directconnect.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *DirectConnectActivities) UntagResource(ctx context.Context, input *directconnect.UntagResourceInput) (*directconnect.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DirectConnectActivities) UpdateDirectConnectGatewayAssociation(ctx context.Context, input *directconnect.UpdateDirectConnectGatewayAssociationInput) (*directconnect.UpdateDirectConnectGatewayAssociationOutput, error) {
    return a.client.UpdateDirectConnectGatewayAssociationWithContext(ctx, input)
}

func (a *DirectConnectActivities) UpdateLag(ctx context.Context, input *directconnect.UpdateLagInput) (*directconnect.Lag, error) {
    return a.client.UpdateLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) UpdateVirtualInterfaceAttributes(ctx context.Context, input *directconnect.UpdateVirtualInterfaceAttributesInput) (*directconnect.UpdateVirtualInterfaceAttributesOutput, error) {
    return a.client.UpdateVirtualInterfaceAttributesWithContext(ctx, input)
}
