"use client";

import OrderStatusDropdown from "@/components/customer/order-status-dropdown";
import OrderStatusFilter from "@/components/customer/order-status-filter";
import SortOrderDropdown from "@/components/customer/sort-order-dropdown";

const OrderFilter = () => {
  

  return (
    <>
      <div className="hidden md:block h-fit">
        <p className="font-semibold text-lg">Filter Options</p>
        <form id="order-filter" className="space-y-4">
          <OrderStatusFilter/>
        </form>
      </div>

      <div className="flex justify-between items-center gap-4">
        <div>
        <OrderStatusDropdown/>
        </div>
        <div>
        <SortOrderDropdown/>
        </div>
      </div>
    </>
  )
}

export default OrderFilter
