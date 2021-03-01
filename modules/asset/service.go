package asset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	assetwarehouseextend "go-boilerplate/modules/asset_warehouse_extend"
	"go-boilerplate/modules/company"
	"go-boilerplate/modules/product"
	siteasset "go-boilerplate/modules/site_asset"
	"go-boilerplate/modules/warehouse"
	"math"
	"time"
)

// Service contains business logic
type Service struct {
	repository      Repository
	warehouse       warehouse.Service
	company         company.Service
	product         product.Service
	siteasset       siteasset.Service
	assetWarehouses assetwarehouseextend.Service
}

// InitAssetService create new asset service
func InitAssetService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	product := product.InitProductService(adapters)
	company := company.InitCompanyService(adapters)
	warehouse := warehouse.InitWarehouseService(adapters)
	siteasset := siteasset.InitSiteAssetService(adapters)
	assetwarehouse := assetwarehouseextend.InitService(adapters)

	return Service{
		repository:      repository,
		product:         product,
		company:         company,
		warehouse:       warehouse,
		siteasset:       siteasset,
		assetWarehouses: assetwarehouse,
	}
}

func monthDiff(datetime time.Time) int {
	now := time.Now()
	months := 0
	month := datetime.Month()
	for datetime.Before(now) {
		datetime = datetime.Add(time.Hour * 24)
		nextMonth := datetime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

func (service Service) mapAssetToAssetGroup(asset entity.Asset) (ag entity.AssetGroup, err error) {
	ag.Asset = asset

	ag.Company, err = service.company.GetByID(ag.SupplierCompanyID)
	if err != nil {
		return
	}

	ag.Product, err = service.product.GetByID(ag.ProductID)
	if err != nil {
		return
	}

	ag.Warehouse, err = service.warehouse.GetAllByAssetID(asset.ID)
	if err != nil {
		return
	}

	// calculate linear valuation
	oldestDate := time.Now()
	sites, _ := service.siteasset.GetAllByAssetID(asset.ID)

	// find oldest installation date
	for _, site := range sites {
		if site.CreatedAt.Before(oldestDate) {
			oldestDate = site.CreatedAt
		}
	}

	depreciationRate := (ag.PurchasePrice - ag.SalvageValue) / float32(ag.Product.Lifetime)
	currentPrice := ag.PurchasePrice - (depreciationRate * float32(monthDiff(oldestDate)))

	if currentPrice <= ag.SalvageValue || math.IsNaN(float64(currentPrice)) {
		ag.CurrentValuation = ag.SalvageValue
	} else {
		ag.CurrentValuation = currentPrice
	}

	return
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{
		repository: repo,
	}
}

// CreateAsset create new asset
func (service Service) CreateAsset(
	productID string,
	serialNumber string,
	status int,
	purchaseDate time.Time,
	purchasePrice float32,
	supplierCompanyID string,
	salvageValue float32,
) (asset entity.Asset, err error) {
	asset, err = entity.NewAsset(
		productID,
		serialNumber,
		status,
		purchaseDate,
		purchasePrice,
		supplierCompanyID,
		salvageValue,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(asset)
	return
}

// GetList get list of asset
func (service Service) GetList(pagination entity.Pagination) (assetGroups []entity.AssetGroup, count int, err error) {
	assets, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	for _, asset := range assets {
		assetGroup, _ := service.mapAssetToAssetGroup(asset)
		assetGroups = append(assetGroups, assetGroup)
	}

	return
}

// Update update asset
func (service Service) Update(id string, changeset entity.AssetChangeSet, warehouseIDs []string) (asset entity.Asset, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return
	}

	if len(warehouseIDs) > 0 {
		err = service.assetWarehouses.DeleteByAssetID(id)
		if err != nil {
			return
		}
		_, err = service.assetWarehouses.CreateBatchAssetWarehouses(id, warehouseIDs)
		if err != nil {
			return
		}
	}
	return service.GetByID(id)
}

// GetDetailByID find assetby id
func (service Service) GetDetailByID(id string) (asset entity.AssetGroup, err error) {
	a, err := service.repository.FindByID(id)
	return service.mapAssetToAssetGroup(a)
}

// GetByID find assetby id
func (service Service) GetByID(id string) (asset entity.Asset, err error) {
	return service.repository.FindByID(id)
}

// GetByWorkOrderID finds asset by work order ID
func (service Service) GetByWorkOrderID(workOrderID string) (assetGroups []entity.AssetGroup, err error) {
	assets, err := service.repository.FindByWorkOrderID(workOrderID)
	if err != nil {
		return
	}
	for _, asset := range assets {
		assetGroup, _ := service.mapAssetToAssetGroup(asset)
		assetGroups = append(assetGroups, assetGroup)
	}
	return
}

// GetBySiteID finds asset by site ID
func (service Service) GetBySiteID(siteID string) (assets []entity.Asset, err error) {
	return service.repository.FindBySiteID(siteID)
}

// DeleteByID delete assetby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
