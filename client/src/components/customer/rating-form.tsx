"use client"
import { CreateReviewFormValues } from '@/schemas/review.schema';
import { Star } from '@mui/icons-material';
import { useState } from 'react';
import { Controller, UseFormReturn } from 'react-hook-form';
import { Field, FieldLabel } from '../ui/field';

const RatingForm = ({
  index,
  form,
}: {
  index: number;
  form: UseFormReturn<CreateReviewFormValues>;
}) => {
  const [hover, setHover] = useState(0);

  return (
    <Controller
      name={`reviews.${index}.rating`}
      control={form.control}
      render={({ field, fieldState }) => {
        const rating = field.value || 0;

        return (
          <Field data-invalid={fieldState.invalid}>
            <FieldLabel>Rating Product</FieldLabel>
              <div className="flex gap-1">
                {[...Array(5)].map((_, i) => {
                  const starValue = i + 1;

                  return (
                    <span
                      key={starValue}
                      onClick={() => field.onChange(starValue)}
                      onMouseEnter={() => setHover(starValue)}
                      onMouseLeave={() => setHover(0)}
                      className="cursor-pointer text-xl md:text-3xl/5 select-none"
                      style={{
                        color:
                          starValue <= (hover || rating)
                            ? 'gold'
                            : 'lightgray',
                      }}
                    >
                      <Star fontSize='inherit' />
                    </span>
                  );
                })}
              </div>
          </Field>
        );
      }}
    />
  );
};

export default RatingForm;