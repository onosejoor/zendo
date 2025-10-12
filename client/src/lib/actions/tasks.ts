import { axiosInstance } from "@/api/api";
import { toast } from "sonner";
import { mutate } from "swr";
import { getErrorMesage } from "../utils";
import dayjs from "dayjs";
import { mutateTeam } from "./teams";
import { returnAssigneeId } from "../functions";

type APIResponse = APIRes & {
  taskId: string;
};

export async function createTask(
  formData: Partial<ITask>,
  subTasks: ISubTask[]
) {
  try {
    const { data } = await axiosInstance.post<APIResponse>("/tasks/new", {
      ...formData,
      dueDate: dayjs(formData.dueDate).format(),
      subTasks,
      ...(formData.projectId && { projectId: formData.projectId }),
    });

    return data;
  } catch (error) {
    return {
      taskId: "",
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export async function updateTask(task: Partial<ITask>) {
  const payload = {
    ...task,
    dueDate: new Date(task.dueDate!),
    assignees: returnAssigneeId(task.assignees || []),
  };

  try {
    const { data } = await axiosInstance.put<APIRes>(
      `/tasks/${task._id}`,
      payload
    );
    return { ...data, taskId: task._id };
  } catch (error) {
    return {
      success: false,
      message: getErrorMesage(error),
      taskId: task._id,
    };
  }
}

export async function deleteTask(id: ITask["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(`/tasks/${id}`);
    return { success: data.success, message: data.message };
  } catch (error) {
    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export const handleDeleteTask = async (taskId: ITask["_id"]) => {
  const { success, message } = await deleteTask(taskId);
  const options = success ? "success" : "error";

  toast[options](message);
  if (success) {
    mutateTasks(taskId);
  }
};

export const handleToggleTask = async (task: ITask) => {
  try {
    const newStatus = task.status === "completed" ? "pending" : "completed";

    const newTask = {
      ...task,
      status: newStatus,
      dueDate: dayjs(task.dueDate).format(),
    };

    const { message, success } = await updateTask(newTask as ITask);

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

export async function searchForTasks(search: string) {
  try {
    const { data } = await axiosInstance.get<APIRes & { tasks: ITask[] }>(
      `/tasks/search?search=${search}`
    );
    return data;
  } catch (error) {
    return { success: false, message: getErrorMesage(error), tasks: [] };
  }
}

export function mutateTasks(taskId?: string, projectId?: string) {
  mutate(`/tasks/${taskId}`);
  mutate((key) => typeof key === "string" && key.startsWith("/tasks?"));
  mutate("/stats");
  mutate(`/projects/${projectId}/tasks`);
}
