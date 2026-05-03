import { RatingStars } from '@/components/customer/rating-star'
import { useReviewFilter } from '@/hooks/review/useReviewFilter'
import { useReviewsByProduct } from '@/hooks/review/useReviewsByProduct'
import { capitalizeAll } from '@/lib/formatText'
import { Review } from '@/types/review'
import React from 'react'

const ReviewSection = ({productId}: {productId: number}) => {
  const {filters} = useReviewFilter()
  const {data} = useReviewsByProduct(productId, filters)
  const reviews = data?.data
  const pagination = data?.pagination

  if (!pagination) {
    return <></>;
  }

  return (
    <section className="p-4 md:p-8 space-y-2 md:space-y-4">
      <h3 className='font-semibold text-xl'>Customer Reviews</h3>
      <div className='grid grid-cols-2 gap-4'>
        {reviews.map((r: Review) => (
          <div className='p-2 rounded-lg bg-white space-y-2' key={r.id}>
            <RatingStars rating={r.rating}/>
            <h4 className='font-semibold text-lg'>{capitalizeAll(r.customer_name)}</h4>
            <p className='text-sm text-gray-500'>{r.comment}</p>
          </div>
        ))}
      </div>
    </section>
  )
}

export default ReviewSection
