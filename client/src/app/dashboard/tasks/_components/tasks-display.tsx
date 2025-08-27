"use client";

import { useCallback, useEffect, useState } from "react";
import TaskCard from "./TaskCards";
import { toast } from "sonner";
import { checkExpired, getErrorMesage } from "@/lib/utils";
import { searchForTasks } from "@/lib/actions/tasks";

type Props = {
  initialTasks: ITask[];
  searchTerm: string;
  filter: Status | "all" | "expired";
};

export default function TasksDisplay({
  initialTasks,
  searchTerm,
  filter,
}: Props) {
  const [remoteTasks, setRemoteTasks] = useState<ITask[] | null>(null);

  const getFilteredTasks = useCallback(() => {
    if (!searchTerm.trim()) return initialTasks;

    if (filter === "expired") {
      return initialTasks.filter(
        (task) =>
          (task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
            (task.description &&
              task.description
                .toLowerCase()
                .includes(searchTerm.toLowerCase()))) &&
          checkExpired(task.dueDate) &&
          task.status !== "completed"
      );
    }

    if (filter === "all") {
      return initialTasks.filter(
        (task) =>
          task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
          (task.description &&
            task.description.toLowerCase().includes(searchTerm.toLowerCase()))
      );
    }

    return initialTasks.filter(
      (task) =>
        (task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
          (task.description &&
            task.description
              .toLowerCase()
              .includes(searchTerm.toLowerCase()))) &&
        task.status === filter &&
        (task.status === "completed" || !checkExpired(task.dueDate))
    );
  }, [initialTasks, searchTerm, filter]);

  useEffect(() => {
    if (!searchTerm.trim()) {
      setRemoteTasks(null);
      return;
    }

    const filteredTasks = getFilteredTasks();
    if (filteredTasks.length > 0) {
      setRemoteTasks(null);
      return;
    }

    // If no local match, search backend
    const fetchTasks = async () => {
      try {
        const res = await searchForTasks(searchTerm);
        if (res.success && res.tasks) {
          setRemoteTasks(res.tasks);
        } else {
          setRemoteTasks([]);
        }
      } catch (error) {
        toast.error(getErrorMesage(error));
        setRemoteTasks([]);
      }
    };
    fetchTasks();
  }, [searchTerm, initialTasks, getFilteredTasks]);

  const displayTasks = remoteTasks ? remoteTasks : getFilteredTasks();

  return (
    <div className="grid gap-5 md:grid-cols-2 grid-cols-1">
      {displayTasks.map((task) => (
        <TaskCard key={task._id} task={task} />
      ))}
    </div>
  );
}
