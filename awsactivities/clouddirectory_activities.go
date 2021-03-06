// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/aws/aws-sdk-go/service/clouddirectory/clouddirectoryiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CloudDirectoryActivities struct {
	client clouddirectoryiface.CloudDirectoryAPI
}

func NewCloudDirectoryActivities(session *session.Session, config ...*aws.Config) *CloudDirectoryActivities {
	client := clouddirectory.New(session, config...)
	return &CloudDirectoryActivities{client: client}
}

func (a *CloudDirectoryActivities) AddFacetToObject(ctx context.Context, input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error) {
	return a.client.AddFacetToObjectWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ApplySchema(ctx context.Context, input *clouddirectory.ApplySchemaInput) (*clouddirectory.ApplySchemaOutput, error) {
	return a.client.ApplySchemaWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) AttachObject(ctx context.Context, input *clouddirectory.AttachObjectInput) (*clouddirectory.AttachObjectOutput, error) {
	return a.client.AttachObjectWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) AttachPolicy(ctx context.Context, input *clouddirectory.AttachPolicyInput) (*clouddirectory.AttachPolicyOutput, error) {
	return a.client.AttachPolicyWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) AttachToIndex(ctx context.Context, input *clouddirectory.AttachToIndexInput) (*clouddirectory.AttachToIndexOutput, error) {
	return a.client.AttachToIndexWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) AttachTypedLink(ctx context.Context, input *clouddirectory.AttachTypedLinkInput) (*clouddirectory.AttachTypedLinkOutput, error) {
	return a.client.AttachTypedLinkWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) BatchRead(ctx context.Context, input *clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error) {
	return a.client.BatchReadWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) BatchWrite(ctx context.Context, input *clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error) {
	return a.client.BatchWriteWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) CreateDirectory(ctx context.Context, input *clouddirectory.CreateDirectoryInput) (*clouddirectory.CreateDirectoryOutput, error) {
	return a.client.CreateDirectoryWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) CreateFacet(ctx context.Context, input *clouddirectory.CreateFacetInput) (*clouddirectory.CreateFacetOutput, error) {
	return a.client.CreateFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) CreateIndex(ctx context.Context, input *clouddirectory.CreateIndexInput) (*clouddirectory.CreateIndexOutput, error) {
	return a.client.CreateIndexWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) CreateObject(ctx context.Context, input *clouddirectory.CreateObjectInput) (*clouddirectory.CreateObjectOutput, error) {
	return a.client.CreateObjectWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) CreateSchema(ctx context.Context, input *clouddirectory.CreateSchemaInput) (*clouddirectory.CreateSchemaOutput, error) {
	return a.client.CreateSchemaWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) CreateTypedLinkFacet(ctx context.Context, input *clouddirectory.CreateTypedLinkFacetInput) (*clouddirectory.CreateTypedLinkFacetOutput, error) {
	return a.client.CreateTypedLinkFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DeleteDirectory(ctx context.Context, input *clouddirectory.DeleteDirectoryInput) (*clouddirectory.DeleteDirectoryOutput, error) {
	return a.client.DeleteDirectoryWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DeleteFacet(ctx context.Context, input *clouddirectory.DeleteFacetInput) (*clouddirectory.DeleteFacetOutput, error) {
	return a.client.DeleteFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DeleteObject(ctx context.Context, input *clouddirectory.DeleteObjectInput) (*clouddirectory.DeleteObjectOutput, error) {
	return a.client.DeleteObjectWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DeleteSchema(ctx context.Context, input *clouddirectory.DeleteSchemaInput) (*clouddirectory.DeleteSchemaOutput, error) {
	return a.client.DeleteSchemaWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DeleteTypedLinkFacet(ctx context.Context, input *clouddirectory.DeleteTypedLinkFacetInput) (*clouddirectory.DeleteTypedLinkFacetOutput, error) {
	return a.client.DeleteTypedLinkFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DetachFromIndex(ctx context.Context, input *clouddirectory.DetachFromIndexInput) (*clouddirectory.DetachFromIndexOutput, error) {
	return a.client.DetachFromIndexWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DetachObject(ctx context.Context, input *clouddirectory.DetachObjectInput) (*clouddirectory.DetachObjectOutput, error) {
	return a.client.DetachObjectWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DetachPolicy(ctx context.Context, input *clouddirectory.DetachPolicyInput) (*clouddirectory.DetachPolicyOutput, error) {
	return a.client.DetachPolicyWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DetachTypedLink(ctx context.Context, input *clouddirectory.DetachTypedLinkInput) (*clouddirectory.DetachTypedLinkOutput, error) {
	return a.client.DetachTypedLinkWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) DisableDirectory(ctx context.Context, input *clouddirectory.DisableDirectoryInput) (*clouddirectory.DisableDirectoryOutput, error) {
	return a.client.DisableDirectoryWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) EnableDirectory(ctx context.Context, input *clouddirectory.EnableDirectoryInput) (*clouddirectory.EnableDirectoryOutput, error) {
	return a.client.EnableDirectoryWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetAppliedSchemaVersion(ctx context.Context, input *clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error) {
	return a.client.GetAppliedSchemaVersionWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetDirectory(ctx context.Context, input *clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error) {
	return a.client.GetDirectoryWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetFacet(ctx context.Context, input *clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error) {
	return a.client.GetFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetLinkAttributes(ctx context.Context, input *clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error) {
	return a.client.GetLinkAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetObjectAttributes(ctx context.Context, input *clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error) {
	return a.client.GetObjectAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetObjectInformation(ctx context.Context, input *clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error) {
	return a.client.GetObjectInformationWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetSchemaAsJson(ctx context.Context, input *clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error) {
	return a.client.GetSchemaAsJsonWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) GetTypedLinkFacetInformation(ctx context.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error) {
	return a.client.GetTypedLinkFacetInformationWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListAppliedSchemaArns(ctx context.Context, input *clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error) {
	return a.client.ListAppliedSchemaArnsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListAttachedIndices(ctx context.Context, input *clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error) {
	return a.client.ListAttachedIndicesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListDevelopmentSchemaArns(ctx context.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error) {
	return a.client.ListDevelopmentSchemaArnsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListDirectories(ctx context.Context, input *clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error) {
	return a.client.ListDirectoriesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListFacetAttributes(ctx context.Context, input *clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error) {
	return a.client.ListFacetAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListFacetNames(ctx context.Context, input *clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error) {
	return a.client.ListFacetNamesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListIncomingTypedLinks(ctx context.Context, input *clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
	return a.client.ListIncomingTypedLinksWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListIndex(ctx context.Context, input *clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error) {
	return a.client.ListIndexWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListManagedSchemaArns(ctx context.Context, input *clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error) {
	return a.client.ListManagedSchemaArnsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListObjectAttributes(ctx context.Context, input *clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error) {
	return a.client.ListObjectAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListObjectChildren(ctx context.Context, input *clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error) {
	return a.client.ListObjectChildrenWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListObjectParentPaths(ctx context.Context, input *clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error) {
	return a.client.ListObjectParentPathsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListObjectParents(ctx context.Context, input *clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error) {
	return a.client.ListObjectParentsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListObjectPolicies(ctx context.Context, input *clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error) {
	return a.client.ListObjectPoliciesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListOutgoingTypedLinks(ctx context.Context, input *clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
	return a.client.ListOutgoingTypedLinksWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListPolicyAttachments(ctx context.Context, input *clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
	return a.client.ListPolicyAttachmentsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListPublishedSchemaArns(ctx context.Context, input *clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error) {
	return a.client.ListPublishedSchemaArnsWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListTagsForResource(ctx context.Context, input *clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListTypedLinkFacetAttributes(ctx context.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error) {
	return a.client.ListTypedLinkFacetAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) ListTypedLinkFacetNames(ctx context.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error) {
	return a.client.ListTypedLinkFacetNamesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) LookupPolicy(ctx context.Context, input *clouddirectory.LookupPolicyInput) (*clouddirectory.LookupPolicyOutput, error) {
	return a.client.LookupPolicyWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) PublishSchema(ctx context.Context, input *clouddirectory.PublishSchemaInput) (*clouddirectory.PublishSchemaOutput, error) {
	return a.client.PublishSchemaWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) PutSchemaFromJson(ctx context.Context, input *clouddirectory.PutSchemaFromJsonInput) (*clouddirectory.PutSchemaFromJsonOutput, error) {
	return a.client.PutSchemaFromJsonWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) RemoveFacetFromObject(ctx context.Context, input *clouddirectory.RemoveFacetFromObjectInput) (*clouddirectory.RemoveFacetFromObjectOutput, error) {
	return a.client.RemoveFacetFromObjectWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) TagResource(ctx context.Context, input *clouddirectory.TagResourceInput) (*clouddirectory.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UntagResource(ctx context.Context, input *clouddirectory.UntagResourceInput) (*clouddirectory.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpdateFacet(ctx context.Context, input *clouddirectory.UpdateFacetInput) (*clouddirectory.UpdateFacetOutput, error) {
	return a.client.UpdateFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpdateLinkAttributes(ctx context.Context, input *clouddirectory.UpdateLinkAttributesInput) (*clouddirectory.UpdateLinkAttributesOutput, error) {
	return a.client.UpdateLinkAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpdateObjectAttributes(ctx context.Context, input *clouddirectory.UpdateObjectAttributesInput) (*clouddirectory.UpdateObjectAttributesOutput, error) {
	return a.client.UpdateObjectAttributesWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpdateSchema(ctx context.Context, input *clouddirectory.UpdateSchemaInput) (*clouddirectory.UpdateSchemaOutput, error) {
	return a.client.UpdateSchemaWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpdateTypedLinkFacet(ctx context.Context, input *clouddirectory.UpdateTypedLinkFacetInput) (*clouddirectory.UpdateTypedLinkFacetOutput, error) {
	return a.client.UpdateTypedLinkFacetWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpgradeAppliedSchema(ctx context.Context, input *clouddirectory.UpgradeAppliedSchemaInput) (*clouddirectory.UpgradeAppliedSchemaOutput, error) {
	return a.client.UpgradeAppliedSchemaWithContext(ctx, input)
}

func (a *CloudDirectoryActivities) UpgradePublishedSchema(ctx context.Context, input *clouddirectory.UpgradePublishedSchemaInput) (*clouddirectory.UpgradePublishedSchemaOutput, error) {
	return a.client.UpgradePublishedSchemaWithContext(ctx, input)
}
