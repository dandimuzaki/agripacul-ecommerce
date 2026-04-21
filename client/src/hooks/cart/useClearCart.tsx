import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { cartService } from "@/services/cart.service"
import { cartKeys } from '../queries/cartKeys'
import { toast } from 'sonner'
import { Cart } from '@/types/cart'

export const useClearCart = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () =>
      cartService.clearCart(),

    onMutate: async () => {
      await queryClient.cancelQueries({
        queryKey: cartKeys.all,
      })

      const previousItems =
        queryClient.getQueryData(
          cartKeys.all
        )

      // Optimistic update: remove cart instantly
      queryClient.setQueryData<Cart>(
        cartKeys.all,
        (old) => {
          if (!old) return undefined
          return ({...old, items: []})
        }
      )

      return { previousItems }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousItems) {
        queryClient.setQueryData(
          cartKeys.all,
          context.previousItems
        )
      }

      toast.error(_err.message || "Failed to clear cart")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: cartKeys.all,
      })
    },

    onSuccess: () => {
      toast.success("Cart cleared successfully")
    },
  })
}