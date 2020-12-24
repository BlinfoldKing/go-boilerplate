package ping

import (
	"go-boilerplate/adapters"

	"github.com/kataras/iris/v12"
)

const name = "/ping"

// Routes init ping
func Routes(app *iris.Application, adapters adapters.Adapters) {
	handler := handler{adapters}

	ping := app.Party(name)
	ping.Get("/", handler.Ping)
	ping.Get("/nats", handler.PingNats)
}
