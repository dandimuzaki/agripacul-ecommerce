import z from "zod";

export const reviewSchema = z.object({
  sku_id: z.number("Product id is required"),
  rating: z.number().min(1, "Minimum rating is 1").max(5, "Maximum rating is 5"),
  comment: z.string().optional()
});

export const createReviewSchema = z.object({
    order_id: z.number("Order id is required"),
    reviews: z.array(reviewSchema)
})

export type CreateReviewFormValues = z.infer<typeof createReviewSchema>;

export const reviewFilterSchema = z.object({
  product_id: z.number().optional(),
  search: z.string().optional(),
  sort_by: z.string().optional(),
  sort_order: z.string().optional(),
  page: z.number().optional(),
  limit: z.number().optional(),
  sort: z.string().optional()
})

export type ReviewFilterFormValues = z.infer<typeof reviewFilterSchema>;
