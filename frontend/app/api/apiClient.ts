import axios from 'axios';
import {AuthResponse, LoginRequest, RegisterRequest, User} from './types';

// Create an axios instance with default configuration
const apiClient = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

// Add response interceptor for handling responses and errors
apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.error('API Error:', error.response?.data || error.message);
    return Promise.reject(error);
  }
);

// Character API functions
export const characterApi = {
  create: async (characterData: any) => {
    const response = await apiClient.post('/characters', characterData);
    return response.data;
  },
  search: async (query: string, top: number = 2) => {
    const response = await apiClient.get('/characters/search', {
      params: {query, top},
    });
    return response.data;
  },
  getById: async (id: string) => {
    const response = await apiClient.get(`/characters/${id}`);
    return response.data;
  },
  generatePdf: async (id: string) => {
    const response = await apiClient.get(`/pdf/character/${id}`, {
      responseType: 'blob',
    });
    return response.data;
  },
};

// World API functions
export const worldApi = {
  create: async (worldData: any) => {
    const response = await apiClient.post('/worlds', worldData);
    return response.data;
  },
  search: async (query: string, top: number = 2) => {
    const response = await apiClient.get('/worlds/search', {
      params: {query, top},
    });
    return response.data;
  },
  getById: async (id: string) => {
    const response = await apiClient.get(`/worlds/${id}`);
    return response.data;
  },
  generatePdf: async (id: string) => {
    const response = await apiClient.get(`/pdf/world/${id}`, {
      responseType: 'blob',
    });
    return response.data;
  },
};

// Campaign API functions
export const campaignApi = {
  create: async (campaignData: any) => {
    const response = await apiClient.post('/campaigns', campaignData);
    return response.data;
  },
  getById: async (id: string) => {
    const response = await apiClient.get(`/campaigns/${id}`);
    return response.data;
  },
  getAll: async () => {
    const response = await apiClient.get('/campaigns');
    return response.data;
  },
};

// Asset API functions
export const assetApi = {
  // TODO: provide asset CDN URL
  getCharacterImage: (filename: string): string => {
    return ``;
  },
  getWorldImage: (filename: string): string => {
    return ``;
  },
};

// Authentication API functions
export const authApi = {
  login: async (credentials: LoginRequest): Promise<AuthResponse> => {
    const response = await apiClient.post('/auth/login', credentials);
    if (response.data.user) {
      sessionStorage.setItem('currentUser', JSON.stringify(response.data.user));
    }
    return response.data;
  },

  logout: async (): Promise<{ message: string }> => {
    const response = await apiClient.post('/auth/logout');
    sessionStorage.removeItem('currentUser');
    return response.data;
  },

  register: async (userData: RegisterRequest): Promise<AuthResponse> => {
    const response = await apiClient.post('/auth/register', userData);
    if (response.data.user) {
      sessionStorage.setItem('currentUser', JSON.stringify(response.data.user));
    }
    return response.data;
  },

  getCurrentUser: async (): Promise<User> => {
    const cachedUser = sessionStorage.getItem('currentUser');
    if (cachedUser) {
      try {
        return JSON.parse(cachedUser);
      } catch (error) {
        sessionStorage.removeItem('currentUser');
      }
    }

    const response = await apiClient.get('/auth/me');
    sessionStorage.setItem('currentUser', JSON.stringify(response.data));
    return response.data;
  },

  isAuthenticated: async (): Promise<boolean> => {
    try {
      await authApi.getCurrentUser();
      return true;
    } catch (error) {
      return false;
    }
  },
};

export default apiClient;
