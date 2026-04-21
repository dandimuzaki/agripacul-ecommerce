import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { orderService } from "@/services/order.service"
import { orderKeys } from '../queries/orderKeys'
import { toast } from 'sonner'
import { OrderDetails } from '@/types/order'

export const useConfirmOrder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: number) => orderService.confirmOrder(id),

    onMutate: async (id) => {
      await queryClient.cancelQueries({
        queryKey: orderKeys.detail(id)
      })

      const previousOrder =
        queryClient.getQueryData<OrderDetails>(
          orderKeys.detail(id)
        )

        queryClient.setQueryData(orderKeys.detail(id), (old: OrderDetails) => ({
          ...old,
          display_status: "confirmed"
        }))

      return { previousOrder }
    },

    onError: (_err, variables, context) => {
      if (context?.previousOrder) {
        queryClient.setQueryData(
          orderKeys.detail(variables),
          context.previousOrder
        )
      }
      toast.error(_err.message || "Failed to confirm order")
    },

    onSettled: (_data, _error, variables) => {
      queryClient.invalidateQueries({
        queryKey: orderKeys.all
      })

      queryClient.invalidateQueries({
        queryKey: orderKeys.detail(variables)
      })
    },

    onSuccess: () => {
      toast.success("Order confirmed successfully")
    }
  })
}