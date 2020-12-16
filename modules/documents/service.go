package documents

import (
	"go-boilerplate/entity"
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
func (service Service) CreateDocument(objectName, bucketName, presignedURL string) (entity.DocumentCreateResponse, error) {
	document, _ := entity.NewDocument(objectName, bucketName)
	err := service.repository.Save(document)

	return entity.DocumentCreateResponse{
		Document: document,
		URL:      presignedURL,
	}, err
}

// GetByID find document by id
func (service Service) GetByID(id string) (document entity.Document, err error) {
	return service.repository.FindByID(id)
}
