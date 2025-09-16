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

type TeamMemberData = { members: IMember[]; role: TeamRole; userId: string };
export function useTeamMembers(teamId: string) {
  return useSWR<{
    success: boolean;
    data: TeamMemberData;
  }>(`/teams/${teamId}/members`, fetcher, {
    revalidateOnFocus: false,
  });
}

export function useTeam(id: string) {
  return useSWR<{ success: boolean; team: ITeam }>(`/teams/${id}`, fetcher, {
    revalidateOnFocus: false,
  });
}

export type TeamStat = {
  total_tasks: number;
  total_team_members: number;
  total_pending_invites?: number;
  role: TeamRole;
};
export function useTeamStats(id: string) {
  return useSWR<{ success: boolean; stats: TeamStat }>(
    `/teams/${id}/stats`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}

type AllTeamStatsRes = {
  number_of_teams: number;
  number_of_tasks_assigned_to_me: number;
  number_of_tasks_due_today: number;
};
export function useAllTeamStats() {
  return useSWR<{ success: boolean; stat: AllTeamStatsRes }>(
    `/teams/stats`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
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

export function useTeamTask(teamId: string, taskId: string) {
  return useSWR<{ success: boolean; data: { task: ITask; role: TeamRole } }>(
    `/teams/${teamId}/tasks/${taskId}`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}

type TeamInviteRes = {
  _id: string;
  email: string;
  expiresAt: Date;
  status: "sent" | "pending" | "failed";
  createdAt: Date;
};
export function useTeamInvites(teamId: string) {
  return useSWR<{ success: boolean; invitees: TeamInviteRes[] }>(
    `/teams/${teamId}/invites`,
    fetcher,
    {
      revalidateOnFocus: false,
    }
  );
}
