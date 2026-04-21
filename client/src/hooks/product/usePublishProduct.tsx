import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { productService } from "@/services/product.service"
import { productKeys } from '../queries/productKeys'
import { toast } from 'sonner'

export const usePublishProduct = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      id,
      isPublished
    }: {
      id: number
      isPublished: boolean
    }) => productService.updatePublishProduct(id, isPublished),

    onError: (_err) => {
      toast.error(_err.message || "Failed to update product")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: productKeys.lists()
      })
    },

    onSuccess: () => {
      toast.success("Product updated successfully")
    },
  })
}