"use client"
import { Input } from '../ui/input'
import { Field, FieldLabel } from '../ui/field'
import { Card, CardContent, CardFooter } from '../ui/card'
import { Textarea } from '../ui/textarea'
import { Audiotrack, Email, Instagram, LinkedIn, WhatsApp, X } from '@mui/icons-material'
import { Button } from '../ui/button'
import { Controller } from 'react-hook-form'
import { useSendMessage } from '@/hooks/company/useSendMessage'
import { ContactFormValues } from '@/schemas/contact.schema'
import { Spinner } from '../ui/spinner'

const Contact = () => {
  const {form, mutate, isPending} = useSendMessage()
  const onSendMessage = (data: ContactFormValues) => {
    mutate(data)
  }
  return (
    <>
      <div className='grid md:grid-cols-[3fr_2fr] gap-4 md:gap-6'>
        <Card>
          <CardContent className=''>
            <form className='grid grid-cols-2 gap-2 md:gap-4' id='contact-form' onSubmit={form.handleSubmit(onSendMessage)}>
              <Controller
                name={`first_name`}
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor='first-name'>First Name</FieldLabel>
                    <Input
                      className='border-none bg-gray-100'
                      id="first-name"
                      type="text"
                      onChange={(e) => {
                        field.onChange(e)
                      }}
                    />
                  </Field>
                )}
              />
              <Controller
                name={`last_name`}
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor='last-name'>Last Name</FieldLabel>
                    <Input
                      className='border-none bg-gray-100'
                      id="last-name"
                      type="text"
                      onChange={(e) => {
                        field.onChange(e)
                      }}
                    />
                  </Field>
                )}
              />
              <Controller
                name={`email`}
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor='email'>Email</FieldLabel>
                    <Input
                      className='border-none bg-gray-100'
                      id="email"
                      type="email"
                      onChange={(e) => {
                        field.onChange(e)
                      }}
                    />
                  </Field>
                )}
              />
              <Controller
                name={`phone_number`}
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor='phone-number'>Phone Number</FieldLabel>
                    <Input
                      className='border-none bg-gray-100'
                      id="phone-number"
                      type="phone"
                      onChange={(e) => {
                        field.onChange(e)
                      }}
                    />
                  </Field>
                )}
              />
              <div className='col-span-2'>
                <Controller
                  name={`subject`}
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor='subject'>Subject</FieldLabel>
                      <Input
                        className='border-none bg-gray-100'
                        id="subject"
                        type="text"
                        onChange={(e) => {
                          field.onChange(e)
                        }}
                      />
                    </Field>
                  )}
                />
              </div>
              <div className='col-span-2'>
                <Controller
                  name={`body`}
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor='body'>Message</FieldLabel>
                      <Textarea
                        {...field}
                        id="body"
                        onChange={(e) => {
                          field.onChange(e)
                        }}
                      ></Textarea>
                    </Field>
                  )}
                />
              </div>
            </form>
          </CardContent>
          <CardFooter className='flex items-center justify-center mt-4'>
            <Button type='submit' form='contact-form' className=''>{isPending ? <><Spinner/> Sending message...</> : <>Send Message</>}</Button>
          </CardFooter>
        </Card>
        <div className='bg-primary text-white p-4 md:p-6 flex flex-col justify-between gap-4 rounded-lg h-full'>
          <div className='space-y-1'>
            <p className='font-semibold text-base md:text-lg'>Address</p>
            <p className='text-sm/5 md:text-base/5'>
              Jl. Cisintok Kadumulya, Cihanjuang, Kec. Parongpong, Kabupaten Bandung Barat
            </p>
          </div>
          <div className=''>
            <p className='font-semibold text-base md:text-lg mb-2'>Contact Us</p>
            <a className='flex gap-2 cursor-pointer items-center mb-1'><WhatsApp />+62 853-2409-1088</a>
            <a className='flex gap-2 cursor-pointer items-center'><Email />agripacul@itb.lpik.org</a>
          </div>
          <div className='space-y-1'>
            <p className='font-semibold text-base md:text-lg'>Working Hours</p>
            <p className='text-sm/5 md:text-base/5'>
              Monday - Friday : 09.00 - 17.00
            </p>
          </div>
          <div className='space-y-1'>
            <p className='font-semibold text-base md:text-lg'>Stay Connected</p>
            <div className='flex gap-2'>
              <Button className="bg-white text-primary hover:text-white aspect-square rounded-full h-full" ><Instagram fontSize='small'/></Button>
              <Button className="bg-white text-primary hover:text-white aspect-square rounded-full h-full"><X fontSize='small'/></Button>
              <Button className="bg-white text-primary hover:text-white aspect-square rounded-full h-full"><Audiotrack fontSize='small'/></Button>
              <Button className="bg-white text-primary hover:text-white aspect-square rounded-full h-full"><LinkedIn fontSize='small'/></Button>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default Contact
