package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codecommit"
	"go.temporal.io/sdk/workflow"
)

type CodeCommitClient interface {
       AssociateApprovalRuleTemplateWithRepository(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error)
       AssociateApprovalRuleTemplateWithRepositoryAsync(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) *CodecommitAssociateApprovalRuleTemplateWithRepositoryResult

       BatchAssociateApprovalRuleTemplateWithRepositories(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error)
       BatchAssociateApprovalRuleTemplateWithRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) *CodecommitBatchAssociateApprovalRuleTemplateWithRepositoriesResult

       BatchDescribeMergeConflicts(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error)
       BatchDescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) *CodecommitBatchDescribeMergeConflictsResult

       BatchDisassociateApprovalRuleTemplateFromRepositories(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error)
       BatchDisassociateApprovalRuleTemplateFromRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) *CodecommitBatchDisassociateApprovalRuleTemplateFromRepositoriesResult

       BatchGetCommits(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error)
       BatchGetCommitsAsync(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) *CodecommitBatchGetCommitsResult

       BatchGetRepositories(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error)
       BatchGetRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) *CodecommitBatchGetRepositoriesResult

       CreateApprovalRuleTemplate(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) (*codecommit.CreateApprovalRuleTemplateOutput, error)
       CreateApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) *CodecommitCreateApprovalRuleTemplateResult

       CreateBranch(ctx workflow.Context, input *codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error)
       CreateBranchAsync(ctx workflow.Context, input *codecommit.CreateBranchInput) *CodecommitCreateBranchResult

       CreateCommit(ctx workflow.Context, input *codecommit.CreateCommitInput) (*codecommit.CreateCommitOutput, error)
       CreateCommitAsync(ctx workflow.Context, input *codecommit.CreateCommitInput) *CodecommitCreateCommitResult

       CreatePullRequestApprovalRule(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) (*codecommit.CreatePullRequestApprovalRuleOutput, error)
       CreatePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) *CodecommitCreatePullRequestApprovalRuleResult

       CreateRepository(ctx workflow.Context, input *codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error)
       CreateRepositoryAsync(ctx workflow.Context, input *codecommit.CreateRepositoryInput) *CodecommitCreateRepositoryResult

       CreateUnreferencedMergeCommit(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) (*codecommit.CreateUnreferencedMergeCommitOutput, error)
       CreateUnreferencedMergeCommitAsync(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) *CodecommitCreateUnreferencedMergeCommitResult

       DeleteApprovalRuleTemplate(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) (*codecommit.DeleteApprovalRuleTemplateOutput, error)
       DeleteApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) *CodecommitDeleteApprovalRuleTemplateResult

       DeleteBranch(ctx workflow.Context, input *codecommit.DeleteBranchInput) (*codecommit.DeleteBranchOutput, error)
       DeleteBranchAsync(ctx workflow.Context, input *codecommit.DeleteBranchInput) *CodecommitDeleteBranchResult

       DeleteCommentContent(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) (*codecommit.DeleteCommentContentOutput, error)
       DeleteCommentContentAsync(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) *CodecommitDeleteCommentContentResult

       DeleteFile(ctx workflow.Context, input *codecommit.DeleteFileInput) (*codecommit.DeleteFileOutput, error)
       DeleteFileAsync(ctx workflow.Context, input *codecommit.DeleteFileInput) *CodecommitDeleteFileResult

       DeletePullRequestApprovalRule(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) (*codecommit.DeletePullRequestApprovalRuleOutput, error)
       DeletePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) *CodecommitDeletePullRequestApprovalRuleResult

       DeleteRepository(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error)
       DeleteRepositoryAsync(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) *CodecommitDeleteRepositoryResult

       DescribeMergeConflicts(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error)
       DescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) *CodecommitDescribeMergeConflictsResult

       DescribePullRequestEvents(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error)
       DescribePullRequestEventsAsync(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) *CodecommitDescribePullRequestEventsResult

       DisassociateApprovalRuleTemplateFromRepository(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error)
       DisassociateApprovalRuleTemplateFromRepositoryAsync(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) *CodecommitDisassociateApprovalRuleTemplateFromRepositoryResult

       EvaluatePullRequestApprovalRules(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error)
       EvaluatePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) *CodecommitEvaluatePullRequestApprovalRulesResult

       GetApprovalRuleTemplate(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error)
       GetApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) *CodecommitGetApprovalRuleTemplateResult

       GetBlob(ctx workflow.Context, input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error)
       GetBlobAsync(ctx workflow.Context, input *codecommit.GetBlobInput) *CodecommitGetBlobResult

       GetBranch(ctx workflow.Context, input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error)
       GetBranchAsync(ctx workflow.Context, input *codecommit.GetBranchInput) *CodecommitGetBranchResult

       GetComment(ctx workflow.Context, input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error)
       GetCommentAsync(ctx workflow.Context, input *codecommit.GetCommentInput) *CodecommitGetCommentResult

       GetCommentReactions(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error)
       GetCommentReactionsAsync(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) *CodecommitGetCommentReactionsResult

       GetCommentsForComparedCommit(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error)
       GetCommentsForComparedCommitAsync(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) *CodecommitGetCommentsForComparedCommitResult

       GetCommit(ctx workflow.Context, input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error)
       GetCommitAsync(ctx workflow.Context, input *codecommit.GetCommitInput) *CodecommitGetCommitResult

       GetDifferences(ctx workflow.Context, input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error)
       GetDifferencesAsync(ctx workflow.Context, input *codecommit.GetDifferencesInput) *CodecommitGetDifferencesResult

       GetFile(ctx workflow.Context, input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error)
       GetFileAsync(ctx workflow.Context, input *codecommit.GetFileInput) *CodecommitGetFileResult

       GetFolder(ctx workflow.Context, input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error)
       GetFolderAsync(ctx workflow.Context, input *codecommit.GetFolderInput) *CodecommitGetFolderResult

       GetMergeCommit(ctx workflow.Context, input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error)
       GetMergeCommitAsync(ctx workflow.Context, input *codecommit.GetMergeCommitInput) *CodecommitGetMergeCommitResult

       GetMergeConflicts(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error)
       GetMergeConflictsAsync(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) *CodecommitGetMergeConflictsResult

       GetMergeOptions(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error)
       GetMergeOptionsAsync(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) *CodecommitGetMergeOptionsResult

       GetPullRequestApprovalStates(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error)
       GetPullRequestApprovalStatesAsync(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) *CodecommitGetPullRequestApprovalStatesResult

       GetPullRequestOverrideState(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error)
       GetPullRequestOverrideStateAsync(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) *CodecommitGetPullRequestOverrideStateResult

       GetRepository(ctx workflow.Context, input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error)
       GetRepositoryAsync(ctx workflow.Context, input *codecommit.GetRepositoryInput) *CodecommitGetRepositoryResult

       GetRepositoryTriggers(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error)
       GetRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) *CodecommitGetRepositoryTriggersResult

       ListApprovalRuleTemplates(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error)
       ListApprovalRuleTemplatesAsync(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) *CodecommitListApprovalRuleTemplatesResult

       ListAssociatedApprovalRuleTemplatesForRepository(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error)
       ListAssociatedApprovalRuleTemplatesForRepositoryAsync(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) *CodecommitListAssociatedApprovalRuleTemplatesForRepositoryResult

       ListBranches(ctx workflow.Context, input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error)
       ListBranchesAsync(ctx workflow.Context, input *codecommit.ListBranchesInput) *CodecommitListBranchesResult

       ListPullRequests(ctx workflow.Context, input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error)
       ListPullRequestsAsync(ctx workflow.Context, input *codecommit.ListPullRequestsInput) *CodecommitListPullRequestsResult

       ListRepositories(ctx workflow.Context, input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error)
       ListRepositoriesAsync(ctx workflow.Context, input *codecommit.ListRepositoriesInput) *CodecommitListRepositoriesResult

       ListRepositoriesForApprovalRuleTemplate(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error)
       ListRepositoriesForApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) *CodecommitListRepositoriesForApprovalRuleTemplateResult

       ListTagsForResource(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) *CodecommitListTagsForResourceResult

       MergeBranchesByFastForward(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) (*codecommit.MergeBranchesByFastForwardOutput, error)
       MergeBranchesByFastForwardAsync(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) *CodecommitMergeBranchesByFastForwardResult

       MergeBranchesBySquash(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) (*codecommit.MergeBranchesBySquashOutput, error)
       MergeBranchesBySquashAsync(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) *CodecommitMergeBranchesBySquashResult

       MergeBranchesByThreeWay(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) (*codecommit.MergeBranchesByThreeWayOutput, error)
       MergeBranchesByThreeWayAsync(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) *CodecommitMergeBranchesByThreeWayResult

       MergePullRequestByFastForward(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) (*codecommit.MergePullRequestByFastForwardOutput, error)
       MergePullRequestByFastForwardAsync(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) *CodecommitMergePullRequestByFastForwardResult

       MergePullRequestBySquash(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) (*codecommit.MergePullRequestBySquashOutput, error)
       MergePullRequestBySquashAsync(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) *CodecommitMergePullRequestBySquashResult

       MergePullRequestByThreeWay(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) (*codecommit.MergePullRequestByThreeWayOutput, error)
       MergePullRequestByThreeWayAsync(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) *CodecommitMergePullRequestByThreeWayResult

       OverridePullRequestApprovalRules(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) (*codecommit.OverridePullRequestApprovalRulesOutput, error)
       OverridePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) *CodecommitOverridePullRequestApprovalRulesResult

       PostCommentForComparedCommit(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) (*codecommit.PostCommentForComparedCommitOutput, error)
       PostCommentForComparedCommitAsync(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) *CodecommitPostCommentForComparedCommitResult

       PostCommentReply(ctx workflow.Context, input *codecommit.PostCommentReplyInput) (*codecommit.PostCommentReplyOutput, error)
       PostCommentReplyAsync(ctx workflow.Context, input *codecommit.PostCommentReplyInput) *CodecommitPostCommentReplyResult

       PutCommentReaction(ctx workflow.Context, input *codecommit.PutCommentReactionInput) (*codecommit.PutCommentReactionOutput, error)
       PutCommentReactionAsync(ctx workflow.Context, input *codecommit.PutCommentReactionInput) *CodecommitPutCommentReactionResult

       PutFile(ctx workflow.Context, input *codecommit.PutFileInput) (*codecommit.PutFileOutput, error)
       PutFileAsync(ctx workflow.Context, input *codecommit.PutFileInput) *CodecommitPutFileResult

       PutRepositoryTriggers(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error)
       PutRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) *CodecommitPutRepositoryTriggersResult

       TagResource(ctx workflow.Context, input *codecommit.TagResourceInput) (*codecommit.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *codecommit.TagResourceInput) *CodecommitTagResourceResult

       TestRepositoryTriggers(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error)
       TestRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) *CodecommitTestRepositoryTriggersResult

       UntagResource(ctx workflow.Context, input *codecommit.UntagResourceInput) (*codecommit.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *codecommit.UntagResourceInput) *CodecommitUntagResourceResult

       UpdateApprovalRuleTemplateContent(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error)
       UpdateApprovalRuleTemplateContentAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) *CodecommitUpdateApprovalRuleTemplateContentResult

       UpdateApprovalRuleTemplateDescription(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error)
       UpdateApprovalRuleTemplateDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) *CodecommitUpdateApprovalRuleTemplateDescriptionResult

       UpdateApprovalRuleTemplateName(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error)
       UpdateApprovalRuleTemplateNameAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) *CodecommitUpdateApprovalRuleTemplateNameResult

       UpdateComment(ctx workflow.Context, input *codecommit.UpdateCommentInput) (*codecommit.UpdateCommentOutput, error)
       UpdateCommentAsync(ctx workflow.Context, input *codecommit.UpdateCommentInput) *CodecommitUpdateCommentResult

       UpdateDefaultBranch(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error)
       UpdateDefaultBranchAsync(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) *CodecommitUpdateDefaultBranchResult

       UpdatePullRequestApprovalRuleContent(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error)
       UpdatePullRequestApprovalRuleContentAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) *CodecommitUpdatePullRequestApprovalRuleContentResult

       UpdatePullRequestApprovalState(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) (*codecommit.UpdatePullRequestApprovalStateOutput, error)
       UpdatePullRequestApprovalStateAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) *CodecommitUpdatePullRequestApprovalStateResult

       UpdatePullRequestDescription(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) (*codecommit.UpdatePullRequestDescriptionOutput, error)
       UpdatePullRequestDescriptionAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) *CodecommitUpdatePullRequestDescriptionResult

       UpdatePullRequestStatus(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) (*codecommit.UpdatePullRequestStatusOutput, error)
       UpdatePullRequestStatusAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) *CodecommitUpdatePullRequestStatusResult

       UpdatePullRequestTitle(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) (*codecommit.UpdatePullRequestTitleOutput, error)
       UpdatePullRequestTitleAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) *CodecommitUpdatePullRequestTitleResult

       UpdateRepositoryDescription(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error)
       UpdateRepositoryDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) *CodecommitUpdateRepositoryDescriptionResult

       UpdateRepositoryName(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error)
       UpdateRepositoryNameAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) *CodecommitUpdateRepositoryNameResult
}

type CodeCommitStub struct {
}

func NewCodeCommitStub() CodeCommitClient {
    return &CodeCommitStub{}
}

type CodecommitAssociateApprovalRuleTemplateWithRepositoryResult struct {
	Result workflow.Future
}

func (r *CodecommitAssociateApprovalRuleTemplateWithRepositoryResult) Get(ctx workflow.Context) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error) {
    var output codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitBatchAssociateApprovalRuleTemplateWithRepositoriesResult struct {
	Result workflow.Future
}

func (r *CodecommitBatchAssociateApprovalRuleTemplateWithRepositoriesResult) Get(ctx workflow.Context) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
    var output codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitBatchDescribeMergeConflictsResult struct {
	Result workflow.Future
}

func (r *CodecommitBatchDescribeMergeConflictsResult) Get(ctx workflow.Context) (*codecommit.BatchDescribeMergeConflictsOutput, error) {
    var output codecommit.BatchDescribeMergeConflictsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitBatchDisassociateApprovalRuleTemplateFromRepositoriesResult struct {
	Result workflow.Future
}

func (r *CodecommitBatchDisassociateApprovalRuleTemplateFromRepositoriesResult) Get(ctx workflow.Context) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
    var output codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitBatchGetCommitsResult struct {
	Result workflow.Future
}

func (r *CodecommitBatchGetCommitsResult) Get(ctx workflow.Context) (*codecommit.BatchGetCommitsOutput, error) {
    var output codecommit.BatchGetCommitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitBatchGetRepositoriesResult struct {
	Result workflow.Future
}

func (r *CodecommitBatchGetRepositoriesResult) Get(ctx workflow.Context) (*codecommit.BatchGetRepositoriesOutput, error) {
    var output codecommit.BatchGetRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitCreateApprovalRuleTemplateResult struct {
	Result workflow.Future
}

func (r *CodecommitCreateApprovalRuleTemplateResult) Get(ctx workflow.Context) (*codecommit.CreateApprovalRuleTemplateOutput, error) {
    var output codecommit.CreateApprovalRuleTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitCreateBranchResult struct {
	Result workflow.Future
}

func (r *CodecommitCreateBranchResult) Get(ctx workflow.Context) (*codecommit.CreateBranchOutput, error) {
    var output codecommit.CreateBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitCreateCommitResult struct {
	Result workflow.Future
}

func (r *CodecommitCreateCommitResult) Get(ctx workflow.Context) (*codecommit.CreateCommitOutput, error) {
    var output codecommit.CreateCommitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitCreatePullRequestApprovalRuleResult struct {
	Result workflow.Future
}

func (r *CodecommitCreatePullRequestApprovalRuleResult) Get(ctx workflow.Context) (*codecommit.CreatePullRequestApprovalRuleOutput, error) {
    var output codecommit.CreatePullRequestApprovalRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitCreateRepositoryResult struct {
	Result workflow.Future
}

func (r *CodecommitCreateRepositoryResult) Get(ctx workflow.Context) (*codecommit.CreateRepositoryOutput, error) {
    var output codecommit.CreateRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitCreateUnreferencedMergeCommitResult struct {
	Result workflow.Future
}

func (r *CodecommitCreateUnreferencedMergeCommitResult) Get(ctx workflow.Context) (*codecommit.CreateUnreferencedMergeCommitOutput, error) {
    var output codecommit.CreateUnreferencedMergeCommitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDeleteApprovalRuleTemplateResult struct {
	Result workflow.Future
}

func (r *CodecommitDeleteApprovalRuleTemplateResult) Get(ctx workflow.Context) (*codecommit.DeleteApprovalRuleTemplateOutput, error) {
    var output codecommit.DeleteApprovalRuleTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDeleteBranchResult struct {
	Result workflow.Future
}

func (r *CodecommitDeleteBranchResult) Get(ctx workflow.Context) (*codecommit.DeleteBranchOutput, error) {
    var output codecommit.DeleteBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDeleteCommentContentResult struct {
	Result workflow.Future
}

func (r *CodecommitDeleteCommentContentResult) Get(ctx workflow.Context) (*codecommit.DeleteCommentContentOutput, error) {
    var output codecommit.DeleteCommentContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDeleteFileResult struct {
	Result workflow.Future
}

func (r *CodecommitDeleteFileResult) Get(ctx workflow.Context) (*codecommit.DeleteFileOutput, error) {
    var output codecommit.DeleteFileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDeletePullRequestApprovalRuleResult struct {
	Result workflow.Future
}

func (r *CodecommitDeletePullRequestApprovalRuleResult) Get(ctx workflow.Context) (*codecommit.DeletePullRequestApprovalRuleOutput, error) {
    var output codecommit.DeletePullRequestApprovalRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDeleteRepositoryResult struct {
	Result workflow.Future
}

func (r *CodecommitDeleteRepositoryResult) Get(ctx workflow.Context) (*codecommit.DeleteRepositoryOutput, error) {
    var output codecommit.DeleteRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDescribeMergeConflictsResult struct {
	Result workflow.Future
}

func (r *CodecommitDescribeMergeConflictsResult) Get(ctx workflow.Context) (*codecommit.DescribeMergeConflictsOutput, error) {
    var output codecommit.DescribeMergeConflictsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDescribePullRequestEventsResult struct {
	Result workflow.Future
}

func (r *CodecommitDescribePullRequestEventsResult) Get(ctx workflow.Context) (*codecommit.DescribePullRequestEventsOutput, error) {
    var output codecommit.DescribePullRequestEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitDisassociateApprovalRuleTemplateFromRepositoryResult struct {
	Result workflow.Future
}

func (r *CodecommitDisassociateApprovalRuleTemplateFromRepositoryResult) Get(ctx workflow.Context) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error) {
    var output codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitEvaluatePullRequestApprovalRulesResult struct {
	Result workflow.Future
}

func (r *CodecommitEvaluatePullRequestApprovalRulesResult) Get(ctx workflow.Context) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error) {
    var output codecommit.EvaluatePullRequestApprovalRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetApprovalRuleTemplateResult struct {
	Result workflow.Future
}

func (r *CodecommitGetApprovalRuleTemplateResult) Get(ctx workflow.Context) (*codecommit.GetApprovalRuleTemplateOutput, error) {
    var output codecommit.GetApprovalRuleTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetBlobResult struct {
	Result workflow.Future
}

func (r *CodecommitGetBlobResult) Get(ctx workflow.Context) (*codecommit.GetBlobOutput, error) {
    var output codecommit.GetBlobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetBranchResult struct {
	Result workflow.Future
}

func (r *CodecommitGetBranchResult) Get(ctx workflow.Context) (*codecommit.GetBranchOutput, error) {
    var output codecommit.GetBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetCommentResult struct {
	Result workflow.Future
}

func (r *CodecommitGetCommentResult) Get(ctx workflow.Context) (*codecommit.GetCommentOutput, error) {
    var output codecommit.GetCommentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetCommentReactionsResult struct {
	Result workflow.Future
}

func (r *CodecommitGetCommentReactionsResult) Get(ctx workflow.Context) (*codecommit.GetCommentReactionsOutput, error) {
    var output codecommit.GetCommentReactionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetCommentsForComparedCommitResult struct {
	Result workflow.Future
}

func (r *CodecommitGetCommentsForComparedCommitResult) Get(ctx workflow.Context) (*codecommit.GetCommentsForComparedCommitOutput, error) {
    var output codecommit.GetCommentsForComparedCommitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetCommitResult struct {
	Result workflow.Future
}

func (r *CodecommitGetCommitResult) Get(ctx workflow.Context) (*codecommit.GetCommitOutput, error) {
    var output codecommit.GetCommitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetDifferencesResult struct {
	Result workflow.Future
}

func (r *CodecommitGetDifferencesResult) Get(ctx workflow.Context) (*codecommit.GetDifferencesOutput, error) {
    var output codecommit.GetDifferencesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetFileResult struct {
	Result workflow.Future
}

func (r *CodecommitGetFileResult) Get(ctx workflow.Context) (*codecommit.GetFileOutput, error) {
    var output codecommit.GetFileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetFolderResult struct {
	Result workflow.Future
}

func (r *CodecommitGetFolderResult) Get(ctx workflow.Context) (*codecommit.GetFolderOutput, error) {
    var output codecommit.GetFolderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetMergeCommitResult struct {
	Result workflow.Future
}

func (r *CodecommitGetMergeCommitResult) Get(ctx workflow.Context) (*codecommit.GetMergeCommitOutput, error) {
    var output codecommit.GetMergeCommitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetMergeConflictsResult struct {
	Result workflow.Future
}

func (r *CodecommitGetMergeConflictsResult) Get(ctx workflow.Context) (*codecommit.GetMergeConflictsOutput, error) {
    var output codecommit.GetMergeConflictsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetMergeOptionsResult struct {
	Result workflow.Future
}

func (r *CodecommitGetMergeOptionsResult) Get(ctx workflow.Context) (*codecommit.GetMergeOptionsOutput, error) {
    var output codecommit.GetMergeOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetPullRequestApprovalStatesResult struct {
	Result workflow.Future
}

func (r *CodecommitGetPullRequestApprovalStatesResult) Get(ctx workflow.Context) (*codecommit.GetPullRequestApprovalStatesOutput, error) {
    var output codecommit.GetPullRequestApprovalStatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetPullRequestOverrideStateResult struct {
	Result workflow.Future
}

func (r *CodecommitGetPullRequestOverrideStateResult) Get(ctx workflow.Context) (*codecommit.GetPullRequestOverrideStateOutput, error) {
    var output codecommit.GetPullRequestOverrideStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetRepositoryResult struct {
	Result workflow.Future
}

func (r *CodecommitGetRepositoryResult) Get(ctx workflow.Context) (*codecommit.GetRepositoryOutput, error) {
    var output codecommit.GetRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitGetRepositoryTriggersResult struct {
	Result workflow.Future
}

func (r *CodecommitGetRepositoryTriggersResult) Get(ctx workflow.Context) (*codecommit.GetRepositoryTriggersOutput, error) {
    var output codecommit.GetRepositoryTriggersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListApprovalRuleTemplatesResult struct {
	Result workflow.Future
}

func (r *CodecommitListApprovalRuleTemplatesResult) Get(ctx workflow.Context) (*codecommit.ListApprovalRuleTemplatesOutput, error) {
    var output codecommit.ListApprovalRuleTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListAssociatedApprovalRuleTemplatesForRepositoryResult struct {
	Result workflow.Future
}

func (r *CodecommitListAssociatedApprovalRuleTemplatesForRepositoryResult) Get(ctx workflow.Context) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
    var output codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListBranchesResult struct {
	Result workflow.Future
}

func (r *CodecommitListBranchesResult) Get(ctx workflow.Context) (*codecommit.ListBranchesOutput, error) {
    var output codecommit.ListBranchesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListPullRequestsResult struct {
	Result workflow.Future
}

func (r *CodecommitListPullRequestsResult) Get(ctx workflow.Context) (*codecommit.ListPullRequestsOutput, error) {
    var output codecommit.ListPullRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListRepositoriesResult struct {
	Result workflow.Future
}

func (r *CodecommitListRepositoriesResult) Get(ctx workflow.Context) (*codecommit.ListRepositoriesOutput, error) {
    var output codecommit.ListRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListRepositoriesForApprovalRuleTemplateResult struct {
	Result workflow.Future
}

func (r *CodecommitListRepositoriesForApprovalRuleTemplateResult) Get(ctx workflow.Context) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {
    var output codecommit.ListRepositoriesForApprovalRuleTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CodecommitListTagsForResourceResult) Get(ctx workflow.Context) (*codecommit.ListTagsForResourceOutput, error) {
    var output codecommit.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitMergeBranchesByFastForwardResult struct {
	Result workflow.Future
}

func (r *CodecommitMergeBranchesByFastForwardResult) Get(ctx workflow.Context) (*codecommit.MergeBranchesByFastForwardOutput, error) {
    var output codecommit.MergeBranchesByFastForwardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitMergeBranchesBySquashResult struct {
	Result workflow.Future
}

func (r *CodecommitMergeBranchesBySquashResult) Get(ctx workflow.Context) (*codecommit.MergeBranchesBySquashOutput, error) {
    var output codecommit.MergeBranchesBySquashOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitMergeBranchesByThreeWayResult struct {
	Result workflow.Future
}

func (r *CodecommitMergeBranchesByThreeWayResult) Get(ctx workflow.Context) (*codecommit.MergeBranchesByThreeWayOutput, error) {
    var output codecommit.MergeBranchesByThreeWayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitMergePullRequestByFastForwardResult struct {
	Result workflow.Future
}

func (r *CodecommitMergePullRequestByFastForwardResult) Get(ctx workflow.Context) (*codecommit.MergePullRequestByFastForwardOutput, error) {
    var output codecommit.MergePullRequestByFastForwardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitMergePullRequestBySquashResult struct {
	Result workflow.Future
}

func (r *CodecommitMergePullRequestBySquashResult) Get(ctx workflow.Context) (*codecommit.MergePullRequestBySquashOutput, error) {
    var output codecommit.MergePullRequestBySquashOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitMergePullRequestByThreeWayResult struct {
	Result workflow.Future
}

func (r *CodecommitMergePullRequestByThreeWayResult) Get(ctx workflow.Context) (*codecommit.MergePullRequestByThreeWayOutput, error) {
    var output codecommit.MergePullRequestByThreeWayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitOverridePullRequestApprovalRulesResult struct {
	Result workflow.Future
}

func (r *CodecommitOverridePullRequestApprovalRulesResult) Get(ctx workflow.Context) (*codecommit.OverridePullRequestApprovalRulesOutput, error) {
    var output codecommit.OverridePullRequestApprovalRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitPostCommentForComparedCommitResult struct {
	Result workflow.Future
}

func (r *CodecommitPostCommentForComparedCommitResult) Get(ctx workflow.Context) (*codecommit.PostCommentForComparedCommitOutput, error) {
    var output codecommit.PostCommentForComparedCommitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitPostCommentReplyResult struct {
	Result workflow.Future
}

func (r *CodecommitPostCommentReplyResult) Get(ctx workflow.Context) (*codecommit.PostCommentReplyOutput, error) {
    var output codecommit.PostCommentReplyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitPutCommentReactionResult struct {
	Result workflow.Future
}

func (r *CodecommitPutCommentReactionResult) Get(ctx workflow.Context) (*codecommit.PutCommentReactionOutput, error) {
    var output codecommit.PutCommentReactionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitPutFileResult struct {
	Result workflow.Future
}

func (r *CodecommitPutFileResult) Get(ctx workflow.Context) (*codecommit.PutFileOutput, error) {
    var output codecommit.PutFileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitPutRepositoryTriggersResult struct {
	Result workflow.Future
}

func (r *CodecommitPutRepositoryTriggersResult) Get(ctx workflow.Context) (*codecommit.PutRepositoryTriggersOutput, error) {
    var output codecommit.PutRepositoryTriggersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitTagResourceResult struct {
	Result workflow.Future
}

func (r *CodecommitTagResourceResult) Get(ctx workflow.Context) (*codecommit.TagResourceOutput, error) {
    var output codecommit.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitTestRepositoryTriggersResult struct {
	Result workflow.Future
}

func (r *CodecommitTestRepositoryTriggersResult) Get(ctx workflow.Context) (*codecommit.TestRepositoryTriggersOutput, error) {
    var output codecommit.TestRepositoryTriggersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUntagResourceResult struct {
	Result workflow.Future
}

func (r *CodecommitUntagResourceResult) Get(ctx workflow.Context) (*codecommit.UntagResourceOutput, error) {
    var output codecommit.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateApprovalRuleTemplateContentResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateApprovalRuleTemplateContentResult) Get(ctx workflow.Context) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error) {
    var output codecommit.UpdateApprovalRuleTemplateContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateApprovalRuleTemplateDescriptionResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateApprovalRuleTemplateDescriptionResult) Get(ctx workflow.Context) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error) {
    var output codecommit.UpdateApprovalRuleTemplateDescriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateApprovalRuleTemplateNameResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateApprovalRuleTemplateNameResult) Get(ctx workflow.Context) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error) {
    var output codecommit.UpdateApprovalRuleTemplateNameOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateCommentResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateCommentResult) Get(ctx workflow.Context) (*codecommit.UpdateCommentOutput, error) {
    var output codecommit.UpdateCommentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateDefaultBranchResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateDefaultBranchResult) Get(ctx workflow.Context) (*codecommit.UpdateDefaultBranchOutput, error) {
    var output codecommit.UpdateDefaultBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdatePullRequestApprovalRuleContentResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdatePullRequestApprovalRuleContentResult) Get(ctx workflow.Context) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error) {
    var output codecommit.UpdatePullRequestApprovalRuleContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdatePullRequestApprovalStateResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdatePullRequestApprovalStateResult) Get(ctx workflow.Context) (*codecommit.UpdatePullRequestApprovalStateOutput, error) {
    var output codecommit.UpdatePullRequestApprovalStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdatePullRequestDescriptionResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdatePullRequestDescriptionResult) Get(ctx workflow.Context) (*codecommit.UpdatePullRequestDescriptionOutput, error) {
    var output codecommit.UpdatePullRequestDescriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdatePullRequestStatusResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdatePullRequestStatusResult) Get(ctx workflow.Context) (*codecommit.UpdatePullRequestStatusOutput, error) {
    var output codecommit.UpdatePullRequestStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdatePullRequestTitleResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdatePullRequestTitleResult) Get(ctx workflow.Context) (*codecommit.UpdatePullRequestTitleOutput, error) {
    var output codecommit.UpdatePullRequestTitleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateRepositoryDescriptionResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateRepositoryDescriptionResult) Get(ctx workflow.Context) (*codecommit.UpdateRepositoryDescriptionOutput, error) {
    var output codecommit.UpdateRepositoryDescriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodecommitUpdateRepositoryNameResult struct {
	Result workflow.Future
}

func (r *CodecommitUpdateRepositoryNameResult) Get(ctx workflow.Context) (*codecommit.UpdateRepositoryNameOutput, error) {
    var output codecommit.UpdateRepositoryNameOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) AssociateApprovalRuleTemplateWithRepository(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error) {
    var output codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.AssociateApprovalRuleTemplateWithRepository", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) AssociateApprovalRuleTemplateWithRepositoryAsync(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) *CodecommitAssociateApprovalRuleTemplateWithRepositoryResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.AssociateApprovalRuleTemplateWithRepository", input)
    return &CodecommitAssociateApprovalRuleTemplateWithRepositoryResult{Result: future}
}

func (a *CodeCommitStub) BatchAssociateApprovalRuleTemplateWithRepositories(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
    var output codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.BatchAssociateApprovalRuleTemplateWithRepositories", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) BatchAssociateApprovalRuleTemplateWithRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) *CodecommitBatchAssociateApprovalRuleTemplateWithRepositoriesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.BatchAssociateApprovalRuleTemplateWithRepositories", input)
    return &CodecommitBatchAssociateApprovalRuleTemplateWithRepositoriesResult{Result: future}
}

func (a *CodeCommitStub) BatchDescribeMergeConflicts(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error) {
    var output codecommit.BatchDescribeMergeConflictsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.BatchDescribeMergeConflicts", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) BatchDescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) *CodecommitBatchDescribeMergeConflictsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.BatchDescribeMergeConflicts", input)
    return &CodecommitBatchDescribeMergeConflictsResult{Result: future}
}

func (a *CodeCommitStub) BatchDisassociateApprovalRuleTemplateFromRepositories(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
    var output codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.BatchDisassociateApprovalRuleTemplateFromRepositories", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) BatchDisassociateApprovalRuleTemplateFromRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) *CodecommitBatchDisassociateApprovalRuleTemplateFromRepositoriesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.BatchDisassociateApprovalRuleTemplateFromRepositories", input)
    return &CodecommitBatchDisassociateApprovalRuleTemplateFromRepositoriesResult{Result: future}
}

func (a *CodeCommitStub) BatchGetCommits(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error) {
    var output codecommit.BatchGetCommitsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.BatchGetCommits", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) BatchGetCommitsAsync(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) *CodecommitBatchGetCommitsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.BatchGetCommits", input)
    return &CodecommitBatchGetCommitsResult{Result: future}
}

func (a *CodeCommitStub) BatchGetRepositories(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error) {
    var output codecommit.BatchGetRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.BatchGetRepositories", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) BatchGetRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) *CodecommitBatchGetRepositoriesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.BatchGetRepositories", input)
    return &CodecommitBatchGetRepositoriesResult{Result: future}
}

func (a *CodeCommitStub) CreateApprovalRuleTemplate(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) (*codecommit.CreateApprovalRuleTemplateOutput, error) {
    var output codecommit.CreateApprovalRuleTemplateOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.CreateApprovalRuleTemplate", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) CreateApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) *CodecommitCreateApprovalRuleTemplateResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.CreateApprovalRuleTemplate", input)
    return &CodecommitCreateApprovalRuleTemplateResult{Result: future}
}

func (a *CodeCommitStub) CreateBranch(ctx workflow.Context, input *codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error) {
    var output codecommit.CreateBranchOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.CreateBranch", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) CreateBranchAsync(ctx workflow.Context, input *codecommit.CreateBranchInput) *CodecommitCreateBranchResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.CreateBranch", input)
    return &CodecommitCreateBranchResult{Result: future}
}

func (a *CodeCommitStub) CreateCommit(ctx workflow.Context, input *codecommit.CreateCommitInput) (*codecommit.CreateCommitOutput, error) {
    var output codecommit.CreateCommitOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.CreateCommit", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) CreateCommitAsync(ctx workflow.Context, input *codecommit.CreateCommitInput) *CodecommitCreateCommitResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.CreateCommit", input)
    return &CodecommitCreateCommitResult{Result: future}
}

func (a *CodeCommitStub) CreatePullRequestApprovalRule(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) (*codecommit.CreatePullRequestApprovalRuleOutput, error) {
    var output codecommit.CreatePullRequestApprovalRuleOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.CreatePullRequestApprovalRule", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) CreatePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) *CodecommitCreatePullRequestApprovalRuleResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.CreatePullRequestApprovalRule", input)
    return &CodecommitCreatePullRequestApprovalRuleResult{Result: future}
}

func (a *CodeCommitStub) CreateRepository(ctx workflow.Context, input *codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error) {
    var output codecommit.CreateRepositoryOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.CreateRepository", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) CreateRepositoryAsync(ctx workflow.Context, input *codecommit.CreateRepositoryInput) *CodecommitCreateRepositoryResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.CreateRepository", input)
    return &CodecommitCreateRepositoryResult{Result: future}
}

func (a *CodeCommitStub) CreateUnreferencedMergeCommit(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) (*codecommit.CreateUnreferencedMergeCommitOutput, error) {
    var output codecommit.CreateUnreferencedMergeCommitOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.CreateUnreferencedMergeCommit", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) CreateUnreferencedMergeCommitAsync(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) *CodecommitCreateUnreferencedMergeCommitResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.CreateUnreferencedMergeCommit", input)
    return &CodecommitCreateUnreferencedMergeCommitResult{Result: future}
}

func (a *CodeCommitStub) DeleteApprovalRuleTemplate(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) (*codecommit.DeleteApprovalRuleTemplateOutput, error) {
    var output codecommit.DeleteApprovalRuleTemplateOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteApprovalRuleTemplate", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DeleteApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) *CodecommitDeleteApprovalRuleTemplateResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteApprovalRuleTemplate", input)
    return &CodecommitDeleteApprovalRuleTemplateResult{Result: future}
}

func (a *CodeCommitStub) DeleteBranch(ctx workflow.Context, input *codecommit.DeleteBranchInput) (*codecommit.DeleteBranchOutput, error) {
    var output codecommit.DeleteBranchOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteBranch", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DeleteBranchAsync(ctx workflow.Context, input *codecommit.DeleteBranchInput) *CodecommitDeleteBranchResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteBranch", input)
    return &CodecommitDeleteBranchResult{Result: future}
}

func (a *CodeCommitStub) DeleteCommentContent(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) (*codecommit.DeleteCommentContentOutput, error) {
    var output codecommit.DeleteCommentContentOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteCommentContent", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DeleteCommentContentAsync(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) *CodecommitDeleteCommentContentResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteCommentContent", input)
    return &CodecommitDeleteCommentContentResult{Result: future}
}

func (a *CodeCommitStub) DeleteFile(ctx workflow.Context, input *codecommit.DeleteFileInput) (*codecommit.DeleteFileOutput, error) {
    var output codecommit.DeleteFileOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteFile", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DeleteFileAsync(ctx workflow.Context, input *codecommit.DeleteFileInput) *CodecommitDeleteFileResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteFile", input)
    return &CodecommitDeleteFileResult{Result: future}
}

func (a *CodeCommitStub) DeletePullRequestApprovalRule(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) (*codecommit.DeletePullRequestApprovalRuleOutput, error) {
    var output codecommit.DeletePullRequestApprovalRuleOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DeletePullRequestApprovalRule", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DeletePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) *CodecommitDeletePullRequestApprovalRuleResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DeletePullRequestApprovalRule", input)
    return &CodecommitDeletePullRequestApprovalRuleResult{Result: future}
}

func (a *CodeCommitStub) DeleteRepository(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error) {
    var output codecommit.DeleteRepositoryOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteRepository", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DeleteRepositoryAsync(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) *CodecommitDeleteRepositoryResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DeleteRepository", input)
    return &CodecommitDeleteRepositoryResult{Result: future}
}

func (a *CodeCommitStub) DescribeMergeConflicts(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error) {
    var output codecommit.DescribeMergeConflictsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DescribeMergeConflicts", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) *CodecommitDescribeMergeConflictsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DescribeMergeConflicts", input)
    return &CodecommitDescribeMergeConflictsResult{Result: future}
}

func (a *CodeCommitStub) DescribePullRequestEvents(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error) {
    var output codecommit.DescribePullRequestEventsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DescribePullRequestEvents", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DescribePullRequestEventsAsync(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) *CodecommitDescribePullRequestEventsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DescribePullRequestEvents", input)
    return &CodecommitDescribePullRequestEventsResult{Result: future}
}

func (a *CodeCommitStub) DisassociateApprovalRuleTemplateFromRepository(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error) {
    var output codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.DisassociateApprovalRuleTemplateFromRepository", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) DisassociateApprovalRuleTemplateFromRepositoryAsync(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) *CodecommitDisassociateApprovalRuleTemplateFromRepositoryResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.DisassociateApprovalRuleTemplateFromRepository", input)
    return &CodecommitDisassociateApprovalRuleTemplateFromRepositoryResult{Result: future}
}

func (a *CodeCommitStub) EvaluatePullRequestApprovalRules(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error) {
    var output codecommit.EvaluatePullRequestApprovalRulesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.EvaluatePullRequestApprovalRules", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) EvaluatePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) *CodecommitEvaluatePullRequestApprovalRulesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.EvaluatePullRequestApprovalRules", input)
    return &CodecommitEvaluatePullRequestApprovalRulesResult{Result: future}
}

func (a *CodeCommitStub) GetApprovalRuleTemplate(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error) {
    var output codecommit.GetApprovalRuleTemplateOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetApprovalRuleTemplate", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) *CodecommitGetApprovalRuleTemplateResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetApprovalRuleTemplate", input)
    return &CodecommitGetApprovalRuleTemplateResult{Result: future}
}

func (a *CodeCommitStub) GetBlob(ctx workflow.Context, input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error) {
    var output codecommit.GetBlobOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetBlob", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetBlobAsync(ctx workflow.Context, input *codecommit.GetBlobInput) *CodecommitGetBlobResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetBlob", input)
    return &CodecommitGetBlobResult{Result: future}
}

func (a *CodeCommitStub) GetBranch(ctx workflow.Context, input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error) {
    var output codecommit.GetBranchOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetBranch", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetBranchAsync(ctx workflow.Context, input *codecommit.GetBranchInput) *CodecommitGetBranchResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetBranch", input)
    return &CodecommitGetBranchResult{Result: future}
}

func (a *CodeCommitStub) GetComment(ctx workflow.Context, input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error) {
    var output codecommit.GetCommentOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetComment", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetCommentAsync(ctx workflow.Context, input *codecommit.GetCommentInput) *CodecommitGetCommentResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetComment", input)
    return &CodecommitGetCommentResult{Result: future}
}

func (a *CodeCommitStub) GetCommentReactions(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error) {
    var output codecommit.GetCommentReactionsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetCommentReactions", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetCommentReactionsAsync(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) *CodecommitGetCommentReactionsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetCommentReactions", input)
    return &CodecommitGetCommentReactionsResult{Result: future}
}

func (a *CodeCommitStub) GetCommentsForComparedCommit(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error) {
    var output codecommit.GetCommentsForComparedCommitOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetCommentsForComparedCommit", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetCommentsForComparedCommitAsync(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) *CodecommitGetCommentsForComparedCommitResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetCommentsForComparedCommit", input)
    return &CodecommitGetCommentsForComparedCommitResult{Result: future}
}

func (a *CodeCommitStub) GetCommit(ctx workflow.Context, input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error) {
    var output codecommit.GetCommitOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetCommit", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetCommitAsync(ctx workflow.Context, input *codecommit.GetCommitInput) *CodecommitGetCommitResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetCommit", input)
    return &CodecommitGetCommitResult{Result: future}
}

func (a *CodeCommitStub) GetDifferences(ctx workflow.Context, input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error) {
    var output codecommit.GetDifferencesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetDifferences", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetDifferencesAsync(ctx workflow.Context, input *codecommit.GetDifferencesInput) *CodecommitGetDifferencesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetDifferences", input)
    return &CodecommitGetDifferencesResult{Result: future}
}

func (a *CodeCommitStub) GetFile(ctx workflow.Context, input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error) {
    var output codecommit.GetFileOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetFile", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetFileAsync(ctx workflow.Context, input *codecommit.GetFileInput) *CodecommitGetFileResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetFile", input)
    return &CodecommitGetFileResult{Result: future}
}

func (a *CodeCommitStub) GetFolder(ctx workflow.Context, input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error) {
    var output codecommit.GetFolderOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetFolder", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetFolderAsync(ctx workflow.Context, input *codecommit.GetFolderInput) *CodecommitGetFolderResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetFolder", input)
    return &CodecommitGetFolderResult{Result: future}
}

func (a *CodeCommitStub) GetMergeCommit(ctx workflow.Context, input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error) {
    var output codecommit.GetMergeCommitOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetMergeCommit", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetMergeCommitAsync(ctx workflow.Context, input *codecommit.GetMergeCommitInput) *CodecommitGetMergeCommitResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetMergeCommit", input)
    return &CodecommitGetMergeCommitResult{Result: future}
}

func (a *CodeCommitStub) GetMergeConflicts(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error) {
    var output codecommit.GetMergeConflictsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetMergeConflicts", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetMergeConflictsAsync(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) *CodecommitGetMergeConflictsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetMergeConflicts", input)
    return &CodecommitGetMergeConflictsResult{Result: future}
}

func (a *CodeCommitStub) GetMergeOptions(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error) {
    var output codecommit.GetMergeOptionsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetMergeOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetMergeOptionsAsync(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) *CodecommitGetMergeOptionsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetMergeOptions", input)
    return &CodecommitGetMergeOptionsResult{Result: future}
}

func (a *CodeCommitStub) GetPullRequestApprovalStates(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error) {
    var output codecommit.GetPullRequestApprovalStatesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetPullRequestApprovalStates", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetPullRequestApprovalStatesAsync(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) *CodecommitGetPullRequestApprovalStatesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetPullRequestApprovalStates", input)
    return &CodecommitGetPullRequestApprovalStatesResult{Result: future}
}

func (a *CodeCommitStub) GetPullRequestOverrideState(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error) {
    var output codecommit.GetPullRequestOverrideStateOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetPullRequestOverrideState", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetPullRequestOverrideStateAsync(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) *CodecommitGetPullRequestOverrideStateResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetPullRequestOverrideState", input)
    return &CodecommitGetPullRequestOverrideStateResult{Result: future}
}

func (a *CodeCommitStub) GetRepository(ctx workflow.Context, input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error) {
    var output codecommit.GetRepositoryOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetRepository", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetRepositoryAsync(ctx workflow.Context, input *codecommit.GetRepositoryInput) *CodecommitGetRepositoryResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetRepository", input)
    return &CodecommitGetRepositoryResult{Result: future}
}

func (a *CodeCommitStub) GetRepositoryTriggers(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error) {
    var output codecommit.GetRepositoryTriggersOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.GetRepositoryTriggers", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) GetRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) *CodecommitGetRepositoryTriggersResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.GetRepositoryTriggers", input)
    return &CodecommitGetRepositoryTriggersResult{Result: future}
}

func (a *CodeCommitStub) ListApprovalRuleTemplates(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error) {
    var output codecommit.ListApprovalRuleTemplatesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListApprovalRuleTemplates", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListApprovalRuleTemplatesAsync(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) *CodecommitListApprovalRuleTemplatesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListApprovalRuleTemplates", input)
    return &CodecommitListApprovalRuleTemplatesResult{Result: future}
}

func (a *CodeCommitStub) ListAssociatedApprovalRuleTemplatesForRepository(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
    var output codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListAssociatedApprovalRuleTemplatesForRepository", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListAssociatedApprovalRuleTemplatesForRepositoryAsync(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) *CodecommitListAssociatedApprovalRuleTemplatesForRepositoryResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListAssociatedApprovalRuleTemplatesForRepository", input)
    return &CodecommitListAssociatedApprovalRuleTemplatesForRepositoryResult{Result: future}
}

func (a *CodeCommitStub) ListBranches(ctx workflow.Context, input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error) {
    var output codecommit.ListBranchesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListBranches", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListBranchesAsync(ctx workflow.Context, input *codecommit.ListBranchesInput) *CodecommitListBranchesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListBranches", input)
    return &CodecommitListBranchesResult{Result: future}
}

func (a *CodeCommitStub) ListPullRequests(ctx workflow.Context, input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error) {
    var output codecommit.ListPullRequestsOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListPullRequests", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListPullRequestsAsync(ctx workflow.Context, input *codecommit.ListPullRequestsInput) *CodecommitListPullRequestsResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListPullRequests", input)
    return &CodecommitListPullRequestsResult{Result: future}
}

func (a *CodeCommitStub) ListRepositories(ctx workflow.Context, input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error) {
    var output codecommit.ListRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListRepositories", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListRepositoriesAsync(ctx workflow.Context, input *codecommit.ListRepositoriesInput) *CodecommitListRepositoriesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListRepositories", input)
    return &CodecommitListRepositoriesResult{Result: future}
}

func (a *CodeCommitStub) ListRepositoriesForApprovalRuleTemplate(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {
    var output codecommit.ListRepositoriesForApprovalRuleTemplateOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListRepositoriesForApprovalRuleTemplate", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListRepositoriesForApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) *CodecommitListRepositoriesForApprovalRuleTemplateResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListRepositoriesForApprovalRuleTemplate", input)
    return &CodecommitListRepositoriesForApprovalRuleTemplateResult{Result: future}
}

func (a *CodeCommitStub) ListTagsForResource(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error) {
    var output codecommit.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) ListTagsForResourceAsync(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) *CodecommitListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.ListTagsForResource", input)
    return &CodecommitListTagsForResourceResult{Result: future}
}

func (a *CodeCommitStub) MergeBranchesByFastForward(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) (*codecommit.MergeBranchesByFastForwardOutput, error) {
    var output codecommit.MergeBranchesByFastForwardOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.MergeBranchesByFastForward", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) MergeBranchesByFastForwardAsync(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) *CodecommitMergeBranchesByFastForwardResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.MergeBranchesByFastForward", input)
    return &CodecommitMergeBranchesByFastForwardResult{Result: future}
}

func (a *CodeCommitStub) MergeBranchesBySquash(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) (*codecommit.MergeBranchesBySquashOutput, error) {
    var output codecommit.MergeBranchesBySquashOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.MergeBranchesBySquash", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) MergeBranchesBySquashAsync(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) *CodecommitMergeBranchesBySquashResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.MergeBranchesBySquash", input)
    return &CodecommitMergeBranchesBySquashResult{Result: future}
}

func (a *CodeCommitStub) MergeBranchesByThreeWay(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) (*codecommit.MergeBranchesByThreeWayOutput, error) {
    var output codecommit.MergeBranchesByThreeWayOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.MergeBranchesByThreeWay", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) MergeBranchesByThreeWayAsync(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) *CodecommitMergeBranchesByThreeWayResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.MergeBranchesByThreeWay", input)
    return &CodecommitMergeBranchesByThreeWayResult{Result: future}
}

func (a *CodeCommitStub) MergePullRequestByFastForward(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) (*codecommit.MergePullRequestByFastForwardOutput, error) {
    var output codecommit.MergePullRequestByFastForwardOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.MergePullRequestByFastForward", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) MergePullRequestByFastForwardAsync(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) *CodecommitMergePullRequestByFastForwardResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.MergePullRequestByFastForward", input)
    return &CodecommitMergePullRequestByFastForwardResult{Result: future}
}

func (a *CodeCommitStub) MergePullRequestBySquash(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) (*codecommit.MergePullRequestBySquashOutput, error) {
    var output codecommit.MergePullRequestBySquashOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.MergePullRequestBySquash", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) MergePullRequestBySquashAsync(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) *CodecommitMergePullRequestBySquashResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.MergePullRequestBySquash", input)
    return &CodecommitMergePullRequestBySquashResult{Result: future}
}

func (a *CodeCommitStub) MergePullRequestByThreeWay(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) (*codecommit.MergePullRequestByThreeWayOutput, error) {
    var output codecommit.MergePullRequestByThreeWayOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.MergePullRequestByThreeWay", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) MergePullRequestByThreeWayAsync(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) *CodecommitMergePullRequestByThreeWayResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.MergePullRequestByThreeWay", input)
    return &CodecommitMergePullRequestByThreeWayResult{Result: future}
}

func (a *CodeCommitStub) OverridePullRequestApprovalRules(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) (*codecommit.OverridePullRequestApprovalRulesOutput, error) {
    var output codecommit.OverridePullRequestApprovalRulesOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.OverridePullRequestApprovalRules", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) OverridePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) *CodecommitOverridePullRequestApprovalRulesResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.OverridePullRequestApprovalRules", input)
    return &CodecommitOverridePullRequestApprovalRulesResult{Result: future}
}

func (a *CodeCommitStub) PostCommentForComparedCommit(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) (*codecommit.PostCommentForComparedCommitOutput, error) {
    var output codecommit.PostCommentForComparedCommitOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.PostCommentForComparedCommit", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) PostCommentForComparedCommitAsync(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) *CodecommitPostCommentForComparedCommitResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.PostCommentForComparedCommit", input)
    return &CodecommitPostCommentForComparedCommitResult{Result: future}
}

func (a *CodeCommitStub) PostCommentReply(ctx workflow.Context, input *codecommit.PostCommentReplyInput) (*codecommit.PostCommentReplyOutput, error) {
    var output codecommit.PostCommentReplyOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.PostCommentReply", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) PostCommentReplyAsync(ctx workflow.Context, input *codecommit.PostCommentReplyInput) *CodecommitPostCommentReplyResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.PostCommentReply", input)
    return &CodecommitPostCommentReplyResult{Result: future}
}

func (a *CodeCommitStub) PutCommentReaction(ctx workflow.Context, input *codecommit.PutCommentReactionInput) (*codecommit.PutCommentReactionOutput, error) {
    var output codecommit.PutCommentReactionOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.PutCommentReaction", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) PutCommentReactionAsync(ctx workflow.Context, input *codecommit.PutCommentReactionInput) *CodecommitPutCommentReactionResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.PutCommentReaction", input)
    return &CodecommitPutCommentReactionResult{Result: future}
}

func (a *CodeCommitStub) PutFile(ctx workflow.Context, input *codecommit.PutFileInput) (*codecommit.PutFileOutput, error) {
    var output codecommit.PutFileOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.PutFile", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) PutFileAsync(ctx workflow.Context, input *codecommit.PutFileInput) *CodecommitPutFileResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.PutFile", input)
    return &CodecommitPutFileResult{Result: future}
}

func (a *CodeCommitStub) PutRepositoryTriggers(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error) {
    var output codecommit.PutRepositoryTriggersOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.PutRepositoryTriggers", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) PutRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) *CodecommitPutRepositoryTriggersResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.PutRepositoryTriggers", input)
    return &CodecommitPutRepositoryTriggersResult{Result: future}
}

func (a *CodeCommitStub) TagResource(ctx workflow.Context, input *codecommit.TagResourceInput) (*codecommit.TagResourceOutput, error) {
    var output codecommit.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) TagResourceAsync(ctx workflow.Context, input *codecommit.TagResourceInput) *CodecommitTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.TagResource", input)
    return &CodecommitTagResourceResult{Result: future}
}

func (a *CodeCommitStub) TestRepositoryTriggers(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error) {
    var output codecommit.TestRepositoryTriggersOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.TestRepositoryTriggers", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) TestRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) *CodecommitTestRepositoryTriggersResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.TestRepositoryTriggers", input)
    return &CodecommitTestRepositoryTriggersResult{Result: future}
}

func (a *CodeCommitStub) UntagResource(ctx workflow.Context, input *codecommit.UntagResourceInput) (*codecommit.UntagResourceOutput, error) {
    var output codecommit.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UntagResourceAsync(ctx workflow.Context, input *codecommit.UntagResourceInput) *CodecommitUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UntagResource", input)
    return &CodecommitUntagResourceResult{Result: future}
}

func (a *CodeCommitStub) UpdateApprovalRuleTemplateContent(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error) {
    var output codecommit.UpdateApprovalRuleTemplateContentOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateApprovalRuleTemplateContent", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateApprovalRuleTemplateContentAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) *CodecommitUpdateApprovalRuleTemplateContentResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateApprovalRuleTemplateContent", input)
    return &CodecommitUpdateApprovalRuleTemplateContentResult{Result: future}
}

func (a *CodeCommitStub) UpdateApprovalRuleTemplateDescription(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error) {
    var output codecommit.UpdateApprovalRuleTemplateDescriptionOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateApprovalRuleTemplateDescription", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateApprovalRuleTemplateDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) *CodecommitUpdateApprovalRuleTemplateDescriptionResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateApprovalRuleTemplateDescription", input)
    return &CodecommitUpdateApprovalRuleTemplateDescriptionResult{Result: future}
}

func (a *CodeCommitStub) UpdateApprovalRuleTemplateName(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error) {
    var output codecommit.UpdateApprovalRuleTemplateNameOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateApprovalRuleTemplateName", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateApprovalRuleTemplateNameAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) *CodecommitUpdateApprovalRuleTemplateNameResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateApprovalRuleTemplateName", input)
    return &CodecommitUpdateApprovalRuleTemplateNameResult{Result: future}
}

func (a *CodeCommitStub) UpdateComment(ctx workflow.Context, input *codecommit.UpdateCommentInput) (*codecommit.UpdateCommentOutput, error) {
    var output codecommit.UpdateCommentOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateComment", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateCommentAsync(ctx workflow.Context, input *codecommit.UpdateCommentInput) *CodecommitUpdateCommentResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateComment", input)
    return &CodecommitUpdateCommentResult{Result: future}
}

func (a *CodeCommitStub) UpdateDefaultBranch(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error) {
    var output codecommit.UpdateDefaultBranchOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateDefaultBranch", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateDefaultBranchAsync(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) *CodecommitUpdateDefaultBranchResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateDefaultBranch", input)
    return &CodecommitUpdateDefaultBranchResult{Result: future}
}

func (a *CodeCommitStub) UpdatePullRequestApprovalRuleContent(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error) {
    var output codecommit.UpdatePullRequestApprovalRuleContentOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestApprovalRuleContent", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdatePullRequestApprovalRuleContentAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) *CodecommitUpdatePullRequestApprovalRuleContentResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestApprovalRuleContent", input)
    return &CodecommitUpdatePullRequestApprovalRuleContentResult{Result: future}
}

func (a *CodeCommitStub) UpdatePullRequestApprovalState(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) (*codecommit.UpdatePullRequestApprovalStateOutput, error) {
    var output codecommit.UpdatePullRequestApprovalStateOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestApprovalState", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdatePullRequestApprovalStateAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) *CodecommitUpdatePullRequestApprovalStateResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestApprovalState", input)
    return &CodecommitUpdatePullRequestApprovalStateResult{Result: future}
}

func (a *CodeCommitStub) UpdatePullRequestDescription(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) (*codecommit.UpdatePullRequestDescriptionOutput, error) {
    var output codecommit.UpdatePullRequestDescriptionOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestDescription", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdatePullRequestDescriptionAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) *CodecommitUpdatePullRequestDescriptionResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestDescription", input)
    return &CodecommitUpdatePullRequestDescriptionResult{Result: future}
}

func (a *CodeCommitStub) UpdatePullRequestStatus(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) (*codecommit.UpdatePullRequestStatusOutput, error) {
    var output codecommit.UpdatePullRequestStatusOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdatePullRequestStatusAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) *CodecommitUpdatePullRequestStatusResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestStatus", input)
    return &CodecommitUpdatePullRequestStatusResult{Result: future}
}

func (a *CodeCommitStub) UpdatePullRequestTitle(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) (*codecommit.UpdatePullRequestTitleOutput, error) {
    var output codecommit.UpdatePullRequestTitleOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestTitle", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdatePullRequestTitleAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) *CodecommitUpdatePullRequestTitleResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdatePullRequestTitle", input)
    return &CodecommitUpdatePullRequestTitleResult{Result: future}
}

func (a *CodeCommitStub) UpdateRepositoryDescription(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error) {
    var output codecommit.UpdateRepositoryDescriptionOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateRepositoryDescription", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateRepositoryDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) *CodecommitUpdateRepositoryDescriptionResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateRepositoryDescription", input)
    return &CodecommitUpdateRepositoryDescriptionResult{Result: future}
}

func (a *CodeCommitStub) UpdateRepositoryName(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error) {
    var output codecommit.UpdateRepositoryNameOutput
    err := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateRepositoryName", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeCommitStub) UpdateRepositoryNameAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) *CodecommitUpdateRepositoryNameResult {
    future := workflow.ExecuteActivity(ctx, "CodeCommit.UpdateRepositoryName", input)
    return &CodecommitUpdateRepositoryNameResult{Result: future}
}
