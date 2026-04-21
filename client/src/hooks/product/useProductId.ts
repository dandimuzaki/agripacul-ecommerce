"use client";

import { useQuery } from "@tanstack/react-query"
import { productService } from "@/services/product.service"
import { productKeys } from "../queries/productKeys";
import { Response } from "@/types/response";
import { ProductDetails } from "@/types/product";

export const useProductId = (id: number) => {  
  return useQuery<Response, Error, ProductDetails>({
    queryKey: productKeys.adminDetail(id),
    queryFn: () => productService.getProductById(id),
    select: (res) => res.data
  })
}