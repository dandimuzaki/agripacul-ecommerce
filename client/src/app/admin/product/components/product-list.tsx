'use client';

import { useMemo, useState } from 'react';
import { Add, Edit, Search } from '@mui/icons-material';
import { ProductSummary } from '@/types/product';
import Image from 'next/image';
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';
import ConfirmDelete from './confirm-delete';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { useProductsAdmin } from '@/hooks/product/useProductsAdmin';
import { formatRupiah } from '@/lib/formatCurrency';
import { useProductFilter } from '@/hooks/product/useProductFilter';
import { ReusablePagination } from '@/components/common/pagination';
import { useRouter, useSearchParams } from 'next/navigation';
import CategoryFilterDropdown from '@/components/admin/category-dropdown';
import SortProductDropdown from '@/components/admin/sort-product-dropdown';
import PublishSwitch from '@/components/admin/publish';
import LoadingProductList from '@/components/admin/loading-product-list';
import NoRows from '@/components/admin/no-rows';

const ProductList = () => {
  const {filters} = useProductFilter()
  const { data: products, isLoading } = useProductsAdmin(filters);
  const [keyword, setKeyword] = useState("")
  const searchParams = useSearchParams()
  const router = useRouter()

  const columns: ColumnDef<ProductSummary>[] = useMemo(() => [
    {
      header: 'Name',
      accessorKey: 'name',
      cell: ({row}) => (<p className='text-left max-w-48'>{row.original.name}</p>)
    },
    {
      header: 'Image',
      accessorKey: 'image',
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
      header: 'Category',
      accessorKey: 'category.name',
    },
    {
      accessorKey: 'min_price',
      header: ({ column }) => (
        <button onClick={() => column.toggleSorting()} className='cursor-pointer'>
          Min Price
        </button>
      ),
      cell: ({row}) => (<>{formatRupiah(row.original.min_price)}</>)
    },
    {
      accessorKey: 'max_price',
      header: ({ column }) => (
        <button onClick={() => column.toggleSorting()} className='cursor-pointer'>
          Max Price 
        </button>
      ),
      cell: ({row}) => (<>{formatRupiah(row.original.max_price)}</>)
    },
    {
      accessorKey: 'sold_count',
      header: ({ column }) => (
        <button onClick={() => column.toggleSorting()} className='cursor-pointer'>
          Sold Count
        </button>
      )
    },
    {
      accessorKey: 'average_rating',
      header: ({ column }) => (
        <button onClick={() => column.toggleSorting()} className='cursor-pointer'>
          Rating 
        </button>
      ),
      cell: ({ row }) => (<p>{row.original?.average_rating}</p>)
    },
    {
      accessorKey: 'is_published',
      header: ({ column }) => (
        <button onClick={() => column.toggleSorting()} className='cursor-pointer'>
          Published 
        </button>
      ),
      cell: ({row}) => (<div className='flex items-center justify-center'><PublishSwitch product={row.original}/></div>)
    },
    {
      header: 'Action',
      cell: ({ row }) => (
        <div className='flex gap-2 justify-center'>
          <Link href={`/admin/product/edit/${row.original.id}`}>
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

  const memoData = useMemo(() => products?.data ?? [], [products?.data])

  const table = useReactTable({
    data: memoData,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  const pagination = products?.pagination;

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.replace(`/admin/product?${params.toString()}`)
  }

  return (
    <>
      <div className='space-y-4'>

        <div className='flex items-center justify-between'>
          <div className='flex'>
            <input value={keyword} onChange={(e) => setKeyword(e.target.value)}  className='bg-white flex-1 rounded-l-md h-8 px-2 border-y border-l border-gray-300 text-sm' type='text' placeholder="Search product" />
            <button onClick={() => onSubmit(keyword)} className='bg-white rounded-r-md h-8 px-2 hover:bg-gray-300 border-y border-r border-gray-300'><Search /></button>
          </div>
          <div className='flex items-center gap-4'>
            <CategoryFilterDropdown/>
            <SortProductDropdown/>
            <Link href={"/admin/product/add"}>
              <Button><Add/><span className="flex-nowrap text-nowrap mr-2">Add Product</span></Button>
            </Link>
            
          </div>
        </div>
      </div>

      {isLoading ? <LoadingProductList/> : <>
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
        {!pagination ? <NoRows colSpan={columns.length} text='products'/> : <tbody>
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
      </>}
    </>
  );
};

export default ProductList;