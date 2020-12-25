package adapters

import (
	"go-boilerplate/adapters/enforcer"
	fb "go-boilerplate/adapters/firebase"
	mg "go-boilerplate/adapters/mailgun"
	"go-boilerplate/adapters/minio"
	"go-boilerplate/adapters/nats"
	"go-boilerplate/adapters/postgres"
	rd "go-boilerplate/adapters/redis"
	validation "go-boilerplate/adapters/validator"

	firebase "firebase.google.com/go"

	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/mailgun/mailgun-go"
)

// Adapters is wrapper for lib/drivers that needed to be injected
type Adapters struct {
	Postgres  *postgres.Postgres
	Validator *validator.Validate
	Redis     *redis.Client
	Enforcer  *casbin.Enforcer
	Minio     *minio.Minio
	Firebase  *firebase.App
	Nats      *nats.Nats
	Mailgun   *mailgun.MailgunImpl
}

// Init create new Adapters
func Init() (Adapters, error) {
	postgres, err := postgres.Init()
	if err != nil {
		return Adapters{}, err
	}

	validator, err := validation.Init()
	if err != nil {
		return Adapters{}, err
	}

	redis, err := rd.Init()
	if err != nil {
		return Adapters{}, err
	}

	enforcer, err := enforcer.Init()
	if err != nil {
		return Adapters{}, err
	}

	minio, err := minio.Init()
	if err != nil {
		return Adapters{}, err
	}

	firebase, err := fb.Init()
	if err != nil {
		return Adapters{}, err
	}

	nats, err := nats.Init()
	if err != nil {
		return Adapters{}, err
	}

	mailgun := mg.Init()

	return Adapters{
		postgres,
		validator,
		redis,
		enforcer,
		minio,
		firebase,
		nats,
		mailgun,
	}, err
}
