"use client";

import { sendEmailToken } from "@/lib/actions/token";
import SuccessComp from "@/app/_components/success-comp";
import { useEffect, useState } from "react";
import Loader from "@/components/loader-card";
import CustomEmailAuthErrorDisplay from "./custom-error-comp";

export default function WaitComponant({ token }: { token: string }) {
  const [result, setResult] = useState<APIRes | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function verifyToken() {
      const res = await sendEmailToken(token);
      setResult(res);
      setLoading(false);
    }

    verifyToken();
  }, [token]);

  if (loading) {
    return <Loader text="Verifying email..." />;
  }

  if (result?.success) {
    return (
      <SuccessComp
        title="Email Verified SuccessFully"
        message="Email successfully verified, redirecting now"
        redirectRoute="/dashboard"
      />
    );
  }

  return (
    <CustomEmailAuthErrorDisplay
      title="Error Verifying Email"
      message={result?.message}
    />
  );
}
