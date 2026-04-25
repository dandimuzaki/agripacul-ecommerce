import ReviewList from "./review-list"

export default function AdminReviewPage() {
  return (
    <div className='space-y-2'>
      <h2 className='font-bold text-2xl'>Review Management</h2>
      <ReviewList/>
    </div>
  )
}