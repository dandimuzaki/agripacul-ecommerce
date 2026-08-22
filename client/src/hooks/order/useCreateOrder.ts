"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { orderKeys } from "../queries/orderKeys";
import { orderService } from "@/services/order.service";
import { OrderFormValues } from "@/schemas/order.schema";

export const useCreateOrder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: OrderFormValues) => {
      return orderService.createOrder(payload)
    },

    onSuccess: () => {
      // refresh order list cache
      queryClient.invalidateQueries({
        queryKey: orderKeys.adminLists()
      })

      queryClient.invalidateQueries({
        queryKey: orderKeys.customerLists()
      })
    }
  })
}