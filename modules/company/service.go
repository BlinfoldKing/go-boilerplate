package company

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	companydocument "go-boilerplate/modules/company_document"
	"go-boilerplate/modules/documents"
)

// Service contains business logic
type Service struct {
	repository       Repository
	companyDocuments companydocument.Service
	documents        documents.Service
}

func InitCompanyService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)

	companyDocuments := companydocument.InitCompanyDocumentService(adapters)
	documents := documents.InitDocumentsService(adapters)
	return CreateService(repository, companyDocuments, documents)
}

// CreateService init service
func CreateService(repo Repository, companyDocuments companydocument.Service, documents documents.Service) Service {
	return Service{repo, companyDocuments, documents}
}

func (service Service) mapCompaniesToCompanyGroups(companies []entity.Company) (companyGroups []entity.CompanyGroup, err error) {
	for _, company := range companies {
		documents, err := service.documents.GetByCompanyID(company.ID)
		if err != nil {
			return []entity.CompanyGroup{}, err
		}
		companyGroup := entity.CompanyGroup{
			Company:   company,
			Documents: documents,
		}

		companyGroups = append(companyGroups, companyGroup)
	}
	return
}

// CreateCompany create new company
func (service Service) CreateCompany(
	name string,
	companyType entity.CompanyType,
	address string,
	phoneNumber string,
	documentIDs []string,
) (company entity.Company, err error) {
	company, err = entity.NewCompany(
		name,
		companyType,
		address,
		phoneNumber,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(company)
	if err != nil {
		return
	}

	_, err = service.companyDocuments.CreateBatchCompanyDocuments(company.ID, documentIDs)
	return
}

// GetList get list of company
func (service Service) GetList(pagination entity.Pagination) (companyGroups []entity.CompanyGroup, count int, err error) {
	companies, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}

	companyGroups, err = service.mapCompaniesToCompanyGroups(companies)
	return
}

// Update update company
func (service Service) Update(id string, changeset entity.CompanyChangeSet) (company entity.CompanyGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.CompanyGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find companyby id
func (service Service) GetByID(id string) (companyGroup entity.CompanyGroup, err error) {
	company, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	documents, err := service.documents.GetByCompanyID(id)
	return entity.CompanyGroup{
		Company:   company,
		Documents: documents,
	}, err
}

// GetByBrandID finds companies by brandID
func (service Service) GetByBrandID(brandID string) (companies []entity.Company, err error) {
	return service.repository.FindByBrandID(brandID)
}

// DeleteByID delete companyby id
func (service Service) DeleteByID(id string) (err error) {
	err = service.repository.DeleteByID(id)
	if err != nil {
		return
	}
	err = service.companyDocuments.DeleteByCompanyID(id)
	return
}
