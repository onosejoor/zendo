import { TeamTaskContainer } from "../_components/team-task-container";
import { getSession } from "@/lib/session/session";


interface TeamTaskPageProps {
  params: Promise<{
    taskId: string;
    teamId: string;
  }>;
}

export default async function TeamTaskPage({ params }: TeamTaskPageProps) {
  const taskId = (await params).taskId;
  const teamId = (await params).teamId;
  const userid = (await getSession()).data?.id;

  
  return <TeamTaskContainer taskId={taskId} teamId={teamId} userId={userid} />;
}
