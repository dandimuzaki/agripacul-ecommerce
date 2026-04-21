export interface Inventory {
  id: number,
  product: string,
  sku_code: string,
  variant_label: string,
  stock: number,
  min_stock: number,
  availability: string,
  status: string
}

export interface InventoryLog {
  id: number,
  sku_id: number,
  type: string,
  quantity_change: number,
  current_stock_after: number,
  reference_id?: string,
  reference_type: string,
  notes: string,
  created_at: string
}