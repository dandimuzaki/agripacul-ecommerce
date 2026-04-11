package request

type AddCartItemRequest struct {
	SKUID    uint `json:"sku_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartItemRequest struct {
	Quantity   *int  `json:"quantity,omitempty"`
	IsSelected *bool `json:"is_selected,omitempty"`
}

type BatchSelectCartItemRequest struct {
	IsSelected bool `json:"is_selected"`
}