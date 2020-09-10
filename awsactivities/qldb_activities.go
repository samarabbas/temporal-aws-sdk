package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldb/qldbiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type QLDBActivities struct {
	client qldbiface.QLDBAPI
}

func NewQLDBActivities(session *session.Session, config ...*aws.Config) *QLDBActivities {
	client := qldb.New(session, config...)
	return &QLDBActivities{client: client}
}

func (a *QLDBActivities) CancelJournalKinesisStream(ctx context.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error) {
	return a.client.CancelJournalKinesisStreamWithContext(ctx, input)
}

func (a *QLDBActivities) CreateLedger(ctx context.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error) {
	return a.client.CreateLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) DeleteLedger(ctx context.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error) {
	return a.client.DeleteLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) DescribeJournalKinesisStream(ctx context.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	return a.client.DescribeJournalKinesisStreamWithContext(ctx, input)
}

func (a *QLDBActivities) DescribeJournalS3Export(ctx context.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error) {
	return a.client.DescribeJournalS3ExportWithContext(ctx, input)
}

func (a *QLDBActivities) DescribeLedger(ctx context.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error) {
	return a.client.DescribeLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) ExportJournalToS3(ctx context.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error) {
	return a.client.ExportJournalToS3WithContext(ctx, input)
}

func (a *QLDBActivities) GetBlock(ctx context.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error) {
	return a.client.GetBlockWithContext(ctx, input)
}

func (a *QLDBActivities) GetDigest(ctx context.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error) {
	return a.client.GetDigestWithContext(ctx, input)
}

func (a *QLDBActivities) GetRevision(ctx context.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error) {
	return a.client.GetRevisionWithContext(ctx, input)
}

func (a *QLDBActivities) ListJournalKinesisStreamsForLedger(ctx context.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	return a.client.ListJournalKinesisStreamsForLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) ListJournalS3Exports(ctx context.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error) {
	return a.client.ListJournalS3ExportsWithContext(ctx, input)
}

func (a *QLDBActivities) ListJournalS3ExportsForLedger(ctx context.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	return a.client.ListJournalS3ExportsForLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) ListLedgers(ctx context.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error) {
	return a.client.ListLedgersWithContext(ctx, input)
}

func (a *QLDBActivities) ListTagsForResource(ctx context.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *QLDBActivities) StreamJournalToKinesis(ctx context.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error) {
	return a.client.StreamJournalToKinesisWithContext(ctx, input)
}

func (a *QLDBActivities) TagResource(ctx context.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *QLDBActivities) UntagResource(ctx context.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *QLDBActivities) UpdateLedger(ctx context.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error) {
	return a.client.UpdateLedgerWithContext(ctx, input)
}
