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
import { ProductDetails } from "@/types/product"
import { useEffect } from "react"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import InputTags from "../../../components/input-tags"
import { useCategories } from "@/hooks/category/useCategories"
import { Spinner } from "@/components/ui/spinner"
import { Category } from "@/types/category"
import BackNavigation from "@/components/common/back-navigation"
import { Skeleton } from "@/components/ui/skeleton"

export default function EditProductForm({
  product,
  onUpdateProduct,
  isPending
}: {
  product: ProductDetails,
  onUpdateProduct: (data: ProductFormValues) => void,
  isPending: boolean
}) {
  const { data: categories, isLoading } = useCategories()

  const form = useForm<ProductFormValues>({
    resolver: zodResolver(productSchema),
    defaultValues: {
      name: product.name,
      category_id: product.category.id,
      description: product.description,
      tags: product.tags,
    }
  })

  const normalizeProductInformation = (product: ProductDetails): ProductFormValues => {
    return {
      name: product.name,
      category_id: product.category.id,
      description: product.description,
      tags: product.tags,
      variants: product.variants
    };
  };

  useEffect(() => {
    if (product && !isLoading && categories?.data?.length) {
      form.reset(normalizeProductInformation(product))
    };
  }, [product, isLoading, categories, form]);

  return (
    <div className="space-y-2">
      <BackNavigation link="/admin/product"/>
      <h2 className="text-xl font-bold">Edit Product Data</h2>
      <Card className="w-full">
        <CardContent className="space-y-4">
          <form id="edit-product" onSubmit={form.handleSubmit(onUpdateProduct)}>
            <FieldGroup>
              <Controller
                name="name"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="edit-product-name">
                      Name
                    </FieldLabel>
                    <Input
                      {...field}
                      id="edit-product-name"
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
              {isLoading ? <Skeleton className="w-1/2 h-8"/> : <Controller
                name="category_id"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid} className="lg:w-1/2">
                    <FieldLabel>Category</FieldLabel>

                    <Select
                      value={field.value ? String(field.value) : undefined}
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
              />}

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
                        id="edit-product-description"
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
              </FieldGroup>
          </form>
        </CardContent>
        <CardFooter className="mt-4">
          <Field orientation="horizontal" className="w-full flex items-center justify-center">
            <Button type="button" className="bg-gray-200 text-black hover:bg-gray-300 hover:text-black" onClick={() => form.reset()}>
              Reset
            </Button>
            <Button className="text-white hover:text-white" type="submit" form="edit-product">
              {isPending ? <><Spinner/>Saving...</> : <>Save Product Updates</>}
            </Button>
          </Field>
        </CardFooter>
      </Card>
    </div>
  )
}
