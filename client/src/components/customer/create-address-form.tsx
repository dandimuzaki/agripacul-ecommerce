'use client'

import { useEffect, useState } from 'react';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '../ui/dialog';
import { Input } from '../ui/input';
import { Button } from '../ui/button';
import { Controller, useForm } from 'react-hook-form';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '../ui/select';
import { zodResolver } from '@hookform/resolvers/zod';
import { AddressFormValues, addressSchema } from '@/schemas/address.schema';
import { Field, FieldError, FieldGroup, FieldLabel } from '../ui/field';
import { InputGroup, InputGroupAddon, InputGroupText, InputGroupTextarea } from '../ui/input-group';
import { useProvince } from '@/hooks/address/useProvince';
import { useCreateAddress } from '@/hooks/address/useCreateAddress';
import { useRegency } from '@/hooks/address/useRegency';
import { useDistrict } from '@/hooks/address/useDistrict';
import { useSubdistrict } from '@/hooks/address/useSubdistrict';

const CreateAddressForm = () => {
  const [onOpen, setOnOpen] = useState(false)
  const { data: provinces } = useProvince();
  const { mutate } = useCreateAddress({
    onSuccess: () => setOnOpen(false)
  })

  const form = useForm<AddressFormValues>({
    resolver: zodResolver(addressSchema),
    defaultValues: {
      recipient_name: "",
      label: "",
      province_id: undefined,
      regency_id: undefined,
      district_id: undefined,
      subdistrict_id: undefined,
      detail_address: "",
      postal_code: "",
      phone_number: ""
    },
  })

  const provinceId = form.watch('province_id');
  const regencyId = form.watch('regency_id');
  const districtId = form.watch('district_id');

  const { data: regencies } = useRegency(provinceId as number)
  const { data: districts } = useDistrict(regencyId as number)
  const { data: subdistricts } = useSubdistrict(districtId as number)

  const onSubmitAddress = (data: AddressFormValues) => {
    mutate(data)
  }

  const handleCloseDialog = () => {
    setOnOpen(false)
    form.reset()
  }

  return (
    <Dialog open={onOpen} onOpenChange={setOnOpen}>
      <DialogTrigger asChild>
        <Button>Add New Address</Button>
      </DialogTrigger>
      <DialogContent className='overflow-y-auto'>
        <DialogHeader>
          <DialogTitle className='text-center'>
            Create Address
          </DialogTitle>
        </DialogHeader>
        <form
          id="create-address" onSubmit={form.handleSubmit(onSubmitAddress)}
        >
          <FieldGroup className='grid grid-cols-2 gap-y-4 gap-x-5'>
            <Controller
              name="recipient_name"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="edit-address-recipient-name">
                    Recipient Name
                  </FieldLabel>
                  <Input
                    {...field}
                    id="edit-address-recipient-name"
                    aria-invalid={fieldState.invalid}
                    placeholder="Fulan"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="label"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="edit-address-label">
                    Label
                  </FieldLabel>
                  <Input
                    {...field}
                    id="edit-address-label"
                    aria-invalid={fieldState.invalid}
                    placeholder="Home"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="phone_number"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="edit-address-phone-number">
                    Label
                  </FieldLabel>
                  <Input
                    {...field}
                    id="edit-address-phone-number"
                    aria-invalid={fieldState.invalid}
                    placeholder="081234567890"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="postal_code"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="edit-address-postal-code">
                    Label
                  </FieldLabel>
                  <Input
                    {...field}
                    id="edit-address-postal-code"
                    aria-invalid={fieldState.invalid}
                    placeholder="40514"
                  />
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="province_id"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel>Province</FieldLabel>

                  <Select
                    value={field.value?.toString()}
                    onValueChange={(value) => field.onChange(Number(value))}
                  >
                    <SelectTrigger aria-invalid={fieldState.invalid}>
                      <SelectValue placeholder="Select province" />
                    </SelectTrigger>

                    <SelectContent>
                      {provinces?.map((p) => (
                        <SelectItem
                          key={p.id}
                          value={p.id.toString()}
                        >
                          {p.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>

                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="regency_id"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel>Regency</FieldLabel>

                  <Select
                    value={field.value?.toString()}
                    onValueChange={(value) => field.onChange(Number(value))}
                  >
                    <SelectTrigger aria-invalid={fieldState.invalid}>
                      <SelectValue placeholder="Select regency" />
                    </SelectTrigger>

                    <SelectContent>
                      {regencies?.map((r) => (
                        <SelectItem
                          key={r.id}
                          value={r.id.toString()}
                        >
                          {r.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>

                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="district_id"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel>District</FieldLabel>

                  <Select
                    value={field.value?.toString()}
                    onValueChange={(value) => field.onChange(Number(value))}
                  >
                    <SelectTrigger aria-invalid={fieldState.invalid}>
                      <SelectValue placeholder="Select district" />
                    </SelectTrigger>

                    <SelectContent>
                      {districts?.map((d) => (
                        <SelectItem
                          key={d.id}
                          value={d.id.toString()}
                        >
                          {d.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>

                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <Controller
              name="subdistrict_id"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel>Subdistrict</FieldLabel>

                  <Select
                    value={field.value?.toString()}
                    onValueChange={(value) => field.onChange(Number(value))}
                  >
                    <SelectTrigger aria-invalid={fieldState.invalid}>
                      <SelectValue placeholder="Select subdistrict" />
                    </SelectTrigger>

                    <SelectContent>
                      {subdistricts?.map((s) => (
                        <SelectItem
                          key={s.id}
                          value={s.id.toString()}
                        >
                          {s.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>

                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />

            <div className="col-span-2">
            <Controller
              name="detail_address"
              control={form.control}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="edit-address-detail">
                    Detail Address
                  </FieldLabel>
                  <InputGroup>
                    <InputGroupTextarea
                      {...field}
                      id="edit-address-detail"
                      placeholder="I'm having an issue with the login button on mobile."
                      rows={6}
                      className="min-h-24 resize-none"
                      aria-invalid={fieldState.invalid}
                    />
                    <InputGroupAddon align="block-end">
                      <InputGroupText className="tabular-nums">
                        {field?.value?.length}/100 characters
                      </InputGroupText>
                    </InputGroupAddon>
                  </InputGroup>
                  {fieldState.invalid && (
                    <FieldError errors={[fieldState.error]} />
                  )}
                </Field>
              )}
            />
            </div>
          
          </FieldGroup>
          <div className="flex justify-center gap-3 mt-3 col-span-2">
            <Button
              className='bg-gray-200 hover:bg-gray-300 text-black' 
              type="button" onClick={handleCloseDialog}>
              Cancel
            </Button>
            <Button type='submit' form='create-address'>
              Save
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default CreateAddressForm;