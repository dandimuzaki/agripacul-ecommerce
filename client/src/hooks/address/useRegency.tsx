"use client";

import { useQuery } from "@tanstack/react-query"
import { addressService } from "@/services/address.service"
import { Response } from "@/types/response";
import { addressKeys } from "../queries/addressKeys";
import { AddressLocation } from "@/types/address";

export const useRegency = (id?: number) => {  
  return useQuery<Response, Error, AddressLocation[]>({
    queryKey: addressKeys.regency(id!),
    queryFn: () => addressService.getRegencyList(id!),
    select: (res) => res.data,
    enabled: !!id,
  })
}