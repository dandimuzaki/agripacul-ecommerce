import { Totals } from "./checkout";
import { variantCombination } from "./sku";

export interface OrderSummary {
  id: number;
  customer: Customer;
  created_at: string;
  updated_at: string;
  display_status: string;
  grand_total: number;
  first_item: OrderItem;
  item_count: number;
  cancellation: boolean | null;
  tracking_number?: string;
  can_review: boolean;
}

export interface Customer {
  name: string;
  email: string;
}

export interface OrderItem {
  id: number,
  sku_id: number,
  name: string;
	main_image_url: string;
	variants: variantCombination[];
	quantity: number;
	price: number;
	total_price: number;
}

export interface OrderStep {
  key: string,
  label: string,
  done: boolean,
  at: string,
  description: string
}

export interface BillingAddress {
  recipient_name: string,
  label: string,
  phone_number: string,
  province: string,
  regency: string,
  district: string,
  subdistrict: string,
  postal_code: string,
  detail_address: string
}

export interface Shipping {
  name: string,
  code: string,
  service: string,
  cost: number,
  etd: number
}

export interface OrderDetails {
  id: number,
  customer: Customer,
  display_status: string,
  steps: OrderStep[],
  cancellation: boolean | null,
  created_at: string,
  billing_address: BillingAddress,
  items: OrderItem[],
  shipping: Shipping,
  totals: Totals,
  notes: string,
  tracking_number?: string,
  can_review: boolean
}