"use client";

import { Card, CardContent } from "@/components/ui/card";

import { useTeam } from "@/hooks/use-teams";
import ErrorDisplay from "@/components/error-display";
import { getErrorMesage } from "@/lib/utils";
import Loader from "@/components/loader-card";

export default function TeamContainer({ teamId }: { teamId: string }) {
  const { data, error, isLoading } = useTeam(teamId);

  if (error) {
    if (error.code === 404) {
      return (
        <ErrorDisplay message={getErrorMesage(error)} title="Team Not Found" />
      );
    }
    return (
      <ErrorDisplay
        message={getErrorMesage(error)}
        title="Failed to fetch team"
      />
    );
  }

  if (isLoading) {
    return <Loader text="Fetching Team..." />;
  }

  const { team } = data!;

  const description =
    team.description.length > 250
      ? team.description.slice(0, 250) + "..."
      : team.description;

  return (
    <>
      <Card className="relative shadow-none border-none">
        <CardContent>
          <div className={"grid grid-cols-1 md:grid-cols-2 gap-6"}>
            <div className="space-y-4">
              <div className="flex items-center space-x-3">
                <p className="text-sm font-medium text-muted-foreground">
                  Name:
                </p>
                <p className="text-foreground line-clamp-1 ">{team.name}</p>
              </div>

              <div className="flex items-center space-x-3">
                <p className="text-sm font-medium text-muted-foreground">
                  Description:
                </p>
                <p className="text-foreground">
                  {description || "No description provided"}
                </p>
              </div>

              <div className="flex items-start space-x-3">
                <p className="text-sm font-medium text-muted-foreground">
                  No Of Members:
                </p>
                <div className="flex">
                  {[...Array(team.members_count)].slice(0, 3).map((_, idx) => {
                    return (
                      <div
                        className="size-5 rounded-full bg-gray-400 border border-white -mr-1.5"
                        key={idx}
                      ></div>
                    );
                  })}
                </div>
                <small className="text-gray-400">
                  {team.members_count} Members In this Team
                </small>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </>
  );
}
