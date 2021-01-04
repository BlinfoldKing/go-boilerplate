package neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4jRepository repository implementation on neo4j
type Neo4jRepository struct {
	session neo4j.Session
	driver neo4j.Driver
}

// CreateNeo4jRepository init Neo4jRepository
func CreateNeo4jRepository(session neo4j.Session, driver neo4j.Driver) Repository {
	return Neo4jRepository{session, driver}
}

// CreateNode create single node to neo4j
func (repo Neo4jRepository) CreateNode(label string, properties map[string]interface{}) error {
	queryn := fmt.Sprintf("CREATE (n:%s $data)", label)

	_, err := repo.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryn, map[string]interface{}{
			"data": properties,
		})

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values(), nil
		}

		return nil, result.Err()
	})

	return err
}

// CreateRelation create single relation to neo4j
func (repo Neo4jRepository) CreateRelation(sourceProps PropertiesVal, destProp PropertiesVal) error {
	queryn := fmt.Sprintf("MATCH (a:%s), (b:%s) WHERE a.id = '%s' AND b.id = '%s' CREATE (a)-[r:RELTYPE]->(b)", sourceProps.Label, destProp.Label, sourceProps.ID, destProp.ID)

	_, err := repo.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryn, map[string]interface{}{})

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values(), nil
		}

		return nil, result.Err()
	})

	return err
}
