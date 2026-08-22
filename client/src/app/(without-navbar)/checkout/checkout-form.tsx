'use client';

import AddressSection from "./address-section";
import CheckoutItems from "./checkout-items";
import ShippingDropdown from "@/components/customer/shipping-dropdown";
import PaymentMethodList from "@/components/customer/payment-method-list";
import CheckoutTotal from "./checkout-total";
import { useCheckoutForm } from "@/hooks/checkout/useCheckoutForm";
import LoadingAddress from "./loading-address";
import LoadingShipping from "./loading-shipping";
import LoadingItems from "./loading-items";
import LoadingPaymentMethodList from "@/components/customer/loading-payment-method";
import LoadingTotal from "./loading-total";
import { useState } from "react";
import AddressList from "@/components/customer/address-list";
import { useCreateOrder } from "@/hooks/order/useCreateOrder";
import { useRouter } from "next/navigation";
import { OrderFormValues } from "@/schemas/order.schema";
import { toast } from "sonner";

const CheckoutForm = () => {
  const { form, isLoading, checkout, shippingAddressId, options } = useCheckoutForm()
  const [openAddress, setOpenAddress] = useState<boolean>(false)
  const { mutateAsync: onCreateOrder, isPending } = useCreateOrder()
  const router = useRouter()

  const handleCreateOrder = async (form: OrderFormValues) => {
    try {
      const data = await onCreateOrder(form)
      console.log("data order created", data)
      router.push(`/purchase?status=success&id=${data.data}`)
    } catch (err) {
      toast.error(String(err))
      router.push(`/purchase?status=error`)
    }
  }

  if (isLoading) return (
    <form id="create-order" 
      className="grid gap-x-4 gap-y-2 lg:grid-cols-[9fr_7fr]"
    >
      <div className="space-y-2 md:space-y-4">
        <LoadingAddress/>
        <LoadingShipping/>
        <LoadingItems/>
      </div>
      <div className="space-y-2 md:space-y-4">
        <LoadingPaymentMethodList/>
        <LoadingTotal/>
      </div>
    </form>
  )

  if (!checkout) return

  return (
    <>
      <form id="create-order" 
        className="grid gap-x-4 gap-y-2 lg:grid-cols-[9fr_7fr]"
        onSubmit={form.handleSubmit(handleCreateOrder)}
      >
        <div className="space-y-2 md:space-y-4">
          <AddressSection shippingAddressId={shippingAddressId} setOpenAddress={setOpenAddress} />
          <ShippingDropdown options={options} form={form} />
          <CheckoutItems checkout={checkout} />
        </div>
        <div className="space-y-2 md:space-y-4">
          <PaymentMethodList form={form} />
          <CheckoutTotal checkout={checkout} isPending={isPending} />
        </div>
      </form>
      <AddressList openAddress={openAddress} setOpenAddress={setOpenAddress} form={form} shippingAddressId={shippingAddressId}/>
    </>
  )
}

export default CheckoutForm
