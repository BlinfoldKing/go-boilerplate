package productdocument

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.ProductDocument) error
	SaveBatch([]entity.ProductDocument) error
	DeleteByID(id string) error
	DeleteByProductID(productID string) error
	FindByID(id string) (entity.ProductDocument, error)
	Update(id string, changeset entity.ProductDocumentChangeSet) error
	GetList(pagination entity.Pagination) (ProductDocuments []entity.ProductDocument, count int, err error)
}
