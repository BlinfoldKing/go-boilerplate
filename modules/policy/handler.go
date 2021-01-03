package policy

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	policy   Service
	adapters adapters.Adapters
}

func (h handler) AddPolicy(ctx iris.Context) {
	request := ctx.Values().Get("body").(*AddPolicyRequest)

	policy, err := h.policy.AddPolicy(request.RoleID, request.Path, request.Method)
	if err != nil {
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(policy).JSON()
	ctx.Next()
}

func (h handler) GetAllPolicies(ctx iris.Context) {
	policies, err := h.policy.GetAllPolicies()
	if err != nil {
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(policies).JSON()
	ctx.Next()
}

func (h handler) DeletePolicy(ctx iris.Context) {
	policies, err := h.policy.GetAllPolicies()
	if err != nil {
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(policies).JSON()
	ctx.Next()
}

func (h handler) GetAllRoutes(ctx iris.Context) {
	request := ctx.Values().Get("body").(*DeletePolicyRequest)

	err := h.policy.DeletePolicy(request.RoleID, request.Path, request.Method)
	if err != nil {
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithMessage("policy deleted").JSON()
	ctx.Next()
}
