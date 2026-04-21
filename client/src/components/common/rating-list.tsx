"use client";

import { useProductFilter } from "@/hooks/product/useProductFilter";
import { Star, StarBorder } from "@mui/icons-material";
import { Controller } from "react-hook-form";
import { Field, FieldLabel } from "../ui/field";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import { useRouter, useSearchParams } from "next/navigation";

export const RatingList = () => {
  const rating = [1,2,3,4,5]
  const {form} = useProductFilter()
  const searchParams = useSearchParams()
  const router = useRouter()

  return (
    <Controller
      name="rating"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <FieldLabel className="text-md border-b border-gray-400 pb-2">
            Rating
          </FieldLabel>
          <RadioGroup
            value={field.value?.toString()}
            onValueChange={(value) => {
              field.onChange(Number(value))
              const params = new URLSearchParams(searchParams.toString())
              params.set("rating", value)
              router.replace(`/products?${params.toString()}`)
            }}
            className='gap-0'
          >
            {rating.map((r) => (
              <label key={r} className={`flex items-center gap-2 p-2 rounded ${(field.value == r) ? 'bg-primary/20' : ''}`}>
                <RadioGroupItem className="border-gray-400" value={String(r)} id={String(r)} />
                <div className="flex items-center text-yellow-500">
                  {Array.from({ length: r }).map((_, i) => (
                    <Star fontSize="small" key={i} />
                  ))}
                  {Array.from({ length: 5-r }).map((_, i) => (
                    <StarBorder fontSize="small" key={i} />
                  ))}
                </div>
              </label>
            ))}
          </RadioGroup>
        </Field>
      )}
    />
  );
};