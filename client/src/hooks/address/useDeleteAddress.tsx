import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { addressService } from "@/services/address.service"
import { addressKeys } from '../queries/addressKeys'
import { toast } from 'sonner'
import { Address } from '@/types/address'
import { Response } from '@/types/response'

export const useDeleteAddress = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: number) =>
      addressService.deleteAddress(id),

    onMutate: async (id) => {
      await queryClient.cancelQueries({
        queryKey: addressKeys.all,
      })

      const previousAddresses =
        queryClient.getQueryData<Address[]>(
          addressKeys.all
        )

      // Optimistic update: remove address instantly
      queryClient.setQueryData<Response>(
        addressKeys.all,
        (old) => {
          if (!old) return old;

          return {
            ...old,
            data: old.data.filter(
              (add: Address) => add.id !== id
            ),
          };
        }
      );

      return { previousAddresses }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousAddresses) {
        queryClient.setQueryData(
          addressKeys.all,
          context.previousAddresses
        )
      }

      toast.error(_err.message || "Failed to delete address")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: addressKeys.all,
      })
    },

    onSuccess: () => {
      toast.success("Address deleted successfully")
    },
  })
}