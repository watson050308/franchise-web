import { ref, computed } from 'vue';
import { authApi } from '@/api/modules/auth';
import type { LoginForm, RegisterForm } from '@/types/auth';
import type { User } from '@/types/auth';

export function useAuth() {
  const getStoredUser = (): User | null => {
    const storedUser = localStorage.getItem('user');
    return storedUser ? JSON.parse(storedUser) : null;
  };


  const user = ref<User | null>(getStoredUser());
  const isLogin = computed(() => !!user.value);
  const error = ref<string | null>(null);

  const login = async (form: LoginForm) => {
    error.value = null;
    
    try {
      const response = await authApi.login(form);
      localStorage.setItem('user', JSON.stringify(response.data));
      user.value = response.data;
      console.log(user.value);
      console.log(isLogin.value);
      localStorage.setItem('token', response.token);
      return response;
    } catch (err) {
      error.value = '登入失敗';
      throw err;
    }
  };

  const register = async (form: RegisterForm) => {
    error.value = null;
    
    try {
      const response = await authApi.register(form);
      return response;
    } catch (err) {
      error.value = '註冊失敗';
      throw err;
    }
  };

  const logout = async () => {
    try {
      await authApi.logout();
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      user.value = null;
    } catch (error) {
      console.error('Logout failed:', error);
      throw error;
    }
  };

  const checkAuth = async () => {
    const token = localStorage.getItem('token');
    const storedUser = localStorage.getItem('user');
    
    if (token && storedUser) {
      user.value = JSON.parse(storedUser);
    }
  };

  checkAuth();

  return {
    user,
    isLogin,
    error,
    login,
    register,
    logout,
    checkAuth
  };
}