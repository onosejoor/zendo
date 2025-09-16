"use client";

import TeamTaskMain from "./team-task-main";

// import BreadCrumbs from "@/components/BreadCrumbs";
import ErrorDisplay from "@/components/error-display";
import DeleteDataDialog from "@/components/dialogs/delete-data-dialog";
import { useTeamTask } from "@/hooks/use-teams";
import TeamTaskDialog from "@/components/dialogs/team-task-dialog";
import Loader from "@/components/loader-card";
import { checkRolesMatch } from "../../actions";
import { getErrorMesage } from "@/lib/utils";

type Props = {
  taskId: string;
  teamId: string;
};

export function TeamTaskContainer({ taskId, teamId }: Props) {
  const { data: taskData, isLoading, error } = useTeamTask(teamId, taskId);

  if (error) {
    if (error.status === 404) {
      return (
        <ErrorDisplay
          dontTryAgain
          title="Team task not found"
          message="The team task you're looking for doesn't exist."
        />
      );
    }
    return <ErrorDisplay message={getErrorMesage(error)} />;
  }

  if (isLoading) {
    return <Loader text="Fetching task..." />;
  }

  const {
    data: { task, role },
  } = taskData!;

  if (!task.team_id) {
    return (
      <ErrorDisplay
        dontTryAgain
        title="Not a team task"
        message="This task is not assigned to a team."
      />
    );
  }

  const isOwner = checkRolesMatch(role, ["owner"]);

  return (
    <>
      <div className="max-w-7xl mx-auto space-y-8">
        {/* <BreadCrumbs /> */}
        {/* Header */}
        <div className="flex sm:items-baseline sm:flex-row flex-col gap-5 justify-between">
          <div>
            <h1 className="text-2xl sm:text-3xl font-bold text-foreground">
              Title: <span className="text-muted-foreground">{task.title}</span>
            </h1>
            <p className="text-muted-foreground mt-1">
              <span className="text-foreground"> Description: </span>
              {task.description || "No description"}
            </p>
          </div>
          {isOwner && (
            <div className="space-x-3 flex">
              <TeamTaskDialog defaultTeamId={teamId} initialData={task} />
              {isOwner && <DeleteDataDialog id={taskId} type="team_task" />}
            </div>
          )}
        </div>

        {/* Team Task Details */}
        <TeamTaskMain task={task} />
      </div>
    </>
  );
}
