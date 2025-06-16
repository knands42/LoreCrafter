import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";

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

export const fetcher = (url: string) =>
    apiClient.get(url).then((res) => res.data);
export const mutator = (url: string, data: any) =>
    apiClient.post(url, data).then((res) => res.data);

export default apiClient;
