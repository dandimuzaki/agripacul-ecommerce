"use client";

import { useQuery } from "@tanstack/react-query"
import { productService } from "@/services/product.service"
import { productKeys } from "../queries/productKeys";
import { Response } from "@/types/response";
import { skuDetails } from "@/types/sku";

export const useSKU = (productId: number) => {  
  return useQuery<Response, Error, skuDetails[]>({
    queryKey: productKeys.sku(productId),
    queryFn: () => productService.getSKUByProductId(productId),
    select: (res) => res.data
  })
}