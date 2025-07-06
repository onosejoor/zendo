import React from "react";
import { LoaderCircle } from "lucide-react";

interface LoaderProps {
  text?: string;
  size?: "sm" | "md" | "lg";
}

const Loader = ({ text = "Loading...", size = "md" }: LoaderProps) => {
    
  const sizeClasses = {
    sm: "w-4 h-4",
    md: "w-8 h-8",
    lg: "w-12 h-12",
  };

  const containerSizeClasses = {
    sm: "gap-2 text-sm",
    md: "gap-3 text-base",
    lg: "gap-4 text-lg",
  };

  return (
    <div className="min-h-screen bg-white flex items-center justify-center px-4">
      <div className="text-center">
        <div
          className={`flex flex-col items-center justify-center ${containerSizeClasses[size]}`}
        >
          <div className="relative">
            <div className="w-16 h-16 bg-blue-50 rounded-full flex items-center justify-center mb-4">
              <LoaderCircle
                className={`${sizeClasses[size]} text-blue-600 animate-spin`}
              />
            </div>
          </div>

          {text && <p className="text-gray-600 font-medium">{text}</p>}
        </div>
      </div>
    </div>
  );
};

export default Loader;
