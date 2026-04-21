'use client';

import { VariantTypeEdit } from "@/types/variant";
import { VariantValueRow } from "./variant-value-row";
import { Controller, UseFormReturn } from "react-hook-form";
import { Field, FieldError, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { VariantAction } from "./variant-field";
import { Button } from "@/components/ui/button";
import { Add, Check, Delete, Edit } from "@mui/icons-material";
import { ProductFormValues } from "@/schemas/product.schema";

export function VariantTypeRow({ 
  variant, 
  dispatch, 
  form,
  index,
}: {
  variant: VariantTypeEdit, 
  dispatch: React.ActionDispatch<[action: VariantAction]>,
  form: UseFormReturn<ProductFormValues>,
  index: number,
}) {
  return (
    <div className="space-y-2 p-4 bg-primary/10 rounded">
      <Controller
        name={`variants.${index}.name`}
        control={form.control}
        render={({ field, fieldState }) => (
          <div className="flex gap-2 w-fit">
            <Field data-invalid={fieldState.invalid} className="grid grid-cols-[110px_auto] flex-1 gap-2">
              <FieldLabel className="whitespace-nowrap">Variant Name : </FieldLabel>

              <Input
                form="variant"
                {...field}
                disabled={!variant.isEditing}
                state={variant.status}
                placeholder="Size"
                onChange={(e) => {
                  field.onChange(e)
                  dispatch({
                    type: "UPDATE_TYPE",
                    payload: {
                      tempId: variant.tempId,
                      name: e.target.value,
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
                  ${variant.isEditing ? 
                    'bg-primary hover:bg-primary/90' :
                    'bg-blue-500 hover:bg-blue-700'
                  }`}
                type="button"
                onClick={() =>
                  dispatch({
                    type: "EDIT_TYPE",
                    payload: { tempId: variant.tempId },
                  })
                }
              >
                {variant.isEditing ? 
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
                    type: "DELETE_TYPE",
                    payload: { tempId: variant.tempId },
                  })
                }
              >
                <Delete fontSize="small"/>
                
              </Button>
            </div>

          </div>
        )}
      />

      <div className="grid grid-cols-[110px_auto] gap-2">
        <FieldLabel className="whitespace-nowrap w-full">Variant Value : </FieldLabel>
        <div className="p-4 rounded-lg">
          {variant.values.map((value, i) =>
            value.status === "deleted" ? null : (
              <VariantValueRow
                key={value.tempId}
                indexType={index}
                indexValue={i}
                value={value}
                typeTempId={variant.tempId}
                form={form}
                dispatch={dispatch}
              />
            )
          )}
        </div>
      </div>

      <div className="grid grid-cols-[110px_auto] gap-2">
        <div></div>
        <Button
          className="bg-yellow-500 text-white hover:text-white w-fit"
          type="button"
          onClick={() =>
            dispatch({
              type: "ADD_VALUE",
              payload: { typeTempId: variant.tempId },
            })
          }
        >
          <Add fontSize="small"/>
          Add Value
        </Button>
      </div>
    </div>
  )
}
