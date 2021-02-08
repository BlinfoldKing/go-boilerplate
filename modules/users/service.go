package users

import (
	"errors"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/modules/company"
	companycontact "go-boilerplate/modules/company_contact"
	"go-boilerplate/modules/contact"
	"go-boilerplate/modules/roles"
	userroles "go-boilerplate/modules/user_roles"
)

// Service contains business logic for users
type Service struct {
	repository      Repository
	roles           roles.Service
	userRoles       userroles.Service
	companyContacts companycontact.Service
	companies       company.Service
	contacts        contact.Service
}

// InitUserService create new user service
func InitUserService(adapters adapters.Adapters) Service {
	roleRepository := roles.CreatePosgresRepository(adapters.Postgres)
	roleService := roles.CreateService(roleRepository)

	userRoleRepository := userroles.CreatePosgresRepository(adapters.Postgres)
	userRoleService := userroles.CreateService(userRoleRepository)

	companyContactService := companycontact.InitCompanyContactService(adapters)
	companyService := company.InitCompanyService(adapters)
	contactService := contact.InitContactService(adapters)

	repository := CreatePosgresRepository(adapters.Postgres)

	return CreateService(repository, roleService, userRoleService, companyContactService, companyService, contactService)
}

// CreateService init service
func CreateService(repo Repository,
	roles roles.Service,
	userRoles userroles.Service,
	companyContacts companycontact.Service,
	companies company.Service,
	contacts contact.Service,
) Service {
	return Service{repo, roles, userRoles, companyContacts, companies, contacts}
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

	company := entity.CompanyGroup{}
	contact := entity.Contact{}

	if user.CompanyContactID != nil {
		companyContact, err := service.companyContacts.GetByID(*user.CompanyContactID)
		if err == nil {
			company, _ = service.companies.GetByID(companyContact.CompanyID)
			contact, _ = service.contacts.GetByID(companyContact.ContactID)
		} else {
			err = nil
		}
	}

	ug = entity.UserGroup{
		User:    user,
		Roles:   roles,
		Company: company.Company,
		Contact: contact,
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
func (service Service) CreateUser(email, password string, companyContactID *string) (res entity.UserGroup, err error) {
	user, err := entity.NewUser(email, password, companyContactID, entity.UserConfig{})
	if err != nil {
		return
	}

	if !config.EMAILACTIVATION() {
		user.ActiveStatus = entity.Active
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

	if user.ActiveStatus == entity.Inactive {
		return entity.UserGroup{}, errors.New("user is inactive")
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

// GetByEmail get user by email
func (service Service) GetByEmail(email string) (entity.UserGroup, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return entity.UserGroup{}, err
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

// GetByWorkOrderID finds user by work order ID
func (service Service) GetByWorkOrderID(workOrderID string) (users []entity.User, err error) {
	return service.repository.FindByWorkOrderID(workOrderID)
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
