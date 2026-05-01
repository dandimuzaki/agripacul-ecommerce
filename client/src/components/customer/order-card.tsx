'use client';

import { formatRupiah } from "@/lib/formatCurrency";
import { formatTime } from "@/lib/formatDate";
import { OrderSummary } from "@/types/order";
import Image from "next/image";
import { Button } from "../ui/button";
import { Card, CardContent } from "../ui/card";
import OrderBadge from "../common/order-badge";
import Link from "next/link";
import ReviewForm from "./review-form";

const OrderCard = ({ order }: {order: OrderSummary}) => {
  return (
    <Card>
      <CardContent className="space-y-2">
        <div className='flex justify-between text-sm items-center'>
          <div>
            <p className='font-bold'>Order placed at</p>
            <p className=''>{formatTime(order.created_at)}</p>
          </div>
          <OrderBadge status={order.display_status}/>
        </div>
        <div className="flex gap-2">
          <div className='row-span-2 h-24 w-24 overflow-hidden rounded-md'>
            {order.first_item?.main_image_url
              ? <Image 
                  src={order.first_item.main_image_url ?? "/loading.png"} 
                  className='aspect-square object-cover'
                  width={100} height={100}
                  alt={order.first_item.name}
                />
              : <div className='w-full h-full bg-gray-200'></div>
            }
          </div>

          <div className="flex-1 flex flex-col justify-between gap-2">
            <div className="flex justify-between gap-4">
              <div className="flex-1">
                <p className='font-semibold'>{order.first_item.name}</p>
                <p className='text-sm'>{order.first_item.quantity} item x <span>{formatRupiah(order.first_item.price)}</span></p>
              </div>
              <div className='text-right w-fit'>
                <p className='text-sm'>Total Bill</p>
                <p className='font-bold'>{formatRupiah(order.grand_total)}</p>
              </div>
            </div>

            <div className="flex justify-between gap-4 items-center">
              <div>
                {order.item_count > 1 && <p className='font-medium'>+{order.item_count - 1} more product</p>}
              </div>
              <div className="flex gap-2 items-center justify-end">
                {order.cancellation && (<Button className="bg-orange-500 hover:bg-orange-600">Cancel</Button>)}
                {order.display_status === "delivered" && (
                  <Button>Mark Complete</Button>
                )}
                <Link href={`/orders/${order.id}`}>
                  <Button variant="outline" className="bg-transparent border-2 shadow-none border-primary text-primary">See Details</Button>
                </Link>
                {(order.display_status === "completed" || order.display_status === "delivered") && order.can_review && (
                  <ReviewForm id={order.id}/>
                )}
              </div>
            </div>
          </div>
        </div>          
      </CardContent>
    </Card>
  );
};

export default OrderCard;