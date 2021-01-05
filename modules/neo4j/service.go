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
func (service Service) CreateNode(label string, data map[string]interface{}) (err error)  {
	err = service.repository.CreateNode(label, data)
	return
}

// CreateRelation create new relation
func (service Service) CreateRelation(sourceProps PropertiesVal, destProp PropertiesVal) (err error)  {
	err = service.repository.CreateRelation(sourceProps, destProp)
	return
}
