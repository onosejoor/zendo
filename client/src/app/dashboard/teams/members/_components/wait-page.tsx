"use client";

import { sendInviteToken } from "@/lib/actions/token";
import SuccessComp from "@/app/_components/success-comp";
import { useEffect, useState } from "react";
import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";

export default function WaitComponant({ token }: { token: string }) {
  const [result, setResult] = useState<(APIRes & { team_id: string }) | null>(
    null
  );
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function verifyToken() {
      const res = await sendInviteToken(token);
      setResult(res);
      setLoading(false);
    }

    verifyToken();
  }, [token]);

  if (loading) {
    return <Loader text="Creating Member..." />;
  }

  if (result?.success) {
    return (
      <SuccessComp
        title="Team member Created Successfully"
        message="You have been added to the team, redirecting now"
        redirectRoute={`/dashboard/teams/${result?.team_id}`}
      />
    );
  }

  return (
    <ErrorDisplay
      title="Error Adding New member To Team"
      message={result?.message}
    />
  );
}
