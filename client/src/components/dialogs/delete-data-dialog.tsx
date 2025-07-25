import { Trash, Trash2, X } from "lucide-react";
import { Button } from "../ui/button";

import { handleDeleteProject } from "@/lib/actions/projects";
import { useState } from "react";
import { handleDeleteTask } from "@/lib/actions/tasks";
import { cn } from "@/lib/utils";
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
} from "../ui/alert-dialog";

type Props = {
  id: string;
  type: "task" | "project";
  card?: boolean;
};

const typePrompts = {
  task: {
    message: "You are about to delete this task",
    action: handleDeleteTask,
  },
  project: {
    message:
      "Are you sure you want to delete this project and all tasks in it?",
    action: handleDeleteProject,
  },
};

export default function DeleteDataDialog({ id, type, card = false }: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const [loading, setLoading] = useState(false);

  const prompt = typePrompts[type];

  const handleAction = async () => {
    setLoading(true);
    await prompt.action(id);
    setLoading(false);

    setOpenDialog(false);
  };

  return (
    <AlertDialog open={openDialog} onOpenChange={setOpenDialog}>
      <AlertDialogTrigger asChild>
        <Button
          onClick={() => setOpenDialog(!openDialog)}
          className={cn(
            "text-red-600 justify-start !px-2",
            card ? "w-full" : "w-fit"
          )}
          variant={card ? "ghost" : "outline"}
        >
          <Trash2 className="h-4 w-4 " />
          {loading ? "Deleting... " : "Delete"}
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>{prompt.message}</AlertDialogDescription>
        </AlertDialogHeader>

        <AlertDialogFooter className="flex !justify-start gap-5 items-center">
          <AlertDialogCancel asChild className="flex space-x-2 items-center">
            <Button variant={"outline"} disabled={loading}>
              <X className="h-4 w-4 " />
              Cancel
            </Button>
          </AlertDialogCancel>

          <AlertDialogAction asChild>
            <Button
              onClick={handleAction}
              disabled={loading}
              className="flex space-x-2 bg-red-500 items-center"
              variant={"destructive"}
            >
              <Trash className="size-5" /> {loading ? "Deleting..." : "Delete"}
            </Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
