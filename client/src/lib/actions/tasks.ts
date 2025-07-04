import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";
import { toast } from "sonner";
import { mutate } from "swr";

type APIRes = {
  success: boolean;
  message: string;
};

export async function updateTask(task: Partial<ITask>) {
  try {
    const { data } = await axiosInstance.put<APIRes>(`/task/${task._id}`, {
      ...task,
      dueDate: new Date(task.dueDate!),
    });
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

export const handleDeleteTask = async (taskId: ITask["_id"]) => {
  if (window.confirm("Are you sure you want to delete this task?")) {
    try {
      const { success, message } = await deleteTask(taskId);

      if (success) {
        mutate(`/task/${taskId}`);
      }
      const options = success ? "success" : "error";

      toast[options](message);
    } catch (error) {
      console.error("Failed to delete task:", error);
      toast.error(error instanceof Error ? error.message : "internal error");
    }
  }
};


