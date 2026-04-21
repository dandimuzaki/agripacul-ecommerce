"use client";

import { useQuery } from "@tanstack/react-query"
import { addressService } from "@/services/address.service"
import { Response } from "@/types/response";
import { addressKeys } from "../queries/addressKeys";
import { AddressLocation } from "@/types/address";

export const useSubdistrict = (id?: number) => {  
  return useQuery<Response, Error, AddressLocation[]>({
    queryKey: addressKeys.subdistrict(id!),
    queryFn: () => addressService.getSubistrictList(id!),
    select: (res) => res.data,
    enabled: !!id,
  })
}