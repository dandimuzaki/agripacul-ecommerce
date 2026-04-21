import { api } from "@/lib/api";
import { BrowseProductFormValues, ProductFormValues, ProductGalleryFormValues, ProductMainImageFormValues } from "@/schemas/product.schema";
import { SKUFormValues } from "@/schemas/sku.schema";
import { Response } from "@/types/response";

export const productService = {
  getProducts(params: BrowseProductFormValues): Promise<Response> {
    return api.get('/products', { params })
  },

  getProductsByAdmin(params: BrowseProductFormValues): Promise<Response> {
    return api.get(`/admin/products`, { params })
  },

  getProductBySlug(slug: string): Promise<Response> {
    return api.get(`/products/details/${slug}`)
  },

  getProductById(id: number): Promise<Response> {
    return api.get(`/admin/products/${id}`)
  },

  getSKUByProductId(id: number): Promise<Response> {
    return api.get(`/admin/products/${id}/sku`)
  },

  createProduct(payload: ProductFormValues) {
    return api.post("/admin/products", payload)
  },

  updateProduct(id: number, payload: ProductFormValues) {
    return api.put(`/admin/products/${id}`, payload)
  },

  updateSKU(productId: number, payload: SKUFormValues) {
    return api.put(`/admin/products/${productId}/sku`, payload.skus)
  },

  updatePublishProduct(id: number, isPublished: boolean) {
    return api.put(`/admin/products/${id}/publish`, {is_published: isPublished})
  },

  deleteProduct(id: number) {
    return api.delete(`/admin/products/${id}`)
  },

  updateProductMainImage(id: number, payload: ProductMainImageFormValues) {
    const formData = new FormData()
    formData.append("image", payload.image)

    return api.post(`/admin/products/${id}`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
  },

  updateProductGallery(id: number, payload: ProductGalleryFormValues) {
    const formData = new FormData()

    payload.images.forEach((file) => {
      formData.append("images", file)
    })

    return api.post(`/admin/products/${id}/gallery`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
  },

  deleteProductImage(productId: number, imageId: number) {
    return api.delete(`/admin/products/${productId}/gallery/${imageId}`)
  },
};