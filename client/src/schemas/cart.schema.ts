import z from "zod";

export const cartSchema = z.object({
  sku_id: z.number(),
  quantity: z.number(),
})

export type CartFormValues = z.input<typeof cartSchema>;

export const updateCartSchema = z.object({
  quantity: z.number().optional(),
  is_selected: z.boolean().optional(),
})

export type UpdateCartFormValues = z.input<typeof updateCartSchema>;