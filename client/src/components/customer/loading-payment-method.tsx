"use client";

import { RadioGroup, RadioGroupItem } from '../ui/radio-group';
import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { Skeleton } from '../ui/skeleton';

const LoadingPaymentMethodList = () => {
  const paymentMethodList = Array(7).fill(1)

  return (
    <Card>
      <CardHeader>
        <CardTitle className='text-primary font-medium text-lg uppercase'>
          <Skeleton className='h-10 w-24' />
        </CardTitle>
      </CardHeader>
      <CardContent>
    <RadioGroup>
      {paymentMethodList.map((payment, index) => (
        <div key={index} className={`flex items-center justify-between space-x-3 p-2 border rounded`}>
          <Skeleton className='h-8 w-16' />
          <RadioGroupItem value={String(index)} id={String(index)} />
        </div>
      ))}
    </RadioGroup>
    </CardContent>
    </Card>
  );
};

export default LoadingPaymentMethodList;