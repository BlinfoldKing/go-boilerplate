package historydocument

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.HistoryDocument) error
	SaveBatch([]entity.HistoryDocument) error
	DeleteByID(id string) error
	DeleteByHistoryID(historyID string) error
	FindByID(id string) (entity.HistoryDocument, error)
	Update(id string, changeset entity.HistoryDocumentChangeSet) error
	GetList(pagination entity.Pagination) (HistoryDocuments []entity.HistoryDocument, count int, err error)
}
