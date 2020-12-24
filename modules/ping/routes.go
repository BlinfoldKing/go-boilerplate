package ping

import (
	"go-boilerplate/adapters"

	"github.com/kataras/iris/v12"
)

const name = "/ping"

// Routes init ping
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	handler := handler{adapters}

	ping := prefix.Party(name)
	ping.Get("/", handler.Ping)
	ping.Get("/nats", handler.PingNats)
}
