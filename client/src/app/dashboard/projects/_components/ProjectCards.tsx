import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { formatDate } from "../../tasks/_components/constants";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { Calendar, Edit, MoreHorizontal } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import Link from "next/link";
import DeleteDataDialog from "@/components/dialogs/delete-data-dialog";

type Props = {
  project: IProject;
  handleEditProject: (project: IProject) => void;
};

export default function ProjectCard({ project, handleEditProject }: Props) {
  return (
    <Link href={`/dashboard/projects/${project._id}`}>
      <Card className="hover:shadow-md transition-shadow h-full">
        <CardHeader className="pb-3">
          <div className="flex items-center justify-between">
            <CardTitle className="text-lg">{project.name}</CardTitle>
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
                <DropdownMenuItem onClick={() => handleEditProject(project)}>
                  <Edit className="h-4 w-4 mr-2" />
                  Edit
                </DropdownMenuItem>
                <DropdownMenuItem asChild>
                  <DeleteDataDialog card type="project" id={project._id} />
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
          <p className="text-sm text-gray-600">
            {project.description || "No description"}
          </p>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <Badge variant="default">Active</Badge>
              <div className="text-sm text-gray-500">
                {project.totalTasks} tasks
              </div>
            </div>

            <div className="flex items-center justify-between text-xs text-gray-500">
              <div className="flex items-center space-x-1">
                <Calendar className="h-3 w-3" />
                <span>Created {formatDate(project.created_at)}</span>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}
