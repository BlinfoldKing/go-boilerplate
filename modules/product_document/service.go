package productdocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

func InitProductDocumentService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateProductDocument create new product_document
func (service Service) CreateProductDocument(productID, documentID string) (productDocument entity.ProductDocument, err error) {
	productDocument, err = entity.NewProductDocument(productID, documentID)
	if err != nil {
		return
	}
	err = service.repository.Save(productDocument)
	return
}

// CreateBatchProductDocument creates a batch of new productDocuments
func (service Service) CreateBatchProductDocument(productID string, documentIDs []string) (productDocuments []entity.ProductDocument, err error) {
	for _, documentID := range documentIDs {
		productDocument, err := entity.NewProductDocument(productID, documentID)
		if err != nil {
			return []entity.ProductDocument{}, err
		}
		productDocuments = append(productDocuments, productDocument)
	}
	err = service.repository.SaveBatch(productDocuments)
	return
}

// GetList get list of product_document
func (service Service) GetList(pagination entity.Pagination) (productDocument []entity.ProductDocument, count int, err error) {
	productDocument, count, err = service.repository.GetList(pagination)
	return
}

// Update update product_document
func (service Service) Update(id string, changeset entity.ProductDocumentChangeSet) (productDocument entity.ProductDocument, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.ProductDocument{}, err
	}
	return service.GetByID(id)
}

// GetByID find product_documentby id
func (service Service) GetByID(id string) (productDocument entity.ProductDocument, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete product_documentby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByProductID delete product_document by product id
func (service Service) DeleteByProductID(productID string) (err error) {
	return service.repository.DeleteByProductID(productID)
}
