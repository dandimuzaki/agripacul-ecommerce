import { formatRupiah } from "@/lib/formatCurrency"
import { CheckoutItem as item } from "@/types/checkout"
import Image from "next/image"

const CheckoutItem = ({item}: {item: item}) => {
  return (
    <div className="flex gap-2">
      <Image 
        className="w-16 h-16 object-cover aspect-square rounded"
        src={"/cherry-tomato.png"} alt={item.product.name} width={100} height={100}/>
      <div className="grid grid-cols-[5fr_2fr] flex-1">
        <div className="space-y-2">
          <p className="font-medium">{item.product.name}</p>
          <div className="flex flex-wrap gap-2">
            {item.variants.map((v, index) => (<p className="bg-gray-200 px-2 py-1 text-xs w-fit h-fit rounded" key={index}>{v.name} : {v.value}</p>))}
          </div>
        </div>
        <div className="space-y-2">
          <div className="text-right">
            {item.quantity} x {formatRupiah(
              item.price.sale_price && item.price.sale_price > 0
                ? item.price.sale_price
                : item.price.unit_price
            )}
          </div>
          <p className="text-lg font-medium text-right h-fit">{formatRupiah(item.subtotal)}</p>
        </div>
      </div>
    </div>
  )
}

export default CheckoutItem
