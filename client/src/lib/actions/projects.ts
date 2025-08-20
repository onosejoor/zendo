import { axiosInstance } from "@/api/api";
import { mutate } from "swr";
import { toast } from "sonner";
import { getErrorMesage } from "../utils";

type APIRes = {
  success: boolean;
  message: string;
};

type CreateProps = {
  name: string;
  description: string;
};

export async function createProject(project: CreateProps) {
  try {
    const { data } = await axiosInstance.post<APIRes & { projectId: string }>(
      `/projects/new`,
      project
    );
    return { success: data.success, message: data.message, id: data.projectId };
  } catch (error) {
    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export async function updateProject(project: Partial<IProject>) {
  try {
    const { data } = await axiosInstance.put<APIRes>(
      `/projects/${project._id}`,
      project
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

export async function deleteProject(id: IProject["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(`/projects/${id}`);
    if (data.success) {
      mutateProject(id);
    }
    return { success: data.success, message: data.message };
  } catch (error) {
    console.log("Error deleting project: ", error);

    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export const handleDeleteProject = async (projectId: IProject["_id"]) => {
  const { success, message } = await deleteProject(projectId);
  const options = success ? "success" : "error";

  toast[options](message);

  if (success) {
    mutateProject(projectId);
  }
};

export function mutateProject(projectId?: string) {
  mutate(`/projects/${projectId}`);
  mutate(`/projects`);
  mutate(`/projects/${projectId}/tasks`);
  mutate("/stats");
}
