import Navbar from "@/components/customer/navbar";
import { Skeleton } from "@/components/ui/skeleton";
import { Suspense } from "react";

export default function NavbarLayout({
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
    </>
  );
}
