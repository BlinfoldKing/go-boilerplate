package policy

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/roles"

	"github.com/kataras/iris/v12"
)

const name = "/policy"

// Routes init auth
func Routes(app *iris.Application, adapters adapters.Adapters) {
	policyRepository := CreateEnforcerRepository(adapters.Enforcer)

	roleRepository := roles.CreatePosgresRepository(adapters.Postgres)
	roleService := roles.CreateService(roleRepository)

	policyService := CreateService(policyRepository, roleService)
	handler := handler{policyService, adapters}

	policy := app.Party(name)

	policy.Post("/", middlewares.ValidateBody(&AddPolicyRequest{}),
		handler.AddPolicy)
}
