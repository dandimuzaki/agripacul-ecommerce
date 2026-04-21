import axios from "axios"

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem("accessToken")

  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

api.interceptors.response.use(
  (response) => {
    // Handle success=false from backend
    if (response && typeof response === "object" && "success" in response) {
      if (!response.success) {
        return Promise.reject(new Error(response.data.message || "Something went wrong"))
      }

      return response.data // 👈 unwrap automatically
    }

    return response.data // fallback (if not wrapped)
  },
  (error) => {
    // Handle 401 globally
    if (error.response?.status === 401) {
      localStorage.removeItem("accessToken")
    }

    // Normalize error message
    const message =
      error.response?.data?.message ||
      error.message ||
      "Unexpected error"

    return Promise.reject(new Error(message))
  }
)