import {apiAuthPrefix, authRoutes, DEFAULT_LOGIN_REDIRECT, publicRoutes} from "@/routes";
import {NextRequest} from "next/server";
import {getCurrentUser} from "@/actions/userAction";

export async function middleware(request: NextRequest) {
    const {nextUrl} = request
    const isLoggedIn = await getCurrentUser()

    const isApiAuthRoute = nextUrl.pathname.startsWith(apiAuthPrefix)
    const isPublicRoute = publicRoutes.includes(nextUrl.pathname)
    const isAuthRoutes = authRoutes.includes(nextUrl.pathname)

    if (isApiAuthRoute) return null
    if (isAuthRoutes) {
        if (isLoggedIn) return Response.redirect(new URL(DEFAULT_LOGIN_REDIRECT, nextUrl))
        return null
    }
    if (!isLoggedIn && !isPublicRoute) return Response.redirect(new URL('/auth/login', nextUrl))
    return null
}

export const config = {
    matcher: [
        // Skip Next.js internals and all static files, unless found in search params
        '/((?!_next|[^?]*\\.(?:html?|css|js(?!on)|jpe?g|webp|png|gif|svg|ttf|woff2?|ico|csv|docx?|xlsx?|zip|webmanifest)).*)',
        // Always run for API routes
        '/(api|trpc)(.*)',
    ],
};