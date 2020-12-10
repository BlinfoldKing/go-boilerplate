package middlewares

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

// Logger Basic logger
func Logger(ctx iris.Context) {
	body, _ := ctx.GetBody()
	log := fmt.Sprintf(
		"path: %s, method: %s, body: %s",
		ctx.Path(), ctx.Request().Method, body)

	logrus.Info(log)

	ctx.Next()
}
