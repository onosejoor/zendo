"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";
import { Plus, Search } from "lucide-react";

const CreateTeamTaskDialog = dynamic(
  () => import("@/components/dialogs/team-task-dialog")
);
const PaginationBtn = dynamic(
  () => import("@/app/dashboard/_components/pagination-btn")
);

const TeamTaskCard = dynamic(
  () => import("../tasks/_components/team-task-card")
);

import Loader from "@/components/loader-card";
import ErrorDisplay from "@/components/error-display";
// import useDebounce from "@/hooks/use-debounce";

import { useTeamTasks } from "@/hooks/use-teams";
import { checkRolesMatch } from "../actions";
import dynamic from "next/dynamic";

export default function TeamTasksSection({ teamId }: { teamId: string }) {
  const [page, setPage] = useState(1);
  const [searchTerm, setSearchTerm] = useState("");

  //   const [filter, setFilter] = useState<Status | "all" | "expired">("all");

  const { data, isLoading, error } = useTeamTasks(teamId, 5, page);
  //   const debouncedSearchTerm = useDebounce(searchTerm, 200);

  if (error) {
    return (
      <ErrorDisplay message="Error Loading team tasks, check internet connection and try again" />
    );
  }

  const {
    data: { tasks = [], role },
  } = data || { data: {} };

  const isAuthorized = checkRolesMatch(role!, ["owner", "admin"]);

  if (!isLoading && tasks.length < 1) {
    return (
      <>
        <Card>
          <CardContent className="p-12 text-center">
            <div className="text-gray-400 mb-4">
              <Plus className="h-12 w-12 mx-auto" />
            </div>
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              No tasks found
            </h3>
            <p className="text-gray-600 mb-4">
              {searchTerm
                ? "Try adjusting your search terms"
                : "Get started by creating your first task"}
            </p>
            {isAuthorized && <CreateTeamTaskDialog defaultTeamId={teamId} />}
          </CardContent>
        </Card>
      </>
    );
  }

  return (
    <>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex justify-end items-center">
          {isAuthorized && <CreateTeamTaskDialog defaultTeamId={teamId} />}
        </div>

        <>
          {/* Search and Filters */}
          <div className="flex space-x-4">
            <div className="relative flex-1">
              <Search className="absolute left-3 top-3 h-4 w-4 text-gray-400" />
              <Input
                placeholder="Search tasks..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="pl-10"
              />
            </div>
          </div>

          {/* Tasks List */}
          {isLoading ? (
            <Loader text="Loading Tasks..." />
          ) : (
            <>
              <div className="grid gap-5 md:grid-cols-2 grid-cols-1">
                {tasks.map((task) => (
                  <TeamTaskCard key={task._id} task={task} userRole={role!} />
                ))}
              </div>
              <PaginationBtn
                page={page}
                setPage={setPage}
                dataLength={tasks!.length}
              />
            </>
          )}
        </>
      </div>
    </>
  );
}
