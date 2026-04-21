import CheckoutItem from '@/components/customer/checkout-item'
import { Card, CardContent } from '@/components/ui/card'
import { CheckoutPreviewResponse } from '@/types/checkout'

const CheckoutItems = ({checkout}: {checkout: CheckoutPreviewResponse}) => {
  const items = checkout.selected_items

  return (
    <Card>
      <CardContent className='space-y-4'>
        {items && items.map((item, index) => (
          <CheckoutItem item={item} key={index}/>
        ))}
      </CardContent>
    </Card>
  )
}

export default CheckoutItems
