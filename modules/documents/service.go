package documents

import (
	"go-boilerplate/entity"

	"github.com/satori/uuid"
)

// Service contains business logic for documents
type Service struct {
	storageRepository StorageRepository
	fileRepository    FileRepository
}

// CreateService init service
func CreateService(storageRepo StorageRepository, fileRepo FileRepository) Service {
	return Service{storageRepo, fileRepo}
}

// CreateDocument create new document
func (service Service) CreateDocument(document entity.Document) (entity.Document, error) {
	document.ID = uuid.NewV4().String()
	err := service.storageRepository.Save(document)

	return document, err
}

// GetByID find document by id
func (service Service) GetByID(id string) (document entity.Document, err error) {
	return service.storageRepository.FindByID(id)
}

// GetByObjectBucketName find document by objectName and bucketName
func (service Service) GetByObjectBucketName(objectName, bucketName string) (document entity.Document, err error) {
	return service.storageRepository.FindByObjectBucketName(objectName, bucketName)
}

// UploadDocument gets the presigned put link for the object
func (service Service) UploadDocument(objectName, bucketName string) (url string, err error) {
	return service.fileRepository.GeneratePutURL(objectName, bucketName)
}

// DownloadDocument gets the presigned get link for the object
func (service Service) DownloadDocument(objectName, bucketName string) (url string, err error) {
	return service.fileRepository.GenerateGetURL(objectName, bucketName)
}
