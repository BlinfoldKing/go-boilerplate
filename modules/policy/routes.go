package policy

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/roles"

	"github.com/kataras/iris/v12"
)

const name = "/policy"

// Routes init auth
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	policyRepository := CreateEnforcerRepository(adapters.Enforcer)

	roleRepository := roles.CreatePosgresRepository(adapters.Postgres)
	roleService := roles.CreateService(roleRepository)

	policyService := CreateService(policyRepository, roleService)
	handler := handler{policyService, adapters}

	policy := prefix.Party(name)

	policy.Post("/", middlewares.ValidateBody(&AddPolicyRequest{}),
		handler.AddPolicy)

	policy.Delete("/", middlewares.ValidateBody(&DeletePolicyRequest{}),
		handler.DeletePolicy)

	policy.Get("/", handler.GetAllPolicies)

	policy.Get(":routes", middlewares.ValidateBody(&AddPolicyRequest{}),
		handler.GetAllRoutes)

}
