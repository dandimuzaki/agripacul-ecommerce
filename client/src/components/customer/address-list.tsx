"use client";

import { useAddress } from '@/hooks/address/useAddress';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '../ui/dialog';
import { Button } from '../ui/button';
import AddressCard from './address-card';
import CreateAddressForm from './create-address-form';
import { CheckoutFormValuesTemp } from '@/schemas/checkout.schema';
import { UseFormReturn } from 'react-hook-form';

const AddressList = ({buttonText, form, shippingAddressId}: {
  buttonText: string,
  form?: UseFormReturn<CheckoutFormValuesTemp>,
  shippingAddressId?: number
}) => {
  const { data: addressList } = useAddress();

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className='text-sm px-3 py-1 cursor-pointer'>
          {buttonText}
        </Button>
      </DialogTrigger>
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