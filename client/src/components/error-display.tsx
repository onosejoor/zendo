"use client";

import { RefreshCcw, ArrowLeft, AlertCircle } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";

interface ErrorDisplayProps {
  title?: string;
  message?: string;
  dontTryAgain?: boolean;
}

const ErrorDisplay = ({
  dontTryAgain = false,
  title = "Something went wrong",
  message = "We encountered an unexpected error. Please try again or return to the home page.",
}: ErrorDisplayProps) => {
  const handleRefresh = () => {
    window.location.reload();
  };

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
          {!dontTryAgain && (
            <Button
              onClick={handleRefresh}
              className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg transition-colors duration-200 flex items-center justify-center gap-2"
            >
              <RefreshCcw className="w-4 h-4" />
              Try Again
            </Button>
          )}

          <Link href={"/dashboard"}>
            <Button
              variant="outline"
              className="border-gray-300 text-gray-700 hover:bg-gray-50 px-6 py-3 rounded-lg transition-colors duration-200 flex items-center justify-center gap-2"
            >
              <ArrowLeft className="w-4 h-4" />
              Back to Home
            </Button>
          </Link>
        </div>

        {/* Additional Help Text */}
        <div className="mt-8 pt-8 border-t border-gray-200">
          <p className="text-sm text-gray-500">
            If this problem persists, please contact our support team.
          </p>
        </div>
      </div>
    </div>
  );
};

export default ErrorDisplay;
