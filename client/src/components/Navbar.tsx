import { CheckSquare, Link } from "lucide-react";
import { Button } from "./ui/button";

export default function Nav() {
  return (
    <header className="border-b bg-white/80 backdrop-blur-sm">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div className="flex items-center space-x-2">
          <CheckSquare className="h-8 w-8 text-blue-600" />
          <span className="text-2xl font-bold text-gray-900">TaskFlow</span>
        </div>
        <div className="space-x-4">
          <Link href="/auth/login">
            <Button variant="ghost">Login</Button>
          </Link>
          <Link href="/auth/signup">
            <Button>Get Started</Button>
          </Link>
        </div>
      </div>
    </header>
  );
}
