"use client";

import { Card, CardContent } from "@/components/ui/card";
import { useProjectTasks } from "@/hooks/use-projects";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { FileText, Plus } from "lucide-react";

import { formatDate } from "@/app/dashboard/tasks/_components/constants";
import { CreateTaskDialog } from "@/components/dialogs/create-task-dialog";
import { useState } from "react";
import { getStatusBadge } from "@/lib/functions";

export default function ProjectTasksTable({
  projectId,
  searchTerm,
}: {
  searchTerm: string;
  projectId: IProject["_id"];
}) {
  const { data: taskData, isLoading, error } = useProjectTasks(projectId);
  const [showCreateTask, setShowCreateTask] = useState(false);

  if (error) {
    return <p>error....</p>;
  }

  if (isLoading) {
    return <Loader />;
  }

  const { tasks = [] } = taskData || {};

  const filteredTasks = tasks?.filter(
    (task) =>
      task.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      task.description.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <>
      <Card>
        <CardContent className="px-5">
          {filteredTasks.length > 0 ? (
            <Table>
              <TableHeader>
                <TableRow className="border-b">
                  <TableHead className="font-semibold text-foreground">
                    Task
                  </TableHead>
                  <TableHead className="font-semibold text-foreground">
                    Status
                  </TableHead>
                  <TableHead className="font-semibold text-foreground">
                    Due Date
                  </TableHead>
                  <TableHead className="w-12"></TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredTasks.map((task) => (
                  <TableRow key={task._id} className="border-b last:border-0">
                    <TableCell>
                      <div>
                        <p className="font-medium text-foreground">
                          {task.title}
                        </p>
                        {task.description && (
                          <p className="text-sm text-muted-foreground mt-1">
                            {task.description}
                          </p>
                        )}
                      </div>
                    </TableCell>
                    <TableCell>
                      {getStatusBadge(task.status, task.dueDate)}
                    </TableCell>
                    <TableCell>
                      <span className="text-sm text-muted-foreground">
                        {formatDate(task.dueDate)}
                      </span>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          ) : (
            <div className="p-12 text-center">
              <div className="text-muted-foreground mb-4">
                <FileText className="h-12 w-12 mx-auto" />
              </div>
              <h3 className="text-lg font-medium text-foreground mb-2">
                No tasks found
              </h3>
              <p className="text-muted-foreground mb-4">
                {searchTerm
                  ? "Try adjusting your search terms"
                  : "Get started by adding your first task to this project"}
              </p>
              <Button onClick={() => setShowCreateTask(true)}>
                <Plus className="h-4 w-4 mr-2" />
                Add Task
              </Button>
            </div>
          )}
        </CardContent>
      </Card>
      <CreateTaskDialog
        open={showCreateTask}
        onOpenChange={setShowCreateTask}
      />
    </>
  );
}

const Loader = () => (
  <Card>
    <CardContent>
      <div className="p-6 space-y-4">
        {[...Array<number>(3)].map((_, i) => (
          <div key={i} className="h-12 bg-muted rounded animate-pulse" />
        ))}
      </div>
    </CardContent>
  </Card>
);
