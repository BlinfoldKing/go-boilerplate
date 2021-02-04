package warehousecontact

// CreateRequest request for create new warehouse_contact
type CreateRequest struct {
	WarehouseID string `json:"warehouse_id" validate:"required"`
	ContactID   string `json:"contact_id" validate:"required"`
}

// UpdateRequest request for update warehouse_contact
type UpdateRequest struct {
	WarehouseID string `json:"warehouse_id"`
	ContactID   string `json:"contact_id"`
}
