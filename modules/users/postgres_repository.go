package users

import (
	"go-boilerplate/entity"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func CreatePosgresRepository(db *gorm.DB) Repository {
	return PostgresRepository{db}
}

func (repo PostgresRepository) Save(user entity.User) error {
	err := repo.db.Model(&user).Create(&user).Error
	return err
}

func (repo PostgresRepository) FindByEmail(email string) (user entity.User, err error) {
	err = repo.db.Model(&user).Where("email = ?").Find(&user).Error
	return
}
