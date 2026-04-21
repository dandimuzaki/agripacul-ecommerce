"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { cartKeys } from "../queries/cartKeys";
import { cartService } from "@/services/cart.service";
import { UpdateCartFormValues } from "@/schemas/cart.schema";

export const useUpdateSelectItem = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({itemId, isSelected}: {itemId: number, isSelected: boolean}) => {
      const payload: UpdateCartFormValues = {
        is_selected: isSelected
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