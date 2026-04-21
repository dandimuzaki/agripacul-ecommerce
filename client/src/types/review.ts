export interface Review {
  id: number,
  product_id: number,
  product_name: string,
  order_id: number,
  rating: number,
  comment: string,
  customer_name: string,
  created_at: string
}