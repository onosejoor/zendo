"use client";

import { sendOauthCode } from "@/lib/actions/token";
import { useEffect, useState } from "react";
import Loader from "@/components/loader-card";
import SuccessComp from "@/app/_components/success-comp";
import ErrorDisplay from "@/components/error-display";

export default function WaitComponant({ code }: { code: string }) {
  const [result, setResult] = useState<APIRes | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function verifyToken() {
      const res = await sendOauthCode(code);
      setResult(res);
      setLoading(false);
    }

    verifyToken();
  }, [code]);

  if (loading) {
    return <Loader text="Signin In..." />;
  }

  if (result?.success) {
    return (
      <SuccessComp
        title="Signin Successfully"
        message=" Oauth Signing Successfull, redirecting now"
        redirectRoute="/dashboard"
      />
    );
  }

  return (
    <ErrorDisplay
      title="Error Signin In"
      message={result?.message}
    />
  );
}
