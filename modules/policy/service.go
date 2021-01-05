package policy

import (
	"go-boilerplate/entity"
	"go-boilerplate/modules/roles"
	"regexp"

	"github.com/kataras/iris/v12"
)

// Service policy service
type Service struct {
	repo  Repository
	roles roles.Service
}

// CreateService create new repo
func CreateService(repo Repository, roles roles.Service) Service {
	return Service{repo, roles}
}

// AddPolicy add new auth policy
func (service Service) AddPolicy(roleID, path, method string) (entity.Policy, error) {
	role, err := service.roles.GetByID(roleID)
	if err != nil {
		return entity.Policy{}, err
	}

	return service.repo.AddPolicy(entity.Policy{
		Method: method,
		Path:   path,
		Role:   role,
	})
}

// DeletePolicy delete auth policy
func (service Service) DeletePolicy(roleID, path, method string) error {
	role, err := service.roles.GetByID(roleID)
	if err != nil {
		return err
	}

	return service.repo.DeletePolicy(role.Slug, path, method)
}

// GetAllPolicies get all policy
func (service Service) GetAllPolicies() ([]entity.Policy, error) {
	policies, err := service.repo.ListPolicy()
	if err != nil {
		return []entity.Policy{}, err
	}

	for i := range policies {
		policies[i].Role, err = service.roles.FindBySlug(policies[i].Role.Slug)
		if err != nil {
			return []entity.Policy{}, err
		}
	}

	return policies, nil
}

// GetAllRoutes get all routes from context
func (service Service) GetAllRoutes(ctx iris.Context) []iris.Map {
	roRoutes := ctx.Application().GetRoutesReadOnly()
	routes := make([]iris.Map, 0)
	for _, route := range roRoutes {
		path := route.Path()
		open := regexp.MustCompile(`{`)
		close := regexp.MustCompile(`\:\w+}`)
		path = open.ReplaceAllLiteralString(path, ":")
		path = close.ReplaceAllLiteralString(path, "")
		routes = append(routes, iris.Map{
			"method": route.Method(),
			"path":   path,
		})
	}

	return routes
}
