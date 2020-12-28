package neo4j

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/neo4j"

// Routes init product
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreateNeo4jRepository(adapters.Neo4j)
	service := CreateService(repository)
	handler := handler{service, adapters}
	product := prefix.Party(name)
	product.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
}
