package file

import (
	"context"
	"mime/multipart"

	"github.com/gva/app/common/service"
	"github.com/gva/internal/bootstrap/database"
	"github.com/minio/minio-go/v7"
)

type FileService struct {
	db   *database.Database
	s3_s *service.S3Service
}

func NewFileService(db *database.Database, s3_s *service.S3Service) *FileService {
	return &FileService{
		db: db,
	}
}

func (s *FileService) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (any, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	info, err := s.s3_s.PutObject(ctx, "/img/",
		fileHeader.Filename,
		file,
		fileHeader.Size,
		func(object *minio.PutObjectOptions) {
			object.ContentType = fileHeader.Header.Get("Content-Type")
			object.ContentEncoding = fileHeader.Header.Get("Content-Encoding")
			object.ContentDisposition = fileHeader.Header.Get("Content-Disposition")
			object.ContentLanguage = fileHeader.Header.Get("Content-Language")
		},
	)

	return info, err
}
