import { useQuery } from "@tanstack/react-query"
import { checkoutService } from "@/services/checkout.service"
import { CheckoutPreviewResponse, ShippingOption } from "@/types/checkout"
import { CheckoutFormValuesTemp } from "@/schemas/checkout.schema";
import { Response } from "@/types/response";
import { useCart } from "../cart/useCart";

export const useCheckoutPreview = (payload: CheckoutFormValuesTemp, selectedShipping?: ShippingOption) => {
  const {data: cart} = useCart()
  
  return useQuery<Response, Error, CheckoutPreviewResponse>({
    queryKey: ["checkout", "preview",
      payload.shipping_address_id,
      payload.selected_shipping_option_id,
      payload.selected_promotion_id,
    ],
    queryFn: () => checkoutService.previewCheckout({...payload, selected_shipping_option: selectedShipping}),
    select: (res) => res.data,
    placeholderData: (prev) => prev,
    enabled: cart?.summary.total_selected_items != 0
  });
};