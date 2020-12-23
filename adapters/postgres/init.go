package postgres

import (
	"go-boilerplate/config"
	"go-boilerplate/helper"

	"context"

	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
	"xorm.io/xorm/contexts"
)

// Postgres extending xorm for extra feature
type Postgres struct {
	*xorm.Engine
}

type loggerHook struct{}

func (hook loggerHook) AfterProcess(c *contexts.ContextHook) error {
	if c.Err != nil {
		helper.Logger.Error(c.Err)
		return nil
	}

	return nil
}

func (hook loggerHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	helper.Logger.WithFields(
		logrus.Fields{
			"query": c.SQL,
			"args":  c.Args,
		}).Debug("SQL")
	return context.Background(), nil
}

// Init create data base driver using xorm
func Init() (db *Postgres, err error) {
	engine, _ := xorm.NewEngine("postgres", config.DBCONFIG())
	engine.AddHook(loggerHook{})
	return &Postgres{engine}, nil
}
