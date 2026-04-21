"use client";

import { Skeleton } from "../ui/skeleton";

const LoadingItem = () => {
  return (
    <div className="flex gap-2">
      <Skeleton
        className="w-16 h-16 object-cover aspect-square rounded"/>
      <div className="grid grid-cols-[5fr_2fr] flex-1 gap-2">
        <Skeleton className="w-36 h-8" />
        <Skeleton className="w-36 h-8" />
        <div className="flex flex-wrap gap-2">
          <Skeleton className="w-16 h-6" />
          <Skeleton className="w-16 h-6" />
          <Skeleton className="w-16 h-6" />
        </div>
        <Skeleton className="w-36 h-8" />
      </div>
    </div>
  )
}

export default LoadingItem
