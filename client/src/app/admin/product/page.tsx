import ProductList from './components/product-list';

const AdminProductPage = () => {
  return (
    <div className='space-y-4'>
      <h2 className='font-bold text-2xl'>Product Management</h2>
      <ProductList/>
    </div>
  )
}

export default AdminProductPage
