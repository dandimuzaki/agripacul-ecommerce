"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";
import { checkoutService } from "@/services/checkout.service";
import { CheckoutFormValuesTemp } from "@/schemas/checkout.schema";

export const useCheckout = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: CheckoutFormValuesTemp) =>
      checkoutService.previewCheckout(payload),

    onSuccess: (data, payload) => {
      queryClient.setQueryData(
        ["checkout", "preview",
          payload.shipping_address_id,
          payload.selected_shipping_option_id,
          payload.selected_promotion_id,
        ],
        data
      )
    }
  })
}