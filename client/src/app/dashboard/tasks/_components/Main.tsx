"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";
import { Plus, Search, SearchX, X } from "lucide-react";

import { CreateTaskDialog } from "@/components/dialogs/create-task-dialog";
import { useTasks } from "@/hooks/use-tasks";
import TaskCard from "./TaskCards";
import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";
import FilterDropdown from "../../_components/filter-dropdown";
import { checkExpired } from "@/lib/utils";
import { Button } from "@/components/ui/button";

export default function TasksPage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [filter, setFilter] = useState<Status | "all" | "expired">("all");

  const { data, isLoading, error } = useTasks();

  if (error) {
    return (
      <ErrorDisplay message="Error Loading tasks, check internet connection and try again" />
    );
  }

  if (isLoading) {
    return <Loader text="Loading Tasks..." />;
  }

  const { tasks } = data!;

  if (tasks.length < 1) {
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

  const getFilteredTasks = () => {
    if (filter === "expired") {
      return tasks.filter(
        (task) =>
          (task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
            task?.description
              .toLowerCase()
              .includes(searchTerm.toLowerCase())) &&
          checkExpired(task.dueDate) &&
          task.status !== "completed"
      );
    }

    if (filter === "all") {
      return tasks.filter(
        (task) =>
          task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
          (task.description &&
            task.description.toLowerCase().includes(searchTerm.toLowerCase()))
      );
    }
    return tasks.filter(
      (task) =>
        (task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
          task?.description.toLowerCase().includes(searchTerm.toLowerCase())) &&
        task.status === filter && !checkExpired(task.dueDate)
    );
  };

  const onFilterChange = (s: Status) => setFilter(s);

  const filteredTasks = getFilteredTasks();

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

        {filteredTasks.length > 0 ? (
          <div className="grid gap-5 md:grid-cols-2 grid-cols-1">
            {filteredTasks.map((task) => (
              <TaskCard key={task._id} task={task} />
            ))}
          </div>
        ) : (
          <div className="p-5 grid place-items-center min-h-[50vh] w-full bg-white">
            <div className="space-y-5 grid items-center">
              <SearchX className="text-gray-500 size-12.5 w-fit mx-auto" />
              <p className="text-gray-600">No Task with the Search or filter</p>

              <Button
                onClick={() => {
                  setSearchTerm("");
                  setFilter("all");
                }}
                className="w-fit mx-auto"
              >
                <X /> Clear Filters
              </Button>
            </div>
          </div>
        )}
      </div>
    </>
  );
}
