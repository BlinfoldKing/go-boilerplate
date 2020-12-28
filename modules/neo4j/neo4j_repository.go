package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4jRepository repository implementation on neo4j
type Neo4jRepository struct {
	neo neo4j.Driver
}

// CreateNeo4jRepository init Neo4jRepository
func CreateNeo4jRepository(db neo4j.Driver) Repository {
	return Neo4jRepository{db}
}

// CreateNode create single node to neo4j
func (repo Neo4jRepository) CreateNode(name string) error {
	session, err := repo.neo.NewSession(neo4j.SessionConfig{})

	if err != nil {
		return err
	}

	defer session.Close()

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("CREATE ($name)", map[string]interface{}{
			"name": name,
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
