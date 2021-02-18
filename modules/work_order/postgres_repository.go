package workorder

import (
	"go-boilerplate/adapters/postgres"
	"go-boilerplate/entity"
	"time"
)

// PostgresRepository repository implementation on postgres
type PostgresRepository struct {
	db *postgres.Postgres
}

// CreatePostgresRepository init PostgresRepository
func CreatePostgresRepository(db *postgres.Postgres) Repository {
	return PostgresRepository{db}
}

// Save save work_order to db
func (repo PostgresRepository) Save(workOrder entity.WorkOrder) error {
	_, err := repo.db.Table("work_orders").Insert(&workOrder)
	return err
}

// ApproveAudit approve audit
func (repo PostgresRepository) ApproveAudit(wo entity.WorkOrderGroup, userid string) error {
	sess := repo.db.Engine.NewSession()

	err := sess.Begin()
	if err != nil {
		return err
	}

	for _, asset := range wo.Asset {
		history, _ := entity.NewHistory(
			userid,
			asset.ID,
			"audit approve",
			"audit approve",
			float64(asset.PurchasePrice),
		)
		_, err = sess.Table("histories").Insert(&history)
		if err != nil {
			sess.Rollback()
			return err
		}
	}

	_, err = repo.db.Table("work_orders").Where("id = ?", wo.ID).Update(&entity.WorkOrderChangeSet{
		Status: entity.AssestmentComplete,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	err = sess.Commit()
	if err != nil {
		return err
	}
	return nil
}

// ApproveAssestment approve assestment
func (repo PostgresRepository) ApproveAssestment(wo entity.WorkOrderGroup, userid string) error {
	sess := repo.db.Engine.NewSession()

	err := sess.Begin()
	if err != nil {
		return err
	}

	for _, asset := range wo.Asset {
		history, _ := entity.NewHistory(
			userid,
			asset.ID,
			"assestment approve",
			"assestment approve",
			float64(asset.PurchasePrice),
		)
		_, err = sess.Table("histories").Insert(&history)
		if err != nil {
			sess.Rollback()
			return err
		}
	}

	_, err = repo.db.Table("work_orders").Where("id = ?", wo.ID).Update(&entity.WorkOrderChangeSet{
		Status: entity.AssestmentComplete,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	err = sess.Commit()
	if err != nil {
		return err
	}
	return nil
}

// ApproveMutationV2 approve mutation
func (repo PostgresRepository) ApproveMutationV2(wo entity.WorkOrderGroup, userid string) error {
	sess := repo.db.Engine.NewSession()

	err := sess.Begin()
	if err != nil {
		return err
	}

	for _, asset := range wo.Asset {
		sa, _ := entity.NewSiteAsset(asset.ID, *wo.NextSiteID)
		_, err = sess.Table("site_assets").Insert(&sa)
		if err != nil {
			sess.Rollback()
			return err
		}
	}

	now := time.Now()
	_, err = repo.db.Table("work_orders").Where("id = ?", wo.ID).Update(&entity.WorkOrderChangeSet{
		Status:             entity.InstallationDeliveryCheckpoint,
		MutationApprovedBy: &userid,
		MutationApprovedAt: &now,
		PreviousSiteID:     wo.SiteID,
		SiteID:             wo.NextSiteID,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	err = sess.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Create save work_order to db
func (repo PostgresRepository) Create(
	workOrder entity.WorkOrder,
	involvedIDs *[]string,
	documentIDs *[]string,
	assets *[]struct {
		ID  string `json:"id" validate:"required"`
		Qty int    `json:"qty" validate:"required"`
	},
) error {
	sess := repo.db.Engine.NewSession()

	err := sess.Begin()
	if err != nil {
		return err
	}

	_, err = sess.Table("work_orders").Insert(&workOrder)
	if err != nil {
		sess.Rollback()
		return err
	}

	if assets != nil {
		for _, asset := range *assets {
			woasset, _ := entity.
				NewWorkOrderAsset(workOrder.ID, asset.ID, asset.Qty)
			_, err = sess.
				Table("work_order_assets").
				Insert(&woasset)

			if err != nil {
				sess.Rollback()
				return err
			}
		}
	}

	if involvedIDs != nil {
		for _, id := range *involvedIDs {
			involveduser, _ := entity.
				NewInvolvedUser(id, workOrder.ID)
			_, err = sess.
				Table("involved_users").
				Insert(&involveduser)

			if err != nil {
				sess.Rollback()
				return err
			}
		}
	}

	if documentIDs != nil {
		for _, id := range *documentIDs {
			document, _ := entity.
				NewWorkOrderDocument(workOrder.ID, id)
			_, err = sess.
				Table("work_order_documents").
				Insert(&document)

			if err != nil {
				sess.Rollback()
				return err
			}
		}
	}

	err = sess.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetList get list of work_order
func (repo PostgresRepository) GetList(pagination entity.Pagination) (workOrders []entity.WorkOrder, count int, err error) {
	count, err = repo.db.
		Paginate("work_orders", &workOrders, pagination)
	return
}

// Update update work_order
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderChangeSet) error {
	_, err := repo.db.Table("work_orders").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order by id
func (repo PostgresRepository) FindByID(id string) (workOrder entity.WorkOrder, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_orders WHERE id = ? AND deleted_at IS null", id).Get(&workOrder)
	return
}

// DeleteByID delete work_order by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_orders").Where("id = ?", id).Delete(&entity.WorkOrder{})
	return err
}
