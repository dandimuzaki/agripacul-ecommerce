import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { addressService } from "@/services/address.service"
import { addressKeys } from '../queries/addressKeys'
import { toast } from 'sonner'
import { Address } from '@/types/address'

export const useDeleteAddress = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: number) =>
      addressService.deleteAddress(id),

    onMutate: async (id) => {
      await queryClient.cancelQueries({
        queryKey: addressKeys.lists(),
      })

      const previousCategories =
        queryClient.getQueryData<Address[]>(
          addressKeys.lists()
        )

      // Optimistic update: remove address instantly
      queryClient.setQueryData<Address[]>(
        addressKeys.lists(),
        (old) => old?.filter((cat) => cat.id !== id) || []
      )

      return { previousCategories }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousCategories) {
        queryClient.setQueryData(
          addressKeys.lists(),
          context.previousCategories
        )
      }

      toast.error(_err.message || "Failed to delete address")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: addressKeys.lists(),
      })
    },

    onSuccess: () => {
      toast.success("Address deleted successfully")
    },
  })
}