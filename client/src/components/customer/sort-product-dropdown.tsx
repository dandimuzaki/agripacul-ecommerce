"use client"

import { useProductFilter } from "@/hooks/product/useProductFilter"
import { Controller } from "react-hook-form"
import { Field, FieldLabel } from "../ui/field"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect } from "react"

const SortProductDropdown = () => {
  const { form } = useProductFilter()
  const searchParams = useSearchParams()
  const router = useRouter()
  const sortOptions = [
    {name: "Most Relevant", value: "relevance" },
    {name: "Lowest Price", value: "price_asc" },
    {name: "Highest Price", value: "price_desc" },
    {name: "Best Rated", value: "rating_desc" },
    {name: "Best Selling", value: "sold_asc" },
  ]

  const getSortValueFromParams = (searchParams: URLSearchParams) => {
    const sort_by = searchParams.get("sort_by") ?? undefined
    const sort_order = searchParams.get("sort_order") ?? undefined

    if (!sort_by || !sort_order) return "relevance"

    return `${sort_by}_${sort_order}`
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
          {/* <FieldLabel className="text-md border-b border-gray-400 pb-2">Sort By</FieldLabel> */}
          <Select
            value={field.value ?? "relevance"}
            onValueChange={(value) => {
              field.onChange(value)
              const params = new URLSearchParams(searchParams.toString())
              if (value !== "relevance") {
                const sort = value.split("_")
                params.delete("page")
                params.set("sort_by", sort[0])
                params.set("sort_order", sort[1])
              } else {
                params.delete("sort_by")
                params.delete("sort_order")
              }
              router.replace(`/products?${params.toString()}`)
            }}
          >
            <SelectTrigger aria-invalid={fieldState.invalid} className="border border-gray-400 w-fit">
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

export default SortProductDropdown
