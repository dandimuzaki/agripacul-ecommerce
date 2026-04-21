import { api } from "@/lib/api";
import { CheckoutFormValues, PromotionFormValues, ShippingFormValues } from "@/schemas/checkout.schema";
import { Response } from "@/types/response";

export const checkoutService = {
  async previewCheckout(payload: CheckoutFormValues): Promise<Response> {
    return api
      .post(`/checkout/preview`, payload)
  },
  async getShippingOptions(payload: ShippingFormValues): Promise<Response> {
    return api.post(`/checkout/shippings`, payload)
  },
  async getValidPromotions(payload: PromotionFormValues): Promise<Response> {
    return api.post(`/checkout/promotions`, payload)
  },
};