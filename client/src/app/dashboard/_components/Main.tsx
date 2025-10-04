"use client";

import { CreateTaskDialog } from "@/components/dialogs/create-task-dialog";
import { CreateProjectDialog } from "@/components/dialogs/create-project-dialog";
import { useUser } from "@/hooks/use-user";
import RecentActivities from "./recent-activities";
import StatCards from "./Stats";
import Loader from "@/components/loader-card";

export default function Dashboard() {
  const { data, isLoading: userLoading, error } = useUser();

  if (userLoading) {
    return <Loader />;
  }

  const { user } = data || {};

  return (
    <>
      <div className="space-y-8">
        {/* Header */}
        <div className="flex justify-between sm:flex-row flex-col gap-5 sm:items-center">
          <div className="truncate">
            <h1 className="text-3xl font-bold text-gray-900 truncate">
              {error ? (
                "Error getting user profile, refresh to try again."
              ) : (
                <>
                  Welcome back,
                  <span className="text-accent-blue">
                    {user?.username || "E"}!
                  </span>
                </>
              )}
            </h1>
            <p className="text-gray-600 mt-1">
              Here&apos;s what&apos;s happening with your tasks today.
            </p>
          </div>
          <div className="flex *:w-fit gap-2 sm:items-center flex-col sm:flex-row">
            <CreateTaskDialog />
            <CreateProjectDialog />
          </div>
        </div>

        {/* Stats Cards */}
        <StatCards />

        {/* Recent Tasks and Projects */}
        <RecentActivities />
      </div>
    </>
  );
}
