import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { productService } from "@/services/product.service"
import { productKeys } from '../queries/productKeys'
import { toast } from 'sonner'
import { ProductSummary } from '@/types/product'

export const useDeleteProduct = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id }: { id: number }) =>
      productService.deleteProduct(id),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: productKeys.lists(),
      })

      const previousProducts =
        queryClient.getQueryData<ProductSummary[]>(
          productKeys.lists()
        )

      // Optimistic update: remove product instantly
      queryClient.setQueryData<ProductSummary[]>(
        productKeys.lists(),
        (old) => old?.filter((cat) => cat.id !== id) || []
      )

      return { previousProducts }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousProducts) {
        queryClient.setQueryData(
          productKeys.lists(),
          context.previousProducts
        )
      }

      toast.error(_err.message || "Failed to delete product")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: productKeys.lists(),
      })
    },

    onSuccess: () => {
      toast.success("Product deleted successfully")
    },
  })
}