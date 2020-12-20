package documents

import (
	"context"
	"go-boilerplate/adapters/minio"
	"go-boilerplate/config"
	"net/url"
)

// MinioRepository repository implementation on postgres
type MinioRepository struct {
	client *minio.Minio
}

// CreateMinioRepository init MinioRepository
func CreateMinioRepository(client *minio.Minio) FileRepository {
	return MinioRepository{client}
}

// GeneratePutURL generates presigned put url for the object
func (m MinioRepository) GeneratePutURL(objectName, bucketName string) (stringURL string, err error) {
	expiry := config.MINIOEXPIRE()
	presignedURL, err := m.client.PresignedPutObject(context.Background(), bucketName, objectName, expiry)
	stringURL = presignedURL.String()
	return
}

// GenerateGetURL generates presigned get url for the object
func (m MinioRepository) GenerateGetURL(objectName, bucketName string) (stringURL string, err error) {
	expiry := config.MINIOEXPIRE()
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename="+objectName)
	presignedURL, err := m.client.PresignedGetObject(context.Background(), bucketName, objectName, expiry, reqParams)
	stringURL = presignedURL.String()
	return
}
