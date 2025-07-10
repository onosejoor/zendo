import { axiosInstance } from "@/api/api";
import { mutateProject } from "./projects";
import { toast } from "sonner";
import { KeyedMutator } from "swr";
import { AvatarProps } from "@/app/dashboard/settings/_components/main";

export async function deleteAllProjects() {
  try {
    const { data } = await axiosInstance.delete<APIRes>("/projects/all");

    const options = data.success ? "success" : "error";
    toast[options](data.message);

    if (data.success) {
      mutateProject();
    }
  } catch (err) {
    toast.error(
      err instanceof Error
        ? err.message
        : "Failed to delete projects. Please try again."
    );
    console.log("Error deleting all projects: ", err);
  }
}

export async function deleteAllTasks() {
  try {
    const { data } = await axiosInstance.delete<APIRes>("/tasks/all");

    const options = data.success ? "success" : "error";
    toast[options](data.message);

    if (data.success) {
      mutateProject();
    }
  } catch (err) {
    toast.error(
      err instanceof Error
        ? err.message
        : "Failed to delete tasks. Please try again."
    );
    console.log("Error deleting all tasks: ", err);
  }
}

export async function updateUser(
  userData: Pick<IUser, "username" | "avatar"> & AvatarProps,
  mutate: KeyedMutator<{ success: boolean; user: IUser }>
) {

  try {
    const formData = new FormData();
    formData.append("username", userData.username!);
    formData.append("avatar", userData.avatar!);
    if (userData.avatarFile) formData.append("avatarFile", userData.avatarFile);

    const { data } = await axiosInstance.put<APIRes>("/auth/user", formData, {
      headers: {
        "Content-Type": "multipart/form-data"
      }
    });

    const options = data.success ? "success" : "error";
    toast[options](data.message);

    if (data.success) {
      mutate();
    
    }
  } catch (err) {
    toast.error(
      err instanceof Error
        ? err.message
        : "Failed to Update user data. Please try again."
    );
    console.log("Error updating user data: ", err);
  }
}
