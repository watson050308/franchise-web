import { ref, computed } from 'vue';
import { authApi } from '@/api/modules/auth';
import type { LoginForm, RegisterForm } from '@/types/auth';
import type { User } from '@/types/auth';

export function useAuth() {
  const loading = ref(false);
  const error = ref<string | null>(null);

  const login = async (form: LoginForm) => {
    loading.value = true;
    error.value = null;
    
    try {
      const response = await authApi.login(form);
      localStorage.setItem('token', response.token);
      return response;
    } catch (err) {
      error.value = '登入失敗';
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const register = async (form: RegisterForm) => {
    loading.value = true;
    error.value = null;
    
    try {
      const response = await authApi.register(form);
      return response;
    } catch (err) {
      error.value = '註冊失敗';
      throw err;
    } finally {
      loading.value = false;
    }
  };

  return {
    loading,
    error,
    login,
    register
  };
}