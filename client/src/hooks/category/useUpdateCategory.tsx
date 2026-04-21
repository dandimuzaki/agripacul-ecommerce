import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { CategoryFormValues } from "@/schemas/category.schema"
import { categoryService } from "@/services/category.service"
import { categoryKeys } from '../queries/categoryKeys'
import { toast } from 'sonner'
import { Category } from '@/types/category'

export const useUpdateCategory = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      id,
      payload
    }: {
      id: number
      payload: CategoryFormValues
    }) => categoryService.updateCategory(id, payload),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: categoryKeys.detail(id)
      })

      const previousCategory =
        queryClient.getQueryData<Category>(
          categoryKeys.detail(id)
        )
      return { previousCategory }
    },

    onError: (_err, variables, context) => {
      if (context?.previousCategory) {
        queryClient.setQueryData(
          categoryKeys.detail(variables.id),
          context.previousCategory
        )
      }
      toast.error(_err.message || "Failed to update category")
    },

    onSettled: (_data, _error, variables) => {
      queryClient.invalidateQueries({
        queryKey: categoryKeys.all
      })

      queryClient.invalidateQueries({
        queryKey: categoryKeys.detail(variables.id)
      })
    },

    onSuccess: () => {
      toast.success("Category updated successfully")
    }
  })
}