export interface ItemVariant {
  name: string;
  value: string;
}

export interface ProductPreview {
  id: number;
  name: string;
  slug: string;
  main_image_url: string;
}

export interface ItemPrice {
  unit_price: number;
  sale_price: number;
  discount_percentage: number;
}

export interface Item {
  id: number;
  sku_id: number;
  sku_code: string;
  product: ProductPreview;
  variants: ItemVariant[];
  price: ItemPrice;
  quantity: number;
  stock: number;
  subtotal: number;
  is_selected: boolean;
  is_available: boolean;
}

export interface CartSummary {
  total_items: number;
  total_price: number;
  total_selected_items: number;
  total_selected_price: number;
}

export interface Cart {
  id: number,
  customer_id: number,
  items: Item[],
  summary: CartSummary
}