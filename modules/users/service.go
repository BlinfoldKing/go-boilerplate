package users

import (
	"errors"
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"go-boilerplate/modules/mail"
	"go-boilerplate/modules/otps"
	"go-boilerplate/modules/roles"
	userroles "go-boilerplate/modules/user_roles"
)

// Service contains business logic for users
type Service struct {
	repository Repository
	roles      roles.Service
	userRoles  userroles.Service
	otps       otps.Service
	mail       mail.Service
}

// CreateService init service
func CreateService(repo Repository,
	roles roles.Service,
	userRoles userroles.Service,
	otps otps.Service,
	mail mail.Service,
) Service {
	return Service{repo, roles, userRoles, otps, mail}
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
		User:  user,
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

func (service Service) generateLink(purpose int, token, email string) string {
	path := "request-activation"
	if purpose == entity.ResetPassword {
		path = "reset-password"
	}
	return fmt.Sprintf(
		"%s%s/auth/%s?token=%s&email=%s",
		config.APPURL(),
		config.PREFIX(),
		path,
		token,
		email,
	)
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

	if config.EMAILACTIVATION() {
		otp, err := service.otps.CreateOTP(email, entity.AccountActivation)
		if err != nil {
			return entity.UserGroup{}, err
		}
		activateLink := service.generateLink(entity.AccountActivation, otp.Token, email)
		emailBody, err := helper.GenerateActivationHTML(email, activateLink)
		if err != nil {
			return entity.UserGroup{}, err
		}
		_, err = service.mail.SendEmail("Semeru", "Account Activation", emailBody, email)
		if err != nil {
			return entity.UserGroup{}, err
		}
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
func (service Service) GetList(pagination entity.Pagination) (ug []entity.UserGroup, count int, err error) {
	users, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}

	ug, err = service.mapUsersToUserGroups(users)
	return
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
