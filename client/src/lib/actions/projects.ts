
import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";

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
    const { data } = await axiosInstance.post<APIRes>(`/projects/new`, project);
    return { success: data.success, message: data.message };
  } catch (error) {
    if (isAxiosError(error)) {
      return { success: false, message: error.response?.data.message };
    }
    return {
      success: false,
      message: error instanceof Error ? error.message : "Internal error",
    };
  }
}

export async function deleteTask(id: ITask["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(`/task/${id}`);
    return { success: data.success, message: data.message };
  } catch (error) {
    if (isAxiosError(error)) {
      return { success: false, message: error.response?.data.message };
    }
    return {
      success: false,
      message: error instanceof Error ? error.message : "Internal error",
    };
  }
}
