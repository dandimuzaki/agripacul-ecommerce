import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { categoryService } from "@/services/category.service"
import { categoryKeys } from '../queries/categoryKeys'
import { toast } from 'sonner'
import { Category } from '@/types/category'

export const useDeleteCategory = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id }: { id: number }) =>
      categoryService.deleteCategory(id),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: categoryKeys.all,
      })

      const previousCategories =
        queryClient.getQueryData<Category[]>(
          categoryKeys.all
        )

      // Optimistic update: remove category instantly
      queryClient.setQueryData<Category[]>(
        categoryKeys.all,
        (old) => old?.filter((cat) => cat.id !== id) || []
      )

      return { previousCategories }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousCategories) {
        queryClient.setQueryData(
          categoryKeys.all,
          context.previousCategories
        )
      }

      toast.error(_err.message || "Failed to delete category")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: categoryKeys.all,
      })
    },

    onSuccess: () => {
      toast.success("Category deleted successfully")
    },
  })
}