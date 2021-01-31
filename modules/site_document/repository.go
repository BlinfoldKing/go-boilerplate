package sitedocument

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.SiteDocument) error
	SaveBatch([]entity.SiteDocument) error
	DeleteByID(id string) error
	FindByID(id string) (entity.SiteDocument, error)
	Update(id string, changeset entity.SiteDocumentChangeSet) error
	GetList(pagination entity.Pagination) (SiteDocuments []entity.SiteDocument, count int, err error)
}
