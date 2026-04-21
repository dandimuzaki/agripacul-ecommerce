import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { addressService } from "@/services/address.service"
import { addressKeys } from '../queries/addressKeys'
import { toast } from 'sonner'

export const useSetDefaultAddress = (options?: {
  onSuccess?: () => void
}) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: number) => addressService.setDefaultAddress(id),

    onError: (_err) => {
      toast.error(_err.message || "Failed to update address")
    },

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: addressKeys.lists()
      })
      
      toast.success("Address updated successfully")

      options?.onSuccess?.()
    },
  })
}