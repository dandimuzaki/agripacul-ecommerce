import { useQuery } from "@tanstack/react-query"
import { orderKeys } from "../queries/orderKeys"
import { orderService } from "@/services/order.service"
import { Response } from "@/types/response"
import { FilterOrderFormValues } from "@/schemas/order.schema"

export const useOrders = (filters: FilterOrderFormValues) => {
  return useQuery<Response>({
    queryKey: orderKeys.adminList(filters),
    queryFn: () => orderService.getAllOrders(filters)
  })
}