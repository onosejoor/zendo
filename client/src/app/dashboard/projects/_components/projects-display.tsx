"use client";

import { useCallback, useEffect, useState } from "react";
import ProjectCard from "./ProjectCards";
import { searchForProjects } from "@/lib/actions/projects";
import { toast } from "sonner";
import { getErrorMesage } from "@/lib/utils";

type Props = {
  initialProjects: IProject[];
  searchTerm: string;
};

export default function ProjectsDisplay({
  initialProjects,
  searchTerm,
}: Props) {
  const [remoteProjects, setRemoteProjects] = useState<IProject[] | null>(null);

  const getFilteredProjects = useCallback(() => {
    if (!searchTerm.trim()) return initialProjects;
    return initialProjects.filter(
      (project) =>
        project.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        (project.description &&
          project.description.toLowerCase().includes(searchTerm.toLowerCase()))
    );
  }, [searchTerm, initialProjects]);

  useEffect(() => {
    if (!searchTerm.trim()) {
      setRemoteProjects(null);
      return;
    }

    const filtered = getFilteredProjects();
    if (filtered.length > 0) {
      setRemoteProjects(null);
      return;
    }

    const fetchProjects = async () => {
      try {
        const res = await searchForProjects(searchTerm);
        if (res.success && res.projects) {
          setRemoteProjects(res.projects);
        } else {
          setRemoteProjects([]);
        }
      } catch (error) {
        toast.error(getErrorMesage(error));
        setRemoteProjects([]);
      }
    };
    fetchProjects();
  }, [searchTerm, initialProjects, getFilteredProjects]);

  const displayProjects = remoteProjects
    ? remoteProjects
    : getFilteredProjects();

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {displayProjects.map((project) => (
        <ProjectCard key={project._id} project={project} />
      ))}
    </div>
  );
}
