"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { useRouter } from "next/navigation"
import { categoryKeys } from "../queries/categoryKeys";
import { categoryService } from "@/services/category.service";
import { CategoryFormValues } from "@/schemas/category.schema";
import { toast } from "sonner";

export const useCreateCategory = () => {
  const queryClient = useQueryClient()
  const router = useRouter()

  return useMutation({
    mutationFn: (payload: CategoryFormValues) =>
      categoryService.createCategory(payload),

    onSuccess: () => {
      toast.success("Category created successfully")

      // refresh category list cache
      queryClient.invalidateQueries({
        queryKey: categoryKeys.all
      })

      // redirect to category list
      setTimeout(() => {
        router.push("/admin/category")
      }, 800)
    },

    onError: (error: any) => {
      toast.error(error.message || "Failed to create category")
    }
  })
}