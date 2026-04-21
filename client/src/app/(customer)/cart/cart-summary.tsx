'use client';

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useCart } from "@/hooks/cart/useCart";
import { formatRupiah } from "@/lib/formatCurrency";
import { useRouter } from "next/navigation";
import LoadingCartSummary from "./loading-cart-summary";
import { useCheckoutForm } from "@/hooks/checkout/useCheckoutForm";
import { Spinner } from "@/components/ui/spinner";
import { Checkbox } from "@/components/ui/checkbox";
import { useSelectAll } from "@/hooks/cart/useSelectAll";

const CartSummary = () => {
  const router = useRouter()
  const {data: cart, isLoading } = useCart()
  const { isLoading: checkoutLoad } = useCheckoutForm()
  const handlePreview = () => {
    router.push("/checkout")
  }

  const isAllSelected = cart && cart.items.length > 0 && cart.items.every(item => item.is_selected);

  const { mutate: selectAll } = useSelectAll()
        
  const onSelectAll = (isSelected: boolean) => {
    selectAll(isSelected)
  }

  if (isLoading) return (
    <LoadingCartSummary/>
  )

  if (!cart) return

  return (
    <>
    <Card className="h-fit hidden md:flex">
      <CardHeader>
        <CardTitle className="text-lg font-bold">Summary</CardTitle>
      </CardHeader>
      <CardContent className="grid grid-cols-2 gap-y-4">
        <p className="">Total Price</p>
        <p className="text-right font-medium">{formatRupiah(cart?.summary.total_selected_price as number)}</p>
        <Button className="col-span-2 font-bold" onClick={() => handlePreview()}>
          {checkoutLoad ? <><Spinner/>Taking you to checkout...</> : <>Checkout Now</>}
        </Button>
      </CardContent>
    </Card>
    <div className="md:hidden shadow-[0px_0px_10px_1px_rgba(0,0,0,0.1)] w-full py-2 px-4 bg-white flex justify-between items-center fixed bottom-0 left-0">
      <div className="flex gap-2 items-center">
      <Checkbox
        className="w-6 h-6"
        checked={isAllSelected}
        onCheckedChange={() => onSelectAll(!isAllSelected)}
      />
      Select All
      </div>
      <div className="flex items-center gap-2">
        <p className="font-bold">{formatRupiah(cart.summary.total_selected_price)}</p>
        <Button className="w-fit" onClick={() => handlePreview()}>
          {checkoutLoad ? <><Spinner/>Taking you to checkout...</> : <>Checkout ({cart.summary.total_selected_items})</>}
        </Button>
      </div>
    </div>
    </>
  )
}

export default CartSummary
