package s3

import (
	"context"
	"fmt"
	"io"
	"path"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type awsS3Storage struct {
	cfg BucketS3Config
}

// NewS3Config creates a new instance of S3Config
func New(cfg BucketS3Config, ctx context.Context) S3Storage {
	s := &awsS3Storage{
		cfg: cfg,
	}
	return s
}
func (s *awsS3Storage) loadConfig() (*aws.Config, error) {
	result, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(s.cfg.AccessKey, s.cfg.SecretKey, ""),
		),
		config.WithRegion(s.cfg.Region),
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *awsS3Storage) getObjectKey(file string) string {
	return fmt.Sprintf("%s/%s", s.cfg.Prefix, file)
}

func (s *awsS3Storage) getStorageClient() (*s3.Client, error) {
	cfg, err := s.loadConfig()
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(*cfg), nil
}

func (s *awsS3Storage) fileUrl(acl *string, file string) string {
	if acl != nil && types.ObjectCannedACL(*acl) == types.ObjectCannedACLPublicRead {
		return fmt.Sprintf("%s/%s", s.cfg.CDNUrl, file)
	}
	return fmt.Sprintf("https://%s.%s.s3.amazonaws.com/%s", s.cfg.BucketName, s.cfg.Region, file)
}

func (s *awsS3Storage) Upload(ctx context.Context, fileName string, fileData io.ReadSeeker, acl *string) (*S3UploadResponse, error) {
	storageClient, err := s.getStorageClient()
	if err != nil {
		return nil, err
	}

	objInput := &s3.PutObjectInput{
		Bucket:             aws.String(s.cfg.BucketName),
		Key:                aws.String(s.getObjectKey(fileName)),
		Body:               fileData,
		ContentDisposition: aws.String("attachment"),
	}
	_, err = storageClient.PutObject(ctx, objInput)
	if err != nil {
		return nil, err
	}

	result := &S3UploadResponse{
		FileUrl:  s.fileUrl(acl, s.getObjectKey(fileName)),
		FileName: s.getObjectKey(fileName),
		FileExt:  path.Ext(fileName),
	}
	return result, nil
}
func (s *awsS3Storage) SignedUrl(ctx context.Context, objectUrl string, expiration time.Duration) (*string, error) {
}
