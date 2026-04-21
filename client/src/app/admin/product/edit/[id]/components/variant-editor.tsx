'use client';
import { useEffect, useMemo, useReducer } from "react"
import { VariantTypeEdit } from "@/types/variant";
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"

import {
  Card,
  CardContent,
  CardFooter,
} from "@/components/ui/card"
import { normalizeVariant, VariantField, variantReducer } from "../../../components/variant-field";
import { ProductDetails } from "@/types/product";
import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";
import { ProductFormValues, productSchema } from "@/schemas/product.schema";
import { Spinner } from "@/components/ui/spinner";

export function VariantEditor({
  product,
  onUpdateProduct,
  isPending
}: {
  product: ProductDetails,
  onUpdateProduct: (product: ProductFormValues) => void,
  isPending: boolean
}) {
  const variants = useMemo(() => {
    if (!product) return []
    return normalizeVariant(product).variants
  }, [product])

  const [variantState, dispatchVariant] = useReducer(
    variantReducer,
    variants
  );

  const form = useForm<ProductFormValues>({
    resolver: zodResolver(productSchema),
    defaultValues: {
      name: product.name,
      category_id: product.category.id,
      description: product.description,
      tags: product.tags,
      variants: [],
    },
  })

  useEffect(() => {
    dispatchVariant({ type: "INIT", payload: variants })

    form.reset({
      name: product.name,
      category_id: product.category.id,
      description: product.description,
      tags: product.tags,
      variants: variants.map(v => ({
        id: v.id,
        name: v.name,
        values: v.values.map(val => ({
          id: val.id,
          value: val.value,
        })),
        status: v.status
      })),
    })
  }, [variants, form, product])

  const buildVariantPayload = (variants: VariantTypeEdit[]): ProductFormValues => {
    if (!product) {
      throw new Error("Product not loaded")
    }

    return {
      name: product.name,
      category_id: product.category.id,
      tags: product.tags,
      description: product.description,
      variants: variants.map(({ tempId, isEditing, values, ...variant }) => ({
        ...variant,
        values: values.map(({ tempId, isEditing, ...val }) => val),
      }))
    }
  }

  const onSaveVariant = () => {
    const payload = buildVariantPayload(variantState)

    const parsed = productSchema.safeParse(payload)

    if (!parsed.success) {
      console.error(parsed.error)
      return
    }

    onUpdateProduct(parsed.data)
  }

  return (
    <div className="space-y-2">
      <h2 className="text-xl font-bold">Edit Product Variant</h2>
      <Card className="w-full">
        <CardContent className="space-y-4">
          <form id="edit-variant" onSubmit={form.handleSubmit(onSaveVariant)}>
            <VariantField
              form={form}
              variantState={variantState}
              dispatchVariant={dispatchVariant}
            />
          </form>
          <CardFooter>
            <Field orientation="horizontal" className="w-full flex items-center justify-center">
              <Button className="text-white" type="submit" form="edit-variant">
                {isPending ? <><Spinner/>Saving...</> : <>Save Variant Updates</>}
              </Button>
            </Field>
          </CardFooter>
        </CardContent>
      </Card>
    </div>
  )
}
