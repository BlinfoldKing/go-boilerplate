package ping

import "github.com/kataras/iris/v12"

const name = "/ping"

func Routes(app *iris.Application) {
	app.Party(name)

	app.Get("/", Ping)
}
