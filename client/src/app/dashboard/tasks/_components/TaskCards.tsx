"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { handleDeleteTask, updateTask } from "@/lib/actions/tasks";
import { Edit, MoreHorizontal, Timer, Trash2 } from "lucide-react";
import { toast } from "sonner";
import { mutate } from "swr";
import { formatDate, getStatusColor } from "./constants";
import { Badge } from "@/components/ui/badge";
import { checkExpired, cn } from "@/lib/utils";
import Link from "next/link";
import { getStatusBadge } from "@/lib/functions";

type Props = {
  task: ITask;
  handleEditTask: (task: ITask) => void;
};

export default function TaskCard({ task, handleEditTask }: Props) {
  const handleToggleTask = async (task: ITask) => {
    try {
      const newStatus = task.status === "completed" ? "pending" : "completed";

      const newTask = { ...task, status: newStatus };

      const { message, success } = await updateTask(newTask as ITask);

      const options = success ? "success" : "error";

      if (success) {
        mutate("/task");
      }
      toast[options](message);
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "internal error");
    }
  };

  const isExpired = checkExpired(task.dueDate);
  const isCompleted = task.status === "completed" || isExpired;

  return (
    <Link href={`/dashboard/tasks/${task._id}`}>
      <Card
        key={task._id}
        className="hover:shadow-md relative h-full !p-0 transition-shadow"
      >
        <CardContent className="p-6">
          {isExpired && (
            <Badge className="absolute bg-red-500 text-white -rotate-40 -left-3 top-1">
              Expired
            </Badge>
          )}
          <div className="flex items-baseline justify-between">
            <div className="flex items-baseline space-x-4 flex-1">
              <input
                type="checkbox"
                checked={task.status === "completed"}
                onChange={() => handleToggleTask(task)}
                className="h-4 w-4 text-blue-600 rounded border-gray-300 focus:ring-accent-blue"
              />
              <div className="flex-1">
                <h3
                  className={cn(
                    `font-medium `,
                    isCompleted ? "line-through text-gray-500" : "text-gray-900"
                  )}
                >
                  {task.title}
                </h3>
                <p
                  className={`text-sm mt-1 ${
                    isCompleted ? "text-gray-400" : "text-gray-600"
                  }`}
                >
                  {task.description || "No Description"}
                </p>
                <div className=" space-y-3 mt-2">
                  <div className="flex items-center space-x-1">
                    <div
                      className={cn(
                        `size-3 rounded-full`,
                        getStatusColor(task.status, task.dueDate)
                      )}
                    />
                    {getStatusBadge(task.status, task.dueDate)}
                  </div>
                  <div className="flex items-center space-x-1 text-sm text-gray-500">
                    <Timer className="size-3" />
                    <span>Due: {formatDate(task.dueDate)}</span>
                  </div>
                </div>
              </div>
            </div>
            <DropdownMenu>
              <DropdownMenuTrigger onClick={(e) => e.stopPropagation()} asChild>
                <Button variant="ghost" size="sm">
                  <MoreHorizontal className="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent
                onClick={(e) => e.stopPropagation()}
                align="end"
              >
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
    </Link>
  );
}
