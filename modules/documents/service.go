package documents

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"time"

	"github.com/satori/uuid"
)

// Service contains business logic for documents
type Service struct {
	storageRepository StorageRepository
	fileRepository    FileRepository
}

// InitDocumentsService init document service
func InitDocumentsService(adapters adapters.Adapters) Service {
	storageRepository := CreatePostgresRepository(adapters.Postgres)
	fileRepository := CreateMinioRepository(adapters.Minio)
	return CreateService(storageRepository, fileRepository)
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

// GetByProductID finds
func (service Service) GetByProductID(productID string) (documents []entity.Document, err error) {
	return service.storageRepository.FindByProductID(productID)
}

// GetByHistoryID finds document by history ID
func (service Service) GetByHistoryID(historyID string) (documents []entity.Document, err error) {
	return service.storageRepository.FindByHistoryID(historyID)
}

// GetByCompanyID finds document by company ID
func (service Service) GetByCompanyID(companyID string) (documents []entity.Document, err error) {
	return service.storageRepository.FindByCompanyID(companyID)
}

// GetByWorkOrderID finds document by work order ID
func (service Service) GetByWorkOrderID(workOrderID string) (documents []entity.Document, err error) {
	return service.storageRepository.FindByWorkOrderID(workOrderID)
}

// GetBySiteID finds document by site ID
func (service Service) GetBySiteID(siteID string) (documents []entity.Document, err error) {
	return service.storageRepository.FindBySiteID(siteID)
}

// GetByObjectBucketName find document by objectName and bucketName
func (service Service) GetByObjectBucketName(objectName, bucketName string) (document entity.Document, err error) {
	return service.storageRepository.FindByObjectBucketName(objectName, bucketName)
}

// UploadDocument gets the presigned put link for the object
func (service Service) UploadDocument(objectName, bucketName string) (url string, err error) {
	objectName = fmt.Sprintf("%d", time.Now().Unix()) + "-" + objectName
	return service.fileRepository.GeneratePutURL(objectName, bucketName)
}

// DownloadDocument gets the presigned get link for the object
func (service Service) DownloadDocument(objectName, bucketName string) (url string, err error) {
	return service.fileRepository.GenerateGetURL(objectName, bucketName)
}
