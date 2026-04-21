'use client';

import { Card, CardContent } from '../ui/card';
import { Skeleton } from '../ui/skeleton';

const LoadingProductCard = () => {
  return (
    <div>
        <div className='flex flex-col gap-2'>
          <div className='aspect-square bg-gray-100 mb-2 '>
            <Skeleton className='bg-gray-200 h-full w-full object-cover' />
          </div>
          <div className='flex justify-between items-center'>
            <Skeleton className='bg-gray-200 h-6 w-full' />
          </div>
          <Skeleton className='bg-gray-200 h-6 w-full' />
          <div className='mt-2 flex justify-between items-center mt-auto'>
            <div>
              <Skeleton className='bg-gray-200 h-8 w-24' />
            </div>
            <button className='bg-primary rounded-lg text-white p-2 flex items-center gap-1'>
              <Skeleton className='bg-gray-200 h-8 w-8 bg-primary' />
            </button>
          </div>
        </div>
    </div>
  )
}

export default LoadingProductCard
