import StrictNavbar from "@/components/customer/strict-navbar";

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <StrictNavbar/>
      {children}
    </>
  );
}
