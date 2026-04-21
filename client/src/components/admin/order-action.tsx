import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { OrderSummary } from "@/types/order"
import { Check, Close, MoreVert, Visibility } from "@mui/icons-material"
import Link from "next/link"
import ConfirmOrder from "./confirm-order"
import { useState } from "react"

export function OrderAction({order}: {order: OrderSummary}) {
  const [openConfirmOrder, setOpenConfirmOrder] = useState(false)
  
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button className="rounded-lg h-10 w-10 bg-primary/20 text-sm text-primary">
          <MoreVert fontSize="inherit"/>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuGroup>
          <Link href={`/admin/order/${order.id}`}>
            <DropdownMenuItem>
              <Visibility fontSize="inherit"/>
              See Details
            </DropdownMenuItem>
          </Link>
          <DropdownMenuItem onClick={() => setOpenConfirmOrder(true)}>
            <Check fontSize="inherit"/>
            Confirm Order
          </DropdownMenuItem>
          {order.cancellation && <DropdownMenuItem>
            <Close fontSize="inherit"/>
            Cancel Order
          </DropdownMenuItem>}
        </DropdownMenuGroup>
      </DropdownMenuContent>
      <ConfirmOrder id={order.id} open={openConfirmOrder} setOpen={setOpenConfirmOrder}/>
    </DropdownMenu>
  )
}
