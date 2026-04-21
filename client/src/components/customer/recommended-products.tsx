"use client";
import ProductCard from './product-card'
import LoadingProductCard from './loading-product-card'
import { ProductSummary } from '@/types/product';

const RecommendedProducts = ({products, isLoading}: {products: ProductSummary[], isLoading: boolean}) => {
  if (isLoading) return <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
    {Array(10).fill(1).map((product, index) => (
      <LoadingProductCard key={index} />
    ))}
    </div>

  return (
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
      {products && products?.length > 0 && products?.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  )
}

export default RecommendedProducts
