package neo4j

import (
	"go-boilerplate/config"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4j extends neo4j
type Neo4j struct {
	neo4j.Driver
}

// Init creates minio client
func Init() (*Neo4j, error) {
	driver, err := neo4j.NewDriver(config.NEO4JENDPOINT(), neo4j.BasicAuth(config.NEO4JUSER(), config.NEO4JPASSWORD(), ""))
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	return &Neo4j{driver}, err
}
