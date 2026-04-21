import ReviewList from "./review-list"

const AdminReviewPage = () => {
  return (
    <div className='space-y-2'>
      <h2 className='font-bold text-2xl'>Review Management</h2>
      <ReviewList/>
    </div>
  )
}

export default AdminReviewPage
