"use client";

import { useQuery } from "@tanstack/react-query"
import { reviewService } from "@/services/review.service"
import { reviewKeys } from "../queries/reviewKeys";
import { Response } from "@/types/response";
import { ReviewFilterFormValues } from "@/schemas/review.schema";

export const useReviews = (filters: ReviewFilterFormValues) => {  
  return useQuery<Response>({
    queryKey: reviewKeys.all,
    queryFn: () => reviewService.getReviews(filters)
  })
}