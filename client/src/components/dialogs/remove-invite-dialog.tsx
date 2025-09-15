import { LoaderIcon, Trash, TrashIcon, X } from "lucide-react";
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
import { handleRemoveInvite } from "@/lib/actions/invite";

type Props = {
  teamId: string;
  email: string;
};

const dialogTexts = {
  remove: {
    button: "Remove Invite",
    loading: "Removing...",
    title: "Are you absolutely sure?",
    description: (email: string) =>
      `Are you sure you want to remove ${email} invite?`,
    action: "Remove",
  },
};

export default function RemoveInviteDialog({ teamId, email }: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const [loading, setLoading] = useState(false);

  const texts = dialogTexts.remove;

  const handleAction = async () => {
    setLoading(true);
    await handleRemoveInvite(email, teamId);
    setLoading(false);
    setOpenDialog(false);
  };

  return (
    <AlertDialog open={openDialog} onOpenChange={setOpenDialog}>
      <AlertDialogTrigger asChild>
        <Button variant="destructive" className="flex items-center">
          {loading ? <LoaderIcon className="animate-spin" /> : <TrashIcon />}
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{texts.title}</AlertDialogTitle>
          <AlertDialogDescription>
            {texts.description(email)}
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
