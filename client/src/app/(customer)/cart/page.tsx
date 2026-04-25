import CartTable from './cart-table'
import CartSummary from './cart-summary'

export default function CartPage() {
  return (
    <section className='grid md:gap-x-4 md:gap-y-2 gap-2 lg:grid-cols-[7fr_3fr] md:px-8 md:py-8 md:pt-24 px-4 py-4 pt-16 pb-16'>
      <h2 className="uppercase text-xl md:text-2xl font-semibold md:col-span-2 text-primary">My Cart</h2>
      <CartTable/>
      <CartSummary/>
    </section>
  )
}
