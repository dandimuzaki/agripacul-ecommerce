"use client";

import { useSearchParams } from "next/navigation"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import {
  FilterInventoryFormValues,
  filterInventorySchema
} from "@/schemas/inventory.schema"
import { useEffect, useMemo } from "react";

export const useInventoryFilter = () => {
  const searchParams = useSearchParams()

  const search = searchParams.get("search")
  const status = searchParams.get("status")
  const stock = searchParams.get("stock")
  const page = searchParams.get("page")
  const limit = searchParams.get("limit")
  const sort_by = searchParams.get("sort_by")
  const sort_order = searchParams.get("sort_order")

  const filters = useMemo(() => ({
    search: search ?? undefined,
    status: status ?? undefined,
    stock: stock ?? undefined,
    page: page
      ? Number(page)
      : undefined,
    limit: limit
      ? Number(limit)
      : undefined,
    sort_by: sort_by ?? undefined,
    sort_order: sort_order ?? undefined,
    sort: "updated_at-desc"
  }), [search, page, limit, sort_by, sort_order])

  const form = useForm<FilterInventoryFormValues>({
    resolver: zodResolver(filterInventorySchema),
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