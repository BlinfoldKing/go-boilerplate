package documents

import (
	"go-boilerplate/adapters"

	"github.com/kataras/iris/v12"
)

const name = "/documents"

// Routes init documents
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	storageRepository := CreatePostgresRepository(adapters.Postgres)
	fileRepository := CreateMinioRepository(adapters.Minio)
	service := CreateService(storageRepository, fileRepository)
	handler := handler{service, adapters}

	documents := prefix.Party(name)

	documents.Post(":upload", handler.Upload)
	documents.Post(":download", handler.Download)

	documents.Get("/{id:string}", handler.GetByID)
	documents.Post("/", handler.Create)
}
