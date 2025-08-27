"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";
import { Plus, Search } from "lucide-react";

import { CreateProjectDialog } from "@/components/dialogs/create-project-dialog";
import { useProjects } from "@/hooks/use-projects";

import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import PaginationBtn from "../../_components/pagination-btn";
import ProjectsDisplay from "./projects-display";
import useDebounce from "@/hooks/use-debounce";

export default function ProjectsContainer() {
  const [searchTerm, setSearchTerm] = useState("");
  const [page, setPage] = useState(1);

  const debouncedSearchTerm = useDebounce(searchTerm, 200);
  const { data, isLoading, error } = useProjects(5, page);

  if (error) {
    return (
      <ErrorDisplay message=" Failed to load projects. Please try again." />
    );
  }

  const { projects } = data || {};

  if (projects && projects.length < 1) {
    return (
      <div className="col-span-full">
        <Card>
          <CardContent className="p-12 text-center">
            <div className="text-gray-400 mb-4">
              <Plus className="h-12 w-12 mx-auto" />
            </div>
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              No projects found
            </h3>
            <p className="text-gray-600 mb-4">
              {searchTerm
                ? "Try adjusting your search terms"
                : "Get started by creating your first project"}
            </p>
            <CreateProjectDialog isVariant />
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Projects</h1>
            <p className="text-gray-600 mt-1">
              Organize your work into projects
            </p>
          </div>
          <CreateProjectDialog isVariant />
        </div>

        {/* Search and Filters */}
        <div className="flex space-x-4">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
            <Input
              placeholder="Search projects..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>

        {isLoading ? (
          <Loader text="Loading Projects..." />
        ) : (
          <>
            <ProjectsDisplay
              initialProjects={projects!}
              searchTerm={debouncedSearchTerm}
            />
            <PaginationBtn
              page={page}
              setPage={setPage}
              dataLength={projects!.length}
            />
          </>
        )}
      </div>
    </>
  );
}
