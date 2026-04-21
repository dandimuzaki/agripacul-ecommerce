"use client";

import { useQuery } from "@tanstack/react-query"
import { inventoryService } from "@/services/inventory.service"
import { inventoryKeys } from "../queries/inventoryKeys";
import { Response } from "@/types/response";
import { Inventory } from "@/types/inventory";

export const useInventoryById = (id: number) => {  
  return useQuery<Response, Error, Inventory>({
    queryKey: inventoryKeys.detail(id),
    queryFn: () => inventoryService.getInventoryBySKUId(id),
    select: (res) => res.data
  })
}