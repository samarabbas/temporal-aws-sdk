// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/acmpca"
	"go.temporal.io/sdk/workflow"
)

type ACMPCAClient interface {
	CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error)
	CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *AcmpcaCreateCertificateAuthorityResult

	CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
	CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *AcmpcaCreateCertificateAuthorityAuditReportResult

	CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error)
	CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *AcmpcaCreatePermissionResult

	DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error)
	DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *AcmpcaDeleteCertificateAuthorityResult

	DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error)
	DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *AcmpcaDeletePermissionResult

	DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *AcmpcaDeletePolicyResult

	DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error)
	DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *AcmpcaDescribeCertificateAuthorityResult

	DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
	DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *AcmpcaDescribeCertificateAuthorityAuditReportResult

	GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error)
	GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *AcmpcaGetCertificateResult

	GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
	GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *AcmpcaGetCertificateAuthorityCertificateResult

	GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error)
	GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *AcmpcaGetCertificateAuthorityCsrResult

	GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *AcmpcaGetPolicyResult

	ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
	ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *AcmpcaImportCertificateAuthorityCertificateResult

	IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error)
	IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *AcmpcaIssueCertificateResult

	ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error)
	ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *AcmpcaListCertificateAuthoritiesResult

	ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error)
	ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *AcmpcaListPermissionsResult

	ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *AcmpcaListTagsResult

	PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error)
	PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *AcmpcaPutPolicyResult

	RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error)
	RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *AcmpcaRestoreCertificateAuthorityResult

	RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error)
	RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *AcmpcaRevokeCertificateResult

	TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error)
	TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *AcmpcaTagCertificateAuthorityResult

	UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error)
	UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *AcmpcaUntagCertificateAuthorityResult

	UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error)
	UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *AcmpcaUpdateCertificateAuthorityResult

	WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error
	WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error
	WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error}

type ACMPCAStub struct{}

func NewACMPCAStub() ACMPCAClient {
	return &ACMPCAStub{}
}

type AcmpcaCreateCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaCreateCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityOutput, error) {
	var output acmpca.CreateCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaCreateCertificateAuthorityAuditReportResult struct {
	Result workflow.Future
}

func (r *AcmpcaCreateCertificateAuthorityAuditReportResult) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.CreateCertificateAuthorityAuditReportOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaCreatePermissionResult struct {
	Result workflow.Future
}

func (r *AcmpcaCreatePermissionResult) Get(ctx workflow.Context) (*acmpca.CreatePermissionOutput, error) {
	var output acmpca.CreatePermissionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaDeleteCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaDeleteCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	var output acmpca.DeleteCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaDeletePermissionResult struct {
	Result workflow.Future
}

func (r *AcmpcaDeletePermissionResult) Get(ctx workflow.Context) (*acmpca.DeletePermissionOutput, error) {
	var output acmpca.DeletePermissionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaDeletePolicyResult struct {
	Result workflow.Future
}

func (r *AcmpcaDeletePolicyResult) Get(ctx workflow.Context) (*acmpca.DeletePolicyOutput, error) {
	var output acmpca.DeletePolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaDescribeCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaDescribeCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	var output acmpca.DescribeCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaDescribeCertificateAuthorityAuditReportResult struct {
	Result workflow.Future
}

func (r *AcmpcaDescribeCertificateAuthorityAuditReportResult) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.DescribeCertificateAuthorityAuditReportOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetCertificateResult) Get(ctx workflow.Context) (*acmpca.GetCertificateOutput, error) {
	var output acmpca.GetCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetCertificateAuthorityCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetCertificateAuthorityCertificateResult) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	var output acmpca.GetCertificateAuthorityCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetCertificateAuthorityCsrResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetCertificateAuthorityCsrResult) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	var output acmpca.GetCertificateAuthorityCsrOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaGetPolicyResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetPolicyResult) Get(ctx workflow.Context) (*acmpca.GetPolicyOutput, error) {
	var output acmpca.GetPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaImportCertificateAuthorityCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaImportCertificateAuthorityCertificateResult) Get(ctx workflow.Context) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	var output acmpca.ImportCertificateAuthorityCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaIssueCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaIssueCertificateResult) Get(ctx workflow.Context) (*acmpca.IssueCertificateOutput, error) {
	var output acmpca.IssueCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaListCertificateAuthoritiesResult struct {
	Result workflow.Future
}

func (r *AcmpcaListCertificateAuthoritiesResult) Get(ctx workflow.Context) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	var output acmpca.ListCertificateAuthoritiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaListPermissionsResult struct {
	Result workflow.Future
}

func (r *AcmpcaListPermissionsResult) Get(ctx workflow.Context) (*acmpca.ListPermissionsOutput, error) {
	var output acmpca.ListPermissionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaListTagsResult struct {
	Result workflow.Future
}

func (r *AcmpcaListTagsResult) Get(ctx workflow.Context) (*acmpca.ListTagsOutput, error) {
	var output acmpca.ListTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaPutPolicyResult struct {
	Result workflow.Future
}

func (r *AcmpcaPutPolicyResult) Get(ctx workflow.Context) (*acmpca.PutPolicyOutput, error) {
	var output acmpca.PutPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaRestoreCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaRestoreCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	var output acmpca.RestoreCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaRevokeCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaRevokeCertificateResult) Get(ctx workflow.Context) (*acmpca.RevokeCertificateOutput, error) {
	var output acmpca.RevokeCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaTagCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaTagCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.TagCertificateAuthorityOutput, error) {
	var output acmpca.TagCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaUntagCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaUntagCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.UntagCertificateAuthorityOutput, error) {
	var output acmpca.UntagCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AcmpcaUpdateCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaUpdateCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	var output acmpca.UpdateCertificateAuthorityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}




func (a *ACMPCAStub) CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
	var output acmpca.CreateCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.CreateCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *AcmpcaCreateCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.CreateCertificateAuthority", input)
	return &AcmpcaCreateCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.CreateCertificateAuthorityAuditReportOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.CreateCertificateAuthorityAuditReport", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *AcmpcaCreateCertificateAuthorityAuditReportResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.CreateCertificateAuthorityAuditReport", input)
	return &AcmpcaCreateCertificateAuthorityAuditReportResult{Result: future}
}

func (a *ACMPCAStub) CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
	var output acmpca.CreatePermissionOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.CreatePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *AcmpcaCreatePermissionResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.CreatePermission", input)
	return &AcmpcaCreatePermissionResult{Result: future}
}

func (a *ACMPCAStub) DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	var output acmpca.DeleteCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.DeleteCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *AcmpcaDeleteCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.DeleteCertificateAuthority", input)
	return &AcmpcaDeleteCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
	var output acmpca.DeletePermissionOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.DeletePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *AcmpcaDeletePermissionResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.DeletePermission", input)
	return &AcmpcaDeletePermissionResult{Result: future}
}

func (a *ACMPCAStub) DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
	var output acmpca.DeletePolicyOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.DeletePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *AcmpcaDeletePolicyResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.DeletePolicy", input)
	return &AcmpcaDeletePolicyResult{Result: future}
}

func (a *ACMPCAStub) DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	var output acmpca.DescribeCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.DescribeCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *AcmpcaDescribeCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.DescribeCertificateAuthority", input)
	return &AcmpcaDescribeCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.DescribeCertificateAuthorityAuditReportOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.DescribeCertificateAuthorityAuditReport", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *AcmpcaDescribeCertificateAuthorityAuditReportResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.DescribeCertificateAuthorityAuditReport", input)
	return &AcmpcaDescribeCertificateAuthorityAuditReportResult{Result: future}
}

func (a *ACMPCAStub) GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
	var output acmpca.GetCertificateOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.GetCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *AcmpcaGetCertificateResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.GetCertificate", input)
	return &AcmpcaGetCertificateResult{Result: future}
}

func (a *ACMPCAStub) GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	var output acmpca.GetCertificateAuthorityCertificateOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.GetCertificateAuthorityCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *AcmpcaGetCertificateAuthorityCertificateResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.GetCertificateAuthorityCertificate", input)
	return &AcmpcaGetCertificateAuthorityCertificateResult{Result: future}
}

func (a *ACMPCAStub) GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	var output acmpca.GetCertificateAuthorityCsrOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.GetCertificateAuthorityCsr", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *AcmpcaGetCertificateAuthorityCsrResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.GetCertificateAuthorityCsr", input)
	return &AcmpcaGetCertificateAuthorityCsrResult{Result: future}
}

func (a *ACMPCAStub) GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
	var output acmpca.GetPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.GetPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *AcmpcaGetPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.GetPolicy", input)
	return &AcmpcaGetPolicyResult{Result: future}
}

func (a *ACMPCAStub) ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	var output acmpca.ImportCertificateAuthorityCertificateOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.ImportCertificateAuthorityCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *AcmpcaImportCertificateAuthorityCertificateResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.ImportCertificateAuthorityCertificate", input)
	return &AcmpcaImportCertificateAuthorityCertificateResult{Result: future}
}

func (a *ACMPCAStub) IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
	var output acmpca.IssueCertificateOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.IssueCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *AcmpcaIssueCertificateResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.IssueCertificate", input)
	return &AcmpcaIssueCertificateResult{Result: future}
}

func (a *ACMPCAStub) ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	var output acmpca.ListCertificateAuthoritiesOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.ListCertificateAuthorities", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *AcmpcaListCertificateAuthoritiesResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.ListCertificateAuthorities", input)
	return &AcmpcaListCertificateAuthoritiesResult{Result: future}
}

func (a *ACMPCAStub) ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
	var output acmpca.ListPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.ListPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *AcmpcaListPermissionsResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.ListPermissions", input)
	return &AcmpcaListPermissionsResult{Result: future}
}

func (a *ACMPCAStub) ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
	var output acmpca.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *AcmpcaListTagsResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.ListTags", input)
	return &AcmpcaListTagsResult{Result: future}
}

func (a *ACMPCAStub) PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
	var output acmpca.PutPolicyOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.PutPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *AcmpcaPutPolicyResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.PutPolicy", input)
	return &AcmpcaPutPolicyResult{Result: future}
}

func (a *ACMPCAStub) RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	var output acmpca.RestoreCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.RestoreCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *AcmpcaRestoreCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.RestoreCertificateAuthority", input)
	return &AcmpcaRestoreCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
	var output acmpca.RevokeCertificateOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.RevokeCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *AcmpcaRevokeCertificateResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.RevokeCertificate", input)
	return &AcmpcaRevokeCertificateResult{Result: future}
}

func (a *ACMPCAStub) TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
	var output acmpca.TagCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.TagCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *AcmpcaTagCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.TagCertificateAuthority", input)
	return &AcmpcaTagCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
	var output acmpca.UntagCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.UntagCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *AcmpcaUntagCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.UntagCertificateAuthority", input)
	return &AcmpcaUntagCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	var output acmpca.UpdateCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "ACMPCA.UpdateCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *ACMPCAStub) UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *AcmpcaUpdateCertificateAuthorityResult {
	future := workflow.ExecuteActivity(ctx, "ACMPCA.UpdateCertificateAuthority", input)
	return &AcmpcaUpdateCertificateAuthorityResult{Result: future}
}

func (a *ACMPCAStub) WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
	return workflow.ExecuteActivity(ctx, "ACMPCA.WaitUntilAuditReportCreated", input).Get(ctx, nil)
}

func (a *ACMPCAStub) WaitUntilAuditReportCreatedAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ACMPCA.WaitUntilAuditReportCreated", input)
}

func (a *ACMPCAStub) WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error {
	return workflow.ExecuteActivity(ctx, "ACMPCA.WaitUntilCertificateAuthorityCSRCreated", input).Get(ctx, nil)
}

func (a *ACMPCAStub) WaitUntilCertificateAuthorityCSRCreatedAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ACMPCA.WaitUntilCertificateAuthorityCSRCreated", input)
}

func (a *ACMPCAStub) WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error {
	return workflow.ExecuteActivity(ctx, "ACMPCA.WaitUntilCertificateIssued", input).Get(ctx, nil)
}

func (a *ACMPCAStub) WaitUntilCertificateIssuedAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "ACMPCA.WaitUntilCertificateIssued", input)
}
