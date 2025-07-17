import { formatDate } from "@/app/dashboard/tasks/_components/constants";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { toggleSubTasks } from "@/lib/actions/subTasks";
import { mutateTasks, updateTask } from "@/lib/actions/tasks";
import { getStatusBadge } from "@/lib/functions";
import { checkExpired, cn } from "@/lib/utils";
import {
  Activity,
  Calendar,
  CheckSquare,
  FileText,
  Link2Icon,
  LinkIcon,
  Timer,
  X,
} from "lucide-react";
import Link from "next/link";
import { toast } from "sonner";

export default function TaskHeader({ task }: { task: ITask }) {
  const isExpired = checkExpired(task.dueDate);

  const handleToggleTask = async (task: ITask) => {
    try {
      const newStatus = task.status === "completed" ? "pending" : "completed";

      const newTask = { ...task, status: newStatus };

      const { message, success } = await updateTask(newTask as ITask);

      const options = success ? "success" : "error";

      if (success) {
        mutateTasks(task._id, task.projectId);
      }
      toast[options](message);
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "internal error");
    }
  };

  const handleToggleSubTask = async (index: number) => {
    try {
      const newStatus = task.subTasks?.[index].completed ? false : true;

      const { message, success } = await toggleSubTasks(
        index,
        task._id,
        newStatus
      );

      const options = success ? "success" : "error";

      if (success) {
        mutateTasks(task._id, task.projectId);
      }
      toast[options](message);
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "internal error");
    }
  };

  const description =
    task.description.length > 250
      ? task.description.slice(0, 250) + "..."
      : task.description;

  return (
    <>
      <Card className="relative">
        <CardContent className={"p-6"}>
          {isExpired && task.status !== "completed" && (
            <Badge className="absolute bg-red-500 text-white -rotate-40 -left-3 top-1">
              Expired
            </Badge>
          )}
          <div className="flex space-x-5 items-center mb-6">
            <input
              type="checkbox"
              checked={task.status === "completed"}
              onChange={() => handleToggleTask(task)}
              className="size-4 text-blue-600 rounded border-gray-300 focus:ring-accent-blue"
            />
            <h2 className="text-lg font-semibold h-fit text-foreground">
              Task Details
            </h2>
          </div>

          <div
            className={cn(
              "grid grid-cols-1 md:grid-cols-2 gap-6",
              task.status === "completed" && "**:!text-gray-400"
            )}
          >
            <div className="space-y-4">
              <div className="flex items-start space-x-3">
                <FileText className="h-5 w-5 shrink-0 text-muted-foreground mt-0.5" />
                <div>
                  <p className="text-sm font-medium text-muted-foreground">
                    Title
                  </p>
                  <p className="text-foreground line-clamp-1 ">{task.title}</p>
                </div>
              </div>

              <div className="flex items-start space-x-3">
                <FileText className="h-5 shrink-0 w-5 text-muted-foreground mt-0.5" />
                <div>
                  <p className="text-sm font-medium text-muted-foreground">
                    Description
                  </p>
                  <p className="text-foreground">
                    {description || "No description provided"}
                  </p>
                </div>
              </div>

              <div className="flex items-start space-x-3">
                <Activity className="h-5 w-5  shrink-0 text-muted-foreground mt-0.5" />
                <div>
                  <p className="text-sm font-medium text-muted-foreground">
                    Status
                  </p>
                  {getStatusBadge(task.status, task.dueDate)}
                </div>
              </div>
            </div>

            <div className="space-y-4">
              <div className="flex items-start space-x-3 text-sm text-gray-500">
                <Timer className="size-5" />
                <div>
                  <p className="text-sm font-medium text-muted-foreground">
                    Due Date:
                  </p>
                  <p className="text-foreground">{formatDate(task.dueDate)}</p>
                </div>
              </div>

              <div className="flex items-start space-x-3">
                <Calendar className="h-5 w-5 text-muted-foreground mt-0.5" />
                <div>
                  <p className="text-sm font-medium text-muted-foreground">
                    Created Date
                  </p>
                  <p className="text-foreground">
                    {new Date(task.created_at).toDateString()}
                  </p>
                </div>
              </div>
              {task.projectId && (
                <div className="flex items-start space-x-3">
                  <LinkIcon className="h-5 w-5 text-muted-foreground mt-0.5" />
                  <div>
                    <p className="text-sm font-medium text-muted-foreground">
                      View Project
                    </p>
                    <Link href={`/dashboard/projects/${task.projectId}`}>
                      <Button
                        variant={"outline"}
                        className="flex gap-3 items-center my-3 text-gray-500"
                      >
                        <Link2Icon /> View Project
                      </Button>
                    </Link>
                  </div>
                </div>
              )}
            </div>
          </div>
        </CardContent>
      </Card>

      {task.subTasks && task.subTasks.length > 0 && (
        <Card>
          <CardContent className="p-6 space-y-3">
            <h3 className="font-semibold">SubTasks</h3>

            <div className="grid gap-5">
              {task.subTasks.map((subTask, i) => (
                <div
                  key={i}
                  className="flex animate-in animation-duration-[200ms] items-center gap-2 group"
                >
                  <button
                    type="button"
                    onClick={() => handleToggleSubTask(i)}
                    className="flex-shrink-0"
                  >
                    <CheckSquare
                      className={`h-4 w-4 ${
                        subTask.completed
                          ? "text-green-600 fill-green-100"
                          : "text-gray-400"
                      }`}
                    />
                  </button>
                  <Input
                    readOnly
                    value={subTask.title}
                    className={cn(
                      `flex-1 h-8`,
                      subTask.completed ? "line-through text-gray-500" : ""
                    )}
                  />
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    // onClick={}
                    className="transition-opacity flex-shrink-0 h-8 w-8 p-0"
                  >
                    <X className="h-3 w-3" />
                  </Button>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      )}
    </>
  );
}
