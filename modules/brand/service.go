package brand

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	brandcompany "go-boilerplate/modules/brand_company"
	"go-boilerplate/modules/company"
)

// Service contains business logic
type Service struct {
	repository     Repository
	brandCompanies brandcompany.Service
	companies      company.Service
}

func InitBrandService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)

	brandCompanyRepository := brandcompany.CreatePostgresRepository(adapters.Postgres)
	brandCompanyService := brandcompany.CreateService(brandCompanyRepository)

	companyRepository := company.CreatePostgresRepository(adapters.Postgres)
	companyService := company.CreateService(companyRepository)

	return CreateService(repository, brandCompanyService, companyService)
}

// CreateService init service
func CreateService(repo Repository, brandCompanies brandcompany.Service, companies company.Service) Service {
	return Service{repo, brandCompanies, companies}
}

func (service Service) mapBrandsToBrandGroups(brands []entity.Brand) (brandGroups []entity.BrandGroup, err error) {
	for _, brand := range brands {
		companies, err := service.companies.GetByBrandID(brand.ID)
		if err != nil {
			return []entity.BrandGroup{}, err
		}

		brandGroup := entity.BrandGroup{
			brand,
			companies,
		}

		brandGroups = append(brandGroups, brandGroup)
	}
	return
}

// CreateBrand create new brand
func (service Service) CreateBrand(name, originCountry string, companyIDs []string) (brand entity.Brand, err error) {
	brand, err = entity.NewBrand(name, originCountry)
	if err != nil {
		return
	}
	err = service.repository.Save(brand)
	if err != nil {
		return
	}

	_, err = service.brandCompanies.CreateBatchBrandCompany(brand.ID, companyIDs)
	return
}

// GetList get list of brand
func (service Service) GetList(pagination entity.Pagination) (brandGroups []entity.BrandGroup, count int, err error) {
	brands, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	brandGroups, err = service.mapBrandsToBrandGroups(brands)
	return
}

// Update update brand
func (service Service) Update(id string, changeset entity.BrandChangeSet) (brand entity.BrandGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return
	}
	return service.GetByID(id)
}

// GetByID find brandby id
func (service Service) GetByID(id string) (brandGroup entity.BrandGroup, err error) {
	brand, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	companies, err := service.companies.GetByBrandID(id)
	if err != nil {
		return
	}
	return entity.BrandGroup{
		brand,
		companies,
	}, nil
}

// DeleteByID delete brandby id
func (service Service) DeleteByID(id string) (err error) {
	err = service.repository.DeleteByID(id)
	if err != nil {
		return
	}
	err = service.brandCompanies.DeleteByBrandID(id)
	return
}
