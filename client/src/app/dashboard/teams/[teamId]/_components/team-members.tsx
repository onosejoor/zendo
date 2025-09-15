"use client";

import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import { Card, CardContent } from "@/components/ui/card";
import { useTeamMembers } from "@/hooks/use-teams";
import { getErrorMesage } from "@/lib/utils";
import MemberCard from "./member-card";
import dynamic from "next/dynamic";
import { checkRolesMatch } from "../actions";

const SendMemberInviteDialog = dynamic(
  () => import("@/components/dialogs/create-member-dialog")
);

export default function TeamMembers({
  teamId,
  userId,
}: {
  teamId: string;
  userId: string;
}) {
  const { data, error, isLoading } = useTeamMembers(teamId);

  if (error) {
    return (
      <ErrorDisplay
        message={getErrorMesage(error)}
        title="Failed to fetch team members"
      />
    );
  }

  if (isLoading) {
    return <Loader text="Fetching Teams..." />;
  }

  const {
    data: { members, role },
  } = data!;

  return (
    members.length > 0 && (
      <Card>
        <CardContent className="space-y-3">
          <div className="flex items-center justify-between mb-12.5">
            <div className="flex space-x-3 items-center">
              <div className="size-2.5 rounded-full bg-accent-blue animate-bounce" />
              <h3 className="font-semibold">Members</h3>
            </div>
            {checkRolesMatch(role, ["owner", "admin"]) && (
              <SendMemberInviteDialog teamId={teamId} />
            )}
          </div>

          <div className="space-y-5 divide-y divide-blue-200">
            {members.map((member) => (
              <MemberCard
                key={member._id}
                member={member}
                userId={userId}
                userRole={role}
                teamId={teamId}
              />
            ))}
          </div>
        </CardContent>
      </Card>
    )
  );
}
