"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";
import { Plus, Search } from "lucide-react";

import { CreateTaskDialog } from "@/components/dialogs/create-task-dialog";
import { useTasks } from "@/hooks/use-tasks";

import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";
import FilterDropdown from "../../_components/filter-dropdown";
import PaginationBtn from "../../_components/pagination-btn";
import TasksDisplay from "./tasks-display";
import useDebounce from "@/hooks/use-debounce";

export default function TasksPage() {
  const [page, setPage] = useState(1);
  const [searchTerm, setSearchTerm] = useState("");

  const [filter, setFilter] = useState<Status | "all" | "expired">("all");

  const { data, isLoading, error } = useTasks(5, page);
  const debouncedSearchTerm = useDebounce(searchTerm, 200);

  if (error) {
    return (
      <ErrorDisplay message="Error Loading tasks, check internet connection and try again" />
    );
  }

  const { tasks } = data || {};

  if (tasks && tasks.length < 1) {
    return (
      <>
        <Card>
          <CardContent className="p-12 text-center">
            <div className="text-gray-400 mb-4">
              <Plus className="h-12 w-12 mx-auto" />
            </div>
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              No tasks found
            </h3>
            <p className="text-gray-600 mb-4">
              {searchTerm
                ? "Try adjusting your search terms"
                : "Get started by creating your first task"}
            </p>
            <CreateTaskDialog />
          </CardContent>
        </Card>
      </>
    );
  }

  const onFilterChange = (s: Status) => setFilter(s);

  return (
    <>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Tasks</h1>
            <p className="text-gray-600 mt-1">
              Manage and track all your tasks
            </p>
          </div>
          <CreateTaskDialog />
        </div>

        <>
          {/* Search and Filters */}
          <div className="flex space-x-4">
            <div className="relative flex-1">
              <Search className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
              <Input
                placeholder="Search tasks..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="pl-10"
              />
            </div>
            <FilterDropdown
              currentFilter={filter}
              onFilterChange={onFilterChange}
            />
          </div>

          {/* Tasks List */}
          {isLoading ? (
            <Loader text="Loading Tasks..." />
          ) : (
            <>
              <TasksDisplay
                searchTerm={debouncedSearchTerm}
                initialTasks={tasks!}
                filter={filter}
              />
              <PaginationBtn
                page={page}
                setPage={setPage}
                dataLength={tasks!.length}
              />
            </>
          )}
        </>
      </div>
    </>
  );
}
