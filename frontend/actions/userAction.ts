"use server"

import {verifyToken} from "@/lib/verifyToken";
import {cookies} from "next/headers";

export const getCurrentUser = async () => {
    const token = (await cookies()).get('auth_token')
    return token ? verifyToken(token.value) : null
}

export const setUserCookies = async (token: string) => {
    (await cookies()).set('auth_token', token, {
        httpOnly: true,
        secure: process.env.NODE_ENV === 'production',
        maxAge: 60 * 60 * 24 * 7, // 1 week
        path: '/',
        sameSite: 'lax',
    });
}

export const clearCookies = async () => {
    (await cookies()).delete('auth_token');
}

export const getAuthToken = async () => {
    const tokenCookie = (await cookies()).get('auth_token');
    return tokenCookie ? tokenCookie.value : null;
}
