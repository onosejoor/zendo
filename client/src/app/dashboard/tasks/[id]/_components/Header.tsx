import {
  formatDate,
  getStatusColor,
} from "@/app/dashboard/tasks/_components/constants";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { cn } from "@/lib/utils";
import { Activity, Calendar, FileText, Timer } from "lucide-react";

export default function TaskHeader({ task }: { task: ITask }) {
  return (
    <Card>
      <CardContent className="p-6">
        <h2 className="text-lg font-semibold text-foreground mb-6">
          Project Details
        </h2>
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

            <div className="flex items-center space-x-3">
              <div
                className={cn(
                  `size-3 rounded-full`,
                  getStatusColor(task.status)
                )}
              />
              <Badge variant="outline">{task.status}</Badge>
            </div>

            {/* <div className="flex items-start space-x-3">
              <ClipboardList className="h-5 w-5 text-muted-foreground mt-0.5" />
              <div>
                <p className="text-sm font-medium text-muted-foreground">
                  Total Tasks
                </p>
                <p className="text-foreground">{task.totalTasks}</p>
              </div>
            </div> */}
          </div>

          <div className="space-y-4">
            <div className="flex items-start space-x-3">
              <Activity className="h-5 w-5 text-muted-foreground mt-0.5" />
              <div>
                <p className="text-sm font-medium text-muted-foreground">
                  Status
                </p>
                <Badge className="bg-accent text-blue-500 border-0 font-medium">
                  Active
                </Badge>
              </div>
            </div>

            <div className="flex items-center space-x-3 text-sm text-gray-500">
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
