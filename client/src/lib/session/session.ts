"use server";

import { cookies } from "next/headers";

export async function getSession() {
  const cookie = await cookies();

  const session = cookie.get("zendo_session_token")?.value;

  if (!session?.trim()) {
    return { isAuth: false, message: "Unauthenticated" };
  }

  return {
    isAuth: true,
    message: "Authenticated",
  };
}

export async function createRedirectUrlCookie(url: string) {
  const cookie = await cookies();

  cookie.set("zendo_redirect_url", encodeURIComponent(url));
}

export async function signOut() {
  const cookie = await cookies();
  cookie.delete("zendo_session_token");
  cookie.delete("zendo_access_token");
}
