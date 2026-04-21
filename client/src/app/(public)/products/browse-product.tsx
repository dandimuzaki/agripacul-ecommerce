'use client'

import LoadingProductCard from "@/components/customer/loading-product-card"
import ProductCard from "@/components/customer/product-card"
import SortProductDropdown from "@/components/customer/sort-product-dropdown"
import { Button } from "@/components/ui/button"
import { useInfiniteProducts } from "@/hooks/product/useInfiniteProducts"
import { useProductFilter } from "@/hooks/product/useProductFilter"

export const BrowseProduct = () => {
  const { filters } = useProductFilter()
  const { data, isLoading, isError, fetchNextPage, hasNextPage } = useInfiniteProducts(filters)
  
  if (isLoading) return <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      {Array(8).fill(1).map((product, index) => (
        <LoadingProductCard key={index} />
      ))}
      </div>

  if (isError) return <p>Error loading products</p>
  if (data?.pages[0].data == null) return <p>No product match your filters</p>
  const products = data?.pages.flatMap(page => page.data) ?? []

  return (
    <>
    <div className="hidden md:flex justify-between items-end gap-6 mb-4">
      <p>Showing <span className="font-semibold">{products.length}</span> results</p>
      <div className='flex items-center gap-4'>
        <SortProductDropdown/>                      
      </div>
    </div>
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-2 md:gap-4">
      {products.length > 0 && products.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
      <Button
        onClick={() => fetchNextPage()}
        disabled={!hasNextPage}
        className="col-span-2 md:col-span-3 lg:col-span-4 w-fit justify-self-center"
        variant="outline"
      >
        Load more
      </Button>
    </div>
    </>
  )
} 