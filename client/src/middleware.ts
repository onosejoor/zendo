import { NextRequest, NextResponse } from "next/server";
import { getSession } from "./lib/session/session";

const protectedRoute = ["/profile"];

export async function middleware(req: NextRequest) {
  const { isAuth } = await getSession();
  if (!isAuth && protectedRoute.includes(req.nextUrl.pathname)) {
    return NextResponse.redirect(new URL("/auth/signin", req.nextUrl));
  }

}

export const config = {
  matcher: [
    "/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)",
  ],
};
