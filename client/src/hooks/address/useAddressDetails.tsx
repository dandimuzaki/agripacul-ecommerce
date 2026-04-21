"use client";

import { useQuery } from "@tanstack/react-query"
import { addressService } from "@/services/address.service"
import { addressKeys } from "../queries/addressKeys";
import { Address } from "@/types/address";
import { Response } from "@/types/response";

export const useAddressDetails = (id?: number) => {
  return useQuery<Response, Error, Address>({
    queryKey: addressKeys.detail(id!),
    queryFn: () => addressService.getAddressById(id!),
    select: (res) => res.data,
    enabled: !!id,
  })
}