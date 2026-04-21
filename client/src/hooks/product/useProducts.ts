"use client";

import { useQuery } from "@tanstack/react-query"
import { productService } from "@/services/product.service"
import { productKeys } from "../queries/productKeys";
import { Response } from "@/types/response";
import { ProductSummary } from "@/types/product";
import { BrowseProductFormValues } from "@/schemas/product.schema";

export const useProducts = (filters: BrowseProductFormValues) => {  
  return useQuery<Response, Error, ProductSummary[]>({
    queryKey: productKeys.list(filters),
    queryFn: () => productService.getProducts(filters),
    select: (res) => res.data
  })
}