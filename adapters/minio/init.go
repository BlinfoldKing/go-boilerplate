package minio

import (
	"context"
	"go-boilerplate/config"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	client *minio.Client
}

func Init() (*Minio, error) {
	client, err := minio.New(config.MINIOENDPOINT(), &minio.Options{
		Creds:  credentials.NewStaticV4(config.MINIOACCESSKEY(), config.MINIOSECRET(), ""),
		Secure: true,
	})
	return &Minio{client}, err
}

func (m Minio) GeneratePutURL(objectName, bucketName string) (url string, err error) {
	expiry := time.Second * 24 * 60 * 60
	presignedURL, err := m.client.PresignedPutObject(context.Background(), bucketName, objectName, expiry)
	url = presignedURL.String()
	return
}
