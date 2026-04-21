"use client"

import { Controller } from "react-hook-form"
import { Field, FieldLabel } from "../ui/field"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect } from "react"
import { useOrderFilter } from "@/hooks/order/useOrderFilter"

const SortOrderDropdown = () => {
  const { form } = useOrderFilter()
  const searchParams = useSearchParams()
  const router = useRouter()
  const sortOptions = [
    {name: "Newest to Oldest", value: "created_at-desc" },
    {name: "Oldest to Newest", value: "created_at-asc" },
    {name: "Recently Modified", value: "updated_at-desc" },
    {name: "Lowest Total", value: "total-asc" },
    {name: "Highest Total", value: "total-desc" },
  ]

  const getSortValueFromParams = (searchParams: URLSearchParams) => {
    const sort_by = searchParams.get("sort_by") ?? undefined
    const sort_order = searchParams.get("sort_order") ?? undefined

    if (!sort_by || !sort_order) return "created_at-desc"

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
          {/* <FieldLabel className="text-md border-b border-gray-400 pb-2">Sort By</FieldLabel> */}
          <Select
            value={field.value}
            onValueChange={(value) => {
              field.onChange(value)
              const params = new URLSearchParams(searchParams.toString())
              const sort = value.split("-")
              params.delete("page")
              params.set("sort_by", sort[0])
              params.set("sort_order", sort[1])
              router.replace(`/orders?${params.toString()}`)
            }}
          >
            <SelectTrigger aria-invalid={fieldState.invalid} className="border border-gray-500 w-fit">
              Sort by: <SelectValue placeholder="Select sort option" defaultValue={"Newest to Oldest"} />
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

export default SortOrderDropdown
