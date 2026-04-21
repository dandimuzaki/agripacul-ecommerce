import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { reviewService } from "@/services/review.service"
import { reviewKeys } from '../queries/reviewKeys'
import { toast } from 'sonner'
import { Review } from '@/types/review'

export const useDeleteReview = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id }: { id: number }) =>
      reviewService.deleteReview(id),

    onMutate: async ({ id }) => {
      await queryClient.cancelQueries({
        queryKey: reviewKeys.lists(),
      })

      const previousReviews =
        queryClient.getQueryData<Review[]>(
          reviewKeys.lists()
        )

      // Optimistic update: remove review instantly
      queryClient.setQueryData<Review[]>(
        reviewKeys.lists(),
        (old) => old?.filter((cat) => cat.id !== id) || []
      )

      return { previousReviews }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousReviews) {
        queryClient.setQueryData(
          reviewKeys.lists(),
          context.previousReviews
        )
      }

      toast.error(_err.message || "Failed to delete review")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: reviewKeys.lists(),
      })
    },

    onSuccess: () => {
      toast.success("Review deleted successfully")
    },
  })
}