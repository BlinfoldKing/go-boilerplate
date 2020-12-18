package documents

import "go-boilerplate/entity"

// StorageRepository abstraction for document storage
type StorageRepository interface {
	Save(entity.Document) error
	FindByID(id string) (entity.Document, error)
	FindByObjectBucketName(objectName string, bucketName string) (entity.Document, error)
}

// FileRepository is abstraction for document file
type FileRepository interface {
	GeneratePutURL(objectName, bucketName string) (stringURL string, err error)
	GenerateGetURL(objectName, bucketName string) (stringURL string, err error)
}
