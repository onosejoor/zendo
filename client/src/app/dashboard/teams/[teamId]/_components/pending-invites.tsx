import { formatDate } from "@/app/dashboard/tasks/_components/constants";
import RemoveInviteDialog from "@/components/dialogs/remove-invite-dialog";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { useTeamInvites } from "@/hooks/use-teams";
import { getErrorMesage } from "@/lib/utils";
import { DialogDescription } from "@radix-ui/react-dialog";
import { SendToBack } from "lucide-react";

export default function PendingInvites({ teamId }: { teamId: string }) {
  const { data, error, isLoading } = useTeamInvites(teamId);

  const { invitees = [] } = data || {};

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          variant={"default"}
          className="flex gap-3 items-center my-3 bg-blue-500 text-white"
        >
          <SendToBack /> Pending Invites
        </Button>
      </DialogTrigger>

      <DialogContent className="ring-2 ring-blue-400 lg:w-4xl max-h-[90vh] overflow-auto ">
        <DialogHeader>
          <DialogTitle>Pending Invites</DialogTitle>
          <DialogDescription className="text-muted-foreground">
            View your team pending invites
          </DialogDescription>
        </DialogHeader>
        {error ? (
          getErrorMesage(error)
        ) : isLoading ? (
          "loading..."
        ) : (
          <ul className="list-inside list-disc">
            {invitees.map((invitee) => {
              return (
                <li
                  key={invitee._id}
                  className="flex border-l-2 border-blue-100 pl-3 justify-between"
                >
                  <div className="flex not-last:pb-5 gap-3 items-center">
                    <div className="space-y-2">
                      <h4 className="text-gray-800 truncate font-medium">
                        {invitee.email}
                      </h4>
                      <h4 className="text-gray-500 text-sm">
                        <span className="font-semibold">Email Status: </span>

                        {getStatusBadge(invitee.status)}
                      </h4>
                      <h4 className="text-gray-500 text-sm">
                        <span className="font-semibold">Sent: </span>
                        {formatDate(invitee.createdAt)}
                      </h4>
                    </div>
                  </div>
                  <RemoveInviteDialog
                    email={invitee.email}
                    id={invitee._id}
                    teamId={teamId}
                  />
                </li>
              );
            })}
          </ul>
        )}
      </DialogContent>
    </Dialog>
  );
}

export function getStatusBadge(status: "sent" | "pending" | "failed") {
  let color = "";

  switch (status) {
    case "pending":
      color = "bg-yellow-500";

      break;

    case "failed":
      color = "bg-red-500";
      break;

    case "sent":
      color = "bg-green-500";
      break;
  }

  return (
    <Badge
      variant={"destructive"}
      className={`h-fit capitalize text-white ${color}`}
    >
      {status}
    </Badge>
  );
}
