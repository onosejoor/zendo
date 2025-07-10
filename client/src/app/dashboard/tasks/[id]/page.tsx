import TaskContainer from "./_components/main";

type Props = {
  params: Promise<{ id: string }>;
};

export default async function TaskDynamicPage({ params }: Props) {
  const taskId = (await params).id;

  return <TaskContainer taskId={taskId} />;
}
