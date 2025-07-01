import { use } from "react"
import ProfilePage from "./_components/profile-page"
import { getUser } from "@/actions/getUser"

export default function Profile() {
  const user = use(getUser())

  if (!user.success) {
    return new Error("Error fetching user")
  }

  

  return <ProfilePage user={user.user!} />
}
