"use client"

import { Label } from "@/components/ui/label"
import { Slider } from "@/components/ui/slider"
import { useEffect, useState } from "react"
import { Button } from "../ui/button"
import { useRouter, useSearchParams } from "next/navigation"
import { formatRupiah } from "@/lib/formatCurrency"
import { useProductFilter } from "@/hooks/product/useProductFilter"

export function PriceRange() {
  const {form} = useProductFilter()
  const searchParams = useSearchParams()
  const min = Number(searchParams.get("min_price") || 0)
  const max = Number(searchParams.get("max_price") || 1000000)

  useEffect(() => {
    if (min) {
      form.setValue("min_price", min)
    }
    if (max) {
      form.setValue("max_price", max)
    }
  }, [min, max])

  const [value, setValue] = useState([min, max])
  const router = useRouter()

  const handlePriceRange = (value: number[]) => {
    const params = new URLSearchParams(searchParams.toString())

    params.set("min_price", String(value[0]))
    params.set("max_price", String(value[1]))

    router.replace(`/products?${params.toString()}`)
  }

  return (
    <div className="mx-auto grid w-full gap-3">
      <Label htmlFor="price-range" className="text-md border-b border-gray-400 pb-2">Price Range</Label>
      <div className="flex items-center justify-between gap-2">
        <span className="text-sm text-muted-foreground">
          {formatRupiah(value[0])}
        </span>
        <span className="text-sm text-muted-foreground">
          {formatRupiah(value[1])}
        </span>
      </div>
      <Slider
        id="price range"
        value={value}
        onValueChange={setValue}
        min={0}
        max={1000000}
        step={10000}
      />
      <Button onClick={() => handlePriceRange(value)} type="button" variant="outline">Set Price</Button>
    </div>
  )
}
