"use client";

import { fetcher } from "@/lib/utils";
import useSWR from "swr";

export function useUser() {
  return useSWR<{
    success: boolean;
    user: IUser;
  }>("/auth/user", fetcher, {
    revalidateOnFocus: false,
    shouldRetryOnError: false,
  });
}
