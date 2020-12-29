package brand

import (
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateBrand create new brand
func (service Service) CreateBrand(name, originCountry string) (brand entity.Brand, err error) {
	brand, err = entity.NewBrand(name, originCountry)
	if err != nil {
		return
	}
	err = service.repository.Save(brand)
	return
}

// GetList get list of brand
func (service Service) GetList(pagination entity.Pagination) (brand []entity.Brand, count int, err error) {
	brand, count, err = service.repository.GetList(pagination)
	return
}

// Update update brand
func (service Service) Update(id string, changeset entity.BrandChangeSet) (brand entity.Brand, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Brand{}, err
	}
	return service.GetByID(id)
}

// GetByID find brandby id
func (service Service) GetByID(id string) (brand entity.Brand, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete brandby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
