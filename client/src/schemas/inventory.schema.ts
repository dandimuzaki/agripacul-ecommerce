import z from "zod";

export const inventorySchema = z.object({
  action: z.string().min(3, "Action is required"),
  quantity_change: z
    .number("Quantity required")
    .int()
    .positive("Quantity must be greater than 0"),
  sku_id: z.coerce.number(),
  notes: z.string().max(100)
});

export type InventoryFormValues = z.input<typeof inventorySchema>;

export const adjustStockSchema = z.object({
  action: z.string().min(3, "Action is required"),
  quantity_change: z
    .number("Quantity required")
    .int(),
  sku_id: z.coerce.number(),
  notes: z.string().max(100)
});

export type AdjustStockFormValues = z.input<typeof adjustStockSchema>;

export const filterInventorySchema = z.object({
  search: z.coerce.string().optional(),
  status: z.string().optional(),
  stock: z.string().optional(),
  page: z.number().optional(),
  limit: z.number().optional(),
  sort_by: z.string().optional(),
  sort_order: z.string().optional(),
  sort: z.string().optional()
})

export type FilterInventoryFormValues = z.input<typeof filterInventorySchema>;

export const inventoryLogSchema = z.object({
  page: z.number().optional(),
  limit: z.number().optional(),
})

export type InventoryLogFormValues = z.input<typeof inventoryLogSchema>;
