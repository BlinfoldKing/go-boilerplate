package middlewares

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"io/ioutil"
	"strings"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// GenerateToken create new token
var GenerateToken func(iris.Context)

// AuthenticateToken check wheter token valid or not
var AuthenticateToken func(iris.Context)

// InvalidateToken invalidate token
var InvalidateToken func(iris.Context)

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

	verifyToken := func(ctx iris.Context) (user entity.UserGroup, err error) {

		if err = j.CheckJWT(ctx); err != nil {
			return
		}

		token := ctx.Values().Get("jwt").(*jwt.Token)
		claim := token.Claims.(jwt.MapClaims)
		userid := claim["user_id"].(string)

		data, err := adapters.Redis.Get(fmt.Sprintf("logged:user:%s", userid)).Result()
		if err != nil {
			return
		}

		json.Unmarshal([]byte(data), &user)

		return
	}

	GenerateToken = func(ctx iris.Context) {
		now := time.Now()
		duration := time.Duration(config.TOKENDURATION()) * time.Second

		u := ctx.Values().Get("user")
		user := u.(entity.UserGroup)

		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"iat":     now.Unix(),
			"exp":     now.Add(duration).Unix(),
		})

		tokenString, _ := token.SignedString(key)

		data, err := json.Marshal(user)
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
			return
		}

		err = adapters.Redis.SetNX(fmt.Sprintf("logged:user:%s", user.ID), data, 0).Err()
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
			return
		}

		helper.CreateResponse(ctx).
			Ok().
			WithData(iris.Map{
				"token": tokenString,
				"user":  user,
			}).
			JSON()

		ctx.Next()
	}

	AuthenticateToken = func(ctx iris.Context) {
		path := ctx.Path()
		path = strings.Replace(path, config.PREFIX(), "", 1)

		helper.Logger.Debug(path)

		sub := "public"
		obj := path
		act := ctx.Method()
		if ok, _ := adapters.Enforcer.Enforce(sub, obj, act); ok {
			ctx.Next()
			return
		}

		user, err := verifyToken(ctx)
		if err != nil {
			helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
			return
		}

		roles := user.Roles
		for _, role := range roles {
			sub = role.Slug
			if ok, err := adapters.Enforcer.Enforce(sub, obj, act); err == nil && ok {
				ctx.Next()
				return
			}

		}

		err = fmt.Errorf("not authorized for user: %s", user.ID)
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return

	}

	InvalidateToken = func(ctx iris.Context) {
		user, err := verifyToken(ctx)
		if err != nil {
			helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
			return
		}

		err = adapters.Redis.Del(fmt.Sprintf("logged:user:%s", user.ID)).Err()
		if err != nil {
			helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
			return
		}

		ctx.Next()
	}

	return nil
}
