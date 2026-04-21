import { Cart } from "./cart";
import { Category } from "./category";
import { CheckoutPreviewResponse, ShippingOption } from "./checkout";
import { Inventory } from "./inventory";
import { OrderDetails, OrderSummary } from "./order";
import { ProductDetails, ProductSummary } from "./product";
import { skuDetails } from "./sku";

export interface Pagination {
  page: number,
  limit: number,
  total: number,
  total_pages: number,
  has_next_page: boolean
}

export interface Response {
  success: boolean,
  message: string,
  data?: any,
  pagination?: Pagination,
  errors: any
}

export interface GetCategoriesResponse {
  success: boolean,
  message: string,
  data: Category[],
  pagination: Pagination
}

export interface GetCategoryResponse {
  success: boolean,
  message: string,
  data: Category,
}

export interface GetProductsResponse {
  success: boolean,
  message: string,
  data: ProductSummary[],
  pagination: Pagination
}

export interface GetProductResponse {
  success: boolean,
  message: string,
  data: ProductDetails,
}

export interface GetSKUsResponse {
  success: boolean,
  message: string,
  data: skuDetails[],
}

export interface GetCartResponse {
  success: boolean,
  message: string,
  data: Cart,
}

export interface GetOrdersResponse {
  success: boolean,
  message: string,
  data: OrderSummary[],
  pagination: Pagination
}

export interface GetOrderResponse {
  success: boolean,
  message: string,
  data: OrderDetails,
}

export interface GetInventoriesResponse {
  success: boolean,
  message: string,
  data: Inventory[],
  pagination: Pagination
}

export interface GetInventoryResponse {
  success: boolean,
  message: string,
  data: Inventory,
}

export interface GetCheckoutResponse {
  success: boolean,
  message: string,
  data: CheckoutPreviewResponse,
}

export interface GetShippingResponse {
  success: boolean,
  message: string,
  data: ShippingOption[],
}