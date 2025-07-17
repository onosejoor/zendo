import { Dispatch, SetStateAction } from "react";
import { getTextNewLength } from "../functions";
import { toast } from "sonner";
import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";

export type SubTaskProps = {
  setFormData: Dispatch<SetStateAction<Partial<ITask>>>;
  newTask: string;
  index: number;
  subTasks: ISubTask[];
};

export const addSubTask = (
  newTask: SubTaskProps["newTask"],
  setFormData: SubTaskProps["setFormData"]
) => {
  if (!newTask.trim()) return;
  const { value: trimmedTitle, isLong } = getTextNewLength({
    id: "subtask",
    value: newTask.trim(),
  });

  if (isLong) {
    toast.error("Subtask title is too long, was shrinked to 70 characters");
  }

  setFormData((prev) => {
    return {
      ...prev,
      subTasks: [{ title: trimmedTitle, completed: false }, ...prev.subTasks!],
    };
  });
};

export const handleRemoveSubTask = ({
  index,
  subTasks,
  setFormData,
}: Omit<SubTaskProps, "newTask">) => {
  const newSubTasks = subTasks.filter((_, i) => i !== index);
  setFormData((prev) => {
    return {
      ...prev,
      subTasks: newSubTasks,
    };
  });
};

export const handleToggleSubTask = (
  setFormData: SubTaskProps["setFormData"],
  index: number
) => {
  setFormData((prev) => {
    return {
      ...prev,
      subTasks: prev.subTasks?.map((subtask, i) =>
        i === index ? { ...subtask, completed: !subtask.completed } : subtask
      ),
    };
  });
};

export async function toggleSubTasks(
  index: number,
  id: ITask["_id"],
  newStatus: boolean
) {
  try {
    const { data } = await axiosInstance.put<APIRes>(
      `/tasks/${id}/subtask/${index}`,
      { newStatus }
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
