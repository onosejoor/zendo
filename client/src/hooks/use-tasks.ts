import { fetcher } from "@/lib/utils";
import useSWR from "swr";

export function useTasks() {
  return useSWR<{ success: boolean; tasks: ITask[] }>(`/task`, fetcher, {
    revalidateOnFocus: false,
  });
}

export function useTask(id: string) {
  return useSWR<{ success: boolean; task: ITask }>(
    id ? `/task/${id}` : null,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}
