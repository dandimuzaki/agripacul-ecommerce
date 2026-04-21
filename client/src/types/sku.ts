export type skuStatus = "active" | "inactive" | "archived"

export type skuDetails = {
  id: number,
  product_id: number,
  sku_code: string,
  price: number,
  sale_price?: number,
  stock: number,
  min_stock: number,
  status: skuStatus,
  weight: number,
  variants: variantCombination[],
  images?: string | null,
}

export type variantCombination = {
  name: string,
  value: string
}

export type SKUEdit = {
  id: number
  sku_code: string
  price: number
  sale_price?: number
  stock: number
  min_stock: number
  weight: number
  status: skuStatus

  originalSKUCode: string
  originalPrice: number
  originalSalePrice: number | undefined
  originalStock: number
  originalMinStock: number
  originalWeight: number
  originalStatus: skuStatus

  isEditing: boolean
  variants: variantCombination[]
}
