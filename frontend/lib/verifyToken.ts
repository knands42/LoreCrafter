import {UserToken} from "@/lib/types";

export const verifyToken = (token: string): UserToken => {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (error) {
        throw new Error(`Invalid token: ${error}`);
    }
}