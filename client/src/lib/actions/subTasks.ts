import { toast } from "sonner";
import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";
import { mutateTasks } from "./tasks";

export async function toggleSubTasks(
  subTaskId: ISubTask["_id"],
  id: ITask["_id"],
  newStatus: boolean
) {
  try {
    const { data } = await axiosInstance.put<APIRes>(
      `/tasks/${id}/subtask/${subTaskId}`,
      { completed: newStatus }
    );
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

export async function deleteSubTasks(
  subTaskId: ISubTask["_id"],
  id: ITask["_id"]
) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(
      `/tasks/${id}/subtask/${subTaskId}`
    );
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

export const handleDeleteSubTask = async (
  subTaskId: ISubTask["_id"],
  taskId: ITask["_id"],
  projectId: ITask["projectId"]
) => {
  try {
    const { message, success } = await deleteSubTasks(subTaskId, taskId);

    const options = success ? "success" : "error";

    if (success) {
      mutateTasks(taskId, projectId);
    }
    toast[options](message);
  } catch (error) {
    toast.error(error instanceof Error ? error.message : "internal error");
  }
};

export const handleToggleSubTask = async (
  subTaskId: ISubTask["_id"],
  task: ITask
) => {
  const subTask = task.subTasks?.find((f) => f._id === subTaskId);

  try {
    const newStatus = subTask?.completed ? false : true;

    const { message, success } = await toggleSubTasks(
      subTaskId,
      task._id,
      newStatus
    );

    const options = success ? "success" : "error";

    if (success) {
      mutateTasks(task._id, task.projectId);
    }
    toast[options](message);
  } catch (error) {
    toast.error(error instanceof Error ? error.message : "internal error");
  }
};
