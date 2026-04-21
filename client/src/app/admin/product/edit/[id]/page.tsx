import UpdateProductSection from "./update-product-section";

export default async function EditProductPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) { 
  const { id } = await params;

  return (
    <div className="space-y-8">
      <UpdateProductSection id={id}/>
    </div>
  );
}
