import ProjectContainer from "./_components/main";

type Props = {
  params: Promise<{ id: string }>;
};

export default async function ProjectDynamicPage({ params }: Props) {
  const projectId = (await params).id;

  return <ProjectContainer projectId={projectId} />;
}
