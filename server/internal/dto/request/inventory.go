package request

type InventoryListRequest struct {
	PaginationRequest
	Status string `form:"status" binding:"omitempty,oneof=active inactive archived"`
	Stock  string `form:"stock" binding:"omitempty,oneof=in out low"`
	Search string `form:"search" binding:"omitempty"`

	// Sorting
	SortBy    string    `form:"sort_by" binding:"omitempty,oneof=created_at updated_at id name stock"`
	SortOrder SortOrder `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type SortInventoryOption string

var (
	SortInventoryByCreatedAt SortInventoryOption = "created_at"
	SortInventoryByUpdatedAt SortInventoryOption = "updated_at"
	SortInventoryByID        SortInventoryOption = "id"
	SortInventoryByName      SortInventoryOption = "name"
	SortInventoryByStock     SortInventoryOption = "stock"
)

type InventoryQueryParams struct {
	Status    string
	Stock     string
	Search    string
	SortBy    string
	SortOrder SortOrder
	Page      int
	Offset    int
	Limit     int
}

func (r InventoryListRequest) ToQuery() InventoryQueryParams {
	return InventoryQueryParams{
		Status:    r.Status,
		Stock:     r.Stock,
		Search:    r.Search,
		Page:      r.GetPage(),
		Offset:    r.GetOffset(),
		Limit:     r.GetPerPage(),
		SortBy:    r.SortBy,
		SortOrder: r.SortOrder,
	}
}

type InventoryAction string

const (
	Restock    InventoryAction = "restock"
	Adjustment InventoryAction = "adjustment"
)

type CreateInventoryLogRequest struct {
	SKUID          uint            `json:"sku_id"`
	Action         InventoryAction `json:"action"`
	QuantityChange int             `json:"quantity_change"`
	Notes          string          `json:"notes"`
}