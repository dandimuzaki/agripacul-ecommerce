import { useInfiniteQuery } from "@tanstack/react-query"
import { productService } from "@/services/product.service"
import { Response } from "@/types/response"
import { BrowseProductFormValues } from "@/schemas/product.schema"
import { useMemo } from "react"
import { productKeys } from "../queries/productKeys"

export const useInfiniteProducts = (filters: BrowseProductFormValues) => {
  const cleanFilters = useMemo(() => {
    return Object.fromEntries(
      Object.entries(filters).filter(([_, v]) => v !== undefined && v !== "")
    )
  }, [filters])

  return useInfiniteQuery<Response>({
    queryKey: productKeys.infinite(cleanFilters),

    queryFn: ({ pageParam = 1 }) =>
      productService.getProducts({
        ...cleanFilters,
        page: pageParam as number,
        limit: 12
      }),

    initialPageParam: 1,

    getNextPageParam: (lastPage) => {
      if (lastPage?.pagination?.has_next_page) {
        return lastPage.pagination.page + 1
      }
      return undefined
    },

    placeholderData: (prev) => prev
  })
}