import Footer from "@/components/common/footer";
import "./../globals.css";
import Navbar from "@/components/customer/navbar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Navbar/>
      {children}
      <Footer/>
    </>
  );
}
