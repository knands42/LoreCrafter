import {
    createContext,
    ReactNode,
    useContext,
    useEffect,
    useState,
} from "react";
import { useRouter } from "next/router";
import { User } from "@/lib/types";
import { getCurrentUser } from "@/actions/userAction";

type AuthContextType = {
    user: User | null;
    loading: boolean;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
        const loggedUser = async () => {
            try {
                const tokenUser = await getCurrentUser();
                setUser(tokenUser);
            } catch (error) {
                console.error("Error fetching user data: ", error);
                setUser(null);
                await router.push("/login");
            } finally {
                setLoading(false);
            }
        };

        loggedUser().then((r) => Promise.resolve(r));
    }, []);

    return (
        <AuthContext.Provider value={{ user, loading }}>
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
