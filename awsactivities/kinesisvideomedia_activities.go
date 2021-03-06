// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type KinesisVideoMediaActivities struct {
	client kinesisvideomediaiface.KinesisVideoMediaAPI
}

func NewKinesisVideoMediaActivities(session *session.Session, config ...*aws.Config) *KinesisVideoMediaActivities {
	client := kinesisvideomedia.New(session, config...)
	return &KinesisVideoMediaActivities{client: client}
}

func (a *KinesisVideoMediaActivities) GetMedia(ctx context.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	return a.client.GetMediaWithContext(ctx, input)
}
