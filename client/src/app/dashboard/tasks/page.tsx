"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import {
  Plus,
  Search,
  Filter,
  MoreHorizontal,
  Edit,
  Trash2,
  Calendar,
} from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { CreateTaskDialog } from "@/components/create-task-dialog";
import { EditTaskDialog } from "@/components/edit-task-dialog";
import { useTasks } from "@/hooks/use-tasks";
import { mutate } from "swr";
import { deleteTask, updateTask } from "@/lib/actions/tasks";
import { toast } from "sonner";

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
          mutate("tasks");
        }
        const options = success ? "success" : "error";

        toast[options](message);
      } catch (error) {
        console.error("Failed to delete task:", error);
        toast.error(error instanceof Error ? error.message : "internal error");
      }
    }
  };

  const handleToggleTask = async (task: ITask) => {
    try {
      const newStatus = task.status === "completed" ? "pending" : "completed";

      const newTask = { ...task, status: newStatus };

      const { message, success } = await updateTask(newTask);

      const options = success ? "success" : "error";

      if (success) {
        mutate("tasks");
      }
      toast[options](message);
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "internal error");
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString();
  };

  const getStatusColor = (status: ITask["status"]) => {
    switch (status) {
      case "completed":
        return "bg-green-500";
      case "in-progress":
        return "bg-blue-500";
      case "pending":
        return "bg-yellow-500";
      default:
        return "bg-gray-500";
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
    return <p>loading...</p>;
  }

  const { tasks } = data!;

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
        <div className="space-y-4">
          {isLoading ? (
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
          ) : filteredTasks.length > 0 ? (
            filteredTasks.map((task) => (
              <Card
                key={task._id}
                className="hover:shadow-md transition-shadow"
              >
                <CardContent className="p-6">
                  <div className="flex items-start justify-between">
                    <div className="flex items-start space-x-4 flex-1">
                      <input
                        type="checkbox"
                        checked={task.status === "completed"}
                        onChange={() => handleToggleTask(task)}
                        className="h-4 w-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500"
                      />
                      <div className="flex-1">
                        <h3
                          className={`font-medium ${
                            task.status === "completed"
                              ? "line-through text-gray-500"
                              : "text-gray-900"
                          }`}
                        >
                          {task.title}
                        </h3>
                        <p
                          className={`text-sm mt-1 ${
                            task.status === "completed"
                              ? "text-gray-400"
                              : "text-gray-600"
                          }`}
                        >
                          {task.description}
                        </p>
                        <div className="flex items-center space-x-2 mt-2">
                          <div className="flex items-center space-x-1">
                            <div
                              className={`w-2 h-2 rounded-full ${getStatusColor(
                                task.status
                              )}`}
                            />
                            <Badge variant="outline">{task.status}</Badge>
                          </div>
                          <div className="flex items-center space-x-1 text-sm text-gray-500">
                            <Calendar className="h-3 w-3" />
                            <span>{formatDate(task.dueDate)}</span>
                          </div>
                        </div>
                      </div>
                    </div>
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button variant="ghost" size="sm">
                          <MoreHorizontal className="h-4 w-4" />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuItem onClick={() => handleEditTask(task)}>
                          <Edit className="h-4 w-4 mr-2" />
                          Edit
                        </DropdownMenuItem>
                        <DropdownMenuItem
                          onClick={() => handleDeleteTask(task._id)}
                          className="text-red-600"
                        >
                          <Trash2 className="h-4 w-4 mr-2" />
                          Delete
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </div>
                </CardContent>
              </Card>
            ))
          ) : (
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
          )}
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
