import { ProductDetails } from '@/types/product';
import React, { Dispatch, SetStateAction } from 'react'

const VariantSelection = ({
  product,
  selectedVariants,
  setSelectedVariants
}: {
  product: ProductDetails,
  selectedVariants: Record<number, number>,
  setSelectedVariants: Dispatch<SetStateAction<Record<number, number>>>
}) => {
  return (
    <div className="my-2 space-y-2">
      {product.variants != null && product.variants.length > 0 && product.variants.map((variant) => (
        <div key={variant.id}>
          <p className="font-medium mb-2">{variant.name}</p>

          <div className="flex flex-wrap gap-2">
            {variant.values != null && variant.values.length > 0 && variant.values.map((value) => {
              const isSelected =
                selectedVariants[variant.id] === value.id;

              return (
                <button
                  key={value.id}
                  onClick={() =>
                    setSelectedVariants((prev) => ({
                      ...prev,
                      [variant.id]: value.id,
                    }))
                  }
                  className={`px-4 py-2 border rounded text-sm
                    ${
                      isSelected
                        ? "border-primary bg-primary/10 text-primary"
                        : "border-gray-300 hover:border-primary"
                    }`}
                >
                  {value.value}
                </button>
              );
            })}
          </div>
        </div>
      ))}
    </div>
  )
}

export default VariantSelection
