import { api } from "@/lib/api";
import { AdjustStockFormValues, FilterInventoryFormValues, InventoryFormValues, InventoryLogFormValues } from "@/schemas/inventory.schema";
import { Response } from "@/types/response";

export const inventoryService = {
  getInventories(params: FilterInventoryFormValues): Promise<Response> {
    return api
      .get(`/admin/inventories`, {params} )
  },

  getInventoryBySKUId(skuId: number): Promise<Response> {
    return api
      .get(`/admin/inventories/${skuId}`)
  },

  getInventoryLogsBySKUId(skuId: number, params: InventoryLogFormValues): Promise<Response> {
    return api
      .get(`/admin/inventories/logs/${skuId}`, {params})
  },

  restock(payload: InventoryFormValues) {
    return api
      .post(`/admin/inventories`, payload)
  },

  adjustment(payload: AdjustStockFormValues) {
    return api
      .post(`/admin/inventories`, payload)
  },
};