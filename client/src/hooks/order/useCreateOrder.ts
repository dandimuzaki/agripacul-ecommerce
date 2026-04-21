"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { useRouter } from "next/navigation"
import { orderKeys } from "../queries/orderKeys";
import { orderService } from "@/services/order.service";
import { OrderFormValues } from "@/schemas/order.schema";

export const useCreateOrder = () => {
  const queryClient = useQueryClient()
  const router = useRouter()

  return useMutation({
    mutationFn: (payload: OrderFormValues) =>
      orderService.createOrder(payload),

    onSuccess: () => {
      // refresh order list cache
      queryClient.invalidateQueries({
        queryKey: orderKeys.adminLists()
      })

      queryClient.invalidateQueries({
        queryKey: orderKeys.customerLists()
      })

      // redirect to order list
      router.push("/purchase?status=success")
    }
  })
}