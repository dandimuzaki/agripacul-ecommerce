"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { reviewKeys } from "../queries/reviewKeys";
import { reviewService } from "@/services/review.service";
import { CreateReviewFormValues } from "@/schemas/review.schema";
import { toast } from "sonner";
import { Dispatch, SetStateAction } from "react";

export const useCreateReview = (setOpenAfterReview: Dispatch<SetStateAction<boolean>>) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: CreateReviewFormValues) =>
      reviewService.createReview(payload),

    onSuccess: () => {
      setOpenAfterReview(true)

      // refresh review list cache
      queryClient.invalidateQueries({
        queryKey: reviewKeys.lists()
      })
    },

    onError: (error: any) => {
      toast.error(error.message || "Failed to create review")
    }
  })
}