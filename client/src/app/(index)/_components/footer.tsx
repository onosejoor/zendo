"use client";

import React from "react";
import Img from "@/components/Img";
import Link from "next/link";
import { Twitter, Github } from "lucide-react";

export default function Footer() {
  return (
    <footer className="border-t bg-white/50 dark:bg-slate-900/50 dark:border-slate-700">
      <div className="container mx-auto px-4 py-6 flex flex-col sm:flex-row items-center justify-between gap-4">
        <div className="flex items-center gap-3">
          <Link href="/" className="flex items-center gap-3">
            <Img
              src={"/images/logo.svg"}
              alt="zendo logo"
              className="h-7.5 w-fit object-cover"
            />
          </Link>

          <span className="hidden sm:inline text-sm text-muted-foreground">
            © {new Date().getFullYear()} Zendo. All rights reserved.
          </span>
        </div>

        <div className="flex items-center gap-3">
          <Link
            href="https://x.com/DevText16"
            target="_blank"
            rel="noopener noreferrer"
            aria-label="Twitter"
            className="p-2 rounded-md hover:bg-slate-100 dark:hover:bg-slate-800"
          >
            <Twitter className="h-5 w-5" />
          </Link>

          <Link
            href="https://github.com/onosejoor/zendo"
            target="_blank"
            rel="noopener noreferrer"
            aria-label="GitHub"
            className="p-2 rounded-md hover:bg-slate-100 dark:hover:bg-slate-800"
          >
            <Github className="h-5 w-5" />
          </Link>
        </div>
      </div>
    </footer>
  );
}
