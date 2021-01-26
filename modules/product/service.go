package product

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/brand"
	"go-boilerplate/modules/documents"
	productcategory "go-boilerplate/modules/product_category"
	productdocument "go-boilerplate/modules/product_document"
	productspecification "go-boilerplate/modules/product_specification"
	"time"
)

// Service contains business logic
type Service struct {
	repository            Repository
	brands                brand.Service
	documents             documents.Service
	productCategories     productcategory.Service
	productDocuments      productdocument.Service
	productSpecifications productspecification.Service
}

// InitProductService create new product
func InitProductService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)

	brandService := brand.InitBrandService(adapters)
	documentService := documents.InitDocumentsService(adapters)
	productCategoryService := productcategory.InitProductCategoryService(adapters)
	productDocumentService := productdocument.InitProductDocumentService(adapters)
	productSpecificationService := productspecification.InitProductSpecificationService(adapters)

	return CreateService(
		repository,
		brandService,
		documentService,
		productCategoryService,
		productDocumentService,
		productSpecificationService,
	)
}

// CreateService init service
func CreateService(
	repo Repository,
	brandService brand.Service,
	documentService documents.Service,
	productCategoryService productcategory.Service,
	productDocumentService productdocument.Service,
	productSpecificationService productspecification.Service,
) Service {
	return Service{
		repo,
		brandService,
		documentService,
		productCategoryService,
		productDocumentService,
		productSpecificationService,
	}
}

func (service Service) mapProductsToProductGroups(products []entity.Product) (productGroups []entity.ProductGroup, err error) {
	for _, product := range products {
		brand, err := service.brands.GetByID(product.BrandID)
		if err != nil {
			return []entity.ProductGroup{}, err
		}

		category, err := service.productCategories.GetByID(product.ProductCategoryID)
		if err != nil {
			return []entity.ProductGroup{}, err
		}

		documents, err := service.documents.GetByProductID(product.ID)
		if err != nil {
			return []entity.ProductGroup{}, err
		}

		specifications, err := service.productSpecifications.GetByProductID(product.ID)
		if err != nil {
			return []entity.ProductGroup{}, err
		}
		productGroup := entity.ProductGroup{
			Product:        product,
			Brand:          brand.Brand,
			Category:       category,
			Documents:      documents,
			Specifications: specifications,
		}

		productGroups = append(productGroups, productGroup)
	}
	return
}

// CreateProduct create new product
func (service Service) CreateProduct(
	name string,
	brandID string,
	productCategoryID string,
	productType string,
	productTags []string,
	lifetime time.Time,
	maintenanceInterval int,
	documentIDs []string,
) (product entity.Product, err error) {
	product, err = entity.NewProduct(
		name,
		brandID,
		productCategoryID,
		productType,
		productTags,
		lifetime,
		maintenanceInterval,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(product)

	_, err = service.productDocuments.CreateBatchProductDocument(product.ID, documentIDs)
	return
}

// GetList get list of product
func (service Service) GetList(pagination entity.Pagination) (productGroups []entity.ProductGroup, count int, err error) {
	products, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	productGroups, err = service.mapProductsToProductGroups(products)
	return
}

// Update update product
func (service Service) Update(id string, changeset entity.ProductChangeSet) (product entity.ProductGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.ProductGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find productby id
func (service Service) GetByID(id string) (productGroup entity.ProductGroup, err error) {
	product, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	brand, err := service.brands.GetByID(product.BrandID)
	if err != nil {
		return
	}

	category, err := service.productCategories.GetByID(product.ProductCategoryID)
	if err != nil {
		return
	}

	documents, err := service.documents.GetByProductID(id)
	if err != nil {
		return
	}

	specifications, err := service.productSpecifications.GetByProductID(id)
	return entity.ProductGroup{
		Product:        product,
		Brand:          brand.Brand,
		Category:       category,
		Documents:      documents,
		Specifications: specifications,
	}, err
}

// DeleteByID delete productby id
func (service Service) DeleteByID(id string) (err error) {
	err = service.repository.DeleteByID(id)
	if err != nil {
		return
	}

	err = service.productDocuments.DeleteByProductID(id)
	if err != nil {
		return
	}

	err = service.productSpecifications.DeleteByProductID(id)
	return
}
