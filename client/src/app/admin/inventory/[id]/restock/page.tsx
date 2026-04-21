import RestockForm from "./restock-form";

export default async function AdminRestockPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) { 
  const { id } = await params;

  return (
    <div className="space-y-8">
      <RestockForm skuId={id} />
    </div>
  );
}
