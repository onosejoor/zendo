import { fetcher } from "@/lib/utils";
import useSWR from "swr";

export function useProjects() {
  return useSWR<{ success: boolean; projects: IProject[] }>(
    "/projects",
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
  return useSWR(projectId ? `/projects/${projectId}/tasks` : null, fetcher, {
    revalidateOnFocus: false,
  });
}
