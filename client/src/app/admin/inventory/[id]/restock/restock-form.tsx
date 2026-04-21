"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { Controller, useForm } from "react-hook-form"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardFooter,
} from "@/components/ui/card"
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupText,
  InputGroupTextarea,
} from "@/components/ui/input-group"
import { InventoryFormValues, inventorySchema } from "@/schemas/inventory.schema"
import { useEffect } from "react"
import { useInventoryById } from "@/hooks/inventory/useInventoryById"
import { useRestock } from "@/hooks/inventory/useRestock"
import { Spinner } from "@/components/ui/spinner"
import StockBadge from "@/components/common/stock-badge"


export default function RestockForm({skuId}: {skuId: number}) {
  const { data: inventory, isLoading } = useInventoryById(skuId)
  const { mutate: restock } = useRestock()

  const form = useForm<InventoryFormValues>({
    resolver: zodResolver(inventorySchema),
    defaultValues: {
      sku_id: skuId,
      action: "restock",
      quantity_change: 0,
      notes: ""
    },
  })

  const stock = (): number => {
    if (inventory) {
      return inventory.stock
    } else {
      return 0
    }
  }

  const quantity = form.watch("quantity_change") ?? 0
  const stockAfterChange = stock() + quantity

  const onRestock = (data: InventoryFormValues) => {
    restock(data)
  };

  useEffect(() => {
    form.reset({
      sku_id: skuId,
      action: "restock"
    })
  }, [skuId, form])

  return (
    <div className="space-y-2">
      <h2 className="text-xl font-bold">Restock</h2>
      <Card className="w-full">
        <CardContent className="space-y-4">
          <form id="restock" onSubmit={form.handleSubmit(onRestock)}>
            <FieldGroup className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <h3 className="font-medium text-sm">Product</h3>
                <p>{inventory?.product}</p>
              </div>
              <div className="space-y-2">
                <h3 className="font-medium text-sm">SKU Code</h3>
                <p className="px-3 py-1 rounded-md bg-gray-200 text-sm w-fit">{inventory?.sku_code}</p>
              </div>
              <div className="space-y-2">
                <h3 className="font-medium text-sm">Variant Label</h3>
                <p className="px-3 py-1 rounded-md bg-gray-200 text-sm w-fit">{inventory?.variant_label}</p>
              </div>
              <div className="space-y-2">
                <h3 className="font-medium text-sm">Availability</h3>
                <StockBadge status={inventory?.availability ?? ""}/>
              </div>
              <Controller
                name={`quantity_change`}
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="restock-quantity-change">
                      Quantity Change
                    </FieldLabel>
                    <Input
                      type="number"
                      {...field}
                      value={field.value}
                      onChange={(e) => {
                        field.onChange(e.target.value === '' ? '' : e.target.valueAsNumber)
                      }}
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <div className="space-y-2">
                <h3 className="font-medium text-sm">Stock After Change</h3>
                <p>{stockAfterChange}</p>
              </div>
              <div className="col-span-2">
                <Controller
                  name="notes"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="restock-notes">
                        Notes
                      </FieldLabel>
                      <InputGroup>
                        <InputGroupTextarea
                          {...field}
                          id="restock-notes"
                          placeholder="Add notes about this restock (e.g., supplier, quantity changes, condition)"
                          rows={6}
                          className="min-h-24 resize-none"
                          aria-invalid={fieldState.invalid}
                        />
                        <InputGroupAddon align="block-end">
                          <InputGroupText className="tabular-nums">
                            {field?.value?.length}/100 characters
                          </InputGroupText>
                        </InputGroupAddon>
                      </InputGroup>
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
                </div>
              </FieldGroup>
          </form>
        </CardContent>
        <CardFooter>
          <Field orientation="horizontal" className="w-full flex items-center justify-center">
            <Button type="button" className="bg-gray-200 text-black hover:bg-gray-300 hover:text-black" onClick={() => form.reset()}>
              Reset
            </Button>
            <Button className="text-white hover:text-white" type="submit" form="restock">
              {isLoading ? <><Spinner/> Adding stock...</> : <>Submit</>}
            </Button>
          </Field>
        </CardFooter>
      </Card>
    </div>
  )
}
