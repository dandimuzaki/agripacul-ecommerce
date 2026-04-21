import { useQuery } from "@tanstack/react-query"
import { checkoutKeys } from "../queries/checkoutKeys"
import { checkoutService } from "@/services/checkout.service"
import { Promotion } from "@/types/checkout"
import { PromotionFormValues } from "@/schemas/checkout.schema";

export const useQueryPromotion = (payload: PromotionFormValues) => {
  return useQuery<Promotion[]>({
    queryKey: checkoutKeys.promotions(payload),
    queryFn: () => checkoutService.getValidPromotions(payload),
  });
};