import { InventoryFormValues } from "@/schemas/inventory.schema"
import { inventoryService } from "@/services/inventory.service"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { inventoryKeys } from "../queries/inventoryKeys"
import { Inventory } from "@/types/inventory"
import { toast } from "sonner"
import { useRouter } from "next/navigation"

export const useRestock = () => {
  const queryClient = useQueryClient()
  const router = useRouter()

  return useMutation({
    mutationFn: (payload: InventoryFormValues) => inventoryService.restock(payload),

    onMutate: async (payload) => {
      await queryClient.cancelQueries({
        queryKey: inventoryKeys.all
      })

      const previousInventory =
        queryClient.getQueryData<Inventory[]>(inventoryKeys.all)

        queryClient.setQueryData<Inventory[]>(
          inventoryKeys.all,
          (old = []) =>
            old.map((item) =>
              item.id === payload.sku_id
                ? {
                    ...item,
                    stock: item.stock + payload.quantity_change
                  }
                : item
            )
        )

      return { previousInventory }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousInventory) {
        queryClient.setQueryData(
          inventoryKeys.all,
          context.previousInventory
        )
      }
      toast.error(_err.message || "Failed to restock the product")
    },

    onSettled: () => {
      queryClient.invalidateQueries({
        queryKey: inventoryKeys.all
      })

      toast.success("Successfully restock the product")

      // redirect to inventory list
      setTimeout(() => {
        router.push("/admin/inventory")
      }, 800)
    }
  })
}