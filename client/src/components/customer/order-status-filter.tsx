"use client"
import { useEffect } from 'react'
import { Controller } from 'react-hook-form'
import { Field, FieldLabel } from '../ui/field'
import { RadioGroup, RadioGroupItem } from '../ui/radio-group'
import { useRouter, useSearchParams } from 'next/navigation'
import { useOrderFilter } from '@/hooks/order/useOrderFilter'
import { capitalize } from '@/lib/formatText'

const OrderStatusFilter = () => {
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
          <FieldLabel className="text-md border-b border-gray-400 pb-2">
            Status
          </FieldLabel>
          <RadioGroup
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
            className='gap-0'
          >
            {statusOptions.map((s) => (
              <div key={s} className={`flex items-center gap-2 p-2 rounded ${(field.value == s) ? 'bg-primary/20' : ''}`}>
                <RadioGroupItem className="border-gray-400" value={s} id={s} />
                <label htmlFor={s} className="flex-1 cursor-pointer text-sm">
                  {capitalize(s)}
                </label>
              </div>
            ))}
          </RadioGroup>
        </Field>
      )}
    />
  )
}

export default OrderStatusFilter
