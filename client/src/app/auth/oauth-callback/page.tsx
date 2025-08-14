import { Suspense } from "react";
import WaitComponant from "./_components/wait-page";
import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";

type Props = {
  searchParams: Promise<{ code: string }>;
};

export default async function VerifyEmailPage({ searchParams }: Props) {
  const code = (await searchParams).code;

  return getComponentToDisplay(code);
}

function getComponentToDisplay(code: string) {
  if (code.trim()) {
    return (
      <Suspense fallback={<Loader text="Signin In..." />}>
        <WaitComponant code={code} />
      </Suspense>
    );
  } else {
    return <ErrorDisplay title="Invalid Code" />;
  }
}
