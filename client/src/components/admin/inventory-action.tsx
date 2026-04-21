import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Inventory } from "@/types/inventory"
import { AssignmentAdd, FactCheck, FormatListBulleted, MoreVert } from "@mui/icons-material"
import Link from "next/link"

export function InventoryAction({inventory}: {inventory: Inventory}) {
  
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button className="rounded-lg h-10 w-10 bg-primary/20 text-sm text-primary">
          <MoreVert fontSize="inherit"/>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuGroup>
          <Link href={`/admin/inventory/${inventory.id}/logs`}>
            <DropdownMenuItem>
              <FormatListBulleted fontSize="inherit"/>
              See Logs
            </DropdownMenuItem>
          </Link>
          <Link href={`/admin/inventory/${inventory.id}/restock`}>
            <DropdownMenuItem>
              <AssignmentAdd fontSize="inherit"/>
              Restock
            </DropdownMenuItem>
          </Link>
          <Link href={`/admin/inventory/${inventory.id}/adjustment`}>
            <DropdownMenuItem>
              <FactCheck fontSize="inherit"/>
              Adjust Stock
            </DropdownMenuItem>
          </Link>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
