"use client";

import { useQuery } from "@tanstack/react-query"
import { addressService } from "@/services/address.service"
import { Response } from "@/types/response";
import { addressKeys } from "../queries/addressKeys";
import { AddressLocation } from "@/types/address";

export const useProvince = () => {  
  return useQuery<Response, Error, AddressLocation[]>({
    queryKey: addressKeys.province,
    queryFn: () => addressService.getProvinceList(),
    select: (res) => res.data
  })
}