import { formatDate } from "@/app/dashboard/tasks/_components/constants";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { updateTask } from "@/lib/actions/tasks";
import { getStatusBadge } from "@/lib/functions";
import { checkExpired } from "@/lib/utils";
import { Activity, Calendar, FileText, Timer } from "lucide-react";
import { toast } from "sonner";
import { mutate } from "swr";

export default function TaskHeader({ task }: { task: ITask }) {
  const isExpired = checkExpired(task.dueDate);

  const handleToggleTask = async (task: ITask) => {
    try {
      const newStatus = task.status === "completed" ? "pending" : "completed";

      const newTask = { ...task, status: newStatus };

      const { message, success } = await updateTask(newTask as ITask);

      const options = success ? "success" : "error";

      if (success) {
        mutate(`/task/${task._id}`);
      }
      toast[options](message);
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "internal error");
    }
  };

  return (
    <Card className="relative">
      <CardContent className="p-6">
        {isExpired && (
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
            Project Details
          </h2>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="space-y-4">
            <div className="flex items-start space-x-3">
              <FileText className="h-5 w-5 text-muted-foreground mt-0.5" />
              <div>
                <p className="text-sm font-medium text-muted-foreground">
                  Title
                </p>
                <p className="text-foreground">{task.title}</p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <FileText className="h-5 w-5 text-muted-foreground mt-0.5" />
              <div>
                <p className="text-sm font-medium text-muted-foreground">
                  Description
                </p>
                <p className="text-foreground">
                  {task.description || "No description provided"}
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <Activity className="h-5 w-5 text-muted-foreground mt-0.5" />
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
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
