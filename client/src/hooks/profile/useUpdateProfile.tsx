import {
  useQueryClient,
  useMutation,
} from '@tanstack/react-query'
import { ProfileFormValues } from "@/schemas/profile.schema"
import { profileService } from "@/services/profile.service"
import { profileKeys } from '../queries/profileKeys'
import { toast } from 'sonner'
import { Dispatch, SetStateAction } from 'react'

export const useUpdateProfile = ({setOnEdit}: {setOnEdit: Dispatch<SetStateAction<boolean>>}) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: ProfileFormValues) => profileService.updateProfile(payload),

    onError: (_err) => {
      toast.error(_err.message || "Failed to update profile")
    },

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: profileKeys.all
      })

      toast.success("Profile updated successfully")

      setOnEdit(false)
    }
  })
}