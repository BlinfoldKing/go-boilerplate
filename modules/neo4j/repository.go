package neo4j

// Repository abstraction for storage
type Repository interface {
	CreateNode(label string, data map[string]interface{}) error
	CreateRelation(sourceProps PropertiesVal, destProp PropertiesVal) error
	DeleteNode(label string, data map[string]interface{}) error
	DeleteRelation(label string, data map[string]interface{}) error
}
