"use client";

import { addressService } from "@/services/address.service";
import { AddressLocation } from "@/types/address";
import { useEffect, useState } from "react";


export const useSubdistrict = (districtId?: number) => {
  const [data, setData] = useState<AddressLocation[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let isMounted = true;

    const fetchSubdistricts = async (districtId: number) => {
      setLoading(true);

      // simulate network delay
      await new Promise((resolve) => setTimeout(resolve, 2000));

      const res = await addressService.getSubistrictList(districtId);

      if (isMounted) {
        setData(res);
        setLoading(false);
      }
    };

    if (districtId) fetchSubdistricts(districtId);

    return () => {
      isMounted = false;
    };
  }, [districtId]);

  return { subdistricts: data, loading };
};