"use client";

import TaskHeader from "./Data";
import { useTask } from "@/hooks/use-tasks";
import { EditTaskDialog } from "@/components/dialogs/edit-task-dialog";

import BreadCrumbs from "@/components/BreadCrumbs";
import ErrorDisplay from "@/components/error-display";
import DeleteDataDialog from "@/components/dialogs/delete-data-dialog";

export default function TaskContainer({ taskId }: { taskId: string }) {
  const { data: taskData, isLoading, error } = useTask(taskId);

  if (error) {
    if (error.status === 404) {
      return (
        <ErrorDisplay
          dontTryAgain
          title="Task not found"
          message=" The task you're looking for doesn't exist."
        />
      );
    }
    return <ErrorDisplay />;
  }

  if (isLoading) {
    return <Loader />;
  }

  const { task } = taskData!;

  return (
    <>
      <div className="max-w-7xl mx-auto space-y-8">
        <BreadCrumbs />
        {/* Header */}
        <div className="flex sm:items-baseline sm:flex-row flex-col gap-5 justify-between">
          <div>
            <h1 className="text-2xl sm:text-3xl font-bold text-foreground">
              Task: {task.title}
            </h1>
            <p className="text-muted-foreground mt-1">
              <span className="text-foreground"> Description:</span>{" "}
              {task.description || "No description"}
            </p>
          </div>
          <div className="space-x-3 flex">
            <EditTaskDialog task={task} />
            <DeleteDataDialog id={taskId} type="task" />
          </div>
        </div>

        {/* Project Details */}
        <TaskHeader task={task} />
      </div>
    </>
  );
}

const Loader = () => (
  <div className="max-w-7xl mx-auto space-y-8">
    <div className="h-8 bg-muted rounded animate-pulse" />
    <div className="h-4 bg-muted rounded animate-pulse w-1/2" />
    <div className="space-y-4">
      <div className="h-6 bg-muted rounded animate-pulse w-1/4" />
      <div className="space-y-2">
        {[...Array(4)].map((_, i) => (
          <div key={i} className="h-4 bg-muted rounded animate-pulse" />
        ))}
      </div>
    </div>
  </div>
);
