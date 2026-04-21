import { OrderTabs } from "../order-tabs";
import OrderDetails from "./order-details";


export default async function OrderDetailsPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) {
  const { id } = await params;
  return (
    <section className='space-y-2'>
      <h2 className="text-xl font-bold col-span-2">Order Details</h2>
      <OrderTabs/>
      <OrderDetails id={id} />
    </section>
  );
}
