'use client';

import { UseFormReturn } from "react-hook-form";
import { VariantTypeRow } from "./variant-type-row";
import { FieldGroup } from "@/components/ui/field";
import { EditVariant, VariantTypeEdit } from "@/types/variant";
import { ProductDetails, Variant } from "@/types/product";
import { Button } from "@/components/ui/button";
import { Add } from "@mui/icons-material";
import { ProductFormValues } from "@/schemas/product.schema";

export type VariantAction =
  | { type: "INIT"; payload: VariantTypeEdit[] }
  | { type: "ADD_TYPE" }
  | {
      type: "UPDATE_TYPE"
      payload: { tempId: string; name: string }
    }
  | {
      type: "EDIT_TYPE"
      payload: { tempId: string }
    }
  | { type: "SAVE_TYPE"; payload: { tempId: string } }
  | {
      type: "DELETE_TYPE"
      payload: { tempId: string }
    }
  | {
      type: "ADD_VALUE"
      payload: { typeTempId: string }
    }
  | {
      type: "UPDATE_VALUE"
      payload: {
        typeTempId: string
        valueTempId: string
        value: string
      }
    }
  | {
      type: "EDIT_VALUE"
      payload: { 
        typeTempId: string
        valueTempId: string 
      }
    }
  | { type: "SAVE_VALUE"; 
      payload: { 
        typeTempId: string
        valueTempId: string 
      } 
    }
  | {
      type: "DELETE_VALUE"
      payload: {
        typeTempId: string
        valueTempId: string
      }
    }

export type VariantState = VariantTypeEdit[]

export const variantReducer: React.Reducer<VariantState, VariantAction> = (
  state,
  action
) => {
  switch (action.type) {
    case "INIT":
      return action.payload

    /* =======================
       VARIANT TYPE
    ======================= */

    case "ADD_TYPE":
      return [
        ...state.map(v => ({ ...v, isEditing: false })),
        {
          tempId: crypto.randomUUID(),
          originalName: "",
          name: "",
          values: [],
          status: "created",
          isEditing: true,
        },
      ]

    case "EDIT_TYPE":
      return state.map(v =>
        v.tempId === action.payload.tempId
          ? { ...v, isEditing: !v.isEditing, values: v.values.map(val => ({...val, isEditing: false })) }
          : { ...v, isEditing: false }
      )

    case "UPDATE_TYPE":
      return state.map(v =>
        v.tempId === action.payload.tempId
          ? {
              ...v,
              name: action.payload.name,
              status:
                v.originalName === action.payload.name
                  ? "clean"
                  : v.status === "created"
                  ? "created"
                  : "updated",
            }
          : v
      )

    case "DELETE_TYPE":
      return state.map(v =>
        v.tempId === action.payload.tempId
          ? { ...v, status: "deleted" }
          : v
      )

    /* =======================
       VARIANT VALUE
    ======================= */

    case "ADD_VALUE":
      return state.map(type =>
        type.tempId === action.payload.typeTempId
          ? {
              ...type,
              values: [
                ...type.values.map(v => ({ ...v, isEditing: false })),
                {
                  tempId: crypto.randomUUID(),
                  originalValue: "",
                  value: "",
                  status: "created",
                  isEditing: true,
                },
              ],
            }
          : type
      )

    case "EDIT_VALUE":
      return state.map(type =>
        type.tempId === action.payload.typeTempId
          ? {
              ...type,
              isEditing: false,
              values: type.values.map(v =>
                v.tempId === action.payload.valueTempId
                  ? { ...v, isEditing: !v.isEditing }
                  : { ...v, isEditing: false }
              ),
            }
          : type
      )

    case "UPDATE_VALUE":
      return state.map(type =>
        type.tempId === action.payload.typeTempId
          ? {
              ...type,
              values: type.values.map(v =>
                v.tempId === action.payload.valueTempId
                  ? {
                      ...v,
                      value: action.payload.value,
                      status:
                        v.originalValue === action.payload.value
                          ? "clean"
                          : v.status === "created"
                          ? "created"
                          : "updated",
                    }
                  : v
              ),
            }
          : type
      )

    case "DELETE_VALUE":
      return state.map(type =>
        type.tempId === action.payload.typeTempId
          ? {
              ...type,
              values: type.values.map(v =>
                v.tempId === action.payload.valueTempId
                  ? { ...v, status: "deleted" }
                  : v
              ),
            }
          : type
      )

    default:
      return state
  }
}

export const normalizeVariant = (product: ProductDetails): EditVariant => {
  let variants: Variant[] = []
  if (product.variants != null) {
    variants = product.variants
  }

  return ({
    ...product,
    category_id: product.category.id,
    variants: variants.map(v => ({
      ...v,
      tempId: crypto.randomUUID(),
      originalName: v.name,
      status: "clean",
      values: v.values ? v.values.map(val => ({
        ...val,
        tempId: crypto.randomUUID(),
        originalValue: val.value,
        status: "clean",
      })) : [],
    })),
  })}

export function VariantField({
  form,
  variantState,
  dispatchVariant,
}: {
  form: UseFormReturn<ProductFormValues>,
  variantState: VariantTypeEdit[],
  dispatchVariant: React.ActionDispatch<[action: VariantAction]>
}) {
  return (
    <div className="space-y-4">
      <FieldGroup>
        {variantState.map((variant, index) =>
          variant.status === "deleted" ? null : (
            <VariantTypeRow
              key={variant.tempId}
              variant={variant}
              dispatch={dispatchVariant}
              form={form}
              index={index}
            />
          )
        )}
      </FieldGroup>

      <Button
        className="bg-orange-500 text-white hover:text-white w-fit"
        type="button"
        onClick={() =>
          dispatchVariant({
            type: "ADD_TYPE",
          })
        }
      >
        <Add fontSize="small"/>
        Add Variant Type
      </Button>
    </div>
  )
}