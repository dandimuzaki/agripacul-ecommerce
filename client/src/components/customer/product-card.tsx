'use client';

import { ProductSummary } from '@/types/product'
import { Icon } from '../ui/icon';
import Image from 'next/image';
import Link from 'next/link';
import { formatRupiah } from '@/lib/formatCurrency';

const ProductCard = ({product}: {product: ProductSummary}) => {
  return (
    <div className='h-full flex justify-between gap-1 flex-col'>
        <Link href={`/products/${product.slug}`} className='flex flex-col'>
            <div className='aspect-square bg-gray-200 mb-2 '>
              <img src={product.main_image_url ?? "/loading.png"} alt={product.name} height={100} width={100} className='h-full w-full object-cover rounded-lg' />
            </div>
            <div className='flex justify-between items-center'>
              <p className='text-xs text-gray-500'>{product.category.name}</p>
              <div className='flex items-center gap-1'>
                <Icon icon='lucide:star' size={16} fill='yellow' className='text-yellow-200' />
                <span className='font-medium'>{product.average_rating}</span>
              </div>
            </div>
            <p className='font-medium text-sm md:text-base'>{product.name}</p>
        </Link>
          <div className='flex justify-between items-center'>
            <div>
              <p className='font-bold text-base md:text-lg'>{formatRupiah(product.min_price)}</p>
            </div>
            <Link href={`/products/${product.slug}`} className='bg-primary rounded-lg text-white p-1 md:p-2 flex items-center gap-1'>
              <Icon icon='mui:add-cart' fontSize="inherit"/>
            </Link>
          </div>
    </div>
  )
}

export default ProductCard
