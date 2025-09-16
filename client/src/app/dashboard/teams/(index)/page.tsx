"use client";

import BreadCrumbs from "@/components/BreadCrumbs";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { ListChecks, Users } from "lucide-react";
import TeamsSection from "./teams-section";
import { TeamCrudDialog } from "@/components/dialogs/team-crud-dialog";
import { useAllTeamStats } from "@/hooks/use-teams";
import ErrorDisplay from "@/components/error-display";
import { getErrorMesage } from "@/lib/utils";
import Loader from "@/components/loader-card";

export default function TeamPage() {
  const { data, error, isLoading } = useAllTeamStats();

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

  const { stat } = data!;
  return (
    <>
      <div className="flex items-center justify-between">
        <BreadCrumbs />
        <TeamCrudDialog isVariant />
      </div>

      <div className="grid gap-7 5 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-5">
        <Card className="min-h-40 grid ">
          <CardContent>
            <h2 className="text-2xl font-semibold">{stat.number_of_teams}</h2>
          </CardContent>
          <CardFooter className="grid gap-1 text-gray-500">
            <Users className="size-5 text-blue-500" />
            <p className="text-sm">
              Total Teams you&apos;re collaborating with
            </p>
          </CardFooter>
        </Card>
        <Card className="min-h-40 grid ">
          <CardContent>
            <h2 className="text-2xl font-semibold">
              {stat.number_of_tasks_assigned_to_me}
            </h2>
          </CardContent>
          <CardFooter className="grid gap-1 text-gray-500">
            <ListChecks className="size-5 text-blue-500" />
            <p className="text-sm">Total Tasks assigned to me</p>
          </CardFooter>
        </Card>
        <Card className="min-h-40 grid ">
          <CardContent>
            <h2 className="text-2xl font-semibold">
              {stat.number_of_tasks_due_today}
            </h2>
          </CardContent>
          <CardFooter className="grid gap-1 text-gray-500">
            <Users className="size-5 text-blue-500" />
            <p className="text-sm">Total Teams Tasks due Today</p>
          </CardFooter>
        </Card>
      </div>

      <hr className="border-t-blue-100" />

      <TeamsSection />
    </>
  );
}
