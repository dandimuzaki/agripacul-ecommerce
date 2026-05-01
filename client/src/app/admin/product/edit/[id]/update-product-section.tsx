'use client';

import { useUpdateProduct } from "@/hooks/product/useUpdateProduct";
import { ProductFormValues } from "@/schemas/product.schema";
import EditProductForm from "./components/product-editor";
import { VariantEditor } from "./components/variant-editor";
import { SKUEditor } from "./components/sku-editor";
import { useProductId } from "@/hooks/product/useProductId";
import MainImageUploader from "./components/main_image";
import ProductGalleryUploader from "./components/gallery";

const UpdateProductSection = ({id}: {id:number}) => {
  const { data: product } = useProductId(id)
  const { mutate, isPending } = useUpdateProduct()
  
  const onUpdateProduct = (data: ProductFormValues) => {
    console.log(data)
    mutate({
      id,
      payload: data
    })
  }

  if (!product) return

  return (
    <section className="space-y-8">
      <EditProductForm product={product} onUpdateProduct={onUpdateProduct} isPending={isPending} />
      <div className="flex items-start gap-6">
        <MainImageUploader id={id} image={product.main_image_url}/>
        <ProductGalleryUploader id={id} images={product.images}/>
      </div>
      <VariantEditor product={product} onUpdateProduct={onUpdateProduct} isPending={isPending}/>
      <SKUEditor productId={product.id}/>
    </section>
  )
}

export default UpdateProductSection
