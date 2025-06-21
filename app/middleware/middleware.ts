import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const protectedRoutes = ["/dashboard", "/profile", "/settings"];

export async function redirectIfNotSignedIn(request: NextRequest) {
  const url = request.nextUrl.clone();
  if (protectedRoutes.some((path) => url.pathname.startsWith(path))) {
    try {
      const verifyResponse = await fetch(
        `${request.nextUrl.origin}/verify-token`,
        {
          method: "GET",
          headers: {
            cookie: request.headers.get("cookie") || "",
          },
        }
      );

      const data = await verifyResponse.json();

      if (!verifyResponse.ok || !data.isAuthenticated) {
        return NextResponse.redirect(url);
      }

      return NextResponse.next();
    } catch (err) {
      console.error("Middleware auth check failed:", err);
      url.pathname = "/signin";
      return NextResponse.redirect(url);
    }
  }

  return NextResponse.next();
}
