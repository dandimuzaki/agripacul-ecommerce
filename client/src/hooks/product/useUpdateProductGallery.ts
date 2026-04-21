import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { ProductGalleryFormValues } from "@/schemas/product.schema"
import { productService } from "@/services/product.service"
import { productKeys } from '../queries/productKeys'
import { PreviewImage, ProductDetails } from '@/types/product'
import { toast } from 'sonner'
import { Dispatch, SetStateAction } from 'react'

export const useUpdateGallery = ({setLocalPreviews}: {setLocalPreviews: Dispatch<SetStateAction<PreviewImage[]>>}) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      id,
      payload
    }: {
      id: number
      payload: ProductGalleryFormValues
    }) => productService.updateProductGallery(id, payload),

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
      toast.error(_err.message || "Failed to update product gallery")
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
      toast.success("Product gallery updated successfully")
      setLocalPreviews([])
    }
  })
}