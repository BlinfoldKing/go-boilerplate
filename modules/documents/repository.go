package documents

import "go-boilerplate/entity"

// Repository abstraction for document storage
type Repository interface {
	Save(entity.Document) error
	FindByID(id string) (entity.Document, error)
	FindByObjectBucketName(objectName string, bucketName string) (entity.Document, error)
}
