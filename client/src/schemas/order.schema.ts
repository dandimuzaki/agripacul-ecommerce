import z from "zod";

export const shippingSchema = z.object({
  name: z.string().optional(),
  code: z.string().optional(),
  service: z.string().optional(),
  description: z.coerce.string().optional(),
  cost: z.coerce.number().optional(),
  etd: z.coerce.string().optional(),
})

export const orderSchema = z.object({
  shipping_address_id: z.coerce.number().optional(),
  selected_shipping_option: shippingSchema.optional(),
  selected_payment_method_id: z.coerce.number().optional(),
  selected_promotion_id: z.coerce.number().optional(),
});

export type OrderFormValues = z.input<typeof orderSchema>;


export const filterOrderSchema = z.object({
  status: z.string().optional(),
  search: z.string().optional(),
  start_date: z.string().optional(),
  end_date: z.string().optional(),
  period: z.string().optional(),
  shipping_method: z.string().optional(),
  page: z.number().optional(),
  limit: z.number().optional(),
  sort_by: z.string().optional(),
  sort_order: z.string().optional(),
  sort: z.string().optional(),
})

export type FilterOrderFormValues = z.input<typeof filterOrderSchema>;