"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal, Subtitles, Timer, Users2Icon } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { checkExpired, cn } from "@/lib/utils";
import Link from "next/link";
import { getStatusBadge } from "@/lib/functions";
import DeleteDataDialog from "@/components/dialogs/delete-data-dialog";

import {
  formatDate,
  getStatusColor,
} from "@/app/dashboard/tasks/_components/constants";
import { checkRolesMatch } from "../../actions";
import TeamTaskDialog from "@/components/dialogs/team-task-dialog";
import { handleToggleTask } from "@/lib/actions/tasks";

type Props = {
  task: ITask;
  userRole: TeamRole;
};

export default function TeamTaskCard({ task, userRole }: Props) {
  const isExpired = checkExpired(task.dueDate);
  const isCompleted = task.status === "completed" || isExpired;

  return (
    <Link href={`/dashboard/teams/${task.team_id}/tasks/${task._id}`}>
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

              <div
                className={cn(
                  "flex-1",
                  isCompleted && "*:line-through *:text-gray-500"
                )}
              >
                <h3
                  className={cn(`font-medium  line-clamp-1`, "text-gray-900")}
                >
                  {task.title}
                </h3>
                <p className={`text-sm mt-1 line-clamp-3 text-gray-600`}>
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
                  <div className="flex items-center space-x-1 text-sm text-gray-500">
                    <Users2Icon className="size-3" />
                    <span>
                      No of members assigned: {task.assignees?.length || 0}
                    </span>
                  </div>
                </div>
              </div>
            </div>
            {checkRolesMatch(userRole, ["owner", "admin"]) && <CrudDialog task={task} />}
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}

const CrudDialog = ({ task }: { task: ITask }) => (
  <DropdownMenu>
    <DropdownMenuTrigger onClick={(e) => e.stopPropagation()} asChild>
      <Button variant="ghost" size="sm">
        <MoreHorizontal className="h-4 w-4" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent onClick={(e) => e.stopPropagation()} align="end">
      <DropdownMenuItem asChild>
        <TeamTaskDialog
          isCard
          defaultTeamId={task.team_id!}
          initialData={task}
        />
      </DropdownMenuItem>
      <DropdownMenuItem asChild>
        <DeleteDataDialog card type="task" id={task._id} />
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
);
