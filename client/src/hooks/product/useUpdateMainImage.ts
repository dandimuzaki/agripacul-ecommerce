import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { ProductFormValues, ProductMainImageFormValues } from "@/schemas/product.schema"
import { productService } from "@/services/product.service"
import { productKeys } from '../queries/productKeys'
import { ProductDetails } from '@/types/product'
import { toast } from 'sonner'

export const useUpdateMainImage = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      id,
      payload
    }: {
      id: number
      payload: ProductMainImageFormValues
    }) => productService.updateProductMainImage(id, payload),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: productKeys.adminDetail(id)
      })

      const previousProduct =
        queryClient.getQueryData<ProductDetails>(
          productKeys.adminDetail(id)
        )
      return { previousProduct }
    },

    onError: (_err, variables, context) => {
      if (context?.previousProduct) {
        queryClient.setQueryData(
          productKeys.adminDetail(variables.id),
          context.previousProduct
        )
      }
      toast.error(_err.message || "Failed to update main image")
    },

    onSettled: (_data, _error, variables) => {
      queryClient.invalidateQueries({
        queryKey: productKeys.lists()
      })

      queryClient.invalidateQueries({
        queryKey: productKeys.adminDetail(variables.id)
      })
    },

    onSuccess: () => {
      toast.success("Main image updated successfully")
    }
  })
}