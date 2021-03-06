// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"go.temporal.io/sdk/workflow"
)

type CloudSearchDomainClient interface {
	Search(ctx workflow.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error)
	SearchAsync(ctx workflow.Context, input *cloudsearchdomain.SearchInput) *CloudsearchdomainSearchResult

	Suggest(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error)
	SuggestAsync(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) *CloudsearchdomainSuggestResult

	UploadDocuments(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error)
	UploadDocumentsAsync(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) *CloudsearchdomainUploadDocumentsResult
}

type CloudSearchDomainStub struct{}

func NewCloudSearchDomainStub() CloudSearchDomainClient {
	return &CloudSearchDomainStub{}
}

type CloudsearchdomainSearchResult struct {
	Result workflow.Future
}

func (r *CloudsearchdomainSearchResult) Get(ctx workflow.Context) (*cloudsearchdomain.SearchOutput, error) {
	var output cloudsearchdomain.SearchOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudsearchdomainSuggestResult struct {
	Result workflow.Future
}

func (r *CloudsearchdomainSuggestResult) Get(ctx workflow.Context) (*cloudsearchdomain.SuggestOutput, error) {
	var output cloudsearchdomain.SuggestOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CloudsearchdomainUploadDocumentsResult struct {
	Result workflow.Future
}

func (r *CloudsearchdomainUploadDocumentsResult) Get(ctx workflow.Context) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	var output cloudsearchdomain.UploadDocumentsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *CloudSearchDomainStub) Search(ctx workflow.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	var output cloudsearchdomain.SearchOutput
	err := workflow.ExecuteActivity(ctx, "CloudSearchDomain.Search", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudSearchDomainStub) SearchAsync(ctx workflow.Context, input *cloudsearchdomain.SearchInput) *CloudsearchdomainSearchResult {
	future := workflow.ExecuteActivity(ctx, "CloudSearchDomain.Search", input)
	return &CloudsearchdomainSearchResult{Result: future}
}

func (a *CloudSearchDomainStub) Suggest(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
	var output cloudsearchdomain.SuggestOutput
	err := workflow.ExecuteActivity(ctx, "CloudSearchDomain.Suggest", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudSearchDomainStub) SuggestAsync(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) *CloudsearchdomainSuggestResult {
	future := workflow.ExecuteActivity(ctx, "CloudSearchDomain.Suggest", input)
	return &CloudsearchdomainSuggestResult{Result: future}
}

func (a *CloudSearchDomainStub) UploadDocuments(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	var output cloudsearchdomain.UploadDocumentsOutput
	err := workflow.ExecuteActivity(ctx, "CloudSearchDomain.UploadDocuments", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudSearchDomainStub) UploadDocumentsAsync(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) *CloudsearchdomainUploadDocumentsResult {
	future := workflow.ExecuteActivity(ctx, "CloudSearchDomain.UploadDocuments", input)
	return &CloudsearchdomainUploadDocumentsResult{Result: future}
}
