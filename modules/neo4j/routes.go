package neo4j

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/neo4j"

// Routes init neo4j
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreateNeo4jRepository(adapters.Neo4j.Session, adapters.Neo4j.Driver)
	service := CreateService(repository)
	handler := handler{service, adapters}
	neo4j := prefix.Party(name)
	neo4j.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
}
