'use client';

import { useEffect, useMemo, useReducer } from "react";
import { skuDetails, SKUEdit } from "@/types/sku";
import { Controller, useForm } from "react-hook-form";
import { skuFormSchema, SKUFormValues } from "@/schemas/sku.schema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useSKU } from "@/hooks/product/useSKU";
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from "@tanstack/react-table";
import ConfirmDeleteSKU from "./delete-sku";
import { Button } from "@/components/ui/button";
import { Check, Edit } from "@mui/icons-material";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Field, FieldError } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { useUpdateSKU } from "@/hooks/product/useUpdateSKU";
import { Spinner } from "@/components/ui/spinner";

type SKUAction =
  | { type: "INIT"; payload: SKUEdit[] }
  | { type: "EDIT_SKU"; payload: { skuId: number } }
  | { type: "SAVE_SKU"; payload: { skuId: number } }

type SKUState = SKUEdit[]

const skuReducer: React.Reducer<SKUState, SKUAction> = (
  state,
  action
) => {
  switch (action.type) {
    case "INIT":
      return action.payload

    case "EDIT_SKU":
      return state?.map(sku =>
        sku.id === action.payload.skuId
          ? { ...sku, isEditing: !sku.isEditing }
          : { ...sku, isEditing: false }
      )
    
    default:
      return state
  }
}

export function SKUEditor({ productId }: { productId: number }) {
  const { data: skus } = useSKU(productId)
  const normalizeSKU = (skus: skuDetails[]): SKUEdit[] =>
    skus.map((s) => ({
      id: s.id,
      sku_code: s.sku_code,
      price: s.price,
      sale_price: s.sale_price,
      stock: s.stock,
      min_stock: s.min_stock,
      weight: s.weight,
      status: s.status,
      originalSKUCode: s.sku_code,
      originalPrice: s.price,
      originalSalePrice: s.sale_price,
      originalStock: s.stock,
      originalMinStock: s.min_stock,
      originalWeight: s.weight,
      originalStatus: s.status,
      variants: s.variants,
      isEditing: false,
    }))
  
  const normalizedSKUs = useMemo(() => {
    if (!skus) return []
    return normalizeSKU(skus)
  }, [skus])

  const form = useForm<SKUFormValues>({
    resolver: zodResolver(skuFormSchema),
    defaultValues: { skus: [] },
  })

  const [skuState, dispatchSKU] = useReducer(
    skuReducer,
    normalizedSKUs ?? []
  );

  useEffect(() => {
    dispatchSKU({ type: "INIT", payload: normalizedSKUs })
    form.reset({
      skus: normalizedSKUs.map((s) => ({
        id: s.id,
        sku_code: s.sku_code,
        price: s.price,
        sale_price: s.sale_price,
        stock: s.stock,
        min_stock: s.min_stock,
        weight: s.weight,
        status: s.status,
        originalSKUCode: s.sku_code,
        originalPrice: s.price,
        originalSalePrice: s.sale_price,
        originalStock: s.stock,
        originalMinStock: s.min_stock,
        originalWeight: s.weight,
        originalStatus: s.status,
        variants: s.variants,
        isEditing: s.isEditing,
      })),
  })}, [normalizedSKUs, form])

  const { mutate, isPending } = useUpdateSKU()
    
  const onUpdateSKU = (data: SKUFormValues) => {
    mutate({
      id: productId,
      payload: data
    })
  }

  const columns: ColumnDef<SKUEdit>[] = useMemo(() => [
    {
      header: 'Variant',
      accessorKey: 'variants',
      cell: ({ row }) => (
        <div className='h-full w-full space-y-2'>
          {row.original.variants.map((v, index) => (
            <div className="space-y-2" key={index}>
              <p>{v.name} :</p>
              <p className="rounded p-1 bg-gray-200">{v.value}</p>
            </div>
          ))}
        </div>
      ),
    },
    {
      header: 'SKU Code',
      accessorKey: 'sku_code',
      cell: ({ row }) => (
        <div className='h-full text-center w-[100px]'>
          <Controller
            name={`skus.${row.index}.sku_code`}
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <Textarea
                  className="whitespace-wrap text-center"
                  {...field}
                  disabled={!row.original.isEditing}
                  placeholder="ABC-123-EFG"
                  onChange={(e) => {
                    field.onChange(e)
                  }}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </div>
      ),
    },
    {
      header: 'Price (Rp)',
      accessorKey: 'price',
      cell: ({ row }) => (
        <div className='h-full w-full text-center'>
          <Controller
            name={`skus.${row.index}.price`}
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <Input
                  type="number"
                  className="text-center"
                  {...field}
                  disabled={!row.original.isEditing}
                  value={field.value ?? ""}
                  onChange={(e) => {
                    field.onChange(e.target.value)
                  }}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </div>
      ),
    },
    {
      header: 'Sale Price (Rp)',
      accessorKey: 'sale_price',
      cell: ({ row }) => (
        <div className='h-full w-full text-center'>
          <Controller
            name={`skus.${row.index}.sale_price`}
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <Input
                  type="number"
                  className="text-center"
                  {...field}
                  disabled={!row.original.isEditing}
                  value={field.value ?? ""}
                  onChange={(e) => {
                    field.onChange(e.target.value)
                  }}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </div>
      ),
    },
    {
      header: 'Minimal Stock',
      accessorKey: 'min_stock',
      cell: ({ row }) => (
        <div className='h-full w-full text-center'>
          <Controller
            name={`skus.${row.index}.min_stock`}
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <Input
                  type="number"
                  className="text-center"
                  disabled={!row.original.isEditing}
                  value={field.value ?? ""}
                  onChange={(e) => {
                    field.onChange(e.target.value)
                  }}
                />

                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </div>
      ),
    },
    {
      header: 'Weight',
      accessorKey: 'weight',
      cell: ({ row }) => (
        <div className='h-full w-full text-center'>
          <Controller
            name={`skus.${row.index}.weight`}
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <Input
                  type="number"
                  className="text-center"
                  {...field}
                  disabled={!row.original.isEditing}
                  value={field.value ?? ""}
                  onChange={(e) => {
                    field.onChange(e.target.value)
                  }}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </div>
      ),
    },
    {
      header: 'Status',
      accessorKey: 'status',
    },
    {
      header: 'Action',
      cell: ({ row }) => (
        <div className='flex gap-2 justify-center'>
          <Button 
            className={`px-1 text-white hover:text-white 
              ${row.original.isEditing ? 
                'bg-primary hover:bg-primary/90' :
                'bg-blue-500 hover:bg-blue-700'
              }`}
            type="button"
            onClick={() =>
              dispatchSKU({
                type: "EDIT_SKU",
                payload: { skuId: row.original.id },
              })
            }
          >
            {row.original.isEditing ? 
              (<>
              <Check fontSize='small'/>
              
              </>) : 
              (<>
              <Edit fontSize='small'/>
              
              </>)}
          </Button>
          <ConfirmDeleteSKU id={row.original.id}/>
        </div>
      )
    },
  ], [form.control]);

  const table = useReactTable({
    data: skuState,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div className="space-y-2">
      <h2 className="text-xl font-bold">Edit Stock Keeping Unit</h2>

      <Card>
        <CardContent className="space-y-4">
          <form id="edit-sku" onSubmit={form.handleSubmit(onUpdateSKU)}>
          <table className='min-w-full text-sm'>
            <thead className=''>
              {table.getHeaderGroups().map((headerGroup) => (
                <tr key={headerGroup.id}>
                  {headerGroup.headers.map((header) => {
                    return (
                      <th key={header.id} className="p-2 bg-primary/50 text-black">
                        {flexRender(header.column.columnDef.header, header.getContext())}
                      </th>
                    )
                  })}
                </tr>
              ))}
            </thead>
            <tbody>
              {table.getRowModel().rows.map((row) => (
                <tr key={row.id}>
                  {row.getVisibleCells().map((cell) => (
                    <td key={cell.id} className={`p-2 border-y border-primary/50 text-center`}>
                      {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </td>
                  ))}
                </tr>
              ))}
            </tbody>
          </table>
          </form>

          <CardFooter>
            <Field orientation="horizontal" className="w-full flex items-center justify-center">
              <Button type="button" className="bg-gray-200 text-black hover:bg-gray-300 hover:text-black" onClick={() => form.reset()}>
                Reset
              </Button>
              <Button className="text-white" form="edit-sku">
                {isPending ? <><Spinner/> Saving changes...</> : <>Save SKU Changes</>}
              </Button>
            </Field>
          </CardFooter>
        </CardContent>
      </Card>
    </div>
  )
}
