package postgres

import (
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/entity"

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
		logrus.Error(c.Err)
		return nil
	}

	return nil
}

func (hook loggerHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	logrus.Println(fmt.Sprintf("[pagination] query: %s, args: %v", c.SQL, c.Args))
	return context.Background(), nil
}

// Init create data base driver using xorm
func Init() (db *Postgres, err error) {
	engine, _ := xorm.NewEngine("postgres", config.DBCONFIG())
	engine.AddHook(loggerHook{})
	return &Postgres{engine}, nil
}
