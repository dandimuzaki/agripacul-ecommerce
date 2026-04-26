import OrderPeriod from "@/components/common/order-period";
import OrderFilter from "./order-filter";
import OrderHistory from "./order-history";
import SortOrderDropdown from "@/components/customer/sort-order-dropdown";
import { Suspense } from "react";
import { Skeleton } from "@/components/ui/skeleton";

export default function OrderListPage() {
  return (
    <section className='grid md:grid-cols-[1fr_3fr] gap-y-2 gap-x-8 md:px-8 md:py-8 md:pt-24 p-4 pt-16'>
      <Suspense fallback={<Skeleton className="h-full w-full" />}>
        <OrderFilter/>
      </Suspense>
      <div className="">
        <h2 className="uppercase text-primary text-xl md:text-2xl font-semibold mb-2">Order History</h2>
        <div className="hidden md:flex justify-between items-center gap-6 mb-4">
          <Suspense fallback={<Skeleton className="h-8 w-16" />}>
            <OrderPeriod/>
          </Suspense>
          <div className='flex items-center gap-4'>
            <Suspense fallback={<Skeleton className="h-8 w-16" />}>
              <SortOrderDropdown/>                      
            </Suspense>
          </div>
        </div>
        <OrderHistory/>
      </div>
    </section>
  );
}
