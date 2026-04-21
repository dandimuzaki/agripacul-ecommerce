export interface AddressSnapshot {
  id: number,
  customer_id: number,
  recipient_name: string,
  label: string,
  province: string,
  regency: string,
  district: string,
  subdistrict: string,
  postal_code: string,
  detail_address: string,
  is_default: boolean
}

export interface AddressLocation {
  id: number,
  name: string,
  raja_ongkir_id?: number | null
}

export interface Address {
  id: number,
  customer_id: number,
  recipient_name: string,
  phone_number: string,
  label: string,
  province: AddressLocation,
  regency: AddressLocation,
  district: AddressLocation,
  subdistrict: AddressLocation,
  postal_code: string,
  detail_address: string,
  is_default: boolean 
}