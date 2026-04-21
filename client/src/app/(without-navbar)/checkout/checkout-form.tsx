'use client';

import AddressSection from "./address-section";
import CheckoutItems from "./checkout-items";
import ShippingDropdown from "@/components/customer/shipping-dropdown";
import PaymentMethodList from "@/components/customer/payment-method-list";
import CheckoutTotal from "./checkout-total";
import { useCheckoutForm } from "@/hooks/checkout/useCheckoutForm";
import { useCreateOrder } from "@/hooks/order/useCreateOrder";
import { OrderFormValues } from "@/schemas/order.schema";
import LoadingAddress from "./loading-address";
import LoadingShipping from "./loading-shipping";
import LoadingItems from "./loading-items";
import LoadingPaymentMethodList from "@/components/customer/loading-payment-method";
import LoadingTotal from "./loading-total";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

const CheckoutForm = () => {
  const { form, isLoading, checkout, options } = useCheckoutForm()
  const { mutateAsync: onCreateOrder, isPending } = useCreateOrder()
  const router = useRouter()

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

  const handleCreateOrder = async (form: OrderFormValues) => {
    try {
      const data = await onCreateOrder(form)
      router.push(`/orders/${data.data.id}`)
    } catch (err) {
      toast.error(err as string)
    }
  }

  return (
    <form id="create-order" 
      onSubmit={form.handleSubmit(handleCreateOrder)}
      className="grid gap-x-4 gap-y-2 lg:grid-cols-[9fr_7fr]"
    >
      <div className="space-y-2 md:space-y-4">
        <AddressSection  />
        <ShippingDropdown options={options} form={form} />
        <CheckoutItems checkout={checkout} />
      </div>
      <div className="space-y-2 md:space-y-4">
        <PaymentMethodList/>
        <CheckoutTotal checkout={checkout} isPending={isPending} />
      </div>
    </form>
  )
}

export default CheckoutForm
