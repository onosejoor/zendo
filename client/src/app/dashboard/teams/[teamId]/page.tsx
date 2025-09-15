import { Suspense } from "react";
import TeamContainer from "./_components/main";
import TabBtns from "./_components/tab-btns";

const TeamMembers = dynamic(() => import("./_components/team-members"));
const TeamTasksSection = dynamic(() => import("./_components/team-tasks"));

import dynamic from "next/dynamic";
import Loader from "@/components/loader-card";
import TeamStats from "./_components/team-stats";
import { getSession } from "@/lib/session/session";

type Props = {
  params: Promise<{ teamId: string }>;
  searchParams: Promise<{ section: string }>;
};

export default async function DynamicTeam({ params, searchParams }: Props) {
  const teamId = (await params).teamId;
  const queryParams = (await searchParams).section;
  const userId = (await getSession()).data?.id;

  return (
    <div className="space-y-7.5">
      <TeamContainer teamId={teamId} />
      <TabBtns section={queryParams} />
      <Suspense fallback={<Loader text="Loading..." />}>
        {returnCompOnQuery(queryParams, teamId, userId!)}
      </Suspense>
    </div>
  );
}

function returnCompOnQuery(query: string, teamId: string, userId: string) {
  switch (query) {
    case "tasks":
      return <TeamTasksSection teamId={teamId} />;

    case "members":
      return <TeamMembers userId={userId} teamId={teamId} />;
    case "overview":
      return <TeamStats teamId={teamId} />;
    default:
      return <TeamMembers userId={userId} teamId={teamId} />;
  }
}
