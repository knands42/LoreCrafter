# Cookie-Based Authentication in LoreCrafter

This document explains how cookie-based authentication is implemented in the LoreCrafter frontend application.

## Overview

The application uses cookies to store user session information. This approach has several advantages:
- Cookies are automatically sent with every request to the same domain
- The server can set HTTP-only cookies that cannot be accessed by JavaScript, enhancing security
- It simplifies authentication management compared to manually handling tokens

## Implementation Details

### Axios Configuration

The main axios instance in `apiClient.ts` is configured to send cookies with every request:

```typescript
const apiClient = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true, // Enable sending cookies with requests
});
```

The `withCredentials: true` option is crucial as it tells axios to include cookies when making requests, even for cross-origin requests.

### Authentication API Functions

The following authentication functions are available in the `authApi` object:

```typescript
// Authentication API functions
export const authApi = {
  login: async (credentials: LoginRequest): Promise<AuthResponse> => {
    const response = await apiClient.post('/auth/login', credentials);
    return response.data;
  },
  logout: async (): Promise<{ message: string }> => {
    const response = await apiClient.post('/auth/logout');
    return response.data;
  },
  register: async (userData: RegisterRequest): Promise<AuthResponse> => {
    const response = await apiClient.post('/auth/register', userData);
    return response.data;
  },
  getCurrentUser: async (): Promise<User> => {
    const response = await apiClient.get('/auth/me');
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
```

### TypeScript Types

The following types are defined in `types/index.ts` to support authentication:

```typescript
export interface User {
  id: string;
  name: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  user: User;
  message?: string;
}
```

## How It Works

1. **Login Process**:
   - User submits login credentials
   - Frontend sends credentials to `/auth/login` endpoint
   - Backend validates credentials and sets a session cookie
   - The cookie is automatically stored by the browser
   - User data is cached in sessionStorage for faster access

2. **Authenticated Requests**:
   - For subsequent requests, the cookie is automatically sent with each request
   - Backend validates the cookie and processes the request if valid

3. **Getting Current User**:
   - First checks sessionStorage for cached user data
   - If found, returns the cached data without making an HTTP request
   - If not found, makes an HTTP request to `/auth/me` and caches the result

4. **Checking Authentication Status**:
   - Use `authApi.isAuthenticated()` to check if the user is currently authenticated
   - This function tries to fetch the current user and returns true if successful
   - Uses the cached user data if available

5. **Logout Process**:
   - Call `authApi.logout()` to log out
   - Backend clears the session cookie
   - Browser removes the cookie
   - Cached user data is removed from sessionStorage

## Usage Examples

### Login

```typescript
import { authApi } from '../app/api/apiClient';

async function handleLogin(email: string, password: string) {
  try {
    const response = await authApi.login({ email, password });
    console.log('Logged in successfully', response.user);
    // Redirect or update UI
  } catch (error) {
    console.error('Login failed', error);
    // Show error message
  }
}
```

### Get Current User

```typescript
import { authApi } from '../app/api/apiClient';

async function fetchCurrentUser() {
  try {
    const user = await authApi.getCurrentUser();
    console.log('Current user:', user);
    // Update UI with user info
  } catch (error) {
    console.error('Not authenticated', error);
    // Redirect to login page
  }
}
```

### Protecting Routes

```typescript
import { authApi } from '../app/api/apiClient';
import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';

function ProtectedPage() {
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    async function checkAuth() {
      const isAuthenticated = await authApi.isAuthenticated();
      if (!isAuthenticated) {
        router.push('/login');
      } else {
        setIsLoading(false);
      }
    }
    checkAuth();
  }, [router]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return <div>Protected content</div>;
}
```

## Client-Side Caching Mechanism

To reduce the number of HTTP requests and improve performance, the application implements a client-side caching mechanism for user information:

1. **How It Works**:
   - User data is cached in `sessionStorage` after successful login or registration
   - The `getCurrentUser` function first checks `sessionStorage` before making an HTTP request
   - If cached data is found, it's returned immediately without an API call
   - If no cached data is found, an HTTP request is made and the result is cached

2. **Benefits**:
   - Reduces the number of HTTP requests to the server
   - Improves application performance and responsiveness
   - Maintains security by still using HTTP-only cookies for authentication
   - Cached data is automatically cleared when the browser session ends

3. **Implementation Details**:
   - User data is stored as a JSON string in `sessionStorage` with the key `'currentUser'`
   - The data is cleared when the user logs out
   - If parsing the cached data fails, it's removed and a new request is made

4. **Security Considerations**:
   - `sessionStorage` is cleared when the browser is closed
   - Only non-sensitive user information should be cached (no passwords or tokens)
   - The authentication itself still relies on HTTP-only cookies

## Backend Requirements

For this cookie-based authentication to work, the backend must:

1. Set appropriate cookies upon successful login
2. Validate cookies on protected endpoints
3. Clear cookies on logout
4. Configure CORS to allow credentials if the frontend and backend are on different domains

The backend should set cookies with these recommended settings:
- `httpOnly: true` - Prevents JavaScript access to the cookie
- `secure: true` - Only sends the cookie over HTTPS (in production)
- `sameSite: 'strict'` or `'lax'` - Provides CSRF protection
