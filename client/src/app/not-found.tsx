import { Button } from "@/components/ui/button";
import { ArrowLeft, Search } from "lucide-react";
import Link from "next/link";

export default function NotFound() {
  return (
    <div className="min-h-screen bg-white flex items-center justify-center px-4">
      <div className="max-w-md w-full text-center">
        <div className="mb-8">
          <div className="w-24 h-24 mx-auto bg-blue-50 rounded-full flex items-center justify-center">
            <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center">
              <Search className="size-6 text-blue-500" />
            </div>
          </div>
        </div>

        <div className="mb-8">
          <h1 className="text-3xl font-bold text-black mb-4">
            404 Page Not Found
          </h1>
          <p className="text-gray-600 text-lg leading-relaxed">
            The page you&apos;re looking for does not exixt, ot might&apos;ve
            been moved
          </p>
        </div>

        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Link href={"/dashboard"} className="block">
            <Button variant="outline">
              <ArrowLeft className="w-4 h-4" />
              Back to Home
            </Button>
          </Link>
        </div>

        <div className="mt-8 pt-8 border-t border-gray-200">
          <p className="text-sm text-gray-500">
            If this problem persists, please contact our support team.
          </p>
          <a href="https://onos-ejoor.vercel.app/contact">
            <Button variant={"link"}>Contact Form</Button>
          </a>
        </div>
      </div>
    </div>
  );
}
