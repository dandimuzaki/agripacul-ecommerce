import { api } from "@/lib/api";
import { Response } from "@/types/response";

export const paymentService = {
  getPaymentMethodList(): Promise<Response> {
    return api.get("/payment-methods");
  },
};