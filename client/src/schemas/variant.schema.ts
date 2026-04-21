import { z } from "zod"

export const variantValueSchema = z.object({
  id: z.number().optional(),
  value: z.string().min(1, "Value required"),
  status: z.string().optional()
})

export const variantTypeSchema = z.object({
  id: z.number().optional(),
  name: z.string().min(1, "Variant name required"),
  values: z.array(variantValueSchema),
  status: z.string().optional()
})

export const variantFormSchema = z.object({
  variants: z.array(variantTypeSchema),
})

export type VariantFormValues = z.infer<typeof variantFormSchema>