"use client";

import LoadingItem from '@/components/customer/loading-item'
import { Card, CardContent } from '@/components/ui/card'

const LoadingItems = () => {
  const items = Array(4).fill(1)

  return (
    <Card>
      <CardContent className='space-y-4'>
        {items.map((item, index) => (
          <LoadingItem key={index}/>
        ))}
      </CardContent>
    </Card>
  )
}

export default LoadingItems
