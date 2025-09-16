"use client";

import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import { useTeams } from "@/hooks/use-teams";
import { getErrorMesage } from "@/lib/utils";
import { useState } from "react";
import PaginationBtn from "../../_components/pagination-btn";
import TeamCard from "./team-card";

export default function TeamsSection() {
  const [page, setPage] = useState(1);
  const { data, isLoading, error } = useTeams(10, page);

  if (error) {
    return (
      <ErrorDisplay
        message={getErrorMesage(error)}
        title="Failed to fetch teams"
      />
    );
  }

  if (isLoading) {
    return <Loader text="Fetching Teams..." />;
  }

  const { teams } = data!;

  return (
    <>
      <h1 className="font-semibold py-5">Teams You Collaborate With</h1>
      <div className="grid gap-5 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-5">
        {teams.map((team) => (
          <TeamCard key={team._id} team={team} />
        ))}
      </div>
      <PaginationBtn page={page} setPage={setPage} dataLength={teams.length} />
    </>
  );
}
