import { z } from "zod";
import { variantTypeSchema } from "./variant.schema";

export const editProductSchema = z.object({
  name: z.string().min(3, "Name is required"),
  description: z.string().min(10).optional(),
  category_id: z.number(),
});

export type EditProductFormValues = z.infer<typeof editProductSchema>;

export const productSchema = z.object({
  name: z.string().min(3, "Name is required"),
  description: z.string().min(10).optional(),
  category_id: z.number(),
  tags: z.array(z.string().min(1)),
  variants: z.array(variantTypeSchema)
});

export type ProductFormValues = z.infer<typeof productSchema>;

const MAX_FILE_SIZE = 2 * 1024 * 1024; // 2MB
const ACCEPTED_IMAGE_TYPES = ["image/jpeg", "image/jpg", "image/png", "image/webp"];

export const productMainImageSchema = z.object({
  image: z.file()
    .refine((file) => file.size <= MAX_FILE_SIZE, "file must be less than 2MB")
    .refine((file) => ACCEPTED_IMAGE_TYPES.includes(file.type),
      "Only .jpg, .jpeg, .png and .webp formats are supported")
})

export type ProductMainImageFormValues = z.input<typeof productMainImageSchema>;

export const productGallerySchema = z.object({
  images: z
    .array(z.instanceof(File)) // Ensure it's an array of File objects
    .min(1, "At least one image is required") // Minimum count
    .max(5, "You can upload up to 5 images") // Maximum count
    .refine(
      (files) => files.every((file) => file.size <= MAX_FILE_SIZE),
      "Each file must be less than 2MB"
    )
    .refine(
      (files) => files.every((file) => ACCEPTED_IMAGE_TYPES.includes(file.type)),
      "Only .jpg, .jpeg, .png and .webp formats are supported"
    ),
})

export type ProductGalleryFormValues = z.input<typeof productGallerySchema>;

export const browseProductSchema = z.object({
  search: z.coerce.string().optional(),
  category_id: z.number().optional(),
  min_price: z.number().optional(),
  max_price: z.number().optional(),
  page: z.number().optional(),
  limit: z.number().optional(),
  rating: z.number().optional(),
  sort_by: z.string().optional(),
  sort_order: z.string().optional(),
  sort: z.string().optional()
})

export type BrowseProductFormValues = z.input<typeof browseProductSchema>;