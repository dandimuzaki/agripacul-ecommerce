import { api } from "@/lib/api";
import { CategoryFormValues } from "@/schemas/category.schema";
import { Response } from "@/types/response";

export const categoryService = {
  getCategories(): Promise<Response> {
    return api.get('/categories')
  },

  getCategoryById(id: number): Promise<Response> {
    return api.get(`/admin/categories/${id}`)
  },

  createCategory(payload: CategoryFormValues) {
    const formData = new FormData()
    if (payload.name) formData.append("name", payload.name)
    if (payload.icon) formData.append("icon", payload.icon)

    return api.post(`/admin/categories`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
  },

  updateCategory(id: number, payload: CategoryFormValues) {
    const formData = new FormData()
    if (payload.name) formData.append("name", payload.name)
    if (payload.icon) formData.append("icon", payload.icon)

    return api.put(`/admin/categories/${id}`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
  },

  deleteCategory(id: number) {
    return api.delete(`/admin/categories/${id}`)
  },
};