import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { fetcher } from "@/lib/utils";
import { CheckSquare, Clock, Users, TrendingUp } from "lucide-react";
import useSWR from "swr";

export default function StatCards() {
  const { data, isLoading, error } = useSWR<{
    stats: IStats;
    success: boolean;
  }>("/stats", fetcher);

  if (error) {
    return <p>error...</p>;
  }

  if (isLoading) {
    return <p>loadin...</p>;
  }

  const { stats } = data!;

  return (
    <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">Total Tasks</CardTitle>
          <CheckSquare className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{stats.total_tasks}</div>
          <p className="text-xs text-muted-foreground">
            {stats.completed_tasks} completed
          </p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">Active Projects</CardTitle>
          <Users className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{stats.total_projects}</div>
          <p className="text-xs text-muted-foreground">Across all workspaces</p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">Completion Rate</CardTitle>
          <TrendingUp className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">
            {Math.round(stats.completion_rate)}%
          </div>
          <Progress value={stats.completion_rate} className="mt-2" />
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">Due Today</CardTitle>
          <Clock className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{stats.dueToday}</div>
          <p className="text-xs text-muted-foreground">Tasks need attention</p>
        </CardContent>
      </Card>
    </div>
  );
}
