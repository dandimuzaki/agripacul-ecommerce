"use client"

import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle } from '@/components/ui/alert-dialog'
import { useConfirmOrder } from '@/hooks/order/useConfirmOrder'
import { Spinner } from '../ui/spinner'
import { Dispatch, SetStateAction } from 'react'

const ConfirmOrder = ({id, open, setOpen}: {id: number, open: boolean, setOpen: Dispatch<SetStateAction<boolean>>}) => {
  const {mutate, isPending} = useConfirmOrder()
  const onConfirmOrder = () => {
    mutate(id)
  }

  return (
    <AlertDialog open={open} onOpenChange={setOpen}>
      <AlertDialogContent>
        <AlertDialogHeader className='text-center'>
          <AlertDialogTitle>Confirm order now?</AlertDialogTitle>
          <AlertDialogDescription>
            Once confirmed, this order will be processed and can’t be changed.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={onConfirmOrder}>
            {isPending ? <><Spinner/>Confirming order...</> : <>Confirm</>}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default ConfirmOrder
