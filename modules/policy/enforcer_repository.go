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
