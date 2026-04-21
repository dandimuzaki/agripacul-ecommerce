import type { Metadata } from "next";
import { Josefin_Sans } from "next/font/google";
import "./globals.css";
import Provider from "./provider";
import PageAnimation from "@/components/common/page-animation";
import { Toaster } from "sonner";

const josefinSans = Josefin_Sans({
  variable: "--font-josefin",
  weight: ["300", "400", "500", "600", "700"],
});

export const metadata: Metadata = {
  title: "Agripacul – Fresh Produce & Garden Essentials",
  description:
    "Naturally grown vegetables, edible flowers, seeds, and gardening tools delivered fresh from our farm to your table.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={josefinSans.className}>
      <body
        className={`${josefinSans.variable} antialiased`}
      >
          <Provider>
            {children}
            <Toaster position="top-right" />
          </Provider>
      </body>
    </html>
  );
}
