package companycontactget

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/company"
	"go-boilerplate/modules/contact"
)

// Service contains business logic
type Service struct {
	repository Repository
	companies  company.Service
	contacts   contact.Service
}

// InitCompanyContactGetService inits company contact service
func InitCompanyContactGetService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	companies := company.InitCompanyService(adapters)
	contacts := contact.InitContactService(adapters)
	return CreateService(repository, companies, contacts)
}

// CreateService init service
func CreateService(repo Repository, companies company.Service, contacts contact.Service) Service {
	return Service{repo, companies, contacts}
}

func (service Service) mapCompanyContactToCompanyContactGroup(companyContact entity.CompanyContact) (companyContactGroup entity.CompanyContactGroup, err error) {
	companyContactGroup.CompanyContact = companyContact
	company, err := service.companies.GetByID(companyContact.CompanyID)
	if err != nil {
		return
	}
	companyContactGroup.Company = company.Company

	companyContactGroup.Contact, err = service.contacts.GetByID(companyContact.ContactID)
	return
}

// GetList get list of company_contact
func (service Service) GetList(pagination entity.Pagination) (companyContactGroups []entity.CompanyContactGroup, count int, err error) {
	companyContacts, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	for _, companyContact := range companyContacts {
		companyContactGroup, _ := service.mapCompanyContactToCompanyContactGroup(companyContact)
		companyContactGroups = append(companyContactGroups, companyContactGroup)
	}
	return
}

// GetByID find company_contactby id
func (service Service) GetByID(id string) (companyContactGroup entity.CompanyContactGroup, err error) {
	companyContact, err := service.repository.FindByID(id)
	if err != nil {
		return
	}
	companyContactGroup, err = service.mapCompanyContactToCompanyContactGroup(companyContact)
	return
}
