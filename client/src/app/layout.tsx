import type React from "react";
import type { Metadata } from "next";
import { Geist } from "next/font/google";
import "./globals.css";
import { Toaster } from "sonner";
import NextTopLoader from "nextjs-toploader";
import { layoutMetadata } from "./metadata";

const geist = Geist({ subsets: ["latin"], variable: "--font-geist-sans" });

export const metadata: Metadata = {
  ...layoutMetadata,
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${geist.className} font-sans `}>
        <NextTopLoader height={5} color="var(--color-accent-blue)" />
        <Toaster />
        {children}
      </body>
    </html>
  );
}
