import { axiosInstance } from "@/api/api";
import { toast } from "sonner";
import { mutate } from "swr";
import { getErrorMesage } from "../utils";
import dayjs from "dayjs";
import { returnAssigneeId } from "../functions";

type APIResponse = APIRes & {
  taskId: string;
};

export async function createTeamTask(
  formData: Partial<ITask>,
  subTasks: ISubTask[]
) {
  try {
    const { data } = await axiosInstance.post<APIResponse>("/tasks/new", {
      ...formData,
      dueDate: dayjs(formData.dueDate).format(),
      subTasks,
      assignees: returnAssigneeId(formData.assignees || []),
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
