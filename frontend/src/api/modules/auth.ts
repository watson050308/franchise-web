import { ApiResponse } from '@/types/auth';
import { LoginForm, RegisterForm, User } from '@/types/auth';
import axiosInstance from '../axios';

export const authApi = {
  login: async (data: LoginForm): Promise<ApiResponse<User>> => {
    try {
      const response = await axiosInstance.post<ApiResponse<User>>('/api/v1/login', data);
      return response.data;
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  },

  register: async (data: RegisterForm): Promise<ApiResponse<User>> => {
    try {
      const response = await axiosInstance.post<ApiResponse<User>>('/api/v1/register', {
        name: data.name,
        email: data.email,
        password: data.password
      });
      return response.data;
    } catch (error) {
      console.error('Register error:', error);    
      throw error;
    }
  },

  logout: async () => {
    try {
      const response = await axiosInstance.delete('/api/v1/logout');
      localStorage.removeItem('token');
      return response.data;
    } catch (error) {
      console.error('Logout error:', error);
      throw error;
    }
  }
};