import { User } from '@/types/user'
import { persist } from 'zustand/middleware'
import { create } from 'zustand'

type AuthState = { 
  user: User | null 
  setUser: (user: User) => void 
  logout: () => void 
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,

      setUser: (user: User) => set({ user }),

      logout: () => {
        localStorage.removeItem("accessToken")
        set({ user: null })
      },
    }),
    {
      name: 'auth-storage', // key in localStorage
    }
  )
)