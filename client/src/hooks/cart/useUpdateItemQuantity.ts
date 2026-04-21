"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { cartKeys } from "../queries/cartKeys";
import { cartService } from "@/services/cart.service";
import { UpdateCartFormValues } from "@/schemas/cart.schema";

export const useUpdateItemQuantity = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({itemId, quantity}: {itemId: number, quantity: number}) => {
      const payload: UpdateCartFormValues = {
        quantity: quantity
      }
      return cartService.updateItemQuantity(itemId, payload)
    },

    onSuccess: () => {
      // refresh cart list cache
      queryClient.invalidateQueries({
        queryKey: cartKeys.all
      })
    }
  })
}