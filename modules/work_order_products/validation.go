package workorderproducts

// CreateRequest request for create new workOrderProducts
type CreateRequest struct {
	WorkOrderID string `json:"work_order_id"`
	ProductID   string `json:"product_id"`
	Qty         int    `json:"qty"`
	Status      int    `json:"status"`
}

// UpdateRequest request for update workOrderProducts
type UpdateRequest struct {
	WorkOrderID string `json:"work_order_id"`
	ProductID   string `json:"product_id"`
	Qty         int    `json:"qty"`
	Status      int    `json:"status"`
}
