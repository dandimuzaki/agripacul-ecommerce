'use client';

import { Skeleton } from "@/components/ui/skeleton";
import { Item } from "@/types/cart";
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { useMemo } from "react";

const LoadingCartTable = () => {
  const columns: ColumnDef<Item>[] = useMemo(() => [
    {
      id: 'select',
      footer: () => <Skeleton className="bg-gray-200 w-6 h-6"/>,
      cell: () => <Skeleton className="bg-gray-200 w-6 h-6"/>
    },
    {
      header: () => <Skeleton className="bg-gray-200 w-24 h-8" />,
      accessorKey: 'items',
      cell: () => (
        <div className='flex items-center gap-2'>
          <div className='w-16 h-16 flex'>
            <Skeleton className='bg-gray-200 flex-1 h-16 w-16'/>
          </div>
          <div className='grid'>
            <Skeleton className='bg-gray-200 h-8 min-w-24 mb-1'/>
            <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>
          </div>
        </div>
      ),
      footer: () => (
        <Skeleton className="bg-gray-200 w-24 h-8" />
      )
    },
    {
      header: () => <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>,
      accessorKey: 'quantity',
      cell: () => <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>,
      footer: () => (
        <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>
      )
    },
    {
      header: () => <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>,
      accessorKey: 'subtotal',
      cell: () => <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>,
      footer: () => (
        <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>
      )
    },
    {
      header: () => <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>,
      id: "action",
      cell: () => <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>,
      footer: () => (
        <div className="flex justify-center"><Skeleton className="bg-gray-200 w-24 h-8" /></div>
      )
    },
  ], []);

  const emptyCart = Array(4).fill(1)

  const table = useReactTable({
    data: emptyCart,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

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
                  ${index === 0 ? "rounded-l-lg" : ""}
                  ${index === headerGroup.headers.length - 1 ? "rounded-r-lg" : ""}
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
              <td key={cell.id} className={`px-2 py-4 border-y border-gray-200 text-center`}>
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
              <th key={header.id} className="p-2">
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

export default LoadingCartTable;