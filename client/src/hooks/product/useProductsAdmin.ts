"use client";

import { useQuery } from "@tanstack/react-query"
import { productService } from "@/services/product.service"
import { productKeys } from "../queries/productKeys";
import { Response } from "@/types/response";
import { BrowseProductFormValues } from "@/schemas/product.schema";

export const useProductsAdmin = (filters: BrowseProductFormValues) => {
  return useQuery<Response>({
    queryKey: productKeys.list(filters),
    queryFn: () => productService.getProductsByAdmin({...filters, limit: 5}),
  })
}