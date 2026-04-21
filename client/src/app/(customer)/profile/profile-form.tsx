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
import { ProfileFormValues, profileSchema } from "@/schemas/profile.schema"
import { Spinner } from "@/components/ui/spinner"
import { useUpdateProfile } from "@/hooks/profile/useUpdateProfile"
import { useProfile } from "@/hooks/profile/useProfile"
import { useEffect, useState } from "react"
import ProfileImageUploader from "./profile-image-uploader"
import { useAuthStore } from "@/store/useAuthStore"
import AddressSection from "./address-section"

export default function ProfileForm() {
  const [onEdit, setOnEdit] = useState(false)
  const { data: profile } = useProfile()
  const { mutate, isPending } = useUpdateProfile({setOnEdit})

  const formatted = (date?: string) => {
    if (!date) return
    return date.split("T")[0]
  }
  
  const onUpdateProfile = (data: ProfileFormValues) => {
    mutate(data)
  }

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(profileSchema),
  })

  useEffect(() => {
    form.reset({
      full_name: profile?.full_name,
      phone_number: profile?.phone_number,
      date_of_birth: formatted(profile?.date_of_birth),
    })
  }, [profile])

  const editProfile = () => {
    setOnEdit(true)
  }

  const { user } = useAuthStore()

  return (
    <div className="grid lg:grid-cols-[192px_auto] gap-x-4 gap-y-2">
      <div className="hidden lg:block"></div>
      <h2 className="text-2xl font-semibold uppercase text-primary text-center lg:text-left">My Profile</h2>
      <ProfileImageUploader onEdit={onEdit} form={form} profileImageURL={profile?.profile_image_url}/>
      <div className="space-y-2 mb-2">
        <Card className="w-full">
          <CardContent>
            <form id="profile-form" onSubmit={form.handleSubmit(onUpdateProfile)}>
              <FieldGroup className="grid md:grid-cols-2 gap-4">
                <Controller
                  name="full_name"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="profile-full-name">
                        Full Name
                      </FieldLabel>
                      <Input
                        {...field}
                        disabled={!onEdit}
                        id="profile-full-name"
                        aria-invalid={fieldState.invalid}
                        placeholder="Enter your full name"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />

                <Field>
                  <FieldLabel htmlFor="email">
                    Email
                  </FieldLabel>
                  <Input
                    disabled
                    id="email"
                    type="email"
                    value={user?.email}
                  />
                </Field>

                <Controller
                  name="phone_number"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="phone-number">
                        Phone Number
                      </FieldLabel>
                      <Input
                        {...field}
                        disabled={!onEdit}
                        id="phone-number"
                        aria-invalid={fieldState.invalid}
                        placeholder="Enter your phone number"
                        type="phone"
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />

                <Controller
                  name="date_of_birth"
                  control={form.control}
                  render={({ fieldState }) => (
                    <Field data-invalid={fieldState.invalid}>
                      <FieldLabel htmlFor="date-of-birth">
                        Date of Birth
                      </FieldLabel>
                      <input
                        className="px-4 py-2 rounded-lg w-fit border border-input"
                        type="date" 
                        {...form.register("date_of_birth", { valueAsDate: true })} 
                        disabled={!onEdit}
                        id="date-of-birth"
                        aria-invalid={fieldState.invalid}
                      />
                      {fieldState.invalid && (
                        <FieldError errors={[fieldState.error]} />
                      )}
                    </Field>
                  )}
                />

              </FieldGroup>
            </form>
          </CardContent>
          <CardFooter>
            <Field orientation="horizontal" className="w-full flex items-center justify-center mt-4">
              {!onEdit ? <Button className="text-white hover:text-white" type="button" onClick={() => editProfile()}>
                Edit Profile
              </Button>
              :
              <>
              <Button type="button" onClick={() => setOnEdit(false)}>Cancel</Button>
              <Button className="text-white hover:text-white" type="submit" form="profile-form">
                {isPending ? <><Spinner/>Saving profile...</> : <>Save Updates</>}
              </Button></>}
            </Field>
          </CardFooter>
        </Card>
      </div>

      <AddressSection/>
    </div>
  )
}
