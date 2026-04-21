'use client';

import { useMemo } from 'react';
import { InventoryLog } from '@/types/inventory';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import { useInventoryFilter } from '@/hooks/inventory/useInventoryFilter';
import { ReusablePagination } from '@/components/common/pagination';
import LoadingInventoryList from '@/components/admin/loading-inventory-list';
import { useInventoryLogs } from '@/hooks/inventory/useInventoryLog';
import InventoryTypeBadge from '@/components/common/inventory-type-badge';
import { formatTime } from '@/lib/formatDate';

const InventoryLogs = ({id}: {id: number}) => {
  const {filters} = useInventoryFilter()
  const { data: logs, isLoading } = useInventoryLogs(id, filters);

  const columns: ColumnDef<InventoryLog>[] = useMemo(() => [
    {
      header: 'Created At',
      accessorKey: 'created_at',
      cell: ({row}) => <>{formatTime(row.original.created_at)}</>
    },
    {
      header: 'Type',
      accessorKey: 'type',
      cell: ({row}) => <div className='flex justify-center'><InventoryTypeBadge type={row.original.type}/></div>
    },
    {
      header: () => (
        <span className="wrap">Quantity Change</span>
      ),
      accessorKey: 'quantity_change',
    },
    {
      header: () => (
        <span className="wrap">Stock After</span>
      ),
      accessorKey: 'current_stock_after',
    },
    {
      header: 'Reference ID',
      accessorKey: 'reference_id',
    },
    {
      header: 'Reference Type',
      accessorKey: 'reference_type',
    },
    {
      header: 'Notes',
      accessorKey: 'notes',
      cell: ({row}) => <p className='text-left w-56'>{row.original.notes}</p>
    }

  ], []);

  const memoData = useMemo(() => logs?.data ?? [], [logs?.data])
  
  const table = useReactTable({
    data: memoData,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  if (isLoading) return <LoadingInventoryList/>

  const pagination = logs?.pagination;

  if (!pagination) {
    return <p>No inventory logs match your filters</p>;
  }

  return (
    <div className='space-y-4'>
      <table className='min-w-full text-sm'>
        <thead>
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
                <td key={cell.id} className={`p-2 border-y border-gray-300 text-center`}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
      <ReusablePagination currentPage={pagination.page} totalPages={pagination.total_pages}/>
    </div>
  );
};

export default InventoryLogs;