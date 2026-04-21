"use client";

import { Address } from '@/types/address';
import { Check } from 'lucide-react';
import { Button } from '../ui/button';
import EditAddressForm from './address-form';
import { useDeleteAddress } from '@/hooks/address/useDeleteAddress';
import { useSetDefaultAddress } from '@/hooks/address/useSetDefaultAddress';
import { UseFormReturn } from 'react-hook-form';
import { CheckoutFormValuesTemp } from '@/schemas/checkout.schema';

const AddressCard = ({ 
  address, form, shippingAddressId
}: { 
  address: Address,
  form?: UseFormReturn<CheckoutFormValuesTemp>,
  shippingAddressId?: number
}) => {
  const { id, is_default, recipient_name, label, province, regency, district, subdistrict, detail_address, phone_number } = address;
  const {mutate: onDelete} = useDeleteAddress()
  const handleDeleteAddress = () => {
    onDelete(id)
  }
  const {mutate: onDefault} = useSetDefaultAddress()
  const handleDefaultAddress = () => {
    onDefault(id)
  }

  const handleShippingAddress = () => {
    if (form) {
      form.setValue("shipping_address_id", id)
    }
  }

  return (
    <div className={`flex justify-center p-4 gap-4 rounded-lg ${is_default ? 'border border-primary bg-primary/20' : 'shadow-[0_0_4px_rgba(0,0,0,0.2)]'}`}>
      <div className='flex-1'>
        <div className='flex gap-2 items-center'>
          <p className='font-semibold text-sm'>{label}</p>
          {address.is_default ? <p className='text-sm px-2 py-1 rounded bg-gray-200'>Default</p> : ''}
        </div>
        <p className='text-base font-bold'>{recipient_name}</p>
        <p className='text-sm'>{phone_number}</p>
        <p className='text-sm'>{detail_address}</p>
        <p className='text-sm'>{subdistrict.name}, {district.name}, {regency.name}, {province.name}</p>
        <div className='mt-2 text-sm flex gap-4'>
          <EditAddressForm id={address.id}/>
          {!is_default && <Button className='bg-gray-200 hover:bg-gray-300 text-black' onClick={handleDefaultAddress}>Set as Default Address</Button>}
          <Button variant="destructive" onClick={handleDeleteAddress}>Delete</Button>
        </div>
      </div>
      {form && <div className='flex justify-center items-center'>
        {id == shippingAddressId ?
          <Check className='text-primary'/>
          :
          <Button
            type='button'
            onClick={handleShippingAddress}
            className='text-sm font-bold rounded py-1 px-3 cursor-pointer'>
          Choose
          </Button>}
      </div>}
    </div>
  );
};

export default AddressCard;