"use client"

import useSWR from "swr"

interface User {
  _id: string
  name: string
  email: string
  avatar?: string
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
    throw new Error("Failed to fetch user")
  }

  return response.json()
}

export function useUser() {
  return useSWR<User>("http://localhost:8080/auth/user", fetcher, {
    revalidateOnFocus: false,
    shouldRetryOnError: false,
  })
}
