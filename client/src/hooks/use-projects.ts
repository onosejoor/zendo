import useSWR from "swr"

interface Project {
  _id: string
  name: string
  description?: string
  ownerId: string
  totalTasks: number
  created_at: string
}

const fetcher = async (url: string) => {
  const token = localStorage.getItem("token")
  if (!token) throw new Error("No token")

  const response = await fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    credentials: "include",
  })

  if (!response.ok) {
    throw new Error("Failed to fetch projects")
  }

  return response.json()
}

export function useProjects() {
  return useSWR<Project[]>("http://localhost:8080/projects", fetcher, {
    revalidateOnFocus: false,
  })
}

export function useProject(id: string) {
  return useSWR<Project>(id ? `http://localhost:8080/projects/${id}` : null, fetcher, {
    revalidateOnFocus: false,
  })
}

export function useProjectTasks(projectId: string) {
  return useSWR(projectId ? `http://localhost:8080/projects/${projectId}/tasks` : null, fetcher, {
    revalidateOnFocus: false,
  })
}
