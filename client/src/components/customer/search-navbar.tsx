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
    <>
    <ul className={`${isHome && !scrolled ? 'flex items-center justify-center gap-4 text-white text-lg font-bold' : 'hidden'}`}>
      <li><a href={"/"} className="px-4 py-2 rounded hover:bg-primary-foreground cursor-pointer">Home</a></li>
      <li><a href={"/products"} className="px-4 py-2 rounded hover:bg-primary-foreground cursor-pointer">Shop</a></li>
      <li><a href={"/about-us"} className="px-4 py-2 rounded hover:bg-primary-foreground cursor-pointer">About Us</a></li>
    </ul>
    <InputGroup className={`${!isHome ? 'md:flex max-w-lg' : scrolled ? 'md:flex max-w-lg' : 'hidden'} hidden rounded-full bg-white border-0 overflow-hidden`}>
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
    </>
  )
}