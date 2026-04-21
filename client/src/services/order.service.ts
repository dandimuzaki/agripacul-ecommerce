import { api } from "@/lib/api";
import { FilterOrderFormValues, OrderFormValues } from "@/schemas/order.schema";
import { Response } from "@/types/response";

export const orderService = {
  getOrderHistory(params: FilterOrderFormValues): Promise<Response> {
    return api.get(`/orders`, { params })
  },

  getAllOrders(params: FilterOrderFormValues): Promise<Response> {
    return api.get(`/admin/orders`, { params })
  },

  getOrderDetails(id: number): Promise<Response> {
    return api.get(`/orders/${id}`)
  },

  getOrderById(id: number): Promise<Response> {
    return api.get(`/admin/orders/${id}`)
  },

  createOrder(payload: OrderFormValues) {
    return api.post("/orders", payload)
  },

  confirmOrder(id: number) {
    return api.put(`/admin/orders/${id}/confirm`)
  },

  completeOrder(id: number) {
    return api.put(`/orders/${id}/complete`)
  },

  cancelOrder(id: number) {
    return api.put(`/orders/${id}/cancel`)
  },

  cancelOrderByAdmin(id: number) {
    return api.put(`/admin/orders/${id}/cancel`)
  },
};