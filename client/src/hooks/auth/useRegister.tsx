"use client"

import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { RegisterFormValues } from "@/schemas/auth.schema"
import { authService } from "@/services/auth.service"
import { useAuthStore } from '@/store/useAuthStore'
import { useRouter } from 'next/navigation'
import { authKeys } from '../queries/authKeys'

export default function useRegister() {
  const queryClient = useQueryClient()
  const router = useRouter()
  const loginStore = useAuthStore((state) => state.login)

  return useMutation({
    mutationFn: (data: RegisterFormValues) => authService.register(data),

    onSuccess: (res) => {
      if (res.success) {
        loginStore(res.data.user, res.data.token)
        queryClient.setQueryData(authKeys.all, res.data.user)
      
      if (res.data.user.role === "customer") {
        router.push("/")
      }
      }
    },
  })
}