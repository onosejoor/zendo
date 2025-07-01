import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

type Props = {
  tasks: ITask[];
  projects: IProject[];
};

export default function RecentActivities({ tasks, projects }: Props) {
  const recentTasks = tasks.slice(0, 5);
  const recentProjects = projects.slice(0, 3);
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
                <div
                  key={task._id}
                  className="flex items-baseline justify-between p-3 border rounded-lg"
                >
                  <div className="flex items-baseline space-x-3">
                    <div
                      className={`size-2 shrink-0 rounded-full ${
                        task.status === "completed"
                          ? "bg-green-500"
                          : "bg-yellow-500"
                      }`}
                    />
                    <div>
                      <p className="font-medium">{task.title}</p>
                      <p className="text-sm text-gray-500">
                        {task.description}
                      </p>
                    </div>
                  </div>
                  <Badge
                    variant={
                      task.status === "completed" ? "default" : "secondary"
                    }
                  >
                    {task.status}
                  </Badge>
                </div>
              ))}
            </div>
          ) : (
            <p className="text-gray-500 text-center py-8">
              No tasks yet. Create your first task!
            </p>
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
                <div key={project._id} className="p-3 border rounded-lg">
                  <div className="flex items-center justify-between mb-2">
                    <h4 className="font-medium">{project.name}</h4>
                    <Badge variant="outline">Active</Badge>
                  </div>
                  <p className="text-sm text-gray-500 mb-2">
                    {project.description}
                  </p>
                  <div className="flex items-center justify-between text-xs text-gray-400">
                    <span>Total Tasks: {project.totalTasks}</span>
                    <span>Progress</span>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <p className="text-gray-500 text-center py-8">
              No projects yet. Create your first project!
            </p>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
