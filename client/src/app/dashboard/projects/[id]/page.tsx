import ProjectContainer from "./_components/container";

type Params = {
  params: Promise<{ id: string }>;
};

export default async function ProjectDetailPage({ params }: Params) {
  const projectId = (await params).id;

  return <ProjectContainer projectId={projectId} />;
}
