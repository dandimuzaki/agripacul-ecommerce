"use client"

import { Controller } from "react-hook-form"
import { Field } from "../ui/field"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect } from "react"
import { useInventoryFilter } from "@/hooks/inventory/useInventoryFilter"

const SortInventoryDropdown = () => {
  const { form } = useInventoryFilter()
  const searchParams = useSearchParams()
  const router = useRouter()
  const sortOptions = [
    {name: "Newest to Oldest", value: "created_at-desc" },
    {name: "Oldest to Newest", value: "created_at-asc" },
    {name: "Recently Modified", value: "updated_at-desc" },
    {name: "Lowest Stock", value: "stock-asc" },
    {name: "Highest Stock", value: "stock-desc" },
    {name: "A to Z", value: "name-asc" },
    {name: "Z to A", value: "name-desc" },
  ]

  const getSortValueFromParams = (searchParams: URLSearchParams) => {
    const sort_by = searchParams.get("sort_by") ?? undefined
    const sort_order = searchParams.get("sort_order") ?? undefined

    if (!sort_by || !sort_order) return "updated_at-desc"

    return `${sort_by}-${sort_order}`
  }

  const currentSort = getSortValueFromParams(searchParams)

  useEffect(() => {
    form.setValue("sort", currentSort)
  }, [currentSort])

  return (
    <Controller
      name="sort"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <Select
            value={field.value}
            onValueChange={(value) => {
              field.onChange(value)
              const params = new URLSearchParams(searchParams.toString())
              const sort = value.split("-")
              params.delete("page")
              params.set("sort_by", sort[0])
              params.set("sort_order", sort[1])
              router.replace(`/admin/inventory?${params.toString()}`)
            }}
          >
            <SelectTrigger aria-invalid={fieldState.invalid}>
              Sort by: <SelectValue placeholder="Select sort option" />
            </SelectTrigger>

            <SelectContent>
              {sortOptions.map((s, i) =>
                <SelectItem key={i} value={s.value}>
                  {s.name}
                </SelectItem>
              )}
            </SelectContent>
          </Select>
        </Field>
      )}
    />
  )
}

export default SortInventoryDropdown
