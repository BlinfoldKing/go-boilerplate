package middlewares

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/cychiuae/casbin-pg-adapter"
	"github.com/go-redis/redis"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"io/ioutil"
	"time"
)

// JWT jwt middleware wrapper
type JWT struct {
	key      []byte
	config   *jwt.Middleware
	redis    *redis.Client
	enforcer *casbin.Enforcer
}

// CreateJWT init JWT struct
func CreateJWT(redis *redis.Client) (JWT, error) {
	key, err := ioutil.ReadFile(".keys/public.pem")
	if err != nil {
		return JWT{}, err
	}

	jwt := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		Expiration:    true,
		SigningMethod: jwt.SigningMethodHS256,
	})

	db, _ := sql.Open("postgres", config.DBCONFIG())
	adapter, err := casbinpgadapter.
		NewAdapter(db, "casbin_rule")
	if err != nil {
		return JWT{}, err
	}

	enforcer, err := casbin.NewEnforcer("./auth_model.conf", adapter)
	if err != nil {
		return JWT{}, err
	}

	err = enforcer.LoadModel()
	if err != nil {
		return JWT{}, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return JWT{}, err
	}

	return JWT{
		[]byte(key),
		jwt,
		redis,
		enforcer,
	}, nil
}

// GenerateToken create new token
func (j JWT) GenerateToken(ctx iris.Context) {
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

	tokenString, _ := token.SignedString(j.key)

	data, _ := json.Marshal(user)
	_ = j.redis.Set(fmt.Sprintf("logged:user:%s", user.ID), data, 0)

	helper.CreateResponse(ctx).
		Ok().
		WithData(iris.Map{
			"token": tokenString,
			"user":  user,
		}).
		JSON()
}

// AuthenticateToken check wheter token valid or not
func (j JWT) AuthenticateToken(ctx iris.Context) {
	path := ctx.Path()
	if path == "/auth/login" || path == "/auth/register" {
		ctx.Next()
		return
	}

	sub := "public"
	obj := path
	act := ctx.Method()
	if ok, _ := j.enforcer.Enforce(sub, obj, act); ok {
		ctx.Next()
		return
	}

	if err := j.config.CheckJWT(ctx); err != nil {
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return
	}

	token := ctx.Values().Get("jwt").(*jwt.Token)
	claim := token.Claims.(jwt.MapClaims)
	userid := claim["user_id"].(string)

	data, _ := j.redis.Get(fmt.Sprintf("logged:user:%s", userid)).Result()

	var user entity.User
	json.Unmarshal([]byte(data), &user)
	sub = user.Role

	if ok, _ := j.enforcer.Enforce(sub, obj, act); !ok {
		helper.CreateErrorResponse(ctx, errors.New("Not Allowed for this role")).Unauthorized().JSON()
		return
	}

	ctx.Next()
}
