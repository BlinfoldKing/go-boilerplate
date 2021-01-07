package companydocument

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.CompanyDocument) error
	SaveBatch([]entity.CompanyDocument) error
	DeleteByID(id string) error
	DeleteByCompanyID(companyID string) error
	FindByID(id string) (entity.CompanyDocument, error)
	Update(id string, changeset entity.CompanyDocumentChangeSet) error
	GetList(pagination entity.Pagination) (CompanyDocuments []entity.CompanyDocument, count int, err error)
}
