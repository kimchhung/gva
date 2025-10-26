package service

import (
	coreerror "backend/core/error"
	"backend/env"
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

const (
	maxFileSizeMB = 5 // 5MB
)

var allowedExtensions = []string{".png", ".jpeg", ".jpg", ".gif", ".webp"}

type S3Service struct {
	// S3 client
	Client *s3.Client

	cfg *env.Config
	log *zap.Logger
}

func NewS3Service(cfg *env.Config, log *zap.Logger) *S3Service {
	config, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.S3.Key, cfg.S3.Secret, cfg.S3.Session),
		),
		awsConfig.WithRegion(cfg.S3.Region),
	)
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(config)
	return &S3Service{
		Client: s3Client,
		cfg:    cfg,
		log:    log.Named("s3"),
	}
}

func (s *S3Service) UploadFile(file *multipart.FileHeader) (*UploadObject, error) {
	// Validate the image
	validatedImage, err := validateImage(file, maxFileSizeMB)
	if err != nil {
		return nil, err
	}

	imageCopy := validatedImage.File
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, imageCopy); err != nil {
		return nil, err
	}

	_, err = imageCopy.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Generate a unique file name
	fileNameByte := sha1.Sum(buf.Bytes())
	fileName := hex.EncodeToString(fileNameByte[:])
	fileName = fmt.Sprintf("%s%s", fileName, validatedImage.Extension)

	// Put the object in the bucket
	_, err = s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.cfg.S3.Bucket),
		Key:         aws.String(fileName),
		Body:        io.ReadSeeker(validatedImage.File),
		ContentType: aws.String(file.Header.Get("Content-Type")),
		Metadata: map[string]string{
			"filename": file.Filename,
			"size":     fmt.Sprintf("%d", file.Size),
		},
	})
	if err != nil {
		return nil, err
	}

	return &UploadObject{
		URL:      fmt.Sprintf("%s/%s", s.cfg.S3.Endpoint, fileName),
		Filename: fileName,
	}, nil
}

type UploadObject struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

type Image struct {
	File      multipart.File
	Extension string
}

func validateImage(file *multipart.FileHeader, maxFileSizeMB int64) (*Image, error) {
	fileData, err := file.Open()
	if err != nil {
		return nil, coreerror.ErrWhileUploading
	}
	defer fileData.Close()

	// Get file extension from header
	ext := strings.ToLower(path.Ext(file.Filename))

	// Check if file extension is allowed
	if !slices.Contains(allowedExtensions, ext) {
		return nil, coreerror.ErrUnsupportedFileFormat
	}

	// Check if file size limit
	if file.Size > maxFileSizeMB*1024*1024 {
		return nil, coreerror.ErrImageTooLarge
	}

	// Read the first 512 bytes of file to check mimetype
	buf := make([]byte, 512)
	_, err = fileData.Read(buf)
	if err != nil {
		return nil, coreerror.ErrWhileUploading
	}

	mimeType := http.DetectContentType(buf)
	if !strings.HasPrefix(mimeType, "image/") {
		return nil, fmt.Errorf("invalid mime type: %s", mimeType)
	}

	// Rewind the file pointer to the beginning of the file
	_, err = fileData.Seek(0, io.SeekStart)
	if err != nil {
		return nil, coreerror.ErrWhileUploading
	}

	return &Image{
		Extension: ext,
		File:      fileData,
	}, nil
}
