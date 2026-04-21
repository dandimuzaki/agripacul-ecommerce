import { Inventory, InventoryLog } from "@/types/inventory"

export const mockInventories: Inventory[] = [
  {
    id: 1,
    product: "Laptop ASUS 124A",
    sku_code: "ASUS-BLK-128",
    variant_label: "Black / 128GB",
    stock: 15,
    min_stock: 5,
    availability: "in stock",
    status: "active"
  },
  {
    id: 2,
    product: "Laptop ASUS 124A",
    sku_code: "ASUS-BLK-256",
    variant_label: "Black / 256GB",
    stock: 7,
    min_stock: 5,
    availability: "in stock",
    status: "active"
  },
  {
    id: 3,
    product: "Laptop ASUS 124A",
    sku_code: "ASUS-GRY-128",
    variant_label: "Grey / 128GB",
    stock: 2,
    min_stock: 5,
    availability: "low stock",
    status: "active"
  },
  {
    id: 4,
    product: "Laptop ASUS 124A",
    sku_code: "ASUS-GRY-256",
    variant_label: "Grey / 256GB",
    stock: 0,
    min_stock: 5,
    availability: "out of stock",
    status: "active"
  },
  {
    id: 5,
    product: "iPhone 15",
    sku_code: "IP15-BLK-128",
    variant_label: "Black / 128GB",
    stock: 25,
    min_stock: 10,
    availability: "in stock",
    status: "active"
  },
  {
    id: 6,
    product: "iPhone 15",
    sku_code: "IP15-BLU-256",
    variant_label: "Blue / 256GB",
    stock: 4,
    min_stock: 8,
    availability: "low stock",
    status: "active"
  }
]

export const mockInventory: Inventory = {
  id: 6,
  product: "iPhone 15",
  sku_code: "IP15-BLU-256",
  variant_label: "Blue / 256GB",
  stock: 4,
  min_stock: 8,
  availability: "low stock",
  status: "active"
}

export const mockInventoryLogs: InventoryLog[] = [
  {
    id: 1,
    sku_id: 1,
    type: "restock",
    quantity_change: 20,
    current_stock_after: 20,
    reference_id: "PO-20240301",
    reference_type: "purchase_order",
    notes: "Initial stock from supplier",
    created_at: "2024-03-01T09:15:00Z"
  },
  {
    id: 2,
    sku_id: 1,
    type: "order",
    quantity_change: -2,
    current_stock_after: 18,
    reference_id: "ORD-20240305",
    reference_type: "order",
    notes: "Customer purchase",
    created_at: "2024-03-05T12:30:00Z"
  },
  {
    id: 3,
    sku_id: 1,
    type: "adjustment",
    quantity_change: -3,
    current_stock_after: 15,
    reference_type: "manual",
    notes: "Damaged items removed",
    created_at: "2024-03-06T10:20:00Z"
  },
  {
    id: 4,
    sku_id: 3,
    type: "order",
    quantity_change: -1,
    current_stock_after: 2,
    reference_id: "ORD-20240308",
    reference_type: "order",
    notes: "Customer purchase",
    created_at: "2024-03-08T15:45:00Z"
  },
  {
    id: 5,
    sku_id: 6,
    type: "restock",
    quantity_change: 10,
    current_stock_after: 10,
    reference_id: "PO-20240310",
    reference_type: "purchase_order",
    notes: "New stock arrived",
    created_at: "2024-03-10T08:00:00Z"
  },
  {
    id: 6,
    sku_id: 6,
    type: "order",
    quantity_change: -6,
    current_stock_after: 4,
    reference_id: "ORD-20240311",
    reference_type: "order",
    notes: "Customer purchase",
    created_at: "2024-03-11T13:10:00Z"
  }
]