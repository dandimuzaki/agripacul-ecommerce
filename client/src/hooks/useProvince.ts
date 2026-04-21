"use client";

import { addressService } from "@/services/address.service";
import { AddressLocation } from "@/types/address";
import { useEffect, useState } from "react";


export const useProvince = () => {
  const [data, setData] = useState<AddressLocation[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let isMounted = true;

    const fetchProvinces = async () => {
      setLoading(true);

      // simulate network delay
      await new Promise((resolve) => setTimeout(resolve, 2000));

      const res = await addressService.getProvinceList();

      if (isMounted) {
        setData(res);
        setLoading(false);
      }
    };

    fetchProvinces();

    return () => {
      isMounted = false;
    };
  }, []);

  return { provinces: data, loading };
};