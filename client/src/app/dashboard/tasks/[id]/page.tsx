"use client";

import { useState } from "react";
import { CreateTaskDialog } from "@/components/create-task-dialog";

import TaskHeader from "./_components/Header";
import { useParams } from "next/navigation";
import { useTask } from "@/hooks/use-tasks";
import { EditTaskDialog } from "@/components/edit-task-dialog";
import { Button } from "@/components/ui/button";
import { Edit } from "lucide-react";

export default function ProjectContainer() {
  const params = useParams();
  const projectId = params.id as string;

  const [showCreateTask, setShowCreateTask] = useState(false);
  const [editingTask, setEditingTask] = useState(false);

  const { data: taskData, isLoading, error } = useTask(projectId);

  if (error) {
    if (error.status === 404) {
      return (
        <div className="max-w-7xl mx-auto text-center py-12">
          <h1 className="text-2xl font-semibold text-foreground mb-2">
            Project not found
          </h1>
          <p className="text-muted-foreground">
            The project you&apos;re looking for doesn&apos;t exist.
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
        {/* Header */}
        <div className="flex sm:items-center sm:flex-row flex-col gap-5 justify-between">
          <div>
            <h1 className="text-3xl font-bold text-foreground">
              Task: {task.title}
            </h1>
            <p className="text-muted-foreground mt-1">
            Description:  {task.description || "No description"}
            </p>
          </div>
          <Button
            className="w-fit"
            onClick={() => setEditingTask(true)}
            variant="outline"
          >
            <Edit className="h-4 w-4 mr-2" />
            Edit Task
          </Button>
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
