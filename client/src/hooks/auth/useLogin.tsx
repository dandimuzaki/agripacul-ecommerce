"use client"

import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { LoginFormValues } from "@/schemas/auth.schema"
import { authService } from "@/services/auth.service"
import { useAuthStore } from '@/store/useAuthStore'
import { useRouter } from 'next/navigation'
import { authKeys } from '../queries/authKeys'

export default function useLogin() {
  const queryClient = useQueryClient()
  const setUser = useAuthStore((state) => state.setUser)
  const router = useRouter()

  return useMutation({
    mutationFn: (data: LoginFormValues) => authService.login(data),

    onSuccess: (res) => {
      if (res.success) {
        localStorage.setItem("accessToken", res.data.token)
        setUser(res.data.user)
        queryClient.setQueryData(authKeys.all, res.data.user)
      
      if (res.data.user.role === "admin") {
        router.push("/admin/product")
      } else {
        router.push("/")
      }
      }
    },
  })
}