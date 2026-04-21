import z from "zod"

export const skuSchema = z.object({
  sku_code: z.string().min(1, "SKU Code required"),
  price: z.number().min(1, "Price required"),
  sale_price: z.number().optional(),
  stock: z.number().min(1, "Stock required"),
  min_stock: z.number(),
  weight: z.number().min(1, "Weight required"),
  status: z.string().min(1, "Status required"),
})

export const skuFormSchema = z.object({
  skus: z.array(skuSchema),
})

export type SKUForm = z.input<typeof skuSchema>

export type SKUFormValues = {
  skus: SKUForm[]
}