'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useAddressDetails } from '@/hooks/address/useAddressDetails';
import { capitalize } from '@/lib/formatText'
import { LocationPin } from '@mui/icons-material'
import { Button } from '@/components/ui/button';
import { Dispatch, SetStateAction } from 'react';

const AddressSection = ({shippingAddressId, setOpenAddress}: {shippingAddressId?: number, setOpenAddress: Dispatch<SetStateAction<boolean>>}) => {
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
          <Button onClick={() => setOpenAddress(true)} type="button">{address ? 'Change' : 'Create'}</Button>
        </div>
      </CardContent>
    </Card>
  )
}

export default AddressSection
