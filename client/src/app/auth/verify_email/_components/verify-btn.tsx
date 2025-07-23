"use client";

import { Button } from "@/components/ui/button";
import { postToken } from "@/lib/actions/token";
import { useState } from "react";
import { toast } from "sonner";

export default function VerifyEmailBtn() {
  const [loading, setLoading] = useState(false);

  const sendNewToken = async () => {
    setLoading(true);
    const data = await postToken();
    toast(data.message);
    setLoading(false);
  };
  
  return (
    <Button
      onClick={async () => await sendNewToken()}
      variant="outline"
      className="border-gray-300 w-fit text-gray-700 hover:bg-gray-50 px-6 py-3 rounded-lg transition-colors duration-200 flex items-center justify-center gap-2"
    >
      {loading ? "Verifying..." : " Verify Email"}
    </Button>
  );
}
