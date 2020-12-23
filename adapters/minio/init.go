package minio

import (
	"go-boilerplate/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Minio extends minio
type Minio struct {
	*minio.Client
}

// Init creates minio client
func Init() (*Minio, error) {
	client, err := minio.New(config.MINIOENDPOINT(), &minio.Options{
		Creds:  credentials.NewStaticV4(config.MINIOACCESSKEY(), config.MINIOSECRET(), ""),
		Secure: true,
	})
	return &Minio{client}, err
}
