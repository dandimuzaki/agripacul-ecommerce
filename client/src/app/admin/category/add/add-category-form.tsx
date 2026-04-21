"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { Controller, useForm } from "react-hook-form"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardFooter,
} from "@/components/ui/card"
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { CategoryFormValues, categorySchema } from "@/schemas/category.schema"
import { useCreateCategory } from "@/hooks/category/useCreateCategory"
import { Spinner } from "@/components/ui/spinner"
import IconUploader from "../icon-uploader"

export default function AddCategoryForm() {
  const values: CategoryFormValues = {
    name: "",
    icon: undefined
  }

  const form = useForm<CategoryFormValues>({
    resolver: zodResolver(categorySchema),
    defaultValues: values,
  })

  const { mutate, isPending } = useCreateCategory()
    
  const onCreateCategory = (data: CategoryFormValues) => {
    mutate(data)
  }

  return (
    <div className="space-y-2">
      <h2 className="text-xl font-bold">Add New Category</h2>
      <Card className="w-full">
        <CardContent className="space-y-4">
          <form id="add-category" onSubmit={form.handleSubmit(onCreateCategory)}>
            <FieldGroup>
              <Controller
                name="name"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="add-category-name">
                      Name
                    </FieldLabel>
                    <Input
                      {...field}
                      id="add-category-name"
                      aria-invalid={fieldState.invalid}
                      placeholder="Fresh Vegetables"
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />

              <div>
                <FieldLabel>Icon</FieldLabel>
                <IconUploader form={form}/>
              </div>
            </FieldGroup>
          </form>
        </CardContent>
        <CardFooter>
          <Field orientation="horizontal" className="w-full flex items-center justify-center">
            <Button type="button" className="bg-gray-200 text-black hover:bg-gray-300 hover:text-black" onClick={() => form.reset()}>
              Reset
            </Button>
            <Button className="text-white hover:text-white" type="submit" form="add-category">
              {isPending ? <><Spinner/>Saving category...</> : <>Submit</>}
            </Button>
          </Field>
        </CardFooter>
      </Card>
    </div>
  )
}
