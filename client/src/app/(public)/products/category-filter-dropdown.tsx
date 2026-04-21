"use client"

import { useProductFilter } from "@/hooks/product/useProductFilter"
import { Controller } from "react-hook-form"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect } from "react"
import { useCategories } from "@/hooks/category/useCategories"
import { capitalizeAll } from "@/lib/formatText"
import { Field } from "@/components/ui/field"

const CategoryFilterDropdown = () => {
  const { form } = useProductFilter()
  const searchParams = useSearchParams()
  const {data} = useCategories()
  const router = useRouter()

  const categories = data?.data

  const categoryId = searchParams.get("category_id") ?? undefined

  useEffect(() => {
    if (categoryId) {
      form.setValue("category_id", Number(categoryId))
    } else {
      form.setValue("category_id", 0)
    }
  }, [categoryId, form])

  return (
    <div className="w-fit">
    <Controller
      name="category_id"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <Select
            value={String(field.value)}
            onValueChange={(value) => {
              field.onChange(Number(value))
              const params = new URLSearchParams(searchParams.toString())
              if (value !== "0") {
                params.delete("page")
                params.set("category_id", value)
              } else {
                params.delete("category_id")
                form.setValue("category_id", 0)
              }
              router.replace(`/products?${params.toString()}`)
            }}
          >
            <SelectTrigger aria-invalid={fieldState.invalid} className="w-fit border border-gray-400">
              Category: <SelectValue placeholder="Select category" />
            </SelectTrigger>

            <SelectContent>
                <SelectItem value={"0"}>
                  <span className="font-semibold">All Categories</span>
                </SelectItem>
              {categories?.map((s, i) =>
                <SelectItem key={i} value={String(s.id)}>
                  <span className="font-semibold">{capitalizeAll(s.name)}</span>
                </SelectItem>
              )}
            </SelectContent>
          </Select>
        </Field>
      )}
    />
    </div>
  )
}

export default CategoryFilterDropdown
