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
