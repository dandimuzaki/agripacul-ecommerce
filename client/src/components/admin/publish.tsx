"use client"

import { usePublishProduct } from "@/hooks/product/usePublishProduct"
import { Label } from "../ui/label"
import { Switch } from "../ui/switch"
import { ProductSummary } from "@/types/product"

export default function PublishSwitch({product}: {product: ProductSummary}) {
  const {mutate} = usePublishProduct()
  const handlePublishProduct = () => {
    mutate({id: product?.id, isPublished: !product?.is_published})
  }
  return (
    <div className="flex items-center space-x-2">
      <Switch onCheckedChange={handlePublishProduct} checked={product.is_published} id="publish" />
      
    </div>
  )
}