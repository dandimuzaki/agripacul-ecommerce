'use client';

import { useMemo } from 'react';
import { Add, Search } from '@mui/icons-material';
import { ProductSummary } from '@/types/product';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { ReusablePagination } from '@/components/common/pagination';
import CategoryFilterDropdown from '@/components/admin/category-dropdown';
import SortProductDropdown from '@/components/admin/sort-product-dropdown';
import { Skeleton } from '../ui/skeleton';

const LoadingProductList = () => {
  const columns: ColumnDef<ProductSummary>[] = useMemo(() => [
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'name',
      cell: ({row}) => (<Skeleton className='text-left w-48 h-16'>{row.original.name}</Skeleton>)
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'image',
      cell: () => (
        <div className='flex justify-center items-center'>
          <Skeleton
            className="h-16 w-16 aspect-square object-cover rounded-md"
          ></Skeleton>
        </div>
      ),
    },
    {
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      accessorKey: 'category.name',
      cell: () => <Skeleton className='bg-gray-200 h-8 w-24'/>
    },
    {
      accessorKey: 'min_price',
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      cell: () => <div className='flex items-center justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      accessorKey: 'max_price',
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      cell: () => <div className='flex items-center justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      accessorKey: 'sold_count',
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      cell: () => <div className='flex items-center justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      accessorKey: 'average_rating',
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      cell: () => <div className='flex items-center justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      accessorKey: 'is_published',
      header: () => <div className='flex justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>,
      cell: () => <div className='flex items-center justify-center'><Skeleton className='bg-gray-200 h-8 w-24'/></div>
    },
    {
      header: 'Action',
      cell: () => (
        <div className='flex gap-2 justify-center'>
            <Skeleton className='h-10 w-10 bg-blue-300'/>
            <Skeleton className='h-10 w-10 bg-red-300'/>
        </div>
      )
    },

  ], []);


  const table = useReactTable({
    data: Array(5).fill(1),
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <>
        <table className='min-w-full text-sm'>
        <thead className=''>
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
    </>
  );
};

export default LoadingProductList;