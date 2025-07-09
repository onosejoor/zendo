import "server-only";

import { jwtVerify } from "jose";

export async function decrypt(session: string) {
  const secretKey = process.env.JWT_SECRET;
  const encodedKey = new TextEncoder().encode(secretKey);

  try {
    const { payload, protectedHeader } = await jwtVerify(session, encodedKey, {
      algorithms: ["HS256"],
    });
    return { header: protectedHeader, payload };
  } catch (error) {
    console.log("Failed to verify session: ", error);
  }
}
