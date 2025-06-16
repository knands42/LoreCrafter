import axios, {AxiosError, AxiosInstance, AxiosResponse} from "axios";
import {AuthResponse, Campaign, CampaignCreationRequest, LoginRequest, RegisterRequest} from "@/lib/types";

const API_BASE_URL =
    process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8000/api";

const apiClient: AxiosInstance = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },
    timeout: 10000, // 10 seconds
    withCredentials: true, // Enable sending cookies with cross-origin requests
});

// Interceptor to add token dynamically
export const setApiClientToken= (token: string) => {
    apiClient.interceptors.request.use(async (config) => {
        config.headers.Authorization = `Bearer ${token}`;
        return config;
    });
}

// Response interceptor
apiClient.interceptors.response.use(
    (response: AxiosResponse) => {
        return response;
    },
    (error: AxiosError) => {
        const { response, request, message } = error;

        if (response) {
            switch (response.status) {
                case 401:
                    console.error(
                        "Unauthorized: You might need to log in again."
                    );
                    break;
                case 403:
                    console.error("Forbidden: You don't have permission.");
                    break;
                case 404:
                    console.error("Not Found: The resource doesn't exist.");
                    break;
                case 500:
                    console.error(
                        "Server error: Something went wrong on the server."
                    );
                    break;
                default:
                    console.error(
                        `Unhandled error: ${response.status}`,
                        response.data
                    );
            }
        } else if (request) {
            console.error("No response received from server", request);
        } else {
            console.error("Request error", message);
        }

        return Promise.reject(error);
    }
);

export const authApi = {
    register: (data: RegisterRequest): Promise<AxiosResponse<AuthResponse>> => 
        apiClient.post("/api/auth/register", data),

    login: (data: LoginRequest): Promise<AxiosResponse<AuthResponse>> => 
        apiClient.post("/api/auth/login", data),
};

export const campaignApi = {
    getAll: (): Promise<AxiosResponse<Campaign>> => apiClient.get("/api/campaigns"),

    getById: (campaignId: string): Promise<AxiosResponse<Array<Campaign>>> =>
        apiClient.get(`/api/campaigns/${campaignId}`),

    create: (data: CampaignCreationRequest): Promise<AxiosResponse<Campaign>> => apiClient.post("/api/campaigns", data),

    update: () => {},

    delete: () => {},
};

export default apiClient;
