import { InventoryTabs } from "./inventory-tabs"

const AdminInventoryPage = () => {
  return (
    <div className='space-y-4'>
      <h2 className='font-bold text-2xl'>Inventory Management</h2>
      <InventoryTabs/>
    </div>
  )
}

export default AdminInventoryPage
