"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Edit, Plus, Search } from "lucide-react";
import { EditProjectDialog } from "@/components/edit-project-dialog";
import { CreateTaskDialog } from "@/components/create-task-dialog";
import { useProject } from "@/hooks/use-projects";

import ProjectHeader from "./Header";
import ProjectTasksTable from "./Tasks";

export default function ProjectContainer({ projectId }: { projectId: string }) {
  const [editingProject, setEditingProject] = useState(false);
  const [showCreateTask, setShowCreateTask] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");

  const {
    data: projectData,
    isLoading: projectLoading,
    error,
  } = useProject(projectId);

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

  if (projectLoading) {
    return <Loader />;
  }

  const { project } = projectData!;

  return (
    <>
      <div className="max-w-7xl mx-auto space-y-8">
        {/* Header */}
        <div className="flex sm:items-center sm:flex-row flex-col gap-5 justify-between">
          <div>
            <h1 className="text-3xl font-bold text-foreground">
              Project: {project.name}
            </h1>
            <p className="text-muted-foreground mt-1">
              {project.description ||
                "Manage all tasks related to this project."}
            </p>
          </div>
          <Button
            className="w-fit"
            onClick={() => setEditingProject(true)}
            variant="outline"
          >
            <Edit className="h-4 w-4 mr-2" />
            Edit Project
          </Button>
        </div>

        {/* Project Details */}
        <ProjectHeader project={project} />

        {/* Tasks Section */}
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <h2 className="text-lg font-semibold text-foreground">Tasks</h2>
            <Button onClick={() => setShowCreateTask(true)}>
              <Plus className="h-4 w-4 mr-2" />
              Add Task
            </Button>
          </div>

          {/* Search */}
          <div className="relative max-w-md">
            <Search className="absolute left-3 top-3 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Search tasks..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="pl-10"
            />
          </div>

          {/* Tasks Table */}
          <ProjectTasksTable searchTerm={searchTerm} projectId={projectId} />
        </div>
      </div>

      {editingProject && (
        <EditProjectDialog
          project={project}
          open={editingProject}
          onOpenChange={setEditingProject}
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
