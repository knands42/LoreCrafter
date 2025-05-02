import axios from 'axios';

// Create an axios instance with default configuration
const apiClient = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add response interceptor for error handling
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    // Handle errors globally
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

export default apiClient;
