import React, { useEffect } from 'react'
import { Controller } from 'react-hook-form'
import { Field, FieldLabel } from '../ui/field'
import { RadioGroup, RadioGroupItem } from '../ui/radio-group'
import { useProductFilter } from '@/hooks/product/useProductFilter'
import { useRouter, useSearchParams } from 'next/navigation'
import { useCategories } from '@/hooks/category/useCategories'
import { Category } from '@/types/category'
import { capitalizeAll } from '@/lib/formatText'

const CategoryFilter = () => {
  const { data } = useCategories()
  const { form } = useProductFilter()
  const searchParams = useSearchParams()
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
    <Controller
      name="category_id"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <FieldLabel className="text-md border-b border-gray-400 pb-2">
            Category
          </FieldLabel>
          <RadioGroup
            value={field.value?.toString()}
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
            className='gap-0'
          >
            <div className={`flex items-center gap-2 p-2 rounded ${(field.value == 0) ? 'bg-primary/20' : ''}`}>
                <RadioGroupItem className="border-gray-400" value={String(0)} id={String(0)} />
                <label htmlFor={String(0)} className="flex-1 cursor-pointer text-sm">
                  All Categories
                </label>
              </div>
            {categories?.map((category: Category) => (
              <div key={category.id} className={`flex items-center gap-2 p-2 rounded ${(field.value == category.id) ? 'bg-primary/20' : ''}`}>
                <RadioGroupItem className="border-gray-400" value={String(category.id)} id={String(category.id)} />
                <label htmlFor={String(category.id)} className="flex-1 cursor-pointer text-sm">
                  {capitalizeAll(category.name)}
                </label>
              </div>
            ))}
          </RadioGroup>
        </Field>
      )}
    />
  )
}

export default CategoryFilter
