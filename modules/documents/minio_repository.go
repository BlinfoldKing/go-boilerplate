package documents

import (
	"context"
	mn "go-boilerplate/adapters/minio"
	"go-boilerplate/config"
	"net/url"

	"github.com/minio/minio-go/v7"
)

// MinioRepository repository implementation on postgres
type MinioRepository struct {
	client *mn.Minio
}

// CreateMinioRepository init MinioRepository
func CreateMinioRepository(client *mn.Minio) FileRepository {
	return MinioRepository{client}
}

// GeneratePutURL generates presigned put url for the object
func (m MinioRepository) GeneratePutURL(objectName, bucketName string) (stringURL string, err error) {
	expiry := config.MINIOEXPIRE()
	ctx := context.Background()
	exists, err := m.client.BucketExists(ctx, bucketName)
	if !exists {
		err = m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: config.MINIOREGION()})
		if err != nil {
			return
		}
	}
	presignedURL, err := m.client.PresignedPutObject(ctx, bucketName, objectName, expiry)
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
