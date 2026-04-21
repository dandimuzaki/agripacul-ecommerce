"use client"

import { useSearchParams, useRouter } from "next/navigation"
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "../ui/pagination"

export function ReusablePagination({ currentPage, totalPages }: { currentPage: number, totalPages: number }) {
  const searchParams = useSearchParams()
  const router = useRouter()

  const goToPage = (page: number) => {
    const params = new URLSearchParams(searchParams.toString())

    params.set("page", String(page))

    router.push(`?${params.toString()}`)
  }

  return (
    <Pagination>
      <PaginationContent>

        {/* Previous */}
        <PaginationItem>
          <PaginationPrevious
            onClick={() => currentPage > 1 && goToPage(currentPage - 1)}
          />
        </PaginationItem>

        {/* Pages */}
        {[...Array(totalPages)].map((_, i) => {
          const page = i + 1

          return (
            <PaginationItem key={page}>
              <PaginationLink
                isActive={page === currentPage}
                onClick={() => goToPage(page)}
              >
                {page}
              </PaginationLink>
            </PaginationItem>
          )
        })}

        {/* Next */}
        <PaginationItem>
          <PaginationNext
            onClick={() =>
              currentPage < totalPages && goToPage(currentPage + 1)
            }
          />
        </PaginationItem>

      </PaginationContent>
    </Pagination>
  )
}