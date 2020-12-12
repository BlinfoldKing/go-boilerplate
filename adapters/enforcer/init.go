package enforcer

import (
	"database/sql"
	"github.com/casbin/casbin/v2"
	"github.com/cychiuae/casbin-pg-adapter"
	"go-boilerplate/config"
)

// Init create Enforcer
func Init() (*casbin.Enforcer, error) {

	db, err := sql.Open("postgres", config.DBCONFIG())
	if err != nil {
		return nil, err
	}

	adapter, err := casbinpgadapter.
		NewAdapter(db, "casbin_rule")
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("./auth_model.conf", adapter)
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadModel()
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, err
}
