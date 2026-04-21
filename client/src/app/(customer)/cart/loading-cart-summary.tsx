'use client';

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

const LoadingCartSummary = () => {
  return (
    <Card className="h-fit">
      <CardHeader>
        <CardTitle className="text-lg font-bold">
          <Skeleton className="w-24 h-10"/>
        </CardTitle>
      </CardHeader>
      <CardContent className="grid grid-cols-2 gap-y-4">
        <Skeleton className="w-24 h-8" />
        <Skeleton className="w-24 h-8 justify-self-end" />
        <Skeleton className="col-span-2 h-10 bg-primary">
          
        </Skeleton>
      </CardContent>
    </Card>
  )
}

export default LoadingCartSummary
