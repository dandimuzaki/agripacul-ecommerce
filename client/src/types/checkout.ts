import { ItemPrice, ItemVariant, ProductPreview } from "./cart";

export interface CheckoutItem {
  id: number;
  sku_id: number;
  sku_code: string;
  product: ProductPreview;
  variants: ItemVariant[];
  price: ItemPrice;
  quantity: number;
  stock: number;
  subtotal: number;
}

export interface ShippingOption {
  name: string;
  code: string;
  service: string;
  description: string;
  cost: number;
  etd: string;
}

export interface ShippingOptionTemp {
  name: string;
  code: string;
  service: string;
  description: string;
  cost: number;
  etd: string;
  id: string;
}

export interface Promotion {
  id: number;
  name: string;
  end_date: string;
  discount_type: "percentage" | "amount";
  discount_value: number;
  minimum_order_value: number;
  maximum_discount: number;
  usage_left: number;
  voucher_code: string | null;
}

export interface Totals {
  payment_method?: string;
  subtotal: number;
  discount_amount: number;
  shipping_cost: number;
  grand_total: number;
}

export interface CheckoutPreviewResponse {
  selected_items: CheckoutItem[];
  totals: Totals;
}