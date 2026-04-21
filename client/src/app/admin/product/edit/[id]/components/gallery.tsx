"use client"

import { useEffect, useMemo, useState } from "react"
import { UploadCloud, X } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { Spinner } from "@/components/ui/spinner"
import { Image, PreviewImage } from "@/types/product"
import { useUpdateGallery } from "@/hooks/product/useUpdateProductGallery"
import { useDeleteImage } from "@/hooks/product/useDeleteImage"
import { Field } from "@/components/ui/field"

const MAX_IMAGES = 5

const ProductGalleryUploader = ({
  id,
  images = [],
}: {
  id: number
  images: Image[]
}) => {
  const [localPreviews, setLocalPreviews] = useState<PreviewImage[]>([])

  const { mutate: uploadImages, isPending } = useUpdateGallery({setLocalPreviews: setLocalPreviews})
  const { mutate: deleteImage } = useDeleteImage()

  // ✅ Load existing images
  const previews = useMemo(() => {
    if (images != null) {
      const existing = images?.map((img) => ({
        id: img.id,
        tempId: `server-${img.id}`,
        image_url: img.image_url,
      }))

      return [...existing, ...localPreviews]
    } else {
      return [...localPreviews]
    }
  }, [images, localPreviews])

  // 🧹 Cleanup blob URLs
  useEffect(() => {
    return () => {
      localPreviews.forEach((img) => {
        if (img.image_url.startsWith("blob:")) {
          URL.revokeObjectURL(img.image_url)
        }
      })
    }
  }, [localPreviews])

  // 📥 Handle upload
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFiles = Array.from(e.target.files || [])

    if ((previews?.length ?? 0) + selectedFiles.length > MAX_IMAGES) {
      alert(`Maximum ${MAX_IMAGES} images allowed`)
      return
    }

    const newPreviews: PreviewImage[] = selectedFiles.map((file) => ({
      tempId: crypto.randomUUID(),
      file,
      image_url: URL.createObjectURL(file),
    }))

    setLocalPreviews((prev) => [...prev, ...newPreviews])
  }

  // ❌ Remove image
  const handleRemove = (tempId: string) => {
    // 🟢 Server image
    if (tempId.startsWith("server-")) {
      const imageId = Number(tempId.replace("server-", ""))

      deleteImage({
        productId: id,
        imageId,
      })

      return
    }

    // 🔵 Local image
    setLocalPreviews((prev) => {
      const target = prev.find((img) => img.tempId === tempId)
      if (!target) return prev

      if (target.image_url.startsWith("blob:")) {
        URL.revokeObjectURL(target.image_url)
      }

      return prev.filter((img) => img.tempId !== tempId)
    })
  }

  // 🚀 Submit
  const handleSubmit = () => {
    const filesToUpload = localPreviews
      .filter((img) => img.file)
      .map((img) => img.file!) // only new files

    if (filesToUpload.length === 0) return

    uploadImages({
      id,
      payload: {
        images: filesToUpload,
      },
    })

    if (!isPending) {
      setLocalPreviews([])
    }
  }

  return (
    <div className="space-y-2 w-full h-full flex-1">
      <h2 className="text-xl font-bold">Product Image Gallery</h2>

      <Card className="h-full w-full">
        <CardContent className="w-fit">
          <div className="grid grid-cols-4 gap-4 flex-1">

            {/* 🖼️ Preview */}
            {previews?.map((img) => (
              <div key={img.tempId} className="relative w-32 h-32">
                <img
                  src={img.image_url}
                  className="w-full h-full object-cover rounded-lg border"
                />

                <Button
                  type="button"
                  size="icon"
                  variant="destructive"
                  className="absolute top-[-4px] right-[-4px] w-6 h-6 rounded-full hover:bg-red-700 hover:scale-120 cursor-pointer"
                  onClick={() => handleRemove(img.tempId)}
                >
                  <X className="w-4 h-4" />
                </Button>
              </div>
            ))}

            {/* ➕ Upload */}
            {(previews?.length ?? 0) < MAX_IMAGES && (
              <label className="w-32 h-32 flex flex-col items-center justify-center border-2 border-dashed rounded-lg cursor-pointer hover:bg-muted transition">
                <UploadCloud className="w-6 h-6 text-muted-foreground mb-1" />
                <span className="text-xs text-muted-foreground">Upload</span>
                <input
                  type="file"
                  multiple
                  accept="image/*"
                  className="hidden"
                  onChange={handleChange}
                />
              </label>
            )}
          </div>

          <p className="text-xs text-muted-foreground mt-3">
            Max {MAX_IMAGES} images. First image will be the main image.
          </p>
        </CardContent>

        <CardFooter className="mt-4">
          <Field orientation="horizontal" className="w-full flex justify-center">
            <Button onClick={handleSubmit} disabled={isPending}>
              {isPending ? (
                <>
                  <Spinner /> Uploading...
                </>
              ) : (
                "Save Images"
              )}
            </Button>
          </Field>
        </CardFooter>
      </Card>
    </div>
  )
}

export default ProductGalleryUploader