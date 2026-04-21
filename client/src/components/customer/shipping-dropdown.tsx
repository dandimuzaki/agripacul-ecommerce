'use client';

import { Controller, UseFormReturn } from 'react-hook-form';
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from '@/components/ui/select';
import { Field } from '../ui/field';
import { formatRupiah } from '@/lib/formatCurrency';
import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { CheckoutFormValuesTemp } from '@/schemas/checkout.schema';
import { ShippingOptionTemp } from '@/types/checkout';

const ShippingDropdown = ({form, options}: {form: UseFormReturn<CheckoutFormValuesTemp>, options: ShippingOptionTemp[]}) => {
  return (
    <Card>
      <CardHeader>
        <CardTitle className='text-primary font-medium text-lg uppercase'>
          Shipping
        </CardTitle>
      </CardHeader>
      <CardContent>
    <Controller
      name="selected_shipping_option_id"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <Select
            value={String(field.value) ?? ""}
            onValueChange={(value) => {
              field.onChange(value);
            }}
            >
            <SelectTrigger className="text-left py-6 w-full justify-between rounded-lg">
              <SelectValue placeholder="Choose shipping option" />
            </SelectTrigger>

            <SelectContent>
              {options?.map((shipping) => (
                <SelectItem key={shipping.id} value={String(shipping.id)}>
                  <div className="space-y-2 w-full">
                    <span className="flex justify-between gap-2">
                      <span>{shipping.name}</span>
                      <span className="font-bold">{formatRupiah(shipping.cost)}</span>
                    </span>
                    <span className="text-sm text-gray-500">
                      Est. {shipping.etd}s
                    </span>
                  </div>
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </Field>
      )}
    />
    </CardContent>
    </Card>
  );
};

export default ShippingDropdown;