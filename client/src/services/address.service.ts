import { api } from "@/lib/api";
import { AddressFormValues } from "@/schemas/address.schema";
import { Response } from "@/types/response";

export const addressService = {
  getAddressList(): Promise<Response> {
    return api.get(`/address`)
  },

  getAddressById(id: number): Promise<Response> {
    return api.get(`/address/${id}`)
  },

  createAddress(payload: AddressFormValues): Promise<Response> {
    return api.post("/address", payload)
  },

  getProvinceList(): Promise<Response> {
    return api.get(`/provinces`)
  },

  getRegencyList(provinceId: number): Promise<Response> {
    return api.get(`/regencies/${provinceId}`)
  },

  getDistrictList(regencyId: number): Promise<Response> {
    return api.get(`/districts/${regencyId}`)
  },

  getSubistrictList(districtId: number): Promise<Response> {
    return api.get(`/subdistricts/${districtId}`)
  },

  updateAddress(id: number, payload: AddressFormValues) {
    return api.put(`/address/${id}`, payload)
  },

  deleteAddress(id: number): Promise<Response> {
    return api.delete(`/address/${id}`)
  },

  setDefaultAddress(id: number): Promise<Response> {
    return api.put(`/address/${id}/default`)
  },
};