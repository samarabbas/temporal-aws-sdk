package awsclients

import (
	"github.com/aws/aws-sdk-go/service/workdocs"
	"go.temporal.io/sdk/workflow"
)

type WorkDocsClient interface {
       AbortDocumentVersionUpload(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) (*workdocs.AbortDocumentVersionUploadOutput, error)
       AbortDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) *WorkdocsAbortDocumentVersionUploadResult

       ActivateUser(ctx workflow.Context, input *workdocs.ActivateUserInput) (*workdocs.ActivateUserOutput, error)
       ActivateUserAsync(ctx workflow.Context, input *workdocs.ActivateUserInput) *WorkdocsActivateUserResult

       AddResourcePermissions(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) (*workdocs.AddResourcePermissionsOutput, error)
       AddResourcePermissionsAsync(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) *WorkdocsAddResourcePermissionsResult

       CreateComment(ctx workflow.Context, input *workdocs.CreateCommentInput) (*workdocs.CreateCommentOutput, error)
       CreateCommentAsync(ctx workflow.Context, input *workdocs.CreateCommentInput) *WorkdocsCreateCommentResult

       CreateCustomMetadata(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) (*workdocs.CreateCustomMetadataOutput, error)
       CreateCustomMetadataAsync(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) *WorkdocsCreateCustomMetadataResult

       CreateFolder(ctx workflow.Context, input *workdocs.CreateFolderInput) (*workdocs.CreateFolderOutput, error)
       CreateFolderAsync(ctx workflow.Context, input *workdocs.CreateFolderInput) *WorkdocsCreateFolderResult

       CreateLabels(ctx workflow.Context, input *workdocs.CreateLabelsInput) (*workdocs.CreateLabelsOutput, error)
       CreateLabelsAsync(ctx workflow.Context, input *workdocs.CreateLabelsInput) *WorkdocsCreateLabelsResult

       CreateNotificationSubscription(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) (*workdocs.CreateNotificationSubscriptionOutput, error)
       CreateNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) *WorkdocsCreateNotificationSubscriptionResult

       CreateUser(ctx workflow.Context, input *workdocs.CreateUserInput) (*workdocs.CreateUserOutput, error)
       CreateUserAsync(ctx workflow.Context, input *workdocs.CreateUserInput) *WorkdocsCreateUserResult

       DeactivateUser(ctx workflow.Context, input *workdocs.DeactivateUserInput) (*workdocs.DeactivateUserOutput, error)
       DeactivateUserAsync(ctx workflow.Context, input *workdocs.DeactivateUserInput) *WorkdocsDeactivateUserResult

       DeleteComment(ctx workflow.Context, input *workdocs.DeleteCommentInput) (*workdocs.DeleteCommentOutput, error)
       DeleteCommentAsync(ctx workflow.Context, input *workdocs.DeleteCommentInput) *WorkdocsDeleteCommentResult

       DeleteCustomMetadata(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) (*workdocs.DeleteCustomMetadataOutput, error)
       DeleteCustomMetadataAsync(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) *WorkdocsDeleteCustomMetadataResult

       DeleteDocument(ctx workflow.Context, input *workdocs.DeleteDocumentInput) (*workdocs.DeleteDocumentOutput, error)
       DeleteDocumentAsync(ctx workflow.Context, input *workdocs.DeleteDocumentInput) *WorkdocsDeleteDocumentResult

       DeleteFolder(ctx workflow.Context, input *workdocs.DeleteFolderInput) (*workdocs.DeleteFolderOutput, error)
       DeleteFolderAsync(ctx workflow.Context, input *workdocs.DeleteFolderInput) *WorkdocsDeleteFolderResult

       DeleteFolderContents(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) (*workdocs.DeleteFolderContentsOutput, error)
       DeleteFolderContentsAsync(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) *WorkdocsDeleteFolderContentsResult

       DeleteLabels(ctx workflow.Context, input *workdocs.DeleteLabelsInput) (*workdocs.DeleteLabelsOutput, error)
       DeleteLabelsAsync(ctx workflow.Context, input *workdocs.DeleteLabelsInput) *WorkdocsDeleteLabelsResult

       DeleteNotificationSubscription(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) (*workdocs.DeleteNotificationSubscriptionOutput, error)
       DeleteNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) *WorkdocsDeleteNotificationSubscriptionResult

       DeleteUser(ctx workflow.Context, input *workdocs.DeleteUserInput) (*workdocs.DeleteUserOutput, error)
       DeleteUserAsync(ctx workflow.Context, input *workdocs.DeleteUserInput) *WorkdocsDeleteUserResult

       DescribeActivities(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error)
       DescribeActivitiesAsync(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) *WorkdocsDescribeActivitiesResult

       DescribeComments(ctx workflow.Context, input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error)
       DescribeCommentsAsync(ctx workflow.Context, input *workdocs.DescribeCommentsInput) *WorkdocsDescribeCommentsResult

       DescribeDocumentVersions(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error)
       DescribeDocumentVersionsAsync(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) *WorkdocsDescribeDocumentVersionsResult

       DescribeFolderContents(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error)
       DescribeFolderContentsAsync(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) *WorkdocsDescribeFolderContentsResult

       DescribeGroups(ctx workflow.Context, input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error)
       DescribeGroupsAsync(ctx workflow.Context, input *workdocs.DescribeGroupsInput) *WorkdocsDescribeGroupsResult

       DescribeNotificationSubscriptions(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error)
       DescribeNotificationSubscriptionsAsync(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) *WorkdocsDescribeNotificationSubscriptionsResult

       DescribeResourcePermissions(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error)
       DescribeResourcePermissionsAsync(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) *WorkdocsDescribeResourcePermissionsResult

       DescribeRootFolders(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error)
       DescribeRootFoldersAsync(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) *WorkdocsDescribeRootFoldersResult

       DescribeUsers(ctx workflow.Context, input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error)
       DescribeUsersAsync(ctx workflow.Context, input *workdocs.DescribeUsersInput) *WorkdocsDescribeUsersResult

       GetCurrentUser(ctx workflow.Context, input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error)
       GetCurrentUserAsync(ctx workflow.Context, input *workdocs.GetCurrentUserInput) *WorkdocsGetCurrentUserResult

       GetDocument(ctx workflow.Context, input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error)
       GetDocumentAsync(ctx workflow.Context, input *workdocs.GetDocumentInput) *WorkdocsGetDocumentResult

       GetDocumentPath(ctx workflow.Context, input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error)
       GetDocumentPathAsync(ctx workflow.Context, input *workdocs.GetDocumentPathInput) *WorkdocsGetDocumentPathResult

       GetDocumentVersion(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error)
       GetDocumentVersionAsync(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) *WorkdocsGetDocumentVersionResult

       GetFolder(ctx workflow.Context, input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error)
       GetFolderAsync(ctx workflow.Context, input *workdocs.GetFolderInput) *WorkdocsGetFolderResult

       GetFolderPath(ctx workflow.Context, input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error)
       GetFolderPathAsync(ctx workflow.Context, input *workdocs.GetFolderPathInput) *WorkdocsGetFolderPathResult

       GetResources(ctx workflow.Context, input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error)
       GetResourcesAsync(ctx workflow.Context, input *workdocs.GetResourcesInput) *WorkdocsGetResourcesResult

       InitiateDocumentVersionUpload(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) (*workdocs.InitiateDocumentVersionUploadOutput, error)
       InitiateDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) *WorkdocsInitiateDocumentVersionUploadResult

       RemoveAllResourcePermissions(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) (*workdocs.RemoveAllResourcePermissionsOutput, error)
       RemoveAllResourcePermissionsAsync(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) *WorkdocsRemoveAllResourcePermissionsResult

       RemoveResourcePermission(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) (*workdocs.RemoveResourcePermissionOutput, error)
       RemoveResourcePermissionAsync(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) *WorkdocsRemoveResourcePermissionResult

       UpdateDocument(ctx workflow.Context, input *workdocs.UpdateDocumentInput) (*workdocs.UpdateDocumentOutput, error)
       UpdateDocumentAsync(ctx workflow.Context, input *workdocs.UpdateDocumentInput) *WorkdocsUpdateDocumentResult

       UpdateDocumentVersion(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) (*workdocs.UpdateDocumentVersionOutput, error)
       UpdateDocumentVersionAsync(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) *WorkdocsUpdateDocumentVersionResult

       UpdateFolder(ctx workflow.Context, input *workdocs.UpdateFolderInput) (*workdocs.UpdateFolderOutput, error)
       UpdateFolderAsync(ctx workflow.Context, input *workdocs.UpdateFolderInput) *WorkdocsUpdateFolderResult

       UpdateUser(ctx workflow.Context, input *workdocs.UpdateUserInput) (*workdocs.UpdateUserOutput, error)
       UpdateUserAsync(ctx workflow.Context, input *workdocs.UpdateUserInput) *WorkdocsUpdateUserResult
}

type WorkDocsStub struct {
}

func NewWorkDocsStub() WorkDocsClient {
    return &WorkDocsStub{}
}

type WorkdocsAbortDocumentVersionUploadResult struct {
	Result workflow.Future
}

func (r *WorkdocsAbortDocumentVersionUploadResult) Get(ctx workflow.Context) (*workdocs.AbortDocumentVersionUploadOutput, error) {
    var output workdocs.AbortDocumentVersionUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsActivateUserResult struct {
	Result workflow.Future
}

func (r *WorkdocsActivateUserResult) Get(ctx workflow.Context) (*workdocs.ActivateUserOutput, error) {
    var output workdocs.ActivateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsAddResourcePermissionsResult struct {
	Result workflow.Future
}

func (r *WorkdocsAddResourcePermissionsResult) Get(ctx workflow.Context) (*workdocs.AddResourcePermissionsOutput, error) {
    var output workdocs.AddResourcePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsCreateCommentResult struct {
	Result workflow.Future
}

func (r *WorkdocsCreateCommentResult) Get(ctx workflow.Context) (*workdocs.CreateCommentOutput, error) {
    var output workdocs.CreateCommentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsCreateCustomMetadataResult struct {
	Result workflow.Future
}

func (r *WorkdocsCreateCustomMetadataResult) Get(ctx workflow.Context) (*workdocs.CreateCustomMetadataOutput, error) {
    var output workdocs.CreateCustomMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsCreateFolderResult struct {
	Result workflow.Future
}

func (r *WorkdocsCreateFolderResult) Get(ctx workflow.Context) (*workdocs.CreateFolderOutput, error) {
    var output workdocs.CreateFolderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsCreateLabelsResult struct {
	Result workflow.Future
}

func (r *WorkdocsCreateLabelsResult) Get(ctx workflow.Context) (*workdocs.CreateLabelsOutput, error) {
    var output workdocs.CreateLabelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsCreateNotificationSubscriptionResult struct {
	Result workflow.Future
}

func (r *WorkdocsCreateNotificationSubscriptionResult) Get(ctx workflow.Context) (*workdocs.CreateNotificationSubscriptionOutput, error) {
    var output workdocs.CreateNotificationSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsCreateUserResult struct {
	Result workflow.Future
}

func (r *WorkdocsCreateUserResult) Get(ctx workflow.Context) (*workdocs.CreateUserOutput, error) {
    var output workdocs.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeactivateUserResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeactivateUserResult) Get(ctx workflow.Context) (*workdocs.DeactivateUserOutput, error) {
    var output workdocs.DeactivateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteCommentResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteCommentResult) Get(ctx workflow.Context) (*workdocs.DeleteCommentOutput, error) {
    var output workdocs.DeleteCommentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteCustomMetadataResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteCustomMetadataResult) Get(ctx workflow.Context) (*workdocs.DeleteCustomMetadataOutput, error) {
    var output workdocs.DeleteCustomMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteDocumentResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteDocumentResult) Get(ctx workflow.Context) (*workdocs.DeleteDocumentOutput, error) {
    var output workdocs.DeleteDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteFolderResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteFolderResult) Get(ctx workflow.Context) (*workdocs.DeleteFolderOutput, error) {
    var output workdocs.DeleteFolderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteFolderContentsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteFolderContentsResult) Get(ctx workflow.Context) (*workdocs.DeleteFolderContentsOutput, error) {
    var output workdocs.DeleteFolderContentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteLabelsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteLabelsResult) Get(ctx workflow.Context) (*workdocs.DeleteLabelsOutput, error) {
    var output workdocs.DeleteLabelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteNotificationSubscriptionResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteNotificationSubscriptionResult) Get(ctx workflow.Context) (*workdocs.DeleteNotificationSubscriptionOutput, error) {
    var output workdocs.DeleteNotificationSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDeleteUserResult struct {
	Result workflow.Future
}

func (r *WorkdocsDeleteUserResult) Get(ctx workflow.Context) (*workdocs.DeleteUserOutput, error) {
    var output workdocs.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeActivitiesResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeActivitiesResult) Get(ctx workflow.Context) (*workdocs.DescribeActivitiesOutput, error) {
    var output workdocs.DescribeActivitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeCommentsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeCommentsResult) Get(ctx workflow.Context) (*workdocs.DescribeCommentsOutput, error) {
    var output workdocs.DescribeCommentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeDocumentVersionsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeDocumentVersionsResult) Get(ctx workflow.Context) (*workdocs.DescribeDocumentVersionsOutput, error) {
    var output workdocs.DescribeDocumentVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeFolderContentsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeFolderContentsResult) Get(ctx workflow.Context) (*workdocs.DescribeFolderContentsOutput, error) {
    var output workdocs.DescribeFolderContentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeGroupsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeGroupsResult) Get(ctx workflow.Context) (*workdocs.DescribeGroupsOutput, error) {
    var output workdocs.DescribeGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeNotificationSubscriptionsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeNotificationSubscriptionsResult) Get(ctx workflow.Context) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
    var output workdocs.DescribeNotificationSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeResourcePermissionsResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeResourcePermissionsResult) Get(ctx workflow.Context) (*workdocs.DescribeResourcePermissionsOutput, error) {
    var output workdocs.DescribeResourcePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeRootFoldersResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeRootFoldersResult) Get(ctx workflow.Context) (*workdocs.DescribeRootFoldersOutput, error) {
    var output workdocs.DescribeRootFoldersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsDescribeUsersResult struct {
	Result workflow.Future
}

func (r *WorkdocsDescribeUsersResult) Get(ctx workflow.Context) (*workdocs.DescribeUsersOutput, error) {
    var output workdocs.DescribeUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetCurrentUserResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetCurrentUserResult) Get(ctx workflow.Context) (*workdocs.GetCurrentUserOutput, error) {
    var output workdocs.GetCurrentUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetDocumentResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetDocumentResult) Get(ctx workflow.Context) (*workdocs.GetDocumentOutput, error) {
    var output workdocs.GetDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetDocumentPathResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetDocumentPathResult) Get(ctx workflow.Context) (*workdocs.GetDocumentPathOutput, error) {
    var output workdocs.GetDocumentPathOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetDocumentVersionResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetDocumentVersionResult) Get(ctx workflow.Context) (*workdocs.GetDocumentVersionOutput, error) {
    var output workdocs.GetDocumentVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetFolderResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetFolderResult) Get(ctx workflow.Context) (*workdocs.GetFolderOutput, error) {
    var output workdocs.GetFolderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetFolderPathResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetFolderPathResult) Get(ctx workflow.Context) (*workdocs.GetFolderPathOutput, error) {
    var output workdocs.GetFolderPathOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsGetResourcesResult struct {
	Result workflow.Future
}

func (r *WorkdocsGetResourcesResult) Get(ctx workflow.Context) (*workdocs.GetResourcesOutput, error) {
    var output workdocs.GetResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsInitiateDocumentVersionUploadResult struct {
	Result workflow.Future
}

func (r *WorkdocsInitiateDocumentVersionUploadResult) Get(ctx workflow.Context) (*workdocs.InitiateDocumentVersionUploadOutput, error) {
    var output workdocs.InitiateDocumentVersionUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsRemoveAllResourcePermissionsResult struct {
	Result workflow.Future
}

func (r *WorkdocsRemoveAllResourcePermissionsResult) Get(ctx workflow.Context) (*workdocs.RemoveAllResourcePermissionsOutput, error) {
    var output workdocs.RemoveAllResourcePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsRemoveResourcePermissionResult struct {
	Result workflow.Future
}

func (r *WorkdocsRemoveResourcePermissionResult) Get(ctx workflow.Context) (*workdocs.RemoveResourcePermissionOutput, error) {
    var output workdocs.RemoveResourcePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsUpdateDocumentResult struct {
	Result workflow.Future
}

func (r *WorkdocsUpdateDocumentResult) Get(ctx workflow.Context) (*workdocs.UpdateDocumentOutput, error) {
    var output workdocs.UpdateDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsUpdateDocumentVersionResult struct {
	Result workflow.Future
}

func (r *WorkdocsUpdateDocumentVersionResult) Get(ctx workflow.Context) (*workdocs.UpdateDocumentVersionOutput, error) {
    var output workdocs.UpdateDocumentVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsUpdateFolderResult struct {
	Result workflow.Future
}

func (r *WorkdocsUpdateFolderResult) Get(ctx workflow.Context) (*workdocs.UpdateFolderOutput, error) {
    var output workdocs.UpdateFolderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type WorkdocsUpdateUserResult struct {
	Result workflow.Future
}

func (r *WorkdocsUpdateUserResult) Get(ctx workflow.Context) (*workdocs.UpdateUserOutput, error) {
    var output workdocs.UpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) AbortDocumentVersionUpload(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) (*workdocs.AbortDocumentVersionUploadOutput, error) {
    var output workdocs.AbortDocumentVersionUploadOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.AbortDocumentVersionUpload", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) AbortDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) *WorkdocsAbortDocumentVersionUploadResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.AbortDocumentVersionUpload", input)
    return &WorkdocsAbortDocumentVersionUploadResult{Result: future}
}

func (a *WorkDocsStub) ActivateUser(ctx workflow.Context, input *workdocs.ActivateUserInput) (*workdocs.ActivateUserOutput, error) {
    var output workdocs.ActivateUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.ActivateUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) ActivateUserAsync(ctx workflow.Context, input *workdocs.ActivateUserInput) *WorkdocsActivateUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.ActivateUser", input)
    return &WorkdocsActivateUserResult{Result: future}
}

func (a *WorkDocsStub) AddResourcePermissions(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) (*workdocs.AddResourcePermissionsOutput, error) {
    var output workdocs.AddResourcePermissionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.AddResourcePermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) AddResourcePermissionsAsync(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) *WorkdocsAddResourcePermissionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.AddResourcePermissions", input)
    return &WorkdocsAddResourcePermissionsResult{Result: future}
}

func (a *WorkDocsStub) CreateComment(ctx workflow.Context, input *workdocs.CreateCommentInput) (*workdocs.CreateCommentOutput, error) {
    var output workdocs.CreateCommentOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.CreateComment", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) CreateCommentAsync(ctx workflow.Context, input *workdocs.CreateCommentInput) *WorkdocsCreateCommentResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.CreateComment", input)
    return &WorkdocsCreateCommentResult{Result: future}
}

func (a *WorkDocsStub) CreateCustomMetadata(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) (*workdocs.CreateCustomMetadataOutput, error) {
    var output workdocs.CreateCustomMetadataOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.CreateCustomMetadata", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) CreateCustomMetadataAsync(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) *WorkdocsCreateCustomMetadataResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.CreateCustomMetadata", input)
    return &WorkdocsCreateCustomMetadataResult{Result: future}
}

func (a *WorkDocsStub) CreateFolder(ctx workflow.Context, input *workdocs.CreateFolderInput) (*workdocs.CreateFolderOutput, error) {
    var output workdocs.CreateFolderOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.CreateFolder", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) CreateFolderAsync(ctx workflow.Context, input *workdocs.CreateFolderInput) *WorkdocsCreateFolderResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.CreateFolder", input)
    return &WorkdocsCreateFolderResult{Result: future}
}

func (a *WorkDocsStub) CreateLabels(ctx workflow.Context, input *workdocs.CreateLabelsInput) (*workdocs.CreateLabelsOutput, error) {
    var output workdocs.CreateLabelsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.CreateLabels", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) CreateLabelsAsync(ctx workflow.Context, input *workdocs.CreateLabelsInput) *WorkdocsCreateLabelsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.CreateLabels", input)
    return &WorkdocsCreateLabelsResult{Result: future}
}

func (a *WorkDocsStub) CreateNotificationSubscription(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) (*workdocs.CreateNotificationSubscriptionOutput, error) {
    var output workdocs.CreateNotificationSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.CreateNotificationSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) CreateNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) *WorkdocsCreateNotificationSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.CreateNotificationSubscription", input)
    return &WorkdocsCreateNotificationSubscriptionResult{Result: future}
}

func (a *WorkDocsStub) CreateUser(ctx workflow.Context, input *workdocs.CreateUserInput) (*workdocs.CreateUserOutput, error) {
    var output workdocs.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.CreateUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) CreateUserAsync(ctx workflow.Context, input *workdocs.CreateUserInput) *WorkdocsCreateUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.CreateUser", input)
    return &WorkdocsCreateUserResult{Result: future}
}

func (a *WorkDocsStub) DeactivateUser(ctx workflow.Context, input *workdocs.DeactivateUserInput) (*workdocs.DeactivateUserOutput, error) {
    var output workdocs.DeactivateUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeactivateUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeactivateUserAsync(ctx workflow.Context, input *workdocs.DeactivateUserInput) *WorkdocsDeactivateUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeactivateUser", input)
    return &WorkdocsDeactivateUserResult{Result: future}
}

func (a *WorkDocsStub) DeleteComment(ctx workflow.Context, input *workdocs.DeleteCommentInput) (*workdocs.DeleteCommentOutput, error) {
    var output workdocs.DeleteCommentOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteComment", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteCommentAsync(ctx workflow.Context, input *workdocs.DeleteCommentInput) *WorkdocsDeleteCommentResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteComment", input)
    return &WorkdocsDeleteCommentResult{Result: future}
}

func (a *WorkDocsStub) DeleteCustomMetadata(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) (*workdocs.DeleteCustomMetadataOutput, error) {
    var output workdocs.DeleteCustomMetadataOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteCustomMetadata", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteCustomMetadataAsync(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) *WorkdocsDeleteCustomMetadataResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteCustomMetadata", input)
    return &WorkdocsDeleteCustomMetadataResult{Result: future}
}

func (a *WorkDocsStub) DeleteDocument(ctx workflow.Context, input *workdocs.DeleteDocumentInput) (*workdocs.DeleteDocumentOutput, error) {
    var output workdocs.DeleteDocumentOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteDocument", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteDocumentAsync(ctx workflow.Context, input *workdocs.DeleteDocumentInput) *WorkdocsDeleteDocumentResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteDocument", input)
    return &WorkdocsDeleteDocumentResult{Result: future}
}

func (a *WorkDocsStub) DeleteFolder(ctx workflow.Context, input *workdocs.DeleteFolderInput) (*workdocs.DeleteFolderOutput, error) {
    var output workdocs.DeleteFolderOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteFolder", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteFolderAsync(ctx workflow.Context, input *workdocs.DeleteFolderInput) *WorkdocsDeleteFolderResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteFolder", input)
    return &WorkdocsDeleteFolderResult{Result: future}
}

func (a *WorkDocsStub) DeleteFolderContents(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) (*workdocs.DeleteFolderContentsOutput, error) {
    var output workdocs.DeleteFolderContentsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteFolderContents", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteFolderContentsAsync(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) *WorkdocsDeleteFolderContentsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteFolderContents", input)
    return &WorkdocsDeleteFolderContentsResult{Result: future}
}

func (a *WorkDocsStub) DeleteLabels(ctx workflow.Context, input *workdocs.DeleteLabelsInput) (*workdocs.DeleteLabelsOutput, error) {
    var output workdocs.DeleteLabelsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteLabels", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteLabelsAsync(ctx workflow.Context, input *workdocs.DeleteLabelsInput) *WorkdocsDeleteLabelsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteLabels", input)
    return &WorkdocsDeleteLabelsResult{Result: future}
}

func (a *WorkDocsStub) DeleteNotificationSubscription(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) (*workdocs.DeleteNotificationSubscriptionOutput, error) {
    var output workdocs.DeleteNotificationSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteNotificationSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) *WorkdocsDeleteNotificationSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteNotificationSubscription", input)
    return &WorkdocsDeleteNotificationSubscriptionResult{Result: future}
}

func (a *WorkDocsStub) DeleteUser(ctx workflow.Context, input *workdocs.DeleteUserInput) (*workdocs.DeleteUserOutput, error) {
    var output workdocs.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DeleteUserAsync(ctx workflow.Context, input *workdocs.DeleteUserInput) *WorkdocsDeleteUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DeleteUser", input)
    return &WorkdocsDeleteUserResult{Result: future}
}

func (a *WorkDocsStub) DescribeActivities(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error) {
    var output workdocs.DescribeActivitiesOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeActivities", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeActivitiesAsync(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) *WorkdocsDescribeActivitiesResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeActivities", input)
    return &WorkdocsDescribeActivitiesResult{Result: future}
}

func (a *WorkDocsStub) DescribeComments(ctx workflow.Context, input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error) {
    var output workdocs.DescribeCommentsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeComments", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeCommentsAsync(ctx workflow.Context, input *workdocs.DescribeCommentsInput) *WorkdocsDescribeCommentsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeComments", input)
    return &WorkdocsDescribeCommentsResult{Result: future}
}

func (a *WorkDocsStub) DescribeDocumentVersions(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error) {
    var output workdocs.DescribeDocumentVersionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeDocumentVersions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeDocumentVersionsAsync(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) *WorkdocsDescribeDocumentVersionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeDocumentVersions", input)
    return &WorkdocsDescribeDocumentVersionsResult{Result: future}
}

func (a *WorkDocsStub) DescribeFolderContents(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error) {
    var output workdocs.DescribeFolderContentsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeFolderContents", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeFolderContentsAsync(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) *WorkdocsDescribeFolderContentsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeFolderContents", input)
    return &WorkdocsDescribeFolderContentsResult{Result: future}
}

func (a *WorkDocsStub) DescribeGroups(ctx workflow.Context, input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error) {
    var output workdocs.DescribeGroupsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeGroupsAsync(ctx workflow.Context, input *workdocs.DescribeGroupsInput) *WorkdocsDescribeGroupsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeGroups", input)
    return &WorkdocsDescribeGroupsResult{Result: future}
}

func (a *WorkDocsStub) DescribeNotificationSubscriptions(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
    var output workdocs.DescribeNotificationSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeNotificationSubscriptions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeNotificationSubscriptionsAsync(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) *WorkdocsDescribeNotificationSubscriptionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeNotificationSubscriptions", input)
    return &WorkdocsDescribeNotificationSubscriptionsResult{Result: future}
}

func (a *WorkDocsStub) DescribeResourcePermissions(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error) {
    var output workdocs.DescribeResourcePermissionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeResourcePermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeResourcePermissionsAsync(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) *WorkdocsDescribeResourcePermissionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeResourcePermissions", input)
    return &WorkdocsDescribeResourcePermissionsResult{Result: future}
}

func (a *WorkDocsStub) DescribeRootFolders(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error) {
    var output workdocs.DescribeRootFoldersOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeRootFolders", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeRootFoldersAsync(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) *WorkdocsDescribeRootFoldersResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeRootFolders", input)
    return &WorkdocsDescribeRootFoldersResult{Result: future}
}

func (a *WorkDocsStub) DescribeUsers(ctx workflow.Context, input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error) {
    var output workdocs.DescribeUsersOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeUsers", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) DescribeUsersAsync(ctx workflow.Context, input *workdocs.DescribeUsersInput) *WorkdocsDescribeUsersResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.DescribeUsers", input)
    return &WorkdocsDescribeUsersResult{Result: future}
}

func (a *WorkDocsStub) GetCurrentUser(ctx workflow.Context, input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error) {
    var output workdocs.GetCurrentUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetCurrentUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetCurrentUserAsync(ctx workflow.Context, input *workdocs.GetCurrentUserInput) *WorkdocsGetCurrentUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetCurrentUser", input)
    return &WorkdocsGetCurrentUserResult{Result: future}
}

func (a *WorkDocsStub) GetDocument(ctx workflow.Context, input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error) {
    var output workdocs.GetDocumentOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetDocument", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetDocumentAsync(ctx workflow.Context, input *workdocs.GetDocumentInput) *WorkdocsGetDocumentResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetDocument", input)
    return &WorkdocsGetDocumentResult{Result: future}
}

func (a *WorkDocsStub) GetDocumentPath(ctx workflow.Context, input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error) {
    var output workdocs.GetDocumentPathOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetDocumentPath", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetDocumentPathAsync(ctx workflow.Context, input *workdocs.GetDocumentPathInput) *WorkdocsGetDocumentPathResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetDocumentPath", input)
    return &WorkdocsGetDocumentPathResult{Result: future}
}

func (a *WorkDocsStub) GetDocumentVersion(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error) {
    var output workdocs.GetDocumentVersionOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetDocumentVersion", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetDocumentVersionAsync(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) *WorkdocsGetDocumentVersionResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetDocumentVersion", input)
    return &WorkdocsGetDocumentVersionResult{Result: future}
}

func (a *WorkDocsStub) GetFolder(ctx workflow.Context, input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error) {
    var output workdocs.GetFolderOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetFolder", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetFolderAsync(ctx workflow.Context, input *workdocs.GetFolderInput) *WorkdocsGetFolderResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetFolder", input)
    return &WorkdocsGetFolderResult{Result: future}
}

func (a *WorkDocsStub) GetFolderPath(ctx workflow.Context, input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error) {
    var output workdocs.GetFolderPathOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetFolderPath", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetFolderPathAsync(ctx workflow.Context, input *workdocs.GetFolderPathInput) *WorkdocsGetFolderPathResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetFolderPath", input)
    return &WorkdocsGetFolderPathResult{Result: future}
}

func (a *WorkDocsStub) GetResources(ctx workflow.Context, input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error) {
    var output workdocs.GetResourcesOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.GetResources", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) GetResourcesAsync(ctx workflow.Context, input *workdocs.GetResourcesInput) *WorkdocsGetResourcesResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.GetResources", input)
    return &WorkdocsGetResourcesResult{Result: future}
}

func (a *WorkDocsStub) InitiateDocumentVersionUpload(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) (*workdocs.InitiateDocumentVersionUploadOutput, error) {
    var output workdocs.InitiateDocumentVersionUploadOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.InitiateDocumentVersionUpload", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) InitiateDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) *WorkdocsInitiateDocumentVersionUploadResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.InitiateDocumentVersionUpload", input)
    return &WorkdocsInitiateDocumentVersionUploadResult{Result: future}
}

func (a *WorkDocsStub) RemoveAllResourcePermissions(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) (*workdocs.RemoveAllResourcePermissionsOutput, error) {
    var output workdocs.RemoveAllResourcePermissionsOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.RemoveAllResourcePermissions", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) RemoveAllResourcePermissionsAsync(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) *WorkdocsRemoveAllResourcePermissionsResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.RemoveAllResourcePermissions", input)
    return &WorkdocsRemoveAllResourcePermissionsResult{Result: future}
}

func (a *WorkDocsStub) RemoveResourcePermission(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) (*workdocs.RemoveResourcePermissionOutput, error) {
    var output workdocs.RemoveResourcePermissionOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.RemoveResourcePermission", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) RemoveResourcePermissionAsync(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) *WorkdocsRemoveResourcePermissionResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.RemoveResourcePermission", input)
    return &WorkdocsRemoveResourcePermissionResult{Result: future}
}

func (a *WorkDocsStub) UpdateDocument(ctx workflow.Context, input *workdocs.UpdateDocumentInput) (*workdocs.UpdateDocumentOutput, error) {
    var output workdocs.UpdateDocumentOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateDocument", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) UpdateDocumentAsync(ctx workflow.Context, input *workdocs.UpdateDocumentInput) *WorkdocsUpdateDocumentResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateDocument", input)
    return &WorkdocsUpdateDocumentResult{Result: future}
}

func (a *WorkDocsStub) UpdateDocumentVersion(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) (*workdocs.UpdateDocumentVersionOutput, error) {
    var output workdocs.UpdateDocumentVersionOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateDocumentVersion", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) UpdateDocumentVersionAsync(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) *WorkdocsUpdateDocumentVersionResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateDocumentVersion", input)
    return &WorkdocsUpdateDocumentVersionResult{Result: future}
}

func (a *WorkDocsStub) UpdateFolder(ctx workflow.Context, input *workdocs.UpdateFolderInput) (*workdocs.UpdateFolderOutput, error) {
    var output workdocs.UpdateFolderOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateFolder", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) UpdateFolderAsync(ctx workflow.Context, input *workdocs.UpdateFolderInput) *WorkdocsUpdateFolderResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateFolder", input)
    return &WorkdocsUpdateFolderResult{Result: future}
}

func (a *WorkDocsStub) UpdateUser(ctx workflow.Context, input *workdocs.UpdateUserInput) (*workdocs.UpdateUserOutput, error) {
    var output workdocs.UpdateUserOutput
    err := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateUser", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkDocsStub) UpdateUserAsync(ctx workflow.Context, input *workdocs.UpdateUserInput) *WorkdocsUpdateUserResult {
    future := workflow.ExecuteActivity(ctx, "WorkDocs.UpdateUser", input)
    return &WorkdocsUpdateUserResult{Result: future}
}
