package companydocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

func InitCompanyDocumentService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateCompanyDocument create new companyDocument
func (service Service) CreateCompanyDocument(companyID, documentID string) (companyDocument entity.CompanyDocument, err error) {
	companyDocument, err = entity.NewCompanyDocument(companyID, documentID)
	if err != nil {
		return
	}
	err = service.repository.Save(companyDocument)
	return
}

// CreateBatchCompanyDocuments creates a batch of new companyDocuments
func (service Service) CreateBatchCompanyDocuments(companyID string, documentIDs []string) (companyDocuments []entity.CompanyDocument, err error) {
	for _, documentID := range documentIDs {
		companyDocument, err := entity.NewCompanyDocument(companyID, documentID)
		if err != nil {
			return []entity.CompanyDocument{}, err
		}
		companyDocuments = append(companyDocuments, companyDocument)
	}
	err = service.repository.SaveBatch(companyDocuments)
	return
}

// GetList get list of companyDocument
func (service Service) GetList(pagination entity.Pagination) (companyDocument []entity.CompanyDocument, count int, err error) {
	companyDocument, count, err = service.repository.GetList(pagination)
	return
}

// Update update companyDocument
func (service Service) Update(id string, changeset entity.CompanyDocumentChangeSet) (companyDocument entity.CompanyDocument, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.CompanyDocument{}, err
	}
	return service.GetByID(id)
}

// GetByID find companyDocumentby id
func (service Service) GetByID(id string) (companyDocument entity.CompanyDocument, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete company_documentby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByCompanyID delete company_document by company id
func (service Service) DeleteByCompanyID(companyID string) (err error) {
	return service.repository.DeleteByCompanyID(companyID)
}
