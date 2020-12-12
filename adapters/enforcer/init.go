package enforcer

import (
	"database/sql"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/cychiuae/casbin-pg-adapter"
	"go-boilerplate/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type policyConfig struct {
	Policy map[string][]struct {
		Method string `yaml:"method"`
		Route  string `yaml:"route"`
	} `yaml:"policy"`
	// Policy interface{} `yaml:"policy"`
}

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

	configFile, err := ioutil.ReadFile("./default_auth_policy.yaml")
	if err != nil {
		return nil, err
	}

	file := string(configFile)
	configFile = []byte(strings.Trim(file, " "))

	conf := policyConfig{}
	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		return nil, err
	}

	for key, items := range conf.Policy {
		fmt.Println(key)
		for _, item := range items {
			fmt.Println(item)
			enforcer.AddPolicy(key, item.Route, item.Method)
		}
	}

	enforcer.SavePolicy()

	return enforcer, err
}
