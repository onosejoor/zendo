import { Button } from "@/components/ui/button";
import { BACKUP_SERVER_URL, SERVER_URl } from "@/lib/utils";
import { Separator } from "@radix-ui/react-separator";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useRef } from "react";

export default function GoogleBtn({ isLoading }: { isLoading: boolean }) {
  const router = useRouter();
  const btnRef = useRef<HTMLButtonElement | null>(null);

  const goToGoogle = async () => {
    if (btnRef.current) {
      btnRef.current.textContent! = "Loading....";
    }

    try {
      await axios.get(`${SERVER_URl}/health`);
      router.push(`${SERVER_URl}/auth/oauth/google`);
    } catch (e) {
      router.push(`${BACKUP_SERVER_URL}/auth/oauth/google`);
      console.log(e);
      
    }
  };
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
        ref={btnRef}
        onClick={async () => await goToGoogle()}
        disabled={isLoading}
        variant="outline"
        className="border-auth-input-border my-4 text-auth-text-primary w-full hover:bg-auth-button-secondary-hover"
      >
        Google
      </Button>
    </>
  );
}
