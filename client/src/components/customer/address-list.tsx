"use client";

import { useAddress } from '@/hooks/address/useAddress';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '../ui/dialog';
import AddressCard from './address-card';
import CreateAddressForm from './create-address-form';
import { CheckoutFormValuesTemp } from '@/schemas/checkout.schema';
import { UseFormReturn } from 'react-hook-form';
import { Dispatch, SetStateAction } from 'react';

const AddressList = ({openAddress, setOpenAddress, form, shippingAddressId}: {
  openAddress: boolean,
  setOpenAddress: Dispatch<SetStateAction<boolean>>,
  form?: UseFormReturn<CheckoutFormValuesTemp>,
  shippingAddressId?: number
}) => {
  const { data: addressList } = useAddress();

  return (
    <Dialog open={openAddress} onOpenChange={setOpenAddress}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle className='text-center'>Address List</DialogTitle>
        </DialogHeader>
        <div className='space-y-4 min-w-10'>
          <CreateAddressForm />
          <div className='space-y-4'>
            {addressList?.map((address, i) =>
              <AddressCard form={form} shippingAddressId={shippingAddressId} address={address} key={i} />
            )}
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default AddressList;