'use client';

import { VariantValueEdit } from "@/types/variant";
import { VariantAction } from "./variant-field";
import { Controller, UseFormReturn } from "react-hook-form";
import { Field, FieldError, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Check, Delete, Edit } from "@mui/icons-material";
import { ProductFormValues } from "@/schemas/product.schema";

type Props = {
  indexType: number
  indexValue: number
  value: VariantValueEdit
  typeTempId: string
  form: UseFormReturn<ProductFormValues>
  dispatch: React.ActionDispatch<[action: VariantAction]>
}

export function VariantValueRow({ indexType, indexValue, value, typeTempId, form, dispatch }: Props) {
  return (
    <div>
      <Controller
        name={`variants.${indexType}.values.${indexValue}.value`}
        control={form.control}
        render={({ field, fieldState }) => (
          <div className="flex gap-2 w-fit">
            <Field data-invalid={fieldState.invalid} className="grid grid-cols-[auto_auto] flex-1 gap-2">
              <FieldLabel className="whitespace-nowrap">Value : </FieldLabel>

              <Input
                {...field}
                disabled={!value.isEditing}
                state={value.status}
                placeholder="Size"
                onChange={(e) => {
                  field.onChange(e)
                  dispatch({
                    type: "UPDATE_VALUE",
                    payload: {
                      typeTempId: typeTempId,
                      valueTempId: value.tempId,
                      value: e.target.value,
                    },
                  })
                }}
              />

              <div></div>

              {fieldState.invalid && (
                <FieldError errors={[fieldState.error]} />
              )}
            </Field>
            
            <div className="flex gap-2">
              <Button
                className={`px-1 text-white hover:text-white 
                  ${value.isEditing ? 
                    'bg-primary hover:bg-primary/90' :
                    'bg-blue-500 hover:bg-blue-700'
                  }`}
                type="button"
                onClick={() =>
                  dispatch({
                    type: "EDIT_VALUE",
                    payload: { typeTempId: typeTempId, valueTempId: value.tempId },
                  })
                }
              >
                {value.isEditing ? 
                (<>
                <Check fontSize='small'/>
                
                </>) : 
                (<>
                <Edit fontSize='small'/>
                
                </>)}
              </Button>

              <Button
                variant="destructive"
                type="button"
                onClick={() =>
                  dispatch({
                    type: "DELETE_VALUE",
                    payload: { typeTempId: typeTempId, valueTempId: value.tempId },
                  })
                }
              >
                <Delete fontSize="small"/>
              </Button>
            </div>
          </div>
        )}
      />
    </div>
  )
}
