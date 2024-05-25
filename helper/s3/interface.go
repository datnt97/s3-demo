package s3

import (
	"context"
	"io"
	"time"
)

type S3Storage interface {
	Upload(ctx context.Context, fileName string, fileData io.ReadSeeker, isCND bool) (*S3UploadResponse, error)
	SignedUrl(ctx context.Context, objectUrl string, expiration time.Duration) (*string, error)
	CloundFrontSignUrl(ctx context.Context, rawURL string, expiration time.Duration) (*string, error)
}
type S3UploadResponse struct {
	FileUrl  string
	FileName string
	FileExt  string
}
