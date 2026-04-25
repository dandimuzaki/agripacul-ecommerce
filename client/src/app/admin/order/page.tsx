import { OrderTabs } from './order-tabs';

export default function AdminOrderPage() {
  return (
    <div className='space-y-2'>
      <h2 className='font-bold text-2xl'>Order Management</h2>
      <OrderTabs/>
    </div>
  )
}
