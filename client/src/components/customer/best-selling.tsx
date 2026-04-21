"use client";
import ProductCard from './product-card'
import { useProducts } from '@/hooks/product/useProducts'
import LoadingProductCard from './loading-product-card'

const BestSelling = () => {
  const {data: products, isLoading} = useProducts({
    limit: 10,
    sort_by: "sold",
    sort_order: "desc"
  })

  if (isLoading) return <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-2 md:gap-4">
    {Array(10).fill(1).map((product, index) => (
      <LoadingProductCard key={index} />
    ))}
    </div>

  return (
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-2 md:gap-4">
      {products && products?.length > 0 && products?.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  )
}

export default BestSelling
