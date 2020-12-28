package middlewares

import (
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

// Logger Basic logger
func Logger(ctx iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)

	helper.Logger.
		WithFields(
			logrus.Fields{
				"method": ctx.Request().Method,
				"path":   ctx.Path(),
				"body":   body,
			}).Info("HTTP")

	ctx.Next()
}
