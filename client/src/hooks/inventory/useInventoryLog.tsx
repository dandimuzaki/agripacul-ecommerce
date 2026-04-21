"use client";

import { useQuery } from "@tanstack/react-query"
import { inventoryService } from "@/services/inventory.service"
import { inventoryKeys } from "../queries/inventoryKeys";
import { Response } from "@/types/response";
import { InventoryLogFormValues } from "@/schemas/inventory.schema";

export const useInventoryLogs = (id: number, filters: InventoryLogFormValues) => {  
  return useQuery<Response>({
    queryKey: inventoryKeys.log(id),
    queryFn: () => inventoryService.getInventoryLogsBySKUId(id, filters),
  })
}