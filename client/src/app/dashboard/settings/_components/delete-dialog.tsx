import { Trash, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";

import { useState } from "react";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { deleteAllProjects, deleteAllTasks } from "@/lib/actions/user";

type Props = {
  type: "tasks" | "projects";
};

const typePrompts = {
  tasks: {
    message:
      "This action cannot be undone. This will permanently delete  all your tasks and remove all associated data from our  servers.",
    action: deleteAllTasks,
  },
  projects: {
    message:
      "This action cannot be undone. This will permanently delete  all your projects and all tasks within those projects from  our servers.",
    action: deleteAllProjects,
  },
};

export default function DeleteAllDataDialog({ type }: Props) {
  const [loading, setLoading] = useState(false);

  const prompt = typePrompts[type];

  const handleAction = async () => {
    setLoading(true);
    await prompt.action();
    setLoading(false);
  };

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="destructive" size="sm" disabled={loading}>
          <Trash2 className="h-4 w-4 mr-2" />
          Delete {type}
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>{prompt.message}</AlertDialogDescription>
        </AlertDialogHeader>

        <AlertDialogFooter className="flex !justify-start gap-5 items-center">
          <AlertDialogCancel
            disabled={loading}
            className="flex space-x-2 items-center"
          >
            Cancel
          </AlertDialogCancel>

          <AlertDialogAction asChild>
            <Button
              onClick={handleAction}
              disabled={loading}
              className="flex space-x-2 bg-red-500 items-center"
              variant={"destructive"}
            >
              <Trash className="size-5" />{" "}
              {loading ? "Deleting..." : `Yes, delete all ${type}`}
            </Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
