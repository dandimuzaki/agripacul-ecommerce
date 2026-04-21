export interface BaseResponse<T = any> {
  data: T;
  message?: string;
  success: boolean;
}

export interface BaseListResponse<T = any> {
  data: T[];
  message?: string;
  success: boolean;
  meta: Meta
}

export interface Meta {
  page: number,
  per_page: number,
  total: number,
  total_pages: number
}

export interface PaginationRequest {
  page?: number;
  limit?: number;
  sort_by?: string;
  sort_order?: string;
}

export interface ErrorResponse {
  errors: string;
  message: string;
  success: boolean;
}