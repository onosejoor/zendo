import { TeamTaskContainer } from "../_components/team-task-container";

interface TeamTaskPageProps {
  params: Promise<{
    taskId: string;
    teamId: string;
  }>;
}

export default async function TeamTaskPage({ params }: TeamTaskPageProps) {
  const taskId = (await params).taskId;
  const teamId = (await params).teamId;
  
  return <TeamTaskContainer taskId={taskId} teamId={teamId} />;
}
