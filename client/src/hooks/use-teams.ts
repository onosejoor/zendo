import { fetcher } from "@/lib/utils";
import useSWR from "swr";

export function useTeams(limit: number = 10, page: number = 1) {
  return useSWR<{ success: boolean; teams: ITeam[] }>(
    `/teams?limit=${limit}&page=${page}`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}

export function useTeamMembers(teamId: string) {
  return useSWR<{
    success: boolean;
    data: { members: IMember[]; role: TeamRole };
  }>(`/teams/${teamId}/members`, fetcher, {
    revalidateOnFocus: false,
  });
}

export function useTeam(id: string) {
  return useSWR<{ success: boolean; team: ITeam }>(`/teams/${id}`, fetcher, {
    revalidateOnFocus: false,
  });
}

export function useTeamTasks(id: string, limit: number = 10, page: number = 1) {
  return useSWR<{ success: boolean; data: { tasks: ITask[]; role: TeamRole } }>(
    `/teams/${id}/tasks?limit=${limit}&page=${page}`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}
