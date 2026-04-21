import { BrowseProductFormValues } from "@/schemas/product.schema";

export const productKeys = {
  all: ["products"] as const,

  lists: () => [...productKeys.all, "lists"] as const,

  list: (filters: BrowseProductFormValues) =>
    [...productKeys.lists(), filters] as const,

  bestSelling: () =>
    [...productKeys.lists(), "best_selling"] as const,

  detail: (slug: string) =>
    [...productKeys.all, "detail", slug] as const,

  adminDetail: (id: number) =>
    [...productKeys.all, "admin_detail", id] as const,

  infinite: (data: any) =>
    [...productKeys.all, "infinite", data] as const,

  sku: (id: number) =>
    [...productKeys.all, "sku", id] as const,
}