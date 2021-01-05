package history

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/documents"
	historydocument "go-boilerplate/modules/history_document"
	"go-boilerplate/modules/users"
)

// Service contains business logic
type Service struct {
	repository       Repository
	assets           asset.Service
	documents        documents.Service
	historyDocuments historydocument.Service
	users            users.Service
}

func InitHistoryService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)

	assetService := asset.InitAssetService(adapters)
	documentService := documents.InitDocumentsService(adapters)
	historyDocumentService := historydocument.InitHistoryDocumentService(adapters)
	userService := users.InitUserService(adapters)
	return CreateService(
		repository,
		assetService,
		documentService,
		historyDocumentService,
		userService,
	)
}

// CreateService init service
func CreateService(repo Repository,
	assets asset.Service,
	documents documents.Service,
	historyDocuments historydocument.Service,
	users users.Service,
) Service {
	return Service{
		repo,
		assets,
		documents,
		historyDocuments,
		users,
	}
}

func (service Service) mapHistoriesToHistoryGroups(histories []entity.History) (historyGroups []entity.HistoryGroup, err error) {
	for _, history := range histories {
		user, err := service.users.GetByID(history.UserID)
		if err != nil {
			return []entity.HistoryGroup{}, err
		}

		asset, err := service.assets.GetByID(history.AssetID)
		if err != nil {
			return []entity.HistoryGroup{}, err
		}

		documents, err := service.documents.GetByHistoryID(history.ID)
		if err != nil {
			return []entity.HistoryGroup{}, err
		}
		historyGroup := entity.HistoryGroup{
			History:   history,
			User:      user.User,
			Asset:     asset,
			Documents: documents,
		}

		historyGroups = append(historyGroups, historyGroup)
	}
	return
}

// CreateHistory create new history
func (service Service) CreateHistory(userID, assetID, action, description string, cost float64, documentIDs []string) (history entity.History, err error) {
	history, err = entity.NewHistory(
		userID,
		assetID,
		action,
		description,
		cost,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(history)
	if err != nil {
		return
	}

	_, err = service.historyDocuments.CreateBatchHistoryDocuments(history.ID, documentIDs)
	return
}

// GetList get list of history
func (service Service) GetList(pagination entity.Pagination) (historyGroups []entity.HistoryGroup, count int, err error) {
	histories, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	historyGroups, err = service.mapHistoriesToHistoryGroups(histories)
	return
}

// Update update history
func (service Service) Update(id string, changeset entity.HistoryChangeSet) (historyGroup entity.HistoryGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.HistoryGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find historyby id
func (service Service) GetByID(id string) (historyGroup entity.HistoryGroup, err error) {
	history, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	user, err := service.users.GetByID(history.UserID)
	if err != nil {
		return
	}

	asset, err := service.assets.GetByID(history.AssetID)
	if err != nil {
		return
	}

	documents, err := service.documents.GetByHistoryID(id)
	return entity.HistoryGroup{
		History:   history,
		User:      user.User,
		Asset:     asset,
		Documents: documents,
	}, err
}

// DeleteByID delete historyby id
func (service Service) DeleteByID(id string) (err error) {
	err = service.repository.DeleteByID(id)
	if err != nil {
		return
	}
	err = service.historyDocuments.DeleteByHistoryID(id)
	return
}
