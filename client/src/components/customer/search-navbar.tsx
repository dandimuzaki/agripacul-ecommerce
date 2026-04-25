"use client"

import { SearchIcon } from "lucide-react"
import { Button } from "../ui/button"
import { InputGroup, InputGroupInput } from "../ui/input-group"
import { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";

export default function SearchInput({isHome, scrolled}: {isHome: boolean, scrolled: boolean}) {
  const [keyword, setKeyword] = useState("");
  const searchParams = useSearchParams();
  const router = useRouter()

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.push(`/products?${params.toString()}`)
  }

  return (
    <InputGroup className={`${!isHome ? 'max-w-lg' : scrolled ? 'max-w-lg' : 'w-0'} hidden md:flex rounded-full bg-white border-0 overflow-hidden`}>
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
  )
}