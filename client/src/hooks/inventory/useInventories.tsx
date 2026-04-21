"use client";

import { useQuery } from "@tanstack/react-query"
import { inventoryService } from "@/services/inventory.service"
import { inventoryKeys } from "../queries/inventoryKeys";
import { Response } from "@/types/response";
import { Inventory } from "@/types/inventory";
import { FilterInventoryFormValues } from "@/schemas/inventory.schema";

export const useInventories = (filters: FilterInventoryFormValues) => {  
  return useQuery<Response>({
    queryKey: inventoryKeys.list(filters),
    queryFn: () => inventoryService.getInventories(filters),
  })
}