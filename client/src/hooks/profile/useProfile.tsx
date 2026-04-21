"use client";

import { useQuery } from "@tanstack/react-query"
import { profileKeys } from "../queries/profileKeys";
import { profileService } from "@/services/profile.service";
import { Response } from "@/types/response";
import { Profile } from "@/types/profile";

export const useProfile = () => {  
  return useQuery<Response, Error, Profile>({
    queryKey: profileKeys.all,
    queryFn: () => profileService.getProfile(),
    select: (res) => res.data
  })
}