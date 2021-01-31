package warehouse

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/contact"
	warehousecontact "go-boilerplate/modules/warehouse_contact"
)

// Service contains business logic
type Service struct {
	repository        Repository
	contacts          contact.Service
	warehouseContacts warehousecontact.Service
}

// InitWarehouseService warehouse
func InitWarehouseService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	contactService := contact.InitContactService(adapters)
	warehouseContactService := warehousecontact.InitWarehouseContactService(adapters)
	return CreateService(repository, contactService, warehouseContactService)
}

// CreateService init service
func CreateService(repo Repository, contacts contact.Service, warehouseContacts warehousecontact.Service) Service {
	return Service{repo, contacts, warehouseContacts}
}

func (service Service) mapWarehousesToWarehouseGroups(warehouses []entity.Warehouse) (warehouseGroups []entity.WarehouseGroup, err error) {
	for _, warehouse := range warehouses {
		contacts, err := service.contacts.GetByWarehouseID(warehouse.ID)
		if err != nil {
			return []entity.WarehouseGroup{}, err
		}
		warehouseGroup := entity.WarehouseGroup{
			Warehouse: warehouse,
			Contacts:  contacts,
		}

		warehouseGroups = append(warehouseGroups, warehouseGroup)
	}
	return
}

// CreateWarehouse create new warehouse
func (service Service) CreateWarehouse(name, description, address string, latitude, longitude float64, contactIDs []string) (warehouse entity.Warehouse, err error) {
	warehouse, err = entity.NewWarehouse(name, description, address, latitude, longitude)
	if err != nil {
		return
	}
	err = service.repository.Save(warehouse)
	if err != nil {
		return
	}

	_, err = service.warehouseContacts.CreateBatchWarehouseContacts(warehouse.ID, contactIDs)
	return
}

// GetList get list of warehouse
func (service Service) GetList(pagination entity.Pagination) (warehouseGroups []entity.WarehouseGroup, count int, err error) {
	warehouses, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}

	warehouseGroups, err = service.mapWarehousesToWarehouseGroups(warehouses)
	return
}

// Update update warehouse
func (service Service) Update(id string, changeset entity.WarehouseChangeSet) (warehouse entity.WarehouseGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WarehouseGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find warehouseby id
func (service Service) GetByID(id string) (warehouseGroup entity.WarehouseGroup, err error) {
	warehouse, err := service.repository.FindByID(id)
	if err != nil {
		return
	}
	contacts, err := service.contacts.GetByWarehouseID(id)
	return entity.WarehouseGroup{
		Warehouse: warehouse,
		Contacts:  contacts,
	}, err
}

// GetAllByAssetID get all workhouse by work order id
func (service Service) GetAllByAssetID(assetID string) (assets []entity.Warehouse, err error) {
	return service.repository.GetAllWarehousebyAssetID(assetID)
}

// DeleteByID delete warehouseby id
func (service Service) DeleteByID(id string) (err error) {
	err = service.repository.DeleteByID(id)
	if err != nil {
		return
	}
	err = service.warehouseContacts.DeleteByWarehouseID(id)
	return
}
