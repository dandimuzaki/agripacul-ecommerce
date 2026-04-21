import Navbar from "@/components/customer/navbar";

export default function NavbarLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Navbar/>
      {children}
    </>
  );
}
