"use client"

import { ProfileFormValues } from "@/schemas/profile.schema"
import { useEffect, useState } from "react"
import { UseFormReturn } from "react-hook-form"

const ProfileImageUploader = ({onEdit, form, profileImageURL}: {onEdit: boolean, form: UseFormReturn<ProfileFormValues>, profileImageURL?: string}) => {
  const [file, setFile] = useState<File | null>(null)
  const [preview, setPreview] = useState<string | null>(null)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selected = e.target.files?.[0]
    if (!selected) return

    setFile(selected)
    setPreview(URL.createObjectURL(selected))
    form.setValue("profile_image", selected)
  }

  useEffect(() => {
    return () => {
      if (preview?.startsWith("blob:")) {
        URL.revokeObjectURL(preview)
      }
    }
  }, [preview])

  const displayImage = preview || profileImageURL || "/profile.png"

  return (
    <div className="flex items-center flex-col gap-4 row-span-2">
      {!preview ? (
        <div className="w-48 h-48 rounded-full relative">
          <img
            src={displayImage}
            alt="Preview"
            className="absolute z-2 hover:z-0 w-48 h-48 object-cover rounded-full border top-0"
          />
        </div>
      ) : (
        <div className="relative w-48 h-48">
          <img
            src={displayImage}
            alt="Preview"
            className="absolute z-2 hover:z-0 w-48 h-48 object-cover rounded-full border top-0"
          />
        </div>
      )}

      {onEdit && <label className="px-4 py-2 rounded-lg bg-gray-200 hover:bg-gray-300 text-black">
        <span className="text-sm text-center">
          Upload
        </span>
        <input
          type="file"
          accept="image/*"
          className="hidden"
          onChange={handleChange}
        />
      </label>}
    </div>
  )
}

export default ProfileImageUploader
