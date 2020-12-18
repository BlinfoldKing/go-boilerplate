package users

import (
	"errors"
	"go-boilerplate/entity"
	"go-boilerplate/modules/roles"
	userroles "go-boilerplate/modules/user_roles"
)

// Service contains business logic for users
type Service struct {
	repository Repository
	roles      roles.Service
	userRoles  userroles.Service
}

// CreateService init service
func CreateService(repo Repository,
	roles roles.Service,
	userRoles userroles.Service,
) Service {
	return Service{repo, roles, userRoles}
}

// helper
func (service Service) mapUserRolesToRoles(ur []entity.UserRole) (roles []entity.Role, err error) {
	for _, role := range ur {
		var r entity.Role
		r, err = service.roles.GetByID(role.RoleID)
		if err != nil {
			return
		}

		roles = append(roles, r)
	}

	return
}

func (service Service) mapUserToUserGroup(user entity.User) (ug entity.UserGroup, err error) {
	ur, err := service.userRoles.GetAllByUserID(user.ID)
	if err != nil {
		return
	}

	roles, err := service.mapUserRolesToRoles(ur)
	if err != nil {
		return
	}

	ug = entity.UserGroup{
		ID:    user.ID,
		Email: user.Email,
		Roles: roles,
	}

	return
}

func (service Service) mapUsersToUserGroups(users []entity.User) (ug []entity.UserGroup, err error) {
	for _, user := range users {
		var u entity.UserGroup
		u, err = service.mapUserToUserGroup(user)
		if err != nil {
			return
		}

		ug = append(ug, u)
	}

	return
}

// CreateUser create new user
func (service Service) CreateUser(email, password string) (res entity.UserGroup, err error) {
	user, err := entity.NewUser(email, password, entity.UserConfig{})
	if err != nil {
		return
	}

	err = service.repository.Save(user)
	if err != nil {
		return
	}

	_, err = service.userRoles.CreateRole(user.ID, entity.DefaultMEMBER)
	if err != nil {
		return
	}

	res, err = service.mapUserToUserGroup(user)

	return
}

// AuthenticateUser create new user
func (service Service) AuthenticateUser(email, password string) (entity.UserGroup, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return entity.UserGroup{}, err
	}

	ok, err := user.ComparePassword(password, entity.UserConfig{})
	if err != nil {
		return entity.UserGroup{}, err
	}

	if !ok {
		return entity.UserGroup{}, errors.New("wrong password")
	}

	return service.mapUserToUserGroup(user)
}

// GetList get list of users
func (service Service) GetList(limit, offset int) ([]entity.UserGroup, error) {
	users, err := service.repository.GetList(limit, offset)
	if err != nil {
		return []entity.UserGroup{}, err
	}

	return service.mapUsersToUserGroups(users)
}

// Update update user
func (service Service) Update(id string, changeset entity.UserChangeSet) (entity.UserGroup, error) {
	err := service.repository.Update(id, changeset)
	if err != nil {
		return entity.UserGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find user by id
func (service Service) GetByID(id string) (user entity.UserGroup, err error) {
	u, err := service.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	return service.mapUserToUserGroup(u)
}

// DeleteByID delete user by id
func (service Service) DeleteByID(id string) error {
	return service.repository.DeleteByID(id)
}
