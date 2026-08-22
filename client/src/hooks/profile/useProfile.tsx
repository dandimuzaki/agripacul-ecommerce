"use client";

import { useQuery } from "@tanstack/react-query"
import { profileKeys } from "../queries/profileKeys";
import { profileService } from "@/services/profile.service";
import { Response } from "@/types/response";
import { Profile } from "@/types/profile";
import { useAuthStore } from "@/store/useAuthStore";

export const useProfile = () => {  
  const token = useAuthStore(state => state.token);

  return useQuery<Response, Error, Profile>({
    queryKey: profileKeys.all,
    queryFn: () => profileService.getProfile(),
    select: (res) => res.data,
    enabled: !!token
  })
}