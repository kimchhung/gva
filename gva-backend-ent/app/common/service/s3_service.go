package service

import (
	"context"
	"fmt"
	"io"

	"github.com/gva/env"
	"github.com/gva/internal/bootstrap/database"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Service struct {
	cfg      *env.Config
	db       *database.Database
	s3Client *minio.Client
}

func NewS3Service(cfg *env.Config, db *database.Database) *S3Service {
	return &S3Service{
		cfg: cfg,
		db:  db,
	}
}

// #initialize S3 bucket
func (s *S3Service) InitClient() error {
	s3Cfg := s.cfg.S3
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(s3Cfg.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(s3Cfg.AccessKey, s3Cfg.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}
	s.s3Client = minioClient
	return nil
}

type (
	PutS3Option func(object *minio.PutObjectOptions)
	GetS3Option func(object *minio.GetObjectOptions)
)

func (s *S3Service) PutObject(
	ctx context.Context,
	filePath string,
	objectName string, reader io.Reader, objectSize int64,
	opts ...PutS3Option,
) (*minio.UploadInfo, error) {
	if err := s.InitClient(); err != nil {
		return nil, err
	}

	object := &minio.PutObjectOptions{}
	for _, o := range opts {
		o(object)
	}

	info, err := s.s3Client.PutObject(ctx,
		s.cfg.S3.BucketName,
		objectName, reader, objectSize,
		*object,
	)

	if err != nil {
		return nil, fmt.Errorf("PutObject err: %v", err)
	}

	return &info, nil
}

func (s *S3Service) GetObject(
	ctx context.Context,
	objectName string,
	opt GetS3Option,
	opts ...GetS3Option,
) (*minio.Object, error) {
	getOpt := &minio.GetObjectOptions{}
	opt(getOpt)

	for _, o := range opts {
		o(getOpt)
	}

	object, err := s.s3Client.GetObject(ctx, s.cfg.S3.BucketName, objectName, *getOpt)
	if err != nil {
		return nil, fmt.Errorf("GetObject err: %v", err)
	}

	return object, nil
}
