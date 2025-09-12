"use client";

import BreadCrumbs from "@/components/BreadCrumbs";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { ListChecks, Users } from "lucide-react";
import TeamsSection from "./teams-section";
import { CreateTeamDialog } from "@/components/dialogs/create-team-dialog";

export default function TeamPage() {
  return (
    <>
      <div className="flex items-center justify-between">
        <BreadCrumbs />
        <CreateTeamDialog isVariant />
      </div>

      <div className="grid gap-7 5 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-5">
        <Card className="min-h-40 grid ">
          <CardContent>
            <h2 className="text-2xl font-semibold">201 </h2>
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
            <h2 className="text-2xl font-semibold">10</h2>
          </CardContent>
          <CardFooter className="grid gap-1 text-gray-500">
            <ListChecks className="size-5 text-blue-500" />
            <p className="text-sm">Total Tasks assigned to me</p>
          </CardFooter>
        </Card>
        <Card className="min-h-40 grid ">
          <CardContent>
            <h2 className="text-2xl font-semibold">201 </h2>
          </CardContent>
          <CardFooter className="grid gap-1 text-gray-500">
            <Users className="size-5 text-blue-500" />
            <p className="text-sm">
              Total Teams you&apos;re collaborating with
            </p>
          </CardFooter>
        </Card>
      </div>

      <hr className="border-t-blue-100" />

      <TeamsSection />
    </>
  );
}
