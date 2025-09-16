import { DeleteIcon, Trash, X } from "lucide-react";
import { Button } from "../ui/button";

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
} from "../ui/alert-dialog";
import { handleRemoveMember } from "@/lib/actions/members";

type Props = {
  id: string;
  teamId: string;
  username: string;
  userId: string;
};

const dialogTexts = {
  remove: {
    button: "Remove Member",
    loading: "Removing...",
    title: "Are you absolutely sure?",
    description: (username: string) =>
      `Are you sure you want to remove ${username} from this team?`,
    action: "Remove",
  },
  leave: {
    button: "Exit Team",
    loading: "Exiting...",
    title: "Exit this team?",
    description: () =>
      "Are you sure you want to leave this team? You will lose access to its tasks and members.",
    action: "Exit",
  },
};

export default function RemoveMemberDialog({
  id,
  teamId,
  username,
  userId,
}: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const [loading, setLoading] = useState(false);

  const mode = userId === id ? "leave" : "remove";
  const texts = dialogTexts[mode];

  const handleAction = async () => {
    setLoading(true);
    await handleRemoveMember(id, teamId);
    setLoading(false);
    setOpenDialog(false);
  };

  return (
    <AlertDialog open={openDialog} onOpenChange={setOpenDialog}>
      <AlertDialogTrigger asChild>
        <Button variant="destructive" className="flex w-fit items-center">
          <DeleteIcon /> {loading ? texts.loading : texts.button}
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{texts.title}</AlertDialogTitle>
          <AlertDialogDescription>
            {texts.description(username)}
          </AlertDialogDescription>
        </AlertDialogHeader>

        <AlertDialogFooter className="flex !justify-start gap-5 items-center">
          <AlertDialogCancel asChild className="flex space-x-2 items-center">
            <Button variant="outline" disabled={loading}>
              <X className="h-4 w-4 " />
              Cancel
            </Button>
          </AlertDialogCancel>

          <AlertDialogAction asChild>
            <Button
              onClick={handleAction}
              disabled={loading}
              className="flex space-x-2 bg-red-500 items-center"
              variant="destructive"
            >
              <Trash className="size-5" />
              {loading ? texts.loading : texts.action}
            </Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
