"use client";

import { useQuery } from "@tanstack/react-query"
import { addressService } from "@/services/address.service"
import { Response } from "@/types/response";
import { Address } from "@/types/address";
import { addressKeys } from "../queries/addressKeys";
import { useMemo } from "react";

export const useAddress = () => {  
  const query = useQuery<Response, Error, Address[]>({
    queryKey: addressKeys.all,
    queryFn: () => addressService.getAddressList(),
    select: (res) => res.data
  })

  const defaultShippingAddress: Address | null = useMemo(() => {
    if (!query.data) return null;

    return query.data.find(a => a.is_default) ?? null;
  }, [query.data]);

  return { ...query, defaultShippingAddress }
}