import OrderPeriod from "@/components/common/order-period";
import OrderFilter from "./order-filter";
import OrderHistory from "./order-history";
import SortOrderDropdown from "@/components/customer/sort-order-dropdown";

export default function OrderListPage() {
  return (
    <section className='grid md:grid-cols-[1fr_3fr] gap-y-2 gap-x-8 md:px-8 md:py-8 md:pt-24 p-4 pt-16'>
      <OrderFilter/>
      <div className="">
        <h2 className="uppercase text-primary text-xl md:text-2xl font-semibold mb-2">Order History</h2>
        <div className="hidden md:flex justify-between items-center gap-6 mb-4">
          <OrderPeriod/>
          <div className='flex items-center gap-4'>
            <SortOrderDropdown/>                      
          </div>
        </div>
        <OrderHistory/>
      </div>
    </section>
  );
}
