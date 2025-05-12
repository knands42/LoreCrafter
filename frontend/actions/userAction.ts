import {verifyToken} from "@/lib/verifyToken";
import {cookies} from "next/headers";

export const getCurrentUser = async () => {
    const token = (await cookies()).get('auth_token')
    
    return token ? verifyToken(token.value) : null
}