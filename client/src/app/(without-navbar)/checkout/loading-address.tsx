'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton';

const LoadingAddress = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle className='text-primary font-medium text-lg uppercase'>
          <Skeleton className='h-10 w-36'/>
        </CardTitle>
      </CardHeader>
      <CardContent className='flex gap-4'>
      <div className='space-y-2 flex-1'>
         <Skeleton className='h-8 w-full'/>
         <Skeleton className='h-8 w-full'/>
      </div>
      <div className='flex justify-center items-center'>
        <Skeleton className='h-7 w-16 bg-primary'/>
      </div>
      </CardContent>
    </Card>
  )
}

export default LoadingAddress
