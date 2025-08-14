"use client";

import { AlertCircle } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";

interface ErrorDisplayProps {
  title?: string;
  message?: string;
  dontTryAgain?: boolean;
}

const CustomOAuthErrorDisplay = ({
  title = "Something went wrong",
  message = "We encountered an unexpected error. Kindly check your internet connection, then try again or return to the home page.",
}: ErrorDisplayProps) => {
  return (
    <div className="min-h-screen bg-white flex items-center justify-center px-4">
      <div className="max-w-md w-full text-center">
        {/* Error Icon */}
        <div className="mb-8">
          <div className="w-24 h-24 mx-auto bg-blue-50 rounded-full flex items-center justify-center">
            <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center">
              <AlertCircle className="size-6 text-blue-500" />
            </div>
          </div>
        </div>

        <div className="mb-8">
          <h1 className="text-3xl font-bold text-black mb-4">{title}</h1>
          <p className="text-gray-600 text-lg leading-relaxed">{message}</p>
        </div>

        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Button variant={"link"}>
            <Link href={"/auth/signin"}>Signin Again </Link>
          </Button>
        </div>
      </div>
    </div>
  );
};

export default CustomOAuthErrorDisplay;
