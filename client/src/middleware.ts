import { NextRequest, NextResponse } from "next/server";
import { getSession } from "./lib/session/session";

const protectedRoute = ["/dashboard"];

export async function middleware(req: NextRequest) {
  const { isAuth } = await getSession();

  const isProtectedRoute = protectedRoute.some((route) =>
    req.nextUrl.pathname.startsWith(route)
  );

  if (!isAuth && isProtectedRoute) {
    return NextResponse.redirect(new URL("/auth/signin", req.nextUrl));
  }
}

export const config = {
  matcher: ["/dashboard/:path*"],
};
