import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { productService } from "@/services/product.service"
import { productKeys } from '../queries/productKeys'
import { toast } from 'sonner'
import { ProductDetails } from '@/types/product'
import { GetProductResponse } from '@/types/response'

export const useDeleteImage = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ productId, imageId }: { productId: number, imageId: number }) =>
      productService.deleteProductImage(productId, imageId),

    onMutate: async ({ productId, imageId }) => {
      await queryClient.cancelQueries({
        queryKey: productKeys.adminDetail(productId),
      })

      const previousProduct = queryClient.getQueryData<ProductDetails>(
        productKeys.adminDetail(productId)
      )

      queryClient.setQueryData<GetProductResponse>(
        productKeys.adminDetail(productId),
        (old) => {
          if (!old) return old

          return {
            ...old,
            images: old.data.images.filter((img) => img.id !== imageId),
          }
        }
      )

      return { previousProduct }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousProduct) {
        queryClient.setQueryData(
          productKeys.lists(),
          context.previousProduct
        )
      }

      toast.error(_err.message || "Failed to delete image")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: productKeys.lists(),
      })
    },

    onSuccess: () => {
      toast.success("Image deleted successfully")
    }
  })
}