package neo4j

// Repository abstraction for storage
type Repository interface {
	CreateNode(neo string) error
}
