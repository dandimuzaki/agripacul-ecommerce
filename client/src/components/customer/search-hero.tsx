"use client"

import { SearchIcon } from "lucide-react"
import { Button } from "../ui/button"
import { InputGroup, InputGroupInput } from "../ui/input-group"
import { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import CategoryFilterDropdown from "./category-dropdown";

export default function SearchHero() {
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
      <form
        onSubmit={(e) => {
          e.preventDefault();
          onSubmit(keyword);
        }}
      >
        <InputGroup className={`p-2 rounded-full bg-white border-0 overflow-hidden`}>
          <CategoryFilterDropdown/>
          <InputGroupInput
            placeholder="Search products..."
            value={keyword}
            onChange={(e) => setKeyword(e.target.value)}
          />
          <Button variant="searchIcon" type='button'>
            <SearchIcon/>
          </Button>
        </InputGroup>
      </form>
    </>
  )
}