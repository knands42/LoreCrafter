"use client"

import {createContext, ReactNode, useContext, useEffect, useState,} from "react";
import {AuthResponse, LoginRequest, User} from "@/lib/types";
import {clearCookies, getCurrentUser, setUserCookies} from "@/actions/userAction";
import {authApi, setApiClientToken} from "@/lib/apiClient";

type AuthContextType = {
    user: User | null;
    loading: boolean;
    login: (credentials: LoginRequest) => Promise<AuthResponse>;
    logout: () => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const loggedUser = async () => {
            try {
                setLoading(true);
                const tokenUser = await getCurrentUser();
                setUser(tokenUser);
            } catch (error) {
                console.error("Error fetching user data: ", error);
                setUser(null);
            } finally {
                setLoading(false);
            }
        };

        loggedUser().then((r) => Promise.resolve(r));
    }, []);

    const login = async (credentials: LoginRequest): Promise<AuthResponse> => {
        setLoading(true);

        const response = await authApi.login(credentials);
        await setUserCookies(response.data.token)
        setApiClientToken(response.data.token)
        
        setLoading(false);
        
        return response.data
    };

    const logout = async () => {
        setUser(null);
        await clearCookies()
    };

    return (
        <AuthContext.Provider value={{ user, loading, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
}
