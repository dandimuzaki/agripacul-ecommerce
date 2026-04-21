'use client';

import { useMemo } from 'react';
import { OrderSummary } from '@/types/order';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { Skeleton } from '../ui/skeleton';

const LoadingInventoryList = () => {
  const columns: ColumnDef<OrderSummary>[] = useMemo(() => [
    {
      header: () => (<div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>),
      accessorKey: 'product',
      cell: () => <div className='space-y-2 text-left'>
        <Skeleton className='bg-gray-200 h-8 w-16'/>
        <Skeleton className='bg-gray-200 h-8 w-24'/>
      </div>
    },
    {
      header: () => (<div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>),
      accessorKey: 'sku_code',
      cell: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      header: () => (<div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>),
      accessorKey: 'variant_label',
      cell: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'stock',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex-col gap-2 flex justify-center'>
        <Skeleton className='bg-gray-200 h-8 w-16'/>
        <Skeleton className='bg-gray-200 h-8 w-24'/>
      </div>,
      accessorKey: 'minimal_stock',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'availability',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'status',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      id: "action",
      cell: () => 
        <div className='flex justify-center'>
          <Skeleton className='h-10 w-10 bg-green-300'/>
        </div>
    },

  ], []);

  const table = useReactTable({
    data: Array(5).fill(1),
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div className='space-y-4'>
      <div className='flex items-center justify-between'>
        <Skeleton className='bg-gray-200 h-8 w-36' />
        <Skeleton className='bg-gray-200 h-8 w-36' />
      </div>

      <table className='min-w-full text-sm'>
        <thead>
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <th key={header.id} className='p-2 bg-gray-300 text-center'>
                  {flexRender(header.column.columnDef.header, header.getContext())}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody>
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <td key={cell.id} className={`p-2 border-y border-gray-300 text-center`}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
      <div className="flex justify-center">
        <Skeleton className='bg-gray-200 h-8 w-48' />
      </div>
    </div>
  );
};

export default LoadingInventoryList;