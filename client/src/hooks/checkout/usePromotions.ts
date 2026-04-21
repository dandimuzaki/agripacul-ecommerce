import { PromotionFormValues } from "@/schemas/checkout.schema"
import { checkoutService } from "@/services/checkout.service"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { checkoutKeys } from "../queries/checkoutKeys"

export const useCheckout = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: PromotionFormValues) =>
      checkoutService.getValidPromotions(payload),

    onSuccess: (data, payload) => {
      queryClient.setQueryData(
        checkoutKeys.promotions(payload),
        data
      )
    }
  })
}