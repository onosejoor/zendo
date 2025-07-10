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
    const { data } = await axiosInstance.put<APIRes>(`/tasks/${task._id}`, {
      ...task,
      dueDate: new Date(task.dueDate!),
    });
    return { success: data.success, message: data.message };
  } catch (error) {
    if (isAxiosError(error)) {
      return {
        success: false,
        message: error.response?.data.message || error.response?.data,
      };
    }
    return {
      success: false,
      message: error instanceof Error ? error.message : "Internal error",
    };
  }
}

export async function deleteTask(id: ITask["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(`/tasks/${id}`);
    return { success: data.success, message: data.message };
  } catch (error) {
    if (isAxiosError(error)) {
      return {
        success: false,
        message: error.response?.data.message || error.response?.data,
      };
    }
    return {
      success: false,
      message: error instanceof Error ? error.message : "Internal error",
    };
  }
}

export const handleDeleteTask = async (taskId: ITask["_id"]) => {
  try {
    const { success, message } = await deleteTask(taskId);
    const options = success ? "success" : "error";

    toast[options](message);
    if (success) {
      mutateTasks(taskId);
      window.location.href = "/dashboard/tasks";
    }
  } catch (error) {
    console.error("Failed to delete task:", error);
    toast.error(error instanceof Error ? error.message : "internal error");
  }
};

export const handleToggleTask = async (task: ITask) => {
  try {
    const newStatus = task.status === "completed" ? "pending" : "completed";

    const newTask = { ...task, status: newStatus };

    const { message, success } = await updateTask(newTask as ITask);

    const options = success ? "success" : "error";

    if (success) {
      mutateTasks(task._id, task.projectId);
    }
    toast[options](message);
  } catch (error) {
    toast.error(error instanceof Error ? error.message : "internal error");
  }
};

export function mutateTasks(taskId?: string, projectId?: string) {
  mutate(`/tasks/${taskId}`);
  mutate(`/task`);
  mutate("/stats");
  mutate(`/projects/${projectId}/tasks`);
}
