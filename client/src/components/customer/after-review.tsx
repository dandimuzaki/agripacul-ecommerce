"use client"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '../ui/dialog';
import { CheckCircle } from '@mui/icons-material';
import { Dispatch, SetStateAction } from 'react';
import { Button } from '../ui/button';

const AfterReview = ({closeAfterReview, openAfterReview, setOpenAfterReview}: {closeAfterReview: () => void, openAfterReview: boolean, setOpenAfterReview: Dispatch<SetStateAction<boolean>>}) => {

  return (
    <Dialog open={openAfterReview} onOpenChange={setOpenAfterReview}>
      <DialogContent className='overflow-y-auto'>
        <DialogHeader>
          <DialogTitle>
            Share Your Review
          </DialogTitle>
        </DialogHeader>
        <div className='flex flex-col items-center'>
          <div className="text-primary text-[120px]/30 mb-2">
            <CheckCircle fontSize="inherit" />
          </div>
          <p className='font-semibold text-base md:text-lg'>Your reviews have been submitted successfully</p>
          <p className='text-center text-sm md:text-base'>Thank you! Your feedback helps us grow better for you</p>
          <Button onClick={closeAfterReview} className='mt-4 text-sm text-black px-4 py-2 bg-gray-300 cursor-pointer hover:bg-gray-400'>
            Close
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default AfterReview;