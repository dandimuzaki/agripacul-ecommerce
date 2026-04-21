"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { Controller, useForm } from "react-hook-form"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardFooter,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupText,
  InputGroupTextarea,
} from "@/components/ui/input-group"
import { ProductFormValues, productSchema } from "@/schemas/product.schema"
import { useEffect, useReducer } from "react"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { VariantField, variantReducer } from "../components/variant-field"
import InputTags from "../components/input-tags"
import { useCreateProduct } from "@/hooks/product/useCreateProduct"
import { Spinner } from "@/components/ui/spinner"
import { useCategories } from "@/hooks/category/useCategories"
import { Category } from "@/types/category"

export default function AddProductForm() {
  const { data: categories } = useCategories()

  const values: ProductFormValues = {
    name: "",
    description: "",
    category_id: undefined,
    tags: [],
    variants: []
  }

  const form = useForm<ProductFormValues>({
    resolver: zodResolver(productSchema),
    defaultValues: values,
  })

  const { mutate, isPending } = useCreateProduct()
    
  const onCreateProduct = (data: ProductFormValues) => {
    mutate(data)
  }

  const [variantState, dispatchVariant] = useReducer(
    variantReducer,
    []
  );

  useEffect(() => {
    if (categories?.data.length === 0) return;
  }, [categories]);

  return (
    <div className="space-y-2">
      <h2 className="text-xl font-bold">Add New Product</h2>
      <Card className="w-full">
        <CardContent className="space-y-4">
          <form id="add-product" onSubmit={form.handleSubmit(onCreateProduct)}>
            <FieldGroup>
              <Controller
                name="name"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="add-product-name">
                      Name
                    </FieldLabel>
                    <Input
                      {...field}
                      id="add-product-name"
                      aria-invalid={fieldState.invalid}
                      placeholder="Salad cup Japanese style"
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <Controller
                name="category_id"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel>Category</FieldLabel>

                    <Select
                      value={field.value?.toString()}
                      onValueChange={(value) => field.onChange(Number(value))}
                    >
                      <SelectTrigger aria-invalid={fieldState.invalid}>
                        <SelectValue placeholder="Select category" />
                      </SelectTrigger>

                      <SelectContent>
                        {categories?.data.map((category: Category) => (
                          <SelectItem
                            key={category.id}
                            value={category.id.toString()}
                          >
                            {category.name}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>

                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />

              <InputTags form={form} />

              <Controller
                name="description"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="edit-product-description">
                      Description
                    </FieldLabel>
                    <InputGroup>
                      <InputGroupTextarea
                        {...field}
                        id="add-product-description"
                        placeholder="Tell customers about your product — its freshness, origin, taste, and how it’s best enjoyed"
                        rows={6}
                        className="min-h-24 resize-none"
                        aria-invalid={fieldState.invalid}
                      />
                      <InputGroupAddon align="block-end">
                        <InputGroupText className="tabular-nums">
                          {field?.value?.length}/100 characters
                        </InputGroupText>
                      </InputGroupAddon>
                    </InputGroup>
                    <FieldDescription>
                      Include steps to reproduce, expected behavior, and what
                      actually happened.
                    </FieldDescription>
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />

              <FieldLabel className="mb-[-8px]">
                Variants
              </FieldLabel>
              <VariantField 
                form={form} 
                variantState={variantState} 
                dispatchVariant={dispatchVariant}
              />
              </FieldGroup>
          </form>
        </CardContent>
        <CardFooter>
          <Field orientation="horizontal" className="w-full flex items-center justify-center">
            <Button type="button" className="bg-gray-200 text-black hover:bg-gray-300 hover:text-black" onClick={() => form.reset()}>
              Reset
            </Button>
            <Button className="text-white hover:text-white" type="submit" form="add-product">
              {isPending ? <><Spinner/>Saving product...</> : <>Submit</>}
            </Button>
          </Field>
        </CardFooter>
      </Card>
    </div>
  )
}
