"use server";

import { cookies } from "next/headers";

export async function getCookie(cookieName: string) {
  const cookieStore = await cookies();
  return cookieStore.get(cookieName)?.value;
}

export async function getAndDeleteCookie(cookieName: string) {
  const cookieStore = await cookies();

  const returnValue = await getCookie(cookieName);

  if (returnValue?.trim()) {
    setTimeout(() => {
      cookieStore.delete(cookieName);
    }, 200);
  }

  return returnValue;
}
