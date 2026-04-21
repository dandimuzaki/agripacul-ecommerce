'use client';

import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { Cancel, CheckCircle } from "@mui/icons-material";
import AnimatedButton from "@/components/common/animated-button";
import { useOrderDetails } from "@/hooks/order/useOrderDetails";
import { Card, CardContent } from "@/components/ui/card";
import { formatRupiah } from "@/lib/formatCurrency";
import { useProducts } from "@/hooks/product/useProducts";
import RecommendedProducts from "@/components/customer/recommended-products";
import { Button } from "@/components/ui/button";

const AfterCheckout = () => {
  const params = useSearchParams();
  const status = params.get("status");
  const orderId = params.get("id")
  const {data: order} = useOrderDetails(Number(orderId))
  const {data: products, isLoading} = useProducts({
    limit: 10,
    sort_by: "rating",
    sort_order: "desc"
  })

  return (
    <>
    <div className='flex gap-4 justify-center items-center py-8'>
      {status === "success" && (
        <Card className="bg-gray-200 text-center max-w-lg">
          <CardContent className="bg-transparent">
          <div className="text-primary text-[60px]/30 md:text-[90px]/30 lg:text-[120px]/30 mb-2">
            <CheckCircle fontSize="inherit" />
          </div>
          <h2 className="font-semibold text-2xl mb-2">Your order has been placed successfully!</h2>
          <p className="text-gray-500 mb-4 text-base/5">
            Our team will review and confirm your order shortly. You’ll receive a notification once it has been approved.
          </p>
          <Link href="/orders">
            <AnimatedButton>
              Track Order
            </AnimatedButton>
          </Link>
          <div className="mt-6 p-4 rounded bg-white text-sm space-y-2">
            <p className='font-semibold text-base text-left'>Order Summary</p>
            <p className='flex justify-between'>Order ID<span>{order?.id}</span></p>
            <p className='flex justify-between'>Shipping Name<span>{order?.shipping.name}</span></p>
            <p className='flex justify-between'>Estimated Time Delivery<span>{order?.shipping.etd} days</span></p>
            <p className='flex justify-between'>Payment Method<span>{order?.totals.payment_method}</span></p>
            <p className='flex justify-between'>Total Items<span>{order?.items.length}</span></p>
            <p className='flex justify-between'>Subtotal<span>{formatRupiah(order?.totals.subtotal ?? 0)}</span></p>
            <p className='flex justify-between'>Shipping Cost<span>{formatRupiah(order?.totals.shipping_cost ?? 0)}</span></p>
            <p className='flex justify-between'>Total Bill<span>{formatRupiah(order?.totals.grand_total ?? 0)}</span></p>
          </div>
          </CardContent>
        </Card>
      )}

      {status === "error" && (
        <Card className="bg-gray-200 text-center max-w-lg">
          <CardContent className="bg-transparent">
          <div className="text-red-500 text-[120px]/30 mb-2">
            <Cancel fontSize="inherit" />
          </div>
          <h2 className="font-semibold text-2xl mb-2">We couldn’t place your order</h2>
          <p className="text-gray-500 mb-4 text-base/5">
            Something went wrong on our side or your request couldn’t be processed. Please try again or contact support if the issue persists.
          </p>
          <Link href="/">
            <AnimatedButton>
              Back to Home
            </AnimatedButton>
          </Link>
        </CardContent>
        </Card>
      )}
    </div>
    {products && <section className="px-16 py-12 space-y-6">
      <div className="flex justify-between items-center">
        <h2 className="text-4xl text-primary font-semibold">Recommended Products</h2>
        <Button>See More</Button>
      </div>
      <RecommendedProducts products={products} isLoading={isLoading} />
    </section>}
    </>
  );
};

export default AfterCheckout;