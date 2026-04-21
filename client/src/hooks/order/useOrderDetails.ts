"use client";

import { useQuery } from "@tanstack/react-query"
import { OrderDetails } from "@/types/order";
import { orderKeys } from "../queries/orderKeys";
import { orderService } from "@/services/order.service";
import { Response } from "@/types/response"

export const useOrderDetails = (id: number) => {  
  return useQuery<Response, Error, OrderDetails>({
    queryKey: orderKeys.detail(id),
    queryFn: () => orderService.getOrderDetails(id),
    select: (res) => res.data
  })
}