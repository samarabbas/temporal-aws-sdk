// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/servicequotas/servicequotasiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ServiceQuotasActivities struct {
	client servicequotasiface.ServiceQuotasAPI
}

func NewServiceQuotasActivities(session *session.Session, config ...*aws.Config) *ServiceQuotasActivities {
	client := servicequotas.New(session, config...)
	return &ServiceQuotasActivities{client: client}
}

func (a *ServiceQuotasActivities) AssociateServiceQuotaTemplate(ctx context.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
	return a.client.AssociateServiceQuotaTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) DeleteServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	return a.client.DeleteServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) DisassociateServiceQuotaTemplate(ctx context.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error) {
	return a.client.DisassociateServiceQuotaTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) GetAWSDefaultServiceQuota(ctx context.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
	return a.client.GetAWSDefaultServiceQuotaWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) GetAssociationForServiceQuotaTemplate(ctx context.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
	return a.client.GetAssociationForServiceQuotaTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) GetRequestedServiceQuotaChange(ctx context.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
	return a.client.GetRequestedServiceQuotaChangeWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) GetServiceQuota(ctx context.Context, input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error) {
	return a.client.GetServiceQuotaWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	return a.client.GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) ListAWSDefaultServiceQuotas(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
	return a.client.ListAWSDefaultServiceQuotasWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) ListRequestedServiceQuotaChangeHistory(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	return a.client.ListRequestedServiceQuotaChangeHistoryWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	return a.client.ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) ListServiceQuotaIncreaseRequestsInTemplate(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	return a.client.ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) ListServiceQuotas(ctx context.Context, input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error) {
	return a.client.ListServiceQuotasWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) ListServices(ctx context.Context, input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error) {
	return a.client.ListServicesWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	return a.client.PutServiceQuotaIncreaseRequestIntoTemplateWithContext(ctx, input)
}

func (a *ServiceQuotasActivities) RequestServiceQuotaIncrease(ctx context.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
	return a.client.RequestServiceQuotaIncreaseWithContext(ctx, input)
}
