import { formatRupiah } from "@/lib/formatCurrency"
import { OrderItem as item } from "@/types/order"
import Image from "next/image"

const OrderItem = ({item}: {item: item}) => {
  return (
    <div className="flex gap-2">
      <img 
        className="w-15 h-15 object-cover aspect-square rounded"
        src={item.main_image_url ?? "/loading.png"} alt={item.name} width={100} height={100}/>
      <div className="grid grid-cols-[5fr_2fr] flex-1">
        <p className="font-medium">{item.name}</p>
        <div className="text-right">{item.quantity} x {formatRupiah(item.price)}</div>
        <div className="flex flex-wrap gap-2">
          {item.variants !== null && item.variants.map((v, index) => (<p className="bg-gray-200 px-2 py-1 text-xs w-fit h-fit rounded" key={index}>{v.name} : {v.value}</p>))}
        </div>
        <p className="text-lg font-medium text-right h-fit">{formatRupiah(item.total_price)}</p>
      </div>
    </div>
  )
}

export default OrderItem
