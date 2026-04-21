import EditCategoryForm from "./edit-category-form";

export default async function EditCategoryPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) { 
  const { id } = await params;

  return (
    <>
      <EditCategoryForm id={id}/>
    </>
  );
}
