import { fetcher } from "@/lib/utils";
import useSWR from "swr";

export function useTasks(limit: number = 10, page: number = 1) {
  return useSWR<{ success: boolean; tasks: ITask[] }>(
    `/tasks?limit=${limit}&page=${page}`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}

export function useTask(id: string) {
  return useSWR<{ success: boolean; task: ITask }>(
    id ? `/tasks/${id}` : null,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}
