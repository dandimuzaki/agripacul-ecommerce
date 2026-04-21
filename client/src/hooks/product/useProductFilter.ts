"use client";

import { useSearchParams } from "next/navigation"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import {
  BrowseProductFormValues,
  browseProductSchema
} from "@/schemas/product.schema"
import { useEffect, useMemo } from "react";

export const useProductFilter = () => {
  const searchParams = useSearchParams()

  const search = searchParams.get("search")
  const category_id = searchParams.get("category_id")
  const min_price = searchParams.get("min_price")
  const max_price = searchParams.get("max_price")
  const rating = searchParams.get("rating")
  const page = searchParams.get("page")
  const limit = searchParams.get("limit")
  const sort_by = searchParams.get("sort_by")
  const sort_order = searchParams.get("sort_order")

  const filters = useMemo(() => ({
    search: search ?? undefined,
    category_id: category_id
      ? Number(category_id)
      : undefined,
    min_price: min_price
      ? Number(min_price)
      : undefined,
    max_price: max_price
      ? Number(max_price)
      : undefined,
    rating: rating
      ? Number(rating)
      : undefined,
    page: page
      ? Number(page)
      : undefined,
    limit: limit
      ? Number(limit)
      : undefined,
    sort_by: sort_by ?? undefined,
    sort_order: sort_order ?? undefined,
    sort: "relevance"
  }), [search, category_id, min_price, max_price, rating, page, limit, sort_by, sort_order])

  const form = useForm<BrowseProductFormValues>({
    resolver: zodResolver(browseProductSchema),
    defaultValues: filters
  })

  useEffect(() => {
    form.reset(filters)
  }, [filters])

  return {
    form,
    filters,
  }
}