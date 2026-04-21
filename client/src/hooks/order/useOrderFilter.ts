"use client";

import { useRouter, useSearchParams } from "next/navigation"
import { FilterOrderFormValues, filterOrderSchema } from "@/schemas/order.schema"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect, useMemo } from "react";

export const useOrderFilter = () => {
  const router = useRouter()
  const searchParams = useSearchParams()

  const filters = useMemo(() => ({
    search: searchParams.get("search") ?? undefined,
    status: searchParams.get("status") ?? "all",
    start_date: searchParams.get("start_date") ?? undefined,
    end_date: searchParams.get("end_date") ?? undefined,
    period: searchParams.get("period") ?? "this_month",
    shipping_method: searchParams.get("shipping_method") ?? undefined,
    sort_by: searchParams.get("sort_by") ?? undefined,
    sort_order: searchParams.get("sort_order") ?? undefined,
    page: searchParams.get("page")
      ? Number(searchParams.get("page"))
      : undefined,
    limit: searchParams.get("limit")
      ? Number(searchParams.get("limit"))
      : undefined,
    sort: "created_at-desc"
  }), [searchParams])

  const form = useForm<FilterOrderFormValues>({
    resolver: zodResolver(filterOrderSchema),
    defaultValues: filters
  })

  useEffect(() => {
    form.reset(filters)
  }, [filters])

  const onSubmitFilter = (values: FilterOrderFormValues) => {
    const params = new URLSearchParams()

    Object.entries(values).forEach(([key, value]) => {
      if (value !== "" && value !== undefined) {
        params.set(key, String(value))
      }
    })

    router.replace(`/orders?${params.toString()}`)
  }

  return {
    filters,
    form,
    onSubmitFilter,
  }
}