'use client';

import { ReusablePagination } from "@/components/common/pagination";
import LoadingOrderCard from "@/components/customer/loading-order-card";
import OrderCard from "@/components/customer/order-card";
import { useOrderFilter } from "@/hooks/order/useOrderFilter";
import { useOrderHistory } from "@/hooks/order/useOrderHistory";
import { OrderSummary } from "@/types/order";

const OrderHistory = () => {
  const {filters} = useOrderFilter()
  const { data, isLoading, isError } = useOrderHistory(filters)

  if (isError) return <p>Error loading orders</p>

  const orders = data?.data
  const pagination = data?.pagination

  if (!pagination) return <p>No orders match your filters</p>

  return (
    <div className="space-y-2 md:space-y-4">
      {isLoading ? Array(4).fill(1).map((o, index) => (
        <LoadingOrderCard key={index}/>
      )) : orders?.map((o: OrderSummary) => (
        <OrderCard key={o.id} order={o}/>
      ))}
      <ReusablePagination currentPage={pagination?.page} totalPages={pagination?.total_pages}/>
    </div>
  )
}

export default OrderHistory
