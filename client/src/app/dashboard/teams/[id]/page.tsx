import { Suspense } from "react";
import TeamContainer from "./_components/main";
import TabBtns from "./_components/tab-btns";

const TeamMembers = dynamic(() => import("./_components/team-members"));
const TeamTasksSection = dynamic(() => import("./_components/team-tasks"));

import dynamic from "next/dynamic";
import Loader from "@/components/loader-card";

type Props = {
  params: Promise<{ id: string }>;
  searchParams: Promise<{ section: string }>;
};

export default async function DynamicTeam({ params, searchParams }: Props) {
  const id = (await params).id;
  const queryParams = (await searchParams).section;

  return (
    <div className="space-y-7.5">
      <TeamContainer teamId={id} />
      <TabBtns section={queryParams} />
      <Suspense fallback={<Loader text="Loading..." />}>
        {returnCompOnQuery(queryParams, id)}
      </Suspense>
    </div>
  );
}

function returnCompOnQuery(query: string, teamId: string) {
  switch (query) {
    case "tasks":
      return <TeamTasksSection teamId={teamId} />;

    case "members":
      return <TeamMembers teamId={teamId} />;

    default:
      return <TeamMembers teamId={teamId} />;
  }
}
