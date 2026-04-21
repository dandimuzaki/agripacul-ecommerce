import { ReviewFilterFormValues } from "@/schemas/review.schema";

export const reviewKeys = {
  all: ["review"] as const,

  lists: () => [...reviewKeys.all, "lists"] as const,
  
  list: (filters: ReviewFilterFormValues) =>
    [...reviewKeys.lists(), filters] as const,

  detail: (id: number) =>
    [...reviewKeys.all, "detail", id] as const
}