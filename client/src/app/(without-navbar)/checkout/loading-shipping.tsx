'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';

const LoadingShipping = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle className='text-primary font-medium text-lg uppercase'>
          <Skeleton className='h-10 w-36'/>
        </CardTitle>
      </CardHeader>
      <CardContent>
    <Skeleton className='h-20 w-full'/>
    </CardContent>
    </Card>
  );
};

export default LoadingShipping;