import { axiosInstance } from "@/api/api";
import { mutate } from "swr";
import { toast } from "sonner";
import { getErrorMesage } from "../utils";

type CreateProps = {
  name: string;
  description: string;
};

export async function createTeam(team: CreateProps) {
  try {
    const { data } = await axiosInstance.post<APIRes & { teamId: string }>(
      `/teams/new`,
      team
    );
    return { success: data.success, message: data.message, id: data.teamId };
  } catch (error) {
    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export async function updateTeam(team: Partial<ITeam>) {
  try {
    const { data } = await axiosInstance.put<APIRes>(
      `/teams/${team._id}`,
      team
    );
    return { success: data.success, message: data.message };
  } catch (error) {
    console.log(error);

    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export async function deleteTeam(id: ITeam["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(`/teams/${id}`);
    if (data.success) {
      mutateTeam(id);
    }
    return { success: data.success, message: data.message };
  } catch (error) {
    console.log("Error deleting team: ", error);

    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export const handleDeleteTeam = async (teamId: ITeam["_id"]) => {
  const { success, message } = await deleteTeam(teamId);
  const options = success ? "success" : "error";

  toast[options](message);

  if (success) {
    mutateTeam(teamId);
  }
};

export async function searchForTeams(search: string) {
  try {
    const { data } = await axiosInstance.get<APIRes & { teams: ITeam[] }>(
      `/teams/search?search=${search}`
    );
    return data;
  } catch (error) {
    return { success: false, message: getErrorMesage(error), teams: [] };
  }
}

export function mutateTeam(teamId?: string) {
  mutate(
    (key) => typeof key === "string" && key.startsWith(`/teams/${teamId}/tasks`)
  );
  mutate((key) => typeof key === "string" && key.startsWith("/teams?"));
  mutate("/stats");
}
