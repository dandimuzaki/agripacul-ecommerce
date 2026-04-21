'use client';

import { useMemo, useState } from 'react';
import { Search } from '@mui/icons-material';
import { OrderSummary } from '@/types/order';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { useOrders } from '@/hooks/order/useOrders';
import { formatTime } from '@/lib/formatDate';
import { formatRupiah } from '@/lib/formatCurrency';
import OrderBadge from '@/components/common/order-badge';
import { OrderAction } from '@/components/admin/order-action';
import { useOrderFilter } from '@/hooks/order/useOrderFilter';
import { useRouter, useSearchParams } from 'next/navigation';
import SortOrderDropdown from '@/components/admin/sort-order-dropdown';
import { ReusablePagination } from '@/components/common/pagination';
import LoadingOrderList from '@/components/admin/loading-order-list';
import NoRows from '@/components/admin/no-rows';

const OrderList = () => {
  const {filters} = useOrderFilter()
  const { data: orders, isLoading } = useOrders(filters);
  const [keyword, setKeyword] = useState("")
  const searchParams = useSearchParams()
  const router = useRouter()

  const columns: ColumnDef<OrderSummary>[] = useMemo(() => [
    {
      header: 'Order ID',
      accessorKey: 'id',
    },
    {
      header: 'Created At',
      accessorKey: 'created_at',
      cell: ({row}) => (<>{formatTime(row.original.created_at)}</>)
    },
    {
      header: 'Customer',
      accessorKey: 'customer.name',
      cell: ({row}) => (<div className='space-y-1 text-left'>
        <p>{row.original.customer.name}</p>
        <p>{row.original.customer.email}</p>
      </div>)
    },
    {
      header: 'Total Items',
      accessorKey: 'item_count',
      cell: ({row}) => (<>{row.original.item_count}</>)
    },
    {
      header: 'Grand Total',
      accessorKey: 'grand_total',
      cell: ({row}) => (<>{formatRupiah(row.original.grand_total)}</>)
    },
    {
      header: 'Status',
      accessorKey: 'display_status',
      cell: ({row}) => (<div className='flex justify-center'><OrderBadge status={row.original.display_status}/></div>)
    },
    {
      header: 'Action',
      cell: ({ row }) => (
        <div className='flex gap-2 justify-center'>
          <OrderAction order={row.original}/>
        </div>
      )
    },

  ], []);

  const memoData = useMemo(() => orders?.data ?? [], [orders?.data])

  const table = useReactTable({
    data: memoData,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  if (isLoading) return <LoadingOrderList/>

  const pagination = orders?.pagination;

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.replace(`/admin/order?${params.toString()}`)
  }

  return (
    <div className='space-y-4'>
      <div className='flex items-center justify-between'>
        <div className='flex'>
          <input value={keyword} onChange={(e) => setKeyword(e.target.value)}  className='bg-white flex-1 rounded-l-md h-8 px-2 border-y border-l border-gray-300 text-sm' type='text' placeholder="Search Order" />
          <button onClick={() => onSubmit(keyword)} className='text-gray-500 bg-white rounded-r-md h-8 px-2 hover:bg-gray-300 border-y border-r border-gray-300'><Search /></button>
        </div>
        <div className='flex items-center gap-4'>
          <SortOrderDropdown/>                      
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
        {!pagination ? <NoRows colSpan={columns.length} text='orders'/> : <tbody>
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

export default OrderList;