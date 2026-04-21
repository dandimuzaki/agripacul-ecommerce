"use client";

import { useQuery } from "@tanstack/react-query"
import { paymentService } from "@/services/payment.service"
import { paymentMethodKeys } from "../queries/paymentKeys";
import { Response } from "@/types/response";
import { PaymentType } from "@/types/payment";

export const usePaymentMethods = () => {  
  return useQuery<Response, Error, PaymentType[]>({
    queryKey: paymentMethodKeys.all,
    queryFn: () => paymentService.getPaymentMethodList(),
    select: (res) => res.data
  })
}