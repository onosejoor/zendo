"use server";

import { cookies } from "next/headers";
import { decodeJwt } from "./decodeJwt";

export async function getSession() {
  const cookie = await cookies();

  const session = cookie.get("auth_session_token")?.value;

  if (!session) {
    return { isAuth: false, message: "Unauthenticated" };
  }
  const decodedData =  decodeJwt(session)

  return {isAuth: true, message: "Authenticated", id: decodedData.payload.id as string}

}
