"use client";

import ErrorDisplay from "@/components/error-display";

type Props = {
  error: Error & { digest?: string };
  reset: () => void;
};

export default function Error({ error }: Props) {
  return <ErrorDisplay title={error.name} />;
}
