"use client"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { Field } from "@/components/ui/field"
import { Spinner } from "@/components/ui/spinner"
import { useProductFilter } from "@/hooks/product/useProductFilter"
import { useUpdateMainImage } from "@/hooks/product/useUpdateMainImage"
import { Delete, Upload } from "@mui/icons-material"
import { UploadCloud, X } from "lucide-react"
import { useEffect, useState } from "react"

const MainImageUploader = ({id, image}: {id: number, image?: string}) => {
  const {form} = useProductFilter()
  const [file, setFile] = useState<File | null>(null)
  const [preview, setPreview] = useState<string | null>(null)
  const {mutate, isPending} = useUpdateMainImage()

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selected = e.target.files?.[0]
    if (!selected) return

    setFile(selected)
    setPreview(URL.createObjectURL(selected))
  }

  const handleRemove = () => {
    setFile(null)
    setPreview(null)
  }

  useEffect(() => {
    return () => {
      if (preview) URL.revokeObjectURL(preview)
    }
  }, [preview])

  useEffect(() => {
    return () => {
      if (image) {
        setPreview(image)
      }
    }
  }, [image])

  const onSubmit = () => {
    if (file) {
      mutate({
        id: id,
        payload: {
          image: file
        }
      })
    }
  }

  return (
    <div className="space-y-2 w-fit">
      <h2 className="text-xl font-bold">Edit Product Main Image</h2>
      <Card className="w-fit">
        <CardContent className="space-y-4">
          <form id="edit-main-image" onSubmit={form.handleSubmit(onSubmit)}>
            <div className="space-y-4">
              {!preview ? (
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
                    src={preview}
                    alt="Preview"
                    className="absolute z-2 hover:z-0 w-48 h-48 object-cover rounded-lg border top-0"
                  />
                  <div className="hover:z-2 hover:bg-black/40 rounded-lg 50 absolute top-0 w-full h-full flex justify-center items-center gap-2">
                      <div className="bg-blue-500 hover:bg-blue-700 p-2 rounded-lg text-white">
                        <Upload className="w-4 h-4"/>
                        <input
                          type="file"
                          accept="image/*"
                          className="hidden"
                          onChange={handleChange}
                        />
                      </div>
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
          </form>
        </CardContent>
        <CardFooter className="mt-4">
          <Field orientation="horizontal" className="w-full flex items-center justify-center">
            <Button type="submit" form="edit-main-image" disabled={isPending}>
              {isPending ? (
                <>
                  <Spinner /> Uploading...
                </>
              ) : (
                "Save Main Image"
              )}
            </Button>
          </Field>
        </CardFooter>
      </Card>
    </div>
  )
}

export default MainImageUploader
