package neo4j

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	neo4j Service
	adapters adapters.Adapters
}

func (h handler) CreateNodes(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequestNodes)

	// Create Bulk Node
	if len(request.Nodes) > 0 {
		for _, node := range request.Nodes {
			err := h.neo4j.CreateNode(node.Label, node.Properties)
			if err != nil {
				helper.
					CreateErrorResponse(ctx, err).
					InternalServer().
					JSON()
				return
			}
		}
	}

	helper.CreateResponse(ctx).Ok().JSON()
	ctx.Next()
}

func (h handler) CreateEdges(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequestEdges)

	// Create Bulk Relation
	if len(request.Edges) > 0 {
		for _, edges := range request.Edges {
			if len(edges.Source) > 0 && len(edges.Destination) > 0 {
				for index, prop := range edges.Source {
					err := h.neo4j.CreateRelation(prop, edges.Destination[index])
					if err != nil {
						helper.
							CreateErrorResponse(ctx, err).
							InternalServer().
							JSON()
						return
					}
				}
			}
		}
	}

	helper.CreateResponse(ctx).Ok().JSON()
	ctx.Next()
}

func (h handler) DeleteNodes(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequestNodes)

	// Create Bulk Node
	if len(request.Nodes) > 0 {
		for _, node := range request.Nodes {
			err := h.neo4j.DeleteNode(node.Label, node.Properties)
			if err != nil {
				helper.
					CreateErrorResponse(ctx, err).
					InternalServer().
					JSON()
				return
			}
		}
	}

	helper.CreateResponse(ctx).Ok().JSON()
	ctx.Next()
}

func (h handler) DeleteRelation(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequestNodes)

	// Create Bulk Node
	if len(request.Nodes) > 0 {
		for _, node := range request.Nodes {
			err := h.neo4j.DeleteRelation(node.Label, node.Properties)
			if err != nil {
				helper.
					CreateErrorResponse(ctx, err).
					InternalServer().
					JSON()
				return
			}
		}
	}

	helper.CreateResponse(ctx).Ok().JSON()
	ctx.Next()
}
