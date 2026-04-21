export type ProductSummary = {
  id: number;
  category: {
    id: number;
    name: string;
  };
  name: string;
  is_published: boolean;
  tags: string[];
  slug: string;
  main_image_url: string;
  average_rating: number;
  review_count: number;
  sold_count: number;
  min_price: number;
  max_price: number;
  badge: string;
  sale_price?: number;
  sale_percentage?: number;
};

export type ProductDetails = {
  id: number;
  category: {
    id: number;
    name: string;
  };
  name: string;
  description: string;
  is_published: boolean;
  tags: string[];
  variants: Variant[];
  skus: SKU[];
  default_sku_id: number;
  slug: string;
  main_image_url: string;
  average_rating: number;
  review_count: number;
  sold_count: number;
  min_price: number;
  max_price: number;
  images: Image[];
};

export type Variant = {
  id: number;
  name: string;
  values: Value[];
};

export type Value = {
  id: number;
  value: string;
};

export type SKU = {
  id: number;
  sku_code: string;
  price: number;
  sale_price?: number | null;
  stock: number;
  min_stock: number;
  weight: number;
  status: "active" | "inactive" | "archived";
  variant_value_ids: number[];
  images: Image[];
};

export type Image = {
  id: number;
  image_url: string;
};

export type PreviewImage = {
  id?: number
  tempId: string
  file?: File
  image_url: string
}