import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { productKeys } from '../queries/productKeys'
import { SKUFormValues } from '@/schemas/sku.schema'
import { productService } from '@/services/product.service'
import { skuDetails } from '@/types/sku'
import { toast } from 'sonner'

export const useUpdateSKU = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      id,
      payload
    }: {
      id: number
      payload: SKUFormValues
    }) => productService.updateSKU(id, payload),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: productKeys.sku(id)
      })

      const previousProduct =
        queryClient.getQueryData<skuDetails[]>(
          productKeys.sku(id)
        )
      return { previousProduct }
    },

    onError: (_err, variables, context) => {
      if (context?.previousProduct) {
        queryClient.setQueryData(
          productKeys.sku(variables.id),
          context.previousProduct
        )
      }
      toast.error(_err.message || "Failed to update stock keeping units")
    },

    onSettled: (_data, _error, variables) => {
      queryClient.invalidateQueries({
        queryKey: productKeys.sku(variables.id)
      })
    },

    onSuccess: () => {
      toast.success("Stock keeping units updated successfully")
    }
  })
}