"use client";

import { PriceRange } from "@/components/common/price-range";
import { RatingList } from "@/components/common/rating-list";
import CategoryFilter from "@/components/customer/category-filter";
import SortProductDropdown from "@/components/customer/sort-product-dropdown";
import { Button } from "@/components/ui/button";
import { InputGroup, InputGroupInput } from "@/components/ui/input-group";
import { SearchIcon } from "lucide-react";
import { useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";
import CategoryFilterDropdown from "./category-filter-dropdown";

const ProductFilter = () => {
  const [keyword, setKeyword] = useState("");
  const searchParams = useSearchParams();
  const router = useRouter()

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.push(`/products?${params.toString()}`)
  }

  return (
    <>
      <div className="hidden md:block h-fit">
        <p className="font-semibold text-lg">Filter Options</p>
        <form id="product-filter" className="space-y-4">
          <CategoryFilter/>
          <PriceRange/>
          <RatingList/>
        </form>
      </div>

      <div className="space-y-2 md:hidden">
        <InputGroup className={`rounded-full bg-white border-0 overflow-hidden`}>
          <InputGroupInput
            className='h-4'
            placeholder="Search products..."
            value={keyword}
            onChange={(e) => setKeyword(e.target.value)}
          />
          <Button variant="searchIcon" type='button' className='bg-white' onClick={() => onSubmit(keyword)}>
            <SearchIcon/>
          </Button>
        </InputGroup>

        <div className="flex items-center justify-between">
          <CategoryFilterDropdown/>
          <div>
          <SortProductDropdown/>
          </div>
        </div>
      </div>
    </>
  )
}

export default ProductFilter
