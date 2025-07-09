"use server";

import { cookies } from "next/headers";

export async function getSession() {
  const cookie = await cookies();

  const session = cookie.get("zendo_session_token")?.value;

  if (!session) {
    return { isAuth: false, message: "Unauthenticated" };
  }

  return {
    isAuth: true,
    message: "Authenticated",
  };
}

export async function signOut() {
  const cookie = await cookies();
  cookie.delete("zendo_session_token");
  cookie.delete("zendo_access_token");
}
