"use client"

import { useOrderFilter } from "@/hooks/order/useOrderFilter"
import { Controller } from "react-hook-form"
import { Field } from "../ui/field"
import { RadioGroup, RadioGroupItem } from "../ui/radio-group"
import { useRouter, useSearchParams } from "next/navigation"

const OrderPeriod = () => {
  const {form} = useOrderFilter()
  const searchParams = useSearchParams()
  const router = useRouter()
  const period = [
    {name: "Today", value: "today"},
    {name: "Last 7 Days", value: "last_7_days"},
    {name: "This Month", value: "this_month"}
  ]
  return (
    <Controller
      name="period"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <RadioGroup
            value={field.value?.toString()}
            onValueChange={(value) => {
              field.onChange(value)
              const params = new URLSearchParams(searchParams.toString())
              params.set("period", value)
              router.replace(`/orders?${params.toString()}`)
            }}
            className="flex gap-2 items-center"
          >
            {period.map((p) => (
              <label htmlFor={p.value} key={p.value} className={`flex items-center gap-2 py-2 px-4 border border-gray-300 rounded ${(field.value == p.value) ? 'bg-primary/50' : 'bg-primary/10'}`}>
                <RadioGroupItem className="sr-only" value={p.value} id={p.value} />
                  {p.name}
              </label>
            ))}
          </RadioGroup>
        </Field>
      )}
    />
  )
}

export default OrderPeriod
