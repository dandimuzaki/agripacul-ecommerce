'use client';

import OrderItem from "@/components/customer/order-item";
import { formatTime } from "@/lib/formatDate";
import { capitalize } from "@mui/material";
import { formatRupiah } from "@/lib/formatCurrency";
import OrderStatus from "@/components/common/order-status";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import OrderBadge from "@/components/common/order-badge";
import { useOrderById } from "@/hooks/order/useOrderById";
import ConfirmOrderById from "@/components/admin/confirm-order-by-id";

const OrderDetails = ({id}: {id: number}) => {
  const { data: order } = useOrderById(id)

  if (!order) return

  return (
    <div className='text-sm grid grid-cols-[3fr_2fr] gap-4'>
      <div className="space-y-4">
        <Card>
          <CardContent className='grid grid-cols-3 gap-4'>
            <div className="space-y-1">
              <p className="font-bold">Order ID</p>
              <p className="font-medium text-gray-500">{order.id}</p>
            </div>
            <div className="space-y-1">
              <p className="font-bold">Purchasing Date</p>
              <p className="font-medium text-gray-500">{formatTime(order.created_at)}</p>
            </div>
            <div className="space-y-1">
              <p className="font-bold">Status</p>
              <OrderBadge status={order.display_status}/>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className='space-y-2'>
            <p className='text-lg font-bold'>Product Detail</p>
            <div className='grid gap-4'>
              {order.items.map((item, index) => (
                <OrderItem key={index} item={item}/>
              ))}
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className='space-y-2'>
            <p className='font-bold text-lg'>Order Tracking</p>
            <div className='grid gap-0'>
              {order.steps.map((step, index) => (
                <OrderStatus
                  key={index}
                  step={step}
                />
              ))}
            </div>
            <div className="flex gap-2 items-center justify-center">
              {order.cancellation && (<Button className="bg-orange-500 hover:bg-orange-600">Cancel</Button>)}
              {order.steps.map((o, index) => (o.key === "confirmed" && !o.done && (
                <ConfirmOrderById key={index} id={order?.id}/>
              )))}
            </div>
          </CardContent>
        </Card>
      </div>
      <div className="space-y-4">
        <Card>
          <CardContent className='grid grid-cols-2 gap-2'>
            <h3 className="font-bold text-lg col-span-2">
              Customer Information
            </h3>
            <p className="font-medium">Name</p>
            <p className="text-right">{order.customer.name}</p>
            <p className="font-medium">Email</p>
            <p className="text-right">{order.customer.email}</p>
            <p className="font-medium">Phone Number</p>
            <p className="text-right">{order.billing_address.phone_number}</p>
          </CardContent>
        </Card>
         <Card>
          <CardContent className='grid grid-cols-2 gap-2'>
            <p className='font-bold text-lg col-span-2'>Shipping Information</p>
            <p className="font-medium">Tracking Number</p>
            <p className="text-right">{order?.tracking_number ?? "-"}</p>
            <p className="font-medium">Courier</p>
            <p className="text-right">{order?.shipping.name}</p>
            <p className="font-medium">Estimated Time Delivery</p>
            <p className="text-right">{order?.shipping.etd} days</p>
            <p className="col-span-2">Address</p>
            <p className="col-span-2">{order.billing_address.detail_address}, {capitalize(order.billing_address.subdistrict)}, {capitalize(order.billing_address.district)}, {capitalize(order.billing_address.regency)}, {capitalize(order.billing_address.province)}, {order.billing_address.phone_number}</p>
        </CardContent>
        </Card>
        <Card>
          <CardContent className='space-y-2'>
            <p className='font-bold text-lg'>Payment Detail</p>
            <p className='flex justify-between'>Payment Method<span>{order.totals.payment_method}</span></p>
            <p className='flex justify-between'>Total Item Price<span>{formatRupiah(order.totals.subtotal)}</span></p>
            <p className='flex justify-between'>Shipping Cost<span>{formatRupiah(order.totals.shipping_cost)}</span></p>

            <p className='flex justify-between'>Total Bill<span>{formatRupiah(order.totals.grand_total)}</span></p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default OrderDetails;