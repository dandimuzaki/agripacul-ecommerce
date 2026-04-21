"use client";

import { useQuery } from "@tanstack/react-query"
import { reviewService } from "@/services/review.service"
import { reviewKeys } from "../queries/reviewKeys";
import { Review } from "@/types/review";
import { Response } from "@/types/response";

export const useReviewById = (id: number) => {  
  return useQuery<Response, Error, Review>({
    queryKey: reviewKeys.detail(id),
    queryFn: () => reviewService.getReviewById(id),
    select: (res) => res.data
  })
}