import AdjustmentForm from "./adjustment-form";

export default async function AdminAdjustmentPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) { 
  const { id } = await params;

  return (
    <div className="space-y-8">
      <AdjustmentForm skuId={id} />
    </div>
  );
}
