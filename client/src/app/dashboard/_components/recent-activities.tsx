import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Link from "next/link";
import { getStatusColor } from "../tasks/_components/constants";
import { getStatusBadge } from "@/lib/functions";
import { useTasks } from "@/hooks/use-tasks";
import { useProjects } from "@/hooks/use-projects";
import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";
import { cn } from "@/lib/utils";

export default function RecentActivities() {
  const {
    data: taskData,
    isLoading: tasksLoading,
    error: taskError,
  } = useTasks();

  const {
    data: projectData,
    isLoading: projectsLoading,
    error: projectError,
  } = useProjects();

  if (taskError || projectError) {
    return (
      <ErrorDisplay
        message="Error getting projects and tasks"
        title={"Error fetching data"}
      />
    );
  }

  if (tasksLoading || projectsLoading) {
    return <Loader text="Loading data..." />;
  }

  const { tasks } = taskData!;
  const { projects } = projectData!;

  const recentTasks = tasks.slice(0, 5);
  const recentProjects = projects.slice(0, 2);
  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <Card>
        <CardHeader>
          <CardTitle>Recent Tasks</CardTitle>
          <CardDescription>Your latest task activity</CardDescription>
        </CardHeader>
        <CardContent>
          {recentTasks.length > 0 ? (
            <div className="space-y-3">
              {recentTasks.map((task) => (
                <TaskCard key={task._id} task={task} />
              ))}
            </div>
          ) : (
            <EmptyText text="task" />
          )}
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Active Projects</CardTitle>
          <CardDescription>Your current project overview</CardDescription>
        </CardHeader>
        <CardContent>
          {recentProjects.length > 0 ? (
            <div className="space-y-3">
              {recentProjects.map((project) => (
                <ProjectCard key={project._id} project={project} />
              ))}
            </div>
          ) : (
            <EmptyText text="project" />
          )}
        </CardContent>
      </Card>
    </div>
  );
}

const TaskCard = ({ task }: { task: ITask }) => {
  const statusColor = getStatusColor(task.status, task.dueDate);

  return (
    <Link href={`/dashboard/tasks/${task._id}`} className="block">
      <div className="flex items-baseline justify-between p-3 border rounded-lg">
        <div className="flex items-baseline space-x-3 ">
          <div className={`size-2.5 shrink-0 rounded-full ${statusColor}`} />
          <div>
            <p className="font-medium">{task.title}</p>
            <p
              className={cn(
                "text-sm text-gray-500 line-clamp-2",
                !task.description.trim() && "italic"
              )}
            >
              {task.description || "No description"}
            </p>
          </div>
        </div>
        {getStatusBadge(task.status, task.dueDate)}
      </div>
    </Link>
  );
};

const ProjectCard = ({ project }: { project: IProject }) => (
  <Link className="block" href={`/dashboard/projects/${project._id}`}>
    <div className="p-3 border rounded-lg">
      <div className="flex items-center justify-between mb-2">
        <h4 className="font-medium">{project.name}</h4>
        <Badge variant="outline">Active</Badge>
      </div>
      <p
        className={cn(
          "text-sm text-gray-500 line-clamp-2",
          !project.description?.trim() && "italic"
        )}
      >
        {project.description || "No description"}
      </p>
      <div className="flex items-center justify-between text-xs text-gray-400">
        <span>Total Tasks: {project.totalTasks}</span>
        <span>Progress</span>
      </div>
    </div>
  </Link>
);

const EmptyText = ({ text }: { text: string }) => (
  <p className="text-gray-500 text-center py-8">
    No {`${text}s`} yet. Create your first {text}!
  </p>
);
