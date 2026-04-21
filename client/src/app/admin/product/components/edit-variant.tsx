'use client';
import { useReducer } from "react"
import { useForm } from "react-hook-form"
import { VariantEditor } from "./variant-editor"
import { SKUEditor } from "../sku-editor";



export type Product = {
  id?: number
  name: string
  category: Category
  variants: VariantTypeProduct[]
}

export type Draft = {
  id?: number
  name: string
  category: Category
  variants: VariantType[]
}

export type SKU = {
  id: number
  skuCode: string
  price: number
  stock: number
  status: string
  variantValues: {
    id: number
    name: string
    label: string
  }[]
}

export type SKUStatus = "active" | "inactive" | "archived"
export type SKURowStatus = "clean" | "updated"

export type SKUDraft = {
  id: number
  skuCode: string
  price: number | ""
  stock: number | ""
  status: SKUStatus

  originalSKUCode: string
  originalPrice: number
  originalStock: number
  originalStatus: SKUStatus

  rowStatus: SKURowStatus
  variantValues: {
    id: number
    name: string
    label: string
  }[]
}

type SKUAction =
  | { type: "INIT"; payload: SKUDraft[] }
  | { type: "UPDATE_SKU_CODE"; id: number; skuCode: string }
  | { type: "UPDATE_PRICE"; payload: {id: number; price: number | ""} }
  | { type: "UPDATE_STOCK"; payload: {id: number; stock: number | ""} }
  | { type: "TOGGLE_STATUS"; id: number }

export function skuReducer(
  state: SKUDraft[],
  action: SKUAction
): SKUDraft[] {
  switch (action.type) {
    case "INIT":
    return action.payload.map(sku => ({
      ...sku,
      price: sku.price,
      stock: sku.stock,
      rowStatus: "clean",
    }))

    case "UPDATE_SKU_CODE":
      return state.map(sku =>
        sku.id === action.id
          ? {
              ...sku,
              skuCode: action.skuCode,
              rowStatus:
                action.skuCode === sku.originalSKUCode ? "clean" : "updated",
            }
          : sku
      )

    case "UPDATE_PRICE":
      return state.map(sku =>
        sku.id === action.payload.id
          ? {
              ...sku,
              price: action.payload.price,
              rowStatus:
                action.payload.price === sku.originalPrice
                  ? "clean"
                  : "updated",
            }
          : sku
      )

    case "UPDATE_STOCK":
      return state.map(sku =>
        sku.id === action.payload.id
          ? {
              ...sku,
              stock: action.payload.stock,
              rowStatus:
                action.payload.stock === sku.originalStock
                  ? "clean"
                  : "updated",
            }
          : sku
      )

    case "TOGGLE_STATUS":
      return state.map(sku => {
        if (sku.id !== action.id) return sku
        if (sku.status === "archived") return sku

        const nextStatus: SKUStatus =
          sku.status === "active" ? "inactive" : "active"

        const nextRowStatus: SKURowStatus =
          nextStatus === sku.originalStatus ? "clean" : "updated"

        return {
          ...sku,
          status: nextStatus,
          rowStatus: nextRowStatus,
        }
      })

    default:
      return state
  }
}

export function EditProductModal() {
  const product = {
    id: 1,
    name: "Laptop",
    category: {
      id: 2,
      name: "electronic"
    },
    variants: [
      {
        id: 1,
        name: "color",
        values: [
          {
            id: 10,
            label: "black"
          },
          {
            id: 11,
            label: "grey"
          }
        ]
      },
      {
        id: 2,
        name: "storage",
        values: [
          {
            id: 15,
            label: "128GB"
          },
          {
            id: 16,
            label: "256GB"
          }
        ]
      }
    ],
    skus: [
      {
        id: 1,
        skuCode: "LPTP-BLACK-128GB",
        price: 0,
        stock: 0,
        status: "active",
        variantValues: [
          {
            id: 10,
            name: "color",
            label: "black"
          },
          {
            id: 15,
            name: "storage",
            label: "128GB"
          }
        ]
      },
      {
        id: 2,
        skuCode: "LPTP-GREY-128GB",
        price: 0,
        stock: 0,
        status: "active",
        variantValues: [
          {
            id: 11,
            name: "color",
            label: "grey"
          },
          {
            id: 15,
            name: "storage",
            label: "128GB"
          }
        ]
      },
      {
        id: 3,
        skuCode: "LPTP-BLACK-256GB",
        price: 0,
        stock: 0,
        status: "active",
        variantValues: [
          {
            id: 10,
            name: "color",
            label: "black"
          },
          {
            id: 16,
            name: "storage",
            label: "256GB"
          }
        ]
      },
      {
        id: 4,
        skuCode: "LPTP-GREY-256GB",
        price: 0,
        stock: 0,
        status: "active",
        variantValues: [
          {
            id: 11,
            name: "color",
            label: "grey"
          },
          {
            id: 16,
            name: "storage",
            label: "256GB"
          }
        ]
      }
    ]
  }

  const normalizeProduct = (product: Product): Draft => ({
    ...product,
    variants: product.variants.map(v => ({
      ...v,
      tempId: crypto.randomUUID(),
      originalName: v.name,
      status: "clean",
      values: v.values.map(val => ({
        ...val,
        tempId: crypto.randomUUID(),
        originalLabel: val.label,
        status: "clean",
      })),
    })),
  })

  const [variantState, dispatchVariant] = useReducer(
    variantReducer,
    normalizeProduct(product).variants
  )

  const form = useForm({
    defaultValues: {
      name: product?.name,
      category_id: product?.category.id,
    },
  })

  const buildVariantPayload = (variants: VariantType[]) => {
  return variants
    .filter(v => v.status !== "deleted")
    .map(({ tempId, isEditing, values, ...variant }) => ({
      ...variant,
      values: values
        .filter(val => val.status !== "deleted")
        .map(({ tempId, isEditing, ...val }) => val),
    }))
  }

  const onSubmit = () => {
    const payload = {
      name: form.getValues("name"),
      category_id: form.getValues("category_id"),
      variants: buildVariantPayload(variantState)
    }

  }

  const normalizeSKUs = (skus: SKU[]): SKUDraft[] => {
    return skus.map((sku): SKUDraft => ({
      id: sku.id,
      skuCode: sku.skuCode,
      price: sku.price,
      stock: sku.stock,
      status: sku.status as SKUStatus,

      originalSKUCode: sku.skuCode,
      originalPrice: sku.price,
      originalStock: sku.stock,
      originalStatus: sku.status as SKUStatus,

      rowStatus: "clean" as SKURowStatus,
      variantValues: sku.variantValues,
    }))
  }

  return (
    <div>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <input {...form.register("name")} />

        <VariantEditor
          variants={variantState}
          dispatch={dispatchVariant}
        />

        <button type="submit">Save Changes</button>
      </form>
      <SKUEditor skus={normalizeSKUs(product.skus)} />
    </div>
  )
}
