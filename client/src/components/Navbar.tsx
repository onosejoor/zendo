import { Button } from "./ui/button";
import Link from "next/link";
import Img from "./Img";

export default function Nav() {
  return (
    <header className="border-b bg-white/80 backdrop-blur-sm">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <Link href={"/"}>
          <Img
            src={"/images/logo.svg"}
            alt="zendo logo"
            className="h-7.5 w-fit object-cover"
          />
        </Link>
        <div className="space-x-4">
          <Link href="/auth/signin" className="sm:inline hidden">
            <Button variant="ghost"> Login</Button>
          </Link>
          <Link href="/auth/signup">
            <Button>Get Started</Button>
          </Link>
        </div>
      </div>
    </header>
  );
}
