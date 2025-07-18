import { Dispatch, SetStateAction } from "react";
import { generateId, getTextNewLength } from "../functions";
import { toast } from "sonner";

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
      subTasks: [
        { _id: generateId(), title: trimmedTitle, completed: false },
        ...(prev.subTasks || []),
      ],
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

export const toggleSubTask = (
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
