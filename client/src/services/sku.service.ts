import { api } from "@/lib/api";
import { skuDetails } from "@/types/sku";

export const skuService = {
  getSKUByProductId(id: number): Promise<skuDetails[]> {
    return api.get(`/admin/products/${id}/sku`);
  },
};