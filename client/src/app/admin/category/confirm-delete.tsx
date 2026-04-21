import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from '@/components/ui/alert-dialog'
import { Button } from '@/components/ui/button'
import { useDeleteCategory } from '@/hooks/category/useDeleteCategory'
import { Delete } from '@mui/icons-material'

const ConfirmDelete = ({id}: {id:number}) => {
  const {mutate: onDelete} = useDeleteCategory()

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="destructive" className='bg-red-500 text-white'><Delete fontSize='small' /></Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader className='text-center'>
          
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone. This will permanently delete the
            category.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction 
            variant="destructive"
            onClick={() => onDelete({id: id})}
          >
            Delete
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default ConfirmDelete
