package documents

import (
	"go-boilerplate/entity"

	"github.com/satori/uuid"
)

// Service contains business logic for documents
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateDocument create new document
func (service Service) CreateDocument(document entity.Document) (entity.Document, error) {
	document.ID = uuid.NewV4().String()
	err := service.repository.Save(document)

	return document, err
}

// GetByID find document by id
func (service Service) GetByID(id string) (document entity.Document, err error) {
	return service.repository.FindByID(id)
}

// GetByObjectBucketName find document by objectName and bucketName
func (service Service) GetByObjectBucketName(objectName, bucketName string) (document entity.Document, err error) {
	return service.repository.FindByObjectBucketName(objectName, bucketName)
}
