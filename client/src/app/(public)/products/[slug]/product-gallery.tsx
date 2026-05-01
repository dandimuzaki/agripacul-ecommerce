"use client"

import { useState } from "react"
import Image from "next/image"

export default function ProductGallery({images}: {images: string[]}) {

  if (images.length <= 0) {
    images.push("/loading.png")
  }

  const [selectedImage, setSelectedImage] = useState(images[0])

  return (
    <div className="flex flex-col gap-4">
      
      {/* Main Image */}
      <div className="w-full h-full relative border aspect-square rounded-xl overflow-hidden">
        <Image
          src={selectedImage}
          alt="Product Image"
          fill
          className="object-cover aspect-square w-full"
        />
      </div>

      {/* Thumbnails */}
      <div className="flex gap-3">
        {images.map((img, index) => (
          <div
            key={index}
            className={`w-20 h-20 relative cursor-pointer border rounded-lg overflow-hidden 
              ${selectedImage === img ? "border-green-500" : "border-gray-300"}`}
            onClick={() => setSelectedImage(img)}
          >
            <Image
              src={img}
              alt="Thumbnail"
              fill
              className={`object-cover ${selectedImage === img ? "" : "filter brightness-60"}`}
            />
          </div>
        ))}
      </div>

    </div>
  )
}