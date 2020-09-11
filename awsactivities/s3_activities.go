
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type S3Activities struct {
    client s3iface.S3API
}

func NewS3Activities(session *session.Session, config ...*aws.Config) *S3Activities {
    client := s3.New(session, config...)
    return &S3Activities{client: client}
}

func (a *S3Activities) AbortMultipartUpload(ctx context.Context, input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
    return a.client.AbortMultipartUploadWithContext(ctx, input)
}

func (a *S3Activities) CompleteMultipartUpload(ctx context.Context, input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
    return a.client.CompleteMultipartUploadWithContext(ctx, input)
}

func (a *S3Activities) CopyObject(ctx context.Context, input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
    return a.client.CopyObjectWithContext(ctx, input)
}

func (a *S3Activities) CreateBucket(ctx context.Context, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
    return a.client.CreateBucketWithContext(ctx, input)
}

func (a *S3Activities) CreateMultipartUpload(ctx context.Context, input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
    return a.client.CreateMultipartUploadWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucket(ctx context.Context, input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
    return a.client.DeleteBucketWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketAnalyticsConfiguration(ctx context.Context, input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
    return a.client.DeleteBucketAnalyticsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketCors(ctx context.Context, input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
    return a.client.DeleteBucketCorsWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketEncryption(ctx context.Context, input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
    return a.client.DeleteBucketEncryptionWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketInventoryConfiguration(ctx context.Context, input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
    return a.client.DeleteBucketInventoryConfigurationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketLifecycle(ctx context.Context, input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
    return a.client.DeleteBucketLifecycleWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketMetricsConfiguration(ctx context.Context, input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
    return a.client.DeleteBucketMetricsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketPolicy(ctx context.Context, input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
    return a.client.DeleteBucketPolicyWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketReplication(ctx context.Context, input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
    return a.client.DeleteBucketReplicationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketTagging(ctx context.Context, input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
    return a.client.DeleteBucketTaggingWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketWebsite(ctx context.Context, input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
    return a.client.DeleteBucketWebsiteWithContext(ctx, input)
}

func (a *S3Activities) DeleteObject(ctx context.Context, input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
    return a.client.DeleteObjectWithContext(ctx, input)
}

func (a *S3Activities) DeleteObjectTagging(ctx context.Context, input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
    return a.client.DeleteObjectTaggingWithContext(ctx, input)
}

func (a *S3Activities) DeleteObjects(ctx context.Context, input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
    return a.client.DeleteObjectsWithContext(ctx, input)
}

func (a *S3Activities) DeletePublicAccessBlock(ctx context.Context, input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
    return a.client.DeletePublicAccessBlockWithContext(ctx, input)
}

func (a *S3Activities) GetBucketAccelerateConfiguration(ctx context.Context, input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
    return a.client.GetBucketAccelerateConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketAcl(ctx context.Context, input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
    return a.client.GetBucketAclWithContext(ctx, input)
}

func (a *S3Activities) GetBucketAnalyticsConfiguration(ctx context.Context, input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
    return a.client.GetBucketAnalyticsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketCors(ctx context.Context, input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
    return a.client.GetBucketCorsWithContext(ctx, input)
}

func (a *S3Activities) GetBucketEncryption(ctx context.Context, input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
    return a.client.GetBucketEncryptionWithContext(ctx, input)
}

func (a *S3Activities) GetBucketInventoryConfiguration(ctx context.Context, input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
    return a.client.GetBucketInventoryConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLifecycle(ctx context.Context, input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
    return a.client.GetBucketLifecycleWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLifecycleConfiguration(ctx context.Context, input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
    return a.client.GetBucketLifecycleConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLocation(ctx context.Context, input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
    return a.client.GetBucketLocationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLogging(ctx context.Context, input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
    return a.client.GetBucketLoggingWithContext(ctx, input)
}

func (a *S3Activities) GetBucketMetricsConfiguration(ctx context.Context, input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
    return a.client.GetBucketMetricsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketNotification(ctx context.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
    return a.client.GetBucketNotificationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketNotificationConfiguration(ctx context.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
    return a.client.GetBucketNotificationConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketPolicy(ctx context.Context, input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
    return a.client.GetBucketPolicyWithContext(ctx, input)
}

func (a *S3Activities) GetBucketPolicyStatus(ctx context.Context, input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
    return a.client.GetBucketPolicyStatusWithContext(ctx, input)
}

func (a *S3Activities) GetBucketReplication(ctx context.Context, input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
    return a.client.GetBucketReplicationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketRequestPayment(ctx context.Context, input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
    return a.client.GetBucketRequestPaymentWithContext(ctx, input)
}

func (a *S3Activities) GetBucketTagging(ctx context.Context, input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
    return a.client.GetBucketTaggingWithContext(ctx, input)
}

func (a *S3Activities) GetBucketVersioning(ctx context.Context, input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
    return a.client.GetBucketVersioningWithContext(ctx, input)
}

func (a *S3Activities) GetBucketWebsite(ctx context.Context, input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
    return a.client.GetBucketWebsiteWithContext(ctx, input)
}

func (a *S3Activities) GetObject(ctx context.Context, input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
    return a.client.GetObjectWithContext(ctx, input)
}

func (a *S3Activities) GetObjectAcl(ctx context.Context, input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
    return a.client.GetObjectAclWithContext(ctx, input)
}

func (a *S3Activities) GetObjectLegalHold(ctx context.Context, input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
    return a.client.GetObjectLegalHoldWithContext(ctx, input)
}

func (a *S3Activities) GetObjectLockConfiguration(ctx context.Context, input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
    return a.client.GetObjectLockConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetObjectRetention(ctx context.Context, input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
    return a.client.GetObjectRetentionWithContext(ctx, input)
}

func (a *S3Activities) GetObjectTagging(ctx context.Context, input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
    return a.client.GetObjectTaggingWithContext(ctx, input)
}

func (a *S3Activities) GetObjectTorrent(ctx context.Context, input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
    return a.client.GetObjectTorrentWithContext(ctx, input)
}

func (a *S3Activities) GetPublicAccessBlock(ctx context.Context, input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
    return a.client.GetPublicAccessBlockWithContext(ctx, input)
}

func (a *S3Activities) HeadBucket(ctx context.Context, input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
    return a.client.HeadBucketWithContext(ctx, input)
}

func (a *S3Activities) HeadObject(ctx context.Context, input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
    return a.client.HeadObjectWithContext(ctx, input)
}

func (a *S3Activities) ListBucketAnalyticsConfigurations(ctx context.Context, input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
    return a.client.ListBucketAnalyticsConfigurationsWithContext(ctx, input)
}

func (a *S3Activities) ListBucketInventoryConfigurations(ctx context.Context, input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
    return a.client.ListBucketInventoryConfigurationsWithContext(ctx, input)
}

func (a *S3Activities) ListBucketMetricsConfigurations(ctx context.Context, input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
    return a.client.ListBucketMetricsConfigurationsWithContext(ctx, input)
}

func (a *S3Activities) ListBuckets(ctx context.Context, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
    return a.client.ListBucketsWithContext(ctx, input)
}

func (a *S3Activities) ListMultipartUploads(ctx context.Context, input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
    return a.client.ListMultipartUploadsWithContext(ctx, input)
}

func (a *S3Activities) ListObjectVersions(ctx context.Context, input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
    return a.client.ListObjectVersionsWithContext(ctx, input)
}

func (a *S3Activities) ListObjects(ctx context.Context, input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
    return a.client.ListObjectsWithContext(ctx, input)
}

func (a *S3Activities) ListObjectsV2(ctx context.Context, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
    return a.client.ListObjectsV2WithContext(ctx, input)
}

func (a *S3Activities) ListParts(ctx context.Context, input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
    return a.client.ListPartsWithContext(ctx, input)
}

func (a *S3Activities) PutBucketAccelerateConfiguration(ctx context.Context, input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
    return a.client.PutBucketAccelerateConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketAcl(ctx context.Context, input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
    return a.client.PutBucketAclWithContext(ctx, input)
}

func (a *S3Activities) PutBucketAnalyticsConfiguration(ctx context.Context, input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
    return a.client.PutBucketAnalyticsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketCors(ctx context.Context, input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
    return a.client.PutBucketCorsWithContext(ctx, input)
}

func (a *S3Activities) PutBucketEncryption(ctx context.Context, input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
    return a.client.PutBucketEncryptionWithContext(ctx, input)
}

func (a *S3Activities) PutBucketInventoryConfiguration(ctx context.Context, input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
    return a.client.PutBucketInventoryConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketLifecycle(ctx context.Context, input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
    return a.client.PutBucketLifecycleWithContext(ctx, input)
}

func (a *S3Activities) PutBucketLifecycleConfiguration(ctx context.Context, input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
    return a.client.PutBucketLifecycleConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketLogging(ctx context.Context, input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
    return a.client.PutBucketLoggingWithContext(ctx, input)
}

func (a *S3Activities) PutBucketMetricsConfiguration(ctx context.Context, input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
    return a.client.PutBucketMetricsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketNotification(ctx context.Context, input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
    return a.client.PutBucketNotificationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketNotificationConfiguration(ctx context.Context, input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
    return a.client.PutBucketNotificationConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketPolicy(ctx context.Context, input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
    return a.client.PutBucketPolicyWithContext(ctx, input)
}

func (a *S3Activities) PutBucketReplication(ctx context.Context, input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
    return a.client.PutBucketReplicationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketRequestPayment(ctx context.Context, input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
    return a.client.PutBucketRequestPaymentWithContext(ctx, input)
}

func (a *S3Activities) PutBucketTagging(ctx context.Context, input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
    return a.client.PutBucketTaggingWithContext(ctx, input)
}

func (a *S3Activities) PutBucketVersioning(ctx context.Context, input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
    return a.client.PutBucketVersioningWithContext(ctx, input)
}

func (a *S3Activities) PutBucketWebsite(ctx context.Context, input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
    return a.client.PutBucketWebsiteWithContext(ctx, input)
}

func (a *S3Activities) PutObject(ctx context.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
    return a.client.PutObjectWithContext(ctx, input)
}

func (a *S3Activities) PutObjectAcl(ctx context.Context, input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
    return a.client.PutObjectAclWithContext(ctx, input)
}

func (a *S3Activities) PutObjectLegalHold(ctx context.Context, input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
    return a.client.PutObjectLegalHoldWithContext(ctx, input)
}

func (a *S3Activities) PutObjectLockConfiguration(ctx context.Context, input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
    return a.client.PutObjectLockConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutObjectRetention(ctx context.Context, input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
    return a.client.PutObjectRetentionWithContext(ctx, input)
}

func (a *S3Activities) PutObjectTagging(ctx context.Context, input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
    return a.client.PutObjectTaggingWithContext(ctx, input)
}

func (a *S3Activities) PutPublicAccessBlock(ctx context.Context, input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
    return a.client.PutPublicAccessBlockWithContext(ctx, input)
}

func (a *S3Activities) RestoreObject(ctx context.Context, input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
    return a.client.RestoreObjectWithContext(ctx, input)
}

func (a *S3Activities) SelectObjectContent(ctx context.Context, input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
    return a.client.SelectObjectContentWithContext(ctx, input)
}

func (a *S3Activities) UploadPart(ctx context.Context, input *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
    return a.client.UploadPartWithContext(ctx, input)
}

func (a *S3Activities) UploadPartCopy(ctx context.Context, input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
    return a.client.UploadPartCopyWithContext(ctx, input)
}

func (a *S3Activities) WaitUntilBucketExists(ctx context.Context, input *s3.HeadBucketInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilBucketExistsWithContext(ctx, input, options...)
	})
}

func (a *S3Activities) WaitUntilBucketNotExists(ctx context.Context, input *s3.HeadBucketInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilBucketNotExistsWithContext(ctx, input, options...)
	})
}

func (a *S3Activities) WaitUntilObjectExists(ctx context.Context, input *s3.HeadObjectInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilObjectExistsWithContext(ctx, input, options...)
	})
}

func (a *S3Activities) WaitUntilObjectNotExists(ctx context.Context, input *s3.HeadObjectInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilObjectNotExistsWithContext(ctx, input, options...)
	})
}
