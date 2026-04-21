import { ShippingFormValues } from "@/schemas/checkout.schema"
import { checkoutService } from "@/services/checkout.service"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { checkoutKeys } from "../queries/checkoutKeys"

export const useCheckout = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: ShippingFormValues) =>
      checkoutService.getShippingOptions(payload),

    onSuccess: (data, payload) => {
      queryClient.setQueryData(
        checkoutKeys.shippings(payload.shipping_address_id as number),
        data.data
      )
    }
  })
}