import { BrowseProduct } from "./browse-product"
import ProductFilter from "./product-filter"

export default function BrowseProductPage() {
  return (
    <section className='grid gap-x-8 gap-y-4 md:gap-y-2 md:grid-cols-[1fr_3fr] md:px-8 md:py-8 md:pt-24 p-4 pt-16'>
      <ProductFilter/>
      <div className="">
        <h2 className="hidden md:block uppercase text-xl md:text-2xl font-semibold text-primary">Browse Products</h2>
        <BrowseProduct/>
      </div>
    </section>
  )
}
