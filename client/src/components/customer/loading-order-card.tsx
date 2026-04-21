'use client';

import { Card, CardContent } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

const LoadingOrderCard = () => {
  return (
    <Card>
      <CardContent className="space-y-2">
        <div className='flex justify-between text-sm items-center'>
          <div className="space-y-2">
            <Skeleton className="h-6 w-24" />
            <Skeleton className="h-6 w-24" />
          </div>
          <Skeleton className="h-8 w-24 bg-green-200" />
        </div>
        <div className="flex gap-2">
          <div className='h-25 aspect-square overflow-hidden rounded-md'>
            <Skeleton className="h-25 w-25" />
          </div>

          <div className="flex-1 flex justify-between flex-col">
            <div className="space-y-2">
              <Skeleton className="h-6 w-36" />
              <Skeleton className="h-6 w-36" />
            </div>
            <Skeleton className="h-6 w-24" />
            
          </div>
          <div className='flex gap-5 justify-end items-end flex-col'>
            <div className='text-right space-y-2'>
              <Skeleton className="h-6 w-24" />
              <Skeleton className="h-6 w-24" />
            </div>
            <div className="flex gap-2 items-center justify-end">
              <Skeleton className="h-8 w-24" />
              <Skeleton className="h-8 w-24 bg-primary" />
            </div>
          </div>
        </div>          
      </CardContent>
    </Card>
  );
};

export default LoadingOrderCard;