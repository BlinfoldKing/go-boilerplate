package sitecontact

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.SiteContact) error
	DeleteByID(id string) error
	FindByID(id string) (entity.SiteContact, error)
	Update(id string, changeset entity.SiteContactChangeSet) error
	GetList(pagination entity.Pagination) (SiteContacts []entity.SiteContact, count int, err error)
}
