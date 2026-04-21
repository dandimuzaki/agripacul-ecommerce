"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { cartKeys } from "../queries/cartKeys";
import { cartService } from "@/services/cart.service";
import { CartFormValues } from "@/schemas/cart.schema";

export const useAddToCart = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (skuID: number) => {
      const payload: CartFormValues = {
        sku_id: skuID,
        quantity: 1
      }
      return cartService.addToCart(payload)
    },

    onSuccess: () => {
      // refresh cart list cache
      queryClient.invalidateQueries({
        queryKey: cartKeys.all
      })
    }
  })
}