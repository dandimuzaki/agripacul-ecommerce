"use client";

import { addressService } from "@/services/address.service";
import { AddressLocation } from "@/types/address";
import { useEffect, useState } from "react";


export const useDistrict = (regencyId?: number) => {
  const [data, setData] = useState<AddressLocation[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let isMounted = true;

    const fetchDistricts = async (regencyId: number) => {
      setLoading(true);

      // simulate network delay
      await new Promise((resolve) => setTimeout(resolve, 2000));

      const res = await addressService.getDistrictList(regencyId);

      if (isMounted) {
        setData(res);
        setLoading(false);
      }
    };

    if (regencyId) fetchDistricts(regencyId);

    return () => {
      isMounted = false;
    };
  }, [regencyId]);

  return { districts: data, loading };
};