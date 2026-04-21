'use client';

import { useMemo, useState } from 'react';
import { Search } from '@mui/icons-material';
import { Inventory } from '@/types/inventory';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { useInventories } from '@/hooks/inventory/useInventories';
import { useInventoryFilter } from '@/hooks/inventory/useInventoryFilter';
import { ReusablePagination } from '@/components/common/pagination';
import { useRouter, useSearchParams } from 'next/navigation';
import { InventoryAction } from '@/components/admin/inventory-action';
import SortInventoryDropdown from '@/components/admin/sort-inventory-dropdown';
import LoadingInventoryList from '@/components/admin/loading-inventory-list';
import StockBadge from '@/components/common/stock-badge';
import SKUStatusBadge from '@/components/common/sku-status-badge';
import NoRows from '@/components/admin/no-rows';

const InventoryList = () => {
  const {filters} = useInventoryFilter()
  const { data: inventories, isLoading } = useInventories(filters);
  const [keyword, setKeyword] = useState("")
  const searchParams = useSearchParams()
  const router = useRouter()

  const columns: ColumnDef<Inventory>[] = useMemo(() => [
    {
      header: () => <p className='text-left'>Product</p>,
      accessorKey: 'product',
      cell: ({row}) => <p className='text-left w-48'>{row.original.product}</p>
    },
    {
      header: 'SKU Code',
      accessorKey: 'sku_code',
      cell: ({row}) => <p className='line-clamp-2 w-36'>{row.original.sku_code}</p>
    },
    {
      header: 'Variant Label',
      accessorKey: 'variant_label',
    },
    {
      header: 'Stock',
      accessorKey: 'stock',
    },
    {
      header: () => (
        <span className="wrap">Minimum Stock</span>
      ),
      accessorKey: 'min_stock',
    },
    {
      header: 'Availability',
      accessorKey: 'availability',
      cell: ({ row }) => <div className='flex justify-center nowrap'><StockBadge status={row.original?.availability}/></div>
    },
    {
      header: 'Status',
      accessorKey: 'status',
      cell: ({ row }) => <div className='flex justify-center'><SKUStatusBadge status={row.original?.status}/></div>
    },
    {
      header: 'Action',
      cell: ({ row }) => (
        <div className='flex gap-2 justify-center'>
          <InventoryAction inventory={row.original}/>
        </div>
      )
    },

  ], []);

  const memoData = useMemo(() => inventories?.data ?? [], [inventories?.data])
  
  const table = useReactTable({
    data: memoData,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  if (isLoading) return <LoadingInventoryList/>

  const pagination = inventories?.pagination;

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.replace(`/admin/inventory?${params.toString()}`)
  }

  return (
    <div className='space-y-4'>
      <div className='flex items-center justify-between'>
        <div className='flex'>
          <input value={keyword} onChange={(e) => setKeyword(e.target.value)}  className='bg-white flex-1 rounded-l-md h-8 px-2 border-y border-l border-gray-300 text-sm' type='text' placeholder="Search inventory" />
          <button onClick={() => onSubmit(keyword)} className='bg-white rounded-r-md h-8 px-2 hover:bg-gray-300 border-y border-r border-gray-300'><Search /></button>
        </div>
      <div className='flex items-center gap-4'>
        <SortInventoryDropdown/>
      </div>
    </div>

      <table className='min-w-full text-sm'>
        <thead>
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header, index) => (
              <th
                key={header.id}
                className={`
                  p-2 bg-primary/50
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
        {!pagination ? <NoRows colSpan={columns.length} text='inventories'/> : <tbody>
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <td key={cell.id} className={`p-2 border-y border-gray-300 text-center`}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>}
      </table>
      <ReusablePagination currentPage={pagination?.page || 1} totalPages={pagination?.total_pages || 1}/>
    </div>
  );
};

export default InventoryList;