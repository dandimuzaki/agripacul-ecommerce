"use client"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { useInventoryFilter } from "@/hooks/inventory/useInventoryFilter"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect, useTransition } from "react"
import InventoryList from "./inventory-list"
import LoadingInventoryList from "@/components/admin/loading-inventory-list"

export function InventoryTabs() {
  const { form } = useInventoryFilter()
  const searchParams = useSearchParams()
  const router = useRouter()

  const [isPending, startTransition] = useTransition()
  const stock = searchParams.get("stock") ?? "all"

  useEffect(() => {
    form.setValue("stock", stock)
  }, [stock, form])

  const handleStatusFilter = (stock: string) => {
    startTransition(() => {
      const params = new URLSearchParams(searchParams.toString())

      if (stock === "all") {
        params.delete("stock")
      } else {
        params.set("stock", stock)
      }

      router.replace(`/admin/inventory?${params.toString()}`)
    })
  }

  return (
    <Tabs
      value={stock}
      onValueChange={handleStatusFilter}
    >
      <TabsList variant="line">
        <TabsTrigger value="all">All</TabsTrigger>
        <TabsTrigger value="in">In Stock</TabsTrigger>
        <TabsTrigger value="low">Low Stock</TabsTrigger>
        <TabsTrigger value="out">Out of Stock</TabsTrigger>
      </TabsList>
      <TabsContent value={stock} className="pt-2">
        {isPending ? <LoadingInventoryList/> : <InventoryList/>}
      </TabsContent>
    </Tabs>
  )
}