package awsclients

import (
	"github.com/aws/aws-sdk-go/service/guardduty"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type GuardDutyClient interface {
    AcceptInvitation(ctx workflow.Context, input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error)
    AcceptInvitationAsync(ctx workflow.Context, input *guardduty.AcceptInvitationInput) *GuarddutyAcceptInvitationResult

    ArchiveFindings(ctx workflow.Context, input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error)
    ArchiveFindingsAsync(ctx workflow.Context, input *guardduty.ArchiveFindingsInput) *GuarddutyArchiveFindingsResult

    CreateDetector(ctx workflow.Context, input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error)
    CreateDetectorAsync(ctx workflow.Context, input *guardduty.CreateDetectorInput) *GuarddutyCreateDetectorResult

    CreateFilter(ctx workflow.Context, input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error)
    CreateFilterAsync(ctx workflow.Context, input *guardduty.CreateFilterInput) *GuarddutyCreateFilterResult

    CreateIPSet(ctx workflow.Context, input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error)
    CreateIPSetAsync(ctx workflow.Context, input *guardduty.CreateIPSetInput) *GuarddutyCreateIPSetResult

    CreateMembers(ctx workflow.Context, input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error)
    CreateMembersAsync(ctx workflow.Context, input *guardduty.CreateMembersInput) *GuarddutyCreateMembersResult

    CreatePublishingDestination(ctx workflow.Context, input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error)
    CreatePublishingDestinationAsync(ctx workflow.Context, input *guardduty.CreatePublishingDestinationInput) *GuarddutyCreatePublishingDestinationResult

    CreateSampleFindings(ctx workflow.Context, input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error)
    CreateSampleFindingsAsync(ctx workflow.Context, input *guardduty.CreateSampleFindingsInput) *GuarddutyCreateSampleFindingsResult

    CreateThreatIntelSet(ctx workflow.Context, input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error)
    CreateThreatIntelSetAsync(ctx workflow.Context, input *guardduty.CreateThreatIntelSetInput) *GuarddutyCreateThreatIntelSetResult

    DeclineInvitations(ctx workflow.Context, input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error)
    DeclineInvitationsAsync(ctx workflow.Context, input *guardduty.DeclineInvitationsInput) *GuarddutyDeclineInvitationsResult

    DeleteDetector(ctx workflow.Context, input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error)
    DeleteDetectorAsync(ctx workflow.Context, input *guardduty.DeleteDetectorInput) *GuarddutyDeleteDetectorResult

    DeleteFilter(ctx workflow.Context, input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error)
    DeleteFilterAsync(ctx workflow.Context, input *guardduty.DeleteFilterInput) *GuarddutyDeleteFilterResult

    DeleteIPSet(ctx workflow.Context, input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error)
    DeleteIPSetAsync(ctx workflow.Context, input *guardduty.DeleteIPSetInput) *GuarddutyDeleteIPSetResult

    DeleteInvitations(ctx workflow.Context, input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error)
    DeleteInvitationsAsync(ctx workflow.Context, input *guardduty.DeleteInvitationsInput) *GuarddutyDeleteInvitationsResult

    DeleteMembers(ctx workflow.Context, input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error)
    DeleteMembersAsync(ctx workflow.Context, input *guardduty.DeleteMembersInput) *GuarddutyDeleteMembersResult

    DeletePublishingDestination(ctx workflow.Context, input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error)
    DeletePublishingDestinationAsync(ctx workflow.Context, input *guardduty.DeletePublishingDestinationInput) *GuarddutyDeletePublishingDestinationResult

    DeleteThreatIntelSet(ctx workflow.Context, input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error)
    DeleteThreatIntelSetAsync(ctx workflow.Context, input *guardduty.DeleteThreatIntelSetInput) *GuarddutyDeleteThreatIntelSetResult

    DescribeOrganizationConfiguration(ctx workflow.Context, input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error)
    DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *guardduty.DescribeOrganizationConfigurationInput) *GuarddutyDescribeOrganizationConfigurationResult

    DescribePublishingDestination(ctx workflow.Context, input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error)
    DescribePublishingDestinationAsync(ctx workflow.Context, input *guardduty.DescribePublishingDestinationInput) *GuarddutyDescribePublishingDestinationResult

    DisableOrganizationAdminAccount(ctx workflow.Context, input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error)
    DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *guardduty.DisableOrganizationAdminAccountInput) *GuarddutyDisableOrganizationAdminAccountResult

    DisassociateFromMasterAccount(ctx workflow.Context, input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error)
    DisassociateFromMasterAccountAsync(ctx workflow.Context, input *guardduty.DisassociateFromMasterAccountInput) *GuarddutyDisassociateFromMasterAccountResult

    DisassociateMembers(ctx workflow.Context, input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error)
    DisassociateMembersAsync(ctx workflow.Context, input *guardduty.DisassociateMembersInput) *GuarddutyDisassociateMembersResult

    EnableOrganizationAdminAccount(ctx workflow.Context, input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error)
    EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *guardduty.EnableOrganizationAdminAccountInput) *GuarddutyEnableOrganizationAdminAccountResult

    GetDetector(ctx workflow.Context, input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error)
    GetDetectorAsync(ctx workflow.Context, input *guardduty.GetDetectorInput) *GuarddutyGetDetectorResult

    GetFilter(ctx workflow.Context, input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error)
    GetFilterAsync(ctx workflow.Context, input *guardduty.GetFilterInput) *GuarddutyGetFilterResult

    GetFindings(ctx workflow.Context, input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error)
    GetFindingsAsync(ctx workflow.Context, input *guardduty.GetFindingsInput) *GuarddutyGetFindingsResult

    GetFindingsStatistics(ctx workflow.Context, input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error)
    GetFindingsStatisticsAsync(ctx workflow.Context, input *guardduty.GetFindingsStatisticsInput) *GuarddutyGetFindingsStatisticsResult

    GetIPSet(ctx workflow.Context, input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error)
    GetIPSetAsync(ctx workflow.Context, input *guardduty.GetIPSetInput) *GuarddutyGetIPSetResult

    GetInvitationsCount(ctx workflow.Context, input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error)
    GetInvitationsCountAsync(ctx workflow.Context, input *guardduty.GetInvitationsCountInput) *GuarddutyGetInvitationsCountResult

    GetMasterAccount(ctx workflow.Context, input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error)
    GetMasterAccountAsync(ctx workflow.Context, input *guardduty.GetMasterAccountInput) *GuarddutyGetMasterAccountResult

    GetMemberDetectors(ctx workflow.Context, input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error)
    GetMemberDetectorsAsync(ctx workflow.Context, input *guardduty.GetMemberDetectorsInput) *GuarddutyGetMemberDetectorsResult

    GetMembers(ctx workflow.Context, input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error)
    GetMembersAsync(ctx workflow.Context, input *guardduty.GetMembersInput) *GuarddutyGetMembersResult

    GetThreatIntelSet(ctx workflow.Context, input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error)
    GetThreatIntelSetAsync(ctx workflow.Context, input *guardduty.GetThreatIntelSetInput) *GuarddutyGetThreatIntelSetResult

    GetUsageStatistics(ctx workflow.Context, input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error)
    GetUsageStatisticsAsync(ctx workflow.Context, input *guardduty.GetUsageStatisticsInput) *GuarddutyGetUsageStatisticsResult

    InviteMembers(ctx workflow.Context, input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error)
    InviteMembersAsync(ctx workflow.Context, input *guardduty.InviteMembersInput) *GuarddutyInviteMembersResult

    ListDetectors(ctx workflow.Context, input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error)
    ListDetectorsAsync(ctx workflow.Context, input *guardduty.ListDetectorsInput) *GuarddutyListDetectorsResult

    ListFilters(ctx workflow.Context, input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error)
    ListFiltersAsync(ctx workflow.Context, input *guardduty.ListFiltersInput) *GuarddutyListFiltersResult

    ListFindings(ctx workflow.Context, input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error)
    ListFindingsAsync(ctx workflow.Context, input *guardduty.ListFindingsInput) *GuarddutyListFindingsResult

    ListIPSets(ctx workflow.Context, input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error)
    ListIPSetsAsync(ctx workflow.Context, input *guardduty.ListIPSetsInput) *GuarddutyListIPSetsResult

    ListInvitations(ctx workflow.Context, input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error)
    ListInvitationsAsync(ctx workflow.Context, input *guardduty.ListInvitationsInput) *GuarddutyListInvitationsResult

    ListMembers(ctx workflow.Context, input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error)
    ListMembersAsync(ctx workflow.Context, input *guardduty.ListMembersInput) *GuarddutyListMembersResult

    ListOrganizationAdminAccounts(ctx workflow.Context, input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error)
    ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *guardduty.ListOrganizationAdminAccountsInput) *GuarddutyListOrganizationAdminAccountsResult

    ListPublishingDestinations(ctx workflow.Context, input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error)
    ListPublishingDestinationsAsync(ctx workflow.Context, input *guardduty.ListPublishingDestinationsInput) *GuarddutyListPublishingDestinationsResult

    ListTagsForResource(ctx workflow.Context, input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *guardduty.ListTagsForResourceInput) *GuarddutyListTagsForResourceResult

    ListThreatIntelSets(ctx workflow.Context, input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error)
    ListThreatIntelSetsAsync(ctx workflow.Context, input *guardduty.ListThreatIntelSetsInput) *GuarddutyListThreatIntelSetsResult

    StartMonitoringMembers(ctx workflow.Context, input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error)
    StartMonitoringMembersAsync(ctx workflow.Context, input *guardduty.StartMonitoringMembersInput) *GuarddutyStartMonitoringMembersResult

    StopMonitoringMembers(ctx workflow.Context, input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error)
    StopMonitoringMembersAsync(ctx workflow.Context, input *guardduty.StopMonitoringMembersInput) *GuarddutyStopMonitoringMembersResult

    TagResource(ctx workflow.Context, input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *guardduty.TagResourceInput) *GuarddutyTagResourceResult

    UnarchiveFindings(ctx workflow.Context, input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error)
    UnarchiveFindingsAsync(ctx workflow.Context, input *guardduty.UnarchiveFindingsInput) *GuarddutyUnarchiveFindingsResult

    UntagResource(ctx workflow.Context, input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *guardduty.UntagResourceInput) *GuarddutyUntagResourceResult

    UpdateDetector(ctx workflow.Context, input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error)
    UpdateDetectorAsync(ctx workflow.Context, input *guardduty.UpdateDetectorInput) *GuarddutyUpdateDetectorResult

    UpdateFilter(ctx workflow.Context, input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error)
    UpdateFilterAsync(ctx workflow.Context, input *guardduty.UpdateFilterInput) *GuarddutyUpdateFilterResult

    UpdateFindingsFeedback(ctx workflow.Context, input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error)
    UpdateFindingsFeedbackAsync(ctx workflow.Context, input *guardduty.UpdateFindingsFeedbackInput) *GuarddutyUpdateFindingsFeedbackResult

    UpdateIPSet(ctx workflow.Context, input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error)
    UpdateIPSetAsync(ctx workflow.Context, input *guardduty.UpdateIPSetInput) *GuarddutyUpdateIPSetResult

    UpdateMemberDetectors(ctx workflow.Context, input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error)
    UpdateMemberDetectorsAsync(ctx workflow.Context, input *guardduty.UpdateMemberDetectorsInput) *GuarddutyUpdateMemberDetectorsResult

    UpdateOrganizationConfiguration(ctx workflow.Context, input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error)
    UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *guardduty.UpdateOrganizationConfigurationInput) *GuarddutyUpdateOrganizationConfigurationResult

    UpdatePublishingDestination(ctx workflow.Context, input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error)
    UpdatePublishingDestinationAsync(ctx workflow.Context, input *guardduty.UpdatePublishingDestinationInput) *GuarddutyUpdatePublishingDestinationResult

    UpdateThreatIntelSet(ctx workflow.Context, input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error)
    UpdateThreatIntelSetAsync(ctx workflow.Context, input *guardduty.UpdateThreatIntelSetInput) *GuarddutyUpdateThreatIntelSetResult
}

type GuarddutyAcceptInvitationResult struct {
	Result workflow.Future
}

func (r *GuarddutyAcceptInvitationResult) Get(ctx workflow.Context) (*guardduty.AcceptInvitationOutput, error) {
    var output guardduty.AcceptInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyArchiveFindingsResult struct {
	Result workflow.Future
}

func (r *GuarddutyArchiveFindingsResult) Get(ctx workflow.Context) (*guardduty.ArchiveFindingsOutput, error) {
    var output guardduty.ArchiveFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreateDetectorResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreateDetectorResult) Get(ctx workflow.Context) (*guardduty.CreateDetectorOutput, error) {
    var output guardduty.CreateDetectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreateFilterResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreateFilterResult) Get(ctx workflow.Context) (*guardduty.CreateFilterOutput, error) {
    var output guardduty.CreateFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreateIPSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreateIPSetResult) Get(ctx workflow.Context) (*guardduty.CreateIPSetOutput, error) {
    var output guardduty.CreateIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreateMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreateMembersResult) Get(ctx workflow.Context) (*guardduty.CreateMembersOutput, error) {
    var output guardduty.CreateMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreatePublishingDestinationResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreatePublishingDestinationResult) Get(ctx workflow.Context) (*guardduty.CreatePublishingDestinationOutput, error) {
    var output guardduty.CreatePublishingDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreateSampleFindingsResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreateSampleFindingsResult) Get(ctx workflow.Context) (*guardduty.CreateSampleFindingsOutput, error) {
    var output guardduty.CreateSampleFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyCreateThreatIntelSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyCreateThreatIntelSetResult) Get(ctx workflow.Context) (*guardduty.CreateThreatIntelSetOutput, error) {
    var output guardduty.CreateThreatIntelSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeclineInvitationsResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeclineInvitationsResult) Get(ctx workflow.Context) (*guardduty.DeclineInvitationsOutput, error) {
    var output guardduty.DeclineInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeleteDetectorResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeleteDetectorResult) Get(ctx workflow.Context) (*guardduty.DeleteDetectorOutput, error) {
    var output guardduty.DeleteDetectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeleteFilterResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeleteFilterResult) Get(ctx workflow.Context) (*guardduty.DeleteFilterOutput, error) {
    var output guardduty.DeleteFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeleteIPSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeleteIPSetResult) Get(ctx workflow.Context) (*guardduty.DeleteIPSetOutput, error) {
    var output guardduty.DeleteIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeleteInvitationsResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeleteInvitationsResult) Get(ctx workflow.Context) (*guardduty.DeleteInvitationsOutput, error) {
    var output guardduty.DeleteInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeleteMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeleteMembersResult) Get(ctx workflow.Context) (*guardduty.DeleteMembersOutput, error) {
    var output guardduty.DeleteMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeletePublishingDestinationResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeletePublishingDestinationResult) Get(ctx workflow.Context) (*guardduty.DeletePublishingDestinationOutput, error) {
    var output guardduty.DeletePublishingDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDeleteThreatIntelSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyDeleteThreatIntelSetResult) Get(ctx workflow.Context) (*guardduty.DeleteThreatIntelSetOutput, error) {
    var output guardduty.DeleteThreatIntelSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDescribeOrganizationConfigurationResult struct {
	Result workflow.Future
}

func (r *GuarddutyDescribeOrganizationConfigurationResult) Get(ctx workflow.Context) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
    var output guardduty.DescribeOrganizationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDescribePublishingDestinationResult struct {
	Result workflow.Future
}

func (r *GuarddutyDescribePublishingDestinationResult) Get(ctx workflow.Context) (*guardduty.DescribePublishingDestinationOutput, error) {
    var output guardduty.DescribePublishingDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDisableOrganizationAdminAccountResult struct {
	Result workflow.Future
}

func (r *GuarddutyDisableOrganizationAdminAccountResult) Get(ctx workflow.Context) (*guardduty.DisableOrganizationAdminAccountOutput, error) {
    var output guardduty.DisableOrganizationAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDisassociateFromMasterAccountResult struct {
	Result workflow.Future
}

func (r *GuarddutyDisassociateFromMasterAccountResult) Get(ctx workflow.Context) (*guardduty.DisassociateFromMasterAccountOutput, error) {
    var output guardduty.DisassociateFromMasterAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyDisassociateMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyDisassociateMembersResult) Get(ctx workflow.Context) (*guardduty.DisassociateMembersOutput, error) {
    var output guardduty.DisassociateMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyEnableOrganizationAdminAccountResult struct {
	Result workflow.Future
}

func (r *GuarddutyEnableOrganizationAdminAccountResult) Get(ctx workflow.Context) (*guardduty.EnableOrganizationAdminAccountOutput, error) {
    var output guardduty.EnableOrganizationAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetDetectorResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetDetectorResult) Get(ctx workflow.Context) (*guardduty.GetDetectorOutput, error) {
    var output guardduty.GetDetectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetFilterResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetFilterResult) Get(ctx workflow.Context) (*guardduty.GetFilterOutput, error) {
    var output guardduty.GetFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetFindingsResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetFindingsResult) Get(ctx workflow.Context) (*guardduty.GetFindingsOutput, error) {
    var output guardduty.GetFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetFindingsStatisticsResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetFindingsStatisticsResult) Get(ctx workflow.Context) (*guardduty.GetFindingsStatisticsOutput, error) {
    var output guardduty.GetFindingsStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetIPSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetIPSetResult) Get(ctx workflow.Context) (*guardduty.GetIPSetOutput, error) {
    var output guardduty.GetIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetInvitationsCountResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetInvitationsCountResult) Get(ctx workflow.Context) (*guardduty.GetInvitationsCountOutput, error) {
    var output guardduty.GetInvitationsCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetMasterAccountResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetMasterAccountResult) Get(ctx workflow.Context) (*guardduty.GetMasterAccountOutput, error) {
    var output guardduty.GetMasterAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetMemberDetectorsResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetMemberDetectorsResult) Get(ctx workflow.Context) (*guardduty.GetMemberDetectorsOutput, error) {
    var output guardduty.GetMemberDetectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetMembersResult) Get(ctx workflow.Context) (*guardduty.GetMembersOutput, error) {
    var output guardduty.GetMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetThreatIntelSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetThreatIntelSetResult) Get(ctx workflow.Context) (*guardduty.GetThreatIntelSetOutput, error) {
    var output guardduty.GetThreatIntelSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyGetUsageStatisticsResult struct {
	Result workflow.Future
}

func (r *GuarddutyGetUsageStatisticsResult) Get(ctx workflow.Context) (*guardduty.GetUsageStatisticsOutput, error) {
    var output guardduty.GetUsageStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyInviteMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyInviteMembersResult) Get(ctx workflow.Context) (*guardduty.InviteMembersOutput, error) {
    var output guardduty.InviteMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListDetectorsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListDetectorsResult) Get(ctx workflow.Context) (*guardduty.ListDetectorsOutput, error) {
    var output guardduty.ListDetectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListFiltersResult struct {
	Result workflow.Future
}

func (r *GuarddutyListFiltersResult) Get(ctx workflow.Context) (*guardduty.ListFiltersOutput, error) {
    var output guardduty.ListFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListFindingsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListFindingsResult) Get(ctx workflow.Context) (*guardduty.ListFindingsOutput, error) {
    var output guardduty.ListFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListIPSetsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListIPSetsResult) Get(ctx workflow.Context) (*guardduty.ListIPSetsOutput, error) {
    var output guardduty.ListIPSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListInvitationsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListInvitationsResult) Get(ctx workflow.Context) (*guardduty.ListInvitationsOutput, error) {
    var output guardduty.ListInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyListMembersResult) Get(ctx workflow.Context) (*guardduty.ListMembersOutput, error) {
    var output guardduty.ListMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListOrganizationAdminAccountsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListOrganizationAdminAccountsResult) Get(ctx workflow.Context) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
    var output guardduty.ListOrganizationAdminAccountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListPublishingDestinationsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListPublishingDestinationsResult) Get(ctx workflow.Context) (*guardduty.ListPublishingDestinationsOutput, error) {
    var output guardduty.ListPublishingDestinationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *GuarddutyListTagsForResourceResult) Get(ctx workflow.Context) (*guardduty.ListTagsForResourceOutput, error) {
    var output guardduty.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyListThreatIntelSetsResult struct {
	Result workflow.Future
}

func (r *GuarddutyListThreatIntelSetsResult) Get(ctx workflow.Context) (*guardduty.ListThreatIntelSetsOutput, error) {
    var output guardduty.ListThreatIntelSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyStartMonitoringMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyStartMonitoringMembersResult) Get(ctx workflow.Context) (*guardduty.StartMonitoringMembersOutput, error) {
    var output guardduty.StartMonitoringMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyStopMonitoringMembersResult struct {
	Result workflow.Future
}

func (r *GuarddutyStopMonitoringMembersResult) Get(ctx workflow.Context) (*guardduty.StopMonitoringMembersOutput, error) {
    var output guardduty.StopMonitoringMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyTagResourceResult struct {
	Result workflow.Future
}

func (r *GuarddutyTagResourceResult) Get(ctx workflow.Context) (*guardduty.TagResourceOutput, error) {
    var output guardduty.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUnarchiveFindingsResult struct {
	Result workflow.Future
}

func (r *GuarddutyUnarchiveFindingsResult) Get(ctx workflow.Context) (*guardduty.UnarchiveFindingsOutput, error) {
    var output guardduty.UnarchiveFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUntagResourceResult struct {
	Result workflow.Future
}

func (r *GuarddutyUntagResourceResult) Get(ctx workflow.Context) (*guardduty.UntagResourceOutput, error) {
    var output guardduty.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateDetectorResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateDetectorResult) Get(ctx workflow.Context) (*guardduty.UpdateDetectorOutput, error) {
    var output guardduty.UpdateDetectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateFilterResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateFilterResult) Get(ctx workflow.Context) (*guardduty.UpdateFilterOutput, error) {
    var output guardduty.UpdateFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateFindingsFeedbackResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateFindingsFeedbackResult) Get(ctx workflow.Context) (*guardduty.UpdateFindingsFeedbackOutput, error) {
    var output guardduty.UpdateFindingsFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateIPSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateIPSetResult) Get(ctx workflow.Context) (*guardduty.UpdateIPSetOutput, error) {
    var output guardduty.UpdateIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateMemberDetectorsResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateMemberDetectorsResult) Get(ctx workflow.Context) (*guardduty.UpdateMemberDetectorsOutput, error) {
    var output guardduty.UpdateMemberDetectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateOrganizationConfigurationResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateOrganizationConfigurationResult) Get(ctx workflow.Context) (*guardduty.UpdateOrganizationConfigurationOutput, error) {
    var output guardduty.UpdateOrganizationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdatePublishingDestinationResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdatePublishingDestinationResult) Get(ctx workflow.Context) (*guardduty.UpdatePublishingDestinationOutput, error) {
    var output guardduty.UpdatePublishingDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuarddutyUpdateThreatIntelSetResult struct {
	Result workflow.Future
}

func (r *GuarddutyUpdateThreatIntelSetResult) Get(ctx workflow.Context) (*guardduty.UpdateThreatIntelSetOutput, error) {
    var output guardduty.UpdateThreatIntelSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GuardDutyStub struct {
    activities awsactivities.GuardDutyActivities
}

func NewGuardDutyStub() GuardDutyClient {
    return &GuardDutyStub{}
}

func (a *GuardDutyStub) AcceptInvitation(ctx workflow.Context, input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error) {
    var output guardduty.AcceptInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) AcceptInvitationAsync(ctx workflow.Context, input *guardduty.AcceptInvitationInput) *GuarddutyAcceptInvitationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input)
    return &GuarddutyAcceptInvitationResult{Result: future}
}

func (a *GuardDutyStub) ArchiveFindings(ctx workflow.Context, input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error) {
    var output guardduty.ArchiveFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ArchiveFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ArchiveFindingsAsync(ctx workflow.Context, input *guardduty.ArchiveFindingsInput) *GuarddutyArchiveFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ArchiveFindings, input)
    return &GuarddutyArchiveFindingsResult{Result: future}
}

func (a *GuardDutyStub) CreateDetector(ctx workflow.Context, input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error) {
    var output guardduty.CreateDetectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDetector, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreateDetectorAsync(ctx workflow.Context, input *guardduty.CreateDetectorInput) *GuarddutyCreateDetectorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDetector, input)
    return &GuarddutyCreateDetectorResult{Result: future}
}

func (a *GuardDutyStub) CreateFilter(ctx workflow.Context, input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error) {
    var output guardduty.CreateFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreateFilterAsync(ctx workflow.Context, input *guardduty.CreateFilterInput) *GuarddutyCreateFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFilter, input)
    return &GuarddutyCreateFilterResult{Result: future}
}

func (a *GuardDutyStub) CreateIPSet(ctx workflow.Context, input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error) {
    var output guardduty.CreateIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreateIPSetAsync(ctx workflow.Context, input *guardduty.CreateIPSetInput) *GuarddutyCreateIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input)
    return &GuarddutyCreateIPSetResult{Result: future}
}

func (a *GuardDutyStub) CreateMembers(ctx workflow.Context, input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error) {
    var output guardduty.CreateMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreateMembersAsync(ctx workflow.Context, input *guardduty.CreateMembersInput) *GuarddutyCreateMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMembers, input)
    return &GuarddutyCreateMembersResult{Result: future}
}

func (a *GuardDutyStub) CreatePublishingDestination(ctx workflow.Context, input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error) {
    var output guardduty.CreatePublishingDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePublishingDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreatePublishingDestinationAsync(ctx workflow.Context, input *guardduty.CreatePublishingDestinationInput) *GuarddutyCreatePublishingDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePublishingDestination, input)
    return &GuarddutyCreatePublishingDestinationResult{Result: future}
}

func (a *GuardDutyStub) CreateSampleFindings(ctx workflow.Context, input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error) {
    var output guardduty.CreateSampleFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSampleFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreateSampleFindingsAsync(ctx workflow.Context, input *guardduty.CreateSampleFindingsInput) *GuarddutyCreateSampleFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSampleFindings, input)
    return &GuarddutyCreateSampleFindingsResult{Result: future}
}

func (a *GuardDutyStub) CreateThreatIntelSet(ctx workflow.Context, input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error) {
    var output guardduty.CreateThreatIntelSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateThreatIntelSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) CreateThreatIntelSetAsync(ctx workflow.Context, input *guardduty.CreateThreatIntelSetInput) *GuarddutyCreateThreatIntelSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateThreatIntelSet, input)
    return &GuarddutyCreateThreatIntelSetResult{Result: future}
}

func (a *GuardDutyStub) DeclineInvitations(ctx workflow.Context, input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error) {
    var output guardduty.DeclineInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeclineInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeclineInvitationsAsync(ctx workflow.Context, input *guardduty.DeclineInvitationsInput) *GuarddutyDeclineInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeclineInvitations, input)
    return &GuarddutyDeclineInvitationsResult{Result: future}
}

func (a *GuardDutyStub) DeleteDetector(ctx workflow.Context, input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error) {
    var output guardduty.DeleteDetectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDetector, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeleteDetectorAsync(ctx workflow.Context, input *guardduty.DeleteDetectorInput) *GuarddutyDeleteDetectorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDetector, input)
    return &GuarddutyDeleteDetectorResult{Result: future}
}

func (a *GuardDutyStub) DeleteFilter(ctx workflow.Context, input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error) {
    var output guardduty.DeleteFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeleteFilterAsync(ctx workflow.Context, input *guardduty.DeleteFilterInput) *GuarddutyDeleteFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFilter, input)
    return &GuarddutyDeleteFilterResult{Result: future}
}

func (a *GuardDutyStub) DeleteIPSet(ctx workflow.Context, input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error) {
    var output guardduty.DeleteIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeleteIPSetAsync(ctx workflow.Context, input *guardduty.DeleteIPSetInput) *GuarddutyDeleteIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input)
    return &GuarddutyDeleteIPSetResult{Result: future}
}

func (a *GuardDutyStub) DeleteInvitations(ctx workflow.Context, input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error) {
    var output guardduty.DeleteInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeleteInvitationsAsync(ctx workflow.Context, input *guardduty.DeleteInvitationsInput) *GuarddutyDeleteInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInvitations, input)
    return &GuarddutyDeleteInvitationsResult{Result: future}
}

func (a *GuardDutyStub) DeleteMembers(ctx workflow.Context, input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error) {
    var output guardduty.DeleteMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeleteMembersAsync(ctx workflow.Context, input *guardduty.DeleteMembersInput) *GuarddutyDeleteMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMembers, input)
    return &GuarddutyDeleteMembersResult{Result: future}
}

func (a *GuardDutyStub) DeletePublishingDestination(ctx workflow.Context, input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error) {
    var output guardduty.DeletePublishingDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePublishingDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeletePublishingDestinationAsync(ctx workflow.Context, input *guardduty.DeletePublishingDestinationInput) *GuarddutyDeletePublishingDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePublishingDestination, input)
    return &GuarddutyDeletePublishingDestinationResult{Result: future}
}

func (a *GuardDutyStub) DeleteThreatIntelSet(ctx workflow.Context, input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error) {
    var output guardduty.DeleteThreatIntelSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteThreatIntelSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DeleteThreatIntelSetAsync(ctx workflow.Context, input *guardduty.DeleteThreatIntelSetInput) *GuarddutyDeleteThreatIntelSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteThreatIntelSet, input)
    return &GuarddutyDeleteThreatIntelSetResult{Result: future}
}

func (a *GuardDutyStub) DescribeOrganizationConfiguration(ctx workflow.Context, input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
    var output guardduty.DescribeOrganizationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *guardduty.DescribeOrganizationConfigurationInput) *GuarddutyDescribeOrganizationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfiguration, input)
    return &GuarddutyDescribeOrganizationConfigurationResult{Result: future}
}

func (a *GuardDutyStub) DescribePublishingDestination(ctx workflow.Context, input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error) {
    var output guardduty.DescribePublishingDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePublishingDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DescribePublishingDestinationAsync(ctx workflow.Context, input *guardduty.DescribePublishingDestinationInput) *GuarddutyDescribePublishingDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePublishingDestination, input)
    return &GuarddutyDescribePublishingDestinationResult{Result: future}
}

func (a *GuardDutyStub) DisableOrganizationAdminAccount(ctx workflow.Context, input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error) {
    var output guardduty.DisableOrganizationAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableOrganizationAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *guardduty.DisableOrganizationAdminAccountInput) *GuarddutyDisableOrganizationAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableOrganizationAdminAccount, input)
    return &GuarddutyDisableOrganizationAdminAccountResult{Result: future}
}

func (a *GuardDutyStub) DisassociateFromMasterAccount(ctx workflow.Context, input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error) {
    var output guardduty.DisassociateFromMasterAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateFromMasterAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DisassociateFromMasterAccountAsync(ctx workflow.Context, input *guardduty.DisassociateFromMasterAccountInput) *GuarddutyDisassociateFromMasterAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateFromMasterAccount, input)
    return &GuarddutyDisassociateFromMasterAccountResult{Result: future}
}

func (a *GuardDutyStub) DisassociateMembers(ctx workflow.Context, input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error) {
    var output guardduty.DisassociateMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) DisassociateMembersAsync(ctx workflow.Context, input *guardduty.DisassociateMembersInput) *GuarddutyDisassociateMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateMembers, input)
    return &GuarddutyDisassociateMembersResult{Result: future}
}

func (a *GuardDutyStub) EnableOrganizationAdminAccount(ctx workflow.Context, input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error) {
    var output guardduty.EnableOrganizationAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableOrganizationAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *guardduty.EnableOrganizationAdminAccountInput) *GuarddutyEnableOrganizationAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableOrganizationAdminAccount, input)
    return &GuarddutyEnableOrganizationAdminAccountResult{Result: future}
}

func (a *GuardDutyStub) GetDetector(ctx workflow.Context, input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error) {
    var output guardduty.GetDetectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDetector, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetDetectorAsync(ctx workflow.Context, input *guardduty.GetDetectorInput) *GuarddutyGetDetectorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDetector, input)
    return &GuarddutyGetDetectorResult{Result: future}
}

func (a *GuardDutyStub) GetFilter(ctx workflow.Context, input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error) {
    var output guardduty.GetFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetFilterAsync(ctx workflow.Context, input *guardduty.GetFilterInput) *GuarddutyGetFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFilter, input)
    return &GuarddutyGetFilterResult{Result: future}
}

func (a *GuardDutyStub) GetFindings(ctx workflow.Context, input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error) {
    var output guardduty.GetFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetFindingsAsync(ctx workflow.Context, input *guardduty.GetFindingsInput) *GuarddutyGetFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFindings, input)
    return &GuarddutyGetFindingsResult{Result: future}
}

func (a *GuardDutyStub) GetFindingsStatistics(ctx workflow.Context, input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error) {
    var output guardduty.GetFindingsStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindingsStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetFindingsStatisticsAsync(ctx workflow.Context, input *guardduty.GetFindingsStatisticsInput) *GuarddutyGetFindingsStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFindingsStatistics, input)
    return &GuarddutyGetFindingsStatisticsResult{Result: future}
}

func (a *GuardDutyStub) GetIPSet(ctx workflow.Context, input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error) {
    var output guardduty.GetIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetIPSetAsync(ctx workflow.Context, input *guardduty.GetIPSetInput) *GuarddutyGetIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input)
    return &GuarddutyGetIPSetResult{Result: future}
}

func (a *GuardDutyStub) GetInvitationsCount(ctx workflow.Context, input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error) {
    var output guardduty.GetInvitationsCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInvitationsCount, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetInvitationsCountAsync(ctx workflow.Context, input *guardduty.GetInvitationsCountInput) *GuarddutyGetInvitationsCountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInvitationsCount, input)
    return &GuarddutyGetInvitationsCountResult{Result: future}
}

func (a *GuardDutyStub) GetMasterAccount(ctx workflow.Context, input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error) {
    var output guardduty.GetMasterAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMasterAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetMasterAccountAsync(ctx workflow.Context, input *guardduty.GetMasterAccountInput) *GuarddutyGetMasterAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMasterAccount, input)
    return &GuarddutyGetMasterAccountResult{Result: future}
}

func (a *GuardDutyStub) GetMemberDetectors(ctx workflow.Context, input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error) {
    var output guardduty.GetMemberDetectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMemberDetectors, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetMemberDetectorsAsync(ctx workflow.Context, input *guardduty.GetMemberDetectorsInput) *GuarddutyGetMemberDetectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMemberDetectors, input)
    return &GuarddutyGetMemberDetectorsResult{Result: future}
}

func (a *GuardDutyStub) GetMembers(ctx workflow.Context, input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error) {
    var output guardduty.GetMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetMembersAsync(ctx workflow.Context, input *guardduty.GetMembersInput) *GuarddutyGetMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMembers, input)
    return &GuarddutyGetMembersResult{Result: future}
}

func (a *GuardDutyStub) GetThreatIntelSet(ctx workflow.Context, input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error) {
    var output guardduty.GetThreatIntelSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetThreatIntelSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetThreatIntelSetAsync(ctx workflow.Context, input *guardduty.GetThreatIntelSetInput) *GuarddutyGetThreatIntelSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetThreatIntelSet, input)
    return &GuarddutyGetThreatIntelSetResult{Result: future}
}

func (a *GuardDutyStub) GetUsageStatistics(ctx workflow.Context, input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error) {
    var output guardduty.GetUsageStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUsageStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) GetUsageStatisticsAsync(ctx workflow.Context, input *guardduty.GetUsageStatisticsInput) *GuarddutyGetUsageStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUsageStatistics, input)
    return &GuarddutyGetUsageStatisticsResult{Result: future}
}

func (a *GuardDutyStub) InviteMembers(ctx workflow.Context, input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error) {
    var output guardduty.InviteMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InviteMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) InviteMembersAsync(ctx workflow.Context, input *guardduty.InviteMembersInput) *GuarddutyInviteMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InviteMembers, input)
    return &GuarddutyInviteMembersResult{Result: future}
}

func (a *GuardDutyStub) ListDetectors(ctx workflow.Context, input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error) {
    var output guardduty.ListDetectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDetectors, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListDetectorsAsync(ctx workflow.Context, input *guardduty.ListDetectorsInput) *GuarddutyListDetectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDetectors, input)
    return &GuarddutyListDetectorsResult{Result: future}
}

func (a *GuardDutyStub) ListFilters(ctx workflow.Context, input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error) {
    var output guardduty.ListFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListFiltersAsync(ctx workflow.Context, input *guardduty.ListFiltersInput) *GuarddutyListFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFilters, input)
    return &GuarddutyListFiltersResult{Result: future}
}

func (a *GuardDutyStub) ListFindings(ctx workflow.Context, input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error) {
    var output guardduty.ListFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListFindingsAsync(ctx workflow.Context, input *guardduty.ListFindingsInput) *GuarddutyListFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input)
    return &GuarddutyListFindingsResult{Result: future}
}

func (a *GuardDutyStub) ListIPSets(ctx workflow.Context, input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error) {
    var output guardduty.ListIPSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListIPSetsAsync(ctx workflow.Context, input *guardduty.ListIPSetsInput) *GuarddutyListIPSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input)
    return &GuarddutyListIPSetsResult{Result: future}
}

func (a *GuardDutyStub) ListInvitations(ctx workflow.Context, input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error) {
    var output guardduty.ListInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListInvitationsAsync(ctx workflow.Context, input *guardduty.ListInvitationsInput) *GuarddutyListInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input)
    return &GuarddutyListInvitationsResult{Result: future}
}

func (a *GuardDutyStub) ListMembers(ctx workflow.Context, input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error) {
    var output guardduty.ListMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListMembersAsync(ctx workflow.Context, input *guardduty.ListMembersInput) *GuarddutyListMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input)
    return &GuarddutyListMembersResult{Result: future}
}

func (a *GuardDutyStub) ListOrganizationAdminAccounts(ctx workflow.Context, input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
    var output guardduty.ListOrganizationAdminAccountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationAdminAccounts, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *guardduty.ListOrganizationAdminAccountsInput) *GuarddutyListOrganizationAdminAccountsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationAdminAccounts, input)
    return &GuarddutyListOrganizationAdminAccountsResult{Result: future}
}

func (a *GuardDutyStub) ListPublishingDestinations(ctx workflow.Context, input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error) {
    var output guardduty.ListPublishingDestinationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPublishingDestinations, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListPublishingDestinationsAsync(ctx workflow.Context, input *guardduty.ListPublishingDestinationsInput) *GuarddutyListPublishingDestinationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPublishingDestinations, input)
    return &GuarddutyListPublishingDestinationsResult{Result: future}
}

func (a *GuardDutyStub) ListTagsForResource(ctx workflow.Context, input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error) {
    var output guardduty.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListTagsForResourceAsync(ctx workflow.Context, input *guardduty.ListTagsForResourceInput) *GuarddutyListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &GuarddutyListTagsForResourceResult{Result: future}
}

func (a *GuardDutyStub) ListThreatIntelSets(ctx workflow.Context, input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error) {
    var output guardduty.ListThreatIntelSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThreatIntelSets, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) ListThreatIntelSetsAsync(ctx workflow.Context, input *guardduty.ListThreatIntelSetsInput) *GuarddutyListThreatIntelSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListThreatIntelSets, input)
    return &GuarddutyListThreatIntelSetsResult{Result: future}
}

func (a *GuardDutyStub) StartMonitoringMembers(ctx workflow.Context, input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error) {
    var output guardduty.StartMonitoringMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMonitoringMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) StartMonitoringMembersAsync(ctx workflow.Context, input *guardduty.StartMonitoringMembersInput) *GuarddutyStartMonitoringMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMonitoringMembers, input)
    return &GuarddutyStartMonitoringMembersResult{Result: future}
}

func (a *GuardDutyStub) StopMonitoringMembers(ctx workflow.Context, input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error) {
    var output guardduty.StopMonitoringMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopMonitoringMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) StopMonitoringMembersAsync(ctx workflow.Context, input *guardduty.StopMonitoringMembersInput) *GuarddutyStopMonitoringMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopMonitoringMembers, input)
    return &GuarddutyStopMonitoringMembersResult{Result: future}
}

func (a *GuardDutyStub) TagResource(ctx workflow.Context, input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error) {
    var output guardduty.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) TagResourceAsync(ctx workflow.Context, input *guardduty.TagResourceInput) *GuarddutyTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &GuarddutyTagResourceResult{Result: future}
}

func (a *GuardDutyStub) UnarchiveFindings(ctx workflow.Context, input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error) {
    var output guardduty.UnarchiveFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnarchiveFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UnarchiveFindingsAsync(ctx workflow.Context, input *guardduty.UnarchiveFindingsInput) *GuarddutyUnarchiveFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnarchiveFindings, input)
    return &GuarddutyUnarchiveFindingsResult{Result: future}
}

func (a *GuardDutyStub) UntagResource(ctx workflow.Context, input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error) {
    var output guardduty.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UntagResourceAsync(ctx workflow.Context, input *guardduty.UntagResourceInput) *GuarddutyUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &GuarddutyUntagResourceResult{Result: future}
}

func (a *GuardDutyStub) UpdateDetector(ctx workflow.Context, input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error) {
    var output guardduty.UpdateDetectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDetector, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateDetectorAsync(ctx workflow.Context, input *guardduty.UpdateDetectorInput) *GuarddutyUpdateDetectorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDetector, input)
    return &GuarddutyUpdateDetectorResult{Result: future}
}

func (a *GuardDutyStub) UpdateFilter(ctx workflow.Context, input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error) {
    var output guardduty.UpdateFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateFilterAsync(ctx workflow.Context, input *guardduty.UpdateFilterInput) *GuarddutyUpdateFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFilter, input)
    return &GuarddutyUpdateFilterResult{Result: future}
}

func (a *GuardDutyStub) UpdateFindingsFeedback(ctx workflow.Context, input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error) {
    var output guardduty.UpdateFindingsFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFindingsFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateFindingsFeedbackAsync(ctx workflow.Context, input *guardduty.UpdateFindingsFeedbackInput) *GuarddutyUpdateFindingsFeedbackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFindingsFeedback, input)
    return &GuarddutyUpdateFindingsFeedbackResult{Result: future}
}

func (a *GuardDutyStub) UpdateIPSet(ctx workflow.Context, input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error) {
    var output guardduty.UpdateIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateIPSetAsync(ctx workflow.Context, input *guardduty.UpdateIPSetInput) *GuarddutyUpdateIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input)
    return &GuarddutyUpdateIPSetResult{Result: future}
}

func (a *GuardDutyStub) UpdateMemberDetectors(ctx workflow.Context, input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error) {
    var output guardduty.UpdateMemberDetectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMemberDetectors, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateMemberDetectorsAsync(ctx workflow.Context, input *guardduty.UpdateMemberDetectorsInput) *GuarddutyUpdateMemberDetectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMemberDetectors, input)
    return &GuarddutyUpdateMemberDetectorsResult{Result: future}
}

func (a *GuardDutyStub) UpdateOrganizationConfiguration(ctx workflow.Context, input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error) {
    var output guardduty.UpdateOrganizationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOrganizationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *guardduty.UpdateOrganizationConfigurationInput) *GuarddutyUpdateOrganizationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateOrganizationConfiguration, input)
    return &GuarddutyUpdateOrganizationConfigurationResult{Result: future}
}

func (a *GuardDutyStub) UpdatePublishingDestination(ctx workflow.Context, input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error) {
    var output guardduty.UpdatePublishingDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePublishingDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdatePublishingDestinationAsync(ctx workflow.Context, input *guardduty.UpdatePublishingDestinationInput) *GuarddutyUpdatePublishingDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePublishingDestination, input)
    return &GuarddutyUpdatePublishingDestinationResult{Result: future}
}

func (a *GuardDutyStub) UpdateThreatIntelSet(ctx workflow.Context, input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error) {
    var output guardduty.UpdateThreatIntelSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateThreatIntelSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GuardDutyStub) UpdateThreatIntelSetAsync(ctx workflow.Context, input *guardduty.UpdateThreatIntelSetInput) *GuarddutyUpdateThreatIntelSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateThreatIntelSet, input)
    return &GuarddutyUpdateThreatIntelSetResult{Result: future}
}