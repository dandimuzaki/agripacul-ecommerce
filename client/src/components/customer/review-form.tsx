import React, { useEffect, useState } from 'react';
import { Dialog, DialogClose, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '../ui/dialog';
import { useCreateReview } from '@/hooks/review/useCreateReview';
import { useOrderDetails } from '@/hooks/order/useOrderDetails';
import { CreateReviewFormValues, createReviewSchema } from '@/schemas/review.schema';
import { Controller, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import RatingForm from './rating-form';
import { Button } from '../ui/button';
import { Field, FieldError, FieldLabel } from '../ui/field';
import { Textarea } from '../ui/textarea';
import Image from 'next/image';
import AfterReview from './after-review';

const ReviewForm = ({id}: {id: number}) => {
  const [ openReview, setOpenReview ] = useState(false)
  const [ openAfterReview, setOpenAfterReview ] = useState(false)
  const {mutate} = useCreateReview(setOpenAfterReview)
  const { data: order } = useOrderDetails(id)
  const form = useForm<CreateReviewFormValues>({
    resolver: zodResolver(createReviewSchema),
  })

  useEffect(() => {
    if (order) {
      form.reset({
        order_id: order.id,
        reviews: order.items.map((item) => ({
          sku_id: item.sku_id,
          rating: 5,
          comment: ""
        }))
      });
    }
  }, [order]);

  const onSubmitReview = (data: CreateReviewFormValues) => {
    mutate(data)
  };

  const closeAfterReview = () => {
    setOpenAfterReview(prev => !prev)
    setOpenReview(prev => !prev)
  }

  return (
    <>
    <Dialog open={openReview} onOpenChange={setOpenReview}>
      <DialogTrigger asChild>
        <Button>Rate Now</Button>
      </DialogTrigger>
      <DialogContent className='overflow-y-auto overflow-x-hidden'>
        <DialogHeader>
          <DialogTitle>
            Share Your Review
          </DialogTitle>
        </DialogHeader>
        <form className='space-y-4' id="review-form" onSubmit={form.handleSubmit(onSubmitReview)}>
          <div className='grid gap-4'>
            {order?.items?.map((item, i) => (
              <div key={item.sku_id} className='space-y-2 bg-white p-4 rounded-lg'>
                <div className='flex gap-2 items-center'>
                  <Image className='h-8 w-8 rounded' src={item.main_image_url ?? "/loading.png"} alt={item.name} height={100} width={100}/>
                  <p className='truncate w-full'>{item.name}</p>
                </div>
                <RatingForm index={i} form={form}/>
                <Controller
                  name={`reviews.${i}.comment`}
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor='review'>Review</FieldLabel>
                      <Textarea
                        className="whitespace-wrap"
                        {...field}
                        id='review'
                        placeholder="What did you love about it"
                        onChange={(e) => {
                          field.onChange(e)
                        }}
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />
              </div>
            ))}
          </div>
        </form>
        <DialogFooter className='flex items-center gap-2'>
          <DialogClose className='text-sm rounded-lg px-4 py-2 bg-gray-300 cursor-pointer hover:bg-gray-400'>
            Cancel
          </DialogClose>
          <Button
            type='submit'
            form='review-form'
          >
            Submit
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
    <AfterReview closeAfterReview={closeAfterReview} openAfterReview={openAfterReview} setOpenAfterReview={setOpenAfterReview} />
    </>
  );
};

export default ReviewForm;