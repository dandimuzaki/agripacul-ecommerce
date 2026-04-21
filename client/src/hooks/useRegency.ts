"use client";

import { addressService } from "@/services/address.service";
import { AddressLocation } from "@/types/address";
import { useEffect, useState } from "react";


export const useRegency = (provinceId?: number) => {
  const [data, setData] = useState<AddressLocation[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let isMounted = true;

    const fetchRegencies = async (provinceId: number) => {
      setLoading(true);

      // simulate network delay
      await new Promise((resolve) => setTimeout(resolve, 2000));

      const res = await addressService.getRegencyList(provinceId);

      if (isMounted) {
        setData(res);
        setLoading(false);
      }
    };

    if (provinceId) fetchRegencies(provinceId);

    return () => {
      isMounted = false;
    };
  }, [provinceId]);

  return { regencies: data, loading };
};