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
	siteasset "go-boilerplate/modules/site_asset"
	"go-boilerplate/modules/users"
	workorderasset "go-boilerplate/modules/work_order_asset"
	workorderdocument "go-boilerplate/modules/work_order_document"
	"time"

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
	siteAsset          siteasset.Service
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
	siteAsset := siteasset.InitSiteAssetService(adapters)

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
		siteAsset,
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
	siteAsset siteasset.Service,
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
		siteAsset,
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
	noOrder string,
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
		noOrder,
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

// VerifyInstallation update work_order
func (service Service) VerifyInstallation(id, userid string) (workOrderGroup entity.WorkOrderGroup, err error) {
	now := time.Now()
	err = service.repository.Update(id, entity.WorkOrderChangeSet{
		VerifiedBy: &userid,
		VerifiedAt: &now,
		Status:     entity.InstallationComplete,
	})
	if err != nil {
		return entity.WorkOrderGroup{}, err
	}
	return service.GetByID(id)
}

// RequestMutation update work_order
func (service Service) RequestMutation(id, userid string) (err error) {
	wo, err := service.GetByID(id)
	if err != nil {
		return
	}

	now := time.Now()
	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status:              entity.InstallationRevision,
		MutationRequestedBy: &userid,
		MutationRequestedAt: &now,
	})

	body, _ := json.Marshal(structs.Map(wo))
	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "mutation request",
		Subtitle: "mutation request",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "mutation request",
			Subtitle: "mutation request",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return
}

// DeclineMutation update work_order
func (service Service) DeclineMutation(id string) (wo entity.WorkOrderGroup, err error) {
	wo, err = service.GetByID(id)
	if err != nil {
		return
	}

	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status: entity.InstallationCheckin,
	})

	body, _ := json.Marshal(structs.Map(wo))
	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "mutation request",
		Subtitle: "mutation request",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "mutation request",
			Subtitle: "mutation request",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return
}

// ApproveMutation update work_order
func (service Service) ApproveMutation(id, userid string) (wo entity.WorkOrderGroup, err error) {
	wo, err = service.GetByID(id)
	if err != nil {
		return
	}

	assets, err := service.workOrderAssets.GetAllByWorkorderID(wo.ID)
	if err != nil {
		return
	}

	var prevSiteID *string = nil
	for _, asset := range assets {
		oldSite, err := service.siteAsset.GetByAssetID(asset.ID)
		if err == nil {
			prevSiteID = &oldSite.SiteID
		}

		_, err = service.siteAsset.CreateAssetSite(asset.AssetID, wo.Site.ID)
		if err != nil {
			return wo, err
		}
	}

	now := time.Now()
	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status:             entity.InstallationInstalling,
		MutationApprovedBy: &userid,
		MutationApprovedAt: &now,
		PreviousSiteID:     prevSiteID,
	})

	body, _ := json.Marshal(structs.Map(wo))
	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "mutation approved",
		Subtitle: "mutation approved",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "mutation approved",
			Subtitle: "mutation approved",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return service.GetByID(id)
}

// RequestAssestment update work_order
func (service Service) RequestAssestment(id string) (wo entity.WorkOrderGroup, err error) {
	wo, err = service.GetByID(id)
	if err != nil {
		return entity.WorkOrderGroup{}, err
	}

	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status: entity.AssestmentRevision,
	})

	body, _ := json.Marshal(structs.Map(wo))
	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "assestment requested",
		Subtitle: "assestment requested",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "assestment requested",
			Subtitle: "assestment requested",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// DeclineAssestment update work_order
func (service Service) DeclineAssestment(id, userid string) (wo entity.WorkOrderGroup, err error) {
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
		Title:    "assestment declined",
		Subtitle: "assestment declined",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "assestment declined",
			Subtitle: "assestment declined",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// ApproveAssestment update work_order
func (service Service) ApproveAssestment(id, userid string) (wo entity.WorkOrderGroup, err error) {
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

	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status: entity.AssestmentComplete,
	})

	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "assestment approved",
		Subtitle: "assestment approved",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "assestment approved",
			Subtitle: "assestment approved",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// RequestAudit update work_order
func (service Service) RequestAudit(id string) (wo entity.WorkOrderGroup, err error) {
	wo, err = service.GetByID(id)
	if err != nil {
		return entity.WorkOrderGroup{}, err
	}

	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status: entity.AuditRevision,
	})

	body, _ := json.Marshal(structs.Map(wo))
	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "audit requested",
		Subtitle: "audit requested",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "audit requested",
			Subtitle: "audit requested",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// DeclineAudit update work_order
func (service Service) DeclineAudit(id, userid string) (wo entity.WorkOrderGroup, err error) {
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
		Title:    "audit declined",
		Subtitle: "audit declined",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "audit declined",
			Subtitle: "audit declined",
			URLLink:  "",
			Body:     string(body),
		})
	}

	return wo, err
}

// ApproveAudit update work_order
func (service Service) ApproveAudit(id, userid string) (wo entity.WorkOrderGroup, err error) {
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

	service.repository.Update(id, entity.WorkOrderChangeSet{
		Status: entity.AuditComplete,
	})

	notifications.PublishToQueue(notifications.Message{
		UserID:   wo.PICID,
		Title:    "audit approved",
		Subtitle: "audit approved",
		URLLink:  "",
		Body:     string(body),
	})

	for _, user := range wo.User {
		notifications.PublishToQueue(notifications.Message{
			UserID:   user.ID,
			Title:    "audit approved",
			Subtitle: "audit approved",
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

	var (
		mutationRequester *entity.User
		mutationApprover  *entity.User
		verifier          *entity.User
	)
	if workOrder.MutationRequestedBy != nil {
		user, err := service.users.GetByID(*workOrder.MutationRequestedBy)
		if err == nil {
			mutationRequester = &user.User
		}
	}

	if workOrder.MutationApprovedBy != nil {
		user, err := service.users.GetByID(*workOrder.MutationApprovedBy)
		if err == nil {
			mutationApprover = &user.User
		}
	}

	if workOrder.VerifiedBy != nil {
		user, err := service.users.GetByID(*workOrder.VerifiedBy)
		if err == nil {
			verifier = &user.User
		}

	}

	documents, err := service.documents.GetByWorkOrderID(workOrder.ID)
	return entity.WorkOrderGroup{
		WorkOrder:               workOrder,
		User:                    users,
		Asset:                   assets,
		Document:                documents,
		MutationRequestedByUser: mutationApprover,
		MutationApprovedByUser:  mutationRequester,
		VerifiedByUser:          verifier,
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
