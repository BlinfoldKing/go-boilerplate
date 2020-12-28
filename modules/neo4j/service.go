package neo4j

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateNode create new node
func (service Service) CreateNode(name string) (err error)  {
	err = service.repository.CreateNode(name)
	return
}
