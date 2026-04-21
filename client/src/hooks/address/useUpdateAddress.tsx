import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { AddressFormValues } from "@/schemas/address.schema"
import { addressService } from "@/services/address.service"
import { addressKeys } from '../queries/addressKeys'
import { toast } from 'sonner'
import { Address } from '@/types/address'

export const useUpdateAddress = (options?: {
  onSuccess?: () => void
}) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      id,
      payload
    }: {
      id: number
      payload: AddressFormValues
    }) => addressService.updateAddress(id, payload),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: addressKeys.detail(id)
      })

      const previousAddress =
        queryClient.getQueryData<Address>(
          addressKeys.detail(id)
        )
      return { previousAddress }
    },

    onError: (_err, variables, context) => {
      if (context?.previousAddress) {
        queryClient.setQueryData(
          addressKeys.detail(variables.id),
          context.previousAddress
        )
      }
      toast.error(_err.message || "Failed to update address")
    },

    onSettled: (_data, _error, variables) => {
      queryClient.invalidateQueries({
        queryKey: addressKeys.lists()
      })

      queryClient.invalidateQueries({
        queryKey: addressKeys.detail(variables.id)
      })
    },

    onSuccess: () => {
      toast.success("Address updated successfully")

      options?.onSuccess?.()
    },
  })
}