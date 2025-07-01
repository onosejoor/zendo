"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";
import { Plus, Search, Filter } from "lucide-react";

import { CreateTaskDialog } from "@/components/create-task-dialog";
import { EditTaskDialog } from "@/components/edit-task-dialog";
import { useTasks } from "@/hooks/use-tasks";
import { mutate } from "swr";
import { deleteTask } from "@/lib/actions/tasks";
import { toast } from "sonner";
import TaskCard from "./TaskCards";

export default function TasksPage() {
  const [showCreateTask, setShowCreateTask] = useState(false);
  const [editingTask, setEditingTask] = useState<ITask | null>(null);
  const [searchTerm, setSearchTerm] = useState("");

  const { data, isLoading, error } = useTasks();

  const handleEditTask = (task: ITask) => {
    setEditingTask(task);
  };

  const handleDeleteTask = async (taskId: ITask["_id"]) => {
    if (window.confirm("Are you sure you want to delete this task?")) {
      try {
        const { success, message } = await deleteTask(taskId);

        if (success) {
          mutate("/task");
        }
        const options = success ? "success" : "error";

        toast[options](message);
      } catch (error) {
        console.error("Failed to delete task:", error);
        toast.error(error instanceof Error ? error.message : "internal error");
      }
    }
  };

  if (error) {
    return (
      <div className="text-center py-8">
        <p className="text-red-600">Failed to load tasks. Please try again.</p>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="space-y-4">
        {[...Array(5)].map((_, i) => (
          <Card key={i}>
            <CardContent className="p-6">
              <div className="h-6 bg-gray-200 rounded animate-pulse mb-2" />
              <div className="h-4 bg-gray-100 rounded animate-pulse w-3/4" />
            </CardContent>
          </Card>
        ))}
      </div>
    );
  }

  const { tasks } = data!;

  if (tasks.length < 1) {
    return (
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
          <Button onClick={() => setShowCreateTask(true)}>
            <Plus className="h-4 w-4 mr-2" />
            Create Task
          </Button>
        </CardContent>
      </Card>
    );
  }

  const filteredTasks = tasks.filter(
    (task) =>
      task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      task.description.toLowerCase().includes(searchTerm.toLowerCase())
  );

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
          <Button onClick={() => setShowCreateTask(true)}>
            <Plus className="h-4 w-4 mr-2" />
            New Task
          </Button>
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
          <Button variant="outline">
            <Filter className="h-4 w-4 mr-2" />
            Filter
          </Button>
        </div>

        {/* Tasks List */}
        <div className="grid gap-5 md:grid-cols-2 grid-cols-1">
          {filteredTasks.map((task) => (
            <TaskCard
              key={task._id}
              task={task}
              handleDeleteTask={handleDeleteTask}
              handleEditTask={handleEditTask}
            />
          ))}
        </div>
      </div>

      <CreateTaskDialog
        open={showCreateTask}
        onOpenChange={setShowCreateTask}
      />
      {editingTask && (
        <EditTaskDialog
          task={editingTask}
          open={!!editingTask}
          onOpenChange={(open) => !open && setEditingTask(null)}
        />
      )}
    </>
  );
}
