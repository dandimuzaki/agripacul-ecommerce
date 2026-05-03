"use client";

import { RadioGroup, RadioGroupItem } from '../ui/radio-group';
import { Controller } from 'react-hook-form';
import { Field } from '../ui/field';
import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { useCheckoutForm } from '@/hooks/checkout/useCheckoutForm';
import { usePaymentMethods } from '@/hooks/payment/usePaymentMethods';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '../ui/accordion';

const PaymentMethodList = () => {
  const { data: payment } = usePaymentMethods()
  const { form } = useCheckoutForm()

  return (
    <Card>
      <CardHeader>
        <CardTitle className='text-primary font-medium text-lg uppercase'>
          Payment Method
        </CardTitle>
      </CardHeader>
      <CardContent>
        <Controller
          name="selected_payment_method_id"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <RadioGroup
                value={field.value?.toString()}
                onValueChange={(value) => {
                  field.onChange(Number(value))
                }}
              >
              <Accordion
                type="single"
                collapsible
                defaultValue="Bank Transfer"
                className="max-w-lg"
              >
                {payment?.map(p => <AccordionItem key={p.id} value={p.name}>
                  <AccordionTrigger className='text-base'>{p.name}</AccordionTrigger>
                  <AccordionContent className='space-y-2'>
                    {p.methods?.map((payment) => (
                      <div key={payment.id} className={`flex items-center justify-between space-x-3 p-2 border rounded ${(field.value == payment.id) ? 'bg-primary/20' : ''}`}>
                        <label htmlFor={String(payment.id)} className="flex-1 cursor-pointer text-sm flex items-center justify-between gap-2">
                          <div className='flex gap-2 items-center'>
                          <div className='flex justify-center h-8 w-16 bg-white items-center'>
                            <img alt={payment.name} src={payment.icon_url} className='h-fit w-fit object-fit'/>
                          </div>
                          {payment.name}
                          </div>
                        <RadioGroupItem value={String(payment.id)} id={String(payment.id)} />
                        </label>
                      </div>
                    ))}
                  </AccordionContent>
                </AccordionItem>)}
              </Accordion>
              </RadioGroup>
            </Field>
          )}
        />
      </CardContent>
    </Card>
  );
};

export default PaymentMethodList;