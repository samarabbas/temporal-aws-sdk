// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"go.temporal.io/sdk/workflow"
)

type ForecastQueryServiceClient interface {
	QueryForecast(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error)
	QueryForecastAsync(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) *ForecastqueryserviceQueryForecastResult
}

type ForecastQueryServiceStub struct{}

func NewForecastQueryServiceStub() ForecastQueryServiceClient {
	return &ForecastQueryServiceStub{}
}

type ForecastqueryserviceQueryForecastResult struct {
	Result workflow.Future
}

func (r *ForecastqueryserviceQueryForecastResult) Get(ctx workflow.Context) (*forecastqueryservice.QueryForecastOutput, error) {
	var output forecastqueryservice.QueryForecastOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *ForecastQueryServiceStub) QueryForecast(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
	var output forecastqueryservice.QueryForecastOutput
	err := workflow.ExecuteActivity(ctx, "ForecastQueryService.QueryForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *ForecastQueryServiceStub) QueryForecastAsync(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) *ForecastqueryserviceQueryForecastResult {
	future := workflow.ExecuteActivity(ctx, "ForecastQueryService.QueryForecast", input)
	return &ForecastqueryserviceQueryForecastResult{Result: future}
}
