"use client";

import { useQuery } from "@tanstack/react-query"
import { categoryService } from "@/services/category.service"
import { categoryKeys } from "../queries/categoryKeys";
import { Response } from "@/types/response";

export const useCategories = () => {  
  return useQuery<Response>({
    queryKey: categoryKeys.all,
    queryFn: () => categoryService.getCategories(),
  })
}