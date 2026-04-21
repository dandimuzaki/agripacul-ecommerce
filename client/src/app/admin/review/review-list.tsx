'use client';

import { useMemo, useState } from 'react';
import { Search } from '@mui/icons-material';
import { Review } from '@/types/review';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { useReviews } from '@/hooks/review/useReviews';
import { formatTime } from '@/lib/formatDate';
import { useReviewFilter } from '@/hooks/review/useReviewFilter';
import { useRouter, useSearchParams } from 'next/navigation';
import { ReusablePagination } from '@/components/common/pagination';
import NoRows from '@/components/admin/no-rows';

const ReviewList = () => {
  const {filters} = useReviewFilter()
  const { data: reviews, isLoading } = useReviews(filters);
  const [keyword, setKeyword] = useState("")
  const searchParams = useSearchParams()
  const router = useRouter()

  const columns: ColumnDef<Review>[] = useMemo(() => [
    {
      header: 'Created At',
      accessorKey: 'created_at',
      cell: ({row}) => (<>{formatTime(row.original.created_at)}</>)
    },
    {
      header: 'Customer',
      accessorKey: 'customer_name',
      cell: ({row}) => (<div className='space-y-1 text-left'>
        <p>{row.original.customer_name}</p>
      </div>)
    },
    {
      header: 'Product',
      accessorKey: 'product_name',
      cell: ({row}) => (<div className='text-left'>{row.original.product_name}</div>)
    },
    {
      header: 'Rating',
      accessorKey: 'rating',
      cell: ({row}) => (<>{row.original.rating}</>)
    },
    {
      header: 'Comment',
      accessorKey: 'comment',
      cell: ({row}) => (<div className='text-left min-w-100'>{row.original.comment}</div>)
    },
  ], []);

  const memoData = useMemo(() => reviews?.data ?? [], [reviews?.data])

  const table = useReactTable({
    data: memoData,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  if (isLoading) return <p>Loading...</p>

  const pagination = reviews?.pagination;

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.replace(`/admin/reviews?${params.toString()}`)
  }

  return (
    <div className='space-y-4'>
      <div className='flex items-center justify-between'>
        <div className='flex'>
          <input value={keyword} onChange={(e) => setKeyword(e.target.value)}  className='bg-white flex-1 rounded-l-md h-8 px-2 breview-y breview-l breview-gray-300 text-sm' type='text' placeholder="Search Review" />
          <button onClick={() => onSubmit(keyword)} className='text-gray-500 bg-white rounded-r-md h-8 px-2 hover:bg-gray-300 breview-y breview-r breview-gray-300'><Search /></button>
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
        {!pagination ? 
          <NoRows colSpan={columns.length} text='reviews'/> : 
          <tbody>
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <td key={cell.id} className={`p-2 breview-y breview-gray-300 text-center`}>
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

export default ReviewList;