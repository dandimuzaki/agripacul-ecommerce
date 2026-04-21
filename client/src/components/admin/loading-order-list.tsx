'use client';

import { useMemo } from 'react';
import { OrderSummary } from '@/types/order';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { Skeleton } from '../ui/skeleton';

const LoadingOrderList = () => {
  const columns: ColumnDef<OrderSummary>[] = useMemo(() => [
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'id',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'created_at',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'customer.name',
      cell: () => <div className='space-y-2 text-left'>
        <Skeleton className='bg-gray-200 h-8 w-24'/>
        <Skeleton className='bg-gray-200 h-8 w-36'/>
      </div>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'item_count',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'grand_total',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'display_status',
      cell: () => <div className='flex justify-center'>
        <Skeleton className='h-10 w-24 bg-yellow-300'/>
      </div>
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

export default LoadingOrderList;