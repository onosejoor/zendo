"use client";

import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import { useTeams } from "@/hooks/use-teams";
import { getErrorMesage } from "@/lib/utils";
import { useState } from "react";
import PaginationBtn from "../../_components/pagination-btn";
import TeamCard from "./team-card";
import { FileText } from "lucide-react";
import { TeamCrudDialog } from "@/components/dialogs/team-crud-dialog";

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
      {teams.length > 0 ? (
        <>
          <div className="grid gap-5 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-5">
            {teams.map((team) => (
              <TeamCard key={team._id} team={team} />
            ))}
          </div>
          <PaginationBtn
            page={page}
            setPage={setPage}
            dataLength={teams.length}
          />
        </>
      ) : (
        <div className="p-12 text-center">
          <div className="text-muted-foreground mb-4">
            <FileText className="h-12 w-12 mx-auto" />
          </div>
          <h3 className="text-lg font-medium text-foreground mb-2">
            No teams yet
          </h3>
          <p className="text-muted-foreground mb-4">
            Get started by adding your first team
          </p>
          <TeamCrudDialog />
        </div>
      )}
    </>
  );
}
