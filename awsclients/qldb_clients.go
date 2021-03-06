// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/qldb"
	"go.temporal.io/sdk/workflow"
)

type QLDBClient interface {
	CancelJournalKinesisStream(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error)
	CancelJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) *QldbCancelJournalKinesisStreamResult

	CreateLedger(ctx workflow.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error)
	CreateLedgerAsync(ctx workflow.Context, input *qldb.CreateLedgerInput) *QldbCreateLedgerResult

	DeleteLedger(ctx workflow.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error)
	DeleteLedgerAsync(ctx workflow.Context, input *qldb.DeleteLedgerInput) *QldbDeleteLedgerResult

	DescribeJournalKinesisStream(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error)
	DescribeJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) *QldbDescribeJournalKinesisStreamResult

	DescribeJournalS3Export(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error)
	DescribeJournalS3ExportAsync(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) *QldbDescribeJournalS3ExportResult

	DescribeLedger(ctx workflow.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error)
	DescribeLedgerAsync(ctx workflow.Context, input *qldb.DescribeLedgerInput) *QldbDescribeLedgerResult

	ExportJournalToS3(ctx workflow.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error)
	ExportJournalToS3Async(ctx workflow.Context, input *qldb.ExportJournalToS3Input) *QldbExportJournalToS3Result

	GetBlock(ctx workflow.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error)
	GetBlockAsync(ctx workflow.Context, input *qldb.GetBlockInput) *QldbGetBlockResult

	GetDigest(ctx workflow.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error)
	GetDigestAsync(ctx workflow.Context, input *qldb.GetDigestInput) *QldbGetDigestResult

	GetRevision(ctx workflow.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error)
	GetRevisionAsync(ctx workflow.Context, input *qldb.GetRevisionInput) *QldbGetRevisionResult

	ListJournalKinesisStreamsForLedger(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error)
	ListJournalKinesisStreamsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) *QldbListJournalKinesisStreamsForLedgerResult

	ListJournalS3Exports(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error)
	ListJournalS3ExportsAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) *QldbListJournalS3ExportsResult

	ListJournalS3ExportsForLedger(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error)
	ListJournalS3ExportsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) *QldbListJournalS3ExportsForLedgerResult

	ListLedgers(ctx workflow.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error)
	ListLedgersAsync(ctx workflow.Context, input *qldb.ListLedgersInput) *QldbListLedgersResult

	ListTagsForResource(ctx workflow.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *qldb.ListTagsForResourceInput) *QldbListTagsForResourceResult

	StreamJournalToKinesis(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error)
	StreamJournalToKinesisAsync(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) *QldbStreamJournalToKinesisResult

	TagResource(ctx workflow.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *qldb.TagResourceInput) *QldbTagResourceResult

	UntagResource(ctx workflow.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *qldb.UntagResourceInput) *QldbUntagResourceResult

	UpdateLedger(ctx workflow.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error)
	UpdateLedgerAsync(ctx workflow.Context, input *qldb.UpdateLedgerInput) *QldbUpdateLedgerResult
}

type QLDBStub struct{}

func NewQLDBStub() QLDBClient {
	return &QLDBStub{}
}

type QldbCancelJournalKinesisStreamResult struct {
	Result workflow.Future
}

func (r *QldbCancelJournalKinesisStreamResult) Get(ctx workflow.Context) (*qldb.CancelJournalKinesisStreamOutput, error) {
	var output qldb.CancelJournalKinesisStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbCreateLedgerResult struct {
	Result workflow.Future
}

func (r *QldbCreateLedgerResult) Get(ctx workflow.Context) (*qldb.CreateLedgerOutput, error) {
	var output qldb.CreateLedgerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbDeleteLedgerResult struct {
	Result workflow.Future
}

func (r *QldbDeleteLedgerResult) Get(ctx workflow.Context) (*qldb.DeleteLedgerOutput, error) {
	var output qldb.DeleteLedgerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbDescribeJournalKinesisStreamResult struct {
	Result workflow.Future
}

func (r *QldbDescribeJournalKinesisStreamResult) Get(ctx workflow.Context) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	var output qldb.DescribeJournalKinesisStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbDescribeJournalS3ExportResult struct {
	Result workflow.Future
}

func (r *QldbDescribeJournalS3ExportResult) Get(ctx workflow.Context) (*qldb.DescribeJournalS3ExportOutput, error) {
	var output qldb.DescribeJournalS3ExportOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbDescribeLedgerResult struct {
	Result workflow.Future
}

func (r *QldbDescribeLedgerResult) Get(ctx workflow.Context) (*qldb.DescribeLedgerOutput, error) {
	var output qldb.DescribeLedgerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbExportJournalToS3Result struct {
	Result workflow.Future
}

func (r *QldbExportJournalToS3Result) Get(ctx workflow.Context) (*qldb.ExportJournalToS3Output, error) {
	var output qldb.ExportJournalToS3Output
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbGetBlockResult struct {
	Result workflow.Future
}

func (r *QldbGetBlockResult) Get(ctx workflow.Context) (*qldb.GetBlockOutput, error) {
	var output qldb.GetBlockOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbGetDigestResult struct {
	Result workflow.Future
}

func (r *QldbGetDigestResult) Get(ctx workflow.Context) (*qldb.GetDigestOutput, error) {
	var output qldb.GetDigestOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbGetRevisionResult struct {
	Result workflow.Future
}

func (r *QldbGetRevisionResult) Get(ctx workflow.Context) (*qldb.GetRevisionOutput, error) {
	var output qldb.GetRevisionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbListJournalKinesisStreamsForLedgerResult struct {
	Result workflow.Future
}

func (r *QldbListJournalKinesisStreamsForLedgerResult) Get(ctx workflow.Context) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	var output qldb.ListJournalKinesisStreamsForLedgerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbListJournalS3ExportsResult struct {
	Result workflow.Future
}

func (r *QldbListJournalS3ExportsResult) Get(ctx workflow.Context) (*qldb.ListJournalS3ExportsOutput, error) {
	var output qldb.ListJournalS3ExportsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbListJournalS3ExportsForLedgerResult struct {
	Result workflow.Future
}

func (r *QldbListJournalS3ExportsForLedgerResult) Get(ctx workflow.Context) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	var output qldb.ListJournalS3ExportsForLedgerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbListLedgersResult struct {
	Result workflow.Future
}

func (r *QldbListLedgersResult) Get(ctx workflow.Context) (*qldb.ListLedgersOutput, error) {
	var output qldb.ListLedgersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *QldbListTagsForResourceResult) Get(ctx workflow.Context) (*qldb.ListTagsForResourceOutput, error) {
	var output qldb.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbStreamJournalToKinesisResult struct {
	Result workflow.Future
}

func (r *QldbStreamJournalToKinesisResult) Get(ctx workflow.Context) (*qldb.StreamJournalToKinesisOutput, error) {
	var output qldb.StreamJournalToKinesisOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbTagResourceResult struct {
	Result workflow.Future
}

func (r *QldbTagResourceResult) Get(ctx workflow.Context) (*qldb.TagResourceOutput, error) {
	var output qldb.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbUntagResourceResult struct {
	Result workflow.Future
}

func (r *QldbUntagResourceResult) Get(ctx workflow.Context) (*qldb.UntagResourceOutput, error) {
	var output qldb.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QldbUpdateLedgerResult struct {
	Result workflow.Future
}

func (r *QldbUpdateLedgerResult) Get(ctx workflow.Context) (*qldb.UpdateLedgerOutput, error) {
	var output qldb.UpdateLedgerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) CancelJournalKinesisStream(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error) {
	var output qldb.CancelJournalKinesisStreamOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.CancelJournalKinesisStream", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) CancelJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) *QldbCancelJournalKinesisStreamResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.CancelJournalKinesisStream", input)
	return &QldbCancelJournalKinesisStreamResult{Result: future}
}

func (a *QLDBStub) CreateLedger(ctx workflow.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error) {
	var output qldb.CreateLedgerOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.CreateLedger", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) CreateLedgerAsync(ctx workflow.Context, input *qldb.CreateLedgerInput) *QldbCreateLedgerResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.CreateLedger", input)
	return &QldbCreateLedgerResult{Result: future}
}

func (a *QLDBStub) DeleteLedger(ctx workflow.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error) {
	var output qldb.DeleteLedgerOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.DeleteLedger", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) DeleteLedgerAsync(ctx workflow.Context, input *qldb.DeleteLedgerInput) *QldbDeleteLedgerResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.DeleteLedger", input)
	return &QldbDeleteLedgerResult{Result: future}
}

func (a *QLDBStub) DescribeJournalKinesisStream(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	var output qldb.DescribeJournalKinesisStreamOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.DescribeJournalKinesisStream", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) DescribeJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) *QldbDescribeJournalKinesisStreamResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.DescribeJournalKinesisStream", input)
	return &QldbDescribeJournalKinesisStreamResult{Result: future}
}

func (a *QLDBStub) DescribeJournalS3Export(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error) {
	var output qldb.DescribeJournalS3ExportOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.DescribeJournalS3Export", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) DescribeJournalS3ExportAsync(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) *QldbDescribeJournalS3ExportResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.DescribeJournalS3Export", input)
	return &QldbDescribeJournalS3ExportResult{Result: future}
}

func (a *QLDBStub) DescribeLedger(ctx workflow.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error) {
	var output qldb.DescribeLedgerOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.DescribeLedger", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) DescribeLedgerAsync(ctx workflow.Context, input *qldb.DescribeLedgerInput) *QldbDescribeLedgerResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.DescribeLedger", input)
	return &QldbDescribeLedgerResult{Result: future}
}

func (a *QLDBStub) ExportJournalToS3(ctx workflow.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error) {
	var output qldb.ExportJournalToS3Output
	err := workflow.ExecuteActivity(ctx, "QLDB.ExportJournalToS3", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) ExportJournalToS3Async(ctx workflow.Context, input *qldb.ExportJournalToS3Input) *QldbExportJournalToS3Result {
	future := workflow.ExecuteActivity(ctx, "QLDB.ExportJournalToS3", input)
	return &QldbExportJournalToS3Result{Result: future}
}

func (a *QLDBStub) GetBlock(ctx workflow.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error) {
	var output qldb.GetBlockOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.GetBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) GetBlockAsync(ctx workflow.Context, input *qldb.GetBlockInput) *QldbGetBlockResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.GetBlock", input)
	return &QldbGetBlockResult{Result: future}
}

func (a *QLDBStub) GetDigest(ctx workflow.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error) {
	var output qldb.GetDigestOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.GetDigest", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) GetDigestAsync(ctx workflow.Context, input *qldb.GetDigestInput) *QldbGetDigestResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.GetDigest", input)
	return &QldbGetDigestResult{Result: future}
}

func (a *QLDBStub) GetRevision(ctx workflow.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error) {
	var output qldb.GetRevisionOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.GetRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) GetRevisionAsync(ctx workflow.Context, input *qldb.GetRevisionInput) *QldbGetRevisionResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.GetRevision", input)
	return &QldbGetRevisionResult{Result: future}
}

func (a *QLDBStub) ListJournalKinesisStreamsForLedger(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	var output qldb.ListJournalKinesisStreamsForLedgerOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.ListJournalKinesisStreamsForLedger", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) ListJournalKinesisStreamsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) *QldbListJournalKinesisStreamsForLedgerResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.ListJournalKinesisStreamsForLedger", input)
	return &QldbListJournalKinesisStreamsForLedgerResult{Result: future}
}

func (a *QLDBStub) ListJournalS3Exports(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error) {
	var output qldb.ListJournalS3ExportsOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.ListJournalS3Exports", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) ListJournalS3ExportsAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) *QldbListJournalS3ExportsResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.ListJournalS3Exports", input)
	return &QldbListJournalS3ExportsResult{Result: future}
}

func (a *QLDBStub) ListJournalS3ExportsForLedger(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	var output qldb.ListJournalS3ExportsForLedgerOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.ListJournalS3ExportsForLedger", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) ListJournalS3ExportsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) *QldbListJournalS3ExportsForLedgerResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.ListJournalS3ExportsForLedger", input)
	return &QldbListJournalS3ExportsForLedgerResult{Result: future}
}

func (a *QLDBStub) ListLedgers(ctx workflow.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error) {
	var output qldb.ListLedgersOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.ListLedgers", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) ListLedgersAsync(ctx workflow.Context, input *qldb.ListLedgersInput) *QldbListLedgersResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.ListLedgers", input)
	return &QldbListLedgersResult{Result: future}
}

func (a *QLDBStub) ListTagsForResource(ctx workflow.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error) {
	var output qldb.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) ListTagsForResourceAsync(ctx workflow.Context, input *qldb.ListTagsForResourceInput) *QldbListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.ListTagsForResource", input)
	return &QldbListTagsForResourceResult{Result: future}
}

func (a *QLDBStub) StreamJournalToKinesis(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error) {
	var output qldb.StreamJournalToKinesisOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.StreamJournalToKinesis", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) StreamJournalToKinesisAsync(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) *QldbStreamJournalToKinesisResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.StreamJournalToKinesis", input)
	return &QldbStreamJournalToKinesisResult{Result: future}
}

func (a *QLDBStub) TagResource(ctx workflow.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error) {
	var output qldb.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) TagResourceAsync(ctx workflow.Context, input *qldb.TagResourceInput) *QldbTagResourceResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.TagResource", input)
	return &QldbTagResourceResult{Result: future}
}

func (a *QLDBStub) UntagResource(ctx workflow.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error) {
	var output qldb.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) UntagResourceAsync(ctx workflow.Context, input *qldb.UntagResourceInput) *QldbUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.UntagResource", input)
	return &QldbUntagResourceResult{Result: future}
}

func (a *QLDBStub) UpdateLedger(ctx workflow.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error) {
	var output qldb.UpdateLedgerOutput
	err := workflow.ExecuteActivity(ctx, "QLDB.UpdateLedger", input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBStub) UpdateLedgerAsync(ctx workflow.Context, input *qldb.UpdateLedgerInput) *QldbUpdateLedgerResult {
	future := workflow.ExecuteActivity(ctx, "QLDB.UpdateLedger", input)
	return &QldbUpdateLedgerResult{Result: future}
}
