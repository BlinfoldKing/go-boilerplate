package neo4j

import (
	"go-boilerplate/config"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4j extends neo4j
type Neo4j struct {
	neo4j.Session
	neo4j.Driver
}

// ConnectToDB creates neo4j client
func ConnectToDB() (neo4j.Session, neo4j.Driver, error) {
	// define driver, session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver(config.NEO4JENDPOINT(), neo4j.BasicAuth(config.NEO4JUSER(), config.NEO4JPASSWORD(), ""), func(c *neo4j.Config) {
		c.Encrypted = false
	}); err != nil {
		return nil, nil, err
	}

	// Open a new session with write access
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, err
	}
	return session, driver, nil
}

// Init creates neo4j client
func Init() (*Neo4j, error) {
	session, driver, err := ConnectToDB()
	if err != nil {
		log.Fatalln("Error connecting to Database")
		log.Fatalln(err)
	}
	log.Println("Connected to Neo4j")
	
	return &Neo4j{session, driver}, err
}
