import { PromotionFormValues } from "@/schemas/checkout.schema";

export const checkoutKeys = {
  all: ["checkout"] as const,

  preview: (payload: any) =>
    [...checkoutKeys.all, "preview", payload] as const,

  shippings: (addressId: number) =>
    [...checkoutKeys.all, "shippings", addressId] as const,

  promotions: (payload: PromotionFormValues) =>
    [...checkoutKeys.all, "promotions", payload] as const,
}