package neo4j

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/topology"

// Routes init neo4j
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreateNeo4jRepository(adapters.Neo4j.Session, adapters.Neo4j.Driver)
	service := CreateService(repository)
	handler := handler{service, adapters}
	neo4j := prefix.Party(name)
	neo4j.Post("/nodes", middlewares.ValidateBody(&CreateRequestNodes{}), handler.CreateNodes)
	neo4j.Post("/edges", middlewares.ValidateBody(&CreateRequestEdges{}), handler.CreateEdges)
	neo4j.Delete("/nodes", middlewares.ValidateBody(&CreateRequestNodes{}), handler.DeleteNodes)
	neo4j.Delete("/edges", middlewares.ValidateBody(&CreateRequestNodes{}), handler.DeleteRelation)
}
