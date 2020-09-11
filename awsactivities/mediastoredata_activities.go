
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"github.com/aws/aws-sdk-go/service/mediastoredata/mediastoredataiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MediaStoreDataActivities struct {
    client mediastoredataiface.MediaStoreDataAPI
}

func NewMediaStoreDataActivities(session *session.Session, config ...*aws.Config) *MediaStoreDataActivities {
    client := mediastoredata.New(session, config...)
    return &MediaStoreDataActivities{client: client}
}

func (a *MediaStoreDataActivities) DeleteObject(ctx context.Context, input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
    return a.client.DeleteObjectWithContext(ctx, input)
}

func (a *MediaStoreDataActivities) DescribeObject(ctx context.Context, input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
    return a.client.DescribeObjectWithContext(ctx, input)
}

func (a *MediaStoreDataActivities) GetObject(ctx context.Context, input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
    return a.client.GetObjectWithContext(ctx, input)
}

func (a *MediaStoreDataActivities) ListItems(ctx context.Context, input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
    return a.client.ListItemsWithContext(ctx, input)
}

func (a *MediaStoreDataActivities) PutObject(ctx context.Context, input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error) {
    return a.client.PutObjectWithContext(ctx, input)
}
