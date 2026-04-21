import z from "zod";
import { shippingSchema } from "./order.schema";

export const shippingSchemaTemp = z.object({
  name: z.string().optional(),
  code: z.string().optional(),
  service: z.string().optional(),
  description: z.coerce.string().optional(),
  cost: z.coerce.number().optional(),
  etd: z.coerce.string().optional(),
  id: z.coerce.string().optional()
})

export const shippingListSchema = z.object({
  shipping_address_id: z.coerce.number().optional(),
});

export type ShippingFormValues = z.input<typeof shippingListSchema>;

export const checkoutSchema = z.object({
  selected_shipping_option: shippingSchema.optional(),
  selected_payment_method_id: z.coerce.number().optional(),
  selected_promotion_id: z.coerce.number().optional(),
});

export type CheckoutFormValues = z.input<typeof checkoutSchema>;

export const checkoutSchemaTemp = z.object({
  shipping_address_id: z.coerce.number().optional(),
  selected_shipping_option_id: z.coerce.string().optional(),
  selected_shipping_option: shippingSchemaTemp.optional(),
  selected_payment_method_id: z.coerce.number().optional(),
  selected_promotion_id: z.coerce.number().optional(),
});

export type CheckoutFormValuesTemp = z.input<typeof checkoutSchemaTemp>;

export const promotionListSchema = z.object({
  page: z.number().optional(),
  limit: z.number().optional(),
  sort_by: z.string().optional(),
  sort_order: z.string().optional(),
});

export type PromotionFormValues = z.input<typeof promotionListSchema>;