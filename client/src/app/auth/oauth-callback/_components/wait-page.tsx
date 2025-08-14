'use client'

import { sendOauthCode } from "@/lib/actions/token";
import SuccessComp from "./success-comp";
import { useEffect, useState } from "react";
import CustomOAuthErrorDisplay from "./custom-error-comp";
import Loader from "@/components/loader-card";

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
    return <SuccessComp />;
  }

  return (
    <CustomOAuthErrorDisplay
      title="Error Signin In"
      message={result?.message}
    />
  );
}
