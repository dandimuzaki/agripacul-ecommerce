import CheckoutForm from './checkout-form';

export default function CheckoutPage() { 
  return (
    <section className='space-y-2 p-4 pt-16 md:px-8 md:py-8 md:pt-24 pb-16'>
      <h2 className="uppercase text-xl md:text-2xl font-semibold col-span-2 text-primary">Checkout</h2>
      <CheckoutForm/>
    </section>
  );
};