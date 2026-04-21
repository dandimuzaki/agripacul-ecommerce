import { useQuery } from "@tanstack/react-query"
import { checkoutKeys } from "../queries/checkoutKeys"
import { checkoutService } from "@/services/checkout.service"
import { ShippingFormValues } from "@/schemas/checkout.schema";
import { Response } from "@/types/response";
import { ShippingOption } from "@/types/checkout";

export const useQueryShipping = (payload: ShippingFormValues) => {
  return useQuery<Response, Error, ShippingOption[]>({
    queryKey: checkoutKeys.shippings(payload.shipping_address_id as number),
    queryFn: () => checkoutService.getShippingOptions(payload),
    select: (res) => res.data,
    enabled: !!payload.shipping_address_id,
  });
};