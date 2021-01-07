package historydocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

func InitHistoryDocumentService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateHistoryDocument create new history_document
func (service Service) CreateHistoryDocument(historyID, documentID string) (historyDocument entity.HistoryDocument, err error) {
	historyDocument, err = entity.NewHistoryDocument(historyID, documentID)
	if err != nil {
		return
	}
	err = service.repository.Save(historyDocument)
	return
}

// CreateBatchHistoryDocuments creates a batch of new historyDocuments
func (service Service) CreateBatchHistoryDocuments(historyID string, documentIDs []string) (historyDocuments []entity.HistoryDocument, err error) {
	for _, documentID := range documentIDs {
		historyDocument, err := entity.NewHistoryDocument(historyID, documentID)
		if err != nil {
			return []entity.HistoryDocument{}, err
		}
		historyDocuments = append(historyDocuments, historyDocument)
	}
	err = service.repository.SaveBatch(historyDocuments)
	return
}

// GetList get list of history_document
func (service Service) GetList(pagination entity.Pagination) (historyDocument []entity.HistoryDocument, count int, err error) {
	historyDocument, count, err = service.repository.GetList(pagination)
	return
}

// Update update history_document
func (service Service) Update(id string, changeset entity.HistoryDocumentChangeSet) (historyDocument entity.HistoryDocument, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.HistoryDocument{}, err
	}
	return service.GetByID(id)
}

// GetByID find history_documentby id
func (service Service) GetByID(id string) (historyDocument entity.HistoryDocument, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete history_documentby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByHistoryID delete history_document by history id
func (service Service) DeleteByHistoryID(historyID string) (err error) {
	return service.repository.DeleteByHistoryID(historyID)
}
