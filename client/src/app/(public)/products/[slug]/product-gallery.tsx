"use client"

import { useState } from "react"

export default function ProductGallery({images}: {images: string[]}) {

  if (images.length <= 0) {
    images.push("/loading.png")
  }

  const [selectedImage, setSelectedImage] = useState(images[0])

  return (
    <div className="flex gap-4">
      
      {/* Main Image */}
      <div className="order-2 w-full h-full relative border aspect-square rounded-xl overflow-hidden">
        <img
          src={selectedImage}
          alt="Product Image"
          className="object-cover aspect-square w-full"
        />
      </div>

      {/* Thumbnails */}
      <div className="order-1 flex flex-col gap-3">
        {images.map((img, index) => (
          <div
            key={index}
            className={`w-20 h-20 relative cursor-pointer border rounded-lg overflow-hidden 
              ${selectedImage === img ? "border-green-500" : "border-gray-300"}`}
            onClick={() => setSelectedImage(img)}
          >
            <img
              src={img}
              alt="Thumbnail"
              className={`object-cover w-full h-full ${selectedImage === img ? "" : "filter brightness-60"}`}
            />
          </div>
        ))}
      </div>

    </div>
  )
}