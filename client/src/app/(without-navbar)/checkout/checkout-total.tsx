'use client';

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Spinner } from "@/components/ui/spinner";
import { useCheckoutForm } from "@/hooks/checkout/useCheckoutForm";
import { formatRupiah } from "@/lib/formatCurrency";
import { CheckoutPreviewResponse } from "@/types/checkout";

const CheckoutTotal = ({checkout, isPending}: {checkout: CheckoutPreviewResponse, isPending: boolean}) => {
  const total = checkout?.totals

  return (
    <>
    <Card className="h-fit">
      <CardHeader>
        <CardTitle className="text-lg font-bold">Shopping Summary</CardTitle>
      </CardHeader>
      <CardContent className="grid grid-cols-2 gap-y-2">
        <p className="">Subtotal</p>
        <p className="text-right font-medium">{total ? formatRupiah(total.subtotal) : <><Spinner/></>}</p>
        <p className="">Shipping Cost</p>
        <p className="text-right font-medium">{total ? formatRupiah(total.shipping_cost) : <><Spinner/></>}</p>
        <p className="">Discount</p>
        <p className="text-right font-medium">{total ? formatRupiah(total.discount_amount) : <><Spinner/></>}</p>
        <p className="font-bold">Shopping Total</p>
        <p className="text-right font-bold">{total ? formatRupiah(total.grand_total) : <><Spinner/></>}</p>
        <Button className="hidden md:flex col-span-2 font-bold mt-4" type="submit" form="create-order">
          {isPending ? <>
            <Spinner/>
            Processing your order...
          </> : <>Buy Now</>}
        </Button>
      </CardContent>
    </Card>
    <div className="md:hidden shadow-[0px_0px_10px_1px_rgba(0,0,0,0.1)] w-full py-2 px-4 bg-white flex justify-end items-center fixed bottom-0 left-0">
      <Button className="font-bold" type="submit" form="create-order">
        {isPending ? <>
          <Spinner/>
          Processing...
        </> : <>Buy Now</>}
      </Button>
    </div>
    </>
  )
}

export default CheckoutTotal
