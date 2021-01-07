package policy

import (
	"go-boilerplate/entity"

	"github.com/casbin/casbin/v2"
)

// EnforcerRepostitory repository implemented in casbin
type EnforcerRepostitory struct {
	enforcer *casbin.Enforcer
}

// CreateEnforcerRepository create new repository
func CreateEnforcerRepository(enforcer *casbin.Enforcer) Repository {
	return EnforcerRepostitory{
		enforcer,
	}
}

// AddPolicy add new policy
func (repo EnforcerRepostitory) AddPolicy(policy entity.Policy) (entity.Policy, error) {
	_, err := repo.enforcer.AddPolicy(policy.Role.Slug, policy.Path, policy.Method)
	return policy, err
}

// DeletePolicy delete policy
func (repo EnforcerRepostitory) DeletePolicy(roleSlug, path, method string) error {
	_, err := repo.enforcer.DeletePermission(roleSlug, path, method)
	return err
}

// ListPolicy get all policy
func (repo EnforcerRepostitory) ListPolicy() ([]entity.Policy, error) {
	policies := repo.enforcer.GetPolicy()

	res := make([]entity.Policy, 0)
	for _, policy := range policies {
		res = append(res, entity.Policy{
			Role: entity.Role{
				Slug: policy[0],
			},
			Path:   policy[1],
			Method: policy[2],
		})
	}

	return res, nil
}
