import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";
import { mutate } from "swr";
import { toast } from "sonner";

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
    if (isAxiosError(error)) {
      return { success: false, message: error.response?.data.message };
    }
    return {
      success: false,
      message: error instanceof Error ? error.message : "Internal error",
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

export async function deleteProject(id: IProject["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(`/projects/${id}`);
    if (data.success) {
      mutate("/projects");
      mutate(`/projects/${id}`);
    }
    return { success: data.success, message: data.message };
  } catch (error) {
    console.log("Error deleting task: ", error);

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

export const handleDeleteProject = async (projectId: IProject["_id"]) => {
  if (
    window.confirm(
      "Are you sure you want to delete this project and all tasks in it?"
    )
  ) {
    try {
      const { success, message } = await deleteProject(projectId);
      const options = success ? "success" : "error";

      toast[options](message);

      if (success) {
        window.location.href = "/dashboard/projects"
      }
    } catch (error) {
      toast.error(
        error instanceof Error ? error.message : "Internal server error"
      );

      console.error("Failed to delete project:", error);
    }
  }
};
