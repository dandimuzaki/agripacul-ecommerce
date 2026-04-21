'use client';

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

const LoadingTotal = () => {
  return (
    <Card className="h-fit">
      <CardHeader>
        <CardTitle className="text-lg font-bold">
          <Skeleton className='h-10 w-24' />
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        {Array(4).fill(1).map((i, index) => <div className="grid grid-cols-2" key={index}>
          <Skeleton className='h-8 w-16' />
          <Skeleton className='h-8 w-16 justify-self-end' />
        </div>)}
        <Skeleton className='h-10 col-span-2 bg-primary' />
      </CardContent>
    </Card>
  )
}

export default LoadingTotal
