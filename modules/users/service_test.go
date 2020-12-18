package users

import (
	"testing"
	"errors"

	entity "go-boilerplate/entity"

	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	assert := assert.New(t)

	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	svc := CreateService(repo)

	id := "1"
	role := "role"
	email := "mendoanjoe@google.com"
	password := "password"

	t.Run("CreateUser will success", func(t *testing.T) {
		// invoke some service methods
		// assert that repository is called as you expect like
		repo.EXPECT().Save(gomock.AssignableToTypeOf(entity.User{})).Return(nil)

		user, err := svc.CreateUser(email, password)
		assert.Equal(err, nil)
		assert.Equal(user.Email, email)
	})

	t.Run("AuthenticateUser will success", func(t *testing.T) {
		// invoke some service methods
		// assert that repository is called as you expect like
		repo.EXPECT().FindByEmail(email).Return(entity.User{}, nil)

		user, err := svc.AuthenticateUser(email, password)
		assert.Equal(err, errors.New("argon2id: hash is not in the correct format"))
		assert.Equal(user, entity.User{})
	})

	t.Run("GetList will success", func(t *testing.T) {
		// invoke some service methods
		// assert that repository is called as you expect like
		repo.EXPECT().GetList(10, 100).Return([]entity.User{}, nil)

		user, err := svc.GetList(10, 100)
		assert.Equal(err, nil)
		assert.Equal(user, []entity.User{})
	})

	t.Run("Update will success", func(t *testing.T) {
		// invoke some service methods
		// assert that repository is called as you expect like
		repo.EXPECT().Update(id, entity.UserChangeSet{ Role: role }).Return(nil)
		repo.EXPECT().FindByID(id).Return(entity.User{}, nil)

		user, err := svc.Update(id, entity.UserChangeSet{ Role: role })
		assert.Equal(err, nil)
		assert.Equal(user, entity.User{})
	})

	t.Run("GetByID will success", func(t *testing.T) {
		// invoke some service methods
		// assert that repository is called as you expect like
		repo.EXPECT().FindByID(id).Return(entity.User{}, nil)

		user, err := svc.GetByID(id)
		assert.Equal(err, nil)
		assert.Equal(user, entity.User{})
	})

	t.Run("DeleteByID will success", func(t *testing.T) {
		// invoke some service methods
		// assert that repository is called as you expect like
		repo.EXPECT().DeleteByID(id).Return(nil)

		err := svc.DeleteByID(id)
		assert.Equal(err, nil)
	})
}
