import { useQuery } from "@tanstack/react-query"
import { orderService } from "@/services/order.service"
import { Response } from "@/types/response"
import { FilterOrderFormValues } from "@/schemas/order.schema"

export const useOrderHistory = (filters: FilterOrderFormValues) => {
  const cleanFilters = Object.fromEntries(
    Object.entries(filters).filter(([_, v]) => v !== undefined || v !== "")
  )

  return useQuery<Response>({
    queryKey: ["orders", "customer_list", 
      filters.status, 
      filters.sort_by, 
      filters.sort_order, 
      filters.period,
      filters.search,
      filters.start_date,
      filters.end_date,
      filters.page,
      filters.limit,
      filters.shipping_method
    ],
    queryFn: () => orderService.getOrderHistory({...cleanFilters, limit: 2}),
    placeholderData: (prev) => prev
  })
}