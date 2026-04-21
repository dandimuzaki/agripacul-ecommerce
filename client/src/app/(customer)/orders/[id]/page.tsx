import OrderDetails from "@/components/customer/order-details";
import { ArrowBackIos } from "@mui/icons-material";
import Link from "next/link";

export default async function OrderDetailsPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) {
  const { id } = await params;
  return (
    <section className='space-y-2 md:px-8 md:py-8 md:pt-24 p-4 pt-16'>
      <Link href={"/orders"} className="flex items-center text-sm md:text-base"><ArrowBackIos fontSize="inherit"/>Back</Link>
      <h2 className="uppercase text-primary text-xl md:text-2xl font-semibold col-span-2">Order Details</h2>
      <OrderDetails id={id} />
    </section>
  );
}
