import { api } from "@/lib/api";
import { ProfileFormValues } from "@/schemas/profile.schema";
import { Response } from "@/types/response";

export const profileService = {
  getProfile(): Promise<Response> {
    return api.get('/profile')
  },

  updateProfile(payload: ProfileFormValues) {
    const formData = new FormData()
    if (payload.full_name) formData.append("full_name", payload.full_name)
    if (payload.phone_number) formData.append("phone_number", payload.phone_number)
    if (payload.date_of_birth) {
      const formatted = (payload.date_of_birth as Date)
        .toISOString()
        .split("T")[0]

      formData.append("date_of_birth", formatted)
    }
    if (payload.profile_image) formData.append("profile_image", payload.profile_image)

    return api.put(`/profile`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
  },
};