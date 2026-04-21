"use client";

import { useSearchParams } from "next/navigation"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import {
  ReviewFilterFormValues,
  reviewFilterSchema
} from "@/schemas/review.schema"
import { useEffect, useMemo } from "react";

export const useReviewFilter = () => {
  const searchParams = useSearchParams()

  const search = searchParams.get("search")
  const product_id = searchParams.get("product_id")
  const page = searchParams.get("page")
  const limit = searchParams.get("limit")
  const sort_by = searchParams.get("sort_by")
  const sort_order = searchParams.get("sort_order")

  const filters = useMemo(() => ({
    search: search ?? undefined,
    product_id: product_id ? Number(product_id) : undefined,
    page: page
      ? Number(page)
      : undefined,
    limit: limit
      ? Number(limit)
      : undefined,
    sort_by: sort_by ?? undefined,
    sort_order: sort_order ?? undefined,
    sort: "created_at-desc"
  }), [search, page, limit, sort_by, sort_order])

  const form = useForm<ReviewFilterFormValues>({
    resolver: zodResolver(reviewFilterSchema),
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