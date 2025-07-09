import { Button } from "@/components/ui/button";
import { SERVER_URl } from "@/lib/utils";
import { Separator } from "@radix-ui/react-separator";
import { useRouter } from "next/navigation";

export default function GoogleBtn({ isLoading }: { isLoading: boolean }) {
  const router = useRouter();
  return (
    <>
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <Separator className="w-full" />
        </div>
        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-auth-card px-2 text-auth-text-muted">
            Or continue with
          </span>
        </div>
      </div>

      <Button
        onClick={() => router.push(`${SERVER_URl}/auth/oauth/google`)}
        disabled={isLoading}
        variant="outline"
        className="border-auth-input-border my-4 text-auth-text-primary w-full hover:bg-auth-button-secondary-hover"
      >
        Google
      </Button>
    </>
  );
}
