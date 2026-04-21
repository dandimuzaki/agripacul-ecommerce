import { EditProductFormValues } from "@/schemas/product.schema"

export type RowStatus = "clean" | "created" | "updated" | "deleted"

export type VariantTypeEdit = {
  id?: number
  tempId: string
  originalName?: string
  name: string
  values: VariantValueEdit[]
  status: RowStatus
  isEditing?: boolean
}

export type VariantValueEdit = {
  id?: number
  tempId: string
  originalValue?: string
  value: string
  status: RowStatus
  isEditing?: boolean
}

export type VariantType = {
  id?: number
  name: string
  values: VariantValue[]
}

export type VariantValue = {
  id?: number
  value: string
}

export type EditVariantType = {
  id?: number
  name: string
  values: VariantValue[]
  status?: string
}

export type EditVariantValue = {
  id?: number
  value: string
  status?: string
}

export type EditVariant = EditProductFormValues & {
  variants: VariantTypeEdit[]
}