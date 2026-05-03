'use client';

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useCart } from "@/hooks/cart/useCart";
import { useSelectAll } from "@/hooks/cart/useSelectAll";
import { useUpdateItemQuantity } from "@/hooks/cart/useUpdateItemQuantity";
import { useUpdateSelectItem } from "@/hooks/cart/useUpdateSelectItem";
import { formatRupiah } from "@/lib/formatCurrency";
import { Item } from "@/types/cart";
import { Add, Delete, Remove } from "@mui/icons-material";
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from "@tanstack/react-table";
import Image from "next/image";
import { useMemo } from "react";
import LoadingCartTable from "./loading-cart-table";
import { useRemoveItem } from "@/hooks/cart/useRemoveItem";
import { useClearCart } from "@/hooks/cart/useClearCart";

const CartTable = () => {
  const {data: cart, isLoading } = useCart()

  const { mutate: updateQuantity, isPending: isPendingQuantity } = useUpdateItemQuantity()
      
  const onUpdateQuantity = (itemId: number, quantity: number) => {
    updateQuantity({itemId, quantity})
  }

  const { mutate: updateSelect, isPending: isPendingSelect } = useUpdateSelectItem()
      
  const onUpdateSelect = (itemId: number, isSelected: boolean) => {
    updateSelect({itemId, isSelected})
  }

  const { mutate: selectAll } = useSelectAll()
      
  const onSelectAll = (isSelected: boolean) => {
    selectAll(isSelected)
  }

  const { mutate: onRemove } = useRemoveItem()
  const { mutate: onClearCart } = useClearCart()

  const columns: ColumnDef<Item>[] = useMemo(() => [
    {
      id: 'select',
      footer: () => {
        const isAllSelected = cart && cart.items.length > 0 && cart.items.every(item => item.is_selected);

        return (
          <div className="hidden md:flex justify-center">
          <Checkbox
            className="w-6 h-6"
            checked={isAllSelected}
            onCheckedChange={() => onSelectAll(!isAllSelected)}
          />
          </div>
        );
      },
      cell: ({ row }) => (
        <Checkbox
          className="w-6 h-6"
          disabled={isPendingSelect}
          checked={row.original.is_selected}
          onCheckedChange={() => {
            onUpdateSelect(row.original.id, !row.original.is_selected);
          }}
        />
      )
    },
    {
      header: () => (<p className="text-left">Items</p>),
      accessorKey: 'items',
      cell: ({ row }) => (
        <div className='flex items-center gap-2'>
          <div className='w-16 h-16 flex'>
            <img className='object-cover w-full h-full rounded-md' src={row.original.product.main_image_url ?? "/loading.png"} alt={row.original.product.name} width={100} height={100} />
          </div>
          <div className='text-left'>
            <p className=''>{row.original.product.name}</p>
            <p className='font-medium'>{formatRupiah(row.original.price.unit_price)}</p>
            <div className='mt-2 md:mt-0 md:hidden grid gap-2 md:gap-4 text-center md:justify-center'>
              {
                !row.original.is_available && <p className='text-red-500 text-xs md:text-sm w-40'>This product is out of stock. We’ll notify you when restock is ready</p>
              }
              {
                (row.original.stock < row.original.quantity) && <p className='text-orange-500 text-xs md:text-sm w-40'>Only {row.original.stock} left in stock. Please reduce your quantity.</p>
              }
              {
                <div className='flex md:gap-2 items-center md:justify-center'>
                  <Button 
                    disabled={isPendingQuantity}
                    onClick={() => onUpdateQuantity(row.original.id, row.original.quantity-1)} 
                    className='w-6 h-6 cursor-pointer bg-transparent hover:bg-primary hover:text-white border-primary shadow-none'
                    variant="outline"
                  >
                    <Remove fontSize='small' />
                  </Button>
                  <div className='w-8 h-6 flex items-center justify-center'>{row.original.quantity}</div>
                  <Button 
                    disabled={row.original.stock < row.original.quantity || isPendingQuantity} 
                    onClick={() => onUpdateQuantity(row.original.id, row.original.quantity+1)}
                    className='w-6 h-6 cursor-pointer bg-transparent hover:bg-primary hover:text-white border-primary shadow-none'
                    variant="outline"
                  >
                    <Add fontSize='small' />
                  </Button>
                </div>
              }
            </div>
          </div>
        </div>
      ),
      footer: () => (
        <p className="hidden md:block text-left px-1 font-medium">Select All</p>
      )
    },
    {
      header: () => (<p className="hidden md:block ">Quantity</p>),
      id: 'quantity',
      cell: ({ row }) => (
        <div className='hidden md:grid gap-4 text-center justify-center'>
          {
            !row.original.is_available && <p className='text-red-500 text-sm w-40'>This product is out of stock. We’ll notify you when restock is ready</p>
          }
          {
            (row.original.stock < row.original.quantity) && <p className='text-orange-500 text-sm w-40'>Only {row.original.stock} left in stock. Please reduce your quantity.</p>
          }
          {
            <div className='flex gap-2 items-center justify-center'>
              <Button 
                disabled={isPendingQuantity}
                onClick={() => onUpdateQuantity(row.original.id, row.original.quantity-1)} 
                className='w-6 h-6 cursor-pointer bg-transparent hover:bg-primary hover:text-white border-primary shadow-none'
                variant="outline"
              >
                <Remove fontSize='small' />
              </Button>
              <div className='w-8 h-6 flex items-center justify-center'>{row.original.quantity}</div>
              <Button 
                disabled={row.original.stock < row.original.quantity || isPendingQuantity} 
                onClick={() => onUpdateQuantity(row.original.id, row.original.quantity+1)}
                className='w-6 h-6 cursor-pointer bg-transparent hover:bg-primary hover:text-white border-primary shadow-none'
                variant="outline"
              >
                <Add fontSize='small' />
              </Button>
            </div>
          }
        </div>
      ),
      footer: () => (
        <div className="hidden md:block w-full text-center font-medium">
          Total Cart
        </div>
      )
    },
    {
      header: () => (<p className="hidden md:block ">Subtotal</p>),
      id: 'subtotal',
      cell: ({ row }) =>
        (<p className='hidden md:block text-lg '>{formatRupiah(row.original.subtotal)}</p>),
      footer: () => (
        <p className="hidden md:block font-medium text-lg">{formatRupiah(cart?.summary.total_price as number)}</p>
      )
    },
    {
      header: 'Delete',
      cell: ({ row }) => (
        <Button className='text-red-500 justify-self-center bg-transparent hover:bg-gray-200 cursor-pointer' onClick={() => onRemove({id: row.original.id})}>
          <Delete/>
        </Button>
      ),
      footer: () => (
        <Button className="hidden bg-yellow-500 justify-self-center  hover:bg-yellow-600 md:flex gap-0 items-center justify-center" onClick={() => onClearCart()}>
          <Delete/>
          Clear
        </Button>
      )
    },
  ], [cart, isPendingQuantity, isPendingSelect]);

  const emptyCart = Array(4)

  const table = useReactTable({
    data: cart?.items ?? emptyCart,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  if (isLoading) return (
    <LoadingCartTable/>
  )

  return (
    <table className='min-w-full'>
      <thead className="bg-primary/50">
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header, index) => (
              <th
                key={header.id}
                className={`
                  p-2
                  ${index === 0 ? "rounded-l-md" : ""}
                  ${index === headerGroup.headers.length - 1 ? "rounded-r-md" : ""}
                `}
              >
                {flexRender(header.column.columnDef.header, header.getContext())}
              </th>
            )
            )}
          </tr>
        ))}
      </thead>
      <tbody>
        {table.getRowModel().rows.map((row) => (
          <tr key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <td key={cell.id} className={`p-1 md:px-2 md:py-4 border-y border-gray-200 text-center`}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
      <tfoot>
        {table.getFooterGroups().map((footerGroup) => (
          <tr key={footerGroup.id}>
            {footerGroup.headers.map((header) => (
              <th key={header.id}>
                {flexRender(
                  header.column.columnDef.footer,
                  header.getContext()
                )}
              </th>
            ))}
          </tr>
        ))}
      </tfoot>
    </table>
  );
};

export default CartTable;