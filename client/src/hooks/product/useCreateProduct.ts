"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { useRouter } from "next/navigation"
import { productKeys } from "../queries/productKeys";
import { productService } from "@/services/product.service";
import { ProductFormValues } from "@/schemas/product.schema";
import { toast } from "sonner";

export const useCreateProduct = () => {
  const queryClient = useQueryClient()
  const router = useRouter()

  return useMutation({
    mutationFn: (payload: ProductFormValues) => {
      return productService.createProduct(payload)},

    onSuccess: () => {
      toast.success("Product created successfully")

      // refresh product list cache
      queryClient.invalidateQueries({
        queryKey: productKeys.lists()
      })

      // redirect to product list
      setTimeout(() => {
        router.push("/admin/product")
      }, 800)
    },

    onError: (error) => {
      toast.error(error.message || "Failed to create product")
    }
  })
}