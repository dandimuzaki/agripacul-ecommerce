import { Skeleton } from "@/components/ui/skeleton";
import LoadingCartTable from "./loading-cart-table";
import LoadingCartSummary from "./loading-cart-summary";

export default function LoadingCart() {
  return (
    <section className='grid gap-x-4 gap-y-2 grid-cols-[7fr_3fr] px-8 py-8 pt-24'>
      <Skeleton className="text-xl w-10"/>
      <LoadingCartTable/>
      <LoadingCartSummary/>
    </section>
  )
}