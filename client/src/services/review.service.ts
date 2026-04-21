import { api } from "@/lib/api";
import { ReviewFilterFormValues, CreateReviewFormValues } from "@/schemas/review.schema";
import { Response } from "@/types/response";

export const reviewService = {
  getReviews(params: ReviewFilterFormValues): Promise<Response> {
    return api.get('/reviews', {params})
  },

  getReviewsByProduct(productId: number, params: ReviewFilterFormValues): Promise<Response> {
    return api.get(`/reviews/${productId}`, {params})
  },

  getReviewById(id: number): Promise<Response> {
    return api.get(`/reviews/details/${id}`)
  },

  createReview(payload: CreateReviewFormValues) {
    return api.post("/reviews", payload)
  },

  deleteReview(id: number) {
    return api.delete(`/reviews/${id}`)
  },
};