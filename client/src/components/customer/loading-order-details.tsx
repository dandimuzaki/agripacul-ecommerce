'use client';

import { Card, CardContent } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

const LoadingOrderDetails = () => {
  return (
    <div className='text-sm grid grid-cols-[3fr_2fr] gap-4'>
      <div className="space-y-4">
        <Card>
          <CardContent className='grid grid-cols-3 gap-4'>
            <div className="space-y-1">
              <Skeleton className="h-8 w-16" />
              <Skeleton className="h-8 w-16" />
            </div>
            <div className="space-y-1">
              <Skeleton className="h-8 w-16" />
              <Skeleton className="h-8 w-16" />
            </div>
            <div className="space-y-1">
              <Skeleton className="h-8 w-16" />
              <Skeleton className="h-8 w-24 bg-green-200" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className='space-y-2'>
            <Skeleton className="h-10 w-24" />
            <div className='grid gap-4'>
              {Array(2).fill(1).map((item, index) => (
                <div key={index} className="flex gap-2">
                  <Skeleton
                    className="w-15 h-15 object-cover aspect-square rounded"/>
                  <div className="grid grid-cols-[5fr_2fr] flex-1 gap-2">
                    <Skeleton className="h-8 w-36"/>
                    <Skeleton className="h-8 w-36"/>
                    <div className="flex flex-wrap gap-2">
                      <Skeleton className="h-6 w-16"/>
                      <Skeleton className="h-6 w-16"/>
                      <Skeleton className="h-6 w-16"/>
                    </div>
                    <Skeleton className="h-8 w-24"/>
                  </div>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className='space-y-2'>
            <Skeleton className="h-10 w-24" />
            <div className='grid gap-0'>
              {Array(4).fill(1).map((step, index) => (
                <div key={index} className='flex gap-4'>
                  <div className='flex-1 grid gap-1 my-2'>
                    <Skeleton className="w-36 h-8" />
                    <Skeleton className="w-full h-8" />
                  </div>
                </div>
              ))}
            </div>
            <div className="flex gap-2 items-center justify-center">
              <Skeleton className="h-10 w-24" />
              <Skeleton className="h-10 w-24 bg-primary" />
            </div>
          </CardContent>
        </Card>
      </div>
      <div className="space-y-4">
        <Card>
          <CardContent className='grid grid-cols-2 gap-2'>
            <Skeleton className="h-10 w-36 col-span-2" />
            <Skeleton className="h-8 w-24" />
            <Skeleton className="h-8 w-24 justify-self-end" />
            <Skeleton className="h-8 w-24" />
            <Skeleton className="h-8 w-24 justify-self-end" />
            <Skeleton className="h-8 w-24" />
            <Skeleton className="h-8 w-24 justify-self-end" />
          </CardContent>
        </Card>
         <Card>
          <CardContent className='grid grid-cols-2 gap-2'>
            <Skeleton className="h-10 w-36 col-span-2" />
            <Skeleton className="h-8 w-24" />
            <Skeleton className="h-8 w-24 justify-self-end" />
            <Skeleton className="h-8 w-24" />
            <Skeleton className="h-8 w-24 justify-self-end" />
            <Skeleton className="h-8 w-24" />
            <Skeleton className="h-8 w-24 justify-self-end" />
            <Skeleton className="h-8 w-24 col-span-2" />
            <Skeleton className="h-8 w-full col-span-2" />
          </CardContent>
        </Card>
        <Card>
          <CardContent className='grid grid-cols-2 gap-2'>
            <Skeleton className="h-10 w-24 col-span-2" />
            <Skeleton className="h-8 w-16" />
            <Skeleton className="h-8 w-16 justify-self-end" />
            <Skeleton className="h-8 w-16" />
            <Skeleton className="h-8 w-16 justify-self-end" />
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default LoadingOrderDetails;