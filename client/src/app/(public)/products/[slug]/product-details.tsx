'use client';
import { RatingStars } from '@/components/customer/rating-star';
import { formatRupiah } from '@/lib/formatCurrency';
import { useMemo, useState } from 'react';
import VariantSelection from './variant-selection';
import { Button } from '@/components/ui/button';
import { Icon } from '@/components/ui/icon';
import { useProductSlug } from '@/hooks/product/useProductSlug';
import { useAddToCart } from '@/hooks/cart/useAddToCart';
import ProductGallery from './product-gallery';
import RecommendedProducts from '@/components/customer/recommended-products';
import { useProducts } from '@/hooks/product/useProducts';
import ReviewSection from './review-section';
import Link from 'next/link';
import { ArrowBackIos } from '@mui/icons-material';

const ProductDetails = ({slug}: {slug: string}) => {
  const { data: product, isLoading } = useProductSlug(slug)
  const {data: products, isLoading: loadProduct} = useProducts({
      limit: 10,
      category_id: product?.category.id
    })

  const defaultVariants = useMemo(() => {
    if (!product) return {};

    const defaultSKU = product?.skus?.find(
      (sku) => sku.id === product.default_sku_id
    );

    if (!defaultSKU) return {};

    const initial: Record<number, number> = {};

    if (product.variants != null && product.variants.length > 0) {
      product.variants.forEach((variant) => {
        const match = defaultSKU.variant_value_ids.find((id) =>
          (variant.values || []).some((v) => v.id === id)
        );

        if (match) initial[variant.id] = match;
      });
    }

    return initial;
  }, [product]);

  const [selectedVariants, setSelectedVariants] = useState<
    Record<number, number>
  >(defaultVariants);

  const { mutate: addToCart, isPending } = useAddToCart()
      
  const onAddToCart = (skuId: number) => {
    addToCart(skuId)
  }

  if (isLoading) return <p>Loading...</p>;
  if (!product) return <p>Product not found</p>;

  const selectedVariantValueIDs = Object.values(selectedVariants);
  const selectedSKU = product?.skus?.find((sku) =>
    selectedVariantValueIDs.every((id) =>
      sku.variant_value_ids.includes(id)
    )
  );

  return (
    <>
    <section className='min-h-screen px-4 pt-16 pb-4 md:pt-32 md:pb-16 md:px-16 space-y-2 md:space-y-4'>
      <Link href={"/products"} className="flex items-center"><ArrowBackIos fontSize="inherit"/>Back</Link>
      <div className='md:flex items-start md:gap-6 space-y-4'>
        <div className=''>
          <ProductGallery/>
        </div>
        <div className='flex-1 w-full h-full'>
          <div className='md:space-y-1'>
            <p className='text-sm md:text-base text-gray-500 font-medium'>{product.category.name}</p>
            <h2 className='text-lg md:text-xl font-bold'>{product.name}</h2>
            <div className='flex gap-1 items-center'>
              <RatingStars rating={product.average_rating} />
              {product.average_rating}
              {` (${product.review_count} Reviews)`}
            </div>
          </div>
          <div className="my-2 md:my-3">
            {selectedSKU ? (
              <p className="text-xl md:text-2xl font-semibold text-primary">
                {formatRupiah(selectedSKU.sale_price ?? selectedSKU.price)}
              </p>
            ) : (
              <div className='flex gap-2 items-center my-2'>
                <p className='text-xl text-primary font-medium'>{formatRupiah(product.min_price)}</p>
                <p className='text-lg text-gray-500 primary line-through'>{formatRupiah(product.max_price)}</p>
              </div>
            )}
          </div>
          <p>{product.description}</p>
          <VariantSelection
            product={product}
            selectedVariants={selectedVariants}
            setSelectedVariants={setSelectedVariants}
          />
          
          <div>
            <Button
              disabled={!selectedSKU || selectedSKU.stock === 0}
              className="mt-4 px-6 py-3 rounded-lg bg-primary text-white disabled:opacity-50"
              onClick={() => {
                if (!selectedSKU) return;

                onAddToCart(selectedSKU.id)
              }}
            >
              {selectedSKU?.stock === 0
                ? ("Empty Stock")
                : (<>
                  <Icon icon='mui:add-cart'/>
                  {"Add to Cart"}
                  </>)}
            </Button>
          </div>
        </div>
      </div>
    </section>
    <ReviewSection productId={product.id}/>
    {products && <section className="p-8 md:p-16 space-y-2 md:space-y-4 bg-white">
      <div className="md:flex justify-between items-center">
        <h2 className="text-xl md:text-2xl text-primary font-semibold uppercase">Recommended Products</h2>
        <Button className='hidden md:flex'>See More</Button>
      </div>
      <RecommendedProducts products={products} isLoading={loadProduct} />
    </section>}
    </>
  )
}

export default ProductDetails
