package work_order_document

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.WorkOrderDocument) error
	DeleteByID(id string) error
	FindByID(id string) (entity.WorkOrderDocument, error)
	Update(id string, changeset entity.WorkOrderDocumentChangeSet) error
	GetList(pagination entity.Pagination) (WorkOrderDocuments []entity.WorkOrderDocument, count int, err error)
}
