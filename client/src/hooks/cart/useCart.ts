"use client";

import { useQuery } from "@tanstack/react-query"
import { cartKeys } from "../queries/cartKeys";
import { cartService } from "@/services/cart.service";
import { Response } from "@/types/response";
import { Cart } from "@/types/cart";

export const useCart = () => {  
  return useQuery<Response, Error, Cart>({
    queryKey: cartKeys.all,
    queryFn: () => cartService.getCart(),
    select: (res) => res.data
  })
}