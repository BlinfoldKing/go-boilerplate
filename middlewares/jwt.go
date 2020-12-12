package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"io/ioutil"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// GenerateToken create new token
var GenerateToken func(iris.Context)

// AuthenticateToken check wheter token valid or not
var AuthenticateToken func(iris.Context)

// InitJWT init JWT struct
func InitJWT(adapters adapters.Adapters) error {
	key, err := ioutil.ReadFile(".keys/public.pem")
	if err != nil {
		return err
	}

	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		Expiration:    true,
		SigningMethod: jwt.SigningMethodHS256,
	})

	GenerateToken = func(ctx iris.Context) {
		path := ctx.Path()
		if path != "/auth/login" && path != "/auth/register" {
			ctx.Next()
			return
		}

		now := time.Now()
		duration := time.Duration(config.TOKENDURATION()) * time.Second

		u := ctx.Values().Get("user")
		user := u.(entity.User)

		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"iat":     now.Unix(),
			"exp":     now.Add(duration).Unix(),
		})

		tokenString, _ := token.SignedString(key)

		data, _ := json.Marshal(user)
		err = adapters.Redis.Set(fmt.Sprintf("logged:user:%s", user.ID), data, 0).Err()
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
		}

		err = adapters.Redis.Expire(fmt.Sprintf("logged:user:%s", user.ID), duration).Err()
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
		}

		helper.CreateResponse(ctx).
			Ok().
			WithData(iris.Map{
				"token": tokenString,
				"user":  user,
			}).
			JSON()
	}

	AuthenticateToken = func(ctx iris.Context) {
		path := ctx.Path()
		if path == "/auth/login" || path == "/auth/register" {
			ctx.Next()
			return
		}

		sub := "public"
		obj := path
		act := ctx.Method()
		if ok, _ := adapters.Enforcer.Enforce(sub, obj, act); ok {
			ctx.Next()
			return
		}

		if err := j.CheckJWT(ctx); err != nil {
			helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
			return
		}

		token := ctx.Values().Get("jwt").(*jwt.Token)
		claim := token.Claims.(jwt.MapClaims)
		userid := claim["user_id"].(string)

		data, err := adapters.Redis.Get(fmt.Sprintf("logged:user:%s", userid)).Result()
		if err != nil {
			helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		}

		var user entity.User
		json.Unmarshal([]byte(data), &user)
		sub = user.Role

		if ok, _ := adapters.Enforcer.Enforce(sub, obj, act); !ok {
			helper.CreateErrorResponse(ctx, errors.New("Not Allowed for this role")).Unauthorized().JSON()
			return
		}

		ctx.Next()
	}

	return nil
}
