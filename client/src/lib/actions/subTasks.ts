import { toast } from "sonner";
import { axiosInstance } from "@/api/api";
import { mutateTasks } from "./tasks";
import { getErrorMesage } from "../utils";
import { mutateTeam } from "./teams";

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
    return {
      success: false,
      message: getErrorMesage(error),
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
    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export const handleDeleteSubTask = async (
  subTaskId: ISubTask["_id"],
  taskId: ITask["_id"],
  projectId: ITask["projectId"],
  teamId: ITask["team_id"]
) => {
  try {
    const { message, success } = await deleteSubTasks(subTaskId, taskId);

    const options = success ? "success" : "error";

    if (success) {
      mutateTasks(taskId, projectId);

      if (teamId) {
        mutateTeam(teamId);
      }
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

      if (task.team_id) {
        mutateTeam(task.team_id);
      }
    }
    toast[options](message);
  } catch (error) {
    toast.error(error instanceof Error ? error.message : "internal error");
  }
};
