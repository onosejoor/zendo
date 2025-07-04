"use client";

import { useState } from "react";
import { CreateTaskDialog } from "@/components/dialogs/create-task-dialog";

import TaskHeader from "./Header";
import { useTask } from "@/hooks/use-tasks";
import { EditTaskDialog } from "@/components/dialogs/edit-task-dialog";
import { Button } from "@/components/ui/button";
import { Edit, Trash2 } from "lucide-react";
import { handleDeleteTask } from "@/lib/actions/tasks";
import BreadCrumbs from "@/components/BreadCrumbs";

export default function ProjectContainer({ projectId }: { projectId: string }) {
  const [showCreateTask, setShowCreateTask] = useState(false);
  const [editingTask, setEditingTask] = useState(false);

  const { data: taskData, isLoading, error } = useTask(projectId);

  if (error) {
    if (error.status === 404) {
      return (
        <div className="max-w-7xl mx-auto text-center py-12">
          <h1 className="text-2xl font-semibold text-foreground mb-2">
            Task not found
          </h1>
          <p className="text-muted-foreground">
            The task you&apos;re looking for doesn&apos;t exist.
          </p>
        </div>
      );
    }
    return (
      <div className="max-w-7xl mx-auto text-center py-12">
        <h1 className="text-2xl font-semibold text-foreground mb-2">
          An Error occured
        </h1>
        <p className="text-muted-foreground">try again</p>
      </div>
    );
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
            <h1 className="text-3xl font-bold text-foreground">
              Task: {task.title}
            </h1>
            <p className="text-muted-foreground mt-1">
              <span className="text-foreground"> Description:</span>{" "}
              {task.description || "No description"}
            </p>
          </div>
          <div className="space-x-3 flex">
            <Button
              className="w-fit"
              onClick={() => setEditingTask(true)}
              variant="outline"
            >
              <Edit className="h-4 w-4 mr-2" />
              Edit Task
            </Button>
            <Button
              onClick={() => handleDeleteTask(task._id)}
              className="text-red-600"
              variant={"outline"}
            >
              <Trash2 className="h-4 w-4 mr-2" />
              Delete
            </Button>
          </div>
        </div>

        {/* Project Details */}
        <TaskHeader task={task} />
      </div>

      {editingTask && (
        <EditTaskDialog
          task={task}
          open={editingTask}
          onOpenChange={setEditingTask}
        />
      )}

      <CreateTaskDialog
        open={showCreateTask}
        onOpenChange={setShowCreateTask}
      />
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
