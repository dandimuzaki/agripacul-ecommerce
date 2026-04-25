import Footer from "@/components/common/footer";
import "./../globals.css";
import Navbar from "@/components/customer/navbar";
import { Suspense } from "react";
import { Skeleton } from "@/components/ui/skeleton";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Suspense fallback={<Skeleton className="h-16 w-full bg-primary" />}>
        <Navbar />
      </Suspense>
      {children}
      <Footer/>
    </>
  );
}
