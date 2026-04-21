"use client";

import { useQuery } from "@tanstack/react-query"
import { categoryService } from "@/services/category.service"
import { categoryKeys } from "../queries/categoryKeys";
import { Category } from "@/types/category";
import { Response } from "@/types/response";

export const useCategoryById = (id: number) => {  
  return useQuery<Response, Error, Category>({
    queryKey: categoryKeys.detail(id),
    queryFn: () => categoryService.getCategoryById(id),
    select: (res) => res.data
  })
}