"use client"
import { useEffect } from 'react'
import { Controller } from 'react-hook-form'
import { Field } from '../ui/field'
import { useRouter, useSearchParams } from 'next/navigation'
import { useOrderFilter } from '@/hooks/order/useOrderFilter'
import { capitalize } from '@/lib/formatText'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '../ui/select'

const OrderStatusDropdown = () => {
  const statusOptions = [
    "all",
    "pending",
    "paid",
    "processing",
    "shipped",
    "delivered",
    "completed",
    "cancelled"
  ]
  const { form } = useOrderFilter()
  const searchParams = useSearchParams()
  const router = useRouter()

  const status = searchParams.get("status") ?? undefined

  useEffect(() => {
    if (status) {
      form.setValue("status", status)
    } else {
      form.setValue("status", "all")
    }
  }, [status])
  
  return (
    <Controller
      name="status"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <Select
            value={field.value}
            onValueChange={(value) => {
              field.onChange(value)
              const params = new URLSearchParams(searchParams.toString())
              if (value == "all") {
                params.delete("status")
              } else {
                params.set("status", value)
              }
              router.replace(`/orders?${params.toString()}`)
            }}
          >
            <SelectTrigger aria-invalid={fieldState.invalid} className='w-fit border-gray-400'>
              Status: <SelectValue placeholder="Select order status" defaultValue={"Newest to Oldest"} />
            </SelectTrigger>

            <SelectContent>
              {statusOptions.map((s) =>
                <SelectItem key={s} value={s}>
                  {capitalize(s)}
                </SelectItem>
              )}
            </SelectContent>
          </Select>
        </Field>
      )}
    />
  )
}

export default OrderStatusDropdown
