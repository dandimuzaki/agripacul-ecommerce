'use client';

import { useMemo } from 'react';
import { Add, Edit } from '@mui/icons-material';
import { Category } from '@/types/category';
import Image from 'next/image';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { useCategories } from '@/hooks/category/useCategories';
import { ReusablePagination } from '@/components/common/pagination';
import ConfirmDelete from './confirm-delete';
import NoRows from '@/components/admin/no-rows';

const CategoryList = () => {
  const { data: categories, isError, error, isLoading } = useCategories();

  const columns: ColumnDef<Category>[] = useMemo(() => [
    {
      header: 'Name',
      accessorKey: 'name',
      cell: ({row}) => (<p className='text-left max-w-48'>{row.original.name}</p>)
    },
    {
      header: 'Icon',
      accessorKey: 'icon_url',
      cell: ({ row }) => (
        <div className='flex justify-center items-center'>
          <Image
            src={"/cherry-tomato.png"}
            alt={row.original?.name}
            width={100}
            height={100}
            className="h-16 w-16 aspect-square object-cover rounded-md"
          />
        </div>
      ),
    },
    {
      header: 'Product Count',
      accessorKey: 'product_count',
    },
    {
      header: 'Action',
      cell: ({ row }) => (
        <div className='flex gap-2 justify-center'>
          <Link href={`/admin/category/edit/${row.original.id}`}>
            <Button 
              className='bg-blue-500 text-white hover:bg-blue-700 hover:text-white' 
              variant="default"
            >
              <Edit fontSize='small' />
            </Button>
          </Link>
          <ConfirmDelete id={row.original.id}/>
        </div>
      )
    },

  ], []);

  const memoData = useMemo(() => categories?.data ?? [], [categories?.data])

  const table = useReactTable({
    data: memoData,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  const pagination = categories?.pagination;

  if (isError) {
    return <p>{error.message}</p>
  }

  return (
    <div className='space-y-4'>
      <div className='flex items-center justify-between'>
        <p className='font-bold text-2xl'>Category Management</p>
        <Link href={"/admin/category/add"}>
          <Button><Add/><span className="flex-nowrap text-nowrap mr-2">Add Category</span></Button>
        </Link>
      </div>

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
        {!pagination ? <NoRows colSpan={columns.length} text='categories'/> : <tbody>
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

export default CategoryList;