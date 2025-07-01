
import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";

type APIRes = {
  success: boolean;
  message: string;
};

export async function updateTask(task: Partial<ITask>) {
  try {
    const { data } = await axiosInstance.put<APIRes>(
      `/task/${task._id}`,
      JSON.stringify({ ...task })
    );
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
