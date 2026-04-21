'use client';

import AddressList from '@/components/customer/address-list'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useCheckoutForm } from '@/hooks/checkout/useCheckoutForm';
import { useAddressDetails } from '@/hooks/address/useAddressDetails';
import { capitalize } from '@/lib/formatText'
import { LocationPin } from '@mui/icons-material'

const AddressSection = () => {
  const { shippingAddressId, form } = useCheckoutForm()
  const { data: address } = useAddressDetails(shippingAddressId)

  return (
    <Card>
      <CardHeader>
        <CardTitle className='text-primary font-medium text-lg uppercase'>
          Delivery Address
        </CardTitle>
      </CardHeader>
      <CardContent className='flex gap-4'>
      <div className='space-y-2 flex-1'>
        {address ?
          <>
            <p className='font-bold ml-[-4px] flex items-center'><LocationPin className='text-[var(--primary)]'/>{address?.label} • {address?.recipient_name}</p>
            <p className='text-sm'>{address?.detail_address}, { capitalize(address?.subdistrict.name)}, {capitalize(address?.district.name)}, {capitalize(address?.regency.name)}, {capitalize(address.province.name)}</p>
          </>
          :
          <p>Please set your address</p>
        }
      </div>
      <div className='flex justify-center items-center'>
        {address ? (<AddressList form={form} shippingAddressId={shippingAddressId} buttonText='Change'/>) : (<AddressList form={form} shippingAddressId={shippingAddressId} buttonText='Create'/>)}
      </div>
      </CardContent>
    </Card>
  )
}

export default AddressSection
