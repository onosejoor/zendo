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
  username: string
};

export default function RemoveMemberDialog({ id, teamId, username }: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleAction = async () => {
    setLoading(true);
    await handleRemoveMember(id, teamId);
    setLoading(false);

    setOpenDialog(false);
  };

  return (
    <AlertDialog open={openDialog} onOpenChange={setOpenDialog}>
      <AlertDialogTrigger asChild>
        <Button variant={"destructive"} className="flex items-center">
          <DeleteIcon /> {loading ? "Removing..." : "Remove Member"}
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>{`Are you sure you want to remove ${username} from this team?`}</AlertDialogDescription>
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
              <Trash className="size-5" /> {loading ? "Removing..." : "Remove"}
            </Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
