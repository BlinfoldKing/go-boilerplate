package site

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Site) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Site, error)
	Update(id string, changeset entity.SiteChangeSet) error
	GetList(pagination entity.Pagination) (Sites []entity.Site, count int, err error)
}
