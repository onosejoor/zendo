import useSWR from "swr"

interface SubTask {
  title: string
  completed: boolean
}

interface Task {
  _id: string
  title: string
  description: string
  userId: string
  subTasks?: SubTask[]
  projectId?: string
  dueDate: string
  status: string
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
    throw new Error("Failed to fetch tasks")
  }

  return response.json()
}

export function useTasks() {
  return useSWR<Task[]>("http://localhost:8080/task", fetcher, {
    revalidateOnFocus: false,
  })
}

export function useTask(id: string) {
  return useSWR<Task>(id ? `http://localhost:8080/task/${id}` : null, fetcher, {
    revalidateOnFocus: false,
  })
}
