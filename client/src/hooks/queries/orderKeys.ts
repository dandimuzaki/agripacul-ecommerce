import { FilterOrderFormValues } from "@/schemas/order.schema";

export const orderKeys = {
  all: ["orders"] as const,

  adminLists: () => [...orderKeys.all, "admin_list"] as const,

  adminList: (filters: FilterOrderFormValues) =>
    [...orderKeys.adminLists(), filters] as const,

  customerLists: () => [...orderKeys.all, "customer_list"] as const,

  detail: (id: number) =>
    [...orderKeys.all, "detail", id] as const,
}