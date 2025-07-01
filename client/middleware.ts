import { NextResponse } from "next/server"
import type { NextRequest } from "next/server"

export function middleware(request: NextRequest) {
  const isDashboard = request.nextUrl.pathname.startsWith("/dashboard")
  const isAuthPage = request.nextUrl.pathname.startsWith("/auth")

  // For dashboard routes, we'll let the client-side handle auth checks
  // since we're using localStorage for token storage
  if (isDashboard || isAuthPage) {
    return NextResponse.next()
  }

  return NextResponse.next()
}

export const config = {
  matcher: ["/dashboard/:path*", "/auth/:path*"],
}
