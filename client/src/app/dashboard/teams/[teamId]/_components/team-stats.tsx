"use client";

import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { useTeamStats } from "@/hooks/use-teams";
import { getTeamRoleColor } from "@/lib/functions";
import { getErrorMesage } from "@/lib/utils";
import { Users } from "lucide-react";
import PendingInvites from "./pending-invites";

export default function TeamStats({ teamId }: { teamId: string }) {
  const { data, error, isLoading } = useTeamStats(teamId);

  if (error) {
    return (
      <ErrorDisplay
        message={getErrorMesage(error)}
        title="Failed to fetch team stat"
      />
    );
  }

  if (isLoading) {
    return <Loader text="Fetching Team stat..." />;
  }

  const { stats } = data!;

  const statArray = [
    { title: "Total Tasks in team", value: stats.total_tasks },
    { title: "Total Team Members", value: stats.total_team_members },
    { title: "My Role", value: getTeamRoleColor(stats.role) },
  ];

  return (
    <Card>
      <CardContent className="space-y-3">
        <div className="flex items-center justify-between mb-12.5">
          <div className="flex space-x-3 items-center">
            <div className="size-2.5 rounded-full bg-accent-blue animate-bounce" />
            <h3 className="font-semibold">Team Overview</h3>
          </div>
          {stats.role === "owner" && <PendingInvites teamId={teamId} />}
        </div>

        <div className="grid gap-5 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-5">
          {statArray.map((stat, idx) => (
            <Card className="min-h-40 grid " key={idx}>
              <CardContent>
                <h2 className="text-2xl font-semibold">{stat.value} </h2>
              </CardContent>
              <CardFooter className="grid gap-1 text-gray-500">
                <Users className="size-5 text-blue-500" />
                <p className="text-sm">{stat.title}</p>
              </CardFooter>
            </Card>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
