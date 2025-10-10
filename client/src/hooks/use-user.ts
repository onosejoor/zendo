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

type UseHomeDataProps = {
  success: boolean;
  data: { total_users: number; avatars: string[] };
};
export function useHomeData() {
  return useSWR<UseHomeDataProps>("/home-stats", fetcher, {
    revalidateOnFocus: false,
    shouldRetryOnError: false,
  });
}
