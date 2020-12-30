package adapters

import (
	"go-boilerplate/adapters/enforcer"
	fb "go-boilerplate/adapters/firebase"
	mailer "go-boilerplate/adapters/gomail"
	mg "go-boilerplate/adapters/mailgun"
	"go-boilerplate/adapters/minio"
	"go-boilerplate/adapters/nats"
	"go-boilerplate/adapters/postgres"
	rd "go-boilerplate/adapters/redis"
	validation "go-boilerplate/adapters/validator"
	"go-boilerplate/helper"

	firebase "firebase.google.com/go"
	"gopkg.in/gomail.v2"

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
	Gomail    *gomail.Dialer
}

// Init create new Adapters
func Init() (Adapters, error) {
	postgres, err := postgres.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to postgres")
	}

	validator, err := validation.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to postgres")
	}

	redis, err := rd.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to redis")
	}

	enforcer, err := enforcer.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to enforcer")
	}

	minio, err := minio.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to minio")
	}

	firebase, err := fb.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to firebase")
	}

	nats, err := nats.Init()
	if err != nil {
		helper.
			Logger.
			WithField("error", err).
			Warn("failed to connect to nats")
	}

	mailgun := mg.Init()

	gomail := mailer.Init()

	return Adapters{
		postgres,
		validator,
		redis,
		enforcer,
		minio,
		firebase,
		nats,
		mailgun,
		gomail,
	}, err
}
