package documents

import (
	"go-boilerplate/adapters"

	"github.com/kataras/iris/v12"
)

const name = "/documents"

// Routes init documents
func Routes(app *iris.Application, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}

	documents := app.Party(name)

	documents.Post("/upload", handler.Create)
	documents.Get("/{id:string}", handler.GetByID)
}
