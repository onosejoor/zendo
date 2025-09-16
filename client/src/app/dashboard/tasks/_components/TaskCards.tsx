"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { handleToggleTask } from "@/lib/actions/tasks";
import { MoreHorizontal, Subtitles, Timer } from "lucide-react";
import { formatDate, getStatusColor } from "./constants";
import { Badge } from "@/components/ui/badge";
import { checkExpired, cn } from "@/lib/utils";
import Link from "next/link";
import { getStatusBadge } from "@/lib/functions";
import DeleteDataDialog from "@/components/dialogs/delete-data-dialog";
import { EditTaskDialog } from "@/components/dialogs/edit-task-dialog";

type Props = {
  task: ITask;
};

export default function TaskCard({ task }: Props) {
  const isExpired = checkExpired(task.dueDate);
  const isCompleted = task.status === "completed" || isExpired;

  return (
    <Link href={`/dashboard/tasks/${task._id}`}>
      <Card className="hover:shadow-md relative h-full !p-0 transition-shadow">
        <CardContent className="p-6">
          {isExpired && task.status !== "completed" && (
            <Badge className="absolute bg-red-500 text-white -rotate-40 -left-3 top-1">
              Expired
            </Badge>
          )}
          <div className="flex items-baseline justify-between">
            <div className="flex items-baseline space-x-4 flex-1">
              <div onClick={(e) => e.stopPropagation()}>
                <input
                  type="checkbox"
                  checked={task.status === "completed"}
                  onChange={() => handleToggleTask(task)}
                  className="h-4 w-4 text-blue-600 rounded border-gray-300 focus:ring-accent-blue"
                />
              </div>

              <div className="flex-1">
                <h3
                  className={cn(
                    `font-medium  line-clamp-1`,
                    isCompleted ? "line-through text-gray-500" : "text-gray-900"
                  )}
                >
                  {task.title}
                </h3>
                <p
                  className={`text-sm mt-1 line-clamp-3 ${
                    isCompleted ? "text-gray-400 " : "text-gray-600"
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
                  <div className="flex items-center space-x-1 text-sm text-gray-500">
                    <Subtitles className="size-3" />
                    <span>Subtasks: {task.subTasks?.length || 0}</span>
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
                <DropdownMenuItem asChild>
                  <EditTaskDialog isCard task={task} />
                </DropdownMenuItem>
                <DropdownMenuItem asChild>
                  <DeleteDataDialog card type="task" id={task._id} />
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}
