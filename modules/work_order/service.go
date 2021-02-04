package workorder

import (
	"encoding/json"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/documents"
	"go-boilerplate/modules/history"
	involveduser "go-boilerplate/modules/involved_user"
	"go-boilerplate/modules/notifications"
	"go-boilerplate/modules/site"
	"go-boilerplate/modules/users"
	workorderasset "go-boilerplate/modules/work_order_asset"
	workorderdocument "go-boilerplate/modules/work_order_document"

	"github.com/fatih/structs"
)

// Service contains business logic
type Service struct {
	repository         Repository
	assets             asset.Service
	documents          documents.Service
	involvedUsers      involveduser.Service
	users              users.Service
	workOrderAssets    workorderasset.Service
	workOrderDocuments workorderdocument.Service
	sites              site.Service
	history            history.Service
}

// InitWorkOrderService is used to init work order service
func InitWorkOrderService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)

	assetService := asset.InitAssetService(adapters)
	documentService := documents.InitDocumentsService(adapters)
	involvedUserService := involveduser.InitInvolvedUserService(adapters)
	userService := users.InitUserService(adapters)
	workOrderAssetService := workorderasset.InitWorkOrderAssetService(adapters)
	workOrderDocumentService := workorderdocument.InitWorkOrderDocumentService(adapters)
	siteService := site.InitSiteService(adapters)
	historyService := history.InitHistoryService(adapters)

	return CreateService(
		repository,
		assetService,
		documentService,
		involvedUserService,
		userService,
		workOrderAssetService,
		workOrderDocumentService,
		siteService,
		historyService,
	)
}

// CreateService init service
func CreateService(
	repo Repository,
	assets asset.Service,
	documents documents.Service,
	involvedUsers involveduser.Service,
	users users.Service,
	workOrderAssets workorderasset.Service,
	workorderDocuments workorderdocument.Service,
	sites site.Service,
	histories history.Service,
) Service {
	return Service{
		repo,
		assets,
		documents,
		involvedUsers,
		users,
		workOrderAssets,
		workorderDocuments,
		sites,
		histories,
	}
}

func (service Service) mapWorkOrdersToWorkOrderGroups(workOrders []entity.WorkOrder) (workOrderGroups []entity.WorkOrderGroup, err error) {
	for _, workOrder := range workOrders {
		users, err := service.users.GetByWorkOrderID(workOrder.ID)
		if err != nil {
			return []entity.WorkOrderGroup{}, err
		}

		assets, err := service.assets.GetByWorkOrderID(workOrder.ID)
		if err != nil {
			return []entity.WorkOrderGroup{}, err
		}

		documents, err := service.documents.GetByWorkOrderID(workOrder.ID)
		if err != nil {
			return []entity.WorkOrderGroup{}, err
		}

		var site *entity.Site
		if workOrder.SiteID != nil {
			s, err := service.sites.GetByID(*workOrder.SiteID)
			if err != nil {
				return []entity.WorkOrderGroup{}, err
			}

			site = &s
		}

		workOrderGroup := entity.WorkOrderGroup{
			WorkOrder: workOrder,
			User:      users,
			Asset:     assets,
			Document:  documents,
			Site:      site,
		}
		workOrderGroups = append(workOrderGroups, workOrderGroup)
	}
	return
}

// CreateWorkOrder create new work_order
func (service Service) CreateWorkOrder(
	picID string,
	siteID *string,
	name,
	description string,
	workOrderType entity.WorkOrderType,
	involvedIDs []string,
	status int,
	assets []struct {
		ID  string `json:"id" validate:"required"`
		Qty int    `json:"qty" validate:"required"`
	},
	documentIDs []string,
) (workOrder entity.WorkOrder, err error) {
	workOrder, err = entity.NewWorkOrder(
		picID,
		siteID,
		name,
		description,
		workOrderType,
		status,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(workOrder)
	if err != nil {
		return
	}

	_, err = service.workOrderAssets.CreateBatchWorkOrderAssets(workOrder.ID, assets)
	if err != nil {
		return
	}

	_, err = service.workOrderDocuments.CreateBatchWorkOrderDocuments(workOrder.ID, documentIDs)
	if err != nil {
		return
	}

	_, err = service.involvedUsers.CreateBatchInvolvedUsers(workOrder.ID, involvedIDs)
	return
}

// GetList get list of work_order
func (service Service) GetList(pagination entity.Pagination) (workOrderGroups []entity.WorkOrderGroup, count int, err error) {
	workOrders, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	workOrderGroups, err = service.mapWorkOrdersToWorkOrderGroups(workOrders)
	return
}

// Update update work_order
func (service Service) Update(id string, changeset entity.WorkOrderChangeSet) (workOrderGroup entity.WorkOrderGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrderGroup{}, err
	}
	return service.GetByID(id)
}

// DeclineMutation update work_order
func (service Service) DeclineMutation(id string) (workOrderGroup entity.WorkOrderGroup, err error) {
	wo, _ := service.GetByID(id)
	body, _ := json.Marshal(structs.Map(wo))

	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "Site mutation declined",
		Subtitle: "Site mutation declined",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "Site mutation declined",
			Subtitle: "Site mutation declined",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// ApproveMutation update work_order
func (service Service) ApproveMutation(id, siteid string) (workOrderGroup entity.WorkOrderGroup, err error) {
	err = service.repository.Update(id, entity.WorkOrderChangeSet{
		SiteID: &siteid,
	})

	if err != nil {
		return entity.WorkOrderGroup{}, err
	}

	wo, _ := service.GetByID(id)
	body, _ := json.Marshal(structs.Map(wo))

	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "Site mutation approved",
		Subtitle: "Site mutation approved",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "Site mutation approved",
			Subtitle: "Site mutation approved",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// DeclineAsset update work_order
func (service Service) DeclineAsset(id, userid string) (wo entity.WorkOrderGroup, err error) {
	wo, err = service.GetByID(id)
	if err != nil {
		return entity.WorkOrderGroup{}, err
	}

	body, _ := json.Marshal(structs.Map(wo))

	for _, asset := range wo.Asset {
		_, err := service.history.CreateHistory(userid, asset.ID, "declined", "declined", float64(asset.PurchasePrice), []string{})
		if err != nil {
			return wo, err
		}
	}

	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "asset declined",
		Subtitle: "asset declined",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "asset declined",
			Subtitle: "asset declined",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// ApproveAsset update work_order
func (service Service) ApproveAsset(id, userid string) (wo entity.WorkOrderGroup, err error) {
	wo, err = service.GetByID(id)
	if err != nil {
		return entity.WorkOrderGroup{}, err
	}

	body, _ := json.Marshal(structs.Map(wo))

	for _, asset := range wo.Asset {
		_, err := service.history.CreateHistory(userid, asset.ID, "approved", "approved", float64(asset.PurchasePrice), []string{})
		if err != nil {
			return wo, err
		}
	}

	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "asset approved",
		Subtitle: "asset approved",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "asset approved",
			Subtitle: "asset approved",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// GetByID find work_orderby id
func (service Service) GetByID(id string) (workOrderGroup entity.WorkOrderGroup, err error) {
	workOrder, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	users, err := service.users.GetByWorkOrderID(workOrder.ID)
	if err != nil {
		return
	}

	assets, err := service.assets.GetByWorkOrderID(workOrder.ID)
	if err != nil {
		return
	}

	documents, err := service.documents.GetByWorkOrderID(workOrder.ID)
	return entity.WorkOrderGroup{
		WorkOrder: workOrder,
		User:      users,
		Asset:     assets,
		Document:  documents,
	}, err
}

// DeleteByID delete work_orderby id
func (service Service) DeleteByID(id string) (err error) {
	err = service.repository.DeleteByID(id)
	if err != nil {
		return
	}

	err = service.workOrderAssets.DeleteByWorkOrderID(id)
	if err != nil {
		return
	}

	err = service.workOrderDocuments.DeleteByWorkOrderID(id)
	if err != nil {
		return
	}

	err = service.involvedUsers.DeleteByWorkOrderID(id)
	return
}
