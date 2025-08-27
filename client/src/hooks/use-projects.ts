import { fetcher } from "@/lib/utils";
import useSWR from "swr";

export function useProjects(limit: number = 10, page: number = 1) {
  return useSWR<{ success: boolean; projects: IProject[] }>(
    `/projects?limit=${limit}&page=${page}`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}

export function useProject(id: string) {
  return useSWR<{ success: boolean; project: IProject }>(
    id ? `/projects/${id}` : null,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}

export function useProjectTasks(projectId: string) {
  return useSWR<{success: boolean, tasks: ITask[]}>(projectId ? `/projects/${projectId}/tasks` : null, fetcher, {
    revalidateOnFocus: false,
  });
}
