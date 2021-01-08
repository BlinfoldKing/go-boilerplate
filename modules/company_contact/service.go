package companycontact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitCompanyContactService inits company contact service
func InitCompanyContactService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateCompanyContact create new company_contact
func (service Service) CreateCompanyContact(companyID, contactID string) (companyContact entity.CompanyContact, err error) {
	companyContact, err = entity.NewCompanyContact(companyID, contactID)
	if err != nil {
		return
	}
	err = service.repository.Save(companyContact)
	return
}

// CreateBatchCompanyContacts creates a batch of new CompanyContacts
func (service Service) CreateBatchCompanyContacts(companyID string, contactIDs []string) (companyContacts []entity.CompanyContact, err error) {
	for _, contactID := range contactIDs {
		companyContact, err := entity.NewCompanyContact(companyID, contactID)
		if err != nil {
			return []entity.CompanyContact{}, err
		}
		companyContacts = append(companyContacts, companyContact)
	}
	err = service.repository.SaveBatch(companyContacts)
	return
}

// GetList get list of company_contact
func (service Service) GetList(pagination entity.Pagination) (companyContact []entity.CompanyContact, count int, err error) {
	companyContact, count, err = service.repository.GetList(pagination)
	return
}

// Update update company_contact
func (service Service) Update(id string, changeset entity.CompanyContactChangeSet) (companyContact entity.CompanyContact, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.CompanyContact{}, err
	}
	return service.GetByID(id)
}

// GetByID find company_contactby id
func (service Service) GetByID(id string) (companyContact entity.CompanyContact, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete company_contactby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByCompanyID delete company contact by company id
func (service Service) DeleteByCompanyID(companyID string) (err error) {
	return service.repository.DeleteByCompanyID(companyID)
}
