"use client"

import { Button } from "@/components/ui/button"
import { CategoryFormValues } from "@/schemas/category.schema"
import { Delete, Upload } from "@mui/icons-material"
import { UploadCloud } from "lucide-react"
import { useEffect, useState } from "react"
import { UseFormReturn } from "react-hook-form"

const IconUploader = ({form, iconURL}: {form: UseFormReturn<CategoryFormValues>, iconURL?: string}) => {
  const [file, setFile] = useState<File | null>(null)
  const [preview, setPreview] = useState<string | null>(null)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selected = e.target.files?.[0]
    if (!selected) return

    setFile(selected)
    setPreview(URL.createObjectURL(selected))
    form.setValue("icon", selected)
    e.target.value = ""
  }

  useEffect(() => {
    return () => {
      if (preview?.startsWith("blob:")) {
        URL.revokeObjectURL(preview)
      }
    }
  }, [preview])

  const displayIcon = preview || iconURL
  
  const handleRemove = () => {
    setFile(null)
    setPreview(null)
  }

  return (
    <div className="space-y-4">
      {!displayIcon ? (
        <label className="w-48 h-48 flex flex-col items-center justify-center border-2 border-dashed rounded-lg p-6 cursor-pointer hover:bg-muted transition">
          <UploadCloud className="w-8 h-8 mb-2 text-muted-foreground" />
          <span className="text-sm text-muted-foreground text-center">
            Click to upload or drag image
          </span>
          <input
            type="file"
            accept="image/*"
            className="hidden"
            onChange={handleChange}
          />
        </label>
      ) : (
        <div className="relative w-48 h-48">
          <img
            src={displayIcon}
            alt="Preview"
            className="absolute z-2 hover:z-0 w-48 h-48 object-cover rounded-lg border top-0"
          />
          <div className="hover:z-3 hover:bg-black/40 rounded-lg 50 absolute top-0 w-full h-full flex justify-center items-center gap-2">
              <label className="bg-blue-500 hover:bg-blue-700 p-2 rounded-lg text-white">
                <Upload className="w-4 h-4"/>
                <input
                  type="file"
                  accept="image/*"
                  className="hidden"
                  onChange={handleChange}
                />
              </label>
            <Button
              size="icon"
              variant="destructive"
              className=""
              onClick={handleRemove}
            >
              <Delete className="w-4 h-4" />
            </Button>
          </div>
        </div>
      )}
    </div>
  )
}

export default IconUploader
