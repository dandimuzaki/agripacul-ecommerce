"use client"

import AddressList from "@/components/customer/address-list"
import { Card, CardContent } from "@/components/ui/card"
import { useAddress } from "@/hooks/address/useAddress"
import { capitalize } from "@/lib/formatText"
import { LocationPin } from "@mui/icons-material"

export default function AddressSection() {
  const { defaultShippingAddress: address } = useAddress()

  return (
    <div className="space-y-2">
      <h2 className="text-2xl font-semibold uppercase text-primary">My Address</h2>
      <Card className="w-full">
        <CardContent className='space-y-2'>
          <p className="text-sm font-smeibold">Default Address</p>
        <div className='space-y-2 flex-1'>
          {address ?
            <>
              <p className='font-bold ml-[-4px] flex items-center'><LocationPin className='text-[var(--primary)]'/>{address?.label} • {address?.recipient_name}</p>
              <p className='text-sm'>{address?.detail_address}, { capitalize(address?.subdistrict.name)}, {capitalize(address?.district.name)}, {capitalize(address?.regency.name)}, {capitalize(address.province.name)}</p>
              <p className="font-semibold">{address?.phone_number}</p>
            </>
            :
            <p>Please set your address</p>
          }
        </div>
        <div className='flex justify-center items-center'>
          <AddressList buttonText="Manage Address"/>
        </div>
        </CardContent>
      </Card>
    </div>
  )
}