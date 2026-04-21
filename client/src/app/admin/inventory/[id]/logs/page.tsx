import InventoryLogs from "./logs-table";

export default async function InventoryLogsPage({
  params,
}: {
  params: Promise<{ id: number }>;
}) { 
  const { id } = await params;

  return (
    <div className='space-y-4'>
      <h2 className='font-bold text-2xl'>Inventory Logs</h2>
      <InventoryLogs id={id}/>
    </div>
  );
}
