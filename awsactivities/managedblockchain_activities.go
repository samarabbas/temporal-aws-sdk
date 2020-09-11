
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/managedblockchain/managedblockchainiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ManagedBlockchainActivities struct {
    client managedblockchainiface.ManagedBlockchainAPI
}

func NewManagedBlockchainActivities(session *session.Session, config ...*aws.Config) *ManagedBlockchainActivities {
    client := managedblockchain.New(session, config...)
    return &ManagedBlockchainActivities{client: client}
}

func (a *ManagedBlockchainActivities) CreateMember(ctx context.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
    return a.client.CreateMemberWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) CreateNetwork(ctx context.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
    return a.client.CreateNetworkWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) CreateNode(ctx context.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
    return a.client.CreateNodeWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) CreateProposal(ctx context.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
    return a.client.CreateProposalWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) DeleteMember(ctx context.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
    return a.client.DeleteMemberWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) DeleteNode(ctx context.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
    return a.client.DeleteNodeWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) GetMember(ctx context.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
    return a.client.GetMemberWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) GetNetwork(ctx context.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
    return a.client.GetNetworkWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) GetNode(ctx context.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
    return a.client.GetNodeWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) GetProposal(ctx context.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
    return a.client.GetProposalWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) ListInvitations(ctx context.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
    return a.client.ListInvitationsWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) ListMembers(ctx context.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
    return a.client.ListMembersWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) ListNetworks(ctx context.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
    return a.client.ListNetworksWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) ListNodes(ctx context.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
    return a.client.ListNodesWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) ListProposalVotes(ctx context.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
    return a.client.ListProposalVotesWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) ListProposals(ctx context.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
    return a.client.ListProposalsWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) RejectInvitation(ctx context.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
    return a.client.RejectInvitationWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) UpdateMember(ctx context.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
    return a.client.UpdateMemberWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) UpdateNode(ctx context.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
    return a.client.UpdateNodeWithContext(ctx, input)
}

func (a *ManagedBlockchainActivities) VoteOnProposal(ctx context.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
    return a.client.VoteOnProposalWithContext(ctx, input)
}
