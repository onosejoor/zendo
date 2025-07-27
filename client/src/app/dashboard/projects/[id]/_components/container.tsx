"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { EditProjectDialog } from "@/components/dialogs/edit-project-dialog";
import { CreateTaskDialog } from "@/components/dialogs/create-task-dialog";
import { useProject } from "@/hooks/use-projects";

import ProjectHeader from "./Header";
import ProjectTasksTable from "./Tasks";
import BreadCrumbs from "@/components/BreadCrumbs";
import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import DeleteDataDialog from "@/components/dialogs/delete-data-dialog";
import { Search } from "lucide-react";

export default function ProjectContainer({ projectId }: { projectId: string }) {
  const [searchTerm, setSearchTerm] = useState("");

  const {
    data: projectData,
    isLoading: projectLoading,
    error,
  } = useProject(projectId);

  if (error) {
    if (error.status === 404) {
      return (
        <ErrorDisplay
          dontTryAgain
          title="Project not found"
          message="The project you're looking for doesn't exist."
        />
      );
    }
    return <ErrorDisplay />;
  }

  if (projectLoading) {
    return <Loader text="loading project" />;
  }

  const { project } = projectData!;

  return (
    <>
      <div className="max-w-7xl mx-auto space-y-8">
        <BreadCrumbs />

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
          <div className="space-x-3 flex">
            <EditProjectDialog project={project} />
            <DeleteDataDialog id={projectId} type="project" />
          </div>
        </div>

        {/* Project Details */}
        <ProjectHeader project={project} />

        {/* Tasks Section */}
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <h2 className="text-lg font-semibold text-foreground">Tasks</h2>
            <CreateTaskDialog />
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
    </>
  );
}
