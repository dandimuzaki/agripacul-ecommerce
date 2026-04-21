"use client";

import { useQuery } from "@tanstack/react-query"
import { productService } from "@/services/product.service"
import { productKeys } from "../queries/productKeys";
import { Response } from "@/types/response";
import { ProductDetails } from "@/types/product";

export const useProductSlug = (slug: string) => {  
  return useQuery<Response, Error, ProductDetails>({
    queryKey: productKeys.detail(slug),
    queryFn: () => productService.getProductBySlug(slug),
    select: (res) => res.data
  })
}