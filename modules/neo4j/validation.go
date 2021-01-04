package neo4j

// NodesVal request for create new node
type NodesVal struct {
	Label		string					`json:"label" validate:"required"`
	Properties	map[string]interface{}	`json:"properties" validate:"required"`
}

// PropertiesVal request for create new node
type PropertiesVal struct {
	ID			string					`json:"id" validate:"required"`
	Label		string					`json:"label" validate:"required"`
}

// EdgesVal request for create new node
type EdgesVal struct {
	Source		[]PropertiesVal			`json:"source" validate:"required"`
	Destination	[]PropertiesVal			`json:"destination" validate:"required"`
}

// CreateRequest request for create new node
type CreateRequest struct {
	Nodes		[]NodesVal				`json:"nodes" validate:"required"`
	Edges		[]EdgesVal				`json:"edges" validate:"required"`
}
