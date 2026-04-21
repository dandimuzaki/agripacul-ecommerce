"use client"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { useOrderFilter } from "@/hooks/order/useOrderFilter"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect, useTransition } from "react"
import OrderList from "./order-list"
import LoadingOrderList from "@/components/admin/loading-order-list"

export function OrderTabs() {
  const { form } = useOrderFilter()
  const searchParams = useSearchParams()
  const router = useRouter()

  const [isPending, startTransition] = useTransition()

  const status = searchParams.get("status") ?? "all"

  useEffect(() => {
    form.setValue("status", status)
  }, [status, form])

  const handleStatusFilter = (status: string) => {
    startTransition(() => {
      const params = new URLSearchParams(searchParams.toString())

      if (status === "all") {
        params.delete("status")
      } else {
        params.set("status", status)
      }

      router.replace(`/admin/order?${params.toString()}`)
    })
  }

  return (
    <Tabs
      value={status}
      onValueChange={handleStatusFilter}
    >
      <TabsList variant="line">
        <TabsTrigger value="all">All</TabsTrigger>
        <TabsTrigger value="pending">Pending</TabsTrigger>
        <TabsTrigger value="shipped">Shipped</TabsTrigger>
        <TabsTrigger value="delivered">Delivered</TabsTrigger>
      </TabsList>
      <TabsContent value={status} className="pt-2">
        {isPending ? <LoadingOrderList/> : <OrderList/>}
      </TabsContent>
    </Tabs>
  )
}