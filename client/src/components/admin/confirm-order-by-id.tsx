"use client"

import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from '@/components/ui/alert-dialog'
import { useConfirmOrder } from '@/hooks/order/useConfirmOrder'
import { Spinner } from '../ui/spinner'
import { Button } from '../ui/button'

const ConfirmOrderById = ({id}: {id: number}) => {
  const {mutate, isPending} = useConfirmOrder()
  const onConfirmOrder = () => {
    mutate(id)
  }

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button>Confirm</Button>
      </AlertDialogTrigger>
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

export default ConfirmOrderById
