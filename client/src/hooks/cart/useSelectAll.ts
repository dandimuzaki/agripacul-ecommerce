"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { cartKeys } from "../queries/cartKeys";
import { cartService } from "@/services/cart.service";
import { Cart, Item } from "@/types/cart";

export const useSelectAll = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (isSelected: boolean) => {
      return cartService.batchSelectItem({is_selected: isSelected})
    },

    onMutate: async (isSelected: boolean) => {
      await queryClient.cancelQueries({ queryKey: cartKeys.all });

      const previousCart = queryClient.getQueryData(cartKeys.all);

      queryClient.setQueryData(cartKeys.all, (old: Cart) => {
        if (!old) return old;

        return {
          ...old,
          items: old.items.map((item: Item) => ({
            ...item,
            is_selected: isSelected,
          })),
        };
      });

      return { previousCart };
    },

    onError: (_err, _vars, context) => {
      queryClient.setQueryData(cartKeys.all, context?.previousCart);
    },

    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: cartKeys.all });
    },
  })
}