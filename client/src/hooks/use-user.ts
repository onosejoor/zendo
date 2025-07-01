"use client";

import { fetcher } from "@/lib/utils";
import useSWR from "swr";

interface User {
  _id: string;
  username: string;
  email: string;
  avatar?: string;
}

export function useUser() {
  return useSWR<{ success: boolean; user: User }>("/auth/user", fetcher, {
    revalidateOnFocus: false,
    shouldRetryOnError: false,
  });
}
