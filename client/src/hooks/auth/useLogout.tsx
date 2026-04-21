"use client"

import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { authService } from "@/services/auth.service"
import { useAuthStore } from '@/store/useAuthStore'
import { useRouter } from 'next/navigation'
import { authKeys } from '../queries/authKeys'

export default function useLogout() {
  const queryClient = useQueryClient()
  const setLogout = useAuthStore((state) => state.logout)
  const router = useRouter()

  return useMutation({
    mutationFn: () => authService.logout(),

    onSuccess: (res) => {
      if (res.success) {
        localStorage.removeItem("accessToken")
        setLogout()
        queryClient.invalidateQueries({queryKey: authKeys.all})
      
        router.push("/")
      }
    },
  })
}