import { api } from "@/lib/api";
import { CartFormValues, UpdateCartFormValues } from "@/schemas/cart.schema";
import { Response } from "@/types/response";

export const cartService = {
  getCart(): Promise<Response> {
    return api.get(`/cart`)
  },

  addToCart(form: CartFormValues) {
    return api.post(`/cart/items`, form)
  },

  updateItemQuantity(itemID: number, form: UpdateCartFormValues) {
    return api.put(`/cart/items/${itemID}`, form)
  },

  updateSelectItem(itemID: number, form: UpdateCartFormValues) {
    return api.put(`/cart/items/${itemID}`, form)
  },

  removeItem(itemID: number) {
    return api.delete(`/cart/items/${itemID}`)
  },

  clearCart() {
    return api.delete(`/cart`)
  },

  batchSelectItem(payload: {is_selected: boolean}) {
    return api.put(`/cart`, payload)
  },
}