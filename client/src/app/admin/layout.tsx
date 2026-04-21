import Sidebar from "@/components/admin/sidebar";

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Sidebar/>
      <div className='ml-64 p-8'>
        {children}
      </div>
    </>
  );
}
