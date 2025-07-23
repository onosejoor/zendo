import { Suspense } from "react";
import WaitComponant from "./_components/wait-page";
import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";

type Props = {
  searchParams: Promise<{ token: string }>;
};

export default async function VerifyEmailPage({ searchParams }: Props) {
  const token = (await searchParams).token;

  return getComponentToDisplay(token);
}

function getComponentToDisplay(token: string) {
  if (token.trim()) {
    if (token.split(".").length === 3) {
      return (
        <Suspense fallback={<Loader text="Verifying email..." />}>
          <WaitComponant token={token} />
        </Suspense>
      );
    } else {
      return <ErrorDisplay title="Invalid Token" />;
    }
  }
  return <></>;
}
