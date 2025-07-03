"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { CreateTaskDialog } from "@/components/create-task-dialog";
import { CreateProjectDialog } from "@/components/create-project-dialog";
import { useTasks } from "@/hooks/use-tasks";
import { useProjects } from "@/hooks/use-projects";
import { useUser } from "@/hooks/use-user";
import RecentActivities from "./recent-activities";
import StatCards from "./Stats";

export default function Dashboard() {
  const [showCreateTask, setShowCreateTask] = useState(false);
  const [showCreateProject, setShowCreateProject] = useState(false);

  const { data, isLoading: userLoading } = useUser();
  const { data: taskData, isLoading: tasksLoading } = useTasks();
  const { data: projectData, isLoading: projectsLoading } = useProjects();

  if (tasksLoading || projectsLoading || userLoading) {
    return <p>loading...</p>;
  }

  const { user } = data!;
  const { tasks } = taskData!;
  const { projects } = projectData!;

  return (
    <>
      <div className="space-y-8">
        {/* Header */}
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">
              Welcome back,{" "}
              <span className="text-blue-500">{user?.username || "User"}!</span>
            </h1>
            <p className="text-gray-600 mt-1">
              Here&apos;s what&apos;s happening with your tasks today.
            </p>
          </div>
          <div className="space-x-2">
            <Button onClick={() => setShowCreateTask(true)}>
              <Plus className="h-4 w-4 mr-2" />
              New Task
            </Button>
            <Button
              variant="outline"
              onClick={() => setShowCreateProject(true)}
            >
              <Plus className="h-4 w-4 mr-2" />
              New Project
            </Button>
          </div>
        </div>

        {/* Stats Cards */}
        <StatCards />

        {/* Recent Tasks and Projects */}
        <RecentActivities tasks={tasks} projects={projects} />
      </div>

      <CreateTaskDialog
        open={showCreateTask}
        onOpenChange={setShowCreateTask}
      />
      <CreateProjectDialog
        open={showCreateProject}
        onOpenChange={setShowCreateProject}
      />
    </>
  );
}
